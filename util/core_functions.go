package util

import(
	"golang.org/x/crypto/bcrypt"
)

// func userNameExists(username string) bool {
// 	row := Conn.QueryRow("SELECT id FROM users WHERE username = '" + username + "'")
// 	var user_id string

// 	row.Scan(&user_id)

// 	if user_id > 0 {
		
// 		return true
// 	}
// 	return false
// }

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}