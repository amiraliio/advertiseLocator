//Package helpers ...
package helpers

import (
	"github.com/amiraliio/advertiselocator/models"
)

func ValidateFileType(fileType, mediaType string) (bool, []string) {
	switch mediaType {
	case models.ImageMediaType:
		return isAcceptableMediaType(models.ImageMimeTypes, fileType)
	case models.VideosMediaType:
		return isAcceptableMediaType(models.VideMimeTypes, fileType)
	case models.AudiosMediaType:
		return isAcceptableMediaType(models.AudioMimeTypes, fileType)
	case models.FilesMediaType:
		return isAcceptableMediaType(models.FileMimeTypes, fileType)
	default:
		return false, nil
	}
}

func isAcceptableMediaType(mimeTypes []string, fileType string) (bool, []string) {
	isAcceptableMediaType, _ := StringSortAndSearch(mimeTypes, fileType)
	if isAcceptableMediaType {
		return true, mimeTypes
	}
	return false, mimeTypes
}
