package internal

import (
	"fmt"
	"net/http"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	EmailRegexp *regexp2.Regexp
}

func NewUserHandler() *UserHandler {
	const EmailRegexpStr = "^\\w+(-+.\\w+)*@\\w+(-.\\w+)*.\\w+(-.\\w+)*$"
	return &UserHandler{EmailRegexp: regexp2.MustCompile(EmailRegexpStr, regexp2.None)}
}

var UserRoute = NewUserHandler()

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
	if req.Password != req.ConfirmPassword {
		gtx.String(http.StatusOK, "密码不一致")
		return
	}
	isMatched, err := u.EmailRegexp.MatchString(req.Email)
	if err != nil {
		gtx.String(http.StatusOK, fmt.Sprintf("邮箱校验超时 %s", err.Error()))
		return
	}

	if !isMatched {
		gtx.String(http.StatusOK, fmt.Sprintf("邮箱不符合规则 %s", req.Email))
		return
	}
	gtx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) ProfileHandler(gtx *gin.Context) {

}

func (u *UserHandler) EditHandler(gtx *gin.Context) {

}
