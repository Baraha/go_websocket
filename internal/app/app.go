package app

import (
	"fmt"
	"sync"

	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/internal/websocket"
	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/pkg/adapters/cache"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func Run() {

	redisdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	cache.SetRedisClient(redisdb)

	var wg sync.WaitGroup
	e := echo.New()
	wg.Add(4)
	go func(e *echo.Echo) {
		e.Use(middleware.Logger())
		wg.Done()
	}(e)
	go func(e *echo.Echo) {
		e.Pre(middleware.RemoveTrailingSlash())
		wg.Done()
	}(e)
	go func(e *echo.Echo) {
		e.Logger.SetLevel(log.DEBUG)
		wg.Done()
	}(e)
	go func(e *echo.Echo) {
		e.GET("/ws", websocket.ServeHTTP)
		wg.Done()
	}(e)
	wg.Wait()
	fmt.Print("server websocket starting")
	e.Logger.Fatal(e.Start("0.0.0.0:9876"))

}
