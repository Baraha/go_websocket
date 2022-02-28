package channels

import (
	"95.85.87.178/BonoboGitServer/IQ-Services-backend.git/pkg/models"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo"
)

func ListenDataCur(msg *redis.Message, c echo.Context) *models.Data {
	var message *models.Data
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(msg.Payload), &message)
	if err != nil {
		c.Logger().Fatal(err)
	}
	return message
}
