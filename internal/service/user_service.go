package service

import (
	"mime/multipart"
	"strconv"
	"welltoon/internal/dto"
	"welltoon/internal/entity"
	"welltoon/internal/repository"
	"welltoon/internal/security"
	"welltoon/pkg"
	"welltoon/pkg/enum"
	"welltoon/pkg/image"
	"welltoon/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(request *dto.UserRegisterRequest) error
	LoginUser(request *dto.UserLoginRequest) (*dto.TokenResponse, error)
	UploadAvatar(userID string, avatar *multipart.FileHeader) error
	UpdateUser(userID string, request *dto.UserUpdateRequest) error
	GetUserByID(userID string) (*dto.UserResponse, error)
}
type userService struct {
	validation     *validator.Validate
	logger         *logrus.Logger
	userRepository repository.UserRepository
	s3Repository   repository.S3Repository
}

func NewUserService(validation *validator.Validate, logger *logrus.Logger, userRepository repository.UserRepository, s3Repository repository.S3Repository) UserService {
	return &userService{
		validation:     validation,
		logger:         logger,
		userRepository: userRepository,
		s3Repository:   s3Repository,
	}
}

func (s *userService) RegisterUser(request *dto.UserRegisterRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", fiber.Map{
			"username": request.Username,
		}).Warn("validation failed")
		return err
	}
	countUserUsername, err := s.userRepository.CountByUsername(request.Username)
	if err != nil {
		s.logger.WithError(err).Error("count user by username failed")
		return err
	}
	if countUserUsername > 0 {
		s.logger.WithField("data", fiber.Map{
			"username": request.Username,
		}).Warn("username already exists")
		return response.Exception(400, "username already exists")
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.WithError(err).Error("bcrypt hash password failed")
		return err
	}

	user := &entity.User{
		Username: request.Username,
		Password: string(newPassword),
		Role:     enum.ROLE_USER,
	}

	if err := s.userRepository.Save(user); err != nil {
		s.logger.WithError(err).Error("save user to database failed")
		return err
	}

	s.logger.WithField("data", fiber.Map{
		"username": request.Username,
	}).Info("register user success")
	return nil
}
func (s *userService) LoginUser(request *dto.UserLoginRequest) (*dto.TokenResponse, error) {
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", fiber.Map{
			"username": request.Username,
		}).Warn("validation failed")
		return nil, err
	}

	user, err := s.userRepository.FindByUsername(request.Username)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"username": request.Username,
		}).Warn("username or password wrong")
		return nil, response.Exception(400, "username or password wrong")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"username": request.Username,
		}).Warn("username or password wrong")
		return nil, response.Exception(400, "username or password wrong")
	}
	token, err := security.JwtCreateToken(user.ID, string(user.Role))
	if err != nil {
		s.logger.WithError(err).Error("jwt create token failed")
		return nil, err
	}
	result := &dto.TokenResponse{
		Token: token,
	}
	s.logger.WithField("data", fiber.Map{
		"username": user.Username,
		"role":     user.Role,
	}).Info("login user success")
	return result, nil
}
func (s *userService) UploadAvatar(userID string, avatar *multipart.FileHeader) error {
	newUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"user_id": userID,
		}).Warn("userID most be number")
		return response.Exception(400, "userID most be number")
	}
	user, err := s.userRepository.FindByID(newUserID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"user_id": newUserID,
		}).Warn("user not found")
		return response.Exception(404, "user not found")
	}

	err = image.Validate(avatar.Filename)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"avatar_filename": avatar.Filename,
		}).Warn(err.Error())
		return response.Exception(400, err.Error())
	}

	if user.AvatarFilename != nil && user.AvatarUrl != nil {
		file, err := image.CompressToWebp(avatar)
		if err != nil {
			s.logger.WithError(err).Error("compress image to webp failed")
			return err
		}

		file.Filename = *user.AvatarFilename
		err = s.s3Repository.PutObject(file)
		if err != nil {
			s.logger.WithError(err).Error("s3 put object failed")
			return err
		}
		s.logger.WithField("data", fiber.Map{
			"user_id":    newUserID,
			"avatar_url": user.AvatarUrl,
		}).Info("upload avatar success")
	} else {
		file, err := image.CompressToWebp(avatar)
		if err != nil {
			s.logger.WithError(err).Error("compress image to webp failed")
			return err
		}

		err = s.s3Repository.PutObject(file)
		if err != nil {
			s.logger.WithError(err).Error("s3 put object failed")
			return err
		}
		avatarUrl, err := pkg.S3GenerateUrl(file.Filename)
		if err != nil {
			s.logger.WithError(err).Error("s3 generate url failed")
			return err
		}
		err = s.userRepository.UpdateAvatar(newUserID, file.Filename, avatarUrl)
		if err != nil {
			err := s.s3Repository.RemoveObject(file.Filename)
			if err != nil {
				s.logger.WithError(err).Error("s3 remove object failed")
				return err
			}
			s.logger.WithError(err).Error("update avatar to database failed")
			return err
		}
		s.logger.WithField("data", fiber.Map{
			"user_id":    newUserID,
			"avatar_url": avatarUrl,
		}).Info("upload avatar success")
	}
	return nil
}
func (s *userService) UpdateUser(userID string, request *dto.UserUpdateRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.logger.WithField("data", fiber.Map{
			"username": request.Username,
		}).Warn("validation failed")
		return err
	}
	newUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"user_id": userID,
		}).Warn("userID most be number")
		return response.Exception(400, "userID most be number")
	}
	user, err := s.userRepository.FindByID(newUserID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"user_id": newUserID,
		}).Warn("user not found")
		return response.Exception(404, "user not found")
	}
	if request.Username != nil {
		user.Username = *request.Username
	}
	if request.OldPassword != nil && request.NewPassword != nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*request.OldPassword))
		if err != nil {
			s.logger.WithField("data", fiber.Map{
				"user_id": newUserID,
			}).Warn("password wrong")
			return response.Exception(400, "password wrong")
		}
		newPassword, err := bcrypt.GenerateFromPassword([]byte(*request.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			s.logger.WithError(err).Error("bcrypt hash password failed")
			return err
		}
		user.Password = string(newPassword)
	} else {
		s.logger.WithField("data", fiber.Map{
			"user_id": newUserID,
		}).Warn("old password and new password most be match")
		return response.Exception(400, "old password and new password most be match")
	}
	if err := s.userRepository.Save(user); err != nil {
		s.logger.WithError(err).Error("save user to database failed")
		return err
	}
	s.logger.WithField("data", fiber.Map{
		"user_id": newUserID,
	}).Info("update user success")
	return nil
}
func (s *userService) GetUserByID(userID string) (*dto.UserResponse, error) {
	newUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"user_id": userID,
		}).Warn("userID most be number")
		return nil, response.Exception(400, "userID most be number")
	}
	user, err := s.userRepository.FindByID(newUserID)
	if err != nil {
		s.logger.WithField("data", fiber.Map{
			"user_id": newUserID,
		}).Warn("user not found")
		return nil, response.Exception(404, "user not found")
	}
	result := &dto.UserResponse{
		ID:             user.ID,
		Username:       user.Username,
		Role:           user.Role,
		AvatarFilename: user.AvatarFilename,
		AvatarUrl:      user.AvatarUrl,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
	s.logger.WithField("data", fiber.Map{
		"user_id": newUserID,
	}).Info("get user by id success")
	return result, nil
}
