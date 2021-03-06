package les2lib

import (
	"errors"
	"github.com/aidarkhanov/nanoid/v2"
	"log"
	"math"
	"unicode"
)

func ChangeLetterCase(char rune) rune {
	if unicode.IsLower(char) {
		return unicode.ToUpper(char)
	} else {
		return unicode.ToLower(char)
	}
}


func FindRoots(a, b, c float64) (float64, float64, error) {
	d := b*b-4*a*c
	if d < 0 {
		return -1,-1,errors.New("discriminant < 0")
	}
	d = math.Sqrt(d)
	r1, r2 := (-b+d)/(2*a), (-b-d)/(2*a)
	return r1,r2,nil
}

func GetUUID() string {
	id, err := nanoid.New() //> "i25_rX9zwDdDn7Sg-ZoaH"
	if err != nil {
		log.Fatalln(err)
	}
	return id
}