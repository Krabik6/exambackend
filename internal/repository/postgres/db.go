package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func NewDB(host, port, user, password, dbname string) (*sql.DB, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Println("dbURL: ", dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при проверке соединения с базой данных: %w", err)
	}

	return db, nil
}

//var once sync.Once
//var instance *sql.DB
//
//// InitDB инициализирует подключение к базе данных и сохраняет его как Singleton.
//func InitDB(dataSourceName string) {
//	once.Do(func() {
//		var err error
//		instance, err = sql.Open("postgres", dataSourceName)
//		if err != nil {
//			log.Fatalf("Не удалось подключиться к базе данных: %v", err)
//		}
//
//		// Проверьте соединение
//		if err = instance.Ping(); err != nil {
//			log.Fatalf("Не удалось выполнить ping базы данных: %v", err)
//		}
//
//		log.Println("Подключение к базе данных успешно установлено")
//	})
//
//	NewTransactionsCategory(instance).Init()
//}
//
//// GetDB возвращает экземпляр подключения к базе данных.
//func GetDB() *sql.DB {
//	if instance == nil {
//		log.Fatal("Инициализация подключения к базе данных не выполнена")
//	}
//	return instance
//}
