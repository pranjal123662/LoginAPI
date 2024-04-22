package model

type LoginData struct {
	IP            string `json:"ip,omitempty"`
	UserNumber    string `json:"usernumber,omitempty"`
	IsAlreadyExit bool   `json:"isalreadyexit,omitempty" bson:"isalreadyexit"`
}
