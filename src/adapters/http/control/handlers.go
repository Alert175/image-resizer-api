package control

import (
	"errors"
	"image-resizer-api/pkg/helpers/logger"

	"github.com/gofiber/fiber/v2"
)

// Parse godoc
// @Description 	Загрузить изображение и сделать из него набор изображений
// @Router      	/control/load [post]
// @Tags        	Control
// @Produce     	form-data
// @Success     	200  {string}  string	"ok"
// @Failure     	500  {string}  string	"error"
func Upload(ctx *fiber.Ctx) error {
	// h := houses.Houses{}
	// h.ParseXml()
	file, errF := ctx.FormFile("file")
	url := ctx.FormValue("url")
	if errF != nil {
		return errF
	}
	if len(url) < 1 {
		return errors.New("url not  valid")
	}
	logger.Log(file.Filename)
	logger.Log(url)
	return ctx.JSON("oka")
}
