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

// Call_GetPresets forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetsResponse.
func Call_GetPresets(ctx context.Context, dev *onvif.Device, request imaging.GetPresets) (imaging.GetPresetsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetsResponse imaging.GetPresetsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPresets")
		return reply.Body.GetPresetsResponse, errors.Annotate(err, "reply")
	}
}
