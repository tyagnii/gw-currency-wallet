package db

import "github.com/google/uuid"

func generateWalletID() (string, error) {
	UUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return UUID.String(), nil
}
