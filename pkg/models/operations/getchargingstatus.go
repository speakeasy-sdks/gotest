// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"testsdkcreation/pkg/models/shared"
)

type GetChargingStatusRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetChargingStatusResponse struct {
	// return EV Charge reading
	ChargeStatus *shared.ChargeStatus
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}
