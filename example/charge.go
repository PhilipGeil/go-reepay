package main

import (
	"github.com/gin-gonic/gin"
	go_reepay "go-reepay"
	"net/http"
)

func CreateChargeSession(c *gin.Context, charge go_reepay.ChargeDTO) error {
	success, erro, err := Reepay.CreateChargeSession(charge)
	if err != nil {
		return err
	}
	if erro != nil {
		c.IndentedJSON(int(erro.HTTPStatus), erro)
		return nil
	}
	c.IndentedJSON(http.StatusOK, success)
	return nil
}
