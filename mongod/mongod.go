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

type MRepository struct {
	Session *mgo.Session
}

func (m *MRepository) Close() {
	m.Session.Close()
}

func (m *MRepository) Collection(dbName, collection string) *mgo.Collection {
	return m.Session.DB(dbName).C(collection)
}
