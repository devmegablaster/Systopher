package handlers

import (
	"github.com/axrav/SysAnalytics/backend/helpers"
	"github.com/axrav/SysAnalytics/backend/types"
	"github.com/gofiber/websocket/v2"
)

func ServerWS(c *websocket.Conn) {
	servers := c.Locals("servers").([]string)
	serverChannel := make(chan []string, 1)
	dataChannel := make(chan types.ServerData)
	serverChannel <- servers
	go helpers.ServerStats(serverChannel, dataChannel, c)
	for {
		data := <-dataChannel
		c.WriteJSON(data)
	}
}