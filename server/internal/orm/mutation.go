package orm

import (
	"errors"
	"strings"
)

// if have choise whom belong creation data, then this func prepare needed options
// 	for ex post creation. post may belong to user or group.
// 	if to user: null,?,?,?,?,?,null
// 	if to group: null,?,?,?,?,null,?
// 	default = first choose
func chooseForeingKeys(datas string, values []interface{}, options []int) (string, []interface{}) {
	valuesCount := len(values)                 // all values
	optionsCount := len(options)               // choose values
	FKsFromIndex := valuesCount - optionsCount // last count will be throw out

	values = values[:FKsFromIndex] // here throw out choose values
	arrFromDatas := strings.Split(datas, ",")

	// here fill throw outed values with null
	for i := FKsFromIndex; i < len(arrFromDatas); i++ {
		arrFromDatas[i] = "null"
	}

	// and choose and fill
	for i, v := range options {
		if v != 0 {
			arrFromDatas[i+FKsFromIndex] = "?"
			values = append(values, options[i])
		}
	}

	datas = strings.Join(arrFromDatas, ",")
	return datas, values
}

// ---------------------Create funcs---------------------------

// Create create one user
func (u *Customer) Create() (int, error) {
	// if u.Nickname == "" || u.Password == "" || u.PhoneNumber == "" {
	// 	return -1, errors.New("н/д")
	// }

	r, e := insertSQL(SQLInsertParams{
		Table:  "Users",
		Datas:  "null,?,?,?",
		Values: MakeArrFromStruct(*u)[1:],
	})
	if e != nil {
		return -1, e
	}
	ID, e := r.LastInsertId()
	return int(ID), e
}

// I think, country & city & travelType & topType is not necessary now

// Create one parsel and return it's ID
func (p *PhotoProject) Create() (int, error) {
	// if p.Description == "" || p.ContactNumber == "" ||
	// 	p.Weight*p.Price*p.CreationDatetime*p.UserID*p.FromID*p.ToID == 0 {
	// 	return -1, errors.New("н/д")
	// }

	params := SQLInsertParams{
		Table:  "Parsels",
		Datas:  "null,?,?,?,?,?,?,?,?,?,?,?,?",
		Values: MakeArrFromStruct(*p),
	}
	params.Datas, params.Values = chooseForeingKeys(params.Datas, params.Values, []int{})
	params.Values = params.Values[1:]

	r, e := insertSQL(params)
	if e != nil {
		return -1, e
	}
	ID, e := r.LastInsertId()
	return int(ID), e
}

// ---------------------Change funcs---------------------------

// Change change user profile
func (u *Customer) Change() error {
	if u.ID == 0 {
		return errors.New("absent/d")
	}

	params := SQLUpdateParams{
		Table:   "Users",
		Couples: map[string]string{},
		Options: DoSQLOption("id=?", "", "", u.ID),
	}

	// if u.PhoneNumber != "" {
	// 	params.Couples["phoneNumber"] = u.PhoneNumber
	// }
	// if u.Nickname != "" {
	// 	params.Couples["nickname"] = u.Nickname
	// }
	// if u.Password != "" {
	// 	params.Couples["password"] = u.Password
	// }
	_, e := updateSQL(params)
	return e
}

// ---------------------Delete funcs---------------------------

// DeleteByParams delete one by id
func DeleteByParams(params SQLDeleteParams) error {
	_, e := deleteSQL(params)
	return e
}
