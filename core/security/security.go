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

package security ;import (_eg "bytes";_e "crypto/aes";_d "crypto/cipher";_b "crypto/md5";_fd "crypto/rand";_gca "crypto/rc4";_db "crypto/sha256";_gd "crypto/sha512";_f "encoding/binary";_gb "errors";_bg "fmt";_fdg "github.com/unidoc/unipdf/v3/common";_gc "hash";
_a "io";_da "math";);

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4 (id0 string ,length int )StdHandler {return stdHandlerR4 {ID0 :id0 ,Length :length }};func (_aa stdHandlerR4 )alg2 (_fc *StdEncryptDict ,_bga []byte )[]byte {_fdg .Log .Trace ("\u0061\u006c\u0067\u0032");_fg :=_aa .paddedPass (_bga );
_de :=_b .New ();_de .Write (_fg );_de .Write (_fc .O );var _eff [4]byte ;_f .LittleEndian .PutUint32 (_eff [:],uint32 (_fc .P ));_de .Write (_eff [:]);_fdg .Log .Trace ("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x",_eff );_de .Write ([]byte (_aa .ID0 ));
_fdg .Log .Trace ("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076",_fc .R ,_fc .EncryptMetadata );if (_fc .R >=4)&&!_fc .EncryptMetadata {_de .Write ([]byte {0xff,0xff,0xff,0xff});
};_eag :=_de .Sum (nil );if _fc .R >=3{_de =_b .New ();for _bgc :=0;_bgc < 50;_bgc ++{_de .Reset ();_de .Write (_eag [0:_aa .Length /8]);_eag =_de .Sum (nil );};};if _fc .R >=3{return _eag [0:_aa .Length /8];};return _eag [0:5];};func (_fea stdHandlerR6 )alg8 (_fdeg *StdEncryptDict ,_ge []byte ,_bde []byte )error {if _bgf :=_af ("\u0061\u006c\u0067\u0038","\u004b\u0065\u0079",32,_ge );
_bgf !=nil {return _bgf ;};var _gad [16]byte ;if _ ,_agd :=_a .ReadFull (_fd .Reader ,_gad [:]);_agd !=nil {return _agd ;};_cca :=_gad [0:8];_efe :=_gad [8:16];_becb :=make ([]byte ,len (_bde )+len (_cca ));_dcc :=copy (_becb ,_bde );copy (_becb [_dcc :],_cca );
_fgba ,_fggb :=_fea .alg2b (_fdeg .R ,_becb ,_bde ,nil );if _fggb !=nil {return _fggb ;};U :=make ([]byte ,len (_fgba )+len (_cca )+len (_efe ));_dcc =copy (U ,_fgba [:32]);_dcc +=copy (U [_dcc :],_cca );copy (U [_dcc :],_efe );_fdeg .U =U ;_dcc =len (_bde );
copy (_becb [_dcc :],_efe );_fgba ,_fggb =_fea .alg2b (_fdeg .R ,_becb ,_bde ,nil );if _fggb !=nil {return _fggb ;};_cgb ,_fggb :=_dg (_fgba [:32]);if _fggb !=nil {return _fggb ;};_agda :=make ([]byte ,_e .BlockSize );_ade :=_d .NewCBCEncrypter (_cgb ,_agda );
UE :=make ([]byte ,32);_ade .CryptBlocks (UE ,_ge [:32]);_fdeg .UE =UE ;return nil ;};func (_faa *ecbDecrypter )CryptBlocks (dst ,src []byte ){if len (src )%_faa ._fdga !=0{_fdg .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_fdg .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_faa ._ff .Decrypt (dst ,src [:_faa ._fdga ]);src =src [_faa ._fdga :];dst =dst [_faa ._fdga :];};};

// AuthEvent is an event type that triggers authentication.
type AuthEvent string ;func _af (_ea ,_gdb string ,_bc int ,_gbg []byte )error {if len (_gbg )< _bc {return errInvalidField {Func :_ea ,Field :_gdb ,Exp :_bc ,Got :len (_gbg )};};return nil ;};func (stdHandlerR4 )paddedPass (_ba []byte )[]byte {_agb :=make ([]byte ,32);
_eae :=copy (_agb ,_ba );for ;_eae < 32;_eae ++{_agb [_eae ]=_eac [_eae -len (_ba )];};return _agb ;};func (_gbf stdHandlerR4 )alg3Key (R int ,_bcb []byte )[]byte {_bd :=_b .New ();_bf :=_gbf .paddedPass (_bcb );_bd .Write (_bf );if R >=3{for _efd :=0;
_efd < 50;_efd ++{_dcf :=_bd .Sum (nil );_bd =_b .New ();_bd .Write (_dcf );};};_eaf :=_bd .Sum (nil );if R ==2{_eaf =_eaf [0:5];}else {_eaf =_eaf [0:_gbf .Length /8];};return _eaf ;};func (_ed *ecbEncrypter )CryptBlocks (dst ,src []byte ){if len (src )%_ed ._fdga !=0{_fdg .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_fdg .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_ed ._ff .Encrypt (dst ,src [:_ed ._fdga ]);src =src [_ed ._fdga :];dst =dst [_ed ._fdga :];};};func (_dda stdHandlerR6 )alg13 (_fead *StdEncryptDict ,_bfb []byte )error {if _dff :=_af ("\u0061\u006c\u00671\u0033","\u004b\u0065\u0079",32,_bfb );
_dff !=nil {return _dff ;};if _eba :=_af ("\u0061\u006c\u00671\u0033","\u0050\u0065\u0072m\u0073",16,_fead .Perms );_eba !=nil {return _eba ;};_bad :=make ([]byte ,16);copy (_bad ,_fead .Perms [:16]);_cbdg ,_dgd :=_e .NewCipher (_bfb [:32]);if _dgd !=nil {return _dgd ;
};_ec :=_bb (_cbdg );_ec .CryptBlocks (_bad ,_bad );if !_eg .Equal (_bad [9:12],[]byte ("\u0061\u0064\u0062")){return _gb .New ("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};_bff :=Permissions (_f .LittleEndian .Uint32 (_bad [0:4]));if _bff !=_fead .P {return _gb .New ("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064");
};var _afe bool ;if _bad [8]=='T'{_afe =true ;}else if _bad [8]=='F'{_afe =false ;}else {return _gb .New ("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};if _afe !=_fead .EncryptMetadata {return _gb .New ("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064");
};return nil ;};const (EventDocOpen =AuthEvent ("\u0044o\u0063\u004f\u0070\u0065\u006e");EventEFOpen =AuthEvent ("\u0045\u0046\u004f\u0070\u0065\u006e"););type errInvalidField struct{Func string ;Field string ;Exp int ;Got int ;};type ecb struct{_ff _d .Block ;
_fdga int ;};func (_ggdg stdHandlerR6 )alg2a (_fdc *StdEncryptDict ,_gcg []byte )([]byte ,Permissions ,error ){if _bgad :=_af ("\u0061\u006c\u00672\u0061","\u004f",48,_fdc .O );_bgad !=nil {return nil ,0,_bgad ;};if _dfc :=_af ("\u0061\u006c\u00672\u0061","\u0055",48,_fdc .U );
_dfc !=nil {return nil ,0,_dfc ;};if len (_gcg )> 127{_gcg =_gcg [:127];};_daad ,_eeg :=_ggdg .alg12 (_fdc ,_gcg );if _eeg !=nil {return nil ,0,_eeg ;};var (_cdg []byte ;_edb []byte ;_cba []byte ;);var _deac Permissions ;if len (_daad )!=0{_deac =PermOwner ;
_edc :=make ([]byte ,len (_gcg )+8+48);_ffc :=copy (_edc ,_gcg );_ffc +=copy (_edc [_ffc :],_fdc .O [40:48]);copy (_edc [_ffc :],_fdc .U [0:48]);_cdg =_edc ;_edb =_fdc .OE ;_cba =_fdc .U [0:48];}else {_daad ,_eeg =_ggdg .alg11 (_fdc ,_gcg );if _eeg ==nil &&len (_daad )==0{_daad ,_eeg =_ggdg .alg11 (_fdc ,[]byte (""));
};if _eeg !=nil {return nil ,0,_eeg ;}else if len (_daad )==0{return nil ,0,nil ;};_deac =_fdc .P ;_ffe :=make ([]byte ,len (_gcg )+8);_bgcb :=copy (_ffe ,_gcg );copy (_ffe [_bgcb :],_fdc .U [40:48]);_cdg =_ffe ;_edb =_fdc .UE ;_cba =nil ;};if _fccf :=_af ("\u0061\u006c\u00672\u0061","\u004b\u0065\u0079",32,_edb );
_fccf !=nil {return nil ,0,_fccf ;};_edb =_edb [:32];_dag ,_eeg :=_ggdg .alg2b (_fdc .R ,_cdg ,_gcg ,_cba );if _eeg !=nil {return nil ,0,_eeg ;};_gae ,_eeg :=_e .NewCipher (_dag [:32]);if _eeg !=nil {return nil ,0,_eeg ;};_efcd :=make ([]byte ,_e .BlockSize );
_bdc :=_d .NewCBCDecrypter (_gae ,_efcd );_ced :=make ([]byte ,32);_bdc .CryptBlocks (_ced ,_edb );if _fdc .R ==5{return _ced ,_deac ,nil ;};_eeg =_ggdg .alg13 (_fdc ,_ced );if _eeg !=nil {return nil ,0,_eeg ;};return _ced ,_deac ,nil ;};func _bec (_dfb []byte ,_effc int ){_beg :=_effc ;
for _beg < len (_dfb ){copy (_dfb [_beg :],_dfb [:_beg ]);_beg *=2;};};func (_dgb stdHandlerR6 )alg2b (R int ,_cdgf ,_fee ,_ada []byte )([]byte ,error ){if R ==5{return _eea (_cdgf );};return _dgc (_cdgf ,_fee ,_ada );};func (_gdd stdHandlerR6 )alg11 (_aaae *StdEncryptDict ,_dbc []byte )([]byte ,error ){if _fcca :=_af ("\u0061\u006c\u00671\u0031","\u0055",48,_aaae .U );
_fcca !=nil {return nil ,_fcca ;};_cdc :=make ([]byte ,len (_dbc )+8);_aec :=copy (_cdc ,_dbc );_aec +=copy (_cdc [_aec :],_aaae .U [32:40]);_ffg ,_gf :=_gdd .alg2b (_aaae .R ,_cdc ,_dbc ,nil );if _gf !=nil {return nil ,_gf ;};_ffg =_ffg [:32];if !_eg .Equal (_ffg ,_aaae .U [:32]){return nil ,nil ;
};return _ffg ,nil ;};func _fa (_bgb _d .Block )*ecb {return &ecb {_ff :_bgb ,_fdga :_bgb .BlockSize ()}};func (_ab stdHandlerR4 )alg4 (_eb []byte ,_feb []byte )([]byte ,error ){_dead ,_agf :=_gca .NewCipher (_eb );if _agf !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_abg :=[]byte (_eac );_fgb :=make ([]byte ,len (_abg ));_dead .XORKeyStream (_fgb ,_abg );return _fgb ,nil ;};type ecbDecrypter ecb ;type ecbEncrypter ecb ;func _c (_dc _d .Block )_d .BlockMode {return (*ecbEncrypter )(_fa (_dc ))};

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32 ;

// Allowed checks if a set of permissions can be granted.
func (_cb Permissions )Allowed (p2 Permissions )bool {return _cb &p2 ==p2 };

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6 ()StdHandler {return stdHandlerR6 {}};func (_edf errInvalidField )Error ()string {return _bg .Sprintf ("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064",_edf .Func ,_edf .Field ,_edf .Exp ,_edf .Got );
};func _dgc (_afb ,_gdf ,_fcf []byte )([]byte ,error ){var (_ccf ,_ebg ,_abe _gc .Hash ;);_ccf =_db .New ();_dcfd :=make ([]byte ,64);_fgc :=_ccf ;_fgc .Write (_afb );K :=_fgc .Sum (_dcfd [:0]);_aed :=make ([]byte ,64*(127+64+48));_dd :=func (_cdda int )([]byte ,error ){_bebf :=len (_gdf )+len (K )+len (_fcf );
_bce :=_aed [:_bebf ];_faf :=copy (_bce ,_gdf );_faf +=copy (_bce [_faf :],K [:]);_faf +=copy (_bce [_faf :],_fcf );if _faf !=_bebf {_fdg .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e");
return nil ,_gb .New ("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065");};K1 :=_aed [:_bebf *64];_bec (K1 ,_bebf );_fdcb ,_bdbc :=_dg (K [0:16]);if _bdbc !=nil {return nil ,_bdbc ;};_cfb :=_d .NewCBCEncrypter (_fdcb ,K [16:32]);_cfb .CryptBlocks (K1 ,K1 );
E :=K1 ;_cbdf :=0;for _fde :=0;_fde < 16;_fde ++{_cbdf +=int (E [_fde ]%3);};var _ead _gc .Hash ;switch _cbdf %3{case 0:_ead =_ccf ;case 1:if _ebg ==nil {_ebg =_gd .New384 ();};_ead =_ebg ;case 2:if _abe ==nil {_abe =_gd .New ();};_ead =_abe ;};_ead .Reset ();
_ead .Write (E );K =_ead .Sum (_dcfd [:0]);return E ,nil ;};for _cee :=0;;{E ,_cdf :=_dd (_cee );if _cdf !=nil {return nil ,_cdf ;};_edcb :=E [len (E )-1];_cee ++;if _cee >=64&&_edcb <=uint8 (_cee -32){break ;};};return K [:32],nil ;};

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct{R int ;P Permissions ;EncryptMetadata bool ;O ,U []byte ;OE ,UE []byte ;Perms []byte ;};func (_efc stdHandlerR4 )alg3 (R int ,_fe ,_cd []byte )([]byte ,error ){var _ee []byte ;if len (_cd )> 0{_ee =_efc .alg3Key (R ,_cd );}else {_ee =_efc .alg3Key (R ,_fe );
};_dcb ,_daf :=_gca .NewCipher (_ee );if _daf !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_dec :=_efc .paddedPass (_fe );_bef :=make ([]byte ,len (_dec ));_dcb .XORKeyStream (_bef ,_dec );
if R >=3{_deg :=make ([]byte ,len (_ee ));for _cbe :=0;_cbe < 19;_cbe ++{for _cf :=0;_cf < len (_ee );_cf ++{_deg [_cf ]=_ee [_cf ]^byte (_cbe +1);};_ca ,_dea :=_gca .NewCipher (_deg );if _dea !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_ca .XORKeyStream (_bef ,_bef );};};return _bef ,nil ;};var _ StdHandler =stdHandlerR6 {};func (_bca stdHandlerR6 )alg9 (_cgd *StdEncryptDict ,_gbb []byte ,_eaa []byte )error {if _eagc :=_af ("\u0061\u006c\u0067\u0039","\u004b\u0065\u0079",32,_gbb );
_eagc !=nil {return _eagc ;};if _dcba :=_af ("\u0061\u006c\u0067\u0039","\u0055",48,_cgd .U );_dcba !=nil {return _dcba ;};var _decg [16]byte ;if _ ,_eeb :=_a .ReadFull (_fd .Reader ,_decg [:]);_eeb !=nil {return _eeb ;};_efg :=_decg [0:8];_bbb :=_decg [8:16];
_eda :=_cgd .U [:48];_fgf :=make ([]byte ,len (_eaa )+len (_efg )+len (_eda ));_fge :=copy (_fgf ,_eaa );_fge +=copy (_fgf [_fge :],_efg );_fge +=copy (_fgf [_fge :],_eda );_cfg ,_gbbf :=_bca .alg2b (_cgd .R ,_fgf ,_eaa ,_eda );if _gbbf !=nil {return _gbbf ;
};O :=make ([]byte ,len (_cfg )+len (_efg )+len (_bbb ));_fge =copy (O ,_cfg [:32]);_fge +=copy (O [_fge :],_efg );_fge +=copy (O [_fge :],_bbb );_cgd .O =O ;_fge =len (_eaa );_fge +=copy (_fgf [_fge :],_bbb );_cfg ,_gbbf =_bca .alg2b (_cgd .R ,_fgf ,_eaa ,_eda );
if _gbbf !=nil {return _gbbf ;};_ffaf ,_gbbf :=_dg (_cfg [:32]);if _gbbf !=nil {return _gbbf ;};_eefg :=make ([]byte ,_e .BlockSize );_fdcg :=_d .NewCBCEncrypter (_ffaf ,_eefg );OE :=make ([]byte ,32);_fdcg .CryptBlocks (OE ,_gbb [:32]);_cgd .OE =OE ;return nil ;
};

// Authenticate implements StdHandler interface.
func (_ffac stdHandlerR6 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){return _ffac .alg2a (d ,pass );};const (PermOwner =Permissions (_da .MaxUint32 );PermPrinting =Permissions (1<<2);PermModify =Permissions (1<<3);PermExtractGraphics =Permissions (1<<4);
PermAnnotate =Permissions (1<<5);PermFillForms =Permissions (1<<8);PermDisabilityExtract =Permissions (1<<9);PermRotateInsert =Permissions (1<<10);PermFullPrintQuality =Permissions (1<<11););type stdHandlerR6 struct{};func _dg (_befa []byte )(_d .Block ,error ){_bcfd ,_ebb :=_e .NewCipher (_befa );
if _ebb !=nil {_fdg .Log .Error ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076",_ebb );
return nil ,_ebb ;};return _bcfd ,nil ;};

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_faab stdHandlerR6 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){_abgg :=make ([]byte ,32);if _ ,_aefa :=_a .ReadFull (_fd .Reader ,_abgg );_aefa !=nil {return nil ,_aefa ;};d .U =nil ;d .O =nil ;d .UE =nil ;d .OE =nil ;
d .Perms =nil ;if len (upass )> 127{upass =upass [:127];};if len (opass )> 127{opass =opass [:127];};if _adc :=_faab .alg8 (d ,_abgg ,upass );_adc !=nil {return nil ,_adc ;};if _cecg :=_faab .alg9 (d ,_abgg ,opass );_cecg !=nil {return nil ,_cecg ;};if d .R ==5{return _abgg ,nil ;
};if _agbc :=_faab .alg10 (d ,_abgg );_agbc !=nil {return nil ,_agbc ;};return _abgg ,nil ;};func (_cea stdHandlerR4 )alg6 (_bda *StdEncryptDict ,_cbd []byte )([]byte ,error ){var (_fdb []byte ;_bdd error ;);_fb :=_cea .alg2 (_bda ,_cbd );if _bda .R ==2{_fdb ,_bdd =_cea .alg4 (_fb ,_cbd );
}else if _bda .R >=3{_fdb ,_bdd =_cea .alg5 (_fb ,_cbd );}else {return nil ,_gb .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};if _bdd !=nil {return nil ,_bdd ;};_fdg .Log .Trace ("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f",string (_fdb ),string (_bda .U ));
_fgg :=_fdb ;_df :=_bda .U ;if _bda .R >=3{if len (_fgg )> 16{_fgg =_fgg [0:16];};if len (_df )> 16{_df =_df [0:16];};};if !_eg .Equal (_fgg ,_df ){return nil ,nil ;};return _fb ,nil ;};var _ StdHandler =stdHandlerR4 {};func (_ef *ecbEncrypter )BlockSize ()int {return _ef ._fdga };
type stdHandlerR4 struct{Length int ;ID0 string ;};func (_fcc stdHandlerR4 )alg5 (_ege []byte ,_eef []byte )([]byte ,error ){_fgd :=_b .New ();_fgd .Write ([]byte (_eac ));_fgd .Write ([]byte (_fcc .ID0 ));_cdd :=_fgd .Sum (nil );_fdg .Log .Trace ("\u0061\u006c\u0067\u0035");
_fdg .Log .Trace ("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x",_ege );_fdg .Log .Trace ("\u0049D\u003a\u0020\u0025\u0020\u0078",_fcc .ID0 );if len (_cdd )!=16{return nil ,_gb .New ("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073");
};_cab ,_fca :=_gca .NewCipher (_ege );if _fca !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_bgca :=make ([]byte ,16);_cab .XORKeyStream (_bgca ,_cdd );_edfa :=make ([]byte ,len (_ege ));
for _eafb :=0;_eafb < 19;_eafb ++{for _egc :=0;_egc < len (_ege );_egc ++{_edfa [_egc ]=_ege [_egc ]^byte (_eafb +1);};_cab ,_fca =_gca .NewCipher (_edfa );if _fca !=nil {return nil ,_gb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_cab .XORKeyStream (_bgca ,_bgca );_fdg .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078",_eafb ,_edfa );_fdg .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078",_eafb ,_bgca );
};_ga :=make ([]byte ,32);for _aaa :=0;_aaa < 16;_aaa ++{_ga [_aaa ]=_bgca [_aaa ];};_ ,_fca =_fd .Read (_ga [16:32]);if _fca !=nil {return nil ,_gb .New ("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r");
};return _ga ,nil ;};const _eac ="\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377"+"\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014"+"\251\xfe\x64\x53\x69\172";

// StdHandler is an interface for standard security handlers.
type StdHandler interface{

// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
// It assumes that R, P and EncryptMetadata are already set.
GenerateParams (_be *StdEncryptDict ,_bee ,_gda []byte )([]byte ,error );

// Authenticate uses encryption dictionary parameters and the password to calculate
// the document encryption key. It also returns permissions that should be granted to a user.
// In case of failed authentication, it returns empty key and zero permissions with no error.
Authenticate (_ae *StdEncryptDict ,_cc []byte )([]byte ,Permissions ,error );};func (_dbed stdHandlerR6 )alg12 (_fda *StdEncryptDict ,_fgbd []byte )([]byte ,error ){if _gef :=_af ("\u0061\u006c\u00671\u0032","\u0055",48,_fda .U );_gef !=nil {return nil ,_gef ;
};if _cbb :=_af ("\u0061\u006c\u00671\u0032","\u004f",48,_fda .O );_cbb !=nil {return nil ,_cbb ;};_dccf :=make ([]byte ,len (_fgbd )+8+48);_efa :=copy (_dccf ,_fgbd );_efa +=copy (_dccf [_efa :],_fda .O [32:40]);_efa +=copy (_dccf [_efa :],_fda .U [0:48]);
_ccc ,_adg :=_dbed .alg2b (_fda .R ,_dccf ,_fgbd ,_fda .U [0:48]);if _adg !=nil {return nil ,_adg ;};_ccc =_ccc [:32];if !_eg .Equal (_ccc ,_fda .O [:32]){return nil ,nil ;};return _ccc ,nil ;};

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_gdg stdHandlerR4 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){O ,_cec :=_gdg .alg3 (d .R ,upass ,opass );if _cec !=nil {_fdg .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_cec );
return nil ,_cec ;};d .O =O ;_fdg .Log .Trace ("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078",O );_bbg :=_gdg .alg2 (d ,upass );U ,_cec :=_gdg .alg5 (_bbg ,upass );if _cec !=nil {_fdg .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_cec );
return nil ,_cec ;};d .U =U ;_fdg .Log .Trace ("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078",U );return _bbg ,nil ;};func _eea (_dae []byte )([]byte ,error ){_bdb :=_db .New ();_bdb .Write (_dae );return _bdb .Sum (nil ),nil ;};func (_ag *ecbDecrypter )BlockSize ()int {return _ag ._fdga };
func (_eeag stdHandlerR6 )alg10 (_cda *StdEncryptDict ,_bdg []byte )error {if _efcb :=_af ("\u0061\u006c\u00671\u0030","\u004b\u0065\u0079",32,_bdg );_efcb !=nil {return _efcb ;};_cbf :=uint64 (uint32 (_cda .P ))|(_da .MaxUint32 <<32);Perms :=make ([]byte ,16);
_f .LittleEndian .PutUint64 (Perms [:8],_cbf );if _cda .EncryptMetadata {Perms [8]='T';}else {Perms [8]='F';};copy (Perms [9:12],"\u0061\u0064\u0062");if _ ,_bcfb :=_a .ReadFull (_fd .Reader ,Perms [12:16]);_bcfb !=nil {return _bcfb ;};_edea ,_ddd :=_dg (_bdg [:32]);
if _ddd !=nil {return _ddd ;};_gga :=_c (_edea );_gga .CryptBlocks (Perms ,Perms );_cda .Perms =Perms [:16];return nil ;};func (_beb stdHandlerR4 )alg7 (_gcc *StdEncryptDict ,_gg []byte )([]byte ,error ){_bcc :=_beb .alg3Key (_gcc .R ,_gg );_ac :=make ([]byte ,len (_gcc .O ));
if _gcc .R ==2{_aac ,_fbb :=_gca .NewCipher (_bcc );if _fbb !=nil {return nil ,_gb .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");};_aac .XORKeyStream (_ac ,_gcc .O );}else if _gcc .R >=3{_bcf :=append ([]byte {},_gcc .O ...);
for _gccb :=0;_gccb < 20;_gccb ++{_ede :=append ([]byte {},_bcc ...);for _cg :=0;_cg < len (_bcc );_cg ++{_ede [_cg ]^=byte (19-_gccb );};_ggd ,_ffb :=_gca .NewCipher (_ede );if _ffb !=nil {return nil ,_gb .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");
};_ggd .XORKeyStream (_ac ,_bcf );_bcf =append ([]byte {},_ac ...);};}else {return nil ,_gb .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};_abd ,_aef :=_beb .alg6 (_gcc ,_ac );if _aef !=nil {return nil ,nil ;};return _abd ,nil ;};

// Authenticate implements StdHandler interface.
func (_dbe stdHandlerR4 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){_fdg .Log .Trace ("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073");
_aefc ,_daa :=_dbe .alg7 (d ,pass );if _daa !=nil {return nil ,0,_daa ;};if _aefc !=nil {_fdg .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _aefc ,PermOwner ,nil ;
};_fdg .Log .Trace ("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073");_aefc ,_daa =_dbe .alg6 (d ,pass );if _daa !=nil {return nil ,0,_daa ;
};if _aefc !=nil {_fdg .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _aefc ,d .P ,nil ;};return nil ,0,nil ;};func _bb (_ce _d .Block )_d .BlockMode {return (*ecbDecrypter )(_fa (_ce ))};
