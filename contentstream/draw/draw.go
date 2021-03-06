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

// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw ;import (_c "fmt";_a "github.com/unidoc/unipdf/v3/contentstream";_fg "github.com/unidoc/unipdf/v3/core";_eb "github.com/unidoc/unipdf/v3/internal/transform";_fb "github.com/unidoc/unipdf/v3/model";_f "math";);

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath ()CubicBezierPath {_df :=CubicBezierPath {};_df .Curves =[]CubicBezierCurve {};return _df ;};

// Offset shifts the Bezier path with the specified offsets.
func (_ge CubicBezierPath )Offset (offX ,offY float64 )CubicBezierPath {for _ag ,_ebf :=range _ge .Curves {_ge .Curves [_ag ]=_ebf .AddOffsetXY (offX ,offY );};return _ge ;};

// Magnitude returns the magnitude of the vector.
func (_fcbc Vector )Magnitude ()float64 {return _f .Sqrt (_f .Pow (_fcbc .Dx ,2.0)+_f .Pow (_fcbc .Dy ,2.0));};

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_aa Path )RemovePoint (number int )Path {if number < 1||number > len (_aa .Points ){return _aa ;};_fd :=number -1;_aa .Points =append (_aa .Points [:_fd ],_aa .Points [_fd +1:]...);return _aa ;};const (LineEndingStyleNone LineEndingStyle =0;LineEndingStyleArrow LineEndingStyle =1;
LineEndingStyleButt LineEndingStyle =2;);

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector (dx ,dy float64 )Vector {_dbc :=Vector {};_dbc .Dx =dx ;_dbc .Dy =dy ;return _dbc };

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;};

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct{Curves []CubicBezierCurve ;BorderWidth float64 ;BorderColor _fb .PdfColor ;FillEnabled bool ;FillColor _fb .PdfColor ;};

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{Points []Point ;};

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor _fb .PdfColor ;BorderEnabled bool ;BorderWidth float64 ;BorderColor _fb .PdfColor ;Opacity float64 ;};

// FlipX flips the sign of the Dx component of the vector.
func (_bda Vector )FlipX ()Vector {_bda .Dx =-_bda .Dx ;return _bda };

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_ebg Point )Rotate (theta float64 )Point {_fgd :=_eb .NewPoint (_ebg .X ,_ebg .Y ).Rotate (theta );return NewPoint (_fgd .X ,_fgd .Y );};

// Offset shifts the path with the specified offsets.
func (_fbc Path )Offset (offX ,offY float64 )Path {for _cbc ,_cga :=range _fbc .Points {_fbc .Points [_cbc ]=_cga .Add (offX ,offY );};return _fbc ;};

// Point represents a two-dimensional point.
type Point struct{X float64 ;Y float64 ;};

// Add adds the specified vector to the current one and returns the result.
func (_dba Vector )Add (other Vector )Vector {_dba .Dx +=other .Dx ;_dba .Dy +=other .Dy ;return _dba };

// GetBounds returns the bounding box of the Bezier curve.
func (_ad CubicBezierCurve )GetBounds ()_fb .PdfRectangle {_cb :=_ad .P0 .X ;_cf :=_ad .P0 .X ;_d :=_ad .P0 .Y ;_cg :=_ad .P0 .Y ;for _ed :=0.0;_ed <=1.0;_ed +=0.001{Rx :=_ad .P0 .X *_f .Pow (1-_ed ,3)+_ad .P1 .X *3*_ed *_f .Pow (1-_ed ,2)+_ad .P2 .X *3*_f .Pow (_ed ,2)*(1-_ed )+_ad .P3 .X *_f .Pow (_ed ,3);
Ry :=_ad .P0 .Y *_f .Pow (1-_ed ,3)+_ad .P1 .Y *3*_ed *_f .Pow (1-_ed ,2)+_ad .P2 .Y *3*_f .Pow (_ed ,2)*(1-_ed )+_ad .P3 .Y *_f .Pow (_ed ,3);if Rx < _cb {_cb =Rx ;};if Rx > _cf {_cf =Rx ;};if Ry < _d {_d =Ry ;};if Ry > _cg {_cg =Ry ;};};_g :=_fb .PdfRectangle {};
_g .Llx =_cb ;_g .Lly =_d ;_g .Urx =_cf ;_g .Ury =_cg ;return _g ;};

// AppendCurve appends the specified Bezier curve to the path.
func (_ab CubicBezierPath )AppendCurve (curve CubicBezierCurve )CubicBezierPath {_ab .Curves =append (_ab .Curves ,curve );return _ab ;};

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_ga PolyBezierCurve )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){if _ga .BorderColor ==nil {_ga .BorderColor =_fb .NewPdfColorDeviceRGB (0,0,0);};_fdd :=NewCubicBezierPath ();for _ ,_ba :=range _ga .Curves {_fdd =_fdd .AppendCurve (_ba );
};_db :=_a .NewContentCreator ();_db .Add_q ();_ga .FillEnabled =_ga .FillEnabled &&_ga .FillColor !=nil ;if _ga .FillEnabled {_db .SetNonStrokingColor (_ga .FillColor );};_db .SetStrokingColor (_ga .BorderColor );_db .Add_w (_ga .BorderWidth );if len (gsName )> 1{_db .Add_gs (_fg .PdfObjectName (gsName ));
};for _ee ,_fgf :=range _fdd .Curves {if _ee ==0{_db .Add_m (_fgf .P0 .X ,_fgf .P0 .Y );}else {_db .Add_l (_fgf .P0 .X ,_fgf .P0 .Y );};_db .Add_c (_fgf .P1 .X ,_fgf .P1 .Y ,_fgf .P2 .X ,_fgf .P2 .Y ,_fgf .P3 .X ,_fgf .P3 .Y );};if _ga .FillEnabled {_db .Add_h ();
_db .Add_B ();}else {_db .Add_S ();};_db .Add_Q ();return _db .Bytes (),_fdd .GetBoundingBox ().ToPdfRectangle (),nil ;};

// GetBoundingBox returns the bounding box of the path.
func (_afa Path )GetBoundingBox ()BoundingBox {_afbd :=BoundingBox {};_dd :=0.0;_ef :=0.0;_ca :=0.0;_bb :=0.0;for _fbf ,_be :=range _afa .Points {if _fbf ==0{_dd =_be .X ;_ef =_be .X ;_ca =_be .Y ;_bb =_be .Y ;continue ;};if _be .X < _dd {_dd =_be .X ;
};if _be .X > _ef {_ef =_be .X ;};if _be .Y < _ca {_ca =_be .Y ;};if _be .Y > _bb {_bb =_be .Y ;};};_afbd .X =_dd ;_afbd .Y =_ca ;_afbd .Width =_ef -_dd ;_afbd .Height =_bb -_ca ;return _afbd ;};

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct{Rings [][]CubicBezierCurve ;FillEnabled bool ;FillColor _fb .PdfColor ;BorderEnabled bool ;BorderColor _fb .PdfColor ;BorderWidth float64 ;};

// Flip changes the sign of the vector: -vector.
func (_cda Vector )Flip ()Vector {_bff :=_cda .Magnitude ();_bca :=_cda .GetPolarAngle ();_cda .Dx =_bff *_f .Cos (_bca +_f .Pi );_cda .Dy =_bff *_f .Sin (_bca +_f .Pi );return _cda ;};

// NewPoint returns a new point with the coordinates x, y.
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct{Points []Point ;LineColor _fb .PdfColor ;LineWidth float64 ;};

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_fde Point )Add (dx ,dy float64 )Point {_fde .X +=dx ;_fde .Y +=dy ;return _fde };

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve (x0 ,y0 ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 )CubicBezierCurve {_fbd :=CubicBezierCurve {};_fbd .P0 =NewPoint (x0 ,y0 );_fbd .P1 =NewPoint (x1 ,y1 );_fbd .P2 =NewPoint (x2 ,y2 );_fbd .P3 =NewPoint (x3 ,y3 );return _fbd ;};

// Copy returns a clone of the path.
func (_fac Path )Copy ()Path {_fda :=Path {};_fda .Points =append (_fda .Points ,_fac .Points ...);return _fda ;};

// AppendPoint adds the specified point to the path.
func (_agg Path )AppendPoint (point Point )Path {_agg .Points =append (_agg .Points ,point );return _agg };

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor _fb .PdfColor ;Opacity float64 ;LineWidth float64 ;LineStyle LineStyle ;};

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator (bpath CubicBezierPath ,creator *_a .ContentCreator ){for _aggf ,_bdd :=range bpath .Curves {if _aggf ==0{creator .Add_m (_bdd .P0 .X ,_bdd .P0 .Y );};creator .Add_c (_bdd .P1 .X ,_bdd .P1 .Y ,_bdd .P2 .X ,_bdd .P2 .Y ,_bdd .P3 .X ,_bdd .P3 .Y );
};};

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar (length float64 ,theta float64 )Vector {_bgf :=Vector {};_bgf .Dx =length *_f .Cos (theta );_bgf .Dy =length *_f .Sin (theta );return _bgf ;};

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct{P0 Point ;P1 Point ;P2 Point ;P3 Point ;};

// GetBoundingBox returns the bounding box of the Bezier path.
func (_bc CubicBezierPath )GetBoundingBox ()Rectangle {_ac :=Rectangle {};_geb :=0.0;_dc :=0.0;_fc :=0.0;_afb :=0.0;for _fgb ,_dg :=range _bc .Curves {_fa :=_dg .GetBounds ();if _fgb ==0{_geb =_fa .Llx ;_dc =_fa .Urx ;_fc =_fa .Lly ;_afb =_fa .Ury ;continue ;
};if _fa .Llx < _geb {_geb =_fa .Llx ;};if _fa .Urx > _dc {_dc =_fa .Urx ;};if _fa .Lly < _fc {_fc =_fa .Lly ;};if _fa .Ury > _afb {_afb =_fa .Ury ;};};_ac .X =_geb ;_ac .Y =_fc ;_ac .Width =_dc -_geb ;_ac .Height =_afb -_fc ;return _ac ;};

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween (a Point ,b Point )Vector {_fddg :=Vector {};_fddg .Dx =b .X -a .X ;_fddg .Dy =b .Y -a .Y ;return _fddg ;};

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int ;

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_fgc Path )GetPointNumber (number int )Point {if number < 1||number > len (_fgc .Points ){return Point {};};return _fgc .Points [number -1];};

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_beg BoundingBox )ToPdfRectangle ()*_fb .PdfRectangle {return &_fb .PdfRectangle {Llx :_beg .X ,Lly :_beg .Y ,Urx :_beg .X +_beg .Width ,Ury :_beg .Y +_beg .Height };};

// FlipY flips the sign of the Dy component of the vector.
func (_fgfd Vector )FlipY ()Vector {_fgfd .Dy =-_fgfd .Dy ;return _fgfd };

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_fbe CurvePolygon )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){_da :=_a .NewContentCreator ();_da .Add_q ();_fbe .FillEnabled =_fbe .FillEnabled &&_fbe .FillColor !=nil ;if _fbe .FillEnabled {_da .SetNonStrokingColor (_fbe .FillColor );
};_fbe .BorderEnabled =_fbe .BorderEnabled &&_fbe .BorderColor !=nil ;if _fbe .BorderEnabled {_da .SetStrokingColor (_fbe .BorderColor );_da .Add_w (_fbe .BorderWidth );};if len (gsName )> 1{_da .Add_gs (_fg .PdfObjectName (gsName ));};_bbb :=NewCubicBezierPath ();
for _ ,_bab :=range _fbe .Rings {for _de ,_ddf :=range _bab {if _de ==0{_da .Add_m (_ddf .P0 .X ,_ddf .P0 .Y );}else {_da .Add_l (_ddf .P0 .X ,_ddf .P0 .Y );};_da .Add_c (_ddf .P1 .X ,_ddf .P1 .Y ,_ddf .P2 .X ,_ddf .P2 .Y ,_ddf .P3 .X ,_ddf .P3 .Y );_bbb =_bbb .AppendCurve (_ddf );
};_da .Add_h ();};if _fbe .FillEnabled &&_fbe .BorderEnabled {_da .Add_B ();}else if _fbe .FillEnabled {_da .Add_f ();}else if _fbe .BorderEnabled {_da .Add_S ();};_da .Add_Q ();return _da .Bytes (),_bbb .GetBoundingBox ().ToPdfRectangle (),nil ;};

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_dad Rectangle )ToPdfRectangle ()*_fb .PdfRectangle {return &_fb .PdfRectangle {Llx :_dad .X ,Lly :_dad .Y ,Urx :_dad .X +_dad .Width ,Ury :_dad .Y +_dad .Height };};

// Length returns the number of points in the path.
func (_aaf Path )Length ()int {return len (_aaf .Points )};

// Copy returns a clone of the Bezier path.
func (_af CubicBezierPath )Copy ()CubicBezierPath {_ff :=CubicBezierPath {};_ff .Curves =append (_ff .Curves ,_af .Curves ...);return _ff ;};

// Draw draws the rectangle. A graphics state can be specified for
// setting additional properties (e.g. opacity). Otherwise pass an empty string
// for the `gsName` parameter. The method returns the content stream as a byte
// array and the bounding box of the shape.
func (_eed Rectangle )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){_cc :=_a .NewContentCreator ();_cc .Add_q ();if _eed .FillEnabled {_cc .SetNonStrokingColor (_eed .FillColor );};if _eed .BorderEnabled {_cc .SetStrokingColor (_eed .BorderColor );
_cc .Add_w (_eed .BorderWidth );};if len (gsName )> 1{_cc .Add_gs (_fg .PdfObjectName (gsName ));};var (_aca ,_bea =_eed .X ,_eed .Y ;_ggf ,_cgb =_eed .Width ,_eed .Height ;_eec =_f .Abs (_eed .BorderRadiusTopLeft );_ega =_f .Abs (_eed .BorderRadiusTopRight );
_fcb =_f .Abs (_eed .BorderRadiusBottomLeft );_gd =_f .Abs (_eed .BorderRadiusBottomRight );_fdf =0.4477;);_egaf :=Path {Points :[]Point {{X :_aca +_ggf -_gd ,Y :_bea },{X :_aca +_ggf ,Y :_bea +_cgb -_ega },{X :_aca +_eec ,Y :_bea +_cgb },{X :_aca ,Y :_bea +_fcb }}};
_gec :=[][7]float64 {{_gd ,_aca +_ggf -_gd *_fdf ,_bea ,_aca +_ggf ,_bea +_gd *_fdf ,_aca +_ggf ,_bea +_gd },{_ega ,_aca +_ggf ,_bea +_cgb -_ega *_fdf ,_aca +_ggf -_ega *_fdf ,_bea +_cgb ,_aca +_ggf -_ega ,_bea +_cgb },{_eec ,_aca +_eec *_fdf ,_bea +_cgb ,_aca ,_bea +_cgb -_eec *_fdf ,_aca ,_bea +_cgb -_eec },{_fcb ,_aca ,_bea +_fcb *_fdf ,_aca +_fcb *_fdf ,_bea ,_aca +_fcb ,_bea }};
_cc .Add_m (_aca +_fcb ,_bea );for _aad :=0;_aad < 4;_aad ++{_dfe :=_egaf .Points [_aad ];_cc .Add_l (_dfe .X ,_dfe .Y );_ffc :=_gec [_aad ];if _edf :=_ffc [0];_edf !=0{_cc .Add_c (_ffc [1],_ffc [2],_ffc [3],_ffc [4],_ffc [5],_ffc [6]);};};_cc .Add_h ();
if _eed .FillEnabled &&_eed .BorderEnabled {_cc .Add_B ();}else if _eed .FillEnabled {_cc .Add_f ();}else if _eed .BorderEnabled {_cc .Add_S ();};_cc .Add_Q ();return _cc .Bytes (),_egaf .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_fdae BasicLine )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){_fabe :=_fdae .LineWidth ;_dagd :=NewPath ();_dagd =_dagd .AppendPoint (NewPoint (_fdae .X1 ,_fdae .Y1 ));_dagd =_dagd .AppendPoint (NewPoint (_fdae .X2 ,_fdae .Y2 ));_dgd :=_a .NewContentCreator ();
_fce :=_dagd .GetBoundingBox ();DrawPathWithCreator (_dagd ,_dgd );if _fdae .LineStyle ==LineStyleDashed {_dgd .Add_d ([]int64 {1,1},0);};_dgd .SetStrokingColor (_fdae .LineColor ).Add_w (_fabe ).Add_S ().Add_Q ();return _dgd .Bytes (),_fce .ToPdfRectangle (),nil ;
};

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_ec Polygon )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){_bac :=_a .NewContentCreator ();_bac .Add_q ();_ec .FillEnabled =_ec .FillEnabled &&_ec .FillColor !=nil ;if _ec .FillEnabled {_bac .SetNonStrokingColor (_ec .FillColor );};_ec .BorderEnabled =_ec .BorderEnabled &&_ec .BorderColor !=nil ;
if _ec .BorderEnabled {_bac .SetStrokingColor (_ec .BorderColor );_bac .Add_w (_ec .BorderWidth );};if len (gsName )> 1{_bac .Add_gs (_fg .PdfObjectName (gsName ));};_age :=NewPath ();for _ ,_fbfb :=range _ec .Points {for _cfd ,_gf :=range _fbfb {_age =_age .AppendPoint (_gf );
if _cfd ==0{_bac .Add_m (_gf .X ,_gf .Y );}else {_bac .Add_l (_gf .X ,_gf .Y );};};_bac .Add_h ();};if _ec .FillEnabled &&_ec .BorderEnabled {_bac .Add_B ();}else if _ec .FillEnabled {_bac .Add_f ();}else if _ec .BorderEnabled {_bac .Add_S ();};_bac .Add_Q ();
return _bac .Bytes (),_age .GetBoundingBox ().ToPdfRectangle (),nil ;};

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator (path Path ,creator *_a .ContentCreator ){for _bf ,_ecf :=range path .Points {if _bf ==0{creator .Add_m (_ecf .X ,_ecf .Y );}else {creator .Add_l (_ecf .X ,_ecf .Y );};};};

// Vector represents a two-dimensional vector.
type Vector struct{Dx float64 ;Dy float64 ;};const (LineStyleSolid LineStyle =0;LineStyleDashed LineStyle =1;);

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_gdc Polyline )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){if _gdc .LineColor ==nil {_gdc .LineColor =_fb .NewPdfColorDeviceRGB (0,0,0);};_fgg :=NewPath ();for _ ,_dagf :=range _gdc .Points {_fgg =_fgg .AppendPoint (_dagf );};_eaa :=_a .NewContentCreator ();
_eaa .Add_q ().SetStrokingColor (_gdc .LineColor ).Add_w (_gdc .LineWidth );if len (gsName )> 1{_eaa .Add_gs (_fg .PdfObjectName (gsName ));};DrawPathWithCreator (_fgg ,_eaa );_eaa .Add_S ();_eaa .Add_Q ();return _eaa .Bytes (),_fgg .GetBoundingBox ().ToPdfRectangle (),nil ;
};

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct{Points [][]Point ;FillEnabled bool ;FillColor _fb .PdfColor ;BorderEnabled bool ;BorderColor _fb .PdfColor ;BorderWidth float64 ;};

// Rotate rotates the vector by the specified angle.
func (_ebdg Vector )Rotate (phi float64 )Vector {_egb :=_ebdg .Magnitude ();_acaa :=_ebdg .GetPolarAngle ();return NewVectorPolar (_egb ,_acaa +phi );};

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_dge Line )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){_fbcc ,_ccc :=_dge .X1 ,_dge .X2 ;_edfa ,_gea :=_dge .Y1 ,_dge .Y2 ;_ccg :=_gea -_edfa ;_dbe :=_ccc -_fbcc ;_fgbe :=_f .Atan2 (_ccg ,_dbe );L :=_f .Sqrt (_f .Pow (_dbe ,2.0)+_f .Pow (_ccg ,2.0));
_fdg :=_dge .LineWidth ;_acb :=_f .Pi ;_gde :=1.0;if _dbe < 0{_gde *=-1.0;};if _ccg < 0{_gde *=-1.0;};VsX :=_gde *(-_fdg /2*_f .Cos (_fgbe +_acb /2));VsY :=_gde *(-_fdg /2*_f .Sin (_fgbe +_acb /2)+_fdg *_f .Sin (_fgbe +_acb /2));V1X :=VsX +_fdg /2*_f .Cos (_fgbe +_acb /2);
V1Y :=VsY +_fdg /2*_f .Sin (_fgbe +_acb /2);V2X :=VsX +_fdg /2*_f .Cos (_fgbe +_acb /2)+L *_f .Cos (_fgbe );V2Y :=VsY +_fdg /2*_f .Sin (_fgbe +_acb /2)+L *_f .Sin (_fgbe );V3X :=VsX +_fdg /2*_f .Cos (_fgbe +_acb /2)+L *_f .Cos (_fgbe )+_fdg *_f .Cos (_fgbe -_acb /2);
V3Y :=VsY +_fdg /2*_f .Sin (_fgbe +_acb /2)+L *_f .Sin (_fgbe )+_fdg *_f .Sin (_fgbe -_acb /2);V4X :=VsX +_fdg /2*_f .Cos (_fgbe -_acb /2);V4Y :=VsY +_fdg /2*_f .Sin (_fgbe -_acb /2);_gc :=NewPath ();_gc =_gc .AppendPoint (NewPoint (V1X ,V1Y ));_gc =_gc .AppendPoint (NewPoint (V2X ,V2Y ));
_gc =_gc .AppendPoint (NewPoint (V3X ,V3Y ));_gc =_gc .AppendPoint (NewPoint (V4X ,V4Y ));_gad :=_dge .LineEndingStyle1 ;_fab :=_dge .LineEndingStyle2 ;_cd :=3*_fdg ;_ddc :=3*_fdg ;_gga :=(_ddc -_fdg )/2;if _fab ==LineEndingStyleArrow {_ceff :=_gc .GetPointNumber (2);
_ea :=NewVectorPolar (_cd ,_fgbe +_acb );_egc :=_ceff .AddVector (_ea );_ebd :=NewVectorPolar (_ddc /2,_fgbe +_acb /2);_dag :=NewVectorPolar (_cd ,_fgbe );_gef :=NewVectorPolar (_gga ,_fgbe +_acb /2);_bd :=_egc .AddVector (_gef );_fdb :=_dag .Add (_ebd .Flip ());
_bge :=_bd .AddVector (_fdb );_edd :=_ebd .Scale (2).Flip ().Add (_fdb .Flip ());_bacg :=_bge .AddVector (_edd );_agee :=_egc .AddVector (NewVectorPolar (_fdg ,_fgbe -_acb /2));_aaec :=NewPath ();_aaec =_aaec .AppendPoint (_gc .GetPointNumber (1));_aaec =_aaec .AppendPoint (_egc );
_aaec =_aaec .AppendPoint (_bd );_aaec =_aaec .AppendPoint (_bge );_aaec =_aaec .AppendPoint (_bacg );_aaec =_aaec .AppendPoint (_agee );_aaec =_aaec .AppendPoint (_gc .GetPointNumber (4));_gc =_aaec ;};if _gad ==LineEndingStyleArrow {_def :=_gc .GetPointNumber (1);
_fdde :=_gc .GetPointNumber (_gc .Length ());_ccf :=NewVectorPolar (_fdg /2,_fgbe +_acb +_acb /2);_eeb :=_def .AddVector (_ccf );_bbe :=NewVectorPolar (_cd ,_fgbe ).Add (NewVectorPolar (_ddc /2,_fgbe +_acb /2));_gb :=_eeb .AddVector (_bbe );_aba :=NewVectorPolar (_gga ,_fgbe -_acb /2);
_acd :=_gb .AddVector (_aba );_abg :=NewVectorPolar (_cd ,_fgbe );_egf :=_fdde .AddVector (_abg );_fae :=NewVectorPolar (_gga ,_fgbe +_acb +_acb /2);_ead :=_egf .AddVector (_fae );_deff :=_eeb ;_gdd :=NewPath ();_gdd =_gdd .AppendPoint (_eeb );_gdd =_gdd .AppendPoint (_gb );
_gdd =_gdd .AppendPoint (_acd );for _ ,_dfc :=range _gc .Points [1:len (_gc .Points )-1]{_gdd =_gdd .AppendPoint (_dfc );};_gdd =_gdd .AppendPoint (_egf );_gdd =_gdd .AppendPoint (_ead );_gdd =_gdd .AppendPoint (_deff );_gc =_gdd ;};_eab :=_a .NewContentCreator ();
_eab .Add_q ().SetNonStrokingColor (_dge .LineColor );if len (gsName )> 1{_eab .Add_gs (_fg .PdfObjectName (gsName ));};_gc =_gc .Offset (_dge .X1 ,_dge .Y1 );_dfg :=_gc .GetBoundingBox ();DrawPathWithCreator (_gc ,_eab );if _dge .LineStyle ==LineStyleDashed {_eab .Add_d ([]int64 {1,1},0).Add_S ().Add_f ().Add_Q ();
}else {_eab .Add_f ().Add_Q ();};return _eab .Bytes (),_dfg .ToPdfRectangle (),nil ;};

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor _fb .PdfColor ;Opacity float64 ;LineWidth float64 ;LineEndingStyle1 LineEndingStyle ;LineEndingStyle2 LineEndingStyle ;LineStyle LineStyle ;};

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{Curves []CubicBezierCurve ;};

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_cbe Vector )GetPolarAngle ()float64 {return _f .Atan2 (_cbe .Dy ,_cbe .Dx )};func (_fcc Point )String ()string {return _c .Sprintf ("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029",_fcc .X ,_fcc .Y );};

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_b CubicBezierCurve )AddOffsetXY (offX ,offY float64 )CubicBezierCurve {_b .P0 .X +=offX ;_b .P1 .X +=offX ;_b .P2 .X +=offX ;_b .P3 .X +=offX ;_b .P0 .Y +=offY ;_b .P1 .Y +=offY ;_b .P2 .Y +=offY ;_b .P3 .Y +=offY ;return _b ;};

// LineStyle refers to how the line will be created.
type LineStyle int ;

// NewPath returns a new empty path.
func NewPath ()Path {return Path {}};

// AddVector adds vector to a point.
func (_cfc Point )AddVector (v Vector )Point {_cfc .X +=v .Dx ;_cfc .Y +=v .Dy ;return _cfc };

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct{

// Position and size properties.
X float64 ;Y float64 ;Width float64 ;Height float64 ;

// Fill properties.
FillEnabled bool ;FillColor _fb .PdfColor ;

// Border properties.
BorderEnabled bool ;BorderColor _fb .PdfColor ;BorderWidth float64 ;BorderRadiusTopLeft float64 ;BorderRadiusTopRight float64 ;BorderRadiusBottomLeft float64 ;BorderRadiusBottomRight float64 ;

// Shape opacity (0-1 interval).
Opacity float64 ;};

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_ce Circle )Draw (gsName string )([]byte ,*_fb .PdfRectangle ,error ){_aae :=_ce .Width /2;_dce :=_ce .Height /2;if _ce .BorderEnabled {_aae -=_ce .BorderWidth /2;_dce -=_ce .BorderWidth /2;};_ae :=0.551784;_cef :=_aae *_ae ;_fe :=_dce *_ae ;_bg :=NewCubicBezierPath ();
_bg =_bg .AppendCurve (NewCubicBezierCurve (-_aae ,0,-_aae ,_fe ,-_cef ,_dce ,0,_dce ));_bg =_bg .AppendCurve (NewCubicBezierCurve (0,_dce ,_cef ,_dce ,_aae ,_fe ,_aae ,0));_bg =_bg .AppendCurve (NewCubicBezierCurve (_aae ,0,_aae ,-_fe ,_cef ,-_dce ,0,-_dce ));
_bg =_bg .AppendCurve (NewCubicBezierCurve (0,-_dce ,-_cef ,-_dce ,-_aae ,-_fe ,-_aae ,0));_bg =_bg .Offset (_aae ,_dce );if _ce .BorderEnabled {_bg =_bg .Offset (_ce .BorderWidth /2,_ce .BorderWidth /2);};if _ce .X !=0||_ce .Y !=0{_bg =_bg .Offset (_ce .X ,_ce .Y );
};_eg :=_a .NewContentCreator ();_eg .Add_q ();if _ce .FillEnabled {_eg .SetNonStrokingColor (_ce .FillColor );};if _ce .BorderEnabled {_eg .SetStrokingColor (_ce .BorderColor );_eg .Add_w (_ce .BorderWidth );};if len (gsName )> 1{_eg .Add_gs (_fg .PdfObjectName (gsName ));
};DrawBezierPathWithCreator (_bg ,_eg );_eg .Add_h ();if _ce .FillEnabled &&_ce .BorderEnabled {_eg .Add_B ();}else if _ce .FillEnabled {_eg .Add_f ();}else if _ce .BorderEnabled {_eg .Add_S ();};_eg .Add_Q ();_gg :=_bg .GetBoundingBox ();if _ce .BorderEnabled {_gg .Height +=_ce .BorderWidth ;
_gg .Width +=_ce .BorderWidth ;_gg .X -=_ce .BorderWidth /2;_gg .Y -=_ce .BorderWidth /2;};return _eg .Bytes (),_gg .ToPdfRectangle (),nil ;};

// Scale scales the vector by the specified factor.
func (_gfg Vector )Scale (factor float64 )Vector {_gbc :=_gfg .Magnitude ();_defb :=_gfg .GetPolarAngle ();_gfg .Dx =factor *_gbc *_f .Cos (_defb );_gfg .Dy =factor *_gbc *_f .Sin (_defb );return _gfg ;};