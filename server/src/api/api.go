package api

import (
	"oruka/api/example"
	"oruka/api/user"
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

func Run(client *ent.Client) {
	r := gin.Default()
	r.GET("/ping", example.Ping)

	user.Router(r, client)

	r.Run(":8080")
}
