package mongod

import (
	mgo "gopkg.in/mgo.v2"
)

func NewSession(url string) (*mgo.Session, func(), error) {
	closeFunc := func() {}
	db, err := mgo.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	closeFunc = func() {
		db.Close()
	}
	return db, closeFunc, nil
}
