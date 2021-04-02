package Controllers

import "github.com/kataras/iris/v12"

var HomeController = struct {
	Test func (ctx iris.Context)
}{
	Test: func(ctx iris.Context) {

	},
}
