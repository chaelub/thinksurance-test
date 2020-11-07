// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chaelub/thinksurance/exchange-rates/models"
)

// GetHistoryReader is a Reader for the GetHistory structure.
type GetHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHistoryOK creates a GetHistoryOK with default headers values
func NewGetHistoryOK() *GetHistoryOK {
	return &GetHistoryOK{}
}

/*GetHistoryOK handles this case with default header values.

an object with exchange rates
*/
type GetHistoryOK struct {
	Payload *models.Rates
}

func (o *GetHistoryOK) Error() string {
	return fmt.Sprintf("[GET /history][%d] getHistoryOK  %+v", 200, o.Payload)
}

func (o *GetHistoryOK) GetPayload() *models.Rates {
	return o.Payload
}

func (o *GetHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}