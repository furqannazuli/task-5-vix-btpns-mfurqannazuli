package photo

import "github.com/furqannazuli/task-5-vix-btpns-mfurqannazuli/models"

type PhotoRegularResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type PhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
	User     models.User
}

func FormatPhoto(photo *models.Photo, typeRes string) interface{} {
	var formatter interface{}

	if typeRes == "regular" {
		formatter = PhotoRegularResponse{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
		}
	} else {
		formatter = PhotoResponse{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			UserID:   photo.User.ID,
			User:     *photo.User,
		}
	}
	return formatter
}
