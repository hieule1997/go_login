package models

import (
	"login_jwt/db"
	"login_jwt/constants"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
 	"login_jwt/utils"
)
type RoleForm struct {
	Code     string `json:"code" bson:"code"`
	Describe string `json:"describe" bson:"describe"`
}
type UserForm struct {
	ID       string   `json:"id" bson:"id"`
	Username string   `json:"username" bson:"username"`
	Password string   `json:"password" bson:"password"`
	Fullname string   `json:"fullname" bson:"fullname"`
	Role     RoleForm `json:"role" bson:"role"`
}
type UpdateUserForm struct{
	Fullname string   `json:"fullname" bson:"fullname"`
}
var dbConnect = db.NewConnection()
var UserCollection = dbConnect.Use(constants.MONGO_COLLECTION_USER)

func FindByUserName(username string)(UserForm,error){
	document := UserForm{}
	err := UserCollection.Find(bson.M{"username":username}).One(&document)
	return document,err
}

func CheckPassword(password,hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func Create(newUser UserForm) error{
	hash,_:= bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	newUser.Password = string(hash)
	newUser.ID =utils.GenUUID()
	err := UserCollection.Insert(&newUser)
	return err
}
func FindById(id string) (UserForm, error) {
	document := UserForm{}
	err := UserCollection.Find(bson.M{"id": id}).One(&document)
	return document, err
}
func UpdateUserById(id string, updateData UpdateUserForm) (UserForm, error) {
	change := bson.M{"$set": bson.M{"fullname": updateData.Fullname}}
	query := bson.M{"id": id}
	err := UserCollection.Update(query, change)
	user, _ := FindById(id)
	return user, err
}