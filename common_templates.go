package afip

import "encoding/xml"

type xmlName string

// Estructura básica del body de la petición
type BaseBody struct {
	XMLName xmlName `xml:"soapenv:Envelope,omitempty"`
	SoapEnv string  `xml:"xmlns:soapenv,attr,omitempty"`
	Header  string  `xml:"soapenv:Header"`
}

type soapFault struct {
	Body struct {
		Fault struct {
			Faultcode   string `xml:"faultcode"`
			Faultstring string `xml:"faultstring"`
			Detail      struct {
				ExceptionName string `xml:"exceptionName"`
				Hostname      string `xml:"hostname"`
			} `xml:"detail"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

type baseRequest struct {
	XMLName xml.Name             `xml:"soapenv:Envelope,omitempty"`
	SoapEnv string               `xml:"xmlns:soapenv,attr,omitempty"`
	Header  string               `xml:"soapenv:Header,omitempty"`
	Xmlns   string               `xml:",omitempty"`
	Body    bodyResponseTemplate `xml:"soapenv:Body,omitempty"`
}

type bodyResponseTemplate interface {
}

type httpHandler interface {
	GetNewRequest() httpRequestHandler
	Request(string, uint16, httpRequestHandler) (httpResponseHandler, error)
}

type httpRequestHandler interface {
	AddHeader(string, string)
	LoadBody([]byte)
	LoadBodyFromString(string)
	SetCharset(string)
	SetVerb(string)
}

type httpResponseHandler interface {
	setBody([]byte)
	GetBody() []byte
	setStatusCode(uint16)
	GetStatusCode() uint16
	StatusOk() bool
	StatusInternalServerError() bool
	StatusServiceUnavailable() bool
	FillWithResponse(*bodyResponseTemplate) error
}
