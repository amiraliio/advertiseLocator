package models

//AuthCollection collection
const AuthCollection string = "auth"

//EmailAuthType auth type
const EmailAuthType string = "email"

//CellPhoneAuthType auth type
const CellPhoneAuthType string = "cellPhone"

//GoogleAuthType auth type
const GoogleAuthType string = "google"

//FaceBookAuthType auth type
const FaceBookAuthType string = "facebook"

//TwitterAuthType auth type
const TwitterAuthType string = "twitter"

//Auth model
type Auth struct {
	BaseUser
	Value    string `json:"value"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
