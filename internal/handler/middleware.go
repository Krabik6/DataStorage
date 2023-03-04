package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func getPassword(c *gin.Context) (string, error) {
	pass := c.Query("password")
	if len(pass) == 0 {
		return "", errors.New("user password not found")
	}

	return pass, nil
}

func getAddress(c *gin.Context) (string, error) {
	address := c.Query("address")
	if len(address) == 0 {
		return "", errors.New("empty address")
	}
	return address, nil
}

func getData(c *gin.Context) (string, error) {
	data := c.Query("data")
	if len(data) == 0 {
		return "", errors.New("empty data")
	}
	return data, nil
}
