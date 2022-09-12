// package financiera
// import "server-go/src/models"
// func ClientesCars_GetSchema_DB() ([]models.Base_DB, string) {
// 	var ClientesCars []models.Base_DB
// 	tableName := "Fina_" + "ClientesCars"
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"c_plac",
// 		Description:"c_plac",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:0.600000,
// 			Max:6,
// 			UpperCase:true,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"n_docu",
// 		Description:"n_docu",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:1.100000,
// 			Max:11,
// 			UpperCase:true,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"l_marc",
// 		Description:"l_marc",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:5.000000,
// 			Max:50,
// 			UpperCase:true,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"l_mode",
// 		Description:"l_mode",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:5.000000,
// 			Max:50,
// 			UpperCase:true,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"l_color",
// 		Description:"l_color",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:7.000000,
// 			Max:70,
// 			UpperCase:true,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"c_year",
// 		Description:"c_year",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:4,
// 			Max:4,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"c_mode",
// 		Description:"c_mode",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:4,
// 			Max:4,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"n_seri",
// 		Description:"n_seri",
// 		Required: true,
// 		Update: true,
// 		Type:"uint64",
// 		Uint: models.Uints{
// 			Max: 10,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"n_pasa",
// 		Description:"n_pasa",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:2.000000,
// 			Max:20,
// 			UpperCase:true,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"l_obse",
// 		Description:"l_obse",
// 		Required: true,
// 		Update: true,
// 		Type:"string",
// 		Strings: models.Strings{
// 			Expr:      *models.Null(),
// 			Min:10.000000,
// 			Max:100,
// 			UpperCase:true,
// 		},
// 	})
// 	ClientesCars = append(ClientesCars, models.Base_DB{
// 		Name:"k_stad",
// 		Description:"k_stad",
// 		Required: true,
// 		Update: true,
// 		Type:"uint64",
// 		Uint: models.Uints{
// 			Max: 10,
// 		},
// 	})
// 	return ClientesCars, tableName
// }

package routes
