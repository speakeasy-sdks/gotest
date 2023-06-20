// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"testsdkcreation/pkg/models/shared"
)

type StartStopChargeRequest struct {
	VehicleID    string               `pathParam:"style=simple,explode=false,name=vehicle_id"`
	ChargeAction *shared.ChargeAction `request:"mediaType=application/json"`
}

type StartStopChargeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Success Response
	SuccessResponse *shared.SuccessResponse
}
