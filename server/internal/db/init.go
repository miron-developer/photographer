package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnToDB is conn to db
var ConnToDB *gorm.DB

func makeMigrations() error {
	if e := migrateSchema(); e != nil {
		return e
	}

	if e := migrateTarrifs(); e != nil {
		return e
	}

	if e := migrateProjectSteps(); e != nil {
		return e
	}
	return nil
}

func migrateSchema() error {
	return ConnToDB.AutoMigrate(
		&Customer{}, &Tariff{}, &BillAccount{}, &Bill{}, &MothlyReport{}, &Session{},
		&ProjectStep{}, &PhotoProject{}, &PhotoCollection{}, &Photo{}, &Link{},
	)
}

func migrateTarrifs() error {
	return ConnToDB.Create(&[]Tariff{
		{Name: "7 дневный пробный, далее бесплатный ограниченный", Cost: 0, Duration: 7},
		{Name: "1 месяц", Cost: 1000, Duration: 30},
		{Name: "6 месяцев", Cost: 5000, Duration: 180},
		{Name: "1 год", Cost: 10000, Duration: 365},
	}).Error
}

func migrateProjectSteps() error {
	return ConnToDB.Create(&[]ProjectStep{
		{Name: "Информация о проекте", Description: "Напишите имя проекту. К примеру, Маржан С."},
		{Name: "Выбор коллекций", Description: "Выберите коллекции, которые хотите отправить клиенту"},
		{Name: "Выбор коллекций клиентом", Description: "Клиент выбирает коллекции, отправленные Вами"},
		{Name: "Заполнение коллекций", Description: "Заполните коллекции фотографиями клиента"},
		{Name: "Клиент выбирает фотографии", Description: "Ваш клиент выбирает нужные фотографии по коллекциям"},
		{Name: "Проект завершен", Description: "Отправьте клиенту ссылку для оплаты. Фотографии будут хранится по ссылке 30дней"}, // how to handle exchange between client and customer?
	}).Error
}

// InitDB init db, settings and tables
func InitDB(log *log.Logger) error {
	// establish connection to db
	log.Println("accessing database...")
	dsn := "host=localhost user=postgres password=postgres dbname=photographer port=5432 sslmode=disable TimeZone=ALMT"
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		return e
	}
	ConnToDB = db
	log.Println("done!")

	// do migrations
	log.Println("migrations...")
	if e = makeMigrations(); e != nil {
		return e
	}
	log.Println("done!")
	return nil
}
