package handler

import (
	"gin-starter/model"
	"gin-starter/services"
	"gin-starter/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandlerMethod struct {
	service *services.AuthServiceMethod
	log     *logrus.Logger
}

func AuthHandler(service *services.AuthServiceMethod, log *logrus.Logger) *AuthHandlerMethod {
	return &AuthHandlerMethod{service, log}
}

func (h *AuthHandlerMethod) LoginHdlr(c *gin.Context) {
	h.log.Println("Execute function LoginHdlr")

	var payloadLogin model.Login
	if err := c.ShouldBindJSON(&payloadLogin); err != nil {
		h.log.Println("Failed to parse payload data login")
		c.JSON(400, utils.ErrorResponse{
			StatusCode: 400,
			Message:    "Error parsing payload data login",
			Error:      err.Error(),
		})
		return
	}

	// validate data
	errorFields, validationError := utils.ValidateData(&payloadLogin)
	if validationError != nil {
		h.log.Println("Validation error in LoginHdlr:", validationError)
		c.JSON(400, utils.ValidatorResponse{
			StatusCode: 400,
			Message:    "Fill The Required Fields",
			Error:      errorFields,
		})
		return
	}

	loginData, errMsg := h.service.LoginSvc(payloadLogin)
	if errMsg != "" {
		h.log.Println("Failed to get data in LoginHdlr")
		c.JSON(401, utils.ErrorResponse{
			StatusCode: 401,
			Message:    errMsg,
			Error:      "Error login",
		})
		return
	}

	c.JSON(200, utils.StandardResponse{
		StatusCode: 200,
		Message:    "Successfully Login!",
		Data:       loginData,
	})

}

func (h *AuthHandlerMethod) VerifyOTPHdlr(c *gin.Context) {
	h.log.Println("Execute function VerifyOTPHdlr")

	var payloadVerify model.VerifyOTP
	if err := c.ShouldBindJSON(&payloadVerify); err != nil {
		h.log.Println("Failed to parse payload data verify OTP")
		c.JSON(400, utils.ErrorResponse{
			StatusCode: 400,
			Message:    "Error parsing payload data verify OTP",
			Error:      err.Error(),
		})
		return
	}

	// validate data
	errorFields, validationError := utils.ValidateData(&payloadVerify)
	if validationError != nil {
		h.log.Println("Validation error in VerifyOTPHdlr:", validationError)
		c.JSON(400, utils.ValidatorResponse{
			StatusCode: 400,
			Message:    "Fill The Required Fields",
			Error:      errorFields,
		})
		return
	}

	verifyOTP, err := h.service.VerifyOTPSvc(payloadVerify)
	if err != "" {
		h.log.Println("Failed to verify data in VerifyOTPHdlr")
		c.JSON(401, utils.ErrorResponse{
			StatusCode: 401,
			Message:    err,
			Error:      "Error verifying OTP",
		})
		return
	}

	c.JSON(200, utils.StandardResponse{
		StatusCode: 200,
		Message:    "Successfully verified OTP",
		Data:       verifyOTP,
	})
}

func (h *AuthHandlerMethod) GetUserProfileHdlr(c *gin.Context) {
	h.log.Println("Execute function GetUserProfileHdlr")

	getUserProfile, err := h.service.GetUserProfileSvc(c)
	if err != nil {
		h.log.Println("Failed to get user profile")
		c.JSON(401, utils.ErrorResponse{
			StatusCode: 401,
			Message:    "Failed to get user profile",
			Error:      err.Error(),
		})
		return
	}

	// example using pagination
	data, meta := utils.GetPaginated(c, 1, 1, getUserProfile)
	c.JSON(200, utils.PaginationResponse{
		StatusCode: 200,
		Message:    "Successfully retrieved user profile",
		Data:       data,
		Meta:       utils.Meta{Pagination: meta},
	})
}
