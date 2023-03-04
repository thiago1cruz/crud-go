package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thiago1cruz/crud-go/src/controller/models/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	fmt.Println(userRequest)

}
