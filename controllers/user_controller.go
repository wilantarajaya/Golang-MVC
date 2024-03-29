package controllers

import (
	"net/http"
	"prakerja6/config"
	"prakerja6/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: user,
	})
}

func DetailUserController(c echo.Context) error {
	var users []models.User 


	var id, _ = strconv.Atoi(c.Param("id"))

	result := config.DB.First(&users, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: users,
	})
}

func GetAllUSer(c echo.Context) error {
	var users []models.User 
	
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: users,
	})
}

func DeleteUser(c echo.Context) error {
	var users []models.User 
	var id, _ = strconv.Atoi(c.Param("id"))

	result := config.DB.Delete(&users, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal Menghapus Data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success Menghapus data", Status: true, Data: users,
	})
}


