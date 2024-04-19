/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stas727/gbot/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Load env error. Please provide env (use .env.example)" + err.Error())
		return
	}
	cmd.Execute()
}
