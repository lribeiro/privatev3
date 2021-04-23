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

package basic ;import _f "github.com/unidoc/unipdf/v3/internal/jbig2/errors";type Stack struct{Data []interface{};Aux *Stack ;};func (_ag IntsMap )Delete (key uint64 ){delete (_ag ,key )};type IntSlice []int ;func (_bcb *Stack )Pop ()(_ed interface{},_de bool ){_ed ,_de =_bcb .peek ();
if !_de {return nil ,_de ;};_bcb .Data =_bcb .Data [:_bcb .top ()];return _ed ,true ;};func Min (x ,y int )int {if x < y {return x ;};return y ;};func (_gd *Stack )peek ()(interface{},bool ){_fbg :=_gd .top ();if _fbg ==-1{return nil ,false ;};return _gd .Data [_fbg ],true ;
};func (_bc IntSlice )Get (index int )(int ,error ){if index > len (_bc )-1{return 0,_f .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );
};return _bc [index ],nil ;};func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};func (_ge *IntSlice )Copy ()*IntSlice {_ea :=IntSlice (make ([]int ,len (*_ge )));
copy (_ea ,*_ge );return &_ea ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func (_g IntsMap )Get (key uint64 )(int ,bool ){_b ,_e :=_g [key ];if !_e {return 0,false ;};if len (_b )==0{return 0,false ;};return _b [0],true ;};func (_geb NumSlice )GetInt (i int )(int ,error ){const _ecf ="\u0047\u0065\u0074\u0049\u006e\u0074";
if i < 0||i > len (_geb )-1{return 0,_f .Errorf (_ecf ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};_ba :=_geb [i ];return int (_ba +Sign (_ba )*0.5),nil ;
};type NumSlice []float32 ;func (_baf *Stack )Push (v interface{}){_baf .Data =append (_baf .Data ,v )};func Max (x ,y int )int {if x > y {return x ;};return y ;};func (_a IntsMap )Add (key uint64 ,value int ){_a [key ]=append (_a [key ],value )};type IntsMap map[uint64 ][]int ;
func (_eg *NumSlice )Add (v float32 ){*_eg =append (*_eg ,v )};func (_fg *NumSlice )AddInt (v int ){*_fg =append (*_fg ,float32 (v ))};func (_cd NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_cd )-1{return 0,_f .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};return _cd [i ],nil ;};func (_c IntsMap )GetSlice (key uint64 )([]int ,bool ){_ee ,_fb :=_c [key ];if !_fb {return nil ,false ;};return _ee ,true ;};func (_cc NumSlice )GetIntSlice ()[]int {_ab :=make ([]int ,len (_cc ));for _bca ,_bg :=range _cc {_ab [_bca ]=int (_bg );
};return _ab ;};func Abs (v int )int {if v > 0{return v ;};return -v ;};func NewIntSlice (i int )*IntSlice {_age :=IntSlice (make ([]int ,i ));return &_age };func (_daf *Stack )Len ()int {return len (_daf .Data )};func (_dd *IntSlice )Add (v int )error {if _dd ==nil {return _f .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");
};*_dd =append (*_dd ,v );return nil ;};func (_ga *Stack )top ()int {return len (_ga .Data )-1};func NewNumSlice (i int )*NumSlice {_da :=NumSlice (make ([]float32 ,i ));return &_da };func (_aa *Stack )Peek ()(_gc interface{},_eeg bool ){return _aa .peek ()};
func (_ec IntSlice )Size ()int {return len (_ec )};