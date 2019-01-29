package kv

import (
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "strings"

    "github.com/dghubble/sling"

    "github.com/bullwark-registrar/common"
    "github.com/bullwark-registrar/common/cli"
)

const (
    consulApiPrefix = "/v1/kv/"
)

func AddKeyValuePair(key string, value interface{}) (*http.Response, error) {

    var sb strings.Builder
    var consulConfig = common.Configs[cli.ConsulConfigType].(cli.ConsulConfig)

    var consulUrl = url.URL{
        Scheme: *consulConfig.Scheme,
        Host: fmt.Sprintf("%s:%d", *consulConfig.Host, *consulConfig.Port),
        Path: consulApiPrefix,
    }

    sb.WriteString(consulUrl.String())
    sb.WriteString(strings.Trim(key, "/"))

    // Put key/value as appropriate for provided value
    //
    var responseBody interface{}
    s := sling.New().
        Put(sb.String())

    if vString, ok := value.(string); ok {
        s.Body(strings.NewReader(vString))
    } else if vJson, ok := value.(map[string]interface{}); ok {
        s.BodyJSON(vJson)
    } else {
        return nil, errors.New("received an unsupported value type for " +
            "keystore")
    }

    // Send request
    //
    res, err := s.ReceiveSuccess(responseBody)

    if err != nil {
        return res, err
    }

    if err == nil {
        err = handleBadResponse(res, responseBody)
    }

    return res, err
}

func RemoveKeyValuePair(key string) (*http.Response, error) {

    var sb strings.Builder
    var consulConfig = common.Configs[cli.ConsulConfigType].(cli.ConsulConfig)

    var consulUrl = url.URL{
        Scheme: *consulConfig.Scheme,
        Host: fmt.Sprintf("%s:%d", *consulConfig.Host, *consulConfig.Port),
        Path: consulApiPrefix,
    }

    sb.WriteString(consulUrl.String())
    sb.WriteString(strings.Trim(key, "/"))

    // Put key/value as appropriate for provided value
    //
    var responseBody interface{}
    s := sling.New().
        Delete(sb.String())

    // Send request
    //
    res, err := s.ReceiveSuccess(responseBody)

    if err != nil {
        return res, err
    }

    if err == nil {
        err = handleBadResponse(res, responseBody)
    }

    return res, err
}

func GetKeyValuePair(key string) (*http.Response, interface{}, error) {

    var sb strings.Builder
    var consulConfig = common.Configs[cli.ConsulConfigType].(cli.ConsulConfig)

    var consulUrl = url.URL{
        Scheme: *consulConfig.Scheme,
        Host: fmt.Sprintf("%s:%d", *consulConfig.Host, *consulConfig.Port),
        Path: consulApiPrefix,
    }

    sb.WriteString(consulUrl.String())
    sb.WriteRune('/')
    sb.WriteString(*consulConfig.Prefix)
    sb.WriteRune('/')
    sb.WriteString(key)

    var responseBody interface{}

    // Get key/value
    //
    res, err := sling.New().
        Get(sb.String()).
        ReceiveSuccess(responseBody)

    if err == nil {
        err = handleBadResponse(res, responseBody)
    }

    return res, responseBody, err
}

func handleBadResponse(response *http.Response, responseBody interface{}) error {
    if response.StatusCode / 100 != 2 {
        return errors.New(fmt.Sprintf("%s", responseBody))
    }

    return nil
}
