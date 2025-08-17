package dto

type ViewResponse struct {
	Daily   int `json:"daily"`
	Weekly  int `json:"weekly"`
	Monthly int `json:"monthly"`
	AllTime int `json:"all_time"`
}
type ViewAddRequest struct {
	ComicID int64 `json:"comic_id" validate:"required"`
}
