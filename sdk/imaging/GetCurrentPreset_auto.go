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

// Call_GetCurrentPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCurrentPresetResponse.
func Call_GetCurrentPreset(ctx context.Context, dev *onvif.Device, request imaging.GetCurrentPreset) (imaging.GetCurrentPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCurrentPresetResponse imaging.GetCurrentPresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCurrentPresetResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCurrentPreset")
		return reply.Body.GetCurrentPresetResponse, errors.Annotate(err, "reply")
	}
}
