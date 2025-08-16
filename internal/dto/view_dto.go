package dto

type ViewResponse struct {
	Daily   int `json:"daily"`
	Weekly  int `json:"weekly"`
	Monthly int `json:"monthly"`
	AllTime int `json:"all_time"`
}
