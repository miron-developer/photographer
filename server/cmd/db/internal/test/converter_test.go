package test

import (
	"errors"
	"photographer/cmd/db/internal"
	"reflect"
	"testing"
)

func compareErrorArrs(a, b []error) bool {
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestMakeArrFromStruct(t *testing.T) {
	tests := []struct {
		sampleStruct interface{}
		want         interface{}
		description  string
	}{
		{internal.Customer{FirstName: "a", LastName: "b"}, []interface{}{"a", "b"}, "check many field"},
		{internal.Customer{FirstName: "a"}, []interface{}{"a"}, "check one field"},
		{internal.Customer{}, []interface{}{}, "no field"},
		{&internal.Customer{FirstName: "a"}, []interface{}{}, "pointer"},
	}

	for _, tt := range tests {
		got := internal.MakeArrFromStruct(tt.sampleStruct)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MakeArrFromStruct:(%v) = %v, want %v. Test: %v\n", tt.sampleStruct, got, tt.want, tt.description)
			continue
		}
		t.Logf("MakeArrFromStruct:(%v) PASS", tt.description)
	}
}

func TestFillStructFromMap(t *testing.T) {
	tests := []struct {
		sampleStruct interface{}
		dataSrc      map[string]interface{}
		want         interface{}
		wantErr      []error
		description  string
	}{
		{
			internal.Customer{},
			map[string]interface{}{"FirstName": "fname", "LastName": "lname", "Email": "email"},
			&internal.Customer{FirstName: "fname", LastName: "lname", Email: "email"},
			nil,
			"test just fill",
		},
		{
			internal.Tariff{},
			map[string]interface{}{"Name": "test wrong type filling", "Cost": "lname", "Duration": "email"},
			&internal.Tariff{Name: "test wrong type filling"},
			[]error{errors.New("provided value type didn't match Duration field type."), errors.New("provided value type didn't match Cost field type.")},
			"test wrong field type",
		},
		{
			internal.Customer{},
			map[string]interface{}{"FirstName": "test wrong field name", "Name": "lname", "mail": "email"},
			&internal.Customer{FirstName: "test wrong field name"},
			[]error{errors.New("no such field: Name in obj."), errors.New("no such field: mail in obj.")},
			"test wrong name",
		},
	}
	for _, tt := range tests {
		got, e := internal.FillStructFromMap(tt.sampleStruct, tt.dataSrc)

		if len(e) != 0 && compareErrorArrs(e, tt.wantErr) {
			t.Error("FillStructFromMap: errors", e)
			continue
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FillStructFromMap:direct(%v, %v) = %v, want %v\n", tt.sampleStruct, tt.dataSrc, got, tt.want)
			continue
		}
		t.Logf("FillStructFromMap:(%v) PASS", tt.description)
	}
}
