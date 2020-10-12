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

// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson ;import (_e "encoding/json";_g "github.com/unidoc/unipdf/v3/core";_gg "github.com/unidoc/unipdf/v3/model";_c "io";_dc "os";);

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_a []fieldValue };

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_ac ,_gdd :=_dc .Open (filePath );if _gdd !=nil {return nil ,_gdd ;};defer _ac .Close ();return LoadFromPDF (_ac );};

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_ea ,_eab :=_dc .Open (filePath );if _eab !=nil {return nil ,_eab ;};defer _ea .Close ();return LoadFromJSON (_ea );};

// JSON returns the field data as a string in JSON format.
func (_egb FieldData )JSON ()(string ,error ){_cd ,_fd :=_e .MarshalIndent (_egb ._a ,"","\u0020\u0020\u0020\u0020");return string (_cd ),_fd ;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _c .ReadSeeker )(*FieldData ,error ){_ead ,_gc :=_gg .NewPdfReader (rs );if _gc !=nil {return nil ,_gc ;};if _ead .AcroForm ==nil {return nil ,nil ;};var _f []fieldValue ;_de :=_ead .AcroForm .AllFields ();for _ ,_gce :=range _de {var _db []string ;_gb :=make (map[string ]struct{});_ef ,_df :=_gce .FullName ();if _df !=nil {return nil ,_df ;};if _cc ,_fc :=_gce .V .(*_g .PdfObjectString );_fc {_f =append (_f ,fieldValue {Name :_ef ,Value :_cc .Decoded ()});continue ;};var _gf string ;for _ ,_fb :=range _gce .Annotations {_ba ,_ffc :=_g .GetName (_fb .AS );if _ffc {_gf =_ba .String ();};_bag ,_gbc :=_g .GetDict (_fb .AP );if !_gbc {continue ;};_dg ,_ :=_g .GetDict (_bag .Get ("\u004e"));for _ ,_ae :=range _dg .Keys (){_dgc :=_ae .String ();if _ ,_bb :=_gb [_dgc ];!_bb {_db =append (_db ,_dgc );_gb [_dgc ]=struct{}{};};};_egd ,_ :=_g .GetDict (_bag .Get ("\u0044"));for _ ,_dea :=range _egd .Keys (){_gcd :=_dea .String ();if _ ,_gad :=_gb [_gcd ];!_gad {_db =append (_db ,_gcd );_gb [_gcd ]=struct{}{};};};};_ffca :=fieldValue {Name :_ef ,Value :_gf ,Options :_db };_f =append (_f ,_ffca );};_gd :=FieldData {_a :_f };return &_gd ,nil ;};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _c .Reader )(*FieldData ,error ){var _b FieldData ;_eg :=_e .NewDecoder (r ).Decode (&_b ._a );if _eg !=nil {return nil ,_eg ;};return &_b ,nil ;};

// FieldValues implements model.FieldValueProvider interface.
func (_ccg *FieldData )FieldValues ()(map[string ]_g .PdfObject ,error ){_fdb :=make (map[string ]_g .PdfObject );for _ ,_fdg :=range _ccg ._a {if len (_fdg .Value )> 0{_fdb [_fdg .Name ]=_g .MakeString (_fdg .Value );};};return _fdb ,nil ;};type fieldValue struct{Name string `json:"name"`;Value string `json:"value"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};