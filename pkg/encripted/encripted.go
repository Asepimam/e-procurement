package encripted

import "golang.org/x/crypto/bcrypt"

type Encripted struct {}


func NewEncripted() *Encripted {
	return &Encripted{}
}

const bycryptCost = 8
func (e *Encripted) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bycryptCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (e *Encripted) CheckPasswordHash(hashedPassword, password string) (bool,error) {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}