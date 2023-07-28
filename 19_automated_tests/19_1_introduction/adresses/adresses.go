package adresses

import "strings"

// IsValidType Verify type of adress
func IsValidType(adress string) string {
	validTypes := []string{"street", "avenue", "road", "highway"}

	// quebrando endereço em um slice
	// comportamento parecido com explode()
	firstWordOfAdress := strings.Split(adress, " ")[0]
	firstWordOfAdressInLowerCase := strings.ToLower(firstWordOfAdress)

	isValidType := false
	for _, validType := range validTypes {
		if validType == firstWordOfAdressInLowerCase {
			isValidType = true
		}
	}

	if isValidType {
		return firstWordOfAdress
	}
	return "Invalid type"
}
