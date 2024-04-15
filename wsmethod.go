package afip

import (
	"encoding/xml"
	"fmt"
)

/*
Un método de un WS tiene:
0 - Un soap action
1 - Un endpoint
2 - Un puerto
3 - Un body http
4 - Un body XML
*/
type afipMethod interface {
	SetPort(uint16)
	SetEndPoint(string)
	SetSoapAction(string)
	FillWithResponse(any) error
	SetHttpBodyRequest(any)
	giveMeAResponse() (httpResponseHandler, error)
	GetSoapAction() *string
	GetHttpBodyRequest() *[]byte
	GetEndPoint() *string
	GetPort() *uint16
	SetHttpRequest(httpRequestHandler)
	GetHttpRequest() *httpRequestHandler
}

type wsMethod struct {
	httpHandler     httpHandler
	httpRequest     httpRequestHandler
	soapAction      string
	endPoint        string
	port            uint16
	httpBodyRequest []byte
	xmlBodyRequest  []byte
	bodyRequest     baseRequest
}

func (w *wsMethod) SetPort(newPort uint16) {
	w.port = newPort
}

func (w *wsMethod) SetEndPoint(newEndPoint string) {
	w.endPoint = newEndPoint
}

func (w *wsMethod) SetSoapAction(newSA string) {
	w.soapAction = newSA
}

func (this *wsMethod) FillWithResponse(any) error {
	return nil
}

func (w *wsMethod) SetHttpBodyRequest(newHttpBodyRequest any) {
	req, _ := xml.MarshalIndent(newHttpBodyRequest, " ", "  ")
	w.httpBodyRequest = req
}

func (this *wsMethod) GetSoapAction() *string {
	return &this.soapAction
}

func (this *wsMethod) GetHttpBodyRequest() *[]byte {
	return &this.httpBodyRequest
}

func (this *wsMethod) GetEndPoint() *string {
	return &this.endPoint
}

func (this *wsMethod) GetPort() *uint16 {
	return &this.port
}

func (this *wsMethod) SetHttpRequest(request httpRequestHandler) {
	this.httpRequest = request
}

func (this *wsMethod) GetHttpRequest() *httpRequestHandler {
	return &this.httpRequest
}

func (this *wsMethod) giveMeAResponse() (httpResponseHandler, error) {
	req := this.httpHandler.GetNewRequest()
	req.AddHeader("SOAPAction", *this.GetSoapAction())
	req.LoadBody(*this.GetHttpBodyRequest())
	this.httpRequest = req

	response, err := this.httpHandler.Request(*this.GetEndPoint(), *this.GetPort(), *this.GetHttpRequest())

	if err != nil {
		return nil, err
	}

	switch {
	case response.StatusServiceUnavailable():
		soapFault := soapFault{}
		xml.Unmarshal(response.GetBody(), &soapFault)
		err = fmt.Errorf("%s", soapFault.Body.Fault.Faultstring)
	default:
		err = fmt.Errorf("los servidores de AFIP no se encuentran disponibles en este momento.")
	}

	return response, err
}

func newAfipMethod(httpHandler httpHandler) wsMethod {
	w := wsMethod{}
	w.httpHandler = httpHandler
	return w
}
