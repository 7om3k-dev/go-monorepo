package main

import (
	"html/template"
	"net/http"
	"slices"
	"uuid"

	"github.com/7om3k/link-share/auth-server/store"
)

type authorizeRequestData struct {
	ClientId    string
	RedirectUri string
}

type authorizeErrorPageData struct {
	errorMsg string
}

type signUpPageData struct {
	ClientId  string
	RequestId string
}

// TODO: Use server side storage like DB for requests.
var requests = make(map[string]authorizeRequestData)

/*
Note:
OAuth 2.0 is designed as an HTTP-based protocol.
Authorization server is expected to be available to the client over HTTP for both the front and back channels
*/

func renderErrorPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./pages/error.html")
	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Authorize - In Oauth client passes its identifier in the client_id parameter
func authorize(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	requestClientId := q.Get("client-id")
	requestRedirectUri := q.Get("redirect-uri")

	// TODO: to log received params

	clientData := store.GetClientData(requestClientId)

	if clientData.IsZeroValue() {
		renderErrorPage(w)
		return
		// TODO: consider to send logs to leg-service with NATS broker
	}

	if !slices.Contains(clientData.RedirectUris, requestRedirectUri) {
		renderErrorPage(w)
		return
	}

	requestId := uuid.NewV4().String()

	requests[requestId] = authorizeRequestData{
		ClientId:    requestClientId,
		RedirectUri: requestRedirectUri,
	}

	t, _ := template.ParseFiles("./pages/sign-up.html")
	err := t.Execute(w, signUpPageData{
		ClientId:  clientData.ClientId,
		RequestId: requestId,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func approve(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Endpoint not ready", http.StatusInternalServerError)
}
