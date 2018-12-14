package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"log"
)

// MyStruct ..

// use a single instance of Validate, it caches struct info
//var validate *validator.Validate

func main() {
	type MyStruct struct {
		String string `validate:"is-awesome"`
	}

	validate := validator.New()
	validate.RegisterValidation("is-awesome", ValidateMyVal)

	s := MyStruct{String: "awesome"}

	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	s.String = "not awesome"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	log.Println("fl.Field().String():", fl.Field().String())
	return fl.Field().String() == "awesome"
}