package controller

import (
	"github.com/orangesys/orangeapi/pkg/config"
	"github.com/orangesys/orangeapi/pkg/firebase"
)

// CheckConsumer with firebase
func CheckConsumer(uuid string) error {
	//	uuid := "iGzNX6QzfudVlwKtR8CQCj0itIU2"
	//	load env ORANGEAPI_FIREBASE_AUTH
	//	     env ORANGEAPI_FIREBASE_URL
	firebaseconfig, err := config.LoadFirebaseConfig()
	if err != nil {
		return err
	}
	user := firebase.FirebaseConfiguration{
		Config: firebaseconfig,
		UUID:   uuid,
	}
	if err = user.CheckUser(); err != nil {
		return err
	}
	return nil
}
