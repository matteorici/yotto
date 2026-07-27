package utils

import "net/http"

func JSON(w http.ResponseWriter) {

	w.Header().Set(

		"Content-Type",

		"application/json",

	)

}
