package responder

import "net/http"

func SuccessResponse(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}

func FailureResponse(w http.ResponseWriter, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	return
}
