// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewRevokeInvitationParams creates a new RevokeInvitationParams object
//
// There are no default values defined in the spec.
func NewRevokeInvitationParams() RevokeInvitationParams {

	return RevokeInvitationParams{}
}

// RevokeInvitationParams contains all the bound params for the revoke invitation operation
// typically these are obtained from a http.Request
//
// swagger:parameters revoke_invitation
type RevokeInvitationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ListID string
	/*
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRevokeInvitationParams() beforehand.
func (o *RevokeInvitationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rListID, rhkListID, _ := route.Params.GetOK("list_id")
	if err := o.bindListID(rListID, rhkListID, route.Formats); err != nil {
		res = append(res, err)
	}

	rUserID, rhkUserID, _ := route.Params.GetOK("user_id")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindListID binds and validates parameter ListID from path.
func (o *RevokeInvitationParams) bindListID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ListID = raw

	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *RevokeInvitationParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UserID = raw

	return nil
}
