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

package endian ;import (_e "encoding/binary";_b "unsafe";);func IsBig ()bool {return _eb };func init (){const _dc =int (_b .Sizeof (0));_a :=1;_gc :=(*[_dc ]byte )(_b .Pointer (&_a ));if _gc [0]==0{_eb =true ;ByteOrder =_e .BigEndian ;}else {ByteOrder =_e .LittleEndian ;
};};var (ByteOrder _e .ByteOrder ;_eb bool ;);func IsLittle ()bool {return !_eb };