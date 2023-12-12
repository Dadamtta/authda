package main

import (
	"dadamtta/pkg/cmd/dadamtta"
	"dadamtta/pkg/utils/logger"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB 연결
	//SSL is not enabled on the server
	var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "funch", "funch12#$", "test")
	println(connectionString)
	// db, err := sql.Open("postgres", connectionString)
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err)
	}
	logger.Debug("진입점")
	// 웹 서버 실행
	router := gin.Default()
	redisStore, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	// cookieStore := cookie.NewStore([]byte("secret"))
	// cookieStore.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // 1Day
	router.Use(sessions.Sessions("session", redisStore), CORSMiddleware())
	dadamtta.NewCommand(router, db)
	router.Run()
	logger.Debug("종료")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Next()
	}
}
