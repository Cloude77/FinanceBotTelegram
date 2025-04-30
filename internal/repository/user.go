// Пакет repository содержит логику работы с базой данных.
package repository

// Импортируем библиотеки.
import (
	"finance-bot/internal/models" // Наши модели
	"github.com/jmoiron/sqlx"     // Для работы с базой
)

// Repository — структура для работы с базой.
type Repository struct {
	db *sqlx.DB // Объект базы данных
}

// New создает новый репозиторий.
func New(db *sqlx.DB) *Repository {
	return &Repository{db: db} // Возвращаем структуру с базой
}

// Create добавляет нового пользователя в базу.
func (r *Repository) Create(user models.User) error {
	// SQL-запрос для вставки пользователя.
	query := `
        INSERT INTO users (telegram_id, username)
        VALUES (:telegram_id, :username)
        ON CONFLICT (telegram_id) DO NOTHING
    `
	// Выполняем запрос, передавая структуру user.
	_, err := r.db.NamedExec(query, user)
	return err // Возвращаем ошибку (или nil, если всё ок)
}
