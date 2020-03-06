package api

import "golang.org/x/crypto/bcrypt"

// 密码加密
func Encrypt(source string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashPwd), err
}

// 密码比对 (传入未加密的密码即可)
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
