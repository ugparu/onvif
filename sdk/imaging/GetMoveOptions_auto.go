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

// Call_GetMoveOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMoveOptionsResponse.
func Call_GetMoveOptions(ctx context.Context, dev *onvif.Device, request imaging.GetMoveOptions) (imaging.GetMoveOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMoveOptionsResponse imaging.GetMoveOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMoveOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetMoveOptions")
		return reply.Body.GetMoveOptionsResponse, errors.Annotate(err, "reply")
	}
}
