package auth

import (
	"errors"
	"github.com/avishkar-activex/chms-auth/models/user"
	log "github.com/sirupsen/logrus"
	"strings"
)

func AuthenticateUser(name, password string) (user.User, error) {
	usr, err := user.FindByName(name)
	if err != nil {
		log.Errorf("user not found %v", err)
		return user.User{}, err
	}
	if strings.Compare(usr.Password, password) != 0 {
		return user.User{}, errors.New("password doesn't match")
	}

	return usr, nil
}
