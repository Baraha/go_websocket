package websocket

import (
	"net/http"
	"sync"

	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/internal/listener"
	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/pkg/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeHTTP(c echo.Context) error {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Fatal(err)
		return err
	}
	defer conn.Close()

	var channels_wg sync.WaitGroup
	var master_mutex sync.Mutex
	var config_channels = models.Config_channels{
		Conn:         conn,
		C:            c,
		Master_mutex: &master_mutex,
		Channels_wg:  channels_wg}

	config_channels.Channels_wg.Add(1)

	go listener.Channel(config_channels, "data_currency", "currency")

	config_channels.Channels_wg.Wait()
	return err
}
