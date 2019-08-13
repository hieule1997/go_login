package db

import (
	container "login_jwt/container"
	"gopkg.in/mgo.v2"
	"fmt"
)
type DBConnection struct {
	session *mgo.Session
}
var config = container.NewContainer()

func NewConnection() (conn *DBConnection) {
	fmt.Println(config.Config.MongoURI)
	session, err := mgo.Dial(config.Config.MongoURI)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}
	return conn
}

func (conn *DBConnection) Use(tableName string) (collection *mgo.Collection) {
	return conn.session.DB(config.Config.DATABASE_NAME).C(tableName)
}

func init() {
	config.Read()
}