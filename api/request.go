package api

import (
	"net/http"

	"fmt"
	"github.com/labstack/echo"
)

type Request struct {
}

func (r *Request) GetTest(c echo.Context) {
	name := c.QueryParam("name")
	return c.JSON(http.StatusOK, fmt.Sprintf("GET %s", name))
}

func (r *Request) PostTest(c echo.Context) {
	return c.JSON(http.StatusOK, "POST")
}

func (r *Request) PutTest(c echo.Context) {
	return c.JSON(http.StatusOK, "PUT")
}

func (r *Request) DeleteTest(c echo.Context) {
	return c.JSON(http.StatusOK, "DELETE")
}
