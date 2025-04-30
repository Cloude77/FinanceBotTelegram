package main

import (
	"finance-bot/internal/config"
	"finance-bot/internal/database"
	"finance-bot/internal/telegram"
	"log"
)

func main() {

	// Загружаем настройки из файла .env.
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Подключаемся к базе данных PostgreSQL
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Создаем и запускаем бота.
	bot, err := telegram.NewBot(cfg.TelegramToken, db)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}
	bot.Start()
}
