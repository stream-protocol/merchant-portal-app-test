// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/oxygenpay/oxygen/pkg/api-kms/v1/model"
)

// CreateMaticTransactionReader is a Reader for the CreateMaticTransaction structure.
type CreateMaticTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMaticTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMaticTransactionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMaticTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateMaticTransactionCreated creates a CreateMaticTransactionCreated with default headers values
func NewCreateMaticTransactionCreated() *CreateMaticTransactionCreated {
	return &CreateMaticTransactionCreated{}
}

/* CreateMaticTransactionCreated describes a response with status code 201, with default header values.

Transaction Created
*/
type CreateMaticTransactionCreated struct {
	Payload *model.MaticTransaction
}

func (o *CreateMaticTransactionCreated) Error() string {
	return fmt.Sprintf("[POST /wallet/{walletId}/transaction/matic][%d] createMaticTransactionCreated  %+v", 201, o.Payload)
}
func (o *CreateMaticTransactionCreated) GetPayload() *model.MaticTransaction {
	return o.Payload
}

func (o *CreateMaticTransactionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.MaticTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMaticTransactionBadRequest creates a CreateMaticTransactionBadRequest with default headers values
func NewCreateMaticTransactionBadRequest() *CreateMaticTransactionBadRequest {
	return &CreateMaticTransactionBadRequest{}
}

/* CreateMaticTransactionBadRequest describes a response with status code 400, with default header values.

Validation error / Not found
*/
type CreateMaticTransactionBadRequest struct {
	Payload *model.ErrorResponse
}

func (o *CreateMaticTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /wallet/{walletId}/transaction/matic][%d] createMaticTransactionBadRequest  %+v", 400, o.Payload)
}
func (o *CreateMaticTransactionBadRequest) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *CreateMaticTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
