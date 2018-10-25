package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params string `form:"params" json:"params"` // params
	ID     string `form:"id" json:"id"`         // id
	Name   string `form:"name" json:"name"`     // name
	Key    string `form:"key" json:"key"`       // key
	Value  string `form:"value" json:"value"`   // value

}

// UserParams UserParams
type UserParams struct {
	UserID string `form:"userid" json:"userid"` // userid
	Name   string `form:"name" json:"name"`     // name
	Pwd    string `form:"pwd" json:"pwd"`       // parssword
	Email  string `form:"email" json:"email"`   // email
	OldPwd string `form:"oldpwd" json:"oldpwd"` // old parssword
}
