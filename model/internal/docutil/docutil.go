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

package docutil ;import (_b "errors";_d "fmt";_g "github.com/unidoc/unipdf/v3/common";_f "github.com/unidoc/unipdf/v3/core";);func (_bada *Document )AddStream (stream *_f .PdfObjectStream ){for _ ,_cab :=range _bada .Objects {if _cab ==stream {return ;
};};_bada .Objects =append (_bada .Objects ,stream );};func (_cd *Catalog )SetMarkInfo (mi _f .PdfObject ){_gff :=_f .MakeIndirectObject (mi );_cd .Object .Set ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f",_gff );_cd ._bd .Objects =append (_cd ._bd .Objects ,_gff );
};func (_acb *Catalog )GetMarkInfo ()(*_f .PdfObjectDictionary ,bool ){_cc ,_ba :=_f .GetDict (_acb .Object .Get ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"));return _cc ,_ba ;};type Page struct{_ab int ;Object *_f .PdfObjectDictionary ;_aef *Document ;
};func (_ac *Catalog )SetMetadata (data []byte )error {_dc ,_gb :=_f .MakeStream (data ,nil );if _gb !=nil {return _gb ;};_dc .Set ("\u0054\u0079\u0070\u0065",_f .MakeName ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));_dc .Set ("\u0053u\u0062\u0074\u0079\u0070\u0065",_f .MakeName ("\u0058\u004d\u004c"));
_ac .Object .Set ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061",_dc );_ac ._bd .Objects =append (_ac ._bd .Objects ,_dc );return nil ;};type Catalog struct{Object *_f .PdfObjectDictionary ;_bd *Document ;};func (_bad *OutputIntents )Len ()int {return _bad ._bb .Len ()};
func (_gaaa *Content )SetData (data []byte )error {_cfbf ,_gdec :=_f .MakeStream (data ,_f .NewFlateEncoder ());if _gdec !=nil {return _gdec ;};_fga ,_ggf :=_f .GetArray (_gaaa ._ggb .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if !_ggf &&_gaaa ._gcg ==0{_gaaa ._ggb .Object .Set ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073",_cfbf );
}else {if _gdec =_fga .Set (_gaaa ._gcg ,_cfbf );_gdec !=nil {return _gdec ;};};_gaaa ._ggb ._aef .Objects =append (_gaaa ._ggb ._aef .Objects ,_cfbf );return nil ;};type OutputIntent struct{Object *_f .PdfObjectDictionary ;};func (_bc *OutputIntents )Get (i int )(OutputIntent ,bool ){if _bc ._bb ==nil {return OutputIntent {},false ;
};if i >=_bc ._bb .Len (){return OutputIntent {},false ;};_bcd :=_bc ._bb .Get (i );_bcdc ,_cac :=_f .GetIndirect (_bcd );if !_cac {_fa ,_bdb :=_f .GetDict (_bcd );return OutputIntent {Object :_fa },_bdb ;};_cdf ,_bgd :=_f .GetDict (_bcdc .PdfObject );
return OutputIntent {Object :_cdf },_bgd ;};func (_fed Page )FindXObjectImages ()([]*Image ,error ){_cbe ,_gdc :=_fed .GetResourcesXObject ();if !_gdc {return nil ,nil ;};var _fea []*Image ;var _gda error ;_age :=map[*_f .PdfObjectStream ]int {};_cff :=map[*_f .PdfObjectStream ]struct{}{};
var _ddcg int ;for _ ,_gfa :=range _cbe .Keys (){_dbg ,_aeg :=_f .GetStream (_cbe .Get (_gfa ));if !_aeg {continue ;};if _ ,_daf :=_age [_dbg ];_daf {continue ;};_aee ,_aagg :=_f .GetName (_dbg .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));if !_aagg ||_aee .String ()!="\u0049\u006d\u0061g\u0065"{continue ;
};_eee :=Image {BitsPerComponent :8,Stream :_dbg ,Name :string (_gfa )};if _eee .Colorspace ,_gda =_ggc (_dbg .PdfObjectDictionary .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));_gda !=nil {_g .Log .Error ("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073",_gda );
continue ;};if _dgf ,_cad :=_f .GetIntVal (_dbg .PdfObjectDictionary .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));_cad {_eee .BitsPerComponent =_dgf ;};if _bcb ,_bbg :=_f .GetIntVal (_dbg .PdfObjectDictionary .Get ("\u0057\u0069\u0064t\u0068"));
_bbg {_eee .Width =_bcb ;};if _dad ,_ceeg :=_f .GetIntVal (_dbg .PdfObjectDictionary .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));_ceeg {_eee .Height =_dad ;};if _egb ,_faf :=_f .GetStream (_dbg .Get ("\u0053\u004d\u0061s\u006b"));_faf {_eee .SMask =&ImageSMask {Image :&_eee ,Stream :_egb };
_cff [_egb ]=struct{}{};};switch _eee .Colorspace {case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":_eee .ColorComponents =3;case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":_eee .ColorComponents =1;case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":_eee .ColorComponents =4;
default:_eee .ColorComponents =-1;};_age [_dbg ]=_ddcg ;_fea =append (_fea ,&_eee );_ddcg ++;};var _bde []int ;for _ ,_badd :=range _fea {if _badd .SMask !=nil {_dbd ,_bbf :=_age [_badd .SMask .Stream ];if _bbf {_bde =append (_bde ,_dbd );};};};_gcf :=make ([]*Image ,len (_fea )-len (_bde ));
_ddcg =0;_egbe :for _gfc ,_fef :=range _fea {for _ ,_cea :=range _bde {if _gfc ==_cea {continue _egbe ;};};_gcf [_ddcg ]=_fef ;_ddcg ++;};return _fea ,nil ;};type Content struct{Stream *_f .PdfObjectStream ;_gcg int ;_ggb Page ;};type ImageSMask struct{Image *Image ;
Stream *_f .PdfObjectStream ;};func (_dd *Catalog )SetVersion (){_dd .Object .Set ("\u0056e\u0072\u0073\u0069\u006f\u006e",_f .MakeName (_d .Sprintf ("\u0025\u0064\u002e%\u0064",_dd ._bd .Version .Major ,_dd ._bd .Version .Minor )));};func (_ddg *Catalog )HasMetadata ()bool {_c :=_ddg .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061");
return _c !=nil ;};func (_ccc *Page )Number ()int {return _ccc ._ab };type OutputIntents struct{_bb *_f .PdfObjectArray ;_bbe *Document ;_bfd *_f .PdfIndirectObject ;};func (_gdg Page )FindXObjectForms ()[]*_f .PdfObjectStream {_dde ,_gbae :=_gdg .GetResourcesXObject ();
if !_gbae {return nil ;};_ece :=map[*_f .PdfObjectStream ]struct{}{};var _bbge func (_gdae *_f .PdfObjectDictionary ,_cdda map[*_f .PdfObjectStream ]struct{});_bbge =func (_cbdd *_f .PdfObjectDictionary ,_gde map[*_f .PdfObjectStream ]struct{}){for _ ,_eced :=range _cbdd .Keys (){_gdgg ,_fad :=_f .GetStream (_cbdd .Get (_eced ));
if !_fad {continue ;};if _ ,_fda :=_gde [_gdgg ];_fda {continue ;};_cceb ,_gcb :=_f .GetName (_gdgg .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));if !_gcb ||_cceb .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_gde [_gdgg ]=struct{}{};_dgca ,_gcb :=_f .GetDict (_gdgg .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));
if !_gcb {continue ;};_ggcb ,_dcf :=_f .GetDict (_dgca .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));if _dcf {_bbge (_ggcb ,_gde );};};};_bbge (_dde ,_ece );var _fc []*_f .PdfObjectStream ;for _ef :=range _ece {_fc =append (_fc ,_ef );};return _fc ;};
func (_aca Page )GetResourcesXObject ()(*_f .PdfObjectDictionary ,bool ){_cbf ,_eg :=_aca .GetResources ();if !_eg {return nil ,false ;};return _f .GetDict (_cbf .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));};func (_caa *OutputIntents )Add (oi _f .PdfObject )error {_cb ,_fd :=oi .(*_f .PdfObjectDictionary );
if !_fd {return _b .New ("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};if _db ,_ga :=_f .GetStream (_cb .Get ("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065"));_ga {_caa ._bbe .Objects =append (_caa ._bbe .Objects ,_db );};_bge ,_cfb :=oi .(*_f .PdfIndirectObject );if !_cfb {_bge =_f .MakeIndirectObject (oi );
};if _caa ._bb ==nil {_caa ._bb =_f .MakeArray (_bge );}else {_caa ._bb .Append (_bge );};_caa ._bbe .Objects =append (_caa ._bbe .Objects ,_bge );return nil ;};func (_cce *Document )FindCatalog ()(*Catalog ,bool ){var _fac *_f .PdfObjectDictionary ;for _ ,_aag :=range _cce .Objects {_gfde ,_ddc :=_f .GetDict (_aag );
if !_ddc {continue ;};if _df ,_dbe :=_f .GetName (_gfde .Get ("\u0054\u0079\u0070\u0065"));_dbe &&*_df =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_fac =_gfde ;break ;};};if _fac ==nil {return nil ,false ;};return &Catalog {Object :_fac ,_bd :_cce },true ;
};func (_bf *Catalog )SetOutputIntents (outputIntents *OutputIntents ){if _dg :=_bf .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");_dg !=nil {for _ag ,_cf :=range _bf ._bd .Objects {if _cf ==_dg {if outputIntents ._bfd ==_dg {return ;
};_bf ._bd .Objects =append (_bf ._bd .Objects [:_ag ],_bf ._bd .Objects [_ag +1:]...);break ;};};};_ee :=outputIntents ._bfd ;if _ee ==nil {_ee =_f .MakeIndirectObject (outputIntents ._bb );};_bf .Object .Set ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073",_ee );
_bf ._bd .Objects =append (_bf ._bd .Objects ,_ee );};func _ggc (_bbc _f .PdfObject )(_f .PdfObjectName ,error ){var _fb *_f .PdfObjectName ;var _fg *_f .PdfObjectArray ;if _cg ,_cdd :=_bbc .(*_f .PdfIndirectObject );_cdd {if _ecc ,_ddcc :=_cg .PdfObject .(*_f .PdfObjectArray );
_ddcc {_fg =_ecc ;}else if _bag ,_gffc :=_cg .PdfObject .(*_f .PdfObjectName );_gffc {_fb =_bag ;};}else if _fgc ,_ce :=_bbc .(*_f .PdfObjectArray );_ce {_fg =_fgc ;}else if _adg ,_fgb :=_bbc .(*_f .PdfObjectName );_fgb {_fb =_adg ;};if _fb !=nil {switch *_fb {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":return *_fb ,nil ;
case "\u0050a\u0074\u0074\u0065\u0072\u006e":return *_fb ,nil ;};};if _fg !=nil &&_fg .Len ()> 0{if _bfe ,_cdff :=_fg .Get (0).(*_f .PdfObjectName );_cdff {switch *_bfe {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":if _fg .Len ()==1{return *_bfe ,nil ;
};case "\u0043a\u006c\u0047\u0072\u0061\u0079","\u0043\u0061\u006c\u0052\u0047\u0042","\u004c\u0061\u0062":return *_bfe ,nil ;case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064","\u0050a\u0074\u0074\u0065\u0072\u006e","\u0049n\u0064\u0065\u0078\u0065\u0064":return *_bfe ,nil ;
case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e","\u0044e\u0076\u0069\u0063\u0065\u004e":return *_bfe ,nil ;};};};return "",nil ;};type Document struct{ID [2]string ;Version _f .Version ;Objects []_f .PdfObject ;Info _f .PdfObject ;Crypt *_f .PdfCrypt ;
UseHashBasedID bool ;};func (_ddca *Document )AddIndirectObject (indirect *_f .PdfIndirectObject ){for _ ,_dfc :=range _ddca .Objects {if _dfc ==indirect {return ;};};_ddca .Objects =append (_ddca .Objects ,indirect );};func (_ge *Catalog )GetPages ()([]Page ,bool ){_bg ,_e :=_f .GetDict (_ge .Object .Get ("\u0050\u0061\u0067e\u0073"));
if !_e {return nil ,false ;};_ec ,_fe :=_f .GetArray (_bg .Get ("\u004b\u0069\u0064\u0073"));if !_fe {return nil ,false ;};_fee :=make ([]Page ,_ec .Len ());for _ae ,_ad :=range _ec .Elements (){_gf ,_aa :=_f .GetDict (_ad );if !_aa {continue ;};_fee [_ae ]=Page {Object :_gf ,_ab :_ae +1,_aef :_ge ._bd };
};return _fee ,true ;};type Image struct{Name string ;Width int ;Height int ;Colorspace _f .PdfObjectName ;ColorComponents int ;BitsPerComponent int ;SMask *ImageSMask ;Stream *_f .PdfObjectStream ;};func (_bdd *Document )GetPages ()([]Page ,bool ){_ade ,_gd :=_bdd .FindCatalog ();
if !_gd {return nil ,false ;};return _ade .GetPages ();};func (_gcfb Content )GetData ()([]byte ,error ){_gfad ,_fcc :=_f .NewEncoderFromStream (_gcfb .Stream );if _fcc !=nil {return nil ,_fcc ;};_fcf ,_fcc :=_gfad .DecodeStream (_gcfb .Stream );if _fcc !=nil {return nil ,_fcc ;
};return _fcf ,nil ;};func (_ca *Catalog )GetMetadata ()(*_f .PdfObjectStream ,bool ){_da ,_gfd :=_f .GetStream (_ca .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));return _da ,_gfd ;};func (_cee Page )GetContents ()([]Content ,bool ){_dagd ,_gaa :=_f .GetArray (_cee .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
if !_gaa {_aed ,_be :=_f .GetStream (_cee .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if !_be {return nil ,false ;};return []Content {{Stream :_aed ,_ggb :_cee ,_gcg :0}},true ;};_ddb :=make ([]Content ,_dagd .Len ());for _egc ,_gc :=range _dagd .Elements (){_ced ,_dge :=_f .GetStream (_gc );
if !_dge {continue ;};_ddb [_egc ]=Content {Stream :_ced ,_ggb :_cee ,_gcg :_egc };};return _ddb ,true ;};func (_add *Catalog )GetOutputIntents ()(*OutputIntents ,bool ){_bgf :=_add .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");
if _bgf ==nil {return nil ,false ;};_gba ,_dgc :=_f .GetIndirect (_bgf );if !_dgc {return nil ,false ;};_addf ,_ged :=_f .GetArray (_gba .PdfObject );if !_ged {return nil ,false ;};return &OutputIntents {_bfd :_gba ,_bb :_addf ,_bbe :_add ._bd },true ;
};func (_cbd Page )GetResources ()(*_f .PdfObjectDictionary ,bool ){return _f .GetDict (_cbd .Object .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));};func (_dca *Catalog )NewOutputIntents ()*OutputIntents {return &OutputIntents {_bbe :_dca ._bd }};
