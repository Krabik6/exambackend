package service

import "golang.org/x/crypto/bcrypt"

// checkPasswordHash сравнивает хэш пароля с паролем, введенным пользователем
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
