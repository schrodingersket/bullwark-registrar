package registration

type Registrar interface {
    Register(request Request) error
    Deregister(request Request) error
}

type Request struct {
    Port         int                    `json:"port"`
    Host         string                 `json:"host"`
    HostProvider string                 `json:"hostprovider"`
    ServiceKey   string                 `json:"serviceKey"`
    ServiceId    string                 `json:"serviceId"`
    BasePath     string                 `json:"basePath"`
    HostMatch    string                 `json:"hostMatch"`
    Rule         string                 `json:"rule"`
    Public       bool                   `json:"public"`
    Scheme       string                 `json:"scheme"`
    Metadata     map[string]interface{} `json:"metadata"`
}

type Response struct {
    Status    string `json:"status"`
    Reason    string `json:"reason"`
    ServiceId string `json:"serviceId"`
}
