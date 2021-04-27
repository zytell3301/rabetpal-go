package UsersRequests

import "rabetpal/Requests"

var RegisterRequest = Requests.ValidationRule{
	Rules: map[string]interface{}{
		"name":          "required,min=2,max=64",
		"lastname":      "required,min=2,max=64",
		"username":      "omitempty,min=8,max=64",
		"mobile":        "required,min=10,max=13",
		"phone":         "omitempty,len=11",
		"national_code": "omitempty,len=10",
		"email":         "required,email",
		"password":      "required,min=8,max=64",
	},
	ErrorMessages: nil,
}
