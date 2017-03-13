package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/tikasan/gae-go-starter/model"
)

type Request struct {
	DB *gorm.DB
}

func (r *Request) GetAllComment(c echo.Context) error {
	var cos []model.Comments
	err := r.DB.Find(&cos).Error
	if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	return c.JSON(http.StatusOK, &cos)
}
