package internal

import (
	"errors"
)

func (service *DB_SERVICE) makeMigrations() error {
	service.Log.Infoln("make migrations...")
	if e := service.migrateSchema(); e != nil {
		return e
	}

	if e := service.migrateTarrifs(); e != nil {
		return e
	}

	if e := service.migrateProjectSteps(); e != nil {
		return e
	}

	if e := service.migrateReport(); e != nil {
		return e
	}
	service.Log.Infoln("done!")
	return nil
}

func (service *DB_SERVICE) migrateSchema() error {
	return service.DBConn.AutoMigrate(
		&Customer{}, &Tariff{}, &BillAccount{}, &MothlyReport{}, &Bill{}, &Session{},
		&ProjectStep{}, &Project{}, &TemplateCollection{}, &ProjectCollection{}, &SelectedCollection{}, &Photo{}, &Link{},
	)
}

func (service *DB_SERVICE) checkExistValues(dst interface{}) (bool, error) {
	if service.DBConn.AutoMigrate(dst) != nil || !service.DBConn.Migrator().HasTable(dst) {
		return false, errors.New("table not found")
	}

	if service.DBConn.First(dst).Error == nil {
		return true, nil
	}
	return false, nil
}

func (service *DB_SERVICE) migrateTarrifs() error {
	tariffs := []Tariff{
		{Name: "7 дневный пробный, далее бесплатный ограниченный", Cost: 0, Duration: 7},
		{Name: "1 месяц", Cost: 1000, Duration: 30},
		{Name: "6 месяцев", Cost: 5000, Duration: 180},
		{Name: "1 год", Cost: 10000, Duration: 365},
	}

	ok, e := service.checkExistValues(&Tariff{})
	if e != nil {
		return e
	}
	if !ok {
		return service.DBConn.Create(&tariffs).Error
	}
	return nil
}

func (service *DB_SERVICE) migrateProjectSteps() error {
	steps := []ProjectStep{
		{Name: "Информация о проекте", Description: "Напишите имя проекту. К примеру, Маржан С."},
		{Name: "Выбор коллекций", Description: "Выберите коллекции, которые хотите отправить клиенту"},
		{Name: "Выбор коллекций клиентом", Description: "Клиент выбирает коллекции, отправленные Вами"},
		{Name: "Заполнение коллекций", Description: "Заполните коллекции фотографиями клиента"},
		{Name: "Клиент выбирает фотографии", Description: "Ваш клиент выбирает нужные фотографии по коллекциям"},
		{Name: "Проект завершен", Description: "Поделитесь с оригиналами фотографий и возьмите свой заработок"},
	}

	ok, e := service.checkExistValues(&ProjectStep{})
	if e != nil {
		return e
	}
	if !ok {
		return service.DBConn.Create(&steps).Error
	}
	return nil
}

func (service *DB_SERVICE) migrateReport() error {
	rp := MothlyReport{Total: 0}
	ok, e := service.checkExistValues(&MothlyReport{})
	if e != nil {
		return e
	}
	if !ok {
		return service.DBConn.Create(&rp).Error
	}
	return nil
}
