package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiago1cruz/crud-go/src/configurations/logger"
	"github.com/thiago1cruz/crud-go/src/configurations/validation"
	"github.com/thiago1cruz/crud-go/src/controller/models/request"
	"github.com/thiago1cruz/crud-go/src/controller/models/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller",
		zap.String("journe", "create user"),
	)
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	logger.Error("Erro trying to validate user info", err, zap.String("journe", "create user"))
	if err != nil {
		errResp := validation.ValidateUserError(err)
		c.JSON(int(errResp.Code), errResp)
		return
	}

	response := response.UserResponse{
		ID:    "teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully", zap.String("journe", "create user"))

	c.JSON(http.StatusOK, response)

}
