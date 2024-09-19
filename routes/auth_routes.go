package routes

import (
	"gin-starter/handler"
	"gin-starter/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	App         *gin.Engine
	AuthHandler *handler.AuthHandlerMethod
}

func (r *AuthRoute) SetupAuthRoute() {
	authRoute := r.App.Group("auth/api/v1")

	authRoute.POST("/login", r.AuthHandler.LoginHdlr)
	authRoute.POST("/login/verify-otp", r.AuthHandler.VerifyOTPHdlr)
	authRoute.GET("/user-profile", middleware.JWTMiddleware, r.AuthHandler.GetUserProfileHdlr)
}
