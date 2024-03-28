package bcrypt

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var (
	errFailedToHashPassword = errors.New("FAILED_TO_HASH_PASSWORD")
	errMissmatchCostType    = errors.New("MISSMATCH_COST_TYPE")
	errMissmatchPassword    = errors.New("MISSMATCH_PASSWORD")
)

func HashPassword(password string) (string, error) {
	costStr := os.Getenv("BCRYPT_COST")
	cost, err := strconv.Atoi(costStr)
	if errors.Is(err, strconv.ErrSyntax) {
		return "", errMissmatchCostType
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", errFailedToHashPassword
	}
	return string(hashedPass), nil
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errMissmatchPassword
	}
	return nil
}
