// Code generated by go-swagger; DO NOT EDIT.

package xtra_db_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateXtraDBClusterReader is a Reader for the UpdateXtraDBCluster structure.
type UpdateXtraDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateXtraDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateXtraDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateXtraDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateXtraDBClusterOK creates a UpdateXtraDBClusterOK with default headers values
func NewUpdateXtraDBClusterOK() *UpdateXtraDBClusterOK {
	return &UpdateXtraDBClusterOK{}
}

/*UpdateXtraDBClusterOK handles this case with default header values.

A successful response.
*/
type UpdateXtraDBClusterOK struct {
	Payload interface{}
}

func (o *UpdateXtraDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBCluster/Update][%d] updateXtraDbClusterOk  %+v", 200, o.Payload)
}

func (o *UpdateXtraDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateXtraDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateXtraDBClusterDefault creates a UpdateXtraDBClusterDefault with default headers values
func NewUpdateXtraDBClusterDefault(code int) *UpdateXtraDBClusterDefault {
	return &UpdateXtraDBClusterDefault{
		_statusCode: code,
	}
}

/*UpdateXtraDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type UpdateXtraDBClusterDefault struct {
	_statusCode int

	Payload *UpdateXtraDBClusterDefaultBody
}

// Code gets the status code for the update xtra DB cluster default response
func (o *UpdateXtraDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *UpdateXtraDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBCluster/Update][%d] UpdateXtraDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateXtraDBClusterDefault) GetPayload() *UpdateXtraDBClusterDefaultBody {
	return o.Payload
}

func (o *UpdateXtraDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateXtraDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateXtraDBClusterBody update xtra DB cluster body
swagger:model UpdateXtraDBClusterBody
*/
type UpdateXtraDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// XtraDB cluster name.
	Name string `json:"name,omitempty"`

	// params
	Params *UpdateXtraDBClusterParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this update xtra DB cluster body
func (o *UpdateXtraDBClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateXtraDBClusterBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateXtraDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateXtraDBClusterBody) UnmarshalBinary(b []byte) error {
	var res UpdateXtraDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateXtraDBClusterDefaultBody update xtra DB cluster default body
swagger:model UpdateXtraDBClusterDefaultBody
*/
type UpdateXtraDBClusterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this update xtra DB cluster default body
func (o *UpdateXtraDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateXtraDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdateXtraDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateXtraDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateXtraDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateXtraDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateXtraDBClusterParamsBodyParams UpdateXtraDBClusterParams represents XtraDB cluster parameters that can be updated.
swagger:model UpdateXtraDBClusterParamsBodyParams
*/
type UpdateXtraDBClusterParamsBodyParams struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// proxysql
	Proxysql *UpdateXtraDBClusterParamsBodyParamsProxysql `json:"proxysql,omitempty"`

	// pxc
	Pxc *UpdateXtraDBClusterParamsBodyParamsPxc `json:"pxc,omitempty"`
}

// Validate validates this update xtra DB cluster params body params
func (o *UpdateXtraDBClusterParamsBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateXtraDBClusterParamsBodyParams) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateXtraDBClusterParamsBodyParams) validatePxc(formats strfmt.Registry) error {

	if swag.IsZero(o.Pxc) { // not required
		return nil
	}

	if o.Pxc != nil {
		if err := o.Pxc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res UpdateXtraDBClusterParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateXtraDBClusterParamsBodyParamsProxysql ProxySQL container parameters.
swagger:model UpdateXtraDBClusterParamsBodyParamsProxysql
*/
type UpdateXtraDBClusterParamsBodyParamsProxysql struct {

	// compute resources
	ComputeResources *UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this update xtra DB cluster params body params proxysql
func (o *UpdateXtraDBClusterParamsBodyParamsProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateXtraDBClusterParamsBodyParamsProxysql) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsProxysql) UnmarshalBinary(b []byte) error {
	var res UpdateXtraDBClusterParamsBodyParamsProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources
*/
type UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this update xtra DB cluster params body params proxysql compute resources
func (o *UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources) UnmarshalBinary(b []byte) error {
	var res UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateXtraDBClusterParamsBodyParamsPxc PXC container parameters.
swagger:model UpdateXtraDBClusterParamsBodyParamsPxc
*/
type UpdateXtraDBClusterParamsBodyParamsPxc struct {

	// compute resources
	ComputeResources *UpdateXtraDBClusterParamsBodyParamsPxcComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this update xtra DB cluster params body params pxc
func (o *UpdateXtraDBClusterParamsBodyParamsPxc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateXtraDBClusterParamsBodyParamsPxc) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsPxc) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsPxc) UnmarshalBinary(b []byte) error {
	var res UpdateXtraDBClusterParamsBodyParamsPxc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateXtraDBClusterParamsBodyParamsPxcComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model UpdateXtraDBClusterParamsBodyParamsPxcComputeResources
*/
type UpdateXtraDBClusterParamsBodyParamsPxcComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this update xtra DB cluster params body params pxc compute resources
func (o *UpdateXtraDBClusterParamsBodyParamsPxcComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsPxcComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateXtraDBClusterParamsBodyParamsPxcComputeResources) UnmarshalBinary(b []byte) error {
	var res UpdateXtraDBClusterParamsBodyParamsPxcComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
