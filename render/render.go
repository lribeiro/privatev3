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

package render ;import (_f "errors";_df "fmt";_ae "github.com/adrg/sysfont";_cae "github.com/unidoc/unipdf/v3/common";_fb "github.com/unidoc/unipdf/v3/contentstream";_ge "github.com/unidoc/unipdf/v3/contentstream/draw";_ca "github.com/unidoc/unipdf/v3/core";
_ad "github.com/unidoc/unipdf/v3/internal/license";_gc "github.com/unidoc/unipdf/v3/internal/transform";_aea "github.com/unidoc/unipdf/v3/model";_bcf "github.com/unidoc/unipdf/v3/render/internal/context";_ef "github.com/unidoc/unipdf/v3/render/internal/context/imagerender";
_bc "golang.org/x/image/draw";_c "image";_gd "image/color";_de "image/draw";_aa "image/jpeg";_b "image/png";_g "math";_d "os";_a "path/filepath";_ee "strings";);func (_ff renderer )renderContentStream (_ecb _bcf .Context ,_cc string ,_da *_aea .PdfPageResources )error {_dgg ,_acg :=_fb .NewContentStreamParser (_cc ).Parse ();
if _acg !=nil {return _acg ;};_adab :=_ecb .TextState ();_adab .GlobalScale =_ff ._baf ;_aaf :=map[string ]*_bcf .TextFont {};_fg :=_ae .NewFinder (&_ae .FinderOpts {Extensions :[]string {"\u002e\u0074\u0074\u0066","\u002e\u0074\u0074\u0063"}});_faf :=_fb .NewContentStreamProcessor (*_dgg );
_faf .AddHandler (_fb .HandlerConditionEnumAllOperands ,"",func (_eae *_fb .ContentStreamOperation ,_eg _fb .GraphicsState ,_fgb *_aea .PdfPageResources )error {_cae .Log .Debug ("\u0050\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0025\u0073",_eae .Operand );
switch _eae .Operand {case "\u0071":_ecb .Push ();case "\u0051":_ecb .Pop ();_adab =_ecb .TextState ();case "\u0063\u006d":if len (_eae .Params )!=6{return _cfd ;};_efg ,_cbg :=_ca .GetNumbersAsFloat (_eae .Params );if _cbg !=nil {return _cbg ;};_aec :=_gc .NewMatrix (_efg [0],_efg [1],_efg [2],_efg [3],_efg [4],_efg [5]);
_cae .Log .Debug ("\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074a\u0074\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078\u003a \u0025\u002b\u0076",_aec );_ecb .SetMatrix (_ecb .Matrix ().Mult (_aec ));case "\u0077":if len (_eae .Params )!=1{return _cfd ;
};_cg ,_dfe :=_ca .GetNumbersAsFloat (_eae .Params );if _dfe !=nil {return _dfe ;};_ecb .SetLineWidth (_cg [0]);case "\u004a":if len (_eae .Params )!=1{return _cfd ;};_acb ,_gaf :=_ca .GetIntVal (_eae .Params [0]);if !_gaf {return _gdc ;};switch _acb {case 0:_ecb .SetLineCap (_bcf .LineCapButt );
case 1:_ecb .SetLineCap (_bcf .LineCapRound );case 2:_ecb .SetLineCap (_bcf .LineCapSquare );default:_cae .Log .Debug ("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u006ee\u0020\u0063\u0061\u0070\u0020\u0073\u0074\u0079\u006c\u0065:\u0020\u0025\u0064",_acb );
return _cfd ;};case "\u006a":if len (_eae .Params )!=1{return _cfd ;};_cga ,_efa :=_ca .GetIntVal (_eae .Params [0]);if !_efa {return _gdc ;};switch _cga {case 0:_ecb .SetLineJoin (_bcf .LineJoinBevel );case 1:_ecb .SetLineJoin (_bcf .LineJoinRound );case 2:_ecb .SetLineJoin (_bcf .LineJoinBevel );
default:_cae .Log .Debug ("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u006e\u0065\u0020\u006a\u006f\u0069\u006e \u0073\u0074\u0079l\u0065:\u0020\u0025\u0064",_cga );return _cfd ;};case "\u004d":if len (_eae .Params )!=1{return _cfd ;};_acf ,_abf :=_ca .GetNumbersAsFloat (_eae .Params );
if _abf !=nil {return _abf ;};_ =_acf ;_cae .Log .Debug ("\u004di\u0074\u0065\u0072\u0020l\u0069\u006d\u0069\u0074\u0020n\u006ft\u0020s\u0075\u0070\u0070\u006f\u0072\u0074\u0065d");case "\u0064":if len (_eae .Params )!=2{return _cfd ;};_cd ,_ed :=_ca .GetArray (_eae .Params [0]);
if !_ed {return _gdc ;};_efac ,_ed :=_ca .GetIntVal (_eae .Params [1]);if !_ed {return _gdc ;};_bf ,_cag :=_ca .GetNumbersAsFloat (_cd .Elements ());if _cag !=nil {return _cag ;};_ecb .SetDash (_bf ...);_ =_efac ;_cae .Log .Debug ("\u004c\u0069n\u0065\u0020\u0064\u0061\u0073\u0068\u0020\u0070\u0068\u0061\u0073\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006frt\u0065\u0064");
case "\u0072\u0069":_cae .Log .Debug ("\u0052\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020i\u006e\u0074\u0065\u006e\u0074\u0020\u006eo\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064");case "\u0069":_cae .Log .Debug ("\u0046\u006c\u0061\u0074\u006e\u0065\u0073\u0073\u0020\u0074\u006f\u006c\u0065\u0072\u0061n\u0063e\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064");
case "\u0067\u0073":if len (_eae .Params )!=1{return _cfd ;};_fbdf ,_aeg :=_ca .GetName (_eae .Params [0]);if !_aeg {return _gdc ;};if _fbdf ==nil {return _cfd ;};_bd ,_aeg :=_fgb .GetExtGState (*_fbdf );if !_aeg {_cae .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006eo\u0074 \u0066i\u006ed\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u003a\u0020\u0025\u0073",*_fbdf );
return _f .New ("\u0072e\u0073o\u0075\u0072\u0063\u0065\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064");};_ece ,_aeg :=_ca .GetDict (_bd );if !_aeg {_cae .Log .Debug ("\u0045\u0052RO\u0052\u003a\u0020c\u006f\u0075\u006c\u0064 ge\u0074 g\u0072\u0061\u0070\u0068\u0069\u0063\u0073 s\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074");
return _gdc ;};_cae .Log .Debug ("G\u0053\u0020\u0064\u0069\u0063\u0074\u003a\u0020\u0025\u0073",_ece .String ());case "\u006d":if len (_eae .Params )!=2{_cae .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e",_cfd );
return nil ;};_bcc ,_bfc :=_ca .GetNumbersAsFloat (_eae .Params );if _bfc !=nil {return _bfc ;};_cae .Log .Debug ("M\u006f\u0076\u0065\u0020\u0074\u006f\u003a\u0020\u0025\u0076",_bcc );_ecb .NewSubPath ();_ecb .MoveTo (_bcc [0],_bcc [1]);case "\u006c":if len (_eae .Params )!=2{_cae .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0073\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e",_cfd );
return nil ;};_ccg ,_ce :=_ca .GetNumbersAsFloat (_eae .Params );if _ce !=nil {return _ce ;};_ecb .LineTo (_ccg [0],_ccg [1]);case "\u0063":if len (_eae .Params )!=6{return _cfd ;};_gdb ,_adf :=_ca .GetNumbersAsFloat (_eae .Params );if _adf !=nil {return _adf ;
};_cae .Log .Debug ("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076",_gdb );_ecb .CubicTo (_gdb [0],_gdb [1],_gdb [2],_gdb [3],_gdb [4],_gdb [5]);case "\u0076","\u0079":if len (_eae .Params )!=4{return _cfd ;
};_dc ,_fc :=_ca .GetNumbersAsFloat (_eae .Params );if _fc !=nil {return _fc ;};_cae .Log .Debug ("\u0043u\u0062\u0069\u0063\u0020\u0062\u0065\u007a\u0069\u0065\u0072\u0020p\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076",_dc );_ecb .QuadraticTo (_dc [0],_dc [1],_dc [2],_dc [3]);
case "\u0068":_ecb .ClosePath ();_ecb .NewSubPath ();case "\u0072\u0065":if len (_eae .Params )!=4{return _cfd ;};_ecg ,_efae :=_ca .GetNumbersAsFloat (_eae .Params );if _efae !=nil {return _efae ;};_ecb .DrawRectangle (_ecg [0],_ecg [1],_ecg [2],_ecg [3]);
_ecb .NewSubPath ();case "\u0053":_fgg ,_geac :=_eg .ColorspaceStroking .ColorToRGB (_eg .ColorStroking );if _geac !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_geac );
return _geac ;};_ccfc ,_abfe :=_fgg .(*_aea .PdfColorDeviceRGB );if !_abfe {_cae .Log .Debug ("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072");return _geac ;};_ecb .SetRGBA (_ccfc .R (),_ccfc .G (),_ccfc .B (),1);
_ecb .Stroke ();case "\u0073":_gdf ,_fdd :=_eg .ColorspaceStroking .ColorToRGB (_eg .ColorStroking );if _fdd !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_fdd );
return _fdd ;};_fad ,_acbb :=_gdf .(*_aea .PdfColorDeviceRGB );if !_acbb {_cae .Log .Debug ("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072");return _fdd ;};_ecb .ClosePath ();_ecb .NewSubPath ();
_ecb .SetRGBA (_fad .R (),_fad .G (),_fad .B (),1);_ecb .Stroke ();case "\u0066","\u0046":_be ,_afc :=_eg .ColorspaceNonStroking .ColorToRGB (_eg .ColorNonStroking );if _afc !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_afc );
return _afc ;};_gag ,_gdbb :=_be .(*_aea .PdfColorDeviceRGB );if !_gdbb {_cae .Log .Debug ("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072");return _afc ;};_ecb .SetRGBA (_gag .R (),_gag .G (),_gag .B (),1);
_ecb .SetFillRule (_bcf .FillRuleWinding );_ecb .Fill ();case "\u0066\u002a":_caed ,_cfda :=_eg .ColorspaceNonStroking .ColorToRGB (_eg .ColorNonStroking );if _cfda !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_cfda );
return _cfda ;};_ag ,_gff :=_caed .(*_aea .PdfColorDeviceRGB );if !_gff {_cae .Log .Debug ("\u0045\u0072\u0072\u006fr \u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006co\u0072");return _cfda ;};_ecb .SetRGBA (_ag .R (),_ag .G (),_ag .B (),1);
_ecb .SetFillRule (_bcf .FillRuleEvenOdd );_ecb .Fill ();case "\u0042":_bbg ,_eb :=_eg .ColorspaceNonStroking .ColorToRGB (_eg .ColorNonStroking );if _eb !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eb );
return _eb ;};_bfcc :=_bbg .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_bfcc .R (),_bfcc .G (),_bfcc .B (),1);_ecb .SetFillRule (_bcf .FillRuleWinding );_ecb .FillPreserve ();_bbg ,_eb =_eg .ColorspaceStroking .ColorToRGB (_eg .ColorStroking );if _eb !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eb );
return _eb ;};_bfcc =_bbg .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_bfcc .R (),_bfcc .G (),_bfcc .B (),1);_ecb .Stroke ();case "\u0042\u002a":_bbe ,_add :=_eg .ColorspaceNonStroking .ColorToRGB (_eg .ColorNonStroking );if _add !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_add );
return _add ;};_gaa :=_bbe .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_gaa .R (),_gaa .G (),_gaa .B (),1);_ecb .SetFillRule (_bcf .FillRuleEvenOdd );_ecb .FillPreserve ();_bbe ,_add =_eg .ColorspaceStroking .ColorToRGB (_eg .ColorStroking );if _add !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_add );
return _add ;};_gaa =_bbe .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_gaa .R (),_gaa .G (),_gaa .B (),1);_ecb .Stroke ();case "\u0062":_bg ,_faa :=_eg .ColorspaceNonStroking .ColorToRGB (_eg .ColorNonStroking );if _faa !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_faa );
return _faa ;};_fga :=_bg .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_fga .R (),_fga .G (),_fga .B (),1);_ecb .ClosePath ();_ecb .NewSubPath ();_ecb .SetFillRule (_bcf .FillRuleWinding );_ecb .FillPreserve ();_bg ,_faa =_eg .ColorspaceStroking .ColorToRGB (_eg .ColorStroking );
if _faa !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_faa );return _faa ;};_fga =_bg .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_fga .R (),_fga .G (),_fga .B (),1);
_ecb .Stroke ();case "\u0062\u002a":_ecb .ClosePath ();_geg ,_fbfd :=_eg .ColorspaceNonStroking .ColorToRGB (_eg .ColorNonStroking );if _fbfd !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_fbfd );
return _fbfd ;};_agg :=_geg .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_agg .R (),_agg .G (),_agg .B (),1);_ecb .NewSubPath ();_ecb .SetFillRule (_bcf .FillRuleEvenOdd );_ecb .FillPreserve ();_geg ,_fbfd =_eg .ColorspaceStroking .ColorToRGB (_eg .ColorStroking );
if _fbfd !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_fbfd );return _fbfd ;};_agg =_geg .(*_aea .PdfColorDeviceRGB );_ecb .SetRGBA (_agg .R (),_agg .G (),_agg .B (),1);
_ecb .Stroke ();case "\u006e":_ecb .ClearPath ();case "\u0057":_ecb .SetFillRule (_bcf .FillRuleWinding );_ecb .ClipPreserve ();case "\u0057\u002a":_ecb .SetFillRule (_bcf .FillRuleEvenOdd );_ecb .ClipPreserve ();case "\u0072\u0067":_acc ,_ecd :=_eg .ColorNonStroking .(*_aea .PdfColorDeviceRGB );
if !_ecd {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorNonStroking );return nil ;};_ecb .SetFillRGBA (_acc .R (),_acc .G (),_acc .B (),1);
case "\u0052\u0047":_aafc ,_bbec :=_eg .ColorStroking .(*_aea .PdfColorDeviceRGB );if !_bbec {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorStroking );
return nil ;};_ecb .SetStrokeRGBA (_aafc .R (),_aafc .G (),_aafc .B (),1);case "\u006b":_addb ,_dag :=_eg .ColorNonStroking .(*_aea .PdfColorDeviceCMYK );if !_dag {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorNonStroking );
return nil ;};_dfd ,_cff :=_eg .ColorspaceNonStroking .ColorToRGB (_addb );if _cff !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorNonStroking );
return nil ;};_acfd ,_dag :=_dfd .(*_aea .PdfColorDeviceRGB );if !_dag {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_dfd );return nil ;
};_ecb .SetFillRGBA (_acfd .R (),_acfd .G (),_acfd .B (),1);case "\u004b":_fe ,_bda :=_eg .ColorStroking .(*_aea .PdfColorDeviceCMYK );if !_bda {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorStroking );
return nil ;};_fcd ,_ggb :=_eg .ColorspaceStroking .ColorToRGB (_fe );if _ggb !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorStroking );
return nil ;};_eec ,_bda :=_fcd .(*_aea .PdfColorDeviceRGB );if !_bda {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_fcd );return nil ;
};_ecb .SetStrokeRGBA (_eec .R (),_eec .G (),_eec .B (),1);case "\u0067":_age ,_db :=_eg .ColorNonStroking .(*_aea .PdfColorDeviceGray );if !_db {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorNonStroking );
return nil ;};_gca ,_bed :=_eg .ColorspaceNonStroking .ColorToRGB (_age );if _bed !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorNonStroking );
return nil ;};_cbc ,_db :=_gca .(*_aea .PdfColorDeviceRGB );if !_db {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_gca );return nil ;
};_ecb .SetFillRGBA (_cbc .R (),_cbc .G (),_cbc .B (),1);case "\u0047":_ffe ,_bbgd :=_eg .ColorStroking .(*_aea .PdfColorDeviceGray );if !_bbgd {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorStroking );
return nil ;};_bdf ,_aegb :=_eg .ColorspaceStroking .ColorToRGB (_ffe );if _aegb !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorStroking );
return nil ;};_eff ,_bbgd :=_bdf .(*_aea .PdfColorDeviceRGB );if !_bbgd {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_bdf );return nil ;
};_ecb .SetStrokeRGBA (_eff .R (),_eff .G (),_eff .B (),1);case "\u0063\u0073","\u0073\u0063","\u0073\u0063\u006e":_ggg ,_ffef :=_eg .ColorspaceNonStroking .ColorToRGB (_eg .ColorNonStroking );if _ffef !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorNonStroking );
return nil ;};_adb ,_faca :=_ggg .(*_aea .PdfColorDeviceRGB );if !_faca {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_ggg );return nil ;
};_ecb .SetFillRGBA (_adb .R (),_adb .G (),_adb .B (),1);case "\u0043\u0053","\u0053\u0043","\u0053\u0043\u004e":_dbe ,_effc :=_eg .ColorspaceStroking .ColorToRGB (_eg .ColorStroking );if _effc !=nil {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_eg .ColorStroking );
return nil ;};_caa ,_gcf :=_dbe .(*_aea .PdfColorDeviceRGB );if !_gcf {_cae .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0063\u006f\u006e\u0076\u0065r\u0074\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072:\u0020\u0025\u0076",_dbe );return nil ;
};_ecb .SetStrokeRGBA (_caa .R (),_caa .G (),_caa .B (),1);case "\u0044\u006f":if len (_eae .Params )!=1{return _cfd ;};_fbfa ,_aed :=_ca .GetName (_eae .Params [0]);if !_aed {return _gdc ;};_ ,_acfg :=_fgb .GetXObjectByName (*_fbfa );switch _acfg {case _aea .XObjectTypeImage :_cae .Log .Debug ("\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067e\u003a\u0020\u0025\u0073",_fbfa .String ());
_bdg ,_egb :=_fgb .GetXObjectImageByName (*_fbfa );if _egb !=nil {return _egb ;};_bdd ,_egb :=_bdg .ToImage ();if _egb !=nil {return _egb ;};if _eceg :=_bdg .ColorSpace ;_eceg !=nil {var _eeg bool ;switch _eceg .(type ){case *_aea .PdfColorspaceSpecialIndexed :_eeg =true ;
};if _eeg {if _dce ,_bgg :=_eceg .ImageToRGB (*_bdd );_bgg !=nil {_cae .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u006fnv\u0065r\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0052G\u0042\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020i\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e");
}else {_bdd =&_dce ;};};};_eag :=_ecb .FillPattern ().ColorAt (0,0);var _ded _c .Image ;if _bdg .Mask !=nil {if _ded ,_egb =_cdf (_bdg .Mask ,_eag );_egb !=nil {_cae .Log .Debug ("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e");
};};var _dae _c .Image ;if _edf ,_ :=_ca .GetBoolVal (_bdg .ImageMask );_edf {_dae =_ccaa (_bdd ,_eag );}else {_dae ,_egb =_bdd .ToGoImage ();if _egb !=nil {return _egb ;};};if _ded !=nil {_dae =_acfde (_dae ,_ded );};_cfa :=_dae .Bounds ();_ecb .Push ();
_ecb .Scale (1.0/float64 (_cfa .Dx ()),-1.0/float64 (_cfa .Dy ()));_ecb .DrawImageAnchored (_dae ,0,0,0,1);_ecb .Pop ();case _aea .XObjectTypeForm :_cae .Log .Debug ("\u0058\u004fb\u006a\u0065\u0063t\u0020\u0066\u006f\u0072\u006d\u003a\u0020\u0025\u0073",_fbfa .String ());
_fbdc ,_cef :=_fgb .GetXObjectFormByName (*_fbfa );if _cef !=nil {return _cef ;};_ede ,_cef :=_fbdc .GetContentStream ();if _cef !=nil {return _cef ;};_bfce :=_fbdc .Resources ;if _bfce ==nil {_bfce =_fgb ;};_ecb .Push ();if _fbdc .Matrix !=nil {_cgc ,_efd :=_ca .GetArray (_fbdc .Matrix );
if !_efd {return _gdc ;};_bga ,_cefe :=_ca .GetNumbersAsFloat (_cgc .Elements ());if _cefe !=nil {return _cefe ;};if len (_bga )!=6{return _cfd ;};_gfb :=_gc .NewMatrix (_bga [0],_bga [1],_bga [2],_bga [3],_bga [4],_bga [5]);_ecb .SetMatrix (_ecb .Matrix ().Mult (_gfb ));
};if _fbdc .BBox !=nil {_ddd ,_gge :=_ca .GetArray (_fbdc .BBox );if !_gge {return _gdc ;};_gde ,_gdcc :=_ca .GetNumbersAsFloat (_ddd .Elements ());if _gdcc !=nil {return _gdcc ;};if len (_gde )!=4{_cae .Log .Debug ("\u004c\u0065\u006e\u0020\u003d\u0020\u0025\u0064",len (_gde ));
return _cfd ;};_ecb .DrawRectangle (_gde [0],_gde [1],_gde [2]-_gde [0],_gde [3]-_gde [1]);_ecb .SetRGBA (1,0,0,1);_ecb .Clip ();}else {_cae .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0052\u0065q\u0075\u0069\u0072e\u0064\u0020\u0042\u0042\u006f\u0078\u0020\u006d\u0069ss\u0069\u006e\u0067 \u006f\u006e \u0058\u004f\u0062\u006a\u0065\u0063t\u0020\u0046o\u0072\u006d");
};_cef =_ff .renderContentStream (_ecb ,string (_ede ),_bfce );if _cef !=nil {return _cef ;};_ecb .Pop ();};case "\u0042\u0049":if len (_eae .Params )!=1{return _cfd ;};_dadf ,_dfeg :=_eae .Params [0].(*_fb .ContentStreamInlineImage );if !_dfeg {return nil ;
};_bge ,_cfg :=_dadf .ToImage (_fgb );if _cfg !=nil {return _cfg ;};_bddc ,_cfg :=_bge .ToGoImage ();if _cfg !=nil {return _cfg ;};_beda :=_bddc .Bounds ();_ecb .Push ();_ecb .Scale (1.0/float64 (_beda .Dx ()),-1.0/float64 (_beda .Dy ()));_ecb .DrawImageAnchored (_bddc ,0,0,0,1);
_ecb .Pop ();case "\u0042\u0054":_adab .Reset ();case "\u0045\u0054":_adab .Reset ();case "\u0054\u0072":if len (_eae .Params )!=1{return _cfd ;};_aaa ,_cbgc :=_ca .GetNumberAsFloat (_eae .Params [0]);if _cbgc !=nil {return _cbgc ;};_adab .Tr =_bcf .TextRenderingMode (_aaa );
case "\u0054\u004c":if len (_eae .Params )!=1{return _cfd ;};_ceg ,_accb :=_ca .GetNumberAsFloat (_eae .Params [0]);if _accb !=nil {return _accb ;};_adab .Tl =_ceg ;case "\u0054\u0063":if len (_eae .Params )!=1{return _cfd ;};_bac ,_eab :=_ca .GetNumberAsFloat (_eae .Params [0]);
if _eab !=nil {return _eab ;};_cae .Log .Debug ("\u0054\u0063\u003a\u0020\u0025\u0076",_bac );_adab .Tc =_bac ;case "\u0054\u0077":if len (_eae .Params )!=1{return _cfd ;};_fff ,_dac :=_ca .GetNumberAsFloat (_eae .Params [0]);if _dac !=nil {return _dac ;
};_cae .Log .Debug ("\u0054\u0077\u003a\u0020\u0025\u0076",_fff );_adab .Tw =_fff ;case "\u0054\u007a":if len (_eae .Params )!=1{return _cfd ;};_ffg ,_cffe :=_ca .GetNumberAsFloat (_eae .Params [0]);if _cffe !=nil {return _cffe ;};_adab .Th =_ffg ;case "\u0054\u0073":if len (_eae .Params )!=1{return _cfd ;
};_ega ,_fed :=_ca .GetNumberAsFloat (_eae .Params [0]);if _fed !=nil {return _fed ;};_adab .Ts =_ega ;case "\u0054\u0064":if len (_eae .Params )!=2{return _cfd ;};_dca ,_gga :=_ca .GetNumbersAsFloat (_eae .Params );if _gga !=nil {return _gga ;};_cae .Log .Debug ("\u0054\u0064\u003a\u0020\u0025\u0076",_dca );
_adab .ProcTd (_dca [0],_dca [1]);case "\u0054\u0044":if len (_eae .Params )!=2{return _cfd ;};_dcc ,_afe :=_ca .GetNumbersAsFloat (_eae .Params );if _afe !=nil {return _afe ;};_cae .Log .Debug ("\u0054\u0044\u003a\u0020\u0025\u0076",_dcc );_adab .ProcTD (_dcc [0],_dcc [1]);
case "\u0054\u002a":_adab .ProcTStar ();case "\u0054\u006d":if len (_eae .Params )!=6{return _cfd ;};_bec ,_eee :=_ca .GetNumbersAsFloat (_eae .Params );if _eee !=nil {return _eee ;};_cae .Log .Debug ("\u0054\u0065x\u0074\u0020\u006da\u0074\u0072\u0069\u0078\u003a\u0020\u0025\u002b\u0076",_bec );
_adab .ProcTm (_bec [0],_bec [1],_bec [2],_bec [3],_bec [4],_bec [5]);case "\u0027":if len (_eae .Params )!=1{return _cfd ;};_deb ,_eef :=_ca .GetStringBytes (_eae .Params [0]);if !_eef {return _gdc ;};_cae .Log .Debug ("\u0027\u0020\u0073t\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_deb ));
_adab .ProcQ (_deb ,_ecb );case "\u0022":if len (_eae .Params )!=3{return _cfd ;};_eabd ,_gb :=_ca .GetNumberAsFloat (_eae .Params [0]);if _gb !=nil {return _gb ;};_cca ,_gb :=_ca .GetNumberAsFloat (_eae .Params [1]);if _gb !=nil {return _gb ;};_aff ,_gba :=_ca .GetStringBytes (_eae .Params [2]);
if !_gba {return _gdc ;};_adab .ProcDQ (_aff ,_eabd ,_cca ,_ecb );case "\u0054\u006a":if len (_eae .Params )!=1{return _cfd ;};_bgf ,_facc :=_ca .GetStringBytes (_eae .Params [0]);if !_facc {return _gdc ;};_cae .Log .Debug ("\u0054j\u0020s\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0060\u0025\u0073\u0060",string (_bgf ));
_adab .ProcTj (_bgf ,_ecb );case "\u0054\u004a":if len (_eae .Params )!=1{return _cfd ;};_aged ,_dbd :=_ca .GetArray (_eae .Params [0]);if !_dbd {_cae .Log .Debug ("\u0054\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_aged );return _gdc ;};_cae .Log .Debug ("\u0054\u004a\u0020\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u002b\u0076",_aged );
for _ ,_acd :=range _aged .Elements (){switch _eac :=_acd .(type ){case *_ca .PdfObjectString :if _eac !=nil {_adab .ProcTj (_eac .Bytes (),_ecb );};case *_ca .PdfObjectFloat ,*_ca .PdfObjectInteger :_eeb ,_caf :=_ca .GetNumberAsFloat (_eac );if _caf ==nil {_adab .Translate (-_eeb *0.001*_adab .Tf .Size *_adab .Th /100.0,0);
};};};case "\u0054\u0066":if len (_eae .Params )!=2{return _cfd ;};_cae .Log .Debug ("\u0025\u0023\u0076",_eae .Params );_bgeg ,_bea :=_ca .GetName (_eae .Params [0]);if !_bea ||_bgeg ==nil {_cae .Log .Debug ("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006e\u0061m\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076",_eae .Params [0]);
return _gdc ;};_cae .Log .Debug ("\u0046\u006f\u006e\u0074\u0020\u006e\u0061\u006d\u0065\u003a\u0020\u0025\u0073",_bgeg .String ());_cbf ,_ecf :=_ca .GetNumberAsFloat (_eae .Params [1]);if _ecf !=nil {_cae .Log .Debug ("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0069z\u0065 \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076",_eae .Params [1]);
return _gdc ;};_cae .Log .Debug ("\u0046\u006f\u006e\u0074\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0076",_cbf );_faff ,_eeba :=_fgb .GetFontByName (*_bgeg );if !_eeba {_cae .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0046\u006f\u006e\u0074\u0020\u0025s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064",_bgeg .String ());
return _f .New ("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};_cae .Log .Debug ("\u0046\u006f\u006e\u0074\u003a\u0020\u0025\u0054",_faff );_efgg ,_bea :=_ca .GetDict (_faff );if !_bea {_cae .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0067e\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074");
return _gdc ;};_dfa ,_ecf :=_aea .NewPdfFontFromPdfObject (_efgg );if _ecf !=nil {_cae .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074");
return _ecf ;};_cec :=_dfa .BaseFont ();if _cec ==""{_cec =_bgeg .String ();};_bbd ,_bea :=_aaf [_cec ];if !_bea {_bbd ,_ecf =_bcf .NewTextFont (_dfa ,_cbf );if _ecf !=nil {_cae .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ecf );};};if _bbd ==nil {if len (_cec )> 7&&_cec [6]=='+'{_cec =_cec [7:];
};_caae :=[]string {_cec ,"\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e","\u0041\u0072\u0069a\u006c","D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"};for _ ,_eagb :=range _caae {_cae .Log .Debug ("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060",_eagb );
if _bbd ,_bea =_aaf [_eagb ];_bea {break ;};_dedg :=_fg .Match (_eagb );if _dedg ==nil {_cae .Log .Debug ("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073",_eagb );
continue ;};_bbd ,_ecf =_bcf .NewTextFontFromPath (_dedg .Filename ,_cbf );if _ecf !=nil {_cae .Log .Debug ("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073",_dedg .Filename );
continue ;};_cae .Log .Debug ("\u0053\u0075\u0062\u0073\u0074\u0069t\u0075\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073 \u0077\u0069\u0074\u0068\u0020\u0025\u0073 \u0028\u0025\u0073\u0029",_cec ,_dedg .Name ,_dedg .Filename );
_aaf [_eagb ]=_bbd ;break ;};};if _bbd ==nil {_cae .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020n\u006f\u0074\u0020\u0066\u0069\u006ed\u0020\u0061\u006e\u0079\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065 \u0066\u006f\u006e\u0074");
return _f .New ("\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0066\u0069\u006e\u0064\u0020a\u006ey\u0020\u0073\u0075\u0069\u0074\u0061\u0062\u006c\u0065\u0020\u0066\u006f\u006e\u0074");};_adab .ProcTf (_bbd .WithSize (_cbf ,_dfa ));case "\u0042\u004d\u0043","\u0042\u0044\u0043":case "\u0045\u004d\u0043":default:_cae .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u006f\u0070\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u0073",_eae .Operand );
};return nil ;});_acg =_faf .Process (_da );if _acg !=nil {return _acg ;};return nil ;};func _ccaa (_ccae *_aea .Image ,_dcb _gd .Color )_c .Image {_fca ,_gae :=int (_ccae .Width ),int (_ccae .Height );_gef :=_c .NewRGBA (_c .Rect (0,0,_fca ,_gae ));for _cfb :=0;
_cfb < _gae ;_cfb ++{for _aca :=0;_aca < _fca ;_aca ++{_dcd ,_eebd :=_ccae .ColorAt (_aca ,_cfb );if _eebd !=nil {_cae .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e",_aca ,_cfb );
continue ;};_aef ,_bedf ,_fef ,_ :=_dcd .RGBA ();var _bca _gd .Color ;if _aef +_bedf +_fef ==0{_bca =_dcb ;}else {_bca =_gd .Transparent ;};_gef .Set (_aca ,_cfb ,_bca );};};return _gef ;};

// Render converts the specified PDF page into an image and returns the result.
func (_cf *ImageDevice )Render (page *_aea .PdfPage )(_c .Image ,error ){_cb ,_fbf :=page .GetMediaBox ();if _fbf !=nil {return nil ,_fbf ;};_cb .Normalize ();_gf :=page .CropBox ;var _bcg ,_ec float64 ;if _gf !=nil {_gf .Normalize ();_bcg ,_ec =_gf .Width (),_gf .Height ();
};_gg :=page .Rotate ;_dd ,_ade ,_ea ,_cac :=_cb .Llx ,_cb .Lly ,_cb .Width (),_cb .Height ();_cbd :=_gc .IdentityMatrix ();if _gg !=nil &&*_gg %360!=0&&*_gg %90==0{_ac :=-float64 (*_gg );_efc :=_daeg (_ea ,_cac ,_ac );_cbd =_cbd .Translate ((_efc .Width -_ea )/2+_ea /2,(_efc .Height -_cac )/2+_cac /2).Rotate (_ac *_g .Pi /180).Translate (-_ea /2,-_cac /2);
_ea ,_cac =_efc .Width ,_efc .Height ;if _gf !=nil {_fd :=_daeg (_bcg ,_ec ,_ac );_bcg ,_ec =_fd .Width ,_fd .Height ;};};if _dd !=0||_ade !=0{_cbd =_cbd .Translate (-_dd ,-_ade );};_cf ._baf =1.0;if _cf .OutputWidth !=0{_fa :=_ea ;if _gf !=nil {_fa =_bcg ;
};_cf ._baf =float64 (_cf .OutputWidth )/_fa ;_ea ,_cac ,_bcg ,_ec =_ea *_cf ._baf ,_cac *_cf ._baf ,_bcg *_cf ._baf ,_ec *_cf ._baf ;_cbd =_gc .ScaleMatrix (_cf ._baf ,_cf ._baf ).Mult (_cbd );};_ga :=_ef .NewContext (int (_ea ),int (_cac ));if _adc :=_cf .renderPage (_ga ,page ,_cbd );
_adc !=nil {return nil ,_adc ;};_af :=_ga .Image ();if _gf !=nil {_efb ,_dfc :=(_gf .Llx -_dd )*_cf ._baf ,(_gf .Lly -_ade )*_cf ._baf ;_dg :=_c .Rect (0,0,int (_bcg ),int (_ec ));_ba :=_c .Pt (int (_efb ),int (_cac -_dfc -_ec ));_ab :=_c .NewRGBA (_dg );
_de .Draw (_ab ,_dg ,_af ,_ba ,_de .Src );_af =_ab ;};return _af ,nil ;};

// NewImageDevice returns a new image device.
func NewImageDevice ()*ImageDevice {const _aae ="r\u0065\u006e\u0064\u0065r.\u004ee\u0077\u0049\u006d\u0061\u0067e\u0044\u0065\u0076\u0069\u0063\u0065";_ad .TrackUse (_aae );return &ImageDevice {};};

// ImageDevice is used to render PDF pages to image targets.
type ImageDevice struct{renderer ;

// OutputWidth represents the width of the rendered images in pixels.
// The heights of the output images are calculated based on the selected
// width and the original height of each rendered page.
OutputWidth int ;};func (_gfg renderer )renderPage (_caeb _bcf .Context ,_ddg *_aea .PdfPage ,_ada _gc .Matrix )error {_fac ,_fbd :=_ddg .GetAllContentStreams ();if _fbd !=nil {return _fbd ;};if _aab :=_ada ;!_aab .Identity (){_fac =_df .Sprintf ("%\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066 \u0025\u002e\u0032\u0066\u0020\u0025\u002e2\u0066\u0020\u0025\u002e\u0032\u0066\u0020\u0063\u006d\u0020%\u0073",_aab [0],_aab [1],_aab [3],_aab [4],_aab [6],_aab [7],_fac );
};_caeb .Translate (0,float64 (_caeb .Height ()));_caeb .Scale (1,-1);_caeb .Push ();_caeb .SetRGBA (1,1,1,1);_caeb .DrawRectangle (0,0,float64 (_caeb .Width ()),float64 (_caeb .Height ()));_caeb .Fill ();_caeb .Pop ();_caeb .SetLineWidth (1.0);_caeb .SetRGBA (0,0,0,1);
return _gfg .renderContentStream (_caeb ,_fac ,_ddg .Resources );};func _ggf (_edfe string ,_cce _c .Image )error {_dab ,_gda :=_d .Create (_edfe );if _gda !=nil {return _gda ;};defer _dab .Close ();return _b .Encode (_dab ,_cce );};func _gbd (_fedd string ,_bdff _c .Image ,_gdec int )error {_aee ,_ddgc :=_d .Create (_fedd );
if _ddgc !=nil {return _ddgc ;};defer _aee .Close ();return _aa .Encode (_aee ,_bdff ,&_aa .Options {Quality :_gdec });};var (_gdc =_f .New ("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072");_cfd =_f .New ("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072");
);func _cdf (_bafa _ca .PdfObject ,_bde _gd .Color )(_c .Image ,error ){_dggb ,_fgbd :=_ca .GetStream (_bafa );if !_fgbd {return nil ,nil ;};_bbde ,_aece :=_aea .NewXObjectImageFromStream (_dggb );if _aece !=nil {return nil ,_aece ;};_dga ,_aece :=_bbde .ToImage ();
if _aece !=nil {return nil ,_aece ;};return _ccaa (_dga ,_bde ),nil ;};

// RenderToPath converts the specified PDF page into an image and saves the
// result at the specified location.
func (_dgc *ImageDevice )RenderToPath (page *_aea .PdfPage ,outputPath string )error {_def ,_gec :=_dgc .Render (page );if _gec !=nil {return _gec ;};_bb :=_ee .ToLower (_a .Ext (outputPath ));if _bb ==""{return _f .New ("\u0063\u006ful\u0064\u0020\u006eo\u0074\u0020\u0072\u0065cog\u006eiz\u0065\u0020\u006f\u0075\u0074\u0070\u0075t \u0066\u0069\u006c\u0065\u0020\u0074\u0079p\u0065");
};switch _bb {case "\u002e\u0070\u006e\u0067":return _ggf (outputPath ,_def );case "\u002e\u006a\u0070\u0067","\u002e\u006a\u0070e\u0067":return _gbd (outputPath ,_def ,100);};return _df .Errorf ("\u0075\u006e\u0072\u0065\u0063\u006fg\u006e\u0069\u007a\u0065\u0064\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020f\u0069\u006c\u0065\u0020\u0074\u0079\u0070e\u003a\u0020\u0025\u0073",_bb );
};type renderer struct{_baf float64 };func _daeg (_ebc ,_bgfe ,_fafb float64 )_ge .BoundingBox {return _ge .Path {Points :[]_ge .Point {_ge .NewPoint (0,0).Rotate (_fafb ),_ge .NewPoint (_ebc ,0).Rotate (_fafb ),_ge .NewPoint (0,_bgfe ).Rotate (_fafb ),_ge .NewPoint (_ebc ,_bgfe ).Rotate (_fafb )}}.GetBoundingBox ();
};func _acfde (_dgb ,_caede _c .Image )_c .Image {_fee ,_ebf :=_caede .Bounds ().Size (),_dgb .Bounds ().Size ();_fgc ,_bafaa :=_fee .X ,_fee .Y ;if _ebf .X > _fgc {_fgc =_ebf .X ;};if _ebf .Y > _bafaa {_bafaa =_ebf .Y ;};_fadf :=_c .Rect (0,0,_fgc ,_bafaa );
if _fee .X !=_fgc ||_fee .Y !=_bafaa {_efcf :=_c .NewRGBA (_fadf );_bc .BiLinear .Scale (_efcf ,_fadf ,_dgb ,_caede .Bounds (),_bc .Over ,nil );_caede =_efcf ;};if _ebf .X !=_fgc ||_ebf .Y !=_bafaa {_fdc :=_c .NewRGBA (_fadf );_bc .BiLinear .Scale (_fdc ,_fadf ,_dgb ,_dgb .Bounds (),_bc .Over ,nil );
_dgb =_fdc ;};_agf :=_c .NewRGBA (_fadf );_bc .DrawMask (_agf ,_fadf ,_dgb ,_c .Point {},_caede ,_c .Point {},_bc .Over );return _agf ;};