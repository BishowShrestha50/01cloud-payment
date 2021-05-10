package main

import (
	"01cloud-payment/internal/controllers"

	"github.com/joho/godotenv"
	"github.com/prometheus/common/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		log.Info("We are getting the env values")
	}
	controllers.Run()

}
