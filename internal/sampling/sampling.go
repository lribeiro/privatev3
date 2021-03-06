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

package sampling ;import (_f "github.com/unidoc/unipdf/v3/internal/bitwise";_a "github.com/unidoc/unipdf/v3/internal/imageutil";_c "io";);type Reader struct{_g _a .ImageBase ;_fb *_f .Reader ;_aa ,_fc ,_gb int ;_af bool ;};type Writer struct{_ecd _a .ImageBase ;
_ea *_f .Writer ;_afeg ,_cd int ;_df bool ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_e []uint32 )error ;};func (_fe *Reader )ReadSample ()(uint32 ,error ){if _fe ._fc ==_fe ._g .Height {return 0,_c .EOF ;};_afc ,_ec :=_fe ._fb .ReadBits (byte (_fe ._g .BitsPerComponent ));
if _ec !=nil {return 0,_ec ;};_fe ._gb --;if _fe ._gb ==0{_fe ._gb =_fe ._g .ColorComponents ;_fe ._aa ++;};if _fe ._aa ==_fe ._g .Width {if _fe ._af {_fe ._fb .ConsumeRemainingBits ();};_fe ._aa =0;_fe ._fc ++;};return uint32 (_afc ),nil ;};func (_aba *Writer )WriteSamples (samples []uint32 )error {for _gf :=0;
_gf < len (samples );_gf ++{if _cfd :=_aba .WriteSample (samples [_gf ]);_cfd !=nil {return _cfd ;};};return nil ;};func NewWriter (img _a .ImageBase )*Writer {return &Writer {_ea :_f .NewWriterMSB (img .Data ),_ecd :img ,_cd :img .ColorComponents ,_df :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _gc []uint32 ;_fcb :=bitsPerSample ;var _eg uint32 ;var _cc byte ;_ed :=0;_fbb :=0;_ee :=0;for _ee < len (data ){if _ed > 0{_ab :=_ed ;if _fcb < _ab {_ab =_fcb ;};_eg =(_eg <<uint (_ab ))|uint32 (_cc >>uint (8-_ab ));
_ed -=_ab ;if _ed > 0{_cc =_cc <<uint (_ab );}else {_cc =0;};_fcb -=_ab ;if _fcb ==0{_gc =append (_gc ,_eg );_fcb =bitsPerSample ;_eg =0;_fbb ++;};}else {_aaa :=data [_ee ];_ee ++;_gbc :=8;if _fcb < _gbc {_gbc =_fcb ;};_ed =8-_gbc ;_eg =(_eg <<uint (_gbc ))|uint32 (_aaa >>uint (_ed ));
if _gbc < 8{_cc =_aaa <<uint (_gbc );};_fcb -=_gbc ;if _fcb ==0{_gc =append (_gc ,_eg );_fcb =bitsPerSample ;_eg =0;_fbb ++;};};};for _ed >=bitsPerSample {_afb :=_ed ;if _fcb < _afb {_afb =_fcb ;};_eg =(_eg <<uint (_afb ))|uint32 (_cc >>uint (8-_afb ));
_ed -=_afb ;if _ed > 0{_cc =_cc <<uint (_afb );}else {_cc =0;};_fcb -=_afb ;if _fcb ==0{_gc =append (_gc ,_eg );_fcb =bitsPerSample ;_eg =0;_fbb ++;};};return _gc ;};func (_ege *Writer )WriteSample (sample uint32 )error {if _ ,_fcd :=_ege ._ea .WriteBits (uint64 (sample ),_ege ._ecd .BitsPerComponent );
_fcd !=nil {return _fcd ;};_ege ._cd --;if _ege ._cd ==0{_ege ._cd =_ege ._ecd .ColorComponents ;_ege ._afeg ++;};if _ege ._afeg ==_ege ._ecd .Width {if _ege ._df {_ege ._ea .FinishByte ();};_ege ._afeg =0;};return nil ;};func (_ac *Reader )ReadSamples (samples []uint32 )(_ff error ){for _ecf :=0;
_ecf < len (samples );_ecf ++{samples [_ecf ],_ff =_ac .ReadSample ();if _ff !=nil {return _ff ;};};return nil ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _ecg []uint32 ;_cb :=bitsPerOutputSample ;
var _b uint32 ;var _dc uint32 ;_egg :=0;_ce :=0;_be :=0;for _be < len (data ){if _egg > 0{_eb :=_egg ;if _cb < _eb {_eb =_cb ;};_b =(_b <<uint (_eb ))|(_dc >>uint (bitsPerInputSample -_eb ));_egg -=_eb ;if _egg > 0{_dc =_dc <<uint (_eb );}else {_dc =0;
};_cb -=_eb ;if _cb ==0{_ecg =append (_ecg ,_b );_cb =bitsPerOutputSample ;_b =0;_ce ++;};}else {_dce :=data [_be ];_be ++;_afe :=bitsPerInputSample ;if _cb < _afe {_afe =_cb ;};_egg =bitsPerInputSample -_afe ;_b =(_b <<uint (_afe ))|(_dce >>uint (_egg ));
if _afe < bitsPerInputSample {_dc =_dce <<uint (_afe );};_cb -=_afe ;if _cb ==0{_ecg =append (_ecg ,_b );_cb =bitsPerOutputSample ;_b =0;_ce ++;};};};for _egg >=bitsPerOutputSample {_afd :=_egg ;if _cb < _afd {_afd =_cb ;};_b =(_b <<uint (_afd ))|(_dc >>uint (bitsPerInputSample -_afd ));
_egg -=_afd ;if _egg > 0{_dc =_dc <<uint (_afd );}else {_dc =0;};_cb -=_afd ;if _cb ==0{_ecg =append (_ecg ,_b );_cb =bitsPerOutputSample ;_b =0;_ce ++;};};if _cb > 0&&_cb < bitsPerOutputSample {_b <<=uint (_cb );_ecg =append (_ecg ,_b );};return _ecg ;
};func NewReader (img _a .ImageBase )*Reader {return &Reader {_fb :_f .NewReader (img .Data ),_g :img ,_gb :img .ColorComponents ,_af :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};type SampleWriter interface{WriteSample (_bc uint32 )error ;
WriteSamples (_cf []uint32 )error ;};