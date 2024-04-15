package afip

/**
*	El WSAA es uno de los webservices de negocio (WSN) de AFIP.
* Tiene la responsabilidad de generar un ticket de acceso (TA) para un WSN de AFIP.
* Para poder generar un ticket de acceso necesito tener un certificado autorizado por afip.
* Un ticket tiene ciertas características a saber:
* 1) Sirve para un único WSN. No puedo acceder al WSFECRED con un ticket para el WSFE.
* 2) Uno para un WSN de homologación no sirve para un WSN de producción y viceversa.
* 3) Para generarlo se necesita una solicitud de acceso firmada por el certificado y convertida a base64. Esta solicitud es un XML, ver el método para más data.
* 4) Tiene una vigencia de 12hs.
 */

type wsaa struct {
	common
	serviceName string
	url         string
}

func (this *wsaa) getTemplateRequest() wsaaTemplateRequest {
	baseReq := this.common.getTemplateRequest()
	req := wsaaTemplateRequest{}
	req.BaseBody = baseReq
	req.Xmlns = "http://wsaa.view.sua.dvadac.desein.afip.gov"
	return req
}

func (w *wsaa) GetUrl() string {
	return w.url
}

func (w *wsaa) getMethodTemplate() afipMethod {
	m := w.common.getMethodTemplate()
	m.SetSoapAction("https://wsaahomo.afip.gov.ar/ws/services/LoginCms/LoginCMS/loginCmsRequest")
	m.SetEndPoint("https://wsaahomo.afip.gov.ar/ws/services/LoginCms")
	m.SetPort(443)
	return m
}

func (w *wsaa) LoginCms(cms *string) (*string, error) {
	var err error
	requestTemplate := w.getTemplateRequest()
	requestTemplate.Body = loginCmsBodyRequest{In0: cms}

	request := w.getMethodTemplate()
	request.SetHttpBodyRequest(requestTemplate)
	_, err = request.giveMeAResponse()

	if err != nil {
		return nil, err
	}

	responseBody := loginCmsBodyResponse{}
	err = request.FillWithResponse(&responseBody)

	ticket := responseBody.Body.LoginCmsResponse.LoginCmsReturn

	return &ticket, err
}
