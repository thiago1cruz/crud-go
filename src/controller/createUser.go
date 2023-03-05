package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago1cruz/crud-go/src/configurations/validation"
	"github.com/thiago1cruz/crud-go/src/controller/models/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		errResp := validation.ValidateUserError(err)
		c.JSON(int(errResp.Code), errResp)
		return
	}

	c.JSON(200, userRequest)

}
