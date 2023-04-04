package handlers

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/webserver/structs"
	"assignment-2/internal/webserver/uptime"
	"encoding/json"
	"errors"
	"net/http"
)

// Webhooks DB
var webhooks []structs.WebhookID

// Init empty list of webhooks
func InitWebhookRegistrations() {
	webhooks = []structs.WebhookID{}
}

// Get number of webhooks
func GetNumberOfWebhooks() int {
	return len(webhooks)
}

// HTTP client
var client = &http.Client{}

// HandlerStatus is a handler for the /status endpoint.
func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to indicate that the response contains JSON data
	w.Header().Add("content-type", "application/json")

	// Return an error if the HTTP method is not GET.
	if r.Method != http.MethodGet {
		http.Error(w, errors.New("method is not supported. Currently only GET are supported").Error(), http.StatusMethodNotAllowed)
		return
	}

	// Get status information.
	status, err := getStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the status information as JSON and send it in the response.
	encoder := json.NewEncoder(w)
	err = encoder.Encode(status)
	if err != nil {
		http.Error(w, errors.New("there were an error during encoding").Error(), http.StatusInternalServerError)
		return
	}
}

func getStatus() (structs.Status, error) {
	// Check the status of the country API.
	url := constants.COUNTRIES_API_URL
	countryApiRequest, _ := http.NewRequest(http.MethodHead, url, nil)

	// Set the content-type header to indicate that the response contains JSON data
	countryApiRequest.Header.Add("content-type", "application/json")

	res, err := client.Do(countryApiRequest)
	if err != nil {
		return structs.Status{}, err
	}

	countriesApiStatus := res.StatusCode

	/*
		// Check the status of the notification db.
		url = constants.NOTIFICATIONDB_URL
		notificationDBRequest, _ := http.NewRequest(http.MethodHead, url, nil)

		res, err = client.Do(notificationDBRequest)
		if err != nil {
			return structs.Status{}, err
		}

		notificationDBStatus := res.StatusCode
	*/

	// get number of registered webhooks
	numWebhooks := GetNumberOfWebhooks()

	// Return a status struct containing information about the uptime and status of the notificationDB and countries APIs.
	return structs.Status{
		CountriesApi: countriesApiStatus,
		//NotificationDB: notificationDBStatus,
		Webhooks: numWebhooks,
		Version:  "v1",
		Uptime:   uptime.GetUptime(),
	}, nil
}
