package utils

import (
	"log"
	"strconv"
)

func StringToUint(str string) uint {
	uInt, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		log.Fatal("Failed to convert string to uint", err)
	}

	return uint(uInt)
}

func StringToBool(str string) bool {
	boolean, err := strconv.ParseBool(str)
	if err != nil {
		log.Fatal(err)
	}

	return boolean
}
