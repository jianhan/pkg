package mongod

import (
	mgo "gopkg.in/mgo.v2"
)

func CreateSession(url string) (*mgo.Session, error) {
	s, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return s, nil
}
