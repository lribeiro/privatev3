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

package mmr ;import (_c "errors";_d "fmt";_ec "github.com/unidoc/unipdf/v3/common";_b "github.com/unidoc/unipdf/v3/internal/bitwise";_cc "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_g "io";);func (_dda *Decoder )createLittleEndianTable (_gbc [][3]int )([]*code ,error ){_acc :=make ([]*code ,_afc +1);
for _abe :=0;_abe < len (_gbc );_abe ++{_egd :=_gf (_gbc [_abe ]);if _egd ._ea <=_feb {_bge :=_feb -_egd ._ea ;_cac :=_egd ._ee <<uint (_bge );for _dfa :=(1<<uint (_bge ))-1;_dfa >=0;_dfa --{_egg :=_cac |_dfa ;_acc [_egg ]=_egd ;};}else {_dff :=_egd ._ee >>uint (_egd ._ea -_feb );
if _acc [_dff ]==nil {var _geb =_gf ([3]int {});_geb ._cf =make ([]*code ,_bf +1);_acc [_dff ]=_geb ;};if _egd ._ea <=_feb +_cfg {_dc :=_feb +_cfg -_egd ._ea ;_gfg :=(_egd ._ee <<uint (_dc ))&_bf ;_acc [_dff ]._f =true ;for _ccf :=(1<<uint (_dc ))-1;_ccf >=0;
_ccf --{_acc [_dff ]._cf [_gfg |_ccf ]=_egd ;};}else {return nil ,_c .New ("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072");
};};};return _acc ,nil ;};var (_ae =[][3]int {{4,0x1,int (_fb )},{3,0x1,int (_af )},{1,0x1,int (_ab )},{3,0x3,int (_cd )},{6,0x3,int (_gbf )},{7,0x3,int (_ef )},{3,0x2,int (_fa )},{6,0x2,int (_fad )},{7,0x2,int (_fg )},{10,0xf,int (_fc )},{12,0xf,int (_eg )},{12,0x1,int (EOL )}};
_ac =[][3]int {{4,0x07,2},{4,0x08,3},{4,0x0B,4},{4,0x0C,5},{4,0x0E,6},{4,0x0F,7},{5,0x12,128},{5,0x13,8},{5,0x14,9},{5,0x1B,64},{5,0x07,10},{5,0x08,11},{6,0x17,192},{6,0x18,1664},{6,0x2A,16},{6,0x2B,17},{6,0x03,13},{6,0x34,14},{6,0x35,15},{6,0x07,1},{6,0x08,12},{7,0x13,26},{7,0x17,21},{7,0x18,28},{7,0x24,27},{7,0x27,18},{7,0x28,24},{7,0x2B,25},{7,0x03,22},{7,0x37,256},{7,0x04,23},{7,0x08,20},{7,0xC,19},{8,0x12,33},{8,0x13,34},{8,0x14,35},{8,0x15,36},{8,0x16,37},{8,0x17,38},{8,0x1A,31},{8,0x1B,32},{8,0x02,29},{8,0x24,53},{8,0x25,54},{8,0x28,39},{8,0x29,40},{8,0x2A,41},{8,0x2B,42},{8,0x2C,43},{8,0x2D,44},{8,0x03,30},{8,0x32,61},{8,0x33,62},{8,0x34,63},{8,0x35,0},{8,0x36,320},{8,0x37,384},{8,0x04,45},{8,0x4A,59},{8,0x4B,60},{8,0x5,46},{8,0x52,49},{8,0x53,50},{8,0x54,51},{8,0x55,52},{8,0x58,55},{8,0x59,56},{8,0x5A,57},{8,0x5B,58},{8,0x64,448},{8,0x65,512},{8,0x67,640},{8,0x68,576},{8,0x0A,47},{8,0x0B,48},{9,0x01,_cdg },{9,0x98,1472},{9,0x99,1536},{9,0x9A,1600},{9,0x9B,1728},{9,0xCC,704},{9,0xCD,768},{9,0xD2,832},{9,0xD3,896},{9,0xD4,960},{9,0xD5,1024},{9,0xD6,1088},{9,0xD7,1152},{9,0xD8,1216},{9,0xD9,1280},{9,0xDA,1344},{9,0xDB,1408},{10,0x01,_cdg },{11,0x01,_cdg },{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560}};
_gd =[][3]int {{2,0x02,3},{2,0x03,2},{3,0x02,1},{3,0x03,4},{4,0x02,6},{4,0x03,5},{5,0x03,7},{6,0x04,9},{6,0x05,8},{7,0x04,10},{7,0x05,11},{7,0x07,12},{8,0x04,13},{8,0x07,14},{9,0x01,_cdg },{9,0x18,15},{10,0x01,_cdg },{10,0x17,16},{10,0x18,17},{10,0x37,0},{10,0x08,18},{10,0x0F,64},{11,0x01,_cdg },{11,0x17,24},{11,0x18,25},{11,0x28,23},{11,0x37,22},{11,0x67,19},{11,0x68,20},{11,0x6C,21},{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560},{12,0x24,52},{12,0x27,55},{12,0x28,56},{12,0x2B,59},{12,0x2C,60},{12,0x33,320},{12,0x34,384},{12,0x35,448},{12,0x37,53},{12,0x38,54},{12,0x52,50},{12,0x53,51},{12,0x54,44},{12,0x55,45},{12,0x56,46},{12,0x57,47},{12,0x58,57},{12,0x59,58},{12,0x5A,61},{12,0x5B,256},{12,0x64,48},{12,0x65,49},{12,0x66,62},{12,0x67,63},{12,0x68,30},{12,0x69,31},{12,0x6A,32},{12,0x6B,33},{12,0x6C,40},{12,0x6D,41},{12,0xC8,128},{12,0xC9,192},{12,0xCA,26},{12,0xCB,27},{12,0xCC,28},{12,0xCD,29},{12,0xD2,34},{12,0xD3,35},{12,0xD4,36},{12,0xD5,37},{12,0xD6,38},{12,0xD7,39},{12,0xDA,42},{12,0xDB,43},{13,0x4A,640},{13,0x4B,704},{13,0x4C,768},{13,0x4D,832},{13,0x52,1280},{13,0x53,1344},{13,0x54,1408},{13,0x55,1472},{13,0x5A,1536},{13,0x5B,1600},{13,0x64,1664},{13,0x65,1728},{13,0x6C,512},{13,0x6D,576},{13,0x72,896},{13,0x73,960},{13,0x74,1024},{13,0x75,1088},{13,0x76,1152},{13,0x77,1216}};
);func (_afg *Decoder )UncompressMMR ()(_ba *_cc .Bitmap ,_ag error ){_ba =_cc .New (_afg ._fca ,_afg ._ecf );_ged :=make ([]int ,_ba .Width +5);_eb :=make ([]int ,_ba .Width +5);_eb [0]=_ba .Width ;_gcc :=1;var _ecg int ;for _bg :=0;_bg < _ba .Height ;
_bg ++{_ecg ,_ag =_afg .uncompress2d (_afg ._gdg ,_eb ,_gcc ,_ged ,_ba .Width );if _ag !=nil {return nil ,_ag ;};if _ecg ==EOF {break ;};if _ecg > 0{_ag =_afg .fillBitmap (_ba ,_bg ,_ged ,_ecg );if _ag !=nil {return nil ,_ag ;};};_eb ,_ged =_ged ,_eb ;
_gcc =_ecg ;};if _ag =_afg .detectAndSkipEOL ();_ag !=nil {return nil ,_ag ;};_afg ._gdg .align ();return _ba ,nil ;};func _gf (_ca [3]int )*code {return &code {_ea :_ca [0],_ee :_ca [1],_db :_ca [2]}};func _be (_bd *_b .SubstreamReader )(*runData ,error ){_fbc :=&runData {_gee :_bd ,_egga :0,_eggc :1};
_aef :=_gb (_fe (_gfga ,int (_bd .Length ())),_faa );_fbc ._accf =make ([]byte ,_aef );if _aea :=_fbc .fillBuffer (0);_aea !=nil {if _aea ==_g .EOF {_fbc ._accf =make ([]byte ,10);_ec .Log .Debug ("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076",_aea );
}else {return nil ,_aea ;};};return _fbc ,nil ;};type Decoder struct{_fca ,_ecf int ;_gdg *runData ;_geg []*code ;_de []*code ;_fec []*code ;};type code struct{_ea int ;_ee int ;_db int ;_cf []*code ;_f bool ;};const (_fb mmrCode =iota ;_af ;_ab ;_cd ;
_gbf ;_ef ;_fa ;_fad ;_fg ;_fc ;_eg ;);func (_cca *runData )fillBuffer (_gde int )error {_cca ._bae =_gde ;_ ,_dfe :=_cca ._gee .Seek (int64 (_gde ),_g .SeekStart );if _dfe !=nil {if _dfe ==_g .EOF {_ec .Log .Debug ("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046");
_cca ._ad =-1;}else {return _dfe ;};};if _dfe ==nil {_cca ._ad ,_dfe =_cca ._gee .Read (_cca ._accf );if _dfe !=nil {if _dfe ==_g .EOF {_ec .Log .Trace ("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046");_cca ._ad =-1;}else {return _dfe ;};};};if _cca ._ad > -1&&_cca ._ad < 3{for _cca ._ad < 3{_geeg ,_cdgc :=_cca ._gee .ReadByte ();
if _cdgc !=nil {if _cdgc ==_g .EOF {_cca ._accf [_cca ._ad ]=0;}else {return _cdgc ;};}else {_cca ._accf [_cca ._ad ]=_geeg &0xFF;};_cca ._ad ++;};};_cca ._ad -=3;if _cca ._ad < 0{_cca ._accf =make ([]byte ,len (_cca ._accf ));_cca ._ad =len (_cca ._accf )-3;
};return nil ;};func New (r _b .StreamReader ,width ,height int ,dataOffset ,dataLength int64 )(*Decoder ,error ){_afa :=&Decoder {_fca :width ,_ecf :height };_ga ,_gc :=_b .NewSubstreamReader (r ,uint64 (dataOffset ),uint64 (dataLength ));if _gc !=nil {return nil ,_gc ;
};_df ,_gc :=_be (_ga );if _gc !=nil {return nil ,_gc ;};_afa ._gdg =_df ;if _ff :=_afa .initTables ();_ff !=nil {return nil ,_ff ;};return _afa ,nil ;};func (_fef *runData )uncompressGetCodeLittleEndian (_egf []*code )(*code ,error ){_efe ,_eda :=_fef .uncompressGetNextCodeLittleEndian ();
if _eda !=nil {_ec .Log .Debug ("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076",_eda );
return nil ,_eda ;};_efe &=0xffffff;_ebg :=_efe >>(_fccc -_feb );_aec :=_egf [_ebg ];if _aec !=nil &&_aec ._f {_ebg =(_efe >>(_fccc -_feb -_cfg ))&_bf ;_aec =_aec ._cf [_ebg ];};return _aec ,nil ;};func (_gfgg *Decoder )initTables ()(_cgf error ){if _gfgg ._geg ==nil {_gfgg ._geg ,_cgf =_gfgg .createLittleEndianTable (_ac );
if _cgf !=nil {return ;};_gfgg ._de ,_cgf =_gfgg .createLittleEndianTable (_gd );if _cgf !=nil {return ;};_gfgg ._fec ,_cgf =_gfgg .createLittleEndianTable (_ae );if _cgf !=nil {return ;};};return nil ;};func (_edc *runData )align (){_edc ._egga =((_edc ._egga +7)>>3)<<3};
func (_dfc *Decoder )uncompress2d (_cb *runData ,_bag []int ,_ebc int ,_deg []int ,_cae int )(int ,error ){var (_fbd int ;_bfa int ;_fdd int ;_ed =true ;_ddg error ;_edd *code ;);_bag [_ebc ]=_cae ;_bag [_ebc +1]=_cae ;_bag [_ebc +2]=_cae +1;_bag [_ebc +3]=_cae +1;
_ddd :for _fdd < _cae {_edd ,_ddg =_cb .uncompressGetCode (_dfc ._fec );if _ddg !=nil {return EOL ,nil ;};if _edd ==nil {_cb ._egga ++;break _ddd ;};_cb ._egga +=_edd ._ea ;switch mmrCode (_edd ._db ){case _ab :_fdd =_bag [_fbd ];case _cd :_fdd =_bag [_fbd ]+1;
case _fa :_fdd =_bag [_fbd ]-1;case _af :for {var _dbc []*code ;if _ed {_dbc =_dfc ._geg ;}else {_dbc =_dfc ._de ;};_edd ,_ddg =_cb .uncompressGetCode (_dbc );if _ddg !=nil {return 0,_ddg ;};if _edd ==nil {break _ddd ;};_cb ._egga +=_edd ._ea ;if _edd ._db < 64{if _edd ._db < 0{_deg [_bfa ]=_fdd ;
_bfa ++;_edd =nil ;break _ddd ;};_fdd +=_edd ._db ;_deg [_bfa ]=_fdd ;_bfa ++;break ;};_fdd +=_edd ._db ;};_bgf :=_fdd ;_fgg :for {var _cff []*code ;if !_ed {_cff =_dfc ._geg ;}else {_cff =_dfc ._de ;};_edd ,_ddg =_cb .uncompressGetCode (_cff );if _ddg !=nil {return 0,_ddg ;
};if _edd ==nil {break _ddd ;};_cb ._egga +=_edd ._ea ;if _edd ._db < 64{if _edd ._db < 0{_deg [_bfa ]=_fdd ;_bfa ++;break _ddd ;};_fdd +=_edd ._db ;if _fdd < _cae ||_fdd !=_bgf {_deg [_bfa ]=_fdd ;_bfa ++;};break _fgg ;};_fdd +=_edd ._db ;};for _fdd < _cae &&_bag [_fbd ]<=_fdd {_fbd +=2;
};continue _ddd ;case _fb :_fbd ++;_fdd =_bag [_fbd ];_fbd ++;continue _ddd ;case _gbf :_fdd =_bag [_fbd ]+2;case _fad :_fdd =_bag [_fbd ]-2;case _ef :_fdd =_bag [_fbd ]+3;case _fg :_fdd =_bag [_fbd ]-3;default:if _cb ._egga ==12&&_edd ._db ==EOL {_cb ._egga =0;
if _ ,_ddg =_dfc .uncompress1d (_cb ,_bag ,_cae );_ddg !=nil {return 0,_ddg ;};_cb ._egga ++;if _ ,_ddg =_dfc .uncompress1d (_cb ,_deg ,_cae );_ddg !=nil {return 0,_ddg ;};_ceee ,_eddc :=_dfc .uncompress1d (_cb ,_bag ,_cae );if _eddc !=nil {return EOF ,_eddc ;
};_cb ._egga ++;return _ceee ,nil ;};_fdd =_cae ;continue _ddd ;};if _fdd <=_cae {_ed =!_ed ;_deg [_bfa ]=_fdd ;_bfa ++;if _fbd > 0{_fbd --;}else {_fbd ++;};for _fdd < _cae &&_bag [_fbd ]<=_fdd {_fbd +=2;};};};if _deg [_bfa ]!=_cae {_deg [_bfa ]=_cae ;
};if _edd ==nil {return EOL ,nil ;};return _bfa ,nil ;};func (_aa *Decoder )detectAndSkipEOL ()error {for {_cg ,_fbf :=_aa ._gdg .uncompressGetCode (_aa ._fec );if _fbf !=nil {return _fbf ;};if _cg !=nil &&_cg ._db ==EOL {_aa ._gdg ._egga +=_cg ._ea ;}else {return nil ;
};};};func _gb (_fd ,_ge int )int {if _fd > _ge {return _ge ;};return _fd ;};func (_dge *Decoder )uncompress1d (_gaa *runData ,_eeg []int ,_fde int )(int ,error ){var (_dfg =true ;_gbb int ;_fdef *code ;_fcc int ;_ddb error ;);_age :for _gbb < _fde {_abg :for {if _dfg {_fdef ,_ddb =_gaa .uncompressGetCode (_dge ._geg );
if _ddb !=nil {return 0,_ddb ;};}else {_fdef ,_ddb =_gaa .uncompressGetCode (_dge ._de );if _ddb !=nil {return 0,_ddb ;};};_gaa ._egga +=_fdef ._ea ;if _fdef ._db < 0{break _age ;};_gbb +=_fdef ._db ;if _fdef ._db < 64{_dfg =!_dfg ;_eeg [_fcc ]=_gbb ;_fcc ++;
break _abg ;};};};if _eeg [_fcc ]!=_fde {_eeg [_fcc ]=_fde ;};_gbe :=EOL ;if _fdef !=nil &&_fdef ._db !=EOL {_gbe =_fcc ;};return _gbe ,nil ;};const (EOF =-3;_cdg =-2;EOL =-1;_feb =8;_afc =(1<<_feb )-1;_cfg =5;_bf =(1<<_cfg )-1;);type mmrCode int ;const (_faa int =1024<<7;
_gfga int =3;_fccc uint =24;);func (_eaf *runData )uncompressGetCode (_bgd []*code )(*code ,error ){return _eaf .uncompressGetCodeLittleEndian (_bgd );};func (_fea *runData )uncompressGetNextCodeLittleEndian ()(int ,error ){_dgd :=_fea ._egga -_fea ._eggc ;
if _dgd < 0||_dgd > 24{_gef :=(_fea ._egga >>3)-_fea ._bae ;if _gef >=_fea ._ad {_gef +=_fea ._bae ;if _aeab :=_fea .fillBuffer (_gef );_aeab !=nil {return 0,_aeab ;};_gef -=_fea ._bae ;};_gdf :=(uint32 (_fea ._accf [_gef ]&0xFF)<<16)|(uint32 (_fea ._accf [_gef +1]&0xFF)<<8)|(uint32 (_fea ._accf [_gef +2]&0xFF));
_add :=uint32 (_fea ._egga &7);_gdf <<=_add ;_fea ._bfaf =int (_gdf );}else {_beb :=_fea ._eggc &7;_ccb :=7-_beb ;if _dgd <=_ccb {_fea ._bfaf <<=uint (_dgd );}else {_bc :=(_fea ._eggc >>3)+3-_fea ._bae ;if _bc >=_fea ._ad {_bc +=_fea ._bae ;if _gfc :=_fea .fillBuffer (_bc );
_gfc !=nil {return 0,_gfc ;};_bc -=_fea ._bae ;};_beb =8-_beb ;for {_fea ._bfaf <<=uint (_beb );_fea ._bfaf |=int (uint (_fea ._accf [_bc ])&0xFF);_dgd -=_beb ;_bc ++;_beb =8;if !(_dgd >=8){break ;};};_fea ._bfaf <<=uint (_dgd );};};_fea ._eggc =_fea ._egga ;
return _fea ._bfaf ,nil ;};func _fe (_gfd ,_a int )int {if _gfd < _a {return _a ;};return _gfd ;};func (_cga *Decoder )fillBitmap (_fac *_cc .Bitmap ,_cgc int ,_afaa []int ,_dg int )error {var _gg byte ;_ce :=0;_cfd :=_fac .GetByteIndex (_ce ,_cgc );for _eef :=0;
_eef < _dg ;_eef ++{_gcd :=byte (1);_eeb :=_afaa [_eef ];if (_eef &1)==0{_gcd =0;};for _ce < _eeb {_gg =(_gg <<1)|_gcd ;_ce ++;if (_ce &7)==0{if _bad :=_fac .SetByte (_cfd ,_gg );_bad !=nil {return _bad ;};_cfd ++;_gg =0;};};};if (_ce &7)!=0{_gg <<=uint (8-(_ce &7));
if _aed :=_fac .SetByte (_cfd ,_gg );_aed !=nil {return _aed ;};};return nil ;};func (_dd *code )String ()string {return _d .Sprintf ("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_dd ._ea ,_dd ._ee ,_dd ._db );};type runData struct{_gee *_b .SubstreamReader ;
_egga int ;_eggc int ;_bfaf int ;_accf []byte ;_bae int ;_ad int ;};