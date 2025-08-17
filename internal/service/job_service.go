package service

import (
	"time"
	"welltoon/internal/entity"
	"welltoon/internal/repository"

	"github.com/go-co-op/gocron/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type JobService struct {
	logger          *logrus.Logger
	schedule        gocron.Scheduler
	comicRepository repository.ComicRepository
	viewRepository  repository.ViewRepository
	cacheRepository repository.CacheRepository
}

func NewJobService(logger *logrus.Logger,
	schedule gocron.Scheduler,
	comicRepository repository.ComicRepository,
	viewRepository repository.ViewRepository,
	cacheRepository repository.CacheRepository) *JobService {
	return &JobService{
		logger:          logger,
		schedule:        schedule,
		comicRepository: comicRepository,
		viewRepository:  viewRepository,
		cacheRepository: cacheRepository,
	}
}
func (s *JobService) JobView() {
	daily, err := s.schedule.NewJob(
		gocron.DailyJob(1,
			gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0)),
		),
		gocron.NewTask(func() {
			total, err := s.cacheRepository.GetView()
			if err != nil {
				s.logger.WithError(err).Error(err)
				return
			}
			if total != 0 {
				view := new(entity.View)
				data, err := s.viewRepository.FindByComicIDIsNull()
				if err != nil {
					s.logger.WithError(err).Warn("view not found, create new")
					view.Daily = total
					view.Weekly = total
					view.Monthly = total
					view.AllTime = total
				} else {
					view = data
					view.Daily = 0
					view.Weekly += total
					view.Monthly += total
					view.AllTime += total
				}
				if err := s.viewRepository.Save(view); err != nil {
					s.logger.WithError(err).Error("view save failed")
					return
				}
				if err := s.cacheRepository.DelView(); err != nil {
					s.logger.WithError(err).Error(err)
					return
				}
				s.logger.WithField("data", view).Info("view save success")
			}
		}),
	)
	if err != nil {
		s.logger.WithError(err).Error("Job view daily failed")
		return
	}
	s.logger.WithField("data", fiber.Map{
		"job_id": daily.ID(),
	}).Info("job view daily success")

	weekly, err := s.schedule.NewJob(
		gocron.WeeklyJob(1, gocron.NewWeekdays(time.Sunday), gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0))),
		gocron.NewTask(func() {
			view, err := s.viewRepository.FindByComicIDIsNull()
			if err != nil {
				s.logger.WithError(err).Warn("view not found")
				return
			}
			view.Weekly = 0
			if err := s.viewRepository.Save(view); err != nil {
				s.logger.WithError(err).Error("view save failed")
				return
			}
			s.logger.WithField("data", view).Info("view save success")
		}),
	)
	if err != nil {
		s.logger.WithError(err).Error("Job view weekly failed")
		return
	}
	s.logger.WithField("data", fiber.Map{
		"job_id": weekly.ID(),
	}).Info("job view weekly success")

	monthly, err := s.schedule.NewJob(
		gocron.MonthlyJob(1, gocron.NewDaysOfTheMonth(-1), gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0))),
		gocron.NewTask(func() {
			view, err := s.viewRepository.FindByComicIDIsNull()
			if err != nil {
				s.logger.WithError(err).Warn("view not found")
				return
			}
			view.Monthly = 0
			if err := s.viewRepository.Save(view); err != nil {
				s.logger.WithError(err).Error("view save failed")
				return
			}
			s.logger.WithField("data", view).Info("view save success")
		}),
	)
	if err != nil {
		s.logger.WithError(err).Error("Job view monthly failed")
		return
	}
	s.logger.WithField("data", fiber.Map{
		"job_id": monthly.ID(),
	}).Info("job view monthly success")
}
