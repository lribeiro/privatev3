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

// Package xmputil provides abstraction used by the pdf document XMP Metadata.
package xmputil ;import (_d "errors";_ab "fmt";_ag "github.com/trimmer-io/go-xmp/models/pdf";_gg "github.com/trimmer-io/go-xmp/models/xmp_mm";_g "github.com/trimmer-io/go-xmp/xmp";_b "github.com/unidoc/unipdf/v3/core";_fb "github.com/unidoc/unipdf/v3/internal/timeutils";
_cb "github.com/unidoc/unipdf/v3/internal/uuid";_dc "github.com/unidoc/unipdf/v3/model/xmputil/pdfaextension";_agd "github.com/unidoc/unipdf/v3/model/xmputil/pdfaid";_c "strconv";_a "time";);

// Marshal the document into xml byte stream.
func (_cf *Document )Marshal ()([]byte ,error ){if _cf ._ge .IsDirty (){if _be :=_cf ._ge .SyncModels ();_be !=nil {return nil ,_be ;};};return _g .Marshal (_cf ._ge );};

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct{Part int ;Conformance string ;};

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct{OriginalDocumentID GUID ;DocumentID GUID ;InstanceID GUID ;VersionID string ;};

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_fa *Document )GetMediaManagement ()(*MediaManagement ,bool ){_ce :=_gg .FindModel (_fa ._ge );if _ce ==nil {return nil ,false ;};_ede :=make ([]MediaManagementVersion ,len (_ce .Versions ));for _ada ,_gea :=range _ce .Versions {_ede [_ada ]=MediaManagementVersion {VersionID :_gea .Version ,ModifyDate :_gea .ModifyDate .Value (),Comments :_gea .Comments ,Modifier :_gea .Modifier };
};_bgbf :=&MediaManagement {OriginalDocumentID :GUID (_ce .OriginalDocumentID .Value ()),DocumentID :GUID (_ce .DocumentID .Value ()),InstanceID :GUID (_ce .InstanceID .Value ()),VersionID :_ce .VersionID ,Versions :_ede };if _ce .DerivedFrom !=nil {_bgbf .DerivedFrom =&MediaManagementDerivedFrom {OriginalDocumentID :GUID (_ce .DerivedFrom .OriginalDocumentID ),DocumentID :GUID (_ce .DerivedFrom .DocumentID ),InstanceID :GUID (_ce .DerivedFrom .InstanceID ),VersionID :_ce .DerivedFrom .VersionID };
};return _bgbf ,true ;};

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_ff *Document )SetPdfAExtension ()error {_dag ,_df :=_dc .MakeModel (_ff ._ge );if _df !=nil {return _df ;};if _df =_dc .FillModel (_ff ._ge ,_dag );_df !=nil {return _df ;};if _df =_dag .SyncToXMP (_ff ._ge );_df !=nil {return _df ;};return nil ;
};

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_agdg *Document )GetPdfaExtensionSchemas ()([]_dc .Schema ,error ){_cd :=_agdg ._ge .FindModel (_dc .Namespace );if _cd ==nil {return nil ,nil ;};_eb ,_bg :=_cd .(*_dc .Model );if !_bg {return nil ,_ab .Errorf ("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054",_cd );
};return _eb .Schemas ,nil ;};

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_e *Document )MarshalIndent (prefix ,indent string )([]byte ,error ){if _e ._ge .IsDirty (){if _agg :=_e ._ge .SyncModels ();_agg !=nil {return nil ,_agg ;};};return _g .MarshalIndent (_e ._ge ,prefix ,indent );};

// SetPdfInfo sets the pdf info into selected document.
func (_cde *Document )SetPdfInfo (options *PdfInfoOptions )error {if options ==nil {return _d .New ("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};_ca ,_ae :=_ag .MakeModel (_cde ._ge );
if _ae !=nil {return _ae ;};if options .Overwrite {*_ca =_ag .PDFInfo {};};if options .InfoDict !=nil {_fc ,_ec :=_b .GetDict (options .InfoDict );if !_ec {return _ab .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",options .InfoDict );
};var _dd *_b .PdfObjectString ;for _ ,_ee :=range _fc .Keys (){switch _ee {case "\u0054\u0069\u0074l\u0065":_dd ,_ec =_b .GetString (_fc .Get ("\u0054\u0069\u0074l\u0065"));if _ec {_ca .Title =_g .NewAltString (_dd );};case "\u0041\u0075\u0074\u0068\u006f\u0072":_dd ,_ec =_b .GetString (_fc .Get ("\u0041\u0075\u0074\u0068\u006f\u0072"));
if _ec {_ca .Author =_g .NewStringList (_dd .String ());};case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":_dd ,_ec =_b .GetString (_fc .Get ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"));if _ec {_ca .Keywords =_dd .String ();};case "\u0043r\u0065\u0061\u0074\u006f\u0072":_dd ,_ec =_b .GetString (_fc .Get ("\u0043r\u0065\u0061\u0074\u006f\u0072"));
if _ec {_ca .Creator =_g .AgentName (_dd .String ());};case "\u0053u\u0062\u006a\u0065\u0063\u0074":_dd ,_ec =_b .GetString (_fc .Get ("\u0053u\u0062\u006a\u0065\u0063\u0074"));if _ec {_ca .Subject =_g .NewAltString (_dd .String ());};case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":_dd ,_ec =_b .GetString (_fc .Get ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"));
if _ec {_ca .Producer =_g .AgentName (_dd .String ());};case "\u0054r\u0061\u0070\u0070\u0065\u0064":_bc ,_fg :=_b .GetName (_fc .Get ("\u0054r\u0061\u0070\u0070\u0065\u0064"));if _fg {switch _bc .String (){case "\u0054\u0072\u0075\u0065":_ca .Trapped =true ;
case "\u0046\u0061\u006cs\u0065":_ca .Trapped =false ;default:_ca .Trapped =true ;};};case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":if _ccb ,_bgd :=_b .GetString (_fc .Get ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"));
_bgd &&_ccb .String ()!=""{_ea ,_ac :=_fb .ParsePdfTime (_ccb .String ());if _ac !=nil {return _ab .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077",_ac );
};_ca .CreationDate =_g .NewDate (_ea );};case "\u004do\u0064\u0044\u0061\u0074\u0065":if _gb ,_age :=_b .GetString (_fc .Get ("\u004do\u0064\u0044\u0061\u0074\u0065"));_age &&_gb .String ()!=""{_gd ,_abb :=_fb .ParsePdfTime (_gb .String ());if _abb !=nil {return _ab .Errorf ("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w",_abb );
};_ca .ModifyDate =_g .NewDate (_gd );};};};};if options .PdfVersion !=""{_ca .PDFVersion =options .PdfVersion ;};if options .Marked {_ca .Marked =_g .Bool (options .Marked );};if options .Copyright !=""{_ca .Copyright =options .Copyright ;};if _ae =_ca .SyncToXMP (_cde ._ge );
_ae !=nil {return _ae ;};return nil ;};

// GetPdfInfo gets the document pdf info.
func (_agf *Document )GetPdfInfo ()(*PdfInfo ,bool ){_daf ,_bee :=_agf ._ge .FindModel (_ag .NsPDF ).(*_ag .PDFInfo );if !_bee {return nil ,false ;};_eec :=PdfInfo {};var _bce *_b .PdfObjectDictionary ;_eec .Copyright =_daf .Copyright ;_eec .PdfVersion =_daf .PDFVersion ;
_eec .Marked =bool (_daf .Marked );_gdf :=func (_bec string ,_ga _b .PdfObject ){if _bce ==nil {_bce =_b .MakeDict ();};_bce .Set (_b .PdfObjectName (_bec ),_ga );};if len (_daf .Title )> 0{_gdf ("\u0054\u0069\u0074l\u0065",_b .MakeString (_daf .Title .Default ()));
};if len (_daf .Author )> 0{_gdf ("\u0041\u0075\u0074\u0068\u006f\u0072",_b .MakeString (_daf .Author [0]));};if _daf .Keywords !=""{_gdf ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",_b .MakeString (_daf .Keywords ));};if len (_daf .Subject )> 0{_gdf ("\u0053u\u0062\u006a\u0065\u0063\u0074",_b .MakeString (_daf .Subject .Default ()));
};if _daf .Creator !=""{_gdf ("\u0043r\u0065\u0061\u0074\u006f\u0072",_b .MakeString (string (_daf .Creator )));};if _daf .Producer !=""{_gdf ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",_b .MakeString (string (_daf .Producer )));};if _daf .Trapped {_gdf ("\u0054r\u0061\u0070\u0070\u0065\u0064",_b .MakeName ("\u0054\u0072\u0075\u0065"));
};if !_daf .CreationDate .IsZero (){_gdf ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_b .MakeString (_fb .FormatPdfTime (_daf .CreationDate .Value ())));};if !_daf .ModifyDate .IsZero (){_gdf ("\u004do\u0064\u0044\u0061\u0074\u0065",_b .MakeString (_fb .FormatPdfTime (_daf .ModifyDate .Value ())));
};_eec .InfoDict =_bce ;return &_eec ,true ;};

// GUID is a string representing a globally unique identifier.
type GUID string ;

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument (stream []byte )(*Document ,error ){_dg :=_g .NewDocument ();if _da :=_g .Unmarshal (stream ,_dg );_da !=nil {return nil ,_da ;};return &Document {_ge :_dg },nil ;};

// PdfInfo is the xmp document pdf info.
type PdfInfo struct{InfoDict _b .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;};

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct{VersionID string ;ModifyDate _a .Time ;Comments string ;Modifier string ;};

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct{InfoDict _b .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;

// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
Overwrite bool ;};

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_bd *Document )SetMediaManagement (options *MediaManagementOptions )error {_bge ,_ed :=_gg .MakeModel (_bd ._ge );if _ed !=nil {return _ed ;};if options ==nil {options =new (MediaManagementOptions );};_cad :=_gg .ResourceRef {};if _bge .OriginalDocumentID .IsZero (){if options .OriginalDocumentID !=""{_bge .OriginalDocumentID =_g .GUID (options .OriginalDocumentID );
}else {_cdea ,_bgb :=_cb .NewUUID ();if _bgb !=nil {return _bgb ;};_bge .OriginalDocumentID =_g .GUID (_cdea .String ());};}else {_cad .OriginalDocumentID =_bge .OriginalDocumentID ;};switch {case options .DocumentID !="":_bge .DocumentID =_g .GUID (options .DocumentID );
case options .NewDocumentID ||_bge .DocumentID .IsZero ():if !_bge .DocumentID .IsZero (){_cad .DocumentID =_bge .DocumentID ;};_cab ,_abf :=_cb .NewUUID ();if _abf !=nil {return _abf ;};_bge .DocumentID =_g .GUID (_cab .String ());};if !_bge .InstanceID .IsZero (){_cad .InstanceID =_bge .InstanceID ;
};_bge .InstanceID =_g .GUID (options .InstanceID );if _bge .InstanceID ==""{_caa ,_ad :=_cb .NewUUID ();if _ad !=nil {return _ad ;};_bge .InstanceID =_g .GUID (_caa .String ());};if !_cad .IsZero (){_bge .DerivedFrom =&_cad ;};_fca :=options .VersionID ;
if _bge .VersionID !=""{_eeb ,_edd :=_c .Atoi (_bge .VersionID );if _edd !=nil {_fca =_c .Itoa (len (_bge .Versions )+1);}else {_fca =_c .Itoa (_eeb +1);};};if _fca ==""{_fca ="\u0031";};_bge .VersionID =_fca ;_ecg :=options .ModifyDate ;if _ecg .IsZero (){_ecg =_a .Now ();
};if _ed =_bge .SyncToXMP (_bd ._ge );_ed !=nil {return _ed ;};return nil ;};

// NewDocument creates a new document without any previous xmp information.
func NewDocument ()*Document {_cc :=_g .NewDocument ();return &Document {_ge :_cc }};

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{_ge *_g .Document };

// MediaManagement are the values from the document media management metadata.
type MediaManagement struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
OriginalDocumentID GUID ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
DocumentID GUID ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
InstanceID GUID ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
DerivedFrom *MediaManagementDerivedFrom ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
VersionID string ;

// Versions is the history of the document versions along with the comments, timestamps and issuers.
Versions []MediaManagementVersion ;};

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_af *Document )GetGoXmpDocument ()*_g .Document {return _af ._ge };

// GetPdfAID gets the pdfaid xmp metadata model.
func (_beea *Document )GetPdfAID ()(*PdfAID ,bool ){_dcd ,_eecg :=_beea ._ge .FindModel (_agd .Namespace ).(*_agd .Model );if !_eecg {return nil ,false ;};return &PdfAID {Part :_dcd .Part ,Conformance :_dcd .Conformance },true ;};

// MediaManagementOptions are the options for the Media management xmp metadata.
type MediaManagementOptions struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
// By default, this value is generated.
OriginalDocumentID string ;

// NewDocumentID is a flag which generates a new Document identifier while setting media management.
// This value should be set to true only if the document is stored and saved as new document.
// Otherwise, if the document is modified and overwrites previous file, it should be set to false.
NewDocumentID bool ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
// By default, this value is generated if NewDocumentID is true or previous doesn't exist.
DocumentID string ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
// By default, this value is generated.
InstanceID string ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
// By default, the derived from structure is filled from previous XMP metadata (if exists).
DerivedFrom string ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
// By default, this values is incremented or set to the next version number.
VersionID string ;

// ModifyComment is a comment to given modification
ModifyComment string ;

// ModifyDate is a custom modification date for the versions.
// By default, this would be set to time.Now().
ModifyDate _a .Time ;

// Modifier is a person who did the modification.
Modifier string ;};

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_cg *Document )SetPdfAID (part int ,conformance string )error {_fe ,_fd :=_agd .MakeModel (_cg ._ge );if _fd !=nil {return _fd ;};_fe .Part =part ;_fe .Conformance =conformance ;if _dfe :=_fe .SyncToXMP (_cg ._ge );_dfe !=nil {return _dfe ;};return nil ;
};