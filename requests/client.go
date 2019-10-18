package requests

//Client model
type Client struct {
	Version   string `json:"version" validate:"required"`
	OsVersion string `json:"osVersion" validate:"required"`
	OsType    string `json:"osType" validate:"required"`
	ID        string `json:"id" validate:"required"`
}
