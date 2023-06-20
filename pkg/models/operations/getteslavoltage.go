// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"testsdkcreation/pkg/models/shared"
)

type GetTeslaVoltageRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetTeslaVoltageResponse struct {
	// returns the voltage of the charger measured by the vehicle.
	ChargeVoltage *shared.ChargeVoltage
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
