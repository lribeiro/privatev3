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

package syncmap ;import _a "sync";func (_ac *ByteRuneMap )Read (b byte )(rune ,bool ){_ac ._d .RLock ();defer _ac ._d .RUnlock ();_de ,_e :=_ac ._f [b ];return _de ,_e ;};func (_dfa *StringRuneMap )Read (g string )(rune ,bool ){_dfa ._dge .RLock ();defer _dfa ._dge .RUnlock ();
_gee ,_bf :=_dfa ._gcd [g ];return _gee ,_bf ;};type StringsTuple struct{Key ,Value string ;};func (_df *RuneStringMap )Write (r rune ,s string ){_df ._aa .Lock ();defer _df ._aa .Unlock ();_df ._aga [r ]=s ;};type RuneStringMap struct{_aga map[rune ]string ;
_aa _a .RWMutex ;};func (_fg *RuneStringMap )Range (f func (_afd rune ,_cba string )(_dabg bool )){_fg ._aa .RLock ();defer _fg ._aa .RUnlock ();for _dae ,_gac :=range _fg ._aga {if f (_dae ,_gac ){break ;};};};type RuneByteMap struct{_def map[rune ]byte ;
_ec _a .RWMutex ;};func (_deg *RuneSet )Write (r rune ){_deg ._ffc .Lock ();defer _deg ._ffc .Unlock ();_deg ._dab [r ]=struct{}{};};func (_dag *RuneSet )Length ()int {_dag ._ffc .RLock ();defer _dag ._ffc .RUnlock ();return len (_dag ._dab );};type RuneUint16Map struct{_fa map[rune ]uint16 ;
_bb _a .RWMutex ;};func (_cbc *RuneUint16Map )Read (r rune )(uint16 ,bool ){_cbc ._bb .RLock ();defer _cbc ._bb .RUnlock ();_cbg ,_cd :=_cbc ._fa [r ];return _cbg ,_cd ;};func NewRuneStringMap (m map[rune ]string )*RuneStringMap {return &RuneStringMap {_aga :m }};
func (_b *RuneByteMap )Write (r rune ,b byte ){_b ._ec .Lock ();defer _b ._ec .Unlock ();_b ._def [r ]=b };func (_cec *RuneByteMap )Length ()int {_cec ._ec .RLock ();defer _cec ._ec .RUnlock ();return len (_cec ._def );};func (_fed *StringsMap )Write (g1 ,g2 string ){_fed ._cc .Lock ();
defer _fed ._cc .Unlock ();_fed ._gbc [g1 ]=g2 ;};func (_ced *ByteRuneMap )Length ()int {_ced ._d .RLock ();defer _ced ._d .RUnlock ();return len (_ced ._f )};type ByteRuneMap struct{_f map[byte ]rune ;_d _a .RWMutex ;};type StringsMap struct{_gbc map[string ]string ;
_cc _a .RWMutex ;};func MakeRuneUint16Map (length int )*RuneUint16Map {return &RuneUint16Map {_fa :make (map[rune ]uint16 ,length )};};func (_dd *RuneSet )Exists (r rune )bool {_dd ._ffc .RLock ();defer _dd ._ffc .RUnlock ();_ ,_eca :=_dd ._dab [r ];return _eca ;
};func (_ff *RuneByteMap )Range (f func (_geg rune ,_ae byte )(_cb bool )){_ff ._ec .RLock ();defer _ff ._ec .RUnlock ();for _fc ,_ee :=range _ff ._def {if f (_fc ,_ee ){break ;};};};func NewStringsMap (tuples []StringsTuple )*StringsMap {_efg :=map[string ]string {};
for _ ,_fcd :=range tuples {_efg [_fcd .Key ]=_fcd .Value ;};return &StringsMap {_gbc :_efg };};func (_ged *StringsMap )Copy ()*StringsMap {_ged ._cc .RLock ();defer _ged ._cc .RUnlock ();_agg :=map[string ]string {};for _eeg ,_bc :=range _ged ._gbc {_agg [_eeg ]=_bc ;
};return &StringsMap {_gbc :_agg };};func (_acg *StringRuneMap )Write (g string ,r rune ){_acg ._dge .Lock ();defer _acg ._dge .Unlock ();_acg ._gcd [g ]=r ;};func (_dc *ByteRuneMap )Write (b byte ,r rune ){_dc ._d .Lock ();defer _dc ._d .Unlock ();_dc ._f [b ]=r };
func NewStringRuneMap (m map[string ]rune )*StringRuneMap {return &StringRuneMap {_gcd :m }};func (_fe *StringsMap )Read (g string )(string ,bool ){_fe ._cc .RLock ();defer _fe ._cc .RUnlock ();_afc ,_faa :=_fe ._gbc [g ];return _afc ,_faa ;};func (_bd *RuneStringMap )Length ()int {_bd ._aa .RLock ();
defer _bd ._aa .RUnlock ();return len (_bd ._aga );};func MakeByteRuneMap (length int )*ByteRuneMap {return &ByteRuneMap {_f :make (map[byte ]rune ,length )}};func (_defc *RuneUint16Map )Length ()int {_defc ._bb .RLock ();defer _defc ._bb .RUnlock ();return len (_defc ._fa );
};type StringRuneMap struct{_gcd map[string ]rune ;_dge _a .RWMutex ;};func NewByteRuneMap (m map[byte ]rune )*ByteRuneMap {return &ByteRuneMap {_f :m }};func (_afe *StringRuneMap )Range (f func (_age string ,_fbd rune )(_gag bool )){_afe ._dge .RLock ();
defer _afe ._dge .RUnlock ();for _aac ,_gb :=range _afe ._gcd {if f (_aac ,_gb ){break ;};};};func (_dee *StringRuneMap )Length ()int {_dee ._dge .RLock ();defer _dee ._dge .RUnlock ();return len (_dee ._gcd );};func (_dcf *RuneUint16Map )RangeDelete (f func (_ba rune ,_ddd uint16 )(_fae bool ,_abg bool )){_dcf ._bb .Lock ();
defer _dcf ._bb .Unlock ();for _gab ,_gc :=range _dcf ._fa {_bdd ,_abc :=f (_gab ,_gc );if _bdd {delete (_dcf ._fa ,_gab );};if _abc {break ;};};};func (_af *RuneByteMap )Read (r rune )(byte ,bool ){_af ._ec .RLock ();defer _af ._ec .RUnlock ();_ab ,_gec :=_af ._def [r ];
return _ab ,_gec ;};func (_c *ByteRuneMap )Range (f func (_ag byte ,_ce rune )(_da bool )){_c ._d .RLock ();defer _c ._d .RUnlock ();for _ga ,_ge :=range _c ._f {if f (_ga ,_ge ){break ;};};};func MakeRuneSet (length int )*RuneSet {return &RuneSet {_dab :make (map[rune ]struct{},length )}};
func (_ed *RuneStringMap )Read (r rune )(string ,bool ){_ed ._aa .RLock ();defer _ed ._aa .RUnlock ();_bg ,_gf :=_ed ._aga [r ];return _bg ,_gf ;};func MakeRuneByteMap (length int )*RuneByteMap {_ca :=make (map[rune ]byte ,length );return &RuneByteMap {_def :_ca };
};func (_db *RuneUint16Map )Delete (r rune ){_db ._bb .Lock ();defer _db ._bb .Unlock ();delete (_db ._fa ,r )};func (_dfb *RuneUint16Map )Range (f func (_dg rune ,_fd uint16 )(_aea bool )){_dfb ._bb .RLock ();defer _dfb ._bb .RUnlock ();for _cgb ,_fb :=range _dfb ._fa {if f (_cgb ,_fb ){break ;
};};};func (_ef *RuneSet )Range (f func (_agf rune )(_cg bool )){_ef ._ffc .RLock ();defer _ef ._ffc .RUnlock ();for _ece :=range _ef ._dab {if f (_ece ){break ;};};};func (_deff *StringsMap )Range (f func (_cee ,_fab string )(_abd bool )){_deff ._cc .RLock ();
defer _deff ._cc .RUnlock ();for _ccc ,_ceeb :=range _deff ._gbc {if f (_ccc ,_ceeb ){break ;};};};func (_gg *RuneUint16Map )Write (r rune ,g uint16 ){_gg ._bb .Lock ();defer _gg ._bb .Unlock ();_gg ._fa [r ]=g ;};type RuneSet struct{_dab map[rune ]struct{};
_ffc _a .RWMutex ;};