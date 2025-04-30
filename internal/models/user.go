// описание структур
package models

// User — это структура, описывающая пользователя.
type User struct {
	ID         int64  `db:"id"`
	TelegramID int64  `db:"telegram_id"`
	Username   string `db:"username"`
	CreatedAt  string `db:"created_at"`
}
