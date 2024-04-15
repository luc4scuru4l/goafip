package afip

/*
Plantilla común para todos los métodos del WSAA
*/
type wsaaTemplateRequest struct {
	BaseBody
	Xmlns string `xml:"xmlns:wsaa,attr,omitempty"`
	Body  any    `xml:"soapenv:Body,omitempty"`
}

// Estructura del Body de la solicitud de acceso a un WSN de AFIP
type loginTicketRequest struct {
	XMLName xmlName         `xml:"loginTicketRequest"`
	Version string          `xml:"version,attr"`
	Header  *headerLoginCms `xml:"header,omitempty"`
	Service string          `xml:"service,omitempty"`
}

// Datos de la cabecera de la solicitud de acceso
type headerLoginCms struct {
	Source         string `xml:"source,omitempty"`
	Destination    string `xml:"destination,omitempty"`
	UniqueID       uint32 `xml:"uniqueId,omitempty"`
	GenerationTime string `xml:"generationTime,omitempty"`
	ExpirationTime string `xml:"expirationTime,omitempty"`
}

// Estructura del Body del XML de petición del método LoginCms
type loginCmsBodyRequest struct {
	In0 *string `xml:"wsaa:loginCms>wsaa:in0"`
}

// Estructura del Body de respuesta del método LoginCms
type loginCmsBodyResponse struct {
	Body struct {
		LoginCmsResponse struct {
			LoginCmsReturn string `xml:"loginCmsReturn"`
		} `xml:"loginCmsResponse"`
	} `xml:"Body"`
}

// Estructura del Ticket de acceso de la respuesta de LoginCms
type loginTicketResponse struct {
	Version     string          `xml:"version,attr"`
	Header      *headerLoginCms `xml:"header,omitempty"`
	Credentials struct {
		Token string `xml:"token"`
		Sign  string `xml:"sign"`
	} `xml:"credentials"`
}
