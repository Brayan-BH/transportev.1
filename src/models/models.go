package models

import "regexp"

type (
	Floats struct {
		Porcentaje bool
		Negativo   bool
		Menor      float64
		Mayor      float64
	}
	Strings struct {
		UpperCase bool
		Encriptar bool
		Cifrar    bool
		Date      bool
		Expr      regexp.Regexp //:=> (I-U-D) ==> expresión regular que debe cumplir el valor que almacenara el campo
		Min       int
		Max       int
	}
	Uints struct {
		Max uint64
	}
	Ints struct {
		Max      int64
		Min      int64
		Negativo bool
	}
	Base struct {
		Name        string      //:=> (I-U-D) ==> nombre del campo
		Description string      //:=> (I-U-D) ==> descripción del campo I-U-D
		Required    bool        //:=> (I)     ==> si el valor del campo es requerido
		Important   bool        //:=> (U-D)   ==> si o si tiene que ser utilizado en un where al actualizar, tiene como valor TRUE ,ó FALSE; I-U-D
		Where       bool        //:=> (U-D)   ==> especifica si el campo sera tomado en cuenta como filtro para el proceso, tiene como valor TRUE ,ó FALSE
		Empty       bool        //:=> (U)     ==> especifica si el campo podrá almacenar valores vacíos
		Type        string      //:=> (I-U-D) ==> tipo del campo que sera evaluado =>ANEXO Types
		Update      bool        //:=> (U)     ==> especifica si el campo podrá ser modificado , tiene como valor TRUE ,ó FALSE
		Default     interface{} //:=> (I)     ==> especifica que valor por defecto tendrá si no se le envía datos
		Strings     Strings     //:=> (I-U)   ==> si almacena el valor en minúsculas
		Float       Floats      //:=> (I-U)   ==> si el valor es un porcentaje
		Uint        Uints       //:=> (I-U)   ==> si el valor es un entero
		Int         Ints        //:=> (I-U)   ==> si el valor es un entero
	}
)

func Null() *regexp.Regexp {
	return regexp.MustCompile(``)
}
func Number_DB() *regexp.Regexp {
	return regexp.MustCompile(`[0-9]{0,}$`)
}
