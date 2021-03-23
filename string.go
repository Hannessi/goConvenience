package goConvenience

// returns true if value is contained within the list
func Contains(value string, list []string) bool {
	for _,listValue := range list {
		if value == listValue {
			return true
		}
	}
	return false
}