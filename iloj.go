package iloj

// Checks if the given string (str) exists in the given array (arr).
func StringInArray(str string, arr []string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}

	return false
}

// An alias for the StringInArray function.
func StrInArr(str string, arr []string) bool {
	return StringInArray(str, arr)
}

func EnvVarValue(varName string, varDefault interface{}) {}
