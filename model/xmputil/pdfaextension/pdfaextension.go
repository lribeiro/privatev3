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

package pdfaextension ;import (_f "fmt";_ada "github.com/trimmer-io/go-xmp/models/dc";_da "github.com/trimmer-io/go-xmp/models/pdf";_ec "github.com/trimmer-io/go-xmp/models/xmp_base";_fb "github.com/trimmer-io/go-xmp/models/xmp_mm";_b "github.com/trimmer-io/go-xmp/xmp";
_de "reflect";_e "sort";_ad "strings";_d "sync";);

// Can implements xmp.Model interface.
func (_gg *Model )Can (nsName string )bool {return Namespace .GetName ()==nsName };func init (){_b .Register (Namespace ,_b .XmpMetadata );_b .Register (SchemaNS );_b .Register (PropertyNS );_b .Register (ValueTypeNS );_b .Register (FieldNS );};var (_ac =map[string ]*Schema {};
_gbc _d .RWMutex ;);

// SetTag implements xmp.Model interface.
func (_bb *Model )SetTag (tag ,value string )error {if _cfb :=_b .SetNativeField (_bb ,tag ,value );_cfb !=nil {return _f .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_cfb );};return nil ;};

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (_cda *PropertyCategory )UnmarshalText (in []byte )error {switch string (in ){case "\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_cda =PropertyCategoryInternal ;case "\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_cda =PropertyCategoryExternal ;
default:*_cda =PropertyCategoryUndefined ;};return nil ;};

// GetSchema for provided namespace.
func GetSchema (namespaceURI string )(*Schema ,bool ){_gbc .RLock ();defer _gbc .RUnlock ();_add ,_cc :=_ac [namespaceURI ];return _add ,_cc ;};

// PropertyCategory is the property category enumerator.
type PropertyCategory int ;

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_gd *Schemas )UnmarshalXMP (d *_b .Decoder ,node *_b .Node ,m _b .Model )error {return _b .UnmarshalArray (d ,node ,_gd .Typ (),_gd );};func _ca (_cac _b .Model ,_bgb *_b .Namespace ,_eg Property )bool {_eea :=_de .ValueOf (_cac ).Elem ();if _eea .Kind ()==_de .Ptr {_eea =_eea .Elem ();
};_ffg :=_eea .Type ();for _caca :=0;_caca < _ffg .NumField ();_caca ++{_gf :=_ffg .Field (_caca );_egd :=_gf .Tag .Get ("\u0078\u006d\u0070");if _egd ==""{continue ;};if !_ad .HasPrefix (_egd ,_bgb .Name ){continue ;};_fbcg :=_ad .IndexRune (_egd ,':');
if _fbcg ==-1{continue ;};_dec :=_egd [_fbcg +1:];if _dec ==_eg .Name {_ea :=_eea .Field (_caca );return !_ea .IsZero ();};};return false ;};

// RegisterSchema registers schema extension definition.
func RegisterSchema (ns *_b .Namespace ,schema *Schema ){_gbc .Lock ();defer _gbc .Unlock ();_ac [ns .URI ]=schema ;};

// NewModel creates a new pdfAExtension model.
func NewModel (name string )_b .Model {return &Model {}};var PdfSchema =Schema {NamespaceURI :"\u0068\u0074\u0074\u0070:\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065.\u0063o\u006d\u002f\u0070\u0064\u0066\u002f\u0031.\u0033\u002f",Prefix :"\u0070\u0064\u0066",Schema :"\u0041\u0064o\u0062\u0065\u0020P\u0044\u0046\u0020\u0053\u0063\u0068\u0065\u006d\u0061",Property :Properties {{Category :PropertyCategoryInternal ,Description :"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",Name :"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0050\u0044F\u0020\u0066\u0069l\u0065\u0020\u0076e\u0072\u0073\u0069\u006f\u006e\u0020\u0028\u0066\u006f\u0072 \u0065\u0078\u0061\u006d\u0070le\u003a\u0020\u0031\u002e\u0030\u002c\u0020\u0031\u002e\u0033\u002c\u0020\u0061\u006e\u0064\u0020\u0073\u006f\u0020\u006f\u006e\u0029\u002e",Name :"\u0050\u0044\u0046\u0056\u0065\u0072\u0073\u0069\u006f\u006e",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0074\u006f\u006fl\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074.",Name :"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",ValueType :ValueTypeNameAgentName }}};


// MarshalText implements encoding.TextMarshaler interface.
func (_ecg PropertyCategory )MarshalText ()([]byte ,error ){switch _ecg {case PropertyCategoryInternal :return []byte ("\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c"),nil ;case PropertyCategoryExternal :return []byte ("\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c"),nil ;
case PropertyCategoryUndefined :return []byte (""),nil ;default:return nil ,_f .Errorf ("\u0075\u006ed\u0065\u0066\u0069\u006ee\u0064\u0020p\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020c\u0061\u0074\u0065\u0067\u006f\u0072\u0079\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_ecg );
};};

// Property is a schema that describes single property.
type Property struct{

// Category is the property category.
Category PropertyCategory `xmp:"pdfaProperty:category"`;

// Description of the property.
Description string `xmp:"pdfaProperty:description"`;

// Name is a property name.
Name string `xmp:"pdfaProperty:name"`;

// ValueType is the property value type.
ValueType ValueTypeName `xmp:"pdfaProperty:valueType"`;};var FieldResourceEventSchema =Schema {NamespaceURI :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u006es\u002e\u0061\u0064ob\u0065\u002e\u0063\u006f\u006d\u002fx\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065\u002f\u0052\u0065s\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065n\u0074\u0023",Prefix :"\u0073\u0074\u0045v\u0074",Schema :"\u0044e\u0066\u0069n\u0069\u0074\u0069\u006fn\u0020\u006f\u0066 \u0062\u0061\u0073\u0069\u0063\u0020\u0076\u0061\u006cue\u0020\u0074\u0079p\u0065\u0020R\u0065\u0073\u006f\u0075\u0072\u0063e\u0045\u0076e\u006e\u0074",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0054he\u0020a\u0063t\u0069\u006f\u006e\u0020\u0074\u0068a\u0074\u0020\u006f\u0063c\u0075\u0072\u0072\u0065\u0064\u002e\u0020\u0044\u0065\u0066\u0069\u006e\u0065d \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0061\u0072\u0065\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065d\u002c \u0063\u006f\u0070\u0069\u0065\u0064\u002c\u0020\u0063\u0072\u0065\u0061\u0074e\u0064\u002c \u0063\u0072\u006fp\u0070\u0065\u0064\u002c\u0020\u0065\u0064\u0069\u0074ed\u002c\u0020\u0066i\u006c\u002d\u0074\u0065r\u0065\u0064\u002c\u0020\u0066\u006fr\u006d\u0061t\u0074\u0065\u0064\u002c\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u005f\u0075\u0070\u0064\u0061\u0074\u0065\u0064\u002c\u0020\u0070\u0072\u0069\u006e\u0074\u0065\u0064\u002c\u0020\u0070ubli\u0073\u0068\u0065\u0064\u002c\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0064\u002c\u0020\u0070\u0072\u006f\u0064\u0075\u0063\u0065\u0064\u002c\u0020\u0072\u0065\u0073i\u007ae\u0064.\u004e\u0065\u0077\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u0076\u0065r\u0062\u0073 \u0069n\u0020\u0074\u0068\u0065\u0020\u0070a\u0073\u0074\u0020\u0074\u0065\u006e\u0073\u0065\u002e",Name :"\u0061\u0063\u0074\u0069\u006f\u006e",ValueType :"O\u0070\u0065\u006e\u0020\u0043\u0068\u006f\u0069\u0063\u0065"},{Category :PropertyCategoryInternal ,Description :"T\u0068\u0065\u0020\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0020\u0049\u0044\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u006d\u006f\u0064\u0069\u0066\u0069\u0065d \u0072\u0065\u0073o\u0075r\u0063\u0065",Name :"\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0041\u0064di\u0074\u0069\u006fn\u0061\u006c\u0020\u0064esc\u0072ip\u0074\u0069\u006f\u006e\u0020\u006f\u0066 t\u0068\u0065\u0020\u0061\u0063\u0074\u0069o\u006e",Name :"\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054h\u0065\u0020s\u006f\u0066\u0074\u0077a\u0072\u0065\u0020a\u0067\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074 p\u0065\u0072\u0066o\u0072\u006de\u0064\u0020\u0074\u0068\u0065\u0020a\u0063\u0074i\u006f\u006e",Name :"\u0073\u006f\u0066\u0074\u0077\u0061\u0072\u0065\u0041\u0067\u0065\u006e\u0074",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u004f\u0070t\u0069\u006f\u006e\u0061\u006c\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u006f\u0066\u0020\u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u006f\u0063\u0063\u0075\u0072\u0072\u0065\u0064",Name :"\u0077\u0068\u0065\u006e",ValueType :ValueTypeNameDate }},ValueType :nil };


// MarshalXMP implements xmp.Marshaler interface.
func (_eeb Properties )MarshalXMP (e *_b .Encoder ,node *_b .Node ,m _b .Model )error {return _b .MarshalArray (e ,node ,_eeb .Typ (),_eeb );};

// IsZero checks if the resources list has no entries.
func (_bbb Schemas )IsZero ()bool {return len (_bbb )==0};func init (){_ac =map[string ]*Schema {_fb .NsXmpMM .URI :&XmpMediaManagementSchema ,"\u0078\u006d\u0070\u0069\u0064\u0071":&XmpIDQualSchema ,_da .NsPDF .URI :&PdfSchema ,"\u0073\u0074\u0045v\u0074":&FieldResourceEventSchema ,"\u0073\u0074\u0056e\u0072":&FieldVersionSchema };
};var (Namespace =_b .NewNamespace ("\u0070\u0064\u0066\u0061\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e","\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002f\u0070\u0064\u0066\u0061/\u006e\u0073\u002f\u0065\u0078t\u0065\u006es\u0069\u006f\u006e\u002f",NewModel );
SchemaNS =_b .NewNamespace ("\u0070\u0064\u0066\u0061\u0053\u0063\u0068\u0065\u006d\u0061","h\u0074\u0074\u0070\u003a\u002f\u002fw\u0077\u0077\u002e\u0061\u0069\u0069m\u002e\u006f\u0072\u0067\u002f\u0070\u0064f\u0061\u002f\u006e\u0073\u002f\u0073\u0063\u0068\u0065\u006da\u0023",nil );
PropertyNS =_b .NewNamespace ("\u0070\u0064\u0066a\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0079","\u0068\u0074t\u0070\u003a\u002f\u002fw\u0077\u0077.\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067/\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0070\u0072\u006f\u0070e\u0072\u0074\u0079\u0023",nil );
ValueTypeNS =_b .NewNamespace ("\u0070\u0064\u0066\u0061\u0054\u0079\u0070\u0065","\u0068\u0074\u0074\u0070\u003a\u002f/\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002ft\u0079\u0070\u0065\u0023",nil );
FieldNS =_b .NewNamespace ("\u0070d\u0066\u0061\u0046\u0069\u0065\u006cd","\u0068\u0074\u0074p:\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069m\u002eo\u0072g\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0066\u0069\u0065\u006c\u0064\u0023",nil );
);

// ValueType is the pdfa extension value type schema.
type ValueType struct{Description string `xmp:"pdfaType:description"`;Field []FieldValueType `xmp:"pdfaType:field"`;NamespaceURI string `xmp:"pdfaType:namespaceURI"`;Prefix string `xmp:"pdfaType:prefix"`;Type string `xmp:"pdfaType:type"`;};

// SyncFromXMP implements xmp.Model interface.
func (_agd *Model )SyncFromXMP (d *_b .Document )error {return nil };var XmpMediaManagementSchema =Schema {NamespaceURI :"\u0068\u0074\u0074p\u003a\u002f\u002f\u006es\u002e\u0061\u0064\u006f\u0062\u0065\u002ec\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u006d\u006d\u002f",Prefix :"\u0078\u006d\u0070M\u004d",Schema :"X\u004dP\u0020\u004d\u0065\u0064\u0069\u0061\u0020\u004da\u006e\u0061\u0067\u0065me\u006e\u0074",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0041\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006fr\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020w\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u006f\u006e\u0065\u0020i\u0073\u0020\u0064e\u0072\u0069\u0076\u0065\u0064\u002e",Name :"D\u0065\u0072\u0069\u0076\u0065\u0064\u0046\u0072\u006f\u006d",ValueType :ValueTypeNameResourceRef },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0072\u0065\u0073\u006fu\u0072\u0063\u0065",Name :"\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"UU\u0049\u0044 \u0062\u0061\u0073\u0065\u0064\u0020\u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0020\u0069\u006e\u0063\u0061\u0072\u006e\u0061\u0074i\u006fn\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u006fc\u0075m\u0065\u006et",Name :"\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074",Name :"\u004fr\u0069g\u0069\u006e\u0061\u006c\u0044o\u0063\u0075m\u0065\u006e\u0074\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0069\u0073 \u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065",Name :"\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073",ValueType :ValueTypeNameRenditionClass },{Category :PropertyCategoryInternal ,Description :"\u0043\u0061n\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020t\u006f\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0072\u0065\u006e\u0064\u0069t\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0074\u0068a\u0074 \u0061r\u0065\u0020\u0074o\u006f\u0020\u0063\u006f\u006d\u0070\u006c\u0065\u0078\u0020\u006f\u0072\u0020\u0076\u0065\u0072\u0062o\u0073\u0065\u0020\u0074\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0069\u006e",Name :"\u0052e\u006ed\u0069\u0074\u0069\u006f\u006e\u0050\u0061\u0072\u0061\u006d\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u0069\u0064\u0065\u006e\u0074i\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0074\u0068i\u0073\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u002e",Name :"\u0056e\u0072\u0073\u0069\u006f\u006e\u0049D",ValueType :ValueTypeNameText },{Category :PropertyCategoryExternal ,Description :"\u0054\u0068\u0065\u0020\u0076\u0065r\u0073\u0069\u006f\u006e\u0020\u0068\u0069\u0073\u0074\u006f\u0072\u0079\u0020\u0061\u0073\u0073\u006f\u0063\u0069\u0061t\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0072e\u0073o\u0075\u0072\u0063\u0065",Name :"\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0073",ValueType :SeqOfValueTypeName (ValueTypeNameVersion )},{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0072\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075r\u0063\u0065\u0027\u0073\u0020\u006d\u0061n\u0061\u0067\u0065\u0072",Name :"\u004da\u006e\u0061\u0067\u0065\u0072",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e",Name :"\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e",Name :"\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074",ValueType :ValueTypeNameText }}};
const (ValueTypeNameBoolean ValueTypeName ="\u0042o\u006f\u006c\u0065\u0061\u006e";ValueTypeNameDate ValueTypeName ="\u0044\u0061\u0074\u0065";ValueTypeNameInteger ValueTypeName ="\u0049n\u0074\u0065\u0067\u0065\u0072";ValueTypeNameReal ValueTypeName ="\u0052\u0065\u0061\u006c";
ValueTypeNameText ValueTypeName ="\u0054\u0065\u0078\u0074";ValueTypeNameAgentName ValueTypeName ="\u0041g\u0065\u006e\u0074\u004e\u0061\u006de";ValueTypeNameProperName ValueTypeName ="\u0050\u0072\u006f\u0070\u0065\u0072\u004e\u0061\u006d\u0065";ValueTypeNameXPath ValueTypeName ="\u0058\u0050\u0061t\u0068";
ValueTypeNameGUID ValueTypeName ="\u0047\u0055\u0049\u0044";ValueTypeNameLocale ValueTypeName ="\u004c\u006f\u0063\u0061\u006c\u0065";ValueTypeNameMIMEType ValueTypeName ="\u004d\u0049\u004d\u0045\u0054\u0079\u0070\u0065";ValueTypeNameRenditionClass ValueTypeName ="\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073";
ValueTypeNameResourceRef ValueTypeName ="R\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0052\u0065\u0066";ValueTypeNameURL ValueTypeName ="\u0055\u0052\u004c";ValueTypeNameURI ValueTypeName ="\u0055\u0052\u0049";ValueTypeNameVersion ValueTypeName ="\u0056e\u0072\u0073\u0069\u006f\u006e";
);

// FillModel fills in the document XMP model.
func FillModel (d *_b .Document ,extModel *Model )error {var _c []*_b .Namespace ;for _ ,_ff :=range d .Namespaces (){_c =append (_c ,_ff );};_e .Slice (_c ,func (_bg ,_fd int )bool {return _c [_bg ].Name < _c [_fd ].Name });for _ ,_ffa :=range _c {switch _ffa {case Namespace ,SchemaNS ,PropertyNS ,ValueTypeNS ,FieldNS ,_ec .NsXmp ,_ada .NsDc :continue ;
default:_ga ,_ab :=GetSchema (_ffa .URI );if !_ab {continue ;};_dc :=d .FindModel (_ffa );_fe :=*_ga ;_fe .Property =Properties {};for _ ,_cd :=range _ga .Property {_cf :=_ca (_dc ,_ffa ,_cd );if _cf {_fe .Property =append (_fe .Property ,_cd );};};if len (_fe .Property )==0{continue ;
};var _cdb bool ;for _bac ,_ge :=range extModel .Schemas {if _ge .Schema ==_fe .Schema {_cdb =true ;extModel .Schemas [_bac ]=_fe ;break ;};};if !_cdb {extModel .Schemas =append (extModel .Schemas ,_fe );};};};return nil ;};

// BagOfValueTypeName gets the ValueTypeName of the bag of provided value type names.
func BagOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0062\u0061\u0067\u0020"+vt };const (PropertyCategoryUndefined PropertyCategory =iota ;PropertyCategoryInternal ;PropertyCategoryExternal ;);

// Model is the pdfa extension metadata model.
type Model struct{Schemas Schemas `xmp:"pdfaExtension:schemas"`;};

// SyncModel implements xmp.Model interface.
func (_db *Model )SyncModel (d *_b .Document )error {return nil };

// GetTag implements xmp.Model interface.
func (_cbg *Model )GetTag (tag string )(string ,error ){_ecf ,_gb :=_b .GetNativeField (_cbg ,tag );if _gb !=nil {return "",_f .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_gb );};return _ecf ,nil ;};func (_abc Schemas )Typ ()_b .ArrayType {return _b .ArrayTypeUnordered };
var XmpIDQualSchema =Schema {NamespaceURI :"\u0068t\u0074\u0070:\u002f\u002f\u006e\u0073.\u0061\u0064\u006fb\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u006d\u0070/I\u0064\u0065\u006et\u0069\u0066i\u0065\u0072\u002f\u0071\u0075\u0061l\u002f\u0031.\u0030\u002f",Prefix :"\u0078\u006d\u0070\u0069\u0064\u0071",Schema :"X\u004dP\u0020\u0078\u006d\u0070\u0069\u0064\u0071\u0020q\u0075\u0061\u006c\u0069fi\u0065\u0072",Property :[]Property {{Category :PropertyCategoryInternal ,Name :"\u0053\u0063\u0068\u0065\u006d\u0065",Description :"\u0041\u0020\u0071\u0075\u0061\u006c\u0069\u0066\u0069\u0065\u0072\u0020\u0070\u0072o\u0076\u0069\u0064i\u006e\u0067\u0020\u0074h\u0065\u0020\u006e\u0020\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u006c \u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0065\u0020\u0075\u0073ed\u0020\u0066\u006fr\u0020\u0061\u006e\u0020\u0069\u0074\u0065\u006d \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0078\u006d\u0070\u003a\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u002e",ValueType :ValueTypeNameText }},ValueType :nil };
const (ValueTypeResourceEvent ValueTypeName ="\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065\u006e\u0074";);

// MakeModel creates or gets a model from document.
func MakeModel (d *_b .Document )(*Model ,error ){_bd ,_g :=d .MakeModel (Namespace );if _g !=nil {return nil ,_g ;};return _bd .(*Model ),nil ;};

// SetSchema sets the schema into given model.
func (_bgg *Model )SetSchema (namespaceURI string ,s Schema ){for _ag :=0;_ag < len (_bgg .Schemas );_ag ++{if _bgg .Schemas [_ag ].NamespaceURI ==namespaceURI {_bgg .Schemas [_ag ]=s ;return ;};};_bgg .Schemas =append (_bgg .Schemas ,s );};

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_dg *Properties )UnmarshalXMP (d *_b .Decoder ,node *_b .Node ,m _b .Model )error {return _b .UnmarshalArray (d ,node ,_dg .Typ (),_dg );};

// SyncToXMP implements xmp.Model interface.
func (_fa *Model )SyncToXMP (d *_b .Document )error {return nil };

// ClosedChoiceValueTypeName gets the closed choice of provided value type name.
func ClosedChoiceValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0043\u006c\u006f\u0073\u0065\u0064\u0020\u0043\u0068\u006f\u0069\u0063e\u0020\u006f\u0066\u0020"+vt ;};

// ValueTypes is the slice of field value types.
type ValueTypes []FieldValueType ;

// Schema is the pdfa extension schema.
type Schema struct{

// NamespaceURI is schema namespace URI.
NamespaceURI string `xmp:"pdfaSchema:namespaceURI"`;

// Prefix is preferred schema namespace prefix.
Prefix string `xmp:"pdfaSchema:prefix"`;

// Schema is optional description.
Schema string `xmp:"pdfaSchema:schema"`;

// Property is description of schema properties.
Property Properties `xmp:"pdfaSchema:property"`;

// ValueType is description of schema-specific value types.
ValueType ValueTypes `xmp:"pdfaSchema:valueType"`;};

// Typ gets the array type of properties.
func (_af Properties )Typ ()_b .ArrayType {return _b .ArrayTypeOrdered };

// MarshalXMP implements xmp.Marshaler interface.
func (_gef Schemas )MarshalXMP (e *_b .Encoder ,node *_b .Node ,m _b .Model )error {return _b .MarshalArray (e ,node ,_gef .Typ (),_gef );};

// Properties is a list of properties.
type Properties []Property ;

// Namespaces implements xmp.Model interface.
func (_cb *Model )Namespaces ()_b .NamespaceList {return _b .NamespaceList {Namespace }};

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_eb *ValueTypes )UnmarshalXMP (d *_b .Decoder ,node *_b .Node ,m _b .Model )error {return _b .UnmarshalArray (d ,node ,_eb .Typ (),_eb );};

// SeqOfValueTypeName gets a value type name of a sequence of input value type names.
func SeqOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0073\u0065\u0071\u0020"+vt };

// Typ gets array type of the field value types.
func (_eaf ValueTypes )Typ ()_b .ArrayType {return _b .ArrayTypeOrdered };

// AltOfValueTypeName gets the ValueTypeName of the alt of given value type names.
func AltOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0061\u006c\u0074\u0020"+vt };

// FieldValueType is a schema that describes a field in a structured type.
type FieldValueType struct{Name string `xmp:"pdfaField:description"`;Description string `xmp:"pdfaField:name"`;ValueType ValueTypeName `xmp:"pdfaField:valueType"`;};

// ValueTypeName is the name of the value type.
type ValueTypeName string ;

// MarshalXMP implements xmp.Marshaler interface.
func (_bf ValueTypes )MarshalXMP (e *_b .Encoder ,node *_b .Node ,m _b .Model )error {return _b .MarshalArray (e ,node ,_bf .Typ (),_bf );};

// CanTag implements xmp.Model interface.
func (_dea *Model )CanTag (tag string )bool {_ ,_agb :=_b .GetNativeField (_dea ,tag );return _agb ==nil ;};var FieldVersionSchema =Schema {NamespaceURI :"\u0068\u0074\u0074p\u003a\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065/\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0023",Prefix :"\u0073\u0074\u0056e\u0072",Schema :"\u0042a\u0073\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074y\u0070\u0065\u0020\u0056\u0065\u0072\u0073\u0069\u006f\u006e",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0043\u006fmm\u0065\u006e\u0074s\u0020\u0063\u006f\u006ecer\u006ein\u0067\u0020\u0077\u0068\u0061\u0074\u0020wa\u0073\u0020\u0063\u0068\u0061\u006e\u0067e\u0064",Name :"\u0063\u006f\u006d\u006d\u0065\u006e\u0074\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0048\u0069\u0067\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u002c\u0020\u0066\u006f\u0072\u006d\u0061\u006c\u0020\u0064e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0077\u0068\u0061\u0074\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0068\u0065\u0020\u0075\u0073\u0065\u0072 \u0070\u0065\u0072f\u006f\u0072\u006d\u0065\u0064\u002e",Name :"\u0065\u0076\u0065n\u0074",ValueType :ValueTypeResourceEvent },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068e\u0020\u0064\u0061\u0074\u0065\u0020\u006f\u006e\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0065\u0063\u006b\u0065\u0064\u0020\u0069\u006e\u002e",Name :"\u006d\u006f\u0064\u0069\u0066\u0079\u0044\u0061\u0074\u0065",ValueType :ValueTypeNameDate },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068e\u0020\u0070\u0065\u0072s\u006f\u006e \u0077\u0068\u006f\u0020\u006d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u002e",Name :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0072",ValueType :ValueTypeNameProperName },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 n\u0065\u0077\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e",Name :"\u0076e\u0072\u0073\u0069\u006f\u006e",ValueType :ValueTypeNameText }},ValueType :nil };


// Schemas is the array of xmp metadata extension resources.
type Schemas []Schema ;