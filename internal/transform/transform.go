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

package transform ;import (_cb "fmt";_d "github.com/unidoc/unipdf/v3/common";_e "math";);func NewMatrix (a ,b ,c ,d ,tx ,ty float64 )Matrix {_g :=Matrix {a ,b ,0,c ,d ,0,tx ,ty ,1};_g .clampRange ();return _g ;};type Point struct{X float64 ;Y float64 ;
};func (_da Matrix )Identity ()bool {return _da [0]==1&&_da [1]==0&&_da [2]==0&&_da [3]==0&&_da [4]==1&&_da [5]==0&&_da [6]==0&&_da [7]==0&&_da [8]==1;};const _gag =1e-6;func (_cg Point )Distance (b Point )float64 {return _e .Hypot (_cg .X -b .X ,_cg .Y -b .Y )};
func (_cad Point )String ()string {return _cb .Sprintf ("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029",_cad .X ,_cad .Y );};func (_bae Matrix )Angle ()float64 {_dbb :=_e .Atan2 (-_bae [1],_bae [0]);if _dbb < 0.0{_dbb +=2*_e .Pi ;};return _dbb /_e .Pi *180.0;
};func (_gf *Matrix )Shear (x ,y float64 ){_gf .Concat (ShearMatrix (x ,y ))};func (_fgb Matrix )Rotate (theta float64 )Matrix {return _fgb .Mult (RotationMatrix (theta ))};func (_ee Matrix )Translation ()(float64 ,float64 ){return _ee [6],_ee [7]};func (_ef Matrix )Scale (xScale ,yScale float64 )Matrix {return _ef .Mult (ScaleMatrix (xScale ,yScale ))};
func (_bac Matrix )Translate (tx ,ty float64 )Matrix {return _bac .Mult (TranslationMatrix (tx ,ty ))};func (_agf *Matrix )clampRange (){for _gbc ,_dd :=range _agf {if _dd > _bgc {_d .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_dd ,_bgc );
_agf [_gbc ]=_bgc ;}else if _dd < -_bgc {_d .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_dd ,-_bgc );_agf [_gbc ]=-_bgc ;};};};func ShearMatrix (x ,y float64 )Matrix {return NewMatrix (1,y ,x ,1,0,0)};
func (_agg *Point )Set (x ,y float64 ){_agg .X ,_agg .Y =x ,y };func (_gfd *Matrix )Clone ()Matrix {return NewMatrix (_gfd [0],_gfd [1],_gfd [3],_gfd [4],_gfd [6],_gfd [7]);};type Matrix [9]float64 ;func (_fd Matrix )Singular ()bool {return _e .Abs (_fd [0]*_fd [4]-_fd [1]*_fd [3])< _cdg };
const _cdg =1e-10;func (_bgd *Point )transformByMatrix (_eab Matrix ){_bgd .X ,_bgd .Y =_eab .Transform (_bgd .X ,_bgd .Y )};func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};func (_gb Matrix )Mult (b Matrix )Matrix {_gb .Concat (b );return _gb };
const _bgc =1e9;func (_bg Matrix )Transform (x ,y float64 )(float64 ,float64 ){_fc :=x *_bg [0]+y *_bg [3]+_bg [6];_dac :=x *_bg [1]+y *_bg [4]+_bg [7];return _fc ,_dac ;};func (_dccf Point )Displace (delta Point )Point {return Point {_dccf .X +delta .X ,_dccf .Y +delta .Y }};
func RotationMatrix (angle float64 )Matrix {_ea :=_e .Cos (angle );_cd :=_e .Sin (angle );return NewMatrix (_ea ,_cd ,-_cd ,_ea ,0,0);};const _gg =1.0e-6;func (_fe Matrix )Unrealistic ()bool {_bdg ,_aa ,_fb ,_cba :=_e .Abs (_fe [0]),_e .Abs (_fe [1]),_e .Abs (_fe [3]),_e .Abs (_fe [4]);
_dcc :=_bdg > _gag &&_cba > _gag ;_gcg :=_aa > _gag &&_fb > _gag ;return !(_dcc ||_gcg );};func NewMatrixFromTransforms (xScale ,yScale ,theta ,tx ,ty float64 )Matrix {return IdentityMatrix ().Scale (xScale ,yScale ).Rotate (theta ).Translate (tx ,ty );
};func (_eef *Point )Transform (a ,b ,c ,d ,tx ,ty float64 ){_ebg :=NewMatrix (a ,b ,c ,d ,tx ,ty );_eef .transformByMatrix (_ebg );};func (_db Matrix )ScalingFactorX ()float64 {return _e .Hypot (_db [0],_db [1])};func (_ab *Matrix )Set (a ,b ,c ,d ,tx ,ty float64 ){_ab [0],_ab [1]=a ,b ;
_ab [3],_ab [4]=c ,d ;_ab [6],_ab [7]=tx ,ty ;_ab .clampRange ();};func TranslationMatrix (tx ,ty float64 )Matrix {return NewMatrix (1,0,0,1,tx ,ty )};func (_de Point )Rotate (theta float64 )Point {_ce :=_e .Hypot (_de .X ,_de .Y );_ffg :=_e .Atan2 (_de .Y ,_de .X );
_gbe ,_beb :=_e .Sincos (_ffg +theta /180.0*_e .Pi );return Point {_ce *_beb ,_ce *_gbe };};func (_ed Point )Interpolate (b Point ,t float64 )Point {return Point {X :(1-t )*_ed .X +t *b .X ,Y :(1-t )*_ed .Y +t *b .Y };};func (_f Matrix )Round (precision float64 )Matrix {for _fg :=range _f {_f [_fg ]=_e .Round (_f [_fg ]/precision )*precision ;
};return _f ;};func IdentityMatrix ()Matrix {return NewMatrix (1,0,0,1,0,0)};func (_ff Matrix )ScalingFactorY ()float64 {return _e .Hypot (_ff [3],_ff [4])};func (_ac Matrix )Inverse ()(Matrix ,bool ){_ag ,_dbc :=_ac [0],_ac [1];_be ,_ga :=_ac [3],_ac [4];
_dbe ,_cdb :=_ac [6],_ac [7];_gc :=_ag *_ga -_dbc *_be ;if _e .Abs (_gc )< _gg {return Matrix {},false ;};_bd ,_eec :=_ga /_gc ,-_dbc /_gc ;_ec ,_bc :=-_be /_gc ,_ag /_gc ;_df :=-(_bd *_dbe +_ec *_cdb );_eag :=-(_eec *_dbe +_bc *_cdb );return NewMatrix (_bd ,_eec ,_ec ,_bc ,_df ,_eag ),true ;
};func ScaleMatrix (x ,y float64 )Matrix {return NewMatrix (x ,0,0,y ,0,0)};func (_a Matrix )String ()string {_eg ,_ca ,_cbe ,_b ,_eb ,_dc :=_a [0],_a [1],_a [3],_a [4],_a [6],_a [7];return _cb .Sprintf ("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d",_eg ,_ca ,_cbe ,_b ,_eb ,_dc );
};func (_ba *Matrix )Concat (b Matrix ){*_ba =Matrix {b [0]*_ba [0]+b [1]*_ba [3],b [0]*_ba [1]+b [1]*_ba [4],0,b [3]*_ba [0]+b [4]*_ba [3],b [3]*_ba [1]+b [4]*_ba [4],0,b [6]*_ba [0]+b [7]*_ba [3]+_ba [6],b [6]*_ba [1]+b [7]*_ba [4]+_ba [7],1};_ba .clampRange ();
};