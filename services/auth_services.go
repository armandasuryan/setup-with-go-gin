package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"gin-starter/middleware"
	"gin-starter/model"
	"gin-starter/repository"
	"gin-starter/utils"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type AuthServiceMethod struct {
	repo  *repository.AuthRepoMethod
	redis *redis.Client
	log   *logrus.Logger
}

func AuthService(repo *repository.AuthRepoMethod, redis *redis.Client, log *logrus.Logger) *AuthServiceMethod {
	return &AuthServiceMethod{repo, redis, log}
}

func (s *AuthServiceMethod) LoginSvc(payload model.Login) (model.OTP, string) {
	s.log.Println("Execute function LoginSvc")

	var response model.OTP
	verifyData := s.repo.VerifyDataUserRepo(payload.Username, payload.Password)
	dataUser, _ := s.repo.GetDataUserRepo(payload.Username)

	response.OTP, _ = s.GenerateOTPCode(6)
	duration := time.Minute * 5
	err := s.redis.Set(context.TODO(), payload.Username+"_otp", response.OTP, duration).Err()
	if err != nil {
		s.log.Error("Failed to store OTP in Redis:", err)
	}

	timeLeft := duration.Seconds()
	minutes := int(timeLeft) / 60
	seconds := int(timeLeft) % 60

	response.TimeLeft = fmt.Sprintf("%02d:%02d", minutes, seconds)

	// send otp to email
	if verifyData == "" {
		from := os.Getenv("SMTP_USER")
		to := dataUser.Email
		subject := "OTP Code"
		body := fmt.Sprintf("Your OTP Code is : %v", response.OTP)

		// send email using go routine
		go utils.SentEmail(from, to, subject, body)
	}

	return response, verifyData
}

func (s *AuthServiceMethod) VerifyOTPSvc(payload model.VerifyOTP) (model.ResponseLogin, string) {
	s.log.Println("Execute function VerifyOTPSvc")

	storedOTP, err := s.redis.Get(context.TODO(), payload.Username+"_otp").Result()
	if err == redis.Nil {
		s.log.Error("OTP not found or expired")
		return model.ResponseLogin{}, "OTP not found or expired"
	} else if err != nil {
		s.log.Error("Error getting OTP from Redis:", err)
		errMsg := fmt.Sprintf("Error getting OTP from Redis : %s", err)
		return model.ResponseLogin{}, errMsg
	}

	storedOTPInt, _ := strconv.Atoi(storedOTP)
	if storedOTPInt != payload.OTPCode {
		s.log.Error("Invalid OTP")
		return model.ResponseLogin{}, "Invalid OTP"
	}

	result, _ := s.repo.GetDataUserRepo(payload.Username)
	result.Token, _ = middleware.CreateJwt(result)
	return result, ""
}

func (s *AuthServiceMethod) GenerateOTPCode(maxDigit int) (int, error) {
	s.log.Println("Execute function GenerateOTPCode")

	codes := make([]byte, maxDigit) // set digit
	if _, err := rand.Read(codes); err != nil {
		s.log.Error("Failed get random code")
		return 0, err
	}

	for i := 0; i < maxDigit; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	result, _ := strconv.Atoi(string(codes))
	return result, nil
}

func (s *AuthServiceMethod) GetUserProfileSvc(c *gin.Context) (model.UserProfile, error) {
	s.log.Println("Execute function GetUserProfileSvc")

	token := c.MustGet("token").(jwt.MapClaims)
	username := token["username"]
	return s.repo.GetUserProfile(username.(string))
}
