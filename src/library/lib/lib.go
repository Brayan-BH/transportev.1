package lib

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//typeof
//recibe un valor interface que no se reconoce su tipo y devuelve un string
func InterfaceToString(params ...interface{}) string {
	typeValue := reflect.TypeOf(params[0]).String()
	value := params[0]
	valueReturn := ""
	if strings.Contains(typeValue, "string") {
		toSql := false
		if len(params) == 2 && reflect.TypeOf(params[1]).Kind() == reflect.Bool {
			toSql = params[1].(bool)
		}

		if toSql {
			valueReturn = fmt.Sprintf("'%s'", value)
		} else {
			valueReturn = fmt.Sprintf("%s", value)
		}
	} else if strings.Contains(typeValue, "int") {
		valueReturn = fmt.Sprintf("%d", value)
	} else if strings.Contains(typeValue, "float") {
		valueReturn = fmt.Sprintf("%f", value)
	} else if strings.Contains(typeValue, "bool") {
		valueReturn = fmt.Sprintf("%t", value)
	}
	return valueReturn
}

func BytesToFloat64(bytes []byte) float64 {

	text := bytes // A decimal value represented as Latin-1 text

	f, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		fmt.Print("Error Conv:", err)
	}

	return f
}
