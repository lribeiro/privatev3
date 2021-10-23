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

// Package fdf provides support for loading form field data from Form Field Data (FDF) files.
package fdf ;import (_cg "bufio";_c "bytes";_ag "encoding/hex";_a "errors";_bg "fmt";_gb "github.com/unidoc/unipdf/v3/common";_ba "github.com/unidoc/unipdf/v3/core";_dg "io";_e "os";_dc "regexp";_ce "sort";_d "strconv";_b "strings";);

// Load loads FDF form data from `r`.
func Load (r _dg .ReadSeeker )(*Data ,error ){_be ,_dce :=_cge (r );if _dce !=nil {return nil ,_dce ;};_de ,_dce :=_be .Root ();if _dce !=nil {return nil ,_dce ;};_bgg ,_f :=_ba .GetArray (_de .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_f {return nil ,_a .New ("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067");
};return &Data {_ge :_bgg ,_ac :_de },nil ;};

// Root returns the Root of the FDF document.
func (_gbe *fdfParser )Root ()(*_ba .PdfObjectDictionary ,error ){if _gbe ._bfa !=nil {if _bbcc ,_bfg :=_gbe .trace (_gbe ._bfa .Get ("\u0052\u006f\u006f\u0074")).(*_ba .PdfObjectDictionary );_bfg {if _ddgd ,_dbfe :=_gbe .trace (_bbcc .Get ("\u0046\u0044\u0046")).(*_ba .PdfObjectDictionary );
_dbfe {return _ddgd ,nil ;};};};var _cgd []int64 ;for _bac :=range _gbe ._egec {_cgd =append (_cgd ,_bac );};_ce .Slice (_cgd ,func (_cff ,_aceb int )bool {return _cgd [_cff ]< _cgd [_aceb ]});for _ ,_faed :=range _cgd {_cce :=_gbe ._egec [_faed ];if _dcf ,_abe :=_gbe .trace (_cce ).(*_ba .PdfObjectDictionary );
_abe {if _bbd ,_gcb :=_gbe .trace (_dcf .Get ("\u0046\u0044\u0046")).(*_ba .PdfObjectDictionary );_gcb {return _bbd ,nil ;};};};return nil ,_a .New ("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};func (_bda *fdfParser )parseString ()(*_ba .PdfObjectString ,error ){_bda ._fde .ReadByte ();
var _ggc _c .Buffer ;_afa :=1;for {_ed ,_gefd :=_bda ._fde .Peek (1);if _gefd !=nil {return _ba .MakeString (_ggc .String ()),_gefd ;};if _ed [0]=='\\'{_bda ._fde .ReadByte ();_gbd ,_acd :=_bda ._fde .ReadByte ();if _acd !=nil {return _ba .MakeString (_ggc .String ()),_acd ;
};if _ba .IsOctalDigit (_gbd ){_abg ,_agd :=_bda ._fde .Peek (2);if _agd !=nil {return _ba .MakeString (_ggc .String ()),_agd ;};var _bad []byte ;_bad =append (_bad ,_gbd );for _ ,_fab :=range _abg {if _ba .IsOctalDigit (_fab ){_bad =append (_bad ,_fab );
}else {break ;};};_bda ._fde .Discard (len (_bad )-1);_gb .Log .Trace ("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022",_bad );_fag ,_agd :=_d .ParseUint (string (_bad ),8,32);if _agd !=nil {return _ba .MakeString (_ggc .String ()),_agd ;
};_ggc .WriteByte (byte (_fag ));continue ;};switch _gbd {case 'n':_ggc .WriteRune ('\n');case 'r':_ggc .WriteRune ('\r');case 't':_ggc .WriteRune ('\t');case 'b':_ggc .WriteRune ('\b');case 'f':_ggc .WriteRune ('\f');case '(':_ggc .WriteRune ('(');case ')':_ggc .WriteRune (')');
case '\\':_ggc .WriteRune ('\\');};continue ;}else if _ed [0]=='('{_afa ++;}else if _ed [0]==')'{_afa --;if _afa ==0{_bda ._fde .ReadByte ();break ;};};_edc ,_ :=_bda ._fde .ReadByte ();_ggc .WriteByte (_edc );};return _ba .MakeString (_ggc .String ()),nil ;
};func (_fac *fdfParser )parseBool ()(_ba .PdfObjectBool ,error ){_fgb ,_fcbae :=_fac ._fde .Peek (4);if _fcbae !=nil {return _ba .PdfObjectBool (false ),_fcbae ;};if (len (_fgb )>=4)&&(string (_fgb [:4])=="\u0074\u0072\u0075\u0065"){_fac ._fde .Discard (4);
return _ba .PdfObjectBool (true ),nil ;};_fgb ,_fcbae =_fac ._fde .Peek (5);if _fcbae !=nil {return _ba .PdfObjectBool (false ),_fcbae ;};if (len (_fgb )>=5)&&(string (_fgb [:5])=="\u0066\u0061\u006cs\u0065"){_fac ._fde .Discard (5);return _ba .PdfObjectBool (false ),nil ;
};return _ba .PdfObjectBool (false ),_a .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};func (_ddg *fdfParser )parseIndirectObject ()(_ba .PdfObject ,error ){_bbc :=_ba .PdfIndirectObject {};
_gb .Log .Trace ("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a");_dfc ,_cfg :=_ddg ._fde .Peek (20);if _cfg !=nil {_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return &_bbc ,_cfg ;};_gb .Log .Trace ("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022",string (_dfc ));_gdfb :=_fc .FindStringSubmatchIndex (string (_dfc ));if len (_gdfb )< 6{_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_dfc ));
return &_bbc ,_a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_ddg ._fde .Discard (_gdfb [0]);_gb .Log .Trace ("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064",_gdfb );_aad :=_gdfb [1]-_gdfb [0];_egf :=make ([]byte ,_aad );_ ,_cfg =_ddg .readAtLeast (_egf ,_aad );if _cfg !=nil {_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073",_cfg );
return nil ,_cfg ;};_gb .Log .Trace ("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073",_egf );_ded :=_fc .FindStringSubmatch (string (_egf ));if len (_ded )< 3{_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_egf ));
return &_bbc ,_a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_feac ,_ :=_d .Atoi (_ded [1]);_fdcd ,_ :=_d .Atoi (_ded [2]);_bbc .ObjectNumber =int64 (_feac );_bbc .GenerationNumber =int64 (_fdcd );for {_eac ,_ada :=_ddg ._fde .Peek (2);if _ada !=nil {return &_bbc ,_ada ;};_gb .Log .Trace ("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_eac ),string (_eac ));
if _ba .IsWhiteSpace (_eac [0]){_ddg .skipSpaces ();}else if _eac [0]=='%'{_ddg .skipComments ();}else if (_eac [0]=='<')&&(_eac [1]=='<'){_gb .Log .Trace ("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074");_bbc .PdfObject ,_ada =_ddg .parseDict ();
_gb .Log .Trace ("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076",_ada );if _ada !=nil {return &_bbc ,_ada ;};_gb .Log .Trace ("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e");
}else if (_eac [0]=='/')||(_eac [0]=='(')||(_eac [0]=='[')||(_eac [0]=='<'){_bbc .PdfObject ,_ada =_ddg .parseObject ();if _ada !=nil {return &_bbc ,_ada ;};_gb .Log .Trace ("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e");
}else {if _eac [0]=='e'{_cfbd ,_aba :=_ddg .readTextLine ();if _aba !=nil {return nil ,_aba ;};if len (_cfbd )>=6&&_cfbd [0:6]=="\u0065\u006e\u0064\u006f\u0062\u006a"{break ;};}else if _eac [0]=='s'{_eac ,_ =_ddg ._fde .Peek (10);if string (_eac [:6])=="\u0073\u0074\u0072\u0065\u0061\u006d"{_aef :=6;
if len (_eac )> 6{if _ba .IsWhiteSpace (_eac [_aef ])&&_eac [_aef ]!='\r'&&_eac [_aef ]!='\n'{_gb .Log .Debug ("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072");
_aef ++;};if _eac [_aef ]=='\r'{_aef ++;if _eac [_aef ]=='\n'{_aef ++;};}else if _eac [_aef ]=='\n'{_aef ++;};};_ddg ._fde .Discard (_aef );_adc ,_effb :=_bbc .PdfObject .(*_ba .PdfObjectDictionary );if !_effb {return nil ,_a .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};_gb .Log .Trace ("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073",_adc );_aee ,_add :=_adc .Get ("\u004c\u0065\u006e\u0067\u0074\u0068").(*_ba .PdfObjectInteger );if !_add {return nil ,_a .New ("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072");
};_fbe :=*_aee ;if _fbe < 0{return nil ,_a .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030");};if int64 (_fbe )> _ddg ._gca {_gb .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065");
return nil ,_a .New ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065");};_dae :=make ([]byte ,_fbe );
_ ,_ada =_ddg .readAtLeast (_dae ,int (_fbe ));if _ada !=nil {_gb .Log .Debug ("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058",len (_dae ),_dae );_gb .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ada );
return nil ,_ada ;};_abd :=_ba .PdfObjectStream {};_abd .Stream =_dae ;_abd .PdfObjectDictionary =_bbc .PdfObject .(*_ba .PdfObjectDictionary );_abd .ObjectNumber =_bbc .ObjectNumber ;_abd .GenerationNumber =_bbc .GenerationNumber ;_ddg .skipSpaces ();
_ddg ._fde .Discard (9);_ddg .skipSpaces ();return &_abd ,nil ;};};_bbc .PdfObject ,_ada =_ddg .parseObject ();return &_bbc ,_ada ;};};_gb .Log .Trace ("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021");
return &_bbc ,nil ;};

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data )FieldDictionaries ()(map[string ]*_ba .PdfObjectDictionary ,error ){_bd :=map[string ]*_ba .PdfObjectDictionary {};for _ca :=0;_ca < fdf ._ge .Len ();_ca ++{_deb ,_ae :=_ba .GetDict (fdf ._ge .Get (_ca ));if _ae {_bc ,_ :=_ba .GetString (_deb .Get ("\u0054"));
if _bc !=nil {_bd [_bc .Str ()]=_deb ;};};};return _bd ,nil ;};func (_dcg *fdfParser )seekToEOFMarker (_fcda int64 )error {_dfe :=int64 (0);_dagd :=int64 (1000);for _dfe < _fcda {if _fcda <=(_dagd +_dfe ){_dagd =_fcda -_dfe ;};_ ,_fcbac :=_dcg ._ga .Seek (-_dfe -_dagd ,_dg .SeekEnd );
if _fcbac !=nil {return _fcbac ;};_faga :=make ([]byte ,_dagd );_dcg ._ga .Read (_faga );_gb .Log .Trace ("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022",string (_faga ));
_cfa :=_abf .FindAllStringIndex (string (_faga ),-1);if _cfa !=nil {_eag :=_cfa [len (_cfa )-1];_gb .Log .Trace ("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064",_cfa );_dcg ._ga .Seek (-_dfe -_dagd +int64 (_eag [0]),_dg .SeekEnd );return nil ;};_gb .Log .Debug ("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067");
_dfe +=_dagd ;};_gb .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e");return _a .New ("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");
};func (_debd *fdfParser )parseNull ()(_ba .PdfObjectNull ,error ){_ ,_edf :=_debd ._fde .Discard (4);return _ba .PdfObjectNull {},_edf ;};type fdfParser struct{_dgb int ;_fa int ;_egec map[int64 ]_ba .PdfObject ;_ga _dg .ReadSeeker ;_fde *_cg .Reader ;
_gca int64 ;_bfa *_ba .PdfObjectDictionary ;};func (_bbb *fdfParser )seekFdfVersionTopDown ()(int ,int ,error ){_bbb ._ga .Seek (0,_dg .SeekStart );_bbb ._fde =_cg .NewReader (_bbb ._ga );_def :=20;_cfe :=make ([]byte ,_def );for {_cfef ,_cbf :=_bbb ._fde .ReadByte ();
if _cbf !=nil {if _cbf ==_dg .EOF {break ;}else {return 0,0,_cbf ;};};if _ba .IsDecimalDigit (_cfef )&&_cfe [_def -1]=='.'&&_ba .IsDecimalDigit (_cfe [_def -2])&&_cfe [_def -3]=='-'&&_cfe [_def -4]=='F'&&_cfe [_def -5]=='D'&&_cfe [_def -6]=='P'{_bdab :=int (_cfe [_def -2]-'0');
_ecc :=int (_cfef -'0');return _bdab ,_ecc ,nil ;};_cfe =append (_cfe [1:_def ],_cfef );};return 0,0,_a .New ("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};func (_dagb *fdfParser )parseDict ()(*_ba .PdfObjectDictionary ,error ){_gb .Log .Trace ("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021");
_fca :=_ba .MakeDict ();_ede ,_ :=_dagb ._fde .ReadByte ();if _ede !='<'{return nil ,_a .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};_ede ,_ =_dagb ._fde .ReadByte ();if _ede !='<'{return nil ,_a .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");
};for {_dagb .skipSpaces ();_dagb .skipComments ();_fcd ,_gfe :=_dagb ._fde .Peek (2);if _gfe !=nil {return nil ,_gfe ;};_gb .Log .Trace ("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_fcd ),string (_fcd ));
if (_fcd [0]=='>')&&(_fcd [1]=='>'){_gb .Log .Trace ("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");_dagb ._fde .ReadByte ();_dagb ._fde .ReadByte ();break ;};_gb .Log .Trace ("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021");
_gcagg ,_gfe :=_dagb .parseName ();_gb .Log .Trace ("\u004be\u0079\u003a\u0020\u0025\u0073",_gcagg );if _gfe !=nil {_gb .Log .Debug ("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073",_gfe );
return nil ,_gfe ;};if len (_gcagg )> 4&&_gcagg [len (_gcagg )-4:]=="\u006e\u0075\u006c\u006c"{_gbdd :=_gcagg [0:len (_gcagg )-4];_gb .Log .Debug ("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029",_gcagg );
_gb .Log .Debug ("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c",_gbdd );_dagb .skipSpaces ();_bag ,_ :=_dagb ._fde .Peek (1);if _bag [0]=='/'{_fca .Set (_gbdd ,_ba .MakeNull ());continue ;};};
_dagb .skipSpaces ();_fagb ,_gfe :=_dagb .parseObject ();if _gfe !=nil {return nil ,_gfe ;};_fca .Set (_gcagg ,_fagb );_gb .Log .Trace ("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073",_gcagg ,_fagb .String ());};_gb .Log .Trace ("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021");
return _fca ,nil ;};var _fe =_dc .MustCompile ("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029");func (_acg *fdfParser )parseArray ()(*_ba .PdfObjectArray ,error ){_aa :=_ba .MakeArray ();
_acg ._fde .ReadByte ();for {_acg .skipSpaces ();_dda ,_ff :=_acg ._fde .Peek (1);if _ff !=nil {return _aa ,_ff ;};if _dda [0]==']'{_acg ._fde .ReadByte ();break ;};_decb ,_ff :=_acg .parseObject ();if _ff !=nil {return _aa ,_ff ;};_aa .Append (_decb );
};return _aa ,nil ;};func _cag (_cdg string )(*fdfParser ,error ){_bba :=fdfParser {};_fge :=[]byte (_cdg );_bggc :=_c .NewReader (_fge );_bba ._ga =_bggc ;_bba ._egec =map[int64 ]_ba .PdfObject {};_fgbe :=_cg .NewReader (_bggc );_bba ._fde =_fgbe ;_bba ._gca =int64 (len (_cdg ));
return &_bba ,_bba .parse ();};func (_ab *fdfParser )getFileOffset ()int64 {_bec ,_ :=_ab ._ga .Seek (0,_dg .SeekCurrent );_bec -=int64 (_ab ._fde .Buffered ());return _bec ;};

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath (fdfPath string )(*Data ,error ){_eg ,_gc :=_e .Open (fdfPath );if _gc !=nil {return nil ,_gc ;};defer _eg .Close ();return Load (_eg );};func (_fg *fdfParser )readComment ()(string ,error ){var _bed _c .Buffer ;_ ,_feeg :=_fg .skipSpaces ();
if _feeg !=nil {return _bed .String (),_feeg ;};_bb :=true ;for {_abfb ,_bde :=_fg ._fde .Peek (1);if _bde !=nil {_gb .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_bde .Error ());return _bed .String (),_bde ;};if _bb &&_abfb [0]!='%'{return _bed .String (),_a .New ("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025");
};_bb =false ;if (_abfb [0]!='\r')&&(_abfb [0]!='\n'){_dfa ,_ :=_fg ._fde .ReadByte ();_bed .WriteByte (_dfa );}else {break ;};};return _bed .String (),nil ;};func (_bdba *fdfParser )parseFdfVersion ()(int ,int ,error ){_bdba ._ga .Seek (0,_dg .SeekStart );
_fbb :=20;_egc :=make ([]byte ,_fbb );_bdba ._ga .Read (_egc );_agdb :=_egb .FindStringSubmatch (string (_egc ));if len (_agdb )< 3{_ace ,_ggge ,_ade :=_bdba .seekFdfVersionTopDown ();if _ade !=nil {_gb .Log .Debug ("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e");
return 0,0,_ade ;};return _ace ,_ggge ,nil ;};_cfb ,_gedc :=_d .Atoi (_agdb [1]);if _gedc !=nil {return 0,0,_gedc ;};_ffg ,_gedc :=_d .Atoi (_agdb [2]);if _gedc !=nil {return 0,0,_gedc ;};_gb .Log .Debug ("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064",_cfb ,_ffg );
return _cfb ,_ffg ,nil ;};func (_gbb *fdfParser )parseName ()(_ba .PdfObjectName ,error ){var _gcag _c .Buffer ;_aec :=false ;for {_afb ,_cf :=_gbb ._fde .Peek (1);if _cf ==_dg .EOF {break ;};if _cf !=nil {return _ba .PdfObjectName (_gcag .String ()),_cf ;
};if !_aec {if _afb [0]=='/'{_aec =true ;_gbb ._fde .ReadByte ();}else if _afb [0]=='%'{_gbb .readComment ();_gbb .skipSpaces ();}else {_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029",_afb ,_afb );
return _ba .PdfObjectName (_gcag .String ()),_bg .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029",_afb [0]);};}else {if _ba .IsWhiteSpace (_afb [0]){break ;}else if (_afb [0]=='/')||(_afb [0]=='[')||(_afb [0]=='(')||(_afb [0]==']')||(_afb [0]=='<')||(_afb [0]=='>'){break ;
}else if _afb [0]=='#'{_faec ,_cc :=_gbb ._fde .Peek (3);if _cc !=nil {return _ba .PdfObjectName (_gcag .String ()),_cc ;};_gbb ._fde .Discard (3);_db ,_cc :=_ag .DecodeString (string (_faec [1:3]));if _cc !=nil {return _ba .PdfObjectName (_gcag .String ()),_cc ;
};_gcag .Write (_db );}else {_gef ,_ :=_gbb ._fde .ReadByte ();_gcag .WriteByte (_gef );};};};return _ba .PdfObjectName (_gcag .String ()),nil ;};func (_daa *fdfParser )trace (_aebf _ba .PdfObject )_ba .PdfObject {switch _cagc :=_aebf .(type ){case *_ba .PdfObjectReference :_gdfa ,_fgf :=_daa ._egec [_cagc .ObjectNumber ].(*_ba .PdfIndirectObject );
if _fgf {return _gdfa .PdfObject ;};_gb .Log .Debug ("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");return nil ;case *_ba .PdfIndirectObject :return _cagc .PdfObject ;};return _aebf ;};

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data )FieldValues ()(map[string ]_ba .PdfObject ,error ){_ef ,_cgg :=fdf .FieldDictionaries ();if _cgg !=nil {return nil ,_cgg ;};var _cd []string ;for _fd :=range _ef {_cd =append (_cd ,_fd );};_ce .Strings (_cd );_aeg :=map[string ]_ba .PdfObject {};
for _ ,_dfd :=range _cd {_gf :=_ef [_dfd ];_bf :=_ba .TraceToDirectObject (_gf .Get ("\u0056"));_aeg [_dfd ]=_bf ;};return _aeg ,nil ;};

// Data represents forms data format (FDF) file data.
type Data struct{_ac *_ba .PdfObjectDictionary ;_ge *_ba .PdfObjectArray ;};func (_dag *fdfParser )skipComments ()error {if _ ,_gad :=_dag .skipSpaces ();_gad !=nil {return _gad ;};_dca :=true ;for {_gff ,_fee :=_dag ._fde .Peek (1);if _fee !=nil {_gb .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_fee .Error ());
return _fee ;};if _dca &&_gff [0]!='%'{return nil ;};_dca =false ;if (_gff [0]!='\r')&&(_gff [0]!='\n'){_dag ._fde .ReadByte ();}else {break ;};};return _dag .skipComments ();};func (_dd *fdfParser )parseNumber ()(_ba .PdfObject ,error ){return _ba .ParseNumber (_dd ._fde )};
var _egb =_dc .MustCompile ("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029");func (_fea *fdfParser )parseObject ()(_ba .PdfObject ,error ){_gb .Log .Trace ("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074");
_fea .skipSpaces ();for {_dgd ,_fbd :=_fea ._fde .Peek (2);if _fbd !=nil {return nil ,_fbd ;};_gb .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_dgd ));if _dgd [0]=='/'{_bcf ,_egbb :=_fea .parseName ();
_gb .Log .Trace ("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027",_bcf );return &_bcf ,_egbb ;}else if _dgd [0]=='('{_gb .Log .Trace ("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!");return _fea .parseString ();}else if _dgd [0]=='['{_gb .Log .Trace ("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021");
return _fea .parseArray ();}else if (_dgd [0]=='<')&&(_dgd [1]=='<'){_gb .Log .Trace ("\u002d>\u0044\u0069\u0063\u0074\u0021");return _fea .parseDict ();}else if _dgd [0]=='<'{_gb .Log .Trace ("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021");
return _fea .parseHexString ();}else if _dgd [0]=='%'{_fea .readComment ();_fea .skipSpaces ();}else {_gb .Log .Trace ("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f");_dgd ,_ =_fea ._fde .Peek (15);_bfd :=string (_dgd );
_gb .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_bfd );if (len (_bfd )> 3)&&(_bfd [:4]=="\u006e\u0075\u006c\u006c"){_gdf ,_dgf :=_fea .parseNull ();return &_gdf ,_dgf ;}else if (len (_bfd )> 4)&&(_bfd [:5]=="\u0066\u0061\u006cs\u0065"){_debde ,_fcbb :=_fea .parseBool ();
return &_debde ,_fcbb ;}else if (len (_bfd )> 3)&&(_bfd [:4]=="\u0074\u0072\u0075\u0065"){_ee ,_cadg :=_fea .parseBool ();return &_ee ,_cadg ;};_dea :=_ggg .FindStringSubmatch (_bfd );if len (_dea )> 1{_dgd ,_ =_fea ._fde .ReadBytes ('R');_gb .Log .Trace ("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027",string (_dgd [:]));
_cac ,_dcea :=_bdg (string (_dgd ));return &_cac ,_dcea ;};_dbf :=_ged .FindStringSubmatch (_bfd );if len (_dbf )> 1{_gb .Log .Trace ("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021");return _fea .parseNumber ();};_dbf =_fe .FindStringSubmatch (_bfd );
if len (_dbf )> 1{_gb .Log .Trace ("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021");_gb .Log .Trace ("\u0025\u0020\u0073",_dbf );return _fea .parseNumber ();};_gb .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)",_bfd );
return nil ,_a .New ("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e");
};};};func (_af *fdfParser )skipSpaces ()(int ,error ){_fce :=0;for {_bggg ,_decg :=_af ._fde .ReadByte ();if _decg !=nil {return 0,_decg ;};if _ba .IsWhiteSpace (_bggg ){_fce ++;}else {_af ._fde .UnreadByte ();break ;};};return _fce ,nil ;};func _cge (_ccf _dg .ReadSeeker )(*fdfParser ,error ){_dgg :=&fdfParser {};
_dgg ._ga =_ccf ;_dgg ._egec =map[int64 ]_ba .PdfObject {};_ddb ,_edcg ,_aeb :=_dgg .parseFdfVersion ();if _aeb !=nil {_gb .Log .Error ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076",_aeb );
return nil ,_aeb ;};_dgg ._dgb =_ddb ;_dgg ._fa =_edcg ;_aeb =_dgg .parse ();return _dgg ,_aeb ;};var _ged =_dc .MustCompile ("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029");func (_da *fdfParser )readAtLeast (_dec []byte ,_bgb int )(int ,error ){_fb :=_bgb ;
_deba :=0;_bgf :=0;for _fb > 0{_cef ,_cde :=_da ._fde .Read (_dec [_deba :]);if _cde !=nil {_gb .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073",_cef ,_bgf ,_cde .Error ());
return _deba ,_a .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067");};_bgf ++;_deba +=_cef ;_fb -=_cef ;};return _deba ,nil ;};var _fc =_dc .MustCompile ("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a");
func (_ea *fdfParser )readTextLine ()(string ,error ){var _gd _c .Buffer ;for {_cad ,_beb :=_ea ._fde .Peek (1);if _beb !=nil {_gb .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_beb .Error ());return _gd .String (),_beb ;};if (_cad [0]!='\r')&&(_cad [0]!='\n'){_fae ,_ :=_ea ._fde .ReadByte ();
_gd .WriteByte (_fae );}else {break ;};};return _gd .String (),nil ;};var _abf =_dc .MustCompile ("\u0025\u0025\u0045O\u0046");func _bdg (_bgc string )(_ba .PdfObjectReference ,error ){_ggce :=_ba .PdfObjectReference {};_fga :=_ggg .FindStringSubmatch (_bgc );
if len (_fga )< 3{_gb .Log .Debug ("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065");return _ggce ,_a .New ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e");
};_bef ,_daga :=_d .Atoi (_fga [1]);if _daga !=nil {_gb .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030",_fga [1]);
return _ggce ,nil ;};_ggce .ObjectNumber =int64 (_bef );_bdb ,_daga :=_d .Atoi (_fga [2]);if _daga !=nil {_gb .Log .Debug ("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030",_fga [2]);
return _ggce ,nil ;};_ggce .GenerationNumber =int64 (_bdb );return _ggce ,nil ;};var _ggg =_dc .MustCompile ("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052");func (_dfad *fdfParser )parseHexString ()(*_ba .PdfObjectString ,error ){_dfad ._fde .ReadByte ();
var _dcd _c .Buffer ;for {_fcb ,_dcb :=_dfad ._fde .Peek (1);if _dcb !=nil {return _ba .MakeHexString (""),_dcb ;};if _fcb [0]=='>'{_dfad ._fde .ReadByte ();break ;};_eb ,_ :=_dfad ._fde .ReadByte ();if !_ba .IsWhiteSpace (_eb ){_dcd .WriteByte (_eb );
};};if _dcd .Len ()%2==1{_dcd .WriteRune ('0');};_fdc ,_fcba :=_ag .DecodeString (_dcd .String ());if _fcba !=nil {_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067",_dcd .String ());
return _ba .MakeHexString (""),nil ;};return _ba .MakeHexString (string (_fdc )),nil ;};func (_cb *fdfParser )parse ()error {_cb ._ga .Seek (0,_dg .SeekStart );_cb ._fde =_cg .NewReader (_cb ._ga );for {_cb .skipComments ();_fcef ,_gda :=_cb ._fde .Peek (20);
if _gda !=nil {_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");return _gda ;};if _b .HasPrefix (string (_fcef ),"\u0074r\u0061\u0069\u006c\u0065\u0072"){_cb ._fde .Discard (7);
_cb .skipSpaces ();_cb .skipComments ();_baga ,_ :=_cb .parseDict ();_cb ._bfa =_baga ;break ;};_bdd :=_fc .FindStringSubmatchIndex (string (_fcef ));if len (_bdd )< 6{_gb .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_fcef ));
return _a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_ec ,_gda :=_cb .parseIndirectObject ();if _gda !=nil {return _gda ;};switch _deg :=_ec .(type ){case *_ba .PdfIndirectObject :_cb ._egec [_deg .ObjectNumber ]=_deg ;case *_ba .PdfObjectStream :_cb ._egec [_deg .ObjectNumber ]=_deg ;default:return _a .New ("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};};return nil ;};func (_egef *fdfParser )setFileOffset (_gbc int64 ){_egef ._ga .Seek (_gbc ,_dg .SeekStart );_egef ._fde =_cg .NewReader (_egef ._ga );};