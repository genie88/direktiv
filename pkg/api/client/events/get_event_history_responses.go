// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetEventHistoryReader is a Reader for the GetEventHistory structure.
type GetEventHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEventHistoryOK creates a GetEventHistoryOK with default headers values
func NewGetEventHistoryOK() *GetEventHistoryOK {
	return &GetEventHistoryOK{}
}

/* GetEventHistoryOK describes a response with status code 200, with default header values.

successfully got events history
*/
type GetEventHistoryOK struct {
}

func (o *GetEventHistoryOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/events][%d] getEventHistoryOK ", 200)
}

func (o *GetEventHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
