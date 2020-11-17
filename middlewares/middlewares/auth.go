package middlewares

import (
	"net/http"
)

func WithAuthorization(whitelistedGroups map[string]struct{}, handler HTTPHandler) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) {
		userGroups, ok := req.Header["User-Group"]
		if !ok {
			res.Write([]byte("missing required `User-Group` header in the request"))
			return
		}

		authorized := false
		for _, v := range userGroups {
			if _, ok := whitelistedGroups[v]; ok {
				authorized = true
				break
			}
		}

		if !authorized {
			res.Write([]byte("user not authorized to access the resource"))
			return
		}

		handler(res, req)
	}
}
