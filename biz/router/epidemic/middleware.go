// Code generated by hertz generator.

package epidemic

import (
	"Hospital/biz/router/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return auth.Auth(0)
}

func _epidemicMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _try3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addepidemiccaseMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _infoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryepidemiccasebyidMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryepidemiccasebypatientMw() []app.HandlerFunc {
	// your code...
	return nil
}
