package converter

import (
	"net/http"
	"server/internal/model"
	"server/pkg/errors"
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseImage(c echo.Context) (*model.Image, error) {
	image := &model.Image{}
	if !strings.HasPrefix(c.Request().Header.Get("Content-Type"), "multipart/form-data") {
		return nil, errors.NewError("invalid content type", errors.BadRequest)
	}
	file, err := c.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return nil, errors.NewError("image is empty", errors.BadRequest)
		}
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileType := file.Header.Get("Content-Type")
	switch fileType {
	case "image/jpeg":
		image.Type = "image/jpeg"
		break
	case "image/png":
		image.Type = "image/png"
		break
	case "image/jpg":
		image.Type = "image/jpg"
		break
	default:
		return nil, errors.NewError("invalid image type", errors.BadRequest)
	}
	image.Data = src
	return image, nil
}
