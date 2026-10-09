package handlers

import (
	"html/template"
	"net/http"
	"slices"
	"uuid"

	"github.com/7om3k-dev/OAuth/internal/store"
)

type authorizeRequestData struct {
	ClientId    string
	RedirectUri string
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

// TODO: serve as static page
func renderErrorPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./templates/error.html")
	err := t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Authorize - In Oauth client passes its identifier in the client_id parameter
func Authorise(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	requestClientId := q.Get("client-id")
	requestRedirectUri := q.Get("redirect-uri")

	// TODO: to log received params
	clientData, clientDataErr := store.GetClientData(requestClientId)

	if clientDataErr != nil {
		renderErrorPage(w)
		// TODO: consider to send logs to leg-service with NATS
	} else if !slices.Contains(clientData.RedirectUris, requestRedirectUri) {
		renderErrorPage(w)
		return
	}

	requestId := uuid.NewV4().String()

	// TODO: store request data in persistent storage
	requests[requestId] = authorizeRequestData{
		ClientId:    requestClientId,
		RedirectUri: requestRedirectUri,
	}

	t, _ := template.ParseFiles("./templates/sign-in.html")
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
