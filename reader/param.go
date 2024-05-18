package reader

import "strings"

func SplitByTwo(value, seperator string) (string, string, error) {
	value = strings.ReplaceAll(value, " ", "")
	splitString := strings.Split(value, seperator)
	if len(splitString) < 2 {
		return "", "", ErrNotEnoughParameter
	}
	return splitString[0], splitString[1], nil
}

func SplitByThree(value, seperator string) (string, string, string, error) {
	value = strings.ReplaceAll(value, " ", "")
	splitString := strings.Split(value, seperator)
	if len(splitString) < 3 {
		return "", "", "", ErrNotEnoughParameter
	}
	return splitString[0], splitString[1], splitString[2], nil
}
