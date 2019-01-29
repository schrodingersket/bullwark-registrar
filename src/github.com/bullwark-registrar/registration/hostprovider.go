package registration

import (
    "fmt"
    "net/url"
)

type HostProviderType int

const (
    StaticUrlHostProviderType HostProviderType = iota
)

var HostProviderTypeStrings = []string{
    "static",
}

func (t HostProviderType) name() string {
    return HostProviderTypeStrings[t]
}


// This file exposes various host providers for services. The most basic is a
// static provider; i.e., the hostname is manually specified by the client
// service.
//
// Future support is planned for Rancher and Kubernetes providers.
//
type HostProvider interface {
    GetBaseUrl(request Request) (string, error)
}

type StaticHostProvider struct {}

func NewStaticHostProvider() HostProvider {
    return new(StaticHostProvider)
}

func (StaticHostProvider) GetBaseUrl(request Request) (string, error) {

    var staticUrl = url.URL{
        Scheme: request.Scheme,
        Host: fmt.Sprintf("%s:%d", request.Host, request.Port),
        Path: request.BasePath,
    }

    return staticUrl.String(), nil
}

// Define available host providers. Currently only the static provider is
// supported.
//
var HostProviders = map[string]HostProvider{
    StaticUrlHostProviderType.name(): NewStaticHostProvider(),
}
