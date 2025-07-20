package main

import (
	"forever/dao"
	"forever/fyne"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dao.InitPgsql()
	fyne.Init()
}
