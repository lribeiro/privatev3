//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sigutil ;import (_e "bytes";_gg "crypto";_ca "crypto/x509";_eg "encoding/asn1";_eb "encoding/pem";_a "errors";_cc "fmt";_dd "github.com/unidoc/timestamp";_cf "github.com/unidoc/unipdf/v3/common";_ee "golang.org/x/crypto/ocsp";_g "io";_ed "io/ioutil";
_ag "net/http";_c "time";);

// Get retrieves the certificate at the specified URL.
func (_ggd *CertClient )Get (url string )(*_ca .Certificate ,error ){if _ggd .HTTPClient ==nil {_ggd .HTTPClient =_bd ();};_f ,_ac :=_ggd .HTTPClient .Get (url );if _ac !=nil {return nil ,_ac ;};defer _f .Body .Close ();_b ,_ac :=_ed .ReadAll (_f .Body );
if _ac !=nil {return nil ,_ac ;};if _eef ,_ :=_eb .Decode (_b );_eef !=nil {_b =_eef .Bytes ;};_aca ,_ac :=_ca .ParseCertificate (_b );if _ac !=nil {return nil ,_ac ;};return _aca ,nil ;};

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_bg *CertClient )IsCA (cert *_ca .Certificate )bool {return cert .IsCA &&_e .Equal (cert .RawIssuer ,cert .RawSubject );};

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct{

// HTTPClient is the HTTP client used to make OCSP requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_ag .Client ;

// Hash is the hash function  used when constructing the OCSP
// requests. If zero, SHA-1 will be used.
Hash _gg .Hash ;};

// NewCRLClient returns a new CRL client.
func NewCRLClient ()*CRLClient {return &CRLClient {HTTPClient :_bd ()}};

// GetIssuer retrieves the issuer of the provided certificate.
func (_ga *CertClient )GetIssuer (cert *_ca .Certificate )(*_ca .Certificate ,error ){for _ ,_ba :=range cert .IssuingCertificateURL {_ef ,_baf :=_ga .Get (_ba );if _baf !=nil {_cf .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076",cert .Subject .CommonName ,_baf );
continue ;};return _ef ,nil ;};return nil ,_cc .Errorf ("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064");};func _bd ()*_ag .Client {return &_ag .Client {Timeout :5*_c .Second }};


// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct{

// HTTPClient is the HTTP client used to make certificate requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_ag .Client ;};

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient ()*OCSPClient {return &OCSPClient {HTTPClient :_bd (),Hash :_gg .SHA1 }};

// NewCertClient returns a new certificate client.
func NewCertClient ()*CertClient {return &CertClient {HTTPClient :_bd ()}};

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_eefe *OCSPClient )MakeRequest (serverURL string ,cert ,issuer *_ca .Certificate )(*_ee .Response ,[]byte ,error ){if _eefe .HTTPClient ==nil {_eefe .HTTPClient =_bd ();};if serverURL ==""{if len (cert .OCSPServer )==0{return nil ,nil ,_a .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073");
};serverURL =cert .OCSPServer [0];};_ea ,_caf :=_ee .CreateRequest (cert ,issuer ,&_ee .RequestOptions {Hash :_eefe .Hash });if _caf !=nil {return nil ,nil ,_caf ;};_ab ,_caf :=_eefe .HTTPClient .Post (serverURL ,"\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074",_e .NewReader (_ea ));
if _caf !=nil {return nil ,nil ,_caf ;};defer _ab .Body .Close ();_eeg ,_caf :=_ed .ReadAll (_ab .Body );if _caf !=nil {return nil ,nil ,_caf ;};if _aeb ,_ :=_eb .Decode (_eeg );_aeb !=nil {_eeg =_aeb .Bytes ;};_cgd ,_caf :=_ee .ParseResponseForCert (_eeg ,cert ,issuer );
if _caf !=nil {return nil ,nil ,_caf ;};return _cgd ,_eeg ,nil ;};

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient ()*TimestampClient {return &TimestampClient {HTTPClient :_bd ()}};

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_cee *TimestampClient )GetEncodedToken (serverURL string ,req *_dd .Request )([]byte ,error ){if serverURL ==""{return nil ,_cc .Errorf ("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c");
};if req ==nil {return nil ,_cc .Errorf ("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c");};_aa ,_cfa :=req .Marshal ();if _cfa !=nil {return nil ,_cfa ;
};_aed :=_cee .HTTPClient ;if _aed ==nil {_aed =_bd ();};_bf ,_cfa :=_aed .Post (serverURL ,"a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079",_e .NewBuffer (_aa ));
if _cfa !=nil {return nil ,_cfa ;};defer _bf .Body .Close ();_abf ,_cfa :=_ed .ReadAll (_bf .Body );if _cfa !=nil {return nil ,_cfa ;};if _bf .StatusCode !=_ag .StatusOK {return nil ,_cc .Errorf ("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064",_bf .StatusCode );
};var _afb struct{Version _eg .RawValue ;Content _eg .RawValue ;};if _ ,_cfa =_eg .Unmarshal (_abf ,&_afb );_cfa !=nil {return nil ,_cfa ;};return _afb .Content .FullBytes ,nil ;};

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct{

// HTTPClient is the HTTP client used to make CRL requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_ag .Client ;};

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct{

// HTTPClient is the HTTP client used to make timestamp requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_ag .Client ;};

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest (body _g .Reader ,opts *_dd .RequestOptions )(*_dd .Request ,error ){if opts ==nil {opts =&_dd .RequestOptions {};};if opts .Hash ==0{opts .Hash =_gg .SHA256 ;};if !opts .Hash .Available (){return nil ,_ca .ErrUnsupportedAlgorithm ;
};_ebd :=opts .Hash .New ();if _ ,_cb :=_g .Copy (_ebd ,body );_cb !=nil {return nil ,_cb ;};return &_dd .Request {HashAlgorithm :opts .Hash ,HashedMessage :_ebd .Sum (nil ),Certificates :opts .Certificates ,TSAPolicyOID :opts .TSAPolicyOID ,Nonce :opts .Nonce },nil ;
};

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_cd *CRLClient )MakeRequest (serverURL string ,cert *_ca .Certificate )([]byte ,error ){if _cd .HTTPClient ==nil {_cd .HTTPClient =_bd ();};if serverURL ==""{if len (cert .CRLDistributionPoints )==0{return nil ,_a .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073");
};serverURL =cert .CRLDistributionPoints [0];};_gb ,_ae :=_cd .HTTPClient .Get (serverURL );if _ae !=nil {return nil ,_ae ;};defer _gb .Body .Close ();_cg ,_ae :=_ed .ReadAll (_gb .Body );if _ae !=nil {return nil ,_ae ;};if _ce ,_ :=_eb .Decode (_cg );
_ce !=nil {_cg =_ce .Bytes ;};return _cg ,nil ;};