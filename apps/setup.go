package apps

import (
	"fmt"
	"gin-starter/config/db"
	rds "gin-starter/config/redis"
	"gin-starter/handler"
	"gin-starter/repository"
	"gin-starter/routes"
	"gin-starter/services"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartApps() {
	app := gin.Default()

	// setup cors
	app.Use(CorsMiddleware)

	// setup handle panic
	app.Use(gin.Recovery())

	// setup logger
	app.Use(gin.Logger())
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006/01/02 15:04:05",
		DisableTimestamp: false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
	})
	log := logrus.New()
	log.SetOutput(os.Stdout)

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}

	// setup db, redis
	mysql := setupMySQLConnection()
	redis := setupRedisConnection()

	authSetting := setupAuth(mysql, redis, log)
	authRouteConfig := routes.AuthRoute{
		App:         app,
		AuthHandler: authSetting,
	}
	authRouteConfig.SetupAuthRoute()

	errApp := app.Run(":8080")
	if errApp != nil {
		log.Fatalf("Error starting Gin app: %v", errApp)
	}

}

func CorsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	c.Next()
}

func setupMySQLConnection() *gorm.DB {
	hostMysql := os.Getenv("DB_HOST")
	usernameMysql := os.Getenv("DB_USERNAME")
	passwordMysql := os.Getenv("DB_PASSWORD")
	dbMysql := os.Getenv("DB_NAME")

	return db.MysqlConnect(hostMysql, usernameMysql, passwordMysql, dbMysql)
}

func setupRedisConnection() *redis.Client {
	hostRedis := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	passwordRedis := os.Getenv("REDIS_PASSWORD")
	dbRedis, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return rds.RedisConnect(hostRedis, passwordRedis, dbRedis)
}

func setupAuth(mysql *gorm.DB, redis *redis.Client, log *logrus.Logger) *handler.AuthHandlerMethod {
	repo := repository.AuthRepo(mysql, log)
	svc := services.AuthService(repo, redis, log)
	return handler.AuthHandler(svc, log)
}
