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

package testutils ;import (_c "crypto/md5";_dad "encoding/hex";_dag "errors";_ee "fmt";_eb "github.com/unidoc/unipdf/v3/common";_a "github.com/unidoc/unipdf/v3/core";_e "image";_b "image/png";_cb "io";_da "os";_dd "os/exec";_ga "path/filepath";_gb "strings";
_d "testing";);var (ErrRenderNotSupported =_dag .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");
);func HashFile (file string )(string ,error ){_eg ,_ge :=_da .Open (file );if _ge !=nil {return "",_ge ;};defer _eg .Close ();_ed :=_c .New ();if _ ,_ge =_cb .Copy (_ed ,_eg );_ge !=nil {return "",_ge ;};return _dad .EncodeToString (_ed .Sum (nil )),nil ;
};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_dda ,_geb :=HashFile (file1 );if _geb !=nil {return false ,_geb ;};_ege ,_geb :=HashFile (file2 );if _geb !=nil {return false ,_geb ;};if _dda ==_ege {return true ,nil ;};_fa ,_geb :=ReadPNG (file1 );
if _geb !=nil {return false ,_geb ;};_fc ,_geb :=ReadPNG (file2 );if _geb !=nil {return false ,_geb ;};if _fa .Bounds ()!=_fc .Bounds (){return false ,nil ;};return CompareImages (_fa ,_fc );};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;
};if _ ,_gg :=_dd .LookPath ("\u0067\u0073");_gg !=nil {return ErrRenderNotSupported ;};return _dd .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_ee .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();
};func ReadPNG (file string )(_e .Image ,error ){_ca ,_cdf :=_da .Open (file );if _cdf !=nil {return nil ,_cdf ;};defer _ca .Close ();return _b .Decode (_ca );};func CopyFile (src ,dst string )error {_gc ,_cd :=_da .Open (src );if _cd !=nil {return _cd ;
};defer _gc .Close ();_de ,_cd :=_da .Create (dst );if _cd !=nil {return _cd ;};defer _de .Close ();_ ,_cd =_cb .Copy (_de ,_gc );return _cd ;};func CompareImages (img1 ,img2 _e .Image )(bool ,error ){_edc :=img1 .Bounds ();_cab :=0;for _f :=0;_f < _edc .Size ().X ;
_f ++{for _ac :=0;_ac < _edc .Size ().Y ;_ac ++{_af ,_eba ,_bg ,_ :=img1 .At (_f ,_ac ).RGBA ();_bb ,_fb ,_dc ,_ :=img2 .At (_f ,_ac ).RGBA ();if _af !=_bb ||_eba !=_fb ||_bg !=_dc {_cab ++;};};};_cdd :=float64 (_cab )/float64 (_edc .Dx ()*_edc .Dy ());
if _cdd > 0.0001{_ee .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_cdd ,_cab );return false ,nil ;};return true ,nil ;};func CompareDictionariesDeep (d1 ,d2 *_a .PdfObjectDictionary )bool {if len (d1 .Keys ())!=len (d2 .Keys ()){_eb .Log .Debug ("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",len (d1 .Keys ()),len (d2 .Keys ()));
_eb .Log .Debug ("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027",d1 .WriteString (),d2 .WriteString ());return false ;};for _ ,_dea :=range d1 .Keys (){if _dea =="\u0050\u0061\u0072\u0065\u006e\u0074"{continue ;
};_ffa :=_a .TraceToDirectObject (d1 .Get (_dea ));_dbf :=_a .TraceToDirectObject (d2 .Get (_dea ));if _ffa ==nil {_eb .Log .Debug ("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l");return false ;};if _dbf ==nil {_eb .Log .Debug ("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l");
return false ;};switch _ggf :=_ffa .(type ){case *_a .PdfObjectDictionary :_bd ,_aaf :=_dbf .(*_a .PdfObjectDictionary );if !_aaf {_eb .Log .Debug ("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054",_ffa ,_dbf );
return false ;};if !CompareDictionariesDeep (_ggf ,_bd ){return false ;};continue ;case *_a .PdfObjectArray :_dcde ,_gcb :=_dbf .(*_a .PdfObjectArray );if !_gcb {_eb .Log .Debug ("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return false ;};if _ggf .Len ()!=_dcde .Len (){_eb .Log .Debug ("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",_ggf .Len (),_dcde .Len ());
return false ;};for _edb :=0;_edb < _ggf .Len ();_edb ++{_bbb :=_a .TraceToDirectObject (_ggf .Get (_edb ));_gab :=_a .TraceToDirectObject (_dcde .Get (_edb ));if _be ,_gdc :=_bbb .(*_a .PdfObjectDictionary );_gdc {_gce ,_cae :=_gab .(*_a .PdfObjectDictionary );
if !_cae {return false ;};if !CompareDictionariesDeep (_be ,_gce ){return false ;};}else {if _bbb .WriteString ()!=_gab .WriteString (){_eb .Log .Debug ("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027",_bbb .WriteString (),_gab .WriteString ());
return false ;};};};continue ;};if _ffa .String ()!=_dbf .String (){_eb .Log .Debug ("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027",_dea ,_ffa .String (),_dbf .String ());
_eb .Log .Debug ("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027",_ffa ,_dbf );_eb .Log .Debug ("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027",_ffa ,_dbf );return false ;
};};return true ;};func _ddc (_ae _a .PdfObject ,_ddce map[int64 ]_a .PdfObject )error {switch _age :=_ae .(type ){case *_a .PdfIndirectObject :_afec :=_age ;_ddc (_afec .PdfObject ,_ddce );case *_a .PdfObjectDictionary :_aga :=_age ;for _ ,_cbc :=range _aga .Keys (){_cfc :=_aga .Get (_cbc );
if _fbg ,_fcb :=_cfc .(*_a .PdfObjectReference );_fcb {_eed ,_bf :=_ddce [_fbg .ObjectNumber ];if !_bf {return _dag .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_aga .Set (_cbc ,_eed );}else {_ddc (_cfc ,_ddce );};};case *_a .PdfObjectArray :_dcd :=_age ;for _dba ,_dcc :=range _dcd .Elements (){if _fd ,_ebf :=_dcc .(*_a .PdfObjectReference );_ebf {_bgg ,_daa :=_ddce [_fd .ObjectNumber ];if !_daa {return _dag .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_dcd .Set (_dba ,_bgg );}else {_ddc (_dcc ,_ddce );};};};return nil ;};func RunRenderTest (t *_d .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ){_gf :=_gb .TrimSuffix (_ga .Base (pdfPath ),_ga .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_eda *_d .T ){_cdc :=_ga .Join (outputDir ,_gf );
_ged :=_cdc +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _ag :=RenderPDFToPNGs (pdfPath ,0,_ged );_ag !=nil {_eda .Skip (_ag );};for _gec :=1;true ;_gec ++{_edcg :=_ee .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_cdc ,_gec );_fag :=_ga .Join (baselineRenderPath ,_ee .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_gf ,_gec ));
if _ ,_bba :=_da .Stat (_edcg );_bba !=nil {break ;};_eda .Logf ("\u0025\u0073",_fag );if saveBaseline {_eda .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_edcg ,_fag );_afe :=CopyFile (_edcg ,_fag );if _afe !=nil {_eda .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_fag ,_afe );
};continue ;};_eda .Run (_ee .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_gec ),func (_cdda *_d .T ){_cdda .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_edcg ,_fag );_eec ,_ea :=ComparePNGFiles (_edcg ,_fag );
if _da .IsNotExist (_ea ){_cdda .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_eec {_cdda .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");
};});};});};func ParseIndirectObjects (rawpdf string )(map[int64 ]_a .PdfObject ,error ){_dbe :=_a .NewParserFromString (rawpdf );_gba :=map[int64 ]_a .PdfObject {};for {_ff ,_bc :=_dbe .ParseIndirectObject ();if _bc !=nil {if _bc ==_cb .EOF {break ;};
return nil ,_bc ;};switch _fbe :=_ff .(type ){case *_a .PdfIndirectObject :_gba [_fbe .ObjectNumber ]=_ff ;case *_a .PdfObjectStream :_gba [_fbe .ObjectNumber ]=_ff ;};};for _ ,_aa :=range _gba {_ddc (_aa ,_gba );};return _gba ,nil ;};