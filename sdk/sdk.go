package sdk

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juju/errors"
	"github.com/ugparu/onvif/xmlu"
)

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {
	if httpReply.StatusCode/100 != 2 {
		return fmt.Errorf("unexpected status %v instead of 2XX", httpReply.StatusCode)
	}
	if b, err := ioutil.ReadAll(httpReply.Body); err != nil {
		return errors.Annotate(err, "read")
	} else {
		err = xmlu.Unmarshal(b, reply)
		return errors.Annotate(err, "decode")
	}
}
