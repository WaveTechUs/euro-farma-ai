package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func Salt()  string{
   return "Salt" 
}

func Encryption(data string) ([]byte, error) {
    cData, err := bcrypt.GenerateFromPassword([]byte(data), 14)
    return cData, err
}

func Descryption() string{
    return "Descryption"
}
