package util

import (
	"fmt"
	"strings"
)

// Helper func to use with hx-vals
func QueryParams(kv ...interface{}) string {
	if len(kv)%2 != 0 {
		fmt.Printf("invalid args, must be even to form key:value")
	}

	var qp strings.Builder
	qp.WriteString("{ ")
	for i := 0; i < len(kv); i += 2 {
		if i > 0 {
			qp.WriteString(", ")
		}

		qp.WriteString(fmt.Sprintf(`"%v": "%v"`, kv[i], kv[i+1]))
	}
	qp.WriteString(" }")

	return qp.String()
}

// Helper func to set url path params
func PathParams(path string, params ...interface{}) string {
	if path[len(path)-1] != '/' {
		path = path + "/"
	}

	var pp strings.Builder
    pp.WriteString(path)
	for i := 0; i < len(params); i ++ {
		if i > 0 {
			pp.WriteString("/")
		}

		pp.WriteString(fmt.Sprintf("%v", params[i]))
	}

	return pp.String()
}
