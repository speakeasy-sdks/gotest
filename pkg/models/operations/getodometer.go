// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"testsdkcreation/pkg/models/shared"
)

type GetOdometerRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetOdometerResponse struct {
	ContentType string
	// return odometer reading
	Odometer    *shared.Odometer
	StatusCode  int
	RawResponse *http.Response
}
