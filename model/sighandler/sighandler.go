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

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_e "bytes";_c "crypto";_aa "crypto/rand";_d "crypto/rsa";_ec "crypto/x509";_dd "crypto/x509/pkix";_fb "encoding/asn1";_fa "errors";_b "fmt";_cf "github.com/unidoc/pkcs7";_df "github.com/unidoc/timestamp";_cg "github.com/unidoc/unipdf/v3/core";
_ed "github.com/unidoc/unipdf/v3/model";_ef "github.com/unidoc/unipdf/v3/model/sigutil";_a "hash";_f "time";);func (_bbe *adobeX509RSASHA1 )sign (_fbc *_ed .PdfSignature ,_gbf _ed .Hasher ,_cbd bool )error {if !_cbd {return _bbe .Sign (_fbc ,_gbf );};_gfa ,_fdb :=_bbe ._gg .PublicKey .(*_d .PublicKey );
if !_fdb {return _b .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",_gfa );};_ecd ,_bgd :=_fb .Marshal (make ([]byte ,_gfa .Size ()));if _bgd !=nil {return _bgd ;
};_fbc .Contents =_cg .MakeHexString (string (_ecd ));return nil ;};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_d .PrivateKey ,certificate *_ec .Certificate )(_ed .SignatureHandler ,error ){return &adobeX509RSASHA1 {_gg :certificate ,_dec :privateKey },nil ;};

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts (certificate *_ec .Certificate ,signFunc SignFunc ,opts *AdobeX509RSASHA1Opts )(_ed .SignatureHandler ,error ){if opts ==nil {opts =&AdobeX509RSASHA1Opts {};};return &adobeX509RSASHA1 {_gg :certificate ,_edg :signFunc ,_cb :opts .EstimateSize },nil ;
};

// InitSignature initialises the PdfSignature.
func (_bd *adobeX509RSASHA1 )InitSignature (sig *_ed .PdfSignature )error {if _bd ._gg ==nil {return _fa .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _bd ._dec ==nil &&_bd ._edg ==nil {return _fa .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_gdc :=*_bd ;sig .Handler =&_gdc ;sig .Filter =_cg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_cg .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");
sig .Cert =_cg .MakeString (string (_gdc ._gg .Raw ));sig .Reference =nil ;_fc ,_gde :=_gdc .NewDigest (sig );if _gde !=nil {return _gde ;};_fc .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _gdc .sign (sig ,_fc ,_bd ._cb );};

// NewDigest creates a new digest.
func (_ab *docTimeStamp )NewDigest (sig *_ed .PdfSignature )(_ed .Hasher ,error ){return _e .NewBuffer (nil ),nil ;};type docTimeStamp struct{_dfa string ;_ggd _c .Hash ;_fbcb int ;};

// InitSignature initialises the PdfSignature.
func (_dfad *docTimeStamp )InitSignature (sig *_ed .PdfSignature )error {_aee :=*_dfad ;sig .Handler =&_aee ;sig .Filter =_cg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_cg .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");
sig .Reference =nil ;if _dfad ._fbcb > 0{sig .Contents =_cg .MakeHexString (string (make ([]byte ,_dfad ._fbcb )));}else {_cbg ,_bdc :=_dfad .NewDigest (sig );if _bdc !=nil {return _bdc ;};_cbg .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
if _bdc =_aee .Sign (sig ,_cbg );_bdc !=nil {return _bdc ;};_dfad ._fbcb =_aee ._fbcb ;};return nil ;};

// Sign sets the Contents fields.
func (_aeb *adobePKCS7Detached )Sign (sig *_ed .PdfSignature ,digest _ed .Hasher )error {if _aeb ._bcb {_fd :=_aeb ._ea ;if _fd <=0{_fd =8192;};sig .Contents =_cg .MakeHexString (string (make ([]byte ,_fd )));return nil ;};_cea :=digest .(*_e .Buffer );
_de ,_ff :=_cf .NewSignedData (_cea .Bytes ());if _ff !=nil {return _ff ;};if _egf :=_de .AddSigner (_aeb ._bc ,_aeb ._bb ,_cf .SignerInfoConfig {});_egf !=nil {return _egf ;};_de .Detach ();_eb ,_ff :=_de .Finish ();if _ff !=nil {return _ff ;};_ba :=make ([]byte ,8192);
copy (_ba ,_eb );sig .Contents =_cg .MakeHexString (string (_ba ));return nil ;};

// Validate validates PdfSignature.
func (_fbf *adobeX509RSASHA1 )Validate (sig *_ed .PdfSignature ,digest _ed .Hasher )(_ed .SignatureValidationResult ,error ){_gaf ,_bg :=_fbf .getCertificate (sig );if _bg !=nil {return _ed .SignatureValidationResult {},_bg ;};_gfg :=sig .Contents .Bytes ();
var _gbg []byte ;if _ ,_ddc :=_fb .Unmarshal (_gfg ,&_gbg );_ddc !=nil {return _ed .SignatureValidationResult {},_ddc ;};_fbb ,_eba :=digest .(_a .Hash );if !_eba {return _ed .SignatureValidationResult {},_fa .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_gcg ,_ :=_def (_gaf .SignatureAlgorithm );if _ecg :=_d .VerifyPKCS1v15 (_gaf .PublicKey .(*_d .PublicKey ),_gcg ,_fbb .Sum (nil ),_gbg );_ecg !=nil {return _ed .SignatureValidationResult {},_ecg ;};return _ed .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;
};

// NewDigest creates a new digest.
func (_bbb *adobePKCS7Detached )NewDigest (sig *_ed .PdfSignature )(_ed .Hasher ,error ){return _e .NewBuffer (nil ),nil ;};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_ed .SignatureHandler ,error ){return &adobePKCS7Detached {_bcb :true ,_ea :signatureLen },nil ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom (certificate *_ec .Certificate ,signFunc SignFunc )(_ed .SignatureHandler ,error ){return &adobeX509RSASHA1 {_gg :certificate ,_edg :signFunc },nil ;};func (_bae *adobeX509RSASHA1 )getCertificate (_gee *_ed .PdfSignature )(*_ec .Certificate ,error ){if _bae ._gg !=nil {return _bae ._gg ,nil ;
};_dfde ,_db :=_gee .GetCerts ();if _db !=nil {return nil ,_db ;};return _dfde [0],nil ;};

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts (timestampServerURL string ,hashAlgorithm _c .Hash ,opts *DocTimeStampOpts )(_ed .SignatureHandler ,error ){if opts ==nil {opts =&DocTimeStampOpts {};};if opts .SignatureSize <=0{opts .SignatureSize =4192;};return &docTimeStamp {_dfa :timestampServerURL ,_ggd :hashAlgorithm ,_fbcb :opts .SignatureSize },nil ;
};

// Validate validates PdfSignature.
func (_ae *adobePKCS7Detached )Validate (sig *_ed .PdfSignature ,digest _ed .Hasher )(_ed .SignatureValidationResult ,error ){_eg :=sig .Contents .Bytes ();_cdb ,_bf :=_cf .Parse (_eg );if _bf !=nil {return _ed .SignatureValidationResult {},_bf ;};_gc :=digest .(*_e .Buffer );
_cdb .Content =_gc .Bytes ();if _bf =_cdb .Verify ();_bf !=nil {return _ed .SignatureValidationResult {},_bf ;};return _ed .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};func _def (_da _ec .SignatureAlgorithm )(_c .Hash ,bool ){var _cfe _c .Hash ;
switch _da {case _ec .SHA1WithRSA :_cfe =_c .SHA1 ;case _ec .SHA256WithRSA :_cfe =_c .SHA256 ;case _ec .SHA384WithRSA :_cfe =_c .SHA384 ;case _ec .SHA512WithRSA :_cfe =_c .SHA512 ;default:return _c .SHA1 ,false ;};return _cfe ,true ;};

// InitSignature initialises the PdfSignature.
func (_ge *adobePKCS7Detached )InitSignature (sig *_ed .PdfSignature )error {if !_ge ._bcb {if _ge ._bc ==nil {return _fa .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _ge ._bb ==nil {return _fa .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_gd :=*_ge ;sig .Handler =&_gd ;sig .Filter =_cg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_cg .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_ce ,_bcc :=_gd .NewDigest (sig );if _bcc !=nil {return _bcc ;};_ce .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _gd .Sign (sig ,_ce );};

// NewDigest creates a new digest.
func (_dfdd *adobeX509RSASHA1 )NewDigest (sig *_ed .PdfSignature )(_ed .Hasher ,error ){_geed ,_baeb :=_dfdd .getCertificate (sig );if _baeb !=nil {return nil ,_baeb ;};_ede ,_ :=_def (_geed .SignatureAlgorithm );return _ede .New (),nil ;};func (_dag *docTimeStamp )getCertificate (_egb *_ed .PdfSignature )(*_ec .Certificate ,error ){_egfb ,_efg :=_egb .GetCerts ();
if _efg !=nil {return nil ,_efg ;};return _egfb [0],nil ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_ga *_ed .PdfSignature ,_gf _ed .Hasher )([]byte ,error );

// Validate validates PdfSignature.
func (_ag *docTimeStamp )Validate (sig *_ed .PdfSignature ,digest _ed .Hasher )(_ed .SignatureValidationResult ,error ){_fba :=sig .Contents .Bytes ();_dbb ,_cge :=_cf .Parse (_fba );if _cge !=nil {return _ed .SignatureValidationResult {},_cge ;};if _cge =_dbb .Verify ();
_cge !=nil {return _ed .SignatureValidationResult {},_cge ;};var _bggc timestampInfo ;_ ,_cge =_fb .Unmarshal (_dbb .Content ,&_bggc );if _cge !=nil {return _ed .SignatureValidationResult {},_cge ;};_ded ,_cge :=_gff (_bggc .MessageImprint .HashAlgorithm .Algorithm );
if _cge !=nil {return _ed .SignatureValidationResult {},_cge ;};_bcbg :=_ded .New ();_ca :=digest .(*_e .Buffer );_bcbg .Write (_ca .Bytes ());_defg :=_bcbg .Sum (nil );_fbcg :=_ed .SignatureValidationResult {IsSigned :true ,IsVerified :_e .Equal (_defg ,_bggc .MessageImprint .HashedMessage ),GeneralizedTime :_bggc .GeneralizedTime };
return _fbcg ,nil ;};

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct{

// SignatureSize is the estimated size of the signature contents in bytes.
// If not provided, a default signature size of 4192 is used.
// The signing process will report the model.ErrSignNotEnoughSpace error
// if the estimated signature size is smaller than the actual size of the
// signature.
SignatureSize int ;};type adobePKCS7Detached struct{_bb *_d .PrivateKey ;_bc *_ec .Certificate ;_bcb bool ;_ea int ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_ffa *adobePKCS7Detached )IsApplicable (sig *_ed .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_agd *docTimeStamp )IsApplicable (sig *_ed .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";
};

// Sign sets the Contents fields for the PdfSignature.
func (_be *docTimeStamp )Sign (sig *_ed .PdfSignature ,digest _ed .Hasher )error {_beb ,_bac :=_ef .NewTimestampRequest (digest .(*_e .Buffer ),&_df .RequestOptions {Hash :_be ._ggd ,Certificates :true });if _bac !=nil {return _bac ;};_bec :=_ef .NewTimestampClient ();
_fgc ,_bac :=_bec .GetEncodedToken (_be ._dfa ,_beb );if _bac !=nil {return _bac ;};_deb :=len (_fgc );if _be ._fbcb > 0&&_deb > _be ._fbcb {return _ed .ErrSignNotEnoughSpace ;};if _deb > 0{_be ._fbcb =_deb +128;};sig .Contents =_cg .MakeHexString (string (_fgc ));
return nil ;};

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_d .PrivateKey ,certificate *_ec .Certificate )(_ed .SignatureHandler ,error ){return &adobePKCS7Detached {_bc :certificate ,_bb :privateKey },nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_gcd *adobeX509RSASHA1 )Sign (sig *_ed .PdfSignature ,digest _ed .Hasher )error {var _bgg []byte ;var _bdf error ;if _gcd ._edg !=nil {_bgg ,_bdf =_gcd ._edg (sig ,digest );if _bdf !=nil {return _bdf ;};}else {_fe ,_gae :=digest .(_a .Hash );if !_gae {return _fa .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_fg ,_ :=_def (_gcd ._gg .SignatureAlgorithm );_bgg ,_bdf =_d .SignPKCS1v15 (_aa .Reader ,_gcd ._dec ,_fg ,_fe .Sum (nil ));if _bdf !=nil {return _bdf ;};};_bgg ,_bdf =_fb .Marshal (_bgg );if _bdf !=nil {return _bdf ;};sig .Contents =_cg .MakeHexString (string (_bgg ));
return nil ;};func _gff (_eaa _fb .ObjectIdentifier )(_c .Hash ,error ){switch {case _eaa .Equal (_cf .OIDDigestAlgorithmSHA1 ),_eaa .Equal (_cf .OIDDigestAlgorithmECDSASHA1 ),_eaa .Equal (_cf .OIDDigestAlgorithmDSA ),_eaa .Equal (_cf .OIDDigestAlgorithmDSASHA1 ),_eaa .Equal (_cf .OIDEncryptionAlgorithmRSA ):return _c .SHA1 ,nil ;
case _eaa .Equal (_cf .OIDDigestAlgorithmSHA256 ),_eaa .Equal (_cf .OIDDigestAlgorithmECDSASHA256 ):return _c .SHA256 ,nil ;case _eaa .Equal (_cf .OIDDigestAlgorithmSHA384 ),_eaa .Equal (_cf .OIDDigestAlgorithmECDSASHA384 ):return _c .SHA384 ,nil ;case _eaa .Equal (_cf .OIDDigestAlgorithmSHA512 ),_eaa .Equal (_cf .OIDDigestAlgorithmECDSASHA512 ):return _c .SHA512 ,nil ;
};return _c .Hash (0),_cf .ErrUnsupportedAlgorithm ;};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _c .Hash )(_ed .SignatureHandler ,error ){return &docTimeStamp {_dfa :timestampServerURL ,_ggd :hashAlgorithm },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_bga *adobeX509RSASHA1 )IsApplicable (sig *_ed .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";
};func (_dfd *adobePKCS7Detached )getCertificate (_ddg *_ed .PdfSignature )(*_ec .Certificate ,error ){if _dfd ._bc !=nil {return _dfd ._bc ,nil ;};_gb ,_cd :=_ddg .GetCerts ();if _cd !=nil {return nil ,_cd ;};return _gb [0],nil ;};type timestampInfo struct{Version int ;
Policy _fb .RawValue ;MessageImprint struct{HashAlgorithm _dd .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _fb .RawValue ;GeneralizedTime _f .Time ;};

// AdobeX509RSASHA1Opts defines options for configuring the adbe.x509.rsa_sha1
// signature handler.
type AdobeX509RSASHA1Opts struct{

// EstimateSize specifies whether the size of the signature contents
// should be estimated based on the modulus size of the public key
// extracted from the signing certificate. If set to false, a mock Sign
// call is made in order to estimate the size of the signature contents.
EstimateSize bool ;};type adobeX509RSASHA1 struct{_dec *_d .PrivateKey ;_gg *_ec .Certificate ;_edg SignFunc ;_cb bool ;};