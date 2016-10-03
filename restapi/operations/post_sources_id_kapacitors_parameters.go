package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/influxdata/mrfusion/models"
)

// NewPostSourcesIDKapacitorsParams creates a new PostSourcesIDKapacitorsParams object
// with the default values initialized.
func NewPostSourcesIDKapacitorsParams() PostSourcesIDKapacitorsParams {
	var ()
	return PostSourcesIDKapacitorsParams{}
}

// PostSourcesIDKapacitorsParams contains all the bound params for the post sources ID kapacitors operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostSourcesIDKapacitors
type PostSourcesIDKapacitorsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of the source
	  Required: true
	  In: path
	*/
	ID string
	/*Configuration options for kapacitor
	  In: body
	*/
	Kapacitor *models.Kapacitor
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostSourcesIDKapacitorsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Kapacitor
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("kapacitor", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Kapacitor = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostSourcesIDKapacitorsParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}
