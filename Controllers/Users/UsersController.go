package Users

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"rabetpal/Auth/Hash"
	"rabetpal/Database/Cassandra/Models"
	"rabetpal/Requests"
	"rabetpal/Requests/Rules/UsersRequests"
)

var UsersController = struct {
	Register func(ctx iris.Context)
}{
	Register: func(ctx iris.Context) {
		validationErrors := Requests.ValidateForm(ctx, UsersRequests.RegisterRequest)
		switch len(validationErrors) == 0 {
		case false:
			fmt.Println("Validation failed. Errors:",validationErrors)
			return
		}

		Models.NewUser(map[string]interface{}{
			"name":          ctx.FormValue("name"),
			"lastname":      ctx.FormValue("lastname"),
			"username":      ctx.FormValue("username"),
			"mobile":        ctx.FormValue("mobile"),
			"phone":         ctx.FormValue("phone"),
			"national_code": ctx.FormValue("national_code"),
			"email":         ctx.FormValue("email"),
			"password":      Hash.GenerateHash(ctx.FormValue("password")),
			"user_type":     "USER",
			"balance":       0,
		}, nil)
	},
}
