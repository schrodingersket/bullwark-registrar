package registration

import (
  "fmt"
  "strings"

  "github.com/bullwark-registrar/common"
  "github.com/bullwark-registrar/common/cli"
  "github.com/bullwark-registrar/kv"
)

type ConsulRegistrar struct {
  Host         string
  Port         int
  Scheme       string
  Prefix       string
}

const ruleName = "pathrule"

// Factory method for creating ConsulRegistrar instances
//
func NewTraefikRegistrar() Registrar {
  r := new(ConsulRegistrar)

  consulConfig, ok := common.Configs[cli.ConsulConfigType].(cli.ConsulConfig)

  if ok {
    r.Host = *consulConfig.Host
    r.Port = *consulConfig.Port
    r.Scheme = *consulConfig.Scheme
    r.Prefix = *consulConfig.Prefix
  }

  return r
}

func (r ConsulRegistrar) Register(req Request) error {

  // Create key URL for backend
  //
  var backendStringBuilder strings.Builder
  backendStringBuilder.WriteRune('/')
  backendStringBuilder.WriteString(r.Prefix)
  backendStringBuilder.WriteString("/backends/")
  backendStringBuilder.WriteString(req.ServiceKey)
  backendStringBuilder.WriteString("/servers/")
  backendStringBuilder.WriteString(req.ServiceId)
  backendStringBuilder.WriteString("/url")

  // Create target endpoint
  //
  var baseUrl, err = HostProviders[req.HostProvider].GetBaseUrl(req)

  if err != nil {
    return err
  }

  _, err = kv.AddKeyValuePair(backendStringBuilder.String(), baseUrl)

  if err != nil {
    return err
  }

  // Front end URL string builder
  //
  var frontendStringBuilder strings.Builder
  frontendStringBuilder.WriteRune('/')
  frontendStringBuilder.WriteString(r.Prefix)
  frontendStringBuilder.WriteString("/frontends/")
  frontendStringBuilder.WriteString(req.ServiceKey)
  frontendStringBuilder.WriteString("/backend")

  _, err = kv.AddKeyValuePair(frontendStringBuilder.String(), req.ServiceKey)


  if err != nil {
    return err
  }

  // Path rule string builder
  //
  var pathRuleStringBuilder strings.Builder
  pathRuleStringBuilder.WriteRune('/')
  pathRuleStringBuilder.WriteString(r.Prefix)
  pathRuleStringBuilder.WriteString("/frontends/")
  pathRuleStringBuilder.WriteString(req.ServiceKey)
  pathRuleStringBuilder.WriteString("/routes/")
  pathRuleStringBuilder.WriteString(ruleName)
  pathRuleStringBuilder.WriteString("/rule")

  _, err = kv.AddKeyValuePair(pathRuleStringBuilder.String(),
    fmt.Sprintf("PathPrefix:/%s", strings.Trim(req.BasePath, "/")))

  if err != nil {
    return err
  }

  // Path rule string builder
  //
  var metadataStringBuilder strings.Builder
  metadataStringBuilder.WriteRune('/')
  metadataStringBuilder.WriteString(r.Prefix)
  metadataStringBuilder.WriteString("/metadata/")
  metadataStringBuilder.WriteString(req.ServiceKey)
  metadataStringBuilder.WriteString("/backends/")
  metadataStringBuilder.WriteString(req.ServiceId)
  metadataStringBuilder.WriteString("/metadata")

  _, err = kv.AddKeyValuePair(metadataStringBuilder.String(), req.Metadata)

  return err
}

func (r ConsulRegistrar) Deregister(req Request) error {

  // Create key URL for backend
  //
  var backendStringBuilder strings.Builder
  backendStringBuilder.WriteRune('/')
  backendStringBuilder.WriteString(r.Prefix)
  backendStringBuilder.WriteString("/backends/")
  backendStringBuilder.WriteString(req.ServiceKey)

  // Can deregister an entire service, or just a particular instance.
  //
  if req.ServiceId != "" {
    backendStringBuilder.WriteString("/servers/")
    backendStringBuilder.WriteString(req.ServiceId)
  }

  _, err := kv.RemoveKeyValuePair(backendStringBuilder.String())

  if err != nil {
    return err
  }

  // Remove front-end if we're removing the entire service
  //
  if req.ServiceId != "" {

    var frontendStringBuilder strings.Builder
    frontendStringBuilder.WriteRune('/')
    frontendStringBuilder.WriteString(r.Prefix)
    frontendStringBuilder.WriteString("/frontends/")
    frontendStringBuilder.WriteString(req.ServiceKey)

    _, err = kv.RemoveKeyValuePair(frontendStringBuilder.String())


    if err != nil {
      return err
    }
  }

  // Path rule string builder
  //
  var metadataStringBuilder strings.Builder
  metadataStringBuilder.WriteRune('/')
  metadataStringBuilder.WriteString(r.Prefix)
  metadataStringBuilder.WriteString("/metadata/")
  metadataStringBuilder.WriteString(req.ServiceKey)

  // Only remove particular service id metadata if serviceId is specified
  //
  if req.ServiceId != "" {
    metadataStringBuilder.WriteString("/backends/")
    metadataStringBuilder.WriteString(req.ServiceId)
  }

  _, err = kv.RemoveKeyValuePair(metadataStringBuilder.String())

  return err
}
