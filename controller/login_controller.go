package controller

import "github.com/gin-gonic/gin"



func (controller *loginController) Login(ctx *gin.Context) string{
	var credentials dto.Credentials
	err := ctx.ShouldBind($credentials)
    if err != nil{
		return ""
	}
	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated{
		return controller.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
