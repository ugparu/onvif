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

// Call_SetImagingSettings forwards the call to dev.CallMethod() then parses the payload of the reply as a SetImagingSettingsResponse.
func Call_SetImagingSettings(ctx context.Context, dev *onvif.Device, request imaging.SetImagingSettings) (imaging.SetImagingSettingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetImagingSettingsResponse imaging.SetImagingSettingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetImagingSettingsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetImagingSettings")
		return reply.Body.SetImagingSettingsResponse, errors.Annotate(err, "reply")
	}
}
