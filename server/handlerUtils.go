package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func formatValidationErrors(validationErrors validator.ValidationErrors) string {
	var formattedErrors string
	for _, err := range validationErrors {
		formattedErrors += fmt.Sprintf("Field '%s' failed on the '%s' tag\n", err.Field(), err.Tag())
	}
	return formattedErrors
}

func parseRequestAndValidate(res http.ResponseWriter, req *http.Request) *WatchList {
	var watchList WatchList
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&watchList); err != nil {
		respondWithParsingError(res, err)
		return nil
	}
	if err := validate.Struct(watchList); err != nil {
		respondWithValidationError(res, formatValidationErrors(err.(validator.ValidationErrors)))
		return nil
	}
	return &watchList
}
func structToString(strct interface{}) {
	var concatStruct string

	valueOfStruct := reflect.ValueOf(strct)
	typeOfStruct := valueOfStruct.Type()

	if typeOfStruct.Kind() == reflect.Struct {
		for i := 0; i < typeOfStruct.NumField(); i++ {
			field := typeOfStruct.Field(i).Name
			fieldType := typeOfStruct.Field(i).Type

			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}

			concatStruct += fmt.Sprintf("Field Name: %s, Field Type: %s\n", field, fieldType)
		}
		log.Printf("%s", concatStruct)
	} else {
		log.Println("Provided interface is not a struct")
	}
}
