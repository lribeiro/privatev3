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

package optimize ;import (_ef "bytes";_da "crypto/md5";_df "errors";_b "github.com/unidoc/unipdf/v3/common";_c "github.com/unidoc/unipdf/v3/contentstream";_bf "github.com/unidoc/unipdf/v3/core";_de "github.com/unidoc/unipdf/v3/extractor";_e "github.com/unidoc/unipdf/v3/internal/imageutil";
_a "github.com/unidoc/unipdf/v3/internal/textencoding";_eb "github.com/unidoc/unipdf/v3/model";_fc "github.com/unidoc/unitype";_f "golang.org/x/image/draw";_fe "math";);func _gfe (_edb []_bf .PdfObject )(_aa map[*_bf .PdfObjectStream ]struct{},_abf error ){_aa =map[*_bf .PdfObjectStream ]struct{}{};
_be :=map[*_eb .PdfFont ]struct{}{};_ce :=_dbba (_edb );for _ ,_bdda :=range _ce ._daff {_abfa ,_ddb :=_bf .GetDict (_bdda .PdfObject );if !_ddb {continue ;};_ada ,_ddb :=_bf .GetDict (_abfa .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));if !_ddb {continue ;
};_gd ,_ :=_fcfd (_abfa .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_fed ,_ee :=_eb .NewPdfPageResourcesFromDict (_ada );if _ee !=nil {return nil ,_ee ;};_bdcd :=[]content {{_aec :_gd ,_bea :_fed }};_bee :=_aeg (_abfa .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));
if _bee !=nil {_bdcd =append (_bdcd ,_bee ...);};for _ ,_aee :=range _bdcd {_eab ,_edbg :=_de .NewFromContents (_aee ._aec ,_aee ._bea );if _edbg !=nil {return nil ,_edbg ;};_eba ,_ ,_ ,_edbg :=_eab .ExtractPageText ();if _edbg !=nil {return nil ,_edbg ;
};for _ ,_aba :=range _eba .Marks ().Elements (){if _aba .Font ==nil {continue ;};if _ ,_cef :=_be [_aba .Font ];!_cef {_be [_aba .Font ]=struct{}{};};};};};_bdf :=map[*_bf .PdfObjectStream ][]*_eb .PdfFont {};for _bbfa :=range _be {_cafb :=_bbfa .FontDescriptor ();
if _cafb ==nil ||_cafb .FontFile2 ==nil {continue ;};_dce ,_fced :=_bf .GetStream (_cafb .FontFile2 );if !_fced {continue ;};_bdf [_dce ]=append (_bdf [_dce ],_bbfa );};for _cg :=range _bdf {var _egd []rune ;var _abba []_fc .GlyphIndex ;for _ ,_eec :=range _bdf [_cg ]{switch _fa :=_eec .Encoder ().(type ){case *_a .IdentityEncoder :_gdg :=_fa .RegisteredRunes ();
_eag :=make ([]_fc .GlyphIndex ,len (_gdg ));for _cfe ,_cbe :=range _gdg {_eag [_cfe ]=_fc .GlyphIndex (_cbe );};_abba =append (_abba ,_eag ...);case *_a .TrueTypeFontEncoder :_bff :=_fa .RegisteredRunes ();_egd =append (_egd ,_bff ...);case _a .SimpleEncoder :_agb :=_fa .Charcodes ();
for _ ,_fcac :=range _agb {_afd ,_fbgd :=_fa .CharcodeToRune (_fcac );if !_fbgd {_b .Log .Debug ("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064",_fcac );
continue ;};_egd =append (_egd ,_afd );};};};_abf =_egb (_cg ,_egd ,_abba );if _abf !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076",_abf );
return nil ,_abf ;};_aa [_cg ]=struct{}{};};return _aa ,nil ;};func _faa (_fddc *_eb .Image ,_edbge float64 )(*_eb .Image ,error ){_gddc ,_ccf :=_fddc .ToGoImage ();if _ccf !=nil {return nil ,_ccf ;};var _ege _e .Image ;_acc ,_bfd :=_gddc .(*_e .Monochrome );
if _bfd {if _ccf =_acc .ResolveDecode ();_ccf !=nil {return nil ,_ccf ;};_ege ,_ccf =_acc .Scale (_edbge );if _ccf !=nil {return nil ,_ccf ;};}else {_gda :=int (_fe .RoundToEven (float64 (_fddc .Width )*_edbge ));_gccd :=int (_fe .RoundToEven (float64 (_fddc .Height )*_edbge ));
_ege ,_ccf =_e .NewImage (_gda ,_gccd ,int (_fddc .BitsPerComponent ),_fddc .ColorComponents ,nil ,nil ,nil );if _ccf !=nil {return nil ,_ccf ;};_f .CatmullRom .Scale (_ege ,_ege .Bounds (),_gddc ,_gddc .Bounds (),_f .Over ,&_f .Options {});};_ceaf :=_ege .Base ();
_eae :=&_eb .Image {Width :int64 (_ceaf .Width ),Height :int64 (_ceaf .Height ),BitsPerComponent :int64 (_ceaf .BitsPerComponent ),ColorComponents :_ceaf .ColorComponents ,Data :_ceaf .Data };_eae .SetDecode (_ceaf .Decode );_eae .SetAlpha (_ceaf .Alpha );
return _eae ,nil ;};type objectStructure struct{_bbca *_bf .PdfObjectDictionary ;_dccc *_bf .PdfObjectDictionary ;_daff []*_bf .PdfIndirectObject ;};func _bgag (_add *_eb .XObjectImage ,_aeac imageModifications )error {_gecd ,_fddd :=_add .ToImage ();if _fddd !=nil {return _fddd ;
};if _aeac .Scale !=0{_gecd ,_fddd =_faa (_gecd ,_aeac .Scale );if _fddd !=nil {return _fddd ;};};if _aeac .Encoding !=nil {_add .Filter =_aeac .Encoding ;};_add .Decode =nil ;switch _gbef :=_add .Filter .(type ){case *_bf .FlateEncoder :if _gbef .Predictor !=1&&_gbef .Predictor !=11{_gbef .Predictor =1;
};};if _fddd =_add .SetImage (_gecd ,nil );_fddd !=nil {_b .Log .Debug ("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076",_fddd );return _fddd ;};_add .ToPdfObject ();return nil ;
};

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{};type imageInfo struct{BitsPerComponent int ;ColorComponents int ;Width int ;Height int ;Stream *_bf .PdfObjectStream ;PPI float64 ;};func _ff (_bfb *_c .ContentStreamOperations )*_c .ContentStreamOperations {if _bfb ==nil {return nil ;
};_ba :=_c .ContentStreamOperations {};for _ ,_ca :=range *_bfb {switch _ca .Operand {case "\u0042\u0044\u0043","\u0042\u004d\u0043","\u0045\u004d\u0043":continue ;case "\u0054\u006d":if len (_ca .Params )==6{if _bg ,_ab :=_bf .GetNumbersAsFloat (_ca .Params );
_ab ==nil {if _bg [0]==1&&_bg [1]==0&&_bg [2]==0&&_bg [3]==1{_ca =&_c .ContentStreamOperation {Params :[]_bf .PdfObject {_ca .Params [4],_ca .Params [5]},Operand :"\u0054\u0064"};};};};};_ba =append (_ba ,_ca );};return &_ba ;};

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{};

// Optimize optimizes PDF objects to decrease PDF size.
func (_gecf *Image )Optimize (objects []_bf .PdfObject )(_ccbf []_bf .PdfObject ,_afb error ){if _gecf .ImageQuality <=0{return objects ,nil ;};_aecf :=_cabe (objects );if len (_aecf )==0{return objects ,nil ;};_dfa :=make (map[_bf .PdfObject ]_bf .PdfObject );
_cdf :=make (map[_bf .PdfObject ]struct{});for _ ,_abad :=range _aecf {_bda :=_abad .Stream .Get ("\u0053\u004d\u0061s\u006b");_cdf [_bda ]=struct{}{};};for _bcbe ,_dga :=range _aecf {_fgga :=_dga .Stream ;if _ ,_bdaf :=_cdf [_fgga ];_bdaf {continue ;};
_fagc ,_gbbb :=_eb .NewXObjectImageFromStream (_fgga );if _gbbb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_gbbb );continue ;};switch _fagc .Filter .(type ){case *_bf .JBIG2Encoder :continue ;case *_bf .CCITTFaxEncoder :continue ;
};_fad ,_gbbb :=_fagc .ToImage ();if _gbbb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_gbbb );continue ;};_baac :=_bf .NewDCTEncoder ();_baac .ColorComponents =_fad .ColorComponents ;_baac .Quality =_gecf .ImageQuality ;
_baac .BitsPerComponent =_dga .BitsPerComponent ;_baac .Width =_dga .Width ;_baac .Height =_dga .Height ;_gdbe ,_gbbb :=_baac .EncodeBytes (_fad .Data );if _gbbb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_gbbb );
continue ;};var _fdef _bf .StreamEncoder ;_fdef =_baac ;{_ddef :=_bf .NewFlateEncoder ();_degge :=_bf .NewMultiEncoder ();_degge .AddEncoder (_ddef );_degge .AddEncoder (_baac );_aege ,_bdgb :=_degge .EncodeBytes (_fad .Data );if _bdgb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_bdgb );
continue ;};if len (_aege )< len (_gdbe ){_b .Log .Trace ("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029",len (_gdbe ),len (_aege ),len (_fgga .Stream ));
_gdbe =_aege ;_fdef =_degge ;};};_ceeca :=len (_fgga .Stream );if _ceeca < len (_gdbe ){continue ;};_aea :=&_bf .PdfObjectStream {Stream :_gdbe };_aea .PdfObjectReference =_fgga .PdfObjectReference ;_aea .PdfObjectDictionary =_bf .MakeDict ();_aea .Merge (_fgga .PdfObjectDictionary );
_aea .Merge (_fdef .MakeStreamDict ());_aea .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_bf .MakeInteger (int64 (len (_gdbe ))));_dfa [_fgga ]=_aea ;_aecf [_bcbe ].Stream =_aea ;};_ccbf =make ([]_bf .PdfObject ,len (objects ));copy (_ccbf ,objects );_gddg (_ccbf ,_dfa );
return _ccbf ,nil ;};

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct{

// Subset embedded fonts if encountered (if true).
// Otherwise attempts to reduce the font program.
Subset bool ;};

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{_cc []_eb .Optimizer };

// Optimize optimizes PDF objects to decrease PDF size.
func (_gcd *CombineDuplicateDirectObjects )Optimize (objects []_bf .PdfObject )(_acg []_bf .PdfObject ,_agg error ){_deaa (objects );_gdb :=make (map[string ][]*_bf .PdfObjectDictionary );var _cccb func (_eee *_bf .PdfObjectDictionary );_cccb =func (_geg *_bf .PdfObjectDictionary ){for _ ,_fgg :=range _geg .Keys (){_gcb :=_geg .Get (_fgg );
if _gbbd ,_cea :=_gcb .(*_bf .PdfObjectDictionary );_cea {_dfg :=_da .New ();_dfg .Write ([]byte (_gbbd .WriteString ()));_bag :=string (_dfg .Sum (nil ));_gdb [_bag ]=append (_gdb [_bag ],_gbbd );_cccb (_gbbd );};};};for _ ,_gde :=range objects {_fdg ,_agc :=_gde .(*_bf .PdfIndirectObject );
if !_agc {continue ;};if _bga ,_fdc :=_fdg .PdfObject .(*_bf .PdfObjectDictionary );_fdc {_cccb (_bga );};};_eda :=make ([]_bf .PdfObject ,0,len (_gdb ));_beaf :=make (map[_bf .PdfObject ]_bf .PdfObject );for _ ,_bffc :=range _gdb {if len (_bffc )< 2{continue ;
};_eeed :=_bf .MakeDict ();_eeed .Merge (_bffc [0]);_gdc :=_bf .MakeIndirectObject (_eeed );_eda =append (_eda ,_gdc );for _ceec :=0;_ceec < len (_bffc );_ceec ++{_dff :=_bffc [_ceec ];_beaf [_dff ]=_gdc ;};};_acg =make ([]_bf .PdfObject ,len (objects ));
copy (_acg ,objects );_acg =append (_eda ,_acg ...);_gddg (_acg ,_beaf );return _acg ,nil ;};func _fcfd (_adca _bf .PdfObject )(_dgdb string ,_eeg []_bf .PdfObject ){var _cca _ef .Buffer ;switch _eccda :=_adca .(type ){case *_bf .PdfIndirectObject :_eeg =append (_eeg ,_eccda );
_adca =_eccda .PdfObject ;};switch _dfae :=_adca .(type ){case *_bf .PdfObjectStream :if _bge ,_gcdd :=_bf .DecodeStream (_dfae );_gcdd ==nil {_cca .Write (_bge );_eeg =append (_eeg ,_dfae );};case *_bf .PdfObjectArray :for _ ,_fdcd :=range _dfae .Elements (){switch _fceda :=_fdcd .(type ){case *_bf .PdfObjectStream :if _fagd ,_ffa :=_bf .DecodeStream (_fceda );
_ffa ==nil {_cca .Write (_fagd );_eeg =append (_eeg ,_fceda );};};};};return _cca .String (),_eeg ;};func _aeg (_cdc _bf .PdfObject )[]content {if _cdc ==nil {return nil ;};_bbc ,_bgf :=_bf .GetArray (_cdc );if !_bgf {_b .Log .Debug ("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return nil ;};var _cafe []content ;for _ ,_agbd :=range _bbc .Elements (){_dfe ,_caa :=_bf .GetDict (_agbd );if !_caa {_b .Log .Debug ("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073");
continue ;};_ec ,_caa :=_bf .GetDict (_dfe .Get ("\u0041\u0050"));if !_caa {_b .Log .Debug ("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");continue ;};_fbd :=_bf .TraceToDirectObject (_ec .Get ("\u004e"));
if _fbd ==nil {_b .Log .Debug ("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067");continue ;};var _gcf *_bf .PdfObjectStream ;switch _cge :=_fbd .(type ){case *_bf .PdfObjectDictionary :_cda ,_bgd :=_bf .GetName (_dfe .Get ("\u0041\u0053"));
if !_bgd {_b .Log .Debug ("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");continue ;};_gcf ,_bgd =_bf .GetStream (_cge .Get (*_cda ));if !_bgd {_b .Log .Debug ("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg");
continue ;};case *_bf .PdfObjectStream :_gcf =_cge ;};if _gcf ==nil {_b .Log .Debug ("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067");
continue ;};_dgg ,_dab :=_eb .NewXObjectFormFromStream (_gcf );if _dab !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_dab );
continue ;};_efgb ,_dab :=_dgg .GetContentStream ();if _dab !=nil {_b .Log .Debug ("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076",_dab );continue ;};_cafe =append (_cafe ,content {_aec :string (_efgb ),_bea :_dgg .Resources });
};return _cafe ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_ebdc *ObjectStreams )Optimize (objects []_bf .PdfObject )(_abfg []_bf .PdfObject ,_egba error ){_cebf :=&_bf .PdfObjectStreams {};_edcg :=make ([]_bf .PdfObject ,0,len (objects ));for _ ,_dda :=range objects {if _gad ,_cgbe :=_dda .(*_bf .PdfIndirectObject );
_cgbe &&_gad .GenerationNumber ==0{_cebf .Append (_dda );}else {_edcg =append (_edcg ,_dda );};};if _cebf .Len ()==0{return _edcg ,nil ;};_abfg =make ([]_bf .PdfObject ,0,len (_edcg )+_cebf .Len ()+1);if _cebf .Len ()> 1{_abfg =append (_abfg ,_cebf );};
_abfg =append (_abfg ,_cebf .Elements ()...);_abfg =append (_abfg ,_edcg ...);return _abfg ,nil ;};

// New creates a optimizers chain from options.
func New (options Options )*Chain {_gege :=new (Chain );if options .CleanFonts ||options .SubsetFonts {_gege .Append (&CleanFonts {Subset :options .SubsetFonts });};if options .CleanContentstream {_gege .Append (new (CleanContentstream ));};if options .ImageUpperPPI > 0{_fbafg :=new (ImagePPI );
_fbafg .ImageUpperPPI =options .ImageUpperPPI ;_gege .Append (_fbafg );};if options .ImageQuality > 0{_caag :=new (Image );_caag .ImageQuality =options .ImageQuality ;_gege .Append (_caag );};if options .CombineDuplicateDirectObjects {_gege .Append (new (CombineDuplicateDirectObjects ));
};if options .CombineDuplicateStreams {_gege .Append (new (CombineDuplicateStreams ));};if options .CombineIdenticalIndirectObjects {_gege .Append (new (CombineIdenticalIndirectObjects ));};if options .UseObjectStreams {_gege .Append (new (ObjectStreams ));
};if options .CompressStreams {_gege .Append (new (CompressStreams ));};return _gege ;};

// Append appends optimizers to the chain.
func (_bc *Chain )Append (optimizers ..._eb .Optimizer ){_bc ._cc =append (_bc ._cc ,optimizers ...)};

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ImageUpperPPI float64 ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_baga *ImagePPI )Optimize (objects []_bf .PdfObject )(_decd []_bf .PdfObject ,_addf error ){if _baga .ImageUpperPPI <=0{return objects ,nil ;};_aabg :=_cabe (objects );if len (_aabg )==0{return objects ,nil ;};_fgb :=make (map[_bf .PdfObject ]struct{});
for _ ,_afde :=range _aabg {_gbeb :=_afde .Stream .PdfObjectDictionary .Get ("\u0053\u004d\u0061s\u006b");_fgb [_gbeb ]=struct{}{};};_ffe :=make (map[*_bf .PdfObjectStream ]*imageInfo );for _ ,_gbbbe :=range _aabg {_ffe [_gbbbe .Stream ]=_gbbbe ;};var _baaca *_bf .PdfObjectDictionary ;
for _ ,_fgbf :=range objects {if _gbag ,_gfdd :=_bf .GetDict (_fgbf );_baaca ==nil &&_gfdd {if _dbf ,_debbf :=_bf .GetName (_gbag .Get ("\u0054\u0079\u0070\u0065"));_debbf &&*_dbf =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_baaca =_gbag ;};};};if _baaca ==nil {return objects ,nil ;
};_cdg ,_ecd :=_bf .GetDict (_baaca .Get ("\u0050\u0061\u0067e\u0073"));if !_ecd {return objects ,nil ;};_eea ,_ceeb :=_bf .GetArray (_cdg .Get ("\u004b\u0069\u0064\u0073"));if !_ceeb {return objects ,nil ;};for _ ,_gaa :=range _eea .Elements (){_ffcb :=make (map[string ]*imageInfo );
_dbb ,_aae :=_bf .GetDict (_gaa );if !_aae {continue ;};_fbed ,_ :=_fcfd (_dbb .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if len (_fbed )==0{continue ;};_cbef ,_eefb :=_bf .GetDict (_dbb .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));
if !_eefb {continue ;};_gebf ,_bfbe :=_eb .NewPdfPageResourcesFromDict (_cbef );if _bfbe !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076",_bfbe );
continue ;};_afg ,_abaf :=_bf .GetDict (_cbef .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));if !_abaf {continue ;};_cgaa :=_afg .Keys ();for _ ,_gbce :=range _cgaa {if _gca ,_dccg :=_bf .GetStream (_afg .Get (_gbce ));_dccg {if _afgd ,_acce :=_ffe [_gca ];
_acce {_ffcb [string (_gbce )]=_afgd ;};};};_eeca :=_c .NewContentStreamParser (_fbed );_fec ,_bfbe :=_eeca .Parse ();if _bfbe !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_bfbe );continue ;};_fddda :=_c .NewContentStreamProcessor (*_fec );
_fddda .AddHandler (_c .HandlerConditionEnumAllOperands ,"",func (_dba *_c .ContentStreamOperation ,_cgbg _c .GraphicsState ,_afgdg *_eb .PdfPageResources )error {switch _dba .Operand {case "\u0044\u006f":if len (_dba .Params )!=1{_b .Log .Debug ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031");
return nil ;};_cbca ,_cce :=_bf .GetName (_dba .Params [0]);if !_cce {_b .Log .Debug ("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072");
return nil ;};if _bcc ,_dee :=_ffcb [string (*_cbca )];_dee {_dfb :=_cgbg .CTM .ScalingFactorX ();_ccfa :=_cgbg .CTM .ScalingFactorY ();_fadd ,_bdfb :=_dfb /72.0,_ccfa /72.0;_cdfg ,_geag :=float64 (_bcc .Width )/_fadd ,float64 (_bcc .Height )/_bdfb ;if _fadd ==0||_bdfb ==0{_cdfg =72.0;
_geag =72.0;};_bcc .PPI =_fe .Max (_bcc .PPI ,_cdfg );_bcc .PPI =_fe .Max (_bcc .PPI ,_geag );};};return nil ;});_bfbe =_fddda .Process (_gebf );if _bfbe !=nil {_b .Log .Debug ("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076",_bfbe );
continue ;};};for _ ,_afgf :=range _aabg {if _ ,_dcd :=_fgb [_afgf .Stream ];_dcd {continue ;};if _afgf .PPI <=_baga .ImageUpperPPI {continue ;};_adcg ,_ddbc :=_eb .NewXObjectImageFromStream (_afgf .Stream );if _ddbc !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_ddbc );
continue ;};var _cdcb imageModifications ;_cdcb .Scale =_baga .ImageUpperPPI /_afgf .PPI ;if _afgf .BitsPerComponent ==1&&_afgf .ColorComponents ==1{_bgg :=_fe .Round (_afgf .PPI /_baga .ImageUpperPPI );_dagf :=_e .NextPowerOf2 (uint (_bgg ));if _e .InDelta (float64 (_dagf ),1/_cdcb .Scale ,0.3){_cdcb .Scale =float64 (1)/float64 (_dagf );
};if _ ,_dgdf :=_adcg .Filter .(*_bf .JBIG2Encoder );!_dgdf {_cdcb .Encoding =_bf .NewJBIG2Encoder ();};};if _ddbc =_bgag (_adcg ,_cdcb );_ddbc !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_ddbc );
continue ;};_cdcb .Encoding =nil ;if _fafb ,_aegg :=_bf .GetStream (_afgf .Stream .PdfObjectDictionary .Get ("\u0053\u004d\u0061s\u006b"));_aegg {_bcec ,_dcb :=_eb .NewXObjectImageFromStream (_fafb );if _dcb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_dcb );
continue ;};if _dcb =_bgag (_bcec ,_cdcb );_dcb !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_dcb );continue ;};};};return objects ,nil ;};type imageModifications struct{Scale float64 ;Encoding _bf .StreamEncoder ;
};func _dbba (_degf []_bf .PdfObject )objectStructure {_ffgc :=objectStructure {};_cgdf :=false ;for _ ,_gef :=range _degf {switch _gdce :=_gef .(type ){case *_bf .PdfIndirectObject :_egg ,_beb :=_bf .GetDict (_gdce );if !_beb {continue ;};_bage ,_beb :=_bf .GetName (_egg .Get ("\u0054\u0079\u0070\u0065"));
if !_beb {continue ;};switch _bage .String (){case "\u0043a\u0074\u0061\u006c\u006f\u0067":_ffgc ._bbca =_egg ;_cgdf =true ;};};if _cgdf {break ;};};if !_cgdf {return _ffgc ;};_fda ,_ecdgb :=_bf .GetDict (_ffgc ._bbca .Get ("\u0050\u0061\u0067e\u0073"));
if !_ecdgb {return _ffgc ;};_ffgc ._dccc =_fda ;_eeaf ,_ecdgb :=_bf .GetArray (_fda .Get ("\u004b\u0069\u0064\u0073"));if !_ecdgb {return _ffgc ;};for _ ,_edcd :=range _eeaf .Elements (){_cfed ,_ggcf :=_bf .GetIndirect (_edcd );if !_ggcf {break ;};_ffgc ._daff =append (_ffgc ._daff ,_cfed );
};return _ffgc ;};func _cabe (_gbc []_bf .PdfObject )[]*imageInfo {_fbe :=_bf .PdfObjectName ("\u0053u\u0062\u0074\u0079\u0070\u0065");_deg :=make (map[*_bf .PdfObjectStream ]struct{});var _fefc []*imageInfo ;for _ ,_fag :=range _gbc {_afee ,_ggg :=_bf .GetStream (_fag );
if !_ggg {continue ;};if _ ,_gbf :=_deg [_afee ];_gbf {continue ;};_deg [_afee ]=struct{}{};_gbee :=_afee .PdfObjectDictionary .Get (_fbe );_db ,_ggg :=_bf .GetName (_gbee );if !_ggg ||string (*_db )!="\u0049\u006d\u0061g\u0065"{continue ;};_bce :=&imageInfo {Stream :_afee ,BitsPerComponent :8};
if _edg ,_dde :=_bf .GetIntVal (_afee .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));_dde {_bce .BitsPerComponent =_edg ;};if _degg ,_fbab :=_bf .GetIntVal (_afee .Get ("\u0057\u0069\u0064t\u0068"));_fbab {_bce .Width =_degg ;
};if _dacb ,_bgb :=_bf .GetIntVal (_afee .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));_bgb {_bce .Height =_dacb ;};_adc ,_bdg :=_eb .NewPdfColorspaceFromPdfObject (_afee .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));if _bdg !=nil {_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_bdg );
continue ;};if _adc ==nil {_baaa ,_gcc :=_bf .GetName (_afee .Get ("\u0046\u0069\u006c\u0074\u0065\u0072"));if _gcc {switch _baaa .String (){case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065","J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":_adc =_eb .NewPdfColorspaceDeviceGray ();
_bce .BitsPerComponent =1;};};};switch _edfe :=_adc .(type ){case *_eb .PdfColorspaceDeviceRGB :_bce .ColorComponents =3;case *_eb .PdfColorspaceDeviceGray :_bce .ColorComponents =1;default:_b .Log .Debug ("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070",_edfe );
continue ;};_fefc =append (_fefc ,_bce );};return _fefc ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_baad *CombineDuplicateStreams )Optimize (objects []_bf .PdfObject )(_aegf []_bf .PdfObject ,_afe error ){_dfd :=make (map[_bf .PdfObject ]_bf .PdfObject );_cgcc :=make (map[_bf .PdfObject ]struct{});_gab :=make (map[string ][]*_bf .PdfObjectStream );
for _ ,_gba :=range objects {if _beea ,_fdd :=_gba .(*_bf .PdfObjectStream );_fdd {_acbe :=_da .New ();_acbe .Write (_beea .Stream );_acbe .Write ([]byte (_beea .PdfObjectDictionary .WriteString ()));_ffb :=string (_acbe .Sum (nil ));_gab [_ffb ]=append (_gab [_ffb ],_beea );
};};for _ ,_geb :=range _gab {if len (_geb )< 2{continue ;};_efaa :=_geb [0];for _edf :=1;_edf < len (_geb );_edf ++{_debb :=_geb [_edf ];_dfd [_debb ]=_efaa ;_cgcc [_debb ]=struct{}{};};};_aegf =make ([]_bf .PdfObject ,0,len (objects )-len (_cgcc ));for _ ,_dfeb :=range objects {if _ ,_cgb :=_cgcc [_dfeb ];
_cgb {continue ;};_aegf =append (_aegf ,_dfeb );};_gddg (_aegf ,_dfd );return _aegf ,nil ;};func _gddg (_agf []_bf .PdfObject ,_fedd map[_bf .PdfObject ]_bf .PdfObject ){if len (_fedd )==0{return ;};for _bcbc ,_ffcc :=range _agf {if _bdcg ,_ddc :=_fedd [_ffcc ];
_ddc {_agf [_bcbc ]=_bdcg ;continue ;};_fedd [_ffcc ]=_ffcc ;switch _cgd :=_ffcc .(type ){case *_bf .PdfObjectArray :_eefa :=make ([]_bf .PdfObject ,_cgd .Len ());copy (_eefa ,_cgd .Elements ());_gddg (_eefa ,_fedd );for _fcec ,_ggc :=range _eefa {_cgd .Set (_fcec ,_ggc );
};case *_bf .PdfObjectStreams :_gddg (_cgd .Elements (),_fedd );case *_bf .PdfObjectStream :_cfg :=[]_bf .PdfObject {_cgd .PdfObjectDictionary };_gddg (_cfg ,_fedd );_cgd .PdfObjectDictionary =_cfg [0].(*_bf .PdfObjectDictionary );case *_bf .PdfObjectDictionary :_ddd :=_cgd .Keys ();
_dbd :=make ([]_bf .PdfObject ,len (_ddd ));for _fbde ,_eecf :=range _ddd {_dbd [_fbde ]=_cgd .Get (_eecf );};_gddg (_dbd ,_fedd );for _gggc ,_aade :=range _ddd {_cgd .Set (_aade ,_dbd [_gggc ]);};case *_bf .PdfIndirectObject :_cbcf :=[]_bf .PdfObject {_cgd .PdfObject };
_gddg (_cbcf ,_fedd );_cgd .PdfObject =_cbcf [0];};};};

// Options describes PDF optimization parameters.
type Options struct{CombineDuplicateStreams bool ;CombineDuplicateDirectObjects bool ;ImageUpperPPI float64 ;ImageQuality int ;UseObjectStreams bool ;CombineIdenticalIndirectObjects bool ;CompressStreams bool ;CleanFonts bool ;SubsetFonts bool ;CleanContentstream bool ;
};

// Optimize optimizes PDF objects to decrease PDF size.
func (_agbe *CleanFonts )Optimize (objects []_bf .PdfObject )(_fde []_bf .PdfObject ,_gfc error ){var _aad map[*_bf .PdfObjectStream ]struct{};if _agbe .Subset {var _acf error ;_aad ,_acf =_gfe (objects );if _acf !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076",_acf );
return nil ,_acf ;};};for _ ,_eaa :=range objects {_febf ,_gfd :=_bf .GetStream (_eaa );if !_gfd {continue ;};if _ ,_cgc :=_aad [_febf ];_cgc {continue ;};_cbc ,_acba :=_bf .NewEncoderFromStream (_febf );if _acba !=nil {_b .Log .Debug ("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067",_acba );
continue ;};_ded ,_acba :=_cbc .DecodeStream (_febf );if _acba !=nil {_b .Log .Debug ("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067",_acba );
continue ;};if len (_ded )< 4{continue ;};_cfcf :=string (_ded [:4]);if _cfcf =="\u004f\u0054\u0054\u004f"{continue ;};if _cfcf !="\u0000\u0001\u0000\u0000"&&_cfcf !="\u0074\u0072\u0075\u0065"{continue ;};_gdd ,_acba :=_fc .Parse (_ef .NewReader (_ded ));
if _acba !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_acba );continue ;};_acba =_gdd .Optimize ();
if _acba !=nil {_b .Log .Debug ("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067",_acba );continue ;};var _bffb _ef .Buffer ;_acba =_gdd .Write (&_bffb );
if _acba !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_acba );continue ;};if _bffb .Len ()> len (_ded ){_b .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
continue ;};_cde ,_acba :=_bf .MakeStream (_bffb .Bytes (),_bf .NewFlateEncoder ());if _acba !=nil {continue ;};*_febf =*_cde ;_febf .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_bf .MakeInteger (int64 (_bffb .Len ())));};return objects ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_gead *CompressStreams )Optimize (objects []_bf .PdfObject )(_fgf []_bf .PdfObject ,_gae error ){_fgf =make ([]_bf .PdfObject ,len (objects ));copy (_fgf ,objects );for _ ,_gfb :=range objects {_daba ,_fef :=_bf .GetStream (_gfb );if !_fef {continue ;
};if _cga :=_daba .Get ("\u0046\u0069\u006c\u0074\u0065\u0072");_cga !=nil {if _ ,_aeea :=_bf .GetName (_cga );_aeea {continue ;};if _bbb ,_daf :=_bf .GetArray (_cga );_daf &&_bbb .Len ()> 0{continue ;};};_gag :=_bf .NewFlateEncoder ();var _aada []byte ;
_aada ,_gae =_gag .EncodeBytes (_daba .Stream );if _gae !=nil {return _fgf ,_gae ;};_deaee :=_gag .MakeStreamDict ();if len (_aada )+len (_deaee .WriteString ())< len (_daba .Stream ){_daba .Stream =_aada ;_daba .PdfObjectDictionary .Merge (_deaee );_daba .PdfObjectDictionary .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_bf .MakeInteger (int64 (len (_daba .Stream ))));
};};return _fgf ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_cad *CleanContentstream )Optimize (objects []_bf .PdfObject )(_ed []_bf .PdfObject ,_dac error ){_dc :=map[*_bf .PdfObjectStream ]struct{}{};var _abb []*_bf .PdfObjectStream ;_efa :=func (_ac *_bf .PdfObjectStream ){if _ ,_feb :=_dc [_ac ];!_feb {_dc [_ac ]=struct{}{};
_abb =append (_abb ,_ac );};};_ebb :=map[_bf .PdfObject ]bool {};_bb :=map[_bf .PdfObject ]bool {};for _ ,_acb :=range objects {switch _fca :=_acb .(type ){case *_bf .PdfIndirectObject :switch _bad :=_fca .PdfObject .(type ){case *_bf .PdfObjectDictionary :if _bbd ,_gf :=_bf .GetName (_bad .Get ("\u0054\u0079\u0070\u0065"));
!_gf ||_bbd .String ()!="\u0050\u0061\u0067\u0065"{continue ;};if _ea ,_ccb :=_bf .GetStream (_bad .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_ccb {_efa (_ea );}else if _deb ,_ae :=_bf .GetArray (_bad .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
_ae {var _edc []*_bf .PdfObjectStream ;for _ ,_ad :=range _deb .Elements (){if _dg ,_dcc :=_bf .GetStream (_ad );_dcc {_edc =append (_edc ,_dg );};};if len (_edc )> 0{var _af _ef .Buffer ;for _ ,_fd :=range _edc {if _bdc ,_bbf :=_bf .DecodeStream (_fd );
_bbf ==nil {_af .Write (_bdc );};_ebb [_fd ]=true ;};_eg ,_dd :=_bf .MakeStream (_af .Bytes (),_bf .NewFlateEncoder ());if _dd !=nil {return nil ,_dd ;};_bb [_eg ]=true ;_bad .Set ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073",_eg );_efa (_eg );};
};};case *_bf .PdfObjectStream :if _ge ,_bbg :=_bf .GetName (_fca .Get ("\u0054\u0079\u0070\u0065"));!_bbg ||_ge .String ()!="\u0058O\u0062\u006a\u0065\u0063\u0074"{continue ;};if _dea ,_cd :=_bf .GetName (_fca .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
!_cd ||_dea .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_efa (_fca );};};for _ ,_gff :=range _abb {_dac =_ffc (_gff );if _dac !=nil {return nil ,_dac ;};};_ed =nil ;for _ ,_cfc :=range objects {if _ebb [_cfc ]{continue ;};_ed =append (_ed ,_cfc );
};for _fg :=range _bb {_ed =append (_ed ,_fg );};return _ed ,nil ;};type content struct{_aec string ;_bea *_eb .PdfPageResources ;};func _ffc (_ccc *_bf .PdfObjectStream )error {_fee ,_fce :=_bf .DecodeStream (_ccc );if _fce !=nil {return _fce ;};_cab :=_c .NewContentStreamParser (string (_fee ));
_bdd ,_fce :=_cab .Parse ();if _fce !=nil {return _fce ;};_bdd =_ff (_bdd );_fba :=_bdd .Bytes ();if len (_fba )>=len (_fee ){return nil ;};_cf ,_fce :=_bf .MakeStream (_bdd .Bytes (),_bf .NewFlateEncoder ());if _fce !=nil {return _fce ;};_ccc .Stream =_cf .Stream ;
_ccc .Merge (_cf .PdfObjectDictionary );return nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_fdb *CombineIdenticalIndirectObjects )Optimize (objects []_bf .PdfObject )(_bcg []_bf .PdfObject ,_dec error ){_deaa (objects );_dgf :=make (map[_bf .PdfObject ]_bf .PdfObject );_cbeg :=make (map[_bf .PdfObject ]struct{});_aff :=make (map[string ][]*_bf .PdfIndirectObject );
for _ ,_fdec :=range objects {_ceb ,_ecc :=_fdec .(*_bf .PdfIndirectObject );if !_ecc {continue ;};if _agbb ,_egbg :=_ceb .PdfObject .(*_bf .PdfObjectDictionary );_egbg {if _cdd ,_fbaf :=_agbb .Get ("\u0054\u0079\u0070\u0065").(*_bf .PdfObjectName );_fbaf &&*_cdd =="\u0050\u0061\u0067\u0065"{continue ;
};_egc :=_da .New ();_egc .Write ([]byte (_agbb .WriteString ()));_ged :=string (_egc .Sum (nil ));_aff [_ged ]=append (_aff [_ged ],_ceb );};};for _ ,_cgbc :=range _aff {if len (_cgbc )< 2{continue ;};_eccd :=_cgbc [0];for _cba :=1;_cba < len (_cgbc );
_cba ++{_eac :=_cgbc [_cba ];_dgf [_eac ]=_eccd ;_cbeg [_eac ]=struct{}{};};};_bcg =make ([]_bf .PdfObject ,0,len (objects )-len (_cbeg ));for _ ,_aab :=range objects {if _ ,_cfa :=_cbeg [_aab ];_cfa {continue ;};_bcg =append (_bcg ,_aab );};_gddg (_bcg ,_dgf );
return _bcg ,nil ;};func _deaa (_gecc []_bf .PdfObject ){for _gfec ,_faff :=range _gecc {switch _ecdg :=_faff .(type ){case *_bf .PdfIndirectObject :_ecdg .ObjectNumber =int64 (_gfec +1);_ecdg .GenerationNumber =0;case *_bf .PdfObjectStream :_ecdg .ObjectNumber =int64 (_gfec +1);
_ecdg .GenerationNumber =0;case *_bf .PdfObjectStreams :_ecdg .ObjectNumber =int64 (_gfec +1);_ecdg .GenerationNumber =0;};};};

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{};

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{};func _egb (_efg *_bf .PdfObjectStream ,_ffd []rune ,_dca []_fc .GlyphIndex )error {_efg ,_ebd :=_bf .GetStream (_efg );if !_ebd {_b .Log .Debug ("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067");
return _df .New ("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_fcee ,_efgd :=_bf .DecodeStream (_efg );if _efgd !=nil {_b .Log .Debug ("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_efgd );
return _efgd ;};_cfd ,_efgd :=_fc .Parse (_ef .NewReader (_fcee ));if _efgd !=nil {_b .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074",len (_efg .Stream ));
return _efgd ;};_cee :=_dca ;if len (_ffd )> 0{_aef :=_cfd .LookupRunes (_ffd );_cee =append (_cee ,_aef ...);};_cfd ,_efgd =_cfd .SubsetKeepIndices (_cee );if _efgd !=nil {_b .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v",_efgd );
return _efgd ;};var _aac _ef .Buffer ;_efgd =_cfd .Write (&_aac );if _efgd !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_efgd );return _efgd ;};if _aac .Len ()> len (_fcee ){_b .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
return nil ;};_bfg ,_efgd :=_bf .MakeStream (_aac .Bytes (),_bf .NewFlateEncoder ());if _efgd !=nil {_b .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_efgd );return _efgd ;
};*_efg =*_bfg ;_efg .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_bf .MakeInteger (int64 (_aac .Len ())));return nil ;};

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{};

// Optimize optimizes PDF objects to decrease PDF size.
func (_fb *Chain )Optimize (objects []_bf .PdfObject )(_bd []_bf .PdfObject ,_g error ){_gg :=objects ;for _ ,_ga :=range _fb ._cc {_ag ,_fbg :=_ga .Optimize (_gg );if _fbg !=nil {_b .Log .Debug ("\u0045\u0052\u0052OR\u0020\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u002b\u0076",_fbg );
continue ;};_gg =_ag ;};return _gg ,nil ;};

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ImageQuality int ;};

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{};