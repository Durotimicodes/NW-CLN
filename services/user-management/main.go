package main

import (
	"github.com/durotimicodes/natwest-clone/user-service/server"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	server.SetUpUserServer(r)

}