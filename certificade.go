package afip

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"os"

	"go.mozilla.org/pkcs7"
	"golang.org/x/crypto/pkcs12"
)

/*
* Un certificado tiene:
* 0 - Un path.
* 1 - Una contraseña.
* 2 - Una clase privada RSA.
* 3 - Un método para firmar contenido.
 */

type pfx struct {
	path     *string
	password *string
	rawData  *[]byte
}

func (p *pfx) GetPath() *string {
	return p.path
}

func (p *pfx) GetPassword() *string {
	return p.password
}

func (p *pfx) Load(path string, password string) error {
	raw, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	p.rawData = &raw
	p.password = &password

	return err
}

/*
* Firma un contenido pasado como argumento con el certificado y la clave privada del certificado cargado previamente con el método Load.
* Retorna el contenido firmado y convertido a base64
 */
func (p *pfx) Sign(content []byte) (*string, error) {
	privateKey, certificade, err := pkcs12.Decode(*p.rawData, *p.password)
	if err != nil {
		return nil, fmt.Errorf("error decodificando el certificado: %s", err.Error())
	}

	rsaPrivateKey, isRsaKey := privateKey.(*rsa.PrivateKey)
	if !isRsaKey {
		return nil, fmt.Errorf("el certificado debe contener una clave privada RSA")
	}

	signedData, err := pkcs7.NewSignedData(content)
	if err != nil {
		return nil, fmt.Errorf("ocurrió un problema al procesar el contenido a firmar: %s", err.Error())
	}

	if err := signedData.AddSigner(certificade, rsaPrivateKey, pkcs7.SignerInfoConfig{}); err != nil {
		return nil, fmt.Errorf("ocurrió un problema al intentar firmar el contenido: %s", err.Error())
	}

	detachedSignature, _ := signedData.Finish()

	contentBase64 := base64.StdEncoding.EncodeToString(detachedSignature)

	return &contentBase64, nil
}

func NewCertificade() *pfx {
	cert := pfx{}
	return &cert
}
