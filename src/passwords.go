package src

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
bcrypt returns a different hash each time because it incorporates a different random value
into the hash. This is known as a "salt". It prevents people from attacking your hashed passwords
with a "rainbow table", a pre-generated table mapping password hashes back to their passwords.
The salt means that instead of there being one hash for a password, there's 2^16 of them.
Too many to store.

The salt is stored as part of the hashed password.
So bcrypt.CompareHashAndPassword(encryptedPassword, plainPassword)
can encrypt plainPassword using the same salt as encryptedPassword and compare them.

https://stackoverflow.com/questions/52121168/bcrypt-encryption-different-every-time-with-same-input
*/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
