package common

import "fmt"

// CheckErr func
func CheckErr(err error) {
	if err != nil {
		fmt.Println("error = " + err.Error())
	}
}
