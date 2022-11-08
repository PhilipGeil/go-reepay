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
	router.POST("/settle", SettleCharge)
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
	var charge go_reepay.ChargeDTO
	if err := c.BindJSON(&charge); err != nil {
		fmt.Println(err)
		return
	}
	err := SettleSession(c, charge)
	if err != nil {
		return
	}
}
