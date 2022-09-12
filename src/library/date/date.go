package date

import "time"

/**
 * Retorna el índice de un elemento en un arreglo de strings
 * date[string] = fecha a validar formato string dd/mm/yyyy
 * return [bool],[error] = [true or false],[descripción del error or  nil]
 */
func CheckDate(date string) error {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return err
	}
	return nil
}
