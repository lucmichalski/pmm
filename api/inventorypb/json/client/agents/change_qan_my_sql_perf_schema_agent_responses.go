// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeQANMySQLPerfSchemaAgentReader is a Reader for the ChangeQANMySQLPerfSchemaAgent structure.
type ChangeQANMySQLPerfSchemaAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeQANMySQLPerfSchemaAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeQANMySQLPerfSchemaAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeQANMySQLPerfSchemaAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeQANMySQLPerfSchemaAgentOK creates a ChangeQANMySQLPerfSchemaAgentOK with default headers values
func NewChangeQANMySQLPerfSchemaAgentOK() *ChangeQANMySQLPerfSchemaAgentOK {
	return &ChangeQANMySQLPerfSchemaAgentOK{}
}

/*ChangeQANMySQLPerfSchemaAgentOK handles this case with default header values.

A successful response.
*/
type ChangeQANMySQLPerfSchemaAgentOK struct {
	Payload *ChangeQANMySQLPerfSchemaAgentOKBody
}

func (o *ChangeQANMySQLPerfSchemaAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANMySQLPerfSchemaAgent][%d] changeQanMySqlPerfSchemaAgentOk  %+v", 200, o.Payload)
}

func (o *ChangeQANMySQLPerfSchemaAgentOK) GetPayload() *ChangeQANMySQLPerfSchemaAgentOKBody {
	return o.Payload
}

func (o *ChangeQANMySQLPerfSchemaAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeQANMySQLPerfSchemaAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeQANMySQLPerfSchemaAgentDefault creates a ChangeQANMySQLPerfSchemaAgentDefault with default headers values
func NewChangeQANMySQLPerfSchemaAgentDefault(code int) *ChangeQANMySQLPerfSchemaAgentDefault {
	return &ChangeQANMySQLPerfSchemaAgentDefault{
		_statusCode: code,
	}
}

/*ChangeQANMySQLPerfSchemaAgentDefault handles this case with default header values.

An error response.
*/
type ChangeQANMySQLPerfSchemaAgentDefault struct {
	_statusCode int

	Payload *ChangeQANMySQLPerfSchemaAgentDefaultBody
}

// Code gets the status code for the change QAN my SQL perf schema agent default response
func (o *ChangeQANMySQLPerfSchemaAgentDefault) Code() int {
	return o._statusCode
}

func (o *ChangeQANMySQLPerfSchemaAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANMySQLPerfSchemaAgent][%d] ChangeQANMySQLPerfSchemaAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeQANMySQLPerfSchemaAgentDefault) GetPayload() *ChangeQANMySQLPerfSchemaAgentDefaultBody {
	return o.Payload
}

func (o *ChangeQANMySQLPerfSchemaAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeQANMySQLPerfSchemaAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeQANMySQLPerfSchemaAgentBody change QAN my SQL perf schema agent body
swagger:model ChangeQANMySQLPerfSchemaAgentBody
*/
type ChangeQANMySQLPerfSchemaAgentBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangeQANMySQLPerfSchemaAgentParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change QAN my SQL perf schema agent body
func (o *ChangeQANMySQLPerfSchemaAgentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLPerfSchemaAgentBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLPerfSchemaAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMySQLPerfSchemaAgentDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeQANMySQLPerfSchemaAgentDefaultBody
*/
type ChangeQANMySQLPerfSchemaAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change QAN my SQL perf schema agent default body
func (o *ChangeQANMySQLPerfSchemaAgentDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLPerfSchemaAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMySQLPerfSchemaAgentOKBody change QAN my SQL perf schema agent OK body
swagger:model ChangeQANMySQLPerfSchemaAgentOKBody
*/
type ChangeQANMySQLPerfSchemaAgentOKBody struct {

	// qan mysql perfschema agent
	QANMysqlPerfschemaAgent *ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent `json:"qan_mysql_perfschema_agent,omitempty"`
}

// Validate validates this change QAN my SQL perf schema agent OK body
func (o *ChangeQANMySQLPerfSchemaAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANMysqlPerfschemaAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANMySQLPerfSchemaAgentOKBody) validateQANMysqlPerfschemaAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlPerfschemaAgent) { // not required
		return nil
	}

	if o.QANMysqlPerfschemaAgent != nil {
		if err := o.QANMysqlPerfschemaAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQanMySqlPerfSchemaAgentOk" + "." + "qan_mysql_perfschema_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLPerfSchemaAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent QANMySQLPerfSchemaAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent
*/
type ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this change QAN my SQL perf schema agent OK body QAN mysql perfschema agent
func (o *ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum = append(changeQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum, v)
	}
}

const (

	// ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTARTING captures enum value "STARTING"
	ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTARTING string = "STARTING"

	// ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusRUNNING captures enum value "RUNNING"
	ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusRUNNING string = "RUNNING"

	// ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusWAITING captures enum value "WAITING"
	ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusWAITING string = "WAITING"

	// ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTOPPING captures enum value "STOPPING"
	ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusSTOPPING string = "STOPPING"

	// ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusDONE captures enum value "DONE"
	ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgentStatusDONE string = "DONE"
)

// prop value enum
func (o *ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, changeQanMySqlPerfSchemaAgentOkBodyQanMysqlPerfschemaAgentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeQanMySqlPerfSchemaAgentOk"+"."+"qan_mysql_perfschema_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLPerfSchemaAgentOKBodyQANMysqlPerfschemaAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANMySQLPerfSchemaAgentParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeQANMySQLPerfSchemaAgentParamsBodyCommon
*/
type ChangeQANMySQLPerfSchemaAgentParamsBodyCommon struct {

	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`
}

// Validate validates this change QAN my SQL perf schema agent params body common
func (o *ChangeQANMySQLPerfSchemaAgentParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANMySQLPerfSchemaAgentParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeQANMySQLPerfSchemaAgentParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}