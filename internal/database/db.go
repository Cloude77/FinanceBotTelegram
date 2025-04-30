package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Драйвер PostgreSQL (импортируем анонимно)
)

// New создает новое подключение к базе данных.

func New(databaseURL string) (*sql.DB, error) {
	// Подключаемся к базе данных.
	db, err := sql.Open("postgres", databaseURL) // Используем URL из .env
	if err != nil {
		return nil, err
	}
	// Проверяем, что подключение работает.
	err = db.Ping() // Отправляем "пинг" в базу
	if err != nil { // Если произошла ошибка
		return nil, err
	}
	return db, nil
}
