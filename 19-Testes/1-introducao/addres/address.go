package address

import "strings"

// TypeAddress verify address is valid and return type
func TypeAddress(address string) string {
	validTypes := []string{"rua", "avenida"}

	addressToLowerCase := strings.ToLower(address)

	getFirstWordAddress := strings.Split(addressToLowerCase, " ")[0]

	addressIsValidType := false

	for _, typeAddress := range validTypes {
		if typeAddress == getFirstWordAddress {
			addressIsValidType = true
		}
	}

	if addressIsValidType {
		return strings.Title(getFirstWordAddress)
	}
	return "Invalid type"
}
