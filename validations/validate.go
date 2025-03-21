package validations

import (
	"discountmodule/entities"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var validate = validator.New()

func ValidateRequest(req entities.RequestDisCountModule) []string {
	var errors []string

	if err := validate.Struct(req); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()

			// ใช้ reflection เพื่อหาค่า json tag
			var jsonTag string
			// ใช้ reflect เพื่อดึง json tag
			structType := reflect.TypeOf(req)

			// ตรวจสอบชนิดของฟิลด์ก่อนว่าฟิลด์นั้นเป็น struct หรือ slice
			fieldStruct, found := structType.FieldByName(field)
			if !found {
				// ถ้าเป็น slice หรือ array ค้นหาจาก sub-struct
				for i := 0; i < structType.NumField(); i++ {
					f := structType.Field(i)
					if f.Type.Kind() == reflect.Slice || f.Type.Kind() == reflect.Array {
						// ดึง json tag ของ element ใน slice
						elemType := f.Type.Elem()
						subField, foundSub := elemType.FieldByName(field)
						if foundSub {
							jsonTag = subField.Tag.Get("json")
							break
						}
					} else if f.Type.Kind() == reflect.Struct {
						// ถ้าเป็น struct ก็หาจาก field ของ struct นั้น
						subField, foundSub := f.Type.FieldByName(field)
						if foundSub {
							jsonTag = subField.Tag.Get("json")
							break
						}
					}
				}
			} else {
				// ถ้าฟิลด์เป็น struct หรือ field ใน struct นั้นๆ
				jsonTag = fieldStruct.Tag.Get("json")
			}

			// ถ้าไม่มี json tag ให้ใช้ชื่อ field แทน
			if jsonTag == "" {
				jsonTag = field
			}

			// แปลง Error Message ให้อ่านง่าย
			switch tag {
			case "required":
				errors = append(errors, fmt.Sprintf("'%s' เป็นข้อมูลที่จำเป็นต้องระบุใน Body", jsonTag))

			default:
				errors = append(errors, fmt.Sprintf("'%s' is invalid", jsonTag))
			}
		}
	}
	return errors
}