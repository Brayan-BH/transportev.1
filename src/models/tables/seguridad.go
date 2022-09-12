package tables

import (
	"transporte/src/models"

	"github.com/google/uuid"
)

func Seguridad_GetSchema() ([]models.Base, string) {
	var Seguridad []models.Base
	tableName := "Seguridad"
	Seguridad = append(Seguridad, models.Base{
		Name:        "id",
		Description: "id",
		Required:    true,
		Important:   true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "users",
		Description: "users",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3,
			Max:       25,
			UpperCase: true,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Number_DB(),
			Min:       8,
			Max:       11,
			UpperCase: true,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "l_nomb",
		Description: "l_nomb",
		Required:    true,
		// Update: true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       100,
			UpperCase: true,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "l_apl1",
		Description: "l_apl1",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			UpperCase: true,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "l_apl2",
		Description: "l_apl2",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			UpperCase: true,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "l_emai",
		Description: "l_emai",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       15,
			Max:       150,
			UpperCase: true,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "n_celu",
		Description: "n_celu",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       9,
			Max:       11,
			UpperCase: true,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "k_carg",
		Description: "k_carg",
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	Seguridad = append(Seguridad, models.Base{
		Name:        "l_pass",
		Description: "l_pass",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       20,
			Max:       200,
			Encriptar: true,
		},
	})
	return Seguridad, tableName
}
