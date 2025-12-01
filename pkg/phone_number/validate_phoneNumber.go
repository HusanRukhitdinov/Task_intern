package phone_number

func IsValidPhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) != 12 {
		return false
	}

	if phoneNumber[:3] != "998" {
		return false
	}

	for _, char := range phoneNumber {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}
