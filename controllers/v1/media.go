//Package controllers ...
package controllers

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

//TODO resize and exif check

func UploadMedia(request echo.Context) (err error) {
	requestedMediaType := request.Param("mediaType")
	mediaTypes := []string{models.ImageMediaType, models.VideosMediaType, models.FilesMediaType, models.AudiosMediaType}
	isCurrentMediaType, _ := helpers.StringSortAndSearch(mediaTypes, requestedMediaType)
	if !isCurrentMediaType {
		return helpers.ResponseError(request, nil, http.StatusUnprocessableEntity, "CM-1000", "Media Type", "requested media type must be one the "+strings.Join(mediaTypes, ", "))
	}
	file, err := request.FormFile("media")
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1001", "Upload File", "file not uploaded")
	}
	mediaSize := viper.GetInt64("MEDIA." + strings.ToUpper(requestedMediaType) + "_SIZE")
	fileSizeInMegaByte, err := helpers.ConvertByte(file.Size, "MB")
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1003", "Get File Size", "cannot estimate file size")
	}
	if int64(fileSizeInMegaByte) > mediaSize {
		return helpers.ResponseError(request, nil, http.StatusBadRequest, "CM-1004", "File Size", "media size must not be more than "+string(mediaSize))
	}
	sourceFile, err := file.Open()
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1005", "Open File", "cannot open the file")
	}
	defer sourceFile.Close()
	success, mimeTypes := helpers.ValidateFileType(file.Header.Get("content-type"), requestedMediaType)
	if !success {
		return helpers.ResponseError(request, nil, http.StatusBadRequest, "CM-1007", "Check Mime Type", requestedMediaType+" must be one of the "+strings.Join(mimeTypes, ", "))
	}
	authData := helpers.AuthData(request)
	storagePath := helpers.Path("storage")
	//TODO move this string to helper
	filePath := "/temp/" + requestedMediaType + "/" + authData.UserID.Hex() + "/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + uuid.New().String()
	fileName := "/" + file.Filename
	//TODO move this mkdir to helpers
	if _, err := os.Stat(storagePath + filePath); os.IsNotExist(err) {
		if err = os.MkdirAll(storagePath+filePath, 0755); err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1008", "Create Directory", "cannot create directory")
		}
	}
	destination, err := os.Create(storagePath + filePath + fileName)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1009", "Write File", "cannot write in the directory")
	}
	defer destination.Close()
	if _, err := io.Copy(destination, sourceFile); err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1010", "Move File", "cannot move file to the directory")
	}
	_ = destination.Sync()
	imageModel := new(models.Image)
	imageModel.OriginalURL = filePath + fileName
	imageModel.URL = filePath + fileName
	imageModel.Size = models.OriginalSize
	imageModel.Type = file.Header.Get("content-type")
	return helpers.ResponseOk(request, http.StatusOK, imageModel)
}

func GetMedia(request echo.Context) (err error) {
	tempPath := helpers.Path("temp")
	mediaType := request.Param("mediaType")
	userID := request.Param("userID")
	date := request.Param("year") + "/" + request.Param("month") + "/" + request.Param("day")
	uniqueID := request.Param("uniqueID")
	filename := request.Param("filename")
	filePath := tempPath + "/" + mediaType + "/" + userID + "/" + date + "/" + uniqueID + "/" + filename
	return request.Inline(filePath, filename)
}
