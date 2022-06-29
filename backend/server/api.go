/*
 * Scheduler API
 *
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"

	"github.com/DennisMaes6/scheduler/backend/model"
)

// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface {
	DbScheduleGet(http.ResponseWriter, *http.Request)
	FileScheduleGet(http.ResponseWriter, *http.Request)
	InstanceDataGetGet(http.ResponseWriter, *http.Request)
	InstanceDataSetPost(http.ResponseWriter, *http.Request)
	ModelParametersGetGet(http.ResponseWriter, *http.Request)
	ModelParametersSetPost(http.ResponseWriter, *http.Request)
	ScheduleGet(http.ResponseWriter, *http.Request)
	GenerateScheduleGet(http.ResponseWriter, *http.Request)
}

// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface {
	DbScheduleGet(context.Context) (ImplResponse, error)
	FileScheduleGet(context.Context) (ImplResponse, error)
	InstanceDataGetGet(context.Context) (ImplResponse, error)
	InstanceDataSetPost(context.Context, model.InstanceData) (ImplResponse, error)
	ModelParametersGetGet(context.Context) (ImplResponse, error)
	ModelParametersSetPost(context.Context, model.ModelParameters) (ImplResponse, error)
	ScheduleGet(context.Context) (ImplResponse, error)
	GenerateScheduleGet(context.Context) (ImplResponse, error)
}
