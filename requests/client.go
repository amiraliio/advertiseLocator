package requests

//Client model
type Client struct {
	Version   string `json:"version" validate:"required,min=1,max=70"`
	OsVersion string `json:"osVersion" validate:"required,min=1,max=70"`
	OsType    string `json:"osType" validate:"required,min=1,max=70"`
	ID        string `json:"id" validate:"required,uuid"`
}
