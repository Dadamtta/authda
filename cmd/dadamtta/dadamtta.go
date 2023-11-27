package main

import (
	"dadamtta/pkg/cmd/dadamtta"
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
	println("진입점")
	// 웹 서버 실행
	router := gin.Default()
	redisStore, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	// cookieStore := cookie.NewStore([]byte("secret"))
	// cookieStore.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // 1Day
	router.Use(sessions.Sessions("session", redisStore))
	dadamtta.NewCommand(router, db)
	router.Run()
	println("종료")
}
