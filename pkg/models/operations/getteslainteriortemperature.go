// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"testsdkcreation/pkg/models/shared"
)

type GetTeslaInteriorTemperatureRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetTeslaInteriorTemperatureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// returns the interior temperature of a Tesla.
	Temperature *shared.Temperature
}
