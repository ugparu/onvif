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

// Call_GetImagingSettings forwards the call to dev.CallMethod() then parses the payload of the reply as a GetImagingSettingsResponse.
func Call_GetImagingSettings(ctx context.Context, dev *onvif.Device, request imaging.GetImagingSettings) (imaging.GetImagingSettingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetImagingSettingsResponse imaging.GetImagingSettingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetImagingSettingsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetImagingSettings")
		return reply.Body.GetImagingSettingsResponse, errors.Annotate(err, "reply")
	}
}
