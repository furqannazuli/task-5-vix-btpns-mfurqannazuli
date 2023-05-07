package controllers

import (
	"net/http"

	photoRes "github.com/furqannazuli/task-5-vix-btpns-mfurqannazuli/app/photo"
	"github.com/furqannazuli/task-5-vix-btpns-mfurqannazuli/helpers"
	"github.com/furqannazuli/task-5-vix-btpns-mfurqannazuli/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type photoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *photoController {
	return &photoController{db}
}

// Get Foto
func (h *photoController) Get(c *gin.Context) {
	var userPhoto models.Photo
	err := h.db.Preload("User").Find(&userPhoto).Error

	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Gagal mendapatkan foto")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if userPhoto.PhotoURL == "" {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Tolong upload foto terlebih dahulu")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photoRes.FormatPhoto(&userPhoto, "")
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "Sukses upload foto")
	c.JSON(http.StatusOK, response)
}

// Upload Foto
func (h *photoController) Create(c *gin.Context) {
	var userPhoto models.Photo
	var countPhoto int64
	currentUser := c.MustGet("currentUser").(models.User)

	h.db.Model(&userPhoto).Where("user_id = ?", currentUser.ID).Count(&countPhoto)
	if countPhoto > 0 {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "Anda sudah memiliki foto")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Gagal untuk upload foto user")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("photo_profile")
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Gagal untuk upload foto user")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Gagal untuk upload foto user")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	h.InsertPhoto(input, path, currentUser.ID)

	data := gin.H{"is_uploaded": true}
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Foto profile berhasil diupload")
	c.JSON(http.StatusOK, response)
}

// Upload Foto
func (h *photoController) InsertPhoto(userPhoto models.Photo, fileLocation string, currUserID int) error {
	savePhoto := models.Photo{
		UserID:   currUserID,
		Title:    userPhoto.Title,
		Caption:  userPhoto.Caption,
		PhotoURL: fileLocation,
	}

	err := h.db.Debug().Create(&savePhoto).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Foto
func (h *photoController) Update(c *gin.Context) {
	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	err := h.db.Where("user_id = ?", currentUser.ID).Find(&userPhoto).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Foto profile gagal diupdate")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo
	err = c.ShouldBind(&input)
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Foto profile gagal diupdate")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("update_profile")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Foto user gagal diupdate")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Foto profile gagal diupdate")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.UpdatePhoto(input, &userPhoto, path)

	data := photoRes.FormatPhoto(&userPhoto, "regular")
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Foto profile berhasil diupdate")
	c.JSON(http.StatusOK, response)
}

// Update Foto
func (h *photoController) UpdatePhoto(oldPhoto models.Photo, newPhoto *models.Photo, path string) error {
	newPhoto.Title = oldPhoto.Title
	newPhoto.Caption = oldPhoto.Caption
	newPhoto.PhotoURL = path

	err := h.db.Save(&newPhoto).Error
	if err != nil {
		return err
	}

	return nil
}

// Hapus Foto
func (h *photoController) Delete(c *gin.Context) {
	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	err := h.db.Where("user_id = ?", currentUser.ID).Delete(&userPhoto).Error
	if err != nil {
		data := gin.H{
			"is_deleted": false,
		}

		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "Gagal untuk menghapus foto")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_deleted": true,
	}

	response := helpers.ApiResponse(http.StatusOK, "success", data, "Foto user berhasil dihapus")
	c.JSON(http.StatusOK, response)
}
