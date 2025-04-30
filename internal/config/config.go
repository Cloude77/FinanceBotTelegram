// Пакет config отвечает за загрузку настроек из файла .env.
package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config — это структура, где мы храним настройки.

type Config struct {
	TelegramToken string
	DatabaseURL   string
}

// Load загружает настройки из .env и возвращает их.
func Load() (*Config, error) {
	// Загружаем файл .env.
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file", err)
	}
	// Создаем структуру Config и заполняем её значениями.
	cfg := &Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"), // Читаем TELEGRAM_TOKEN из окружения
		DatabaseURL:   os.Getenv("DATABASE_URL"),   // Читаем DATABASE_URL из окружения
	}

	// Проверяем, что токен и URL не пустые.
	if cfg.TelegramToken == "" { // Если токен пустой
		return nil, errors.New("Telegram Token is empty") // возвращаем ошибку
	}
	if cfg.DatabaseURL == "" { // Если URL пустой
		return nil, errors.New("Database URL is empty")
	}

	return cfg, nil
}
