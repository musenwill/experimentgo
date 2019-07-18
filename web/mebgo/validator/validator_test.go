package validator

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestUser struct {
	Name  string
	Age   int
	Phone string
	Email string
}

func (*TestUser) GetValidators() map[string][]FieldValidator {
	validators := map[string][]FieldValidator{
		"Age":   {Required(0), Min(18)},
		"Name":  {Required(nil), MinLen(0)},
		"Phone": {Required("110aa"), MaxLen(10)},
	}
	return validators
}

func TestValidatorOK(t *testing.T) {
	JSON := `{"Name":"musenwill", "Age":20, "Email":"musenwill@qq.com"}`
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(JSON), &data)
	if nil != err {
		t.Fatal(err)
	}

	var user *TestUser
	validators := user.GetValidators()
	errors := make([]error, 0)
	for field, validators := range validators {
		for _, validator := range validators {
			err = validator(data, field)
			if nil != err {
				errors = append(errors, err)
			}
		}
	}

	fmt.Println(data)
	fmt.Println(errors)
}
