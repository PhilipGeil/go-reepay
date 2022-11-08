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

func SettleSession(c *gin.Context, charge go_reepay.SettleDTO, handle string) error {
	settle, erro, err := Reepay.SettleChargeSession(charge, handle)
	if err != nil {
		return err
	}
	if erro != nil {
		c.IndentedJSON(int(erro.HTTPStatus), erro)
		return nil
	}
	c.IndentedJSON(http.StatusOK, settle)
	return nil
}

func CancelSession(c *gin.Context, handle string) error {
	settle, erro, err := Reepay.CancelChargeSession(handle)
	if err != nil {
		return err
	}
	if erro != nil {
		c.IndentedJSON(int(erro.HTTPStatus), erro)
		return nil
	}
	c.IndentedJSON(http.StatusOK, settle)
	return nil
}

func RefundChargeSession(c *gin.Context, charge go_reepay.RefundDTO) error {
	success, erro, err := Reepay.RefundChargeSession(charge)
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
