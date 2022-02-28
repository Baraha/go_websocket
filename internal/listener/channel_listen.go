package listener

import (
	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/internal/listener/channels"
	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/pkg/adapters/cache"
	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/pkg/models"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

func ConfigScheme(msg *redis.Message, c echo.Context, listen_channel string) interface{} {
	// разбираем событие в зависимости от типа канала
	switch listen_channel {
	case "data_currency":
		message := channels.ListenDataCur(msg, c)
		return message

	}

	return nil
}

func Channel(config_channels models.Config_channels, listen_channel string, category string) {
	// status := config_channels.Redisdb.Ping()
	sub := cache.Sub(listen_channel)
	channel := sub.Channel()

	for msg := range channel {
		go func(msg *redis.Message, conn *websocket.Conn) {
			message := ConfigScheme(msg, config_channels.C, listen_channel)

			var event = models.Json_event{Message: message, Category: category}

			config_channels.Master_mutex.Lock()
			defer config_channels.Master_mutex.Unlock()
			if err := conn.WriteJSON(&event); err != nil {
				config_channels.C.Logger().Error(err)
				err_close_channel := sub.Close()
				if err_close_channel != nil {
					config_channels.C.Logger().Error(err_close_channel)
				}
				err_close_conn := conn.Close()
				if err_close_conn != nil {
					config_channels.C.Logger().Error(err_close_conn)
				}

			}

		}(msg, config_channels.Conn)
	}
	config_channels.Channels_wg.Done()
}
