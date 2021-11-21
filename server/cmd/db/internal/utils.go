package internal

// relation between string table and origin table
var tables = map[string]interface{}{
	"customer":           Customer{},
	"tariff":             Tariff{},
	"billAccount":        BillAccount{},
	"bill":               Bill{},
	"report":             MothlyReport{},
	"session":            Session{},
	"step":               ProjectStep{},
	"project":            Project{},
	"templateCollection": TemplateCollection{},
	"projectCollection":  ProjectCollection{},
	"selectedCollection": SelectedCollection{},
	"photo":              Photo{},
	"link":               Link{},
}

func DefineModel(dst string) (interface{}, bool) {
	model, ok := tables[dst]
	return model, ok
}

func (service *DB_SERVICE) Create(model interface{}, values map[string]interface{}) (interface{}, error) {
	result := service.DBConn.Model(model).Create(values)
	if result.Error != nil {
		return nil, result.Error
	}

	obj, e := FillStructFromMap(model, values)
	if len(e) != 0 {
		return nil, e[0]
	}

	return MakeArrFromStruct(obj)[0], nil
}
