package requestbuilder

import (
	"github.com/Turing-Chu/huobi_golang/internal/model"
	model2 "github.com/Turing-Chu/huobi_golang/pkg/model"
	"time"
)

type WebSocketV1RequestBuilder struct {
	akKey   string
	akValue string
	smKey   string
	smValue string
	svKey   string
	svValue string
	tKey    string
	tValue  string

	host string
	path string

	signer *Signer
}

func (p *WebSocketV1RequestBuilder) Init(accessKey string, secretKey string, host string, path string) *WebSocketV1RequestBuilder {
	p.akKey = "AccessKeyId"
	p.akValue = accessKey
	p.smKey = "SignatureMethod"
	p.smValue = "HmacSHA256"
	p.svKey = "SignatureVersion"
	p.svValue = "2"
	p.tKey = "Timestamp"

	p.host = host
	p.path = path

	p.signer = new(Signer).Init(secretKey)

	return p
}

func (p *WebSocketV1RequestBuilder) Build() (string, error) {
	time := time.Now().UTC()
	return p.build(time)
}

func (p *WebSocketV1RequestBuilder) build(utcDate time.Time) (string, error) {
	time := utcDate.Format("2006-01-02T15:04:05")

	req := new(model2.GetRequest).Init()
	req.AddParam(p.akKey, p.akValue)
	req.AddParam(p.smKey, p.smValue)
	req.AddParam(p.svKey, p.svValue)
	req.AddParam(p.tKey, time)

	signature := p.signer.Sign("GET", p.host, p.path, req.BuildParams())

	auth := new(model.WebSocketV1AuthenticationRequest).Init()
	auth.AccessKeyId = p.akValue
	auth.Timestamp = time
	auth.Signature = signature

	result, err := model2.ToJson(auth)
	return result, err
}
