package mongo

import (
	"gopkg.in/mgo.v2"
)

// New creates connection to mongo database
func New(url, username, password string) *mgo.Session {
	session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	credentials := mgo.Credential{
		Username: username,
		Password: password,
	}

	err = session.Login(&credentials)

	if err != nil {
		panic(err)
	}

	return session
}
