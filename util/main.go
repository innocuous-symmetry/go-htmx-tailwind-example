package util

import (
	"net/http"
	"strconv"
	"strings"
)

func GetIDFromPath(r *http.Request) (id int, err error) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	last := segments[len(segments)-1]
	id, err = strconv.Atoi(last)
	return
}
