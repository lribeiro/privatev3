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

package testutils ;import (_ge "crypto/md5";_ea "encoding/hex";_cf "errors";_a "fmt";_ff "github.com/unidoc/unipdf/v3/common";_dg "github.com/unidoc/unipdf/v3/core";_ce "image";_fe "image/png";_e "io";_gee "os";_f "os/exec";_c "path/filepath";_d "strings";
_b "testing";);func CompareImages (img1 ,img2 _ce .Image )(bool ,error ){_eb :=img1 .Bounds ();_bb :=0;for _bc :=0;_bc < _eb .Size ().X ;_bc ++{for _ec :=0;_ec < _eb .Size ().Y ;_ec ++{_gf ,_ed ,_bgc ,_ :=img1 .At (_bc ,_ec ).RGBA ();_beg ,_gfg ,_bgg ,_ :=img2 .At (_bc ,_ec ).RGBA ();
if _gf !=_beg ||_ed !=_gfg ||_bgc !=_bgg {_bb ++;};};};_gd :=float64 (_bb )/float64 (_eb .Dx ()*_eb .Dy ());if _gd > 0.0001{_a .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_gd ,_bb );
return false ,nil ;};return true ,nil ;};func CopyFile (src ,dst string )error {_ead ,_gc :=_gee .Open (src );if _gc !=nil {return _gc ;};defer _ead .Close ();_cc ,_gc :=_gee .Create (dst );if _gc !=nil {return _gc ;};defer _cc .Close ();_ ,_gc =_e .Copy (_cc ,_ead );
return _gc ;};func ReadPNG (file string )(_ce .Image ,error ){_fg ,_cfe :=_gee .Open (file );if _cfe !=nil {return nil ,_cfe ;};defer _fg .Close ();return _fe .Decode (_fg );};func HashFile (file string )(string ,error ){_bg ,_cee :=_gee .Open (file );
if _cee !=nil {return "",_cee ;};defer _bg .Close ();_gg :=_ge .New ();if _ ,_cee =_e .Copy (_gg ,_bg );_cee !=nil {return "",_cee ;};return _ea .EncodeToString (_gg .Sum (nil )),nil ;};func CompareDictionariesDeep (d1 ,d2 *_dg .PdfObjectDictionary )bool {if len (d1 .Keys ())!=len (d2 .Keys ()){_ff .Log .Debug ("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",len (d1 .Keys ()),len (d2 .Keys ()));
_ff .Log .Debug ("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027",d1 .WriteString (),d2 .WriteString ());return false ;};for _ ,_ggc :=range d1 .Keys (){if _ggc =="\u0050\u0061\u0072\u0065\u006e\u0074"{continue ;
};_bedb :=_dg .TraceToDirectObject (d1 .Get (_ggc ));_gec :=_dg .TraceToDirectObject (d2 .Get (_ggc ));if _bedb ==nil {_ff .Log .Debug ("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l");return false ;};if _gec ==nil {_ff .Log .Debug ("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l");
return false ;};switch _fgd :=_bedb .(type ){case *_dg .PdfObjectDictionary :_agea ,_cfa :=_gec .(*_dg .PdfObjectDictionary );if !_cfa {_ff .Log .Debug ("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054",_bedb ,_gec );
return false ;};if !CompareDictionariesDeep (_fgd ,_agea ){return false ;};continue ;case *_dg .PdfObjectArray :_gfgc ,_dgc :=_gec .(*_dg .PdfObjectArray );if !_dgc {_ff .Log .Debug ("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return false ;};if _fgd .Len ()!=_gfgc .Len (){_ff .Log .Debug ("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",_fgd .Len (),_gfgc .Len ());
return false ;};for _bfa :=0;_bfa < _fgd .Len ();_bfa ++{_aec :=_dg .TraceToDirectObject (_fgd .Get (_bfa ));_ffae :=_dg .TraceToDirectObject (_gfgc .Get (_bfa ));if _agead ,_gege :=_aec .(*_dg .PdfObjectDictionary );_gege {_egc ,_cac :=_ffae .(*_dg .PdfObjectDictionary );
if !_cac {return false ;};if !CompareDictionariesDeep (_agead ,_egc ){return false ;};}else {if _aec .WriteString ()!=_ffae .WriteString (){_ff .Log .Debug ("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027",_aec .WriteString (),_ffae .WriteString ());
return false ;};};};continue ;};if _bedb .String ()!=_gec .String (){_ff .Log .Debug ("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027",_ggc ,_bedb .String (),_gec .String ());
_ff .Log .Debug ("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027",_bedb ,_gec );_ff .Log .Debug ("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027",_bedb ,_gec );return false ;
};};return true ;};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_gda :=_f .LookPath ("\u0067\u0073");_gda !=nil {return ErrRenderNotSupported ;};return _f .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_a .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();
};func _fgb (_ca _dg .PdfObject ,_bed map[int64 ]_dg .PdfObject )error {switch _acc :=_ca .(type ){case *_dg .PdfIndirectObject :_cec :=_acc ;_fgb (_cec .PdfObject ,_bed );case *_dg .PdfObjectDictionary :_gbe :=_acc ;for _ ,_ba :=range _gbe .Keys (){_gad :=_gbe .Get (_ba );
if _fed ,_ffde :=_gad .(*_dg .PdfObjectReference );_ffde {_cdf ,_eeb :=_bed [_fed .ObjectNumber ];if !_eeb {return _cf .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_gbe .Set (_ba ,_cdf );}else {_fgb (_gad ,_bed );};};case *_dg .PdfObjectArray :_fedd :=_acc ;for _dcf ,_aaa :=range _fedd .Elements (){if _bgf ,_bf :=_aaa .(*_dg .PdfObjectReference );_bf {_fb ,_acf :=_bed [_bgf .ObjectNumber ];if !_acf {return _cf .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_fedd .Set (_dcf ,_fb );}else {_fgb (_aaa ,_bed );};};};return nil ;};var (ErrRenderNotSupported =_cf .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");
);func ParseIndirectObjects (rawpdf string )(map[int64 ]_dg .PdfObject ,error ){_age :=_dg .NewParserFromString (rawpdf );_adf :=map[int64 ]_dg .PdfObject {};for {_ae ,_edfb :=_age .ParseIndirectObject ();if _edfb !=nil {if _edfb ==_e .EOF {break ;};return nil ,_edfb ;
};switch _afb :=_ae .(type ){case *_dg .PdfIndirectObject :_adf [_afb .ObjectNumber ]=_ae ;case *_dg .PdfObjectStream :_adf [_afb .ObjectNumber ]=_ae ;};};for _ ,_fd :=range _adf {_fgb (_fd ,_adf );};return _adf ,nil ;};func RunRenderTest (t *_b .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ){_gca :=_d .TrimSuffix (_c .Base (pdfPath ),_c .Ext (pdfPath ));
t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_ee *_b .T ){_ffa :=_c .Join (outputDir ,_gca );_edf :=_ffa +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _ag :=RenderPDFToPNGs (pdfPath ,0,_edf );_ag !=nil {_ee .Skip (_ag );};for _ab :=1;true ;_ab ++{_eg :=_a .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_ffa ,_ab );
_bbf :=_c .Join (baselineRenderPath ,_a .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_gca ,_ab ));if _ ,_ad :=_gee .Stat (_eg );_ad !=nil {break ;};_ee .Logf ("\u0025\u0073",_bbf );if saveBaseline {_ee .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_eg ,_bbf );
_cd :=CopyFile (_eg ,_bbf );if _cd !=nil {_ee .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_bbf ,_cd );};continue ;};_ee .Run (_a .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_ab ),func (_aga *_b .T ){_aga .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_eg ,_bbf );
_ffd ,_ga :=ComparePNGFiles (_eg ,_bbf );if _gee .IsNotExist (_ga ){_aga .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_ffd {_aga .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");
};});};});};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_af ,_edc :=HashFile (file1 );if _edc !=nil {return false ,_edc ;};_cbd ,_edc :=HashFile (file2 );if _edc !=nil {return false ,_edc ;};if _af ==_cbd {return true ,nil ;};_dc ,_edc :=ReadPNG (file1 );
if _edc !=nil {return false ,_edc ;};_gb ,_edc :=ReadPNG (file2 );if _edc !=nil {return false ,_edc ;};if _dc .Bounds ()!=_gb .Bounds (){return false ,nil ;};return CompareImages (_dc ,_gb );};