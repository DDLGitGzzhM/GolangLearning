package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"GolangLearning/webook/backend/pkg"
)

type UserHandler struct {
}

var UserRoute UserHandler

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	up := server.Group("/users")
	up.GET("/profile", u.ProfileHandler)
	up.POST("/login", u.ProfileHandler)
	up.POST("/signup", u.SignUpHandler)
	up.POST("/edit", u.EditHandler)
}

func (u *UserHandler) LoginHandler(gtx *gin.Context) {

}

func (u *UserHandler) SignUpHandler(gtx *gin.Context) {
	type SignUpHandler struct {
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}
	var req SignUpHandler
	if err := gtx.Bind(&req); err != nil {
		return
	}
	fmt.Println(pkg.JsonMarshal(req))
	gtx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) ProfileHandler(gtx *gin.Context) {

}

func (u *UserHandler) EditHandler(gtx *gin.Context) {

}
