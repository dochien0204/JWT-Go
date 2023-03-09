package routes

import (
	"jwt-project/controllers/forgot"
	"jwt-project/controllers/login"
	"jwt-project/controllers/register"
	forgotHandler "jwt-project/handlers/forgot"
	handlersLogin "jwt-project/handlers/login"
	logoutHandler "jwt-project/handlers/logout"
	registerHandler "jwt-project/handlers/register"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(route *gin.Engine) {

	LoginRepository := login.NewLoginRepository()
	loginService := login.NewServiceLogin(LoginRepository)
	loginHandler := handlersLogin.NewHandlerLogin(loginService)

	registerRepository := register.NewRegisterRepository()
	registerService := register.NewRegisterService(registerRepository)
	registerHandler := registerHandler.NewHandlerRegister(registerService)

	forgotRepository := forgot.NewForgotRepository()
	forgotService := forgot.NewForgotService(forgotRepository)
	forgotHandler := forgotHandler.NewForgotHandler(forgotService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/logout", logoutHandler.LogoutHandler)
	groupRoute.POST("/signin", loginHandler.LoginHandler)
	groupRoute.POST("/forgot", forgotHandler.ForgotHandler)
}