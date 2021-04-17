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

package transform ;import (_f "fmt";_af "github.com/unidoc/unipdf/v3/common";_d "math";);func (_aeg *Matrix )Shear (x ,y float64 ){_aeg .Concat (ShearMatrix (x ,y ))};func (_ea *Matrix )Concat (b Matrix ){*_ea =Matrix {b [0]*_ea [0]+b [1]*_ea [3],b [0]*_ea [1]+b [1]*_ea [4],0,b [3]*_ea [0]+b [4]*_ea [3],b [3]*_ea [1]+b [4]*_ea [4],0,b [6]*_ea [0]+b [7]*_ea [3]+_ea [6],b [6]*_ea [1]+b [7]*_ea [4]+_ea [7],1};
_ea .clampRange ();};func (_bf Matrix )Translate (tx ,ty float64 )Matrix {return _bf .Mult (TranslationMatrix (tx ,ty ))};func TranslationMatrix (tx ,ty float64 )Matrix {return NewMatrix (1,0,0,1,tx ,ty )};func (_fa Point )Distance (b Point )float64 {return _d .Hypot (_fa .X -b .X ,_fa .Y -b .Y )};
func (_gfg Point )String ()string {return _f .Sprintf ("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029",_gfg .X ,_gfg .Y );};func (_db *Matrix )Set (a ,b ,c ,d ,tx ,ty float64 ){_db [0],_db [1]=a ,b ;_db [3],_db [4]=c ,d ;_db [6],_db [7]=tx ,ty ;
_db .clampRange ();};func (_fgf Matrix )Transform (x ,y float64 )(float64 ,float64 ){_dfa :=x *_fgf [0]+y *_fgf [3]+_fgf [6];_bfg :=x *_fgf [1]+y *_fgf [4]+_fgf [7];return _dfa ,_bfg ;};func (_bg Matrix )ScalingFactorX ()float64 {return _d .Hypot (_bg [0],_bg [1])};
func (_ac Matrix )Identity ()bool {return _ac [0]==1&&_ac [1]==0&&_ac [2]==0&&_ac [3]==0&&_ac [4]==1&&_ac [5]==0&&_ac [6]==0&&_ac [7]==0&&_ac [8]==1;};func (_ba Matrix )Singular ()bool {return _d .Abs (_ba [0]*_ba [4]-_ba [1]*_ba [3])< _ab };func (_ec *Matrix )Clone ()Matrix {return NewMatrix (_ec [0],_ec [1],_ec [3],_ec [4],_ec [6],_ec [7])};
const _ge =1.0e-6;func (_dff Matrix )ScalingFactorY ()float64 {return _d .Hypot (_dff [3],_dff [4])};func NewMatrix (a ,b ,c ,d ,tx ,ty float64 )Matrix {_g :=Matrix {a ,b ,0,c ,d ,0,tx ,ty ,1};_g .clampRange ();return _g ;};func (_cgd *Point )Transform (a ,b ,c ,d ,tx ,ty float64 ){_dd :=NewMatrix (a ,b ,c ,d ,tx ,ty );
_cgd .transformByMatrix (_dd );};const _ab =1e-10;func (_dc Matrix )String ()string {_fg ,_e ,_aff ,_c ,_ad ,_de :=_dc [0],_dc [1],_dc [3],_dc [4],_dc [6],_dc [7];return _f .Sprintf ("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d",_fg ,_e ,_aff ,_c ,_ad ,_de );
};func (_egd Point )Displace (delta Point )Point {return Point {_egd .X +delta .X ,_egd .Y +delta .Y }};func RotationMatrix (angle float64 )Matrix {_ae :=_d .Cos (angle );_df :=_d .Sin (angle );return NewMatrix (_ae ,_df ,-_df ,_ae ,0,0);};func (_cda *Point )transformByMatrix (_bc Matrix ){_cda .X ,_cda .Y =_bc .Transform (_cda .X ,_cda .Y )};
func (_eca Point )Interpolate (b Point ,t float64 )Point {return Point {X :(1-t )*_eca .X +t *b .X ,Y :(1-t )*_eca .Y +t *b .Y };};const _dbf =1e9;func (_cd Point )Rotate (theta float64 )Point {_bde :=_d .Hypot (_cd .X ,_cd .Y );_feg :=_d .Atan2 (_cd .Y ,_cd .X );
_ecf ,_ggb :=_d .Sincos (_feg +theta /180.0*_d .Pi );return Point {_bde *_ggb ,_bde *_ecf };};type Point struct{X float64 ;Y float64 ;};func (_cc Matrix )Mult (b Matrix )Matrix {_cc .Concat (b );return _cc };func (_afa Matrix )Angle ()float64 {_gb :=_d .Atan2 (-_afa [1],_afa [0]);
if _gb < 0.0{_gb +=2*_d .Pi ;};return _gb /_d .Pi *180.0;};func (_add Matrix )Translation ()(float64 ,float64 ){return _add [6],_add [7]};func (_ee *Point )Set (x ,y float64 ){_ee .X ,_ee .Y =x ,y };func (_fdf *Matrix )clampRange (){for _eaf ,_gfc :=range _fdf {if _gfc > _dbf {_af .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_gfc ,_dbf );
_fdf [_eaf ]=_dbf ;}else if _gfc < -_dbf {_af .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_gfc ,-_dbf );_fdf [_eaf ]=-_dbf ;};};};func (_fdb Matrix )Unrealistic ()bool {_fba ,_be ,_beg ,_ce :=_d .Abs (_fdb [0]),_d .Abs (_fdb [1]),_d .Abs (_fdb [3]),_d .Abs (_fdb [4]);
_ca :=_fba > _cg &&_ce > _cg ;_feb :=_be > _cg &&_beg > _cg ;return !(_ca ||_feb );};func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};func (_b Matrix )Rotate (theta float64 )Matrix {return _b .Mult (RotationMatrix (theta ))};func NewMatrixFromTransforms (xScale ,yScale ,theta ,tx ,ty float64 )Matrix {return IdentityMatrix ().Scale (xScale ,yScale ).Rotate (theta ).Translate (tx ,ty );
};func IdentityMatrix ()Matrix {return NewMatrix (1,0,0,1,0,0)};const _cg =1e-6;func (_fge Matrix )Inverse ()(Matrix ,bool ){_gg ,_fe :=_fge [0],_fge [1];_ecc ,_gf :=_fge [3],_fge [4];_eg ,_gbc :=_fge [6],_fge [7];_gd :=_gg *_gf -_fe *_ecc ;if _d .Abs (_gd )< _ge {return Matrix {},false ;
};_ggc ,_abg :=_gf /_gd ,-_fe /_gd ;_cb ,_aa :=-_ecc /_gd ,_gg /_gd ;_gdf :=-(_ggc *_eg +_cb *_gbc );_bd :=-(_abg *_eg +_aa *_gbc );return NewMatrix (_ggc ,_abg ,_cb ,_aa ,_gdf ,_bd ),true ;};type Matrix [9]float64 ;func (_fb Matrix )Scale (xScale ,yScale float64 )Matrix {return _fb .Mult (ScaleMatrix (xScale ,yScale ))};
func (_fd Matrix )Round (precision float64 )Matrix {for _ff :=range _fd {_fd [_ff ]=_d .Round (_fd [_ff ]/precision )*precision ;};return _fd ;};func ShearMatrix (x ,y float64 )Matrix {return NewMatrix (1,y ,x ,1,0,0)};func ScaleMatrix (x ,y float64 )Matrix {return NewMatrix (x ,0,0,y ,0,0)};
