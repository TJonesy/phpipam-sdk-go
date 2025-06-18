// Package provider provides types and methods for working with the nameservers
// controller.
package nameservers

import (
	"fmt"

	"github.com/tjonesy/phpipam-sdk-go/phpipam/client"
	"github.com/tjonesy/phpipam-sdk-go/phpipam/session"
)

// Nameserver represents a PHPIPAM subnet.
type Nameserver struct {

	// The subnet address, in dotted quad format (i.e. A.B.C.D).
	Name string `json:"name,omitempty"`

	// The nameservers in dotted quad format (i.e. A.B.C.D) delimited by semicolons.
	NameSrv1 string `json:"namesrv1,omitempty"`

	// A detailed description of the subnet.
	Description string `json:"description,omitempty"`

	// A JSON object, stringified, that represents the permissions for this
	// section.
	Permissions string `json:"permissions,omitempty"`

	// The date of the last edit to this resource.
	EditDate string `json:"editDate,omitempty"`

	// The subnet ID.
	ID int `json:"id,omitempty"`

	// A map[string]interface{} of custom fields to set on the resource. Note
	// that this functionality requires PHPIPAM 1.3 or higher with the "Nest
	// custom fields" flag set on the specific API integration. If this is not
	// enabled, this map will be nil on GETs and POSTs and PATCHes with this
	// field set will fail. Use the explicit custom field functions instead.
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// Controller is the base client for the Nameservers controller.
type Controller struct {
	client.Client
}

// NewController returns a new instance of the client for the Nameservers controller.
func NewController(sess *session.Session) *Controller {
	c := &Controller{
		Client: *client.NewClient(sess),
	}

	return c
}

// CreateNameserver creates a subnet by sending a POST request.
func (c *Controller) CreateNameserver(in Nameserver) (message string, err error) {
	err = c.SendRequest("POST", "/tools/nameservers/", &in, &message)

	return
}

// GetNameserverByID GETs a subnet via its ID.
func (c *Controller) GetNameserverByID(id int) (out Nameserver, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/tools/nameservers/%d/", id), &struct{}{}, &out)

	return
}

// UpdateNameserver updates a subnet by sending a PATCH request.
//
// Note you cannot use this function to update a subnet's CIDR - to split,
// grow, or renumber a subnet, you need to use other methods that are currently
// not implemented in this SDK. See the API spec for more details.
func (c *Controller) UpdateNameserver(in Nameserver) (message string, err error) {
	err = c.SendRequest("PATCH", "/tools/nameservers/", &in, &message)

	return
}

// DeleteNameserver deletes a subnet by its ID.
func (c *Controller) DeleteNameserver(id int) (message string, err error) {
	err = c.SendRequest("DELETE", fmt.Sprintf("/tools/nameservers/%d/", id), &struct{}{}, &message)

	return
}
