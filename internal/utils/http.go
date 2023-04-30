package utils

import "net/http"

func StatusText(code int) string {
	var message string
	switch code {
	case http.StatusBadRequest:
		message = "The request had invalid inputs or otherwise cannot be served."
	case http.StatusUnauthorized:
		message = "Authorization information is missing or invalid."
	case http.StatusForbidden:
		message = "Access denied to the resource."
	case http.StatusNotFound:
		message = "Unable to find requested record."
	case http.StatusNotAcceptable:
		message = "Not acceptable for the database."
	case http.StatusRequestTimeout:
		message = "Request took too long to process."
	case http.StatusConflict:
		message = "A conflict has occurred."
	case http.StatusRequestedRangeNotSatisfiable:
		message = "No resource available, unable to fulfill the request."
	case http.StatusTooManyRequests:
		message = "Request rate too high, requests from this this user are throttled."
	case http.StatusInternalServerError:
		message = "An error was encountered."
	case http.StatusServiceUnavailable:
		message = "The service is unavailable, please try again later."
	case http.StatusGatewayTimeout:
		message = "The service timed out waiting for an upstream response. Try again later."
	}

	return message
}
