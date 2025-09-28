package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Объявляем глобальную переменную DB, которая хранит соединение с БД
// Тип *gorm.DB — это объект, через который потом будем выполнять запросы
// Мы делаем её глобальной, чтобы из других пакетов (handlers, repository) можно было легко обращаться к БД

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	// dsn = Data Source Name = строка с параметрами подключения к БД

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Открываем соединение с БД через GORM:
	// postgres.Open(dsn) — указываем драйвер Postgres и строку подключения
	// &gorm.Config{} — конфигурация GORM (здесь пустая, можно позже добавлять настройки)
	// Если всё успешно, gorm.Open вернёт объект *gorm.DB

	if err != nil {
		log.Fatalf("Failed to connect to database:%v", err)
	}
	fmt.Println("Successfully connected to database")
	// Проверяем ошибку, если соединение не удалось → выводим сообщение и останавливаем программу (log.Fatal)
	// Если ошибок нет — выводим сообщение, что подключение прошло успешно

	DB = db
	// Сохраняем соединение в глобальную переменную DB, чтобы использовать его в других местах проекта
}
