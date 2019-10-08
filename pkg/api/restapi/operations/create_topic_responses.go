// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// CreateTopicOKCode is the HTTP code returned for type CreateTopicOK
const CreateTopicOKCode int = 200

/*CreateTopicOK Topic created

swagger:response createTopicOK
*/
type CreateTopicOK struct {

	/*
	  In: Body
	*/
	Payload *models.Topic `json:"body,omitempty"`
}

// NewCreateTopicOK creates CreateTopicOK with default headers values
func NewCreateTopicOK() *CreateTopicOK {

	return &CreateTopicOK{}
}

// WithPayload adds the payload to the create topic o k response
func (o *CreateTopicOK) WithPayload(payload *models.Topic) *CreateTopicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create topic o k response
func (o *CreateTopicOK) SetPayload(payload *models.Topic) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTopicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTopicDefault Error

swagger:response createTopicDefault
*/
type CreateTopicDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.DefaultError `json:"body,omitempty"`
}

// NewCreateTopicDefault creates CreateTopicDefault with default headers values
func NewCreateTopicDefault(code int) *CreateTopicDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTopicDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create topic default response
func (o *CreateTopicDefault) WithStatusCode(code int) *CreateTopicDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create topic default response
func (o *CreateTopicDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create topic default response
func (o *CreateTopicDefault) WithPayload(payload *models.DefaultError) *CreateTopicDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create topic default response
func (o *CreateTopicDefault) SetPayload(payload *models.DefaultError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTopicDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
