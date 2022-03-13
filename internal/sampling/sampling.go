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

package sampling ;import (_ag "github.com/unidoc/unipdf/v3/internal/bitwise";_b "github.com/unidoc/unipdf/v3/internal/imageutil";_d "io";);func (_da *Reader )ReadSamples (samples []uint32 )(_bbg error ){for _ac :=0;_ac < len (samples );_ac ++{samples [_ac ],_bbg =_da .ReadSample ();
if _bbg !=nil {return _bbg ;};};return nil ;};type SampleWriter interface{WriteSample (_cae uint32 )error ;WriteSamples (_gf []uint32 )error ;};func (_dc *Writer )WriteSample (sample uint32 )error {if _ ,_gb :=_dc ._ab .WriteBits (uint64 (sample ),_dc ._ea .BitsPerComponent );
_gb !=nil {return _gb ;};_dc ._fe --;if _dc ._fe ==0{_dc ._fe =_dc ._ea .ColorComponents ;_dc ._ega ++;};if _dc ._ega ==_dc ._ea .Width {if _dc ._feb {_dc ._ab .FinishByte ();};_dc ._ega =0;};return nil ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _gee []uint32 ;
_geg :=bitsPerOutputSample ;var _ce uint32 ;var _bd uint32 ;_ef :=0;_ba :=0;_bfe :=0;for _bfe < len (data ){if _ef > 0{_de :=_ef ;if _geg < _de {_de =_geg ;};_ce =(_ce <<uint (_de ))|(_bd >>uint (bitsPerInputSample -_de ));_ef -=_de ;if _ef > 0{_bd =_bd <<uint (_de );
}else {_bd =0;};_geg -=_de ;if _geg ==0{_gee =append (_gee ,_ce );_geg =bitsPerOutputSample ;_ce =0;_ba ++;};}else {_cb :=data [_bfe ];_bfe ++;_fb :=bitsPerInputSample ;if _geg < _fb {_fb =_geg ;};_ef =bitsPerInputSample -_fb ;_ce =(_ce <<uint (_fb ))|(_cb >>uint (_ef ));
if _fb < bitsPerInputSample {_bd =_cb <<uint (_fb );};_geg -=_fb ;if _geg ==0{_gee =append (_gee ,_ce );_geg =bitsPerOutputSample ;_ce =0;_ba ++;};};};for _ef >=bitsPerOutputSample {_dg :=_ef ;if _geg < _dg {_dg =_geg ;};_ce =(_ce <<uint (_dg ))|(_bd >>uint (bitsPerInputSample -_dg ));
_ef -=_dg ;if _ef > 0{_bd =_bd <<uint (_dg );}else {_bd =0;};_geg -=_dg ;if _geg ==0{_gee =append (_gee ,_ce );_geg =bitsPerOutputSample ;_ce =0;_ba ++;};};if _geg > 0&&_geg < bitsPerOutputSample {_ce <<=uint (_geg );_gee =append (_gee ,_ce );};return _gee ;
};func (_aga *Reader )ReadSample ()(uint32 ,error ){if _aga ._fc ==_aga ._c .Height {return 0,_d .EOF ;};_aa ,_ed :=_aga ._fd .ReadBits (byte (_aga ._c .BitsPerComponent ));if _ed !=nil {return 0,_ed ;};_aga ._ad --;if _aga ._ad ==0{_aga ._ad =_aga ._c .ColorComponents ;
_aga ._bb ++;};if _aga ._bb ==_aga ._c .Width {if _aga ._e {_aga ._fd .ConsumeRemainingBits ();};_aga ._bb =0;_aga ._fc ++;};return uint32 (_aa ),nil ;};func (_acc *Writer )WriteSamples (samples []uint32 )error {for _gbc :=0;_gbc < len (samples );_gbc ++{if _dd :=_acc .WriteSample (samples [_gbc ]);
_dd !=nil {return _dd ;};};return nil ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_f []uint32 )error ;};func NewWriter (img _b .ImageBase )*Writer {return &Writer {_ab :_ag .NewWriterMSB (img .Data ),_ea :img ,_fe :img .ColorComponents ,_feb :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};type Reader struct{_c _b .ImageBase ;_fd *_ag .Reader ;_bb ,_fc ,_ad int ;_e bool ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _eg []uint32 ;_fda :=bitsPerSample ;var _bg uint32 ;var _g byte ;_bgg :=0;_ca :=0;_ec :=0;for _ec < len (data ){if _bgg > 0{_ge :=_bgg ;
if _fda < _ge {_ge =_fda ;};_bg =(_bg <<uint (_ge ))|uint32 (_g >>uint (8-_ge ));_bgg -=_ge ;if _bgg > 0{_g =_g <<uint (_ge );}else {_g =0;};_fda -=_ge ;if _fda ==0{_eg =append (_eg ,_bg );_fda =bitsPerSample ;_bg =0;_ca ++;};}else {_bf :=data [_ec ];_ec ++;
_gg :=8;if _fda < _gg {_gg =_fda ;};_bgg =8-_gg ;_bg =(_bg <<uint (_gg ))|uint32 (_bf >>uint (_bgg ));if _gg < 8{_g =_bf <<uint (_gg );};_fda -=_gg ;if _fda ==0{_eg =append (_eg ,_bg );_fda =bitsPerSample ;_bg =0;_ca ++;};};};for _bgg >=bitsPerSample {_cc :=_bgg ;
if _fda < _cc {_cc =_fda ;};_bg =(_bg <<uint (_cc ))|uint32 (_g >>uint (8-_cc ));_bgg -=_cc ;if _bgg > 0{_g =_g <<uint (_cc );}else {_g =0;};_fda -=_cc ;if _fda ==0{_eg =append (_eg ,_bg );_fda =bitsPerSample ;_bg =0;_ca ++;};};return _eg ;};type Writer struct{_ea _b .ImageBase ;
_ab *_ag .Writer ;_ega ,_fe int ;_feb bool ;};func NewReader (img _b .ImageBase )*Reader {return &Reader {_fd :_ag .NewReader (img .Data ),_c :img ,_ad :img .ColorComponents ,_e :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};