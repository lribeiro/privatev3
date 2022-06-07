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

// Package common contains common properties used by the subpackages.
package common ;import (_e "fmt";_g "io";_f "os";_be "path/filepath";_b "runtime";_cg "time";);

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Trace logs trace message.
func (_ffd WriterLogger )Trace (format string ,args ...interface{}){if _ffd .LogLevel >=LogLevelTrace {_gc :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_ffd .logToWriter (_ffd .Output ,_gc ,format ,args ...);};};const _acea ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";


// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };func (_fed WriterLogger )logToWriter (_ecb _g .Writer ,_eda string ,_ade string ,_aee ...interface{}){_cca (_ecb ,_eda ,_ade ,_aee );};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _g .Writer )*WriterLogger {_ggf :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_ggf ;};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };func _cca (_cgd _g .Writer ,_bb string ,_fg string ,_ecc ...interface{}){_ ,_dea ,_ebab ,_bec :=_b .Caller (3);if !_bec {_dea ="\u003f\u003f\u003f";_ebab =0;}else {_dea =_be .Base (_dea );};_bf :=_e .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_bb ,_dea ,_ebab )+_fg +"\u000a";
_e .Fprintf (_cgd ,_bf ,_ecc ...);};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _g .Writer ;};

// LogLevel is the verbosity level for logging.
type LogLevel int ;const Version ="\u0033\u002e\u0033\u0035\u002e\u0030";

// Info logs info message.
func (_cfe ConsoleLogger )Info (format string ,args ...interface{}){if _cfe .LogLevel >=LogLevelInfo {_de :="\u005bI\u004e\u0046\u004f\u005d\u0020";_cfe .output (_f .Stdout ,_de ,format ,args ...);};};const _fbc =6;

// Debug logs debug message.
func (_ebe WriterLogger )Debug (format string ,args ...interface{}){if _ebe .LogLevel >=LogLevelDebug {_dd :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ebe .logToWriter (_ebe .Output ,_dd ,format ,args ...);};};

// Warning logs warning message.
func (_gba WriterLogger )Warning (format string ,args ...interface{}){if _gba .LogLevel >=LogLevelWarning {_bc :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_gba .logToWriter (_gba .Output ,_bc ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Error logs error message.
func (_ed WriterLogger )Error (format string ,args ...interface{}){if _ed .LogLevel >=LogLevelError {_df :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ed .logToWriter (_ed .Output ,_df ,format ,args ...);};};var ReleasedAt =_cg .Date (_ccc ,_fbc ,_cgg ,_ee ,_ebf ,0,0,_cg .UTC );


// Debug logs debug message.
func (_age ConsoleLogger )Debug (format string ,args ...interface{}){if _age .LogLevel >=LogLevelDebug {_ad :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_age .output (_f .Stdout ,_ad ,format ,args ...);};};const _ccc =2022;

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_caf ConsoleLogger )IsLogLevel (level LogLevel )bool {return _caf .LogLevel >=level };

// Notice logs notice message.
func (_fb WriterLogger )Notice (format string ,args ...interface{}){if _fb .LogLevel >=LogLevelNotice {_dfe :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_fb .logToWriter (_fb .Output ,_dfe ,format ,args ...);};};

// Notice logs notice message.
func (_bg ConsoleLogger )Notice (format string ,args ...interface{}){if _bg .LogLevel >=LogLevelNotice {_fe :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_bg .output (_f .Stdout ,_fe ,format ,args ...);};};

// Trace logs trace message.
func (_fec ConsoleLogger )Trace (format string ,args ...interface{}){if _fec .LogLevel >=LogLevelTrace {_ae :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_fec .output (_f .Stdout ,_ae ,format ,args ...);};};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Error logs error message.
func (_ec ConsoleLogger )Error (format string ,args ...interface{}){if _ec .LogLevel >=LogLevelError {_ggd :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ec .output (_f .Stdout ,_ggd ,format ,args ...);};};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Info logs info message.
func (_db WriterLogger )Info (format string ,args ...interface{}){if _db .LogLevel >=LogLevelInfo {_cb :="\u005bI\u004e\u0046\u004f\u005d\u0020";_db .logToWriter (_db .Output ,_cb ,format ,args ...);};};

// DummyLogger does nothing.
type DummyLogger struct{};var Log Logger =DummyLogger {};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};const _cgg =6;const _ee =15;func (_ace ConsoleLogger )output (_adg _g .Writer ,_cc string ,_ff string ,_bd ...interface{}){_cca (_adg ,_cc ,_ff ,_bd ...);};const _ebf =30;

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_cf string ,_a ...interface{});Warning (_gg string ,_ca ...interface{});Notice (_eb string ,_ge ...interface{});Info (_fc string ,_eba ...interface{});Debug (_d string ,_ac ...interface{});Trace (_dc string ,_gd ...interface{});
IsLogLevel (_ag LogLevel )bool ;};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fa WriterLogger )IsLogLevel (level LogLevel )bool {return _fa .LogLevel >=level };

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _cg .Time )string {return t .Format (_acea )+"\u0020\u0055\u0054\u0043"};

// Warning logs warning message.
func (_gb ConsoleLogger )Warning (format string ,args ...interface{}){if _gb .LogLevel >=LogLevelWarning {_fd :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_gb .output (_f .Stdout ,_fd ,format ,args ...);};};