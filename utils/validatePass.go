package utils

import "golang.org/x/crypto/bcrypt"

func ValidatePass(pass string, passDb string)(bool){
	hashedPassword := ([]byte)(passDb)
	password := ([]byte)(pass)
	err := bcrypt.CompareHashAndPassword(hashedPassword,password)

	if err != nil{
		return false
	} else {
		return true
	}
}