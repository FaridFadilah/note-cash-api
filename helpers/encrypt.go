package helpers

import (
	"crypto/sha1"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

/*
parameters:
 - data : value which will encrypted,
 - t : value which will encrypted,
*/
func Encrypt(data string, t string) (v string){
	if t == "password"{
		b, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
		if err != nil{
			panic(err.Error())
		}
		v = string(b)
	} else if t == "uuid" || t == "token"{
		hasher := sha1.New()
	  hasher.Write([]byte(data))
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))	

		v = sha
	}
	return v
}