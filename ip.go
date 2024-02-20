package validator

import (
	"context"
	"net"
)

type IP struct {
	message               string
	ipv4NotAllowedMessage string
	ipv6NotAllowedMessage string
	allowV4               bool
	allowV6               bool
}

func NewIP() IP {
	return IP{
		message:               "Must be a valid IP address.",
		ipv4NotAllowedMessage: "Must not be an IPv4 address.",
		ipv6NotAllowedMessage: "Must not be an IPv6 address.",
		allowV4:               true,
		allowV6:               true,
	}
}

func (s IP) ValidateValue(_ context.Context, value any) error {
	v, ok := toString(value)
	if !ok {
		return NewResult().WithError(NewValidationError(s.message))
	}

	ip := net.ParseIP(v)
	if ip == nil {
		return NewResult().WithError(NewValidationError(s.message))
	}

	// TODO: implement ipv4 and ipv4 validations

	return nil
}
