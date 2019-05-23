package test

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	log "bitbucket.org/tekion/tbaas/log/v1"
	api "bitbucket.org/tekion/tbaas/tapi"
	"net/http"
)

const servicePort = "8081"

// Start - start the service console app
func Start() {

	api.AddNoAuthServiceRoute(
		"get dealer by _id",
		http.MethodGet,
		"/dealer",
		getDealer,
	)
	api.AddNoAuthServiceRoute(
		"post dealer",
		http.MethodPost,
		"/add",
		addNewDealer,
	)
	api.AddNoAuthServiceRoute(
		"edit dealer",
		http.MethodPut,
		"/edit/{id}",
		editByDealer,
	)

	api.AddNoAuthServiceRoute(
		"add communication collection",
		http.MethodPost,
		"/communication/add",
		addCommunication,
	)

	api.AddNoAuthServiceRoute(
		"get communication collection",
		http.MethodGet,
		"/communication/get",
		getCommunication,
	)

	log.GenericInfo(apiContext.TContext{}, "Started TAP backend Service", log.FieldsMap{"port": servicePort})
	api.Start(servicePort, "/test")
}
