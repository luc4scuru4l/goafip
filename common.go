package afip

import (
	"encoding/xml"
	"time"
)

/*
* Rutinas que son comunes a todos los WS.
 */

type common struct {
}

func (c *common) getTemplateRequest() BaseBody {
	b := BaseBody{SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/"}
	return b
}

func (c *common) getMethodTemplate() afipMethod {
	m := wsMethod{}
	return &m
}

func CreateAccessRequest(serviceName string) []byte {
	expiration := time.Now().Add(10 * time.Minute)
	generationTime := time.Now().Add(-10 * time.Minute).Format(time.RFC3339)
	expirationTime := expiration.Format(time.RFC3339)

	request := loginTicketRequest{
		Version: "1.0",
		Header: &headerLoginCms{
			UniqueID:       1,
			GenerationTime: generationTime,
			ExpirationTime: expirationTime,
		},
		Service: serviceName,
	}
	requestXml, _ := xml.MarshalIndent(request, " ", "  ")

	return requestXml
}
