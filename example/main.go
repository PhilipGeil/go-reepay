package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	go_reepay "go-reepay"
)

var Reepay = go_reepay.Reepay{
	Key:        "priv_8ed48c74022948202c43bbe22fdf3a20",
	SuccessURL: "https://test.dk",
	CancelURL:  "https://test.dk",
}

func main() {
	router := gin.Default()

	router.POST("/charge", CreateCharge)
	err := router.Run(":8080")
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
