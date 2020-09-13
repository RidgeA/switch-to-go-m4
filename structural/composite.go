package structural

import "net/http"

func ComposeMiddlewares(fns ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, fn := range fns {
			fn(w, r)
		}
	}
}
