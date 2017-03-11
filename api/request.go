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
	r.DB.Find(&cos)
	return c.JSON(http.StatusOK, &cos)
}
