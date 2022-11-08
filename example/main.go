package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	go_reepay "go-reepay"
	"log"
	"os"
)

var Reepay go_reepay.Reepay

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Reepay = go_reepay.Reepay{
		Key:        os.Getenv("REEPAY_PRIV_KEY"),
		SuccessURL: os.Getenv("ACCEPT_URL"),
		CancelURL:  os.Getenv("CANCEL_URL"),
	}
	router := gin.Default()

	router.POST("/charge", CreateCharge)
	router.POST("/refund", RefundCharge)
	router.POST("/settle/:handle", SettleCharge)
	router.POST("/cancel/:handle", CancelCharge)
	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateCharge(c *gin.Context) {
	var charge go_reepay.ChargeDTO
	if err := c.BindJSON(&charge); err != nil {
		fmt.Println(err)
		return
	}
	err := CreateChargeSession(c, charge)
	if err != nil {
		return
	}
}

func SettleCharge(c *gin.Context) {
	var charge go_reepay.SettleDTO
	handle := c.Param("handle")
	if err := c.BindJSON(&charge); err != nil {
		fmt.Println(err)
		return
	}
	err := SettleSession(c, charge, handle)
	if err != nil {
		return
	}
}
func CancelCharge(c *gin.Context) {
	handle := c.Param("handle")
	err := CancelSession(c, handle)
	if err != nil {
		return
	}
}

func RefundCharge(c *gin.Context) {
	var charge go_reepay.RefundDTO
	if err := c.BindJSON(&charge); err != nil {
		fmt.Println(err)
		return
	}
	err := RefundChargeSession(c, charge)
	if err != nil {
		return
	}
}
