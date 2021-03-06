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

package timeutils ;import (_e "errors";_ac "fmt";_ea "regexp";_ee "strconv";_c "time";);func ParsePdfTime (pdfTime string )(_c .Time ,error ){_cee :=_ff .FindAllStringSubmatch (pdfTime ,1);if len (_cee )< 1{return _c .Time {},_ac .Errorf ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029",pdfTime );
};if len (_cee [0])!=10{return _c .Time {},_e .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030");
};_bb ,_ :=_ee .ParseInt (_cee [0][1],10,32);_da ,_ :=_ee .ParseInt (_cee [0][2],10,32);_aa ,_ :=_ee .ParseInt (_cee [0][3],10,32);_ad ,_ :=_ee .ParseInt (_cee [0][4],10,32);_ge ,_ :=_ee .ParseInt (_cee [0][5],10,32);_bd ,_ :=_ee .ParseInt (_cee [0][6],10,32);
var (_ed byte ;_ab int64 ;_adg int64 ;);if len (_cee [0][7])> 0{_ed =_cee [0][7][0];}else {_ed ='+';};if len (_cee [0][8])> 0{_ab ,_ =_ee .ParseInt (_cee [0][8],10,32);}else {_ab =0;};if len (_cee [0][9])> 0{_adg ,_ =_ee .ParseInt (_cee [0][9],10,32);}else {_adg =0;
};_bc :=int (_ab *60*60+_adg *60);switch _ed {case '-':_bc =-_bc ;case 'Z':_bc =0;};_cd :=_ac .Sprintf ("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064",_ed ,_ab ,_adg );_eff :=_c .FixedZone (_cd ,_bc );return _c .Date (int (_bb ),_c .Month (_da ),int (_aa ),int (_ad ),int (_ge ),int (_bd ),0,_eff ),nil ;
};var _ff =_ea .MustCompile ("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f");
func FormatPdfTime (in _c .Time )string {_d :=in .Format ("\u002d\u0030\u0037\u003a\u0030\u0030");_b ,_ :=_ee .ParseInt (_d [1:3],10,32);_fg ,_ :=_ee .ParseInt (_d [4:6],10,32);_ce :=int64 (in .Year ());_ef :=int64 (in .Month ());_cg :=int64 (in .Day ());
_af :=int64 (in .Hour ());_fa :=int64 (in .Minute ());_fc :=int64 (in .Second ());_be :=_d [0];return _ac .Sprintf ("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027",_ce ,_ef ,_cg ,_af ,_fa ,_fc ,_be ,_b ,_fg );
};