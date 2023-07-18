// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"testsdkcreation/pkg/models/shared"
)

type ListVehiclesRequest struct {
	// Number of vehicles to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Index to start vehicle list at
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *ListVehiclesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListVehiclesRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type ListVehiclesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A list of vehicles
	VehiclesResponse *shared.VehiclesResponse
}

func (o *ListVehiclesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListVehiclesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListVehiclesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListVehiclesResponse) GetVehiclesResponse() *shared.VehiclesResponse {
	if o == nil {
		return nil
	}
	return o.VehiclesResponse
}
