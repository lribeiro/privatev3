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

package endian ;import (_ff "encoding/binary";_ffc "unsafe";);func IsLittle ()bool {return !_a };func init (){const _ffce =int (_ffc .Sizeof (0));_fe :=1;_g :=(*[_ffce ]byte )(_ffc .Pointer (&_fe ));if _g [0]==0{_a =true ;ByteOrder =_ff .BigEndian ;}else {ByteOrder =_ff .LittleEndian ;
};};var (ByteOrder _ff .ByteOrder ;_a bool ;);func IsBig ()bool {return _a };