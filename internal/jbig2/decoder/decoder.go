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

package decoder ;import (_a "github.com/unidoc/unipdf/v3/internal/bitwise";_e "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_b "github.com/unidoc/unipdf/v3/internal/jbig2/document";_g "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_d "image";);func (_cd *Decoder )decodePage (_aae int )([]byte ,error ){const _fa ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";if _aae < 0{return nil ,_g .Errorf (_fa ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_aae );};if _aae > int (_cd ._ef .NumberOfPages ){return nil ,_g .Errorf (_fa ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_aae );};_cca ,_aac :=_cd ._ef .GetPage (_aae );if _aac !=nil {return nil ,_g .Wrap (_aac ,_fa ,"");};_cda ,_aac :=_cca .GetBitmap ();if _aac !=nil {return nil ,_g .Wrap (_aac ,_fa ,"");};_cda .InverseData ();if !_cd ._de .UnpaddedData {return _cda .Data ,nil ;};return _cda .GetUnpaddedData ();};func (_ea *Decoder )DecodePageImage (pageNumber int )(_d .Image ,error ){const _ae ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";_cbg ,_cf :=_ea .decodePageImage (pageNumber );if _cf !=nil {return nil ,_g .Wrap (_cf ,_ae ,"");};return _cbg ,nil ;};func (_f *Decoder )DecodeNextPage ()([]byte ,error ){_f ._cb ++;_cbb :=_f ._cb ;return _f .decodePage (_cbb );};func (_dg *Decoder )PageNumber ()(int ,error ){const _cc ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";if _dg ._ef ==nil {return 0,_g .Error (_cc ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");};return int (_dg ._ef .NumberOfPages ),nil ;};type Parameters struct{UnpaddedData bool ;Color _e .Color ;};func (_ed *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _ed .decodePage (pageNumber )};type Decoder struct{_aa _a .StreamReader ;_ef *_b .Document ;_cb int ;_de Parameters ;};func Decode (input []byte ,parameters Parameters ,globals *_b .Globals )(*Decoder ,error ){_ac :=_a .NewReader (input );_ff ,_cbbd :=_b .DecodeDocument (_ac ,globals );if _cbbd !=nil {return nil ,_cbbd ;};return &Decoder {_aa :_ac ,_ef :_ff ,_de :parameters },nil ;};func (_ce *Decoder )decodePageImage (_dd int )(_d .Image ,error ){const _dc ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _dd < 0{return nil ,_g .Errorf (_dc ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_dd );};if _dd > int (_ce ._ef .NumberOfPages ){return nil ,_g .Errorf (_dc ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_dd );};_deb ,_ceb :=_ce ._ef .GetPage (_dd );if _ceb !=nil {return nil ,_g .Wrap (_ceb ,_dc ,"");};_cg ,_ceb :=_deb .GetBitmap ();if _ceb !=nil {return nil ,_g .Wrap (_ceb ,_dc ,"");};return _cg .ToImage (),nil ;};