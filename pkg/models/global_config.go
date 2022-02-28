package models

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type Config_channels struct {
	Conn         *websocket.Conn
	C            echo.Context
	Channels_wg  sync.WaitGroup
	Master_mutex *sync.Mutex
}
