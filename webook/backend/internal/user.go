package internal

import "github.com/gin-gonic/gin"

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

}

func (u *UserHandler) ProfileHandler(gtx *gin.Context) {

}

func (u *UserHandler) EditHandler(gtx *gin.Context) {

}
