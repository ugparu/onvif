// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package imaging

import (
	"context"
	"github.com/juju/errors"
	"github.com/ugparu/onvif/lib"
	"github.com/ugparu/onvif/sdk"
	"github.com/ugparu/onvif/imaging"
)

// Call_GetStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStatusResponse.
func Call_GetStatus(ctx context.Context, dev *onvif.Device, request imaging.GetStatus) (imaging.GetStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStatusResponse imaging.GetStatusResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStatusResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStatus")
		return reply.Body.GetStatusResponse, errors.Annotate(err, "reply")
	}
}
