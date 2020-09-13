package structural

import (
	"flag"
	"net/http"
)

func AbstractControllerAction(userId string) {
	// do some work
}

// net/http HandleFunc adapter
func ControllerActionHTTPHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("x-user-id")
	AbstractControllerAction(userId)
}

// CLI adapter
func ControllerActionCLIHandler(args []string) {
	fs := flag.NewFlagSet("", flag.PanicOnError)
	var userId string
	fs.StringVar(&userId, "userId", "", "UserId")
	_ = fs.Parse(args)
	AbstractControllerAction(userId)
}
