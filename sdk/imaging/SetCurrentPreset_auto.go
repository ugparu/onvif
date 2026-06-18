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

// Call_SetCurrentPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a SetCurrentPresetResponse.
func Call_SetCurrentPreset(ctx context.Context, dev *onvif.Device, request imaging.SetCurrentPreset) (imaging.SetCurrentPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetCurrentPresetResponse imaging.SetCurrentPresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetCurrentPresetResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetCurrentPreset")
		return reply.Body.SetCurrentPresetResponse, errors.Annotate(err, "reply")
	}
}
