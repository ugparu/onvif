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

// Call_GetOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOptionsResponse.
func Call_GetOptions(ctx context.Context, dev *onvif.Device, request imaging.GetOptions) (imaging.GetOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOptionsResponse imaging.GetOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetOptions")
		return reply.Body.GetOptionsResponse, errors.Annotate(err, "reply")
	}
}
