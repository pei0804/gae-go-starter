package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/tikasan/gae-go-starter/model"
)

type BBS struct {
	DB *gorm.DB
}

func (b *BBS) Index(c echo.Context) error {
	var cos []model.Comments
	err := b.DB.Order("id desc").Find(&cos).Error; if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	data := struct {
		Comments []model.Comments
	}{
		Comments: cos,
	}
	return c.Render(http.StatusOK, "index", data)
}

func (b *BBS) Show(c echo.Context) error {
	id := c.Param("id")
	var comment model.Comments
	err := b.DB.Where("id = ?", id).Find(&comment).Error; if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	data := struct {
		model.Comments
	}{
		comment,
	}
	return c.Render(http.StatusOK, "show", data)
}

func (b *BBS) New(c echo.Context) error {
	return c.Render(http.StatusOK, "new", "")
}

func (b *BBS) Save(c echo.Context) error {
	comment := c.FormValue("comment")
	name := c.FormValue("name")
	cos := model.Comments{Name: name, Comment: comment}
	err := b.DB.Create(&cos).Error; if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (b *BBS) Edit(c echo.Context) error {
	id := c.Param("id")
	var comment model.Comments
	err := b.DB.Where("id = ?", id).Find(&comment).Error; if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	data := struct {
		model.Comments
	}{
		comment,
	}
	return c.Render(http.StatusOK, "edit", data)
}

func (b *BBS) Update(c echo.Context) error {
	id := c.FormValue("id")
	comment := c.FormValue("comment")
	name := c.FormValue("name")
	var co model.Comments
	err := b.DB.Model(&co).Where("id = ?", id).Update(model.Comments{Name: name, Comment: comment}).Error; if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (b *BBS) DeleteConf(c echo.Context) error {
	id := c.Param("id")
	var comment model.Comments
	err := b.DB.Where("id = ?", id).Find(&comment).Error; if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	data := struct {
		model.Comments
	}{
		comment,
	}
	return c.Render(http.StatusOK, "delete", data)
}

func (b *BBS) Delete(c echo.Context) error {
	id := c.FormValue("id")
	var co model.Comments
	err := b.DB.Where("id = ?", id).Delete(&co).Error; if err != nil {
		return c.Render(http.StatusOK, "err", err.Error())
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func (b *BBS) GetTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "aoue")
}
