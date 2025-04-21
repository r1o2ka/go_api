package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/t-shimpo/go-echo-gorm-rest/model"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	users := []model.User{}
	model.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
    // URLパラメータからIDを取得
    id := c.Param("id")
    // 既存のユーザーを取得
    user := model.User{}
    if err := model.DB.First(&user, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
    }
    // リクエストボディから更新データをバインド
    updateData := model.User{}
    if err := c.Bind(&updateData); err != nil {
        return err
    }
    // 更新するフィールドを設定
    if updateData.Name != "" {
        user.Name = updateData.Name
    }
    // データベースを更新
    if err := model.DB.Save(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
    }
    return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Delete(&user)
	return c.JSON(http.StatusOK, user)
}