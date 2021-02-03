package useragent

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/stripe/stripe-cli/pkg/version"
)

//
// Public functions
//

// GetEncodedStripeUserAgent returns the string to be used as the value for
// the `X-Stripe-Client-User-Agent` HTTP header.
func GetEncodedStripeUserAgent() string {
	return encodedStripeUserAgent
}

// GetEncodedUserAgent returns the string to be used as the value for
// the `User-Agent` HTTP header.
func GetEncodedUserAgent() string {
	return encodedUserAgent
}

//
// Private types
//

// stripeClientUserAgent contains information about the current runtime which
// is serialized and sent in the `X-Stripe-Client-User-Agent` as additional
// debugging information.
type stripeClientUserAgent struct {
	Name      string `json:"name"`
	OS        string `json:"os"`
	Publisher string `json:"publisher"`
	Uname     string `json:"uname"`
	Version   string `json:"version"`
	Term      string `json:"term_program"`
}

//
// Private variables
//

var encodedStripeUserAgent string
var encodedUserAgent string

//
// Private functions
//

func init() {
	initUserAgent()
}

func initUserAgent() {
	termProgram := os.Getenv("TERM_PROGRAM")

	encodedUserAgent = "Stripe/v1 stripe-cli/" + version.Version
	if termProgram != "" {
		encodedUserAgent += "; term_program/" + termProgram
	}

	stripeUserAgent := &stripeClientUserAgent{
		Name:      "stripe-cli",
		Version:   version.Version,
		Publisher: "stripe",
		OS:        runtime.GOOS,
		Uname:     getUname(),
		Term:      termProgram,
	}
	marshaled, err := json.Marshal(stripeUserAgent)
	// Encoding this struct should never be a problem, so we're okay to panic
	// in case it is for some reason.
	if err != nil {
		panic(err)
	}

	encodedStripeUserAgent = string(marshaled)
}
