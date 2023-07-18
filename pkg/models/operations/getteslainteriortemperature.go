// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"testsdkcreation/pkg/models/shared"
)

type GetTeslaInteriorTemperatureRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTeslaInteriorTemperatureRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTeslaInteriorTemperatureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// returns the interior temperature of a Tesla.
	Temperature *shared.Temperature
}

func (o *GetTeslaInteriorTemperatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeslaInteriorTemperatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeslaInteriorTemperatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeslaInteriorTemperatureResponse) GetTemperature() *shared.Temperature {
	if o == nil {
		return nil
	}
	return o.Temperature
}
