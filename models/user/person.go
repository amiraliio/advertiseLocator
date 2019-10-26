package models


//TODO order struct field from high to low field type size

//PersonCollection collection
const PersonCollection string = "persons"

//PersonUserType user type
const PersonUserType string = "person"

//AdminUserType user type
const AdminUserType string = "admin"


//Person model
type Person struct {
	BaseUser  ",inline"
	Location  Location `json:"location" bson:"location"`
	Avatar    Image    `json:"avatar" bson:"avatar"`
	FirstName string   `json:"firstName" bson:"firstName"`
	LastName  string   `json:"lastName" bson:"lastName"`
	CellPhone string   `json:"cellPhone" bson:"cellPhone"`
	Email     string   `json:"email" bson:"email"`
	Radius    uint16   `json:"radius" bson:"radius"`
}