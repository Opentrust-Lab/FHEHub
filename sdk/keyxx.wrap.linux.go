package sdk

import "C"

import (
	"math"
	"encoding/binary"
	"github.com/ebitengine/purego"
)

type Linux_SDK struct {
	UserKey string
	libPath string
	keyxx uintptr
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	println(bits)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(bits))
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	for i := len(bytes); i < 8; i++ {
		bytes = append(bytes, 0)
	}
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

func (sdk *Linux_SDK) getCIntPtr(val int) C.int {
	return C.int(val)
}

func (sdk *Linux_SDK) getCStrPtr(val string) *C.char {
	return C.CString(val)
}

func (sdk *Linux_SDK) getCFloat64Ptr(val float64) *C.char {
	return sdk.getCStrPtr(string(Float64ToByte(val)))
}

func (sdk *Linux_SDK) getReturnString(ptr *C.char) string {
	return C.GoString(ptr)
}

func (sdk *Linux_SDK) getReturnFloat64(ptr *C.char) float64 {
    return ByteToFloat64([]byte(sdk.getReturnString(ptr)))
}

func (sdk *Linux_SDK) SetLibPath(libPath_in string) {
	sdk.libPath = libPath_in
}

func (sdk *Linux_SDK) InitLib() {
	var err error
	if sdk.libPath == "" {
		sdk.libPath = "libkeyxx.core.so"
	}
	sdk.keyxx, err = purego.Dlopen(sdk.libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic("LoadLibrary failed: " + err.Error() + "\n")
	} else {
		println("LoadLibrary Success...\n")
	}
}

func (sdk *Linux_SDK) Release() {
	purego.Dlclose(sdk.keyxx)
}

func (sdk *Linux_SDK) Init(m_in int, n_in int, q_in int, p_in int)  {
	var func_Init func(C.int, C.int, C.int, C.int) 
	purego.RegisterLibFunc(&func_Init, sdk.keyxx, "Init")
	p0 := sdk.getCIntPtr(m_in)
	p1 := sdk.getCIntPtr(n_in)
	p2 := sdk.getCIntPtr(q_in)
	p3 := sdk.getCIntPtr(p_in)
	func_Init(p0, p1, p2, p3)
}

func (sdk *Linux_SDK) InitCodeType(ct int)  {
	var func_InitCodeType func(C.int) 
	purego.RegisterLibFunc(&func_InitCodeType, sdk.keyxx, "InitCodeType")
	p0 := sdk.getCIntPtr(ct)
	func_InitCodeType(p0)
}

func (sdk *Linux_SDK) InitSerializeType(st int)  {
	var func_InitSerializeType func(C.int) 
	purego.RegisterLibFunc(&func_InitSerializeType, sdk.keyxx, "InitSerializeType")
	p0 := sdk.getCIntPtr(st)
	func_InitSerializeType(p0)
}

func (sdk *Linux_SDK) InitStrCaseChangeMode(scc int)  {
	var func_InitStrCaseChangeMode func(C.int) 
	purego.RegisterLibFunc(&func_InitStrCaseChangeMode, sdk.keyxx, "InitStrCaseChangeMode")
	p0 := sdk.getCIntPtr(scc)
	func_InitStrCaseChangeMode(p0)
}

func (sdk *Linux_SDK) InitIsFastString(iss int)  {
	var func_InitIsFastString func(C.int) 
	purego.RegisterLibFunc(&func_InitIsFastString, sdk.keyxx, "InitIsFastString")
	p0 := sdk.getCIntPtr(iss)
	func_InitIsFastString(p0)
}

func (sdk *Linux_SDK) SM3(input string) string {
	var func_SM3 func(*C.char) *C.char
	purego.RegisterLibFunc(&func_SM3, sdk.keyxx, "SM3")
	p0 := sdk.getCStrPtr(input)
	r := func_SM3(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) LoadPrivKey(skbfile string)  {
	var func_LoadPrivKey func(*C.char) 
	purego.RegisterLibFunc(&func_LoadPrivKey, sdk.keyxx, "LoadPrivKey")
	p0 := sdk.getCStrPtr(skbfile)
	func_LoadPrivKey(p0)
}

func (sdk *Linux_SDK) LoadPubKey(pkbfile string)  {
	var func_LoadPubKey func(*C.char) 
	purego.RegisterLibFunc(&func_LoadPubKey, sdk.keyxx, "LoadPubKey")
	p0 := sdk.getCStrPtr(pkbfile)
	func_LoadPubKey(p0)
}

func (sdk *Linux_SDK) LoadDictionary(dictbfile string)  {
	var func_LoadDictionary func(*C.char) 
	purego.RegisterLibFunc(&func_LoadDictionary, sdk.keyxx, "LoadDictionary")
	p0 := sdk.getCStrPtr(dictbfile)
	func_LoadDictionary(p0)
}

func (sdk *Linux_SDK) LoadPrivKeyString(skbstring string)  {
	var func_LoadPrivKeyString func(*C.char) 
	purego.RegisterLibFunc(&func_LoadPrivKeyString, sdk.keyxx, "LoadPrivKeyString")
	p0 := sdk.getCStrPtr(skbstring)
	func_LoadPrivKeyString(p0)
}

func (sdk *Linux_SDK) LoadPubKeyString(pkbstring string)  {
	var func_LoadPubKeyString func(*C.char) 
	purego.RegisterLibFunc(&func_LoadPubKeyString, sdk.keyxx, "LoadPubKeyString")
	p0 := sdk.getCStrPtr(pkbstring)
	func_LoadPubKeyString(p0)
}

func (sdk *Linux_SDK) LoadDictionaryString(dictbstring string)  {
	var func_LoadDictionaryString func(*C.char) 
	purego.RegisterLibFunc(&func_LoadDictionaryString, sdk.keyxx, "LoadDictionaryString")
	p0 := sdk.getCStrPtr(dictbstring)
	func_LoadDictionaryString(p0)
}

func (sdk *Linux_SDK) LoadBiasKey(bkfile string)  {
	var func_LoadBiasKey func(*C.char) 
	purego.RegisterLibFunc(&func_LoadBiasKey, sdk.keyxx, "LoadBiasKey")
	p0 := sdk.getCStrPtr(bkfile)
	func_LoadBiasKey(p0)
}

func (sdk *Linux_SDK) LoadExKey(dictExbfile string)  {
	var func_LoadExKey func(*C.char) 
	purego.RegisterLibFunc(&func_LoadExKey, sdk.keyxx, "LoadExKey")
	p0 := sdk.getCStrPtr(dictExbfile)
	func_LoadExKey(p0)
}

func (sdk *Linux_SDK) LoadPrivkeyFloat(skffile string)  {
	var func_LoadPrivkeyFloat func(*C.char) 
	purego.RegisterLibFunc(&func_LoadPrivkeyFloat, sdk.keyxx, "LoadPrivkeyFloat")
	p0 := sdk.getCStrPtr(skffile)
	func_LoadPrivkeyFloat(p0)
}

func (sdk *Linux_SDK) LoadDictionaryFloat(dictffile string)  {
	var func_LoadDictionaryFloat func(*C.char) 
	purego.RegisterLibFunc(&func_LoadDictionaryFloat, sdk.keyxx, "LoadDictionaryFloat")
	p0 := sdk.getCStrPtr(dictffile)
	func_LoadDictionaryFloat(p0)
}

func (sdk *Linux_SDK) LoadExpressionBlockKeys()  {
	var func_LoadExpressionBlockKeys func() 
	purego.RegisterLibFunc(&func_LoadExpressionBlockKeys, sdk.keyxx, "LoadExpressionBlockKeys")
	func_LoadExpressionBlockKeys()
}

func (sdk *Linux_SDK) LoadExpressionFloatKeys()  {
	var func_LoadExpressionFloatKeys func() 
	purego.RegisterLibFunc(&func_LoadExpressionFloatKeys, sdk.keyxx, "LoadExpressionFloatKeys")
	func_LoadExpressionFloatKeys()
}

func (sdk *Linux_SDK) InitExpression()  {
	var func_InitExpression func() 
	purego.RegisterLibFunc(&func_InitExpression, sdk.keyxx, "InitExpression")
	func_InitExpression()
}

func (sdk *Linux_SDK) ClearExpressionVariable()  {
	var func_ClearExpressionVariable func() 
	purego.RegisterLibFunc(&func_ClearExpressionVariable, sdk.keyxx, "ClearExpressionVariable")
	func_ClearExpressionVariable()
}

func (sdk *Linux_SDK) AddExpressionVariable(_name string, _type string, _value string)  {
	var func_AddExpressionVariable func(*C.char, *C.char, *C.char) 
	purego.RegisterLibFunc(&func_AddExpressionVariable, sdk.keyxx, "AddExpressionVariable")
	p0 := sdk.getCStrPtr(_name)
	p1 := sdk.getCStrPtr(_type)
	p2 := sdk.getCStrPtr(_value)
	func_AddExpressionVariable(p0, p1, p2)
}

func (sdk *Linux_SDK) ExpressionCalculation(formula string) string {
	var func_ExpressionCalculation func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ExpressionCalculation, sdk.keyxx, "ExpressionCalculation")
	p0 := sdk.getCStrPtr(formula)
	r := func_ExpressionCalculation(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) GenSKB(m int, n int, q int, p int, filename string)  {
	var func_GenSKB func(C.int, C.int, C.int, C.int, *C.char) 
	purego.RegisterLibFunc(&func_GenSKB, sdk.keyxx, "GenSKB")
	p0 := sdk.getCIntPtr(m)
	p1 := sdk.getCIntPtr(n)
	p2 := sdk.getCIntPtr(q)
	p3 := sdk.getCIntPtr(p)
	p4 := sdk.getCStrPtr(filename)
	func_GenSKB(p0, p1, p2, p3, p4)
}

func (sdk *Linux_SDK) GenPKB(skbfile string, filename string)  {
	var func_GenPKB func(*C.char, *C.char) 
	purego.RegisterLibFunc(&func_GenPKB, sdk.keyxx, "GenPKB")
	p0 := sdk.getCStrPtr(skbfile)
	p1 := sdk.getCStrPtr(filename)
	func_GenPKB(p0, p1)
}

func (sdk *Linux_SDK) GenDictB(skbfile string, filename string, delta float64)  {
	var func_GenDictB func(*C.char, *C.char, *C.char) 
	purego.RegisterLibFunc(&func_GenDictB, sdk.keyxx, "GenDictB")
	p0 := sdk.getCStrPtr(skbfile)
	p1 := sdk.getCStrPtr(filename)
	p2 := sdk.getCFloat64Ptr(delta)
	func_GenDictB(p0, p1, p2)
}

func (sdk *Linux_SDK) GenSKBString(m int, n int, q int, p int) string {
	var func_GenSKBString func(C.int, C.int, C.int, C.int) *C.char
	purego.RegisterLibFunc(&func_GenSKBString, sdk.keyxx, "GenSKBString")
	p0 := sdk.getCIntPtr(m)
	p1 := sdk.getCIntPtr(n)
	p2 := sdk.getCIntPtr(q)
	p3 := sdk.getCIntPtr(p)
	r := func_GenSKBString(p0, p1, p2, p3)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) GenPKBString(skbfile string) string {
	var func_GenPKBString func(*C.char) *C.char
	purego.RegisterLibFunc(&func_GenPKBString, sdk.keyxx, "GenPKBString")
	p0 := sdk.getCStrPtr(skbfile)
	r := func_GenPKBString(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) GenDictBString(skbfile string, delta float64) string {
	var func_GenDictBString func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_GenDictBString, sdk.keyxx, "GenDictBString")
	p0 := sdk.getCStrPtr(skbfile)
	p1 := sdk.getCFloat64Ptr(delta)
	r := func_GenDictBString(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) GenBK(n int, q int, p int, filename string)  {
	var func_GenBK func(C.int, C.int, C.int, *C.char) 
	purego.RegisterLibFunc(&func_GenBK, sdk.keyxx, "GenBK")
	p0 := sdk.getCIntPtr(n)
	p1 := sdk.getCIntPtr(q)
	p2 := sdk.getCIntPtr(p)
	p3 := sdk.getCStrPtr(filename)
	func_GenBK(p0, p1, p2, p3)
}

func (sdk *Linux_SDK) GenDictExB(bkfile string, pkbfile string, dictbfile string, filename string)  {
	var func_GenDictExB func(*C.char, *C.char, *C.char, *C.char) 
	purego.RegisterLibFunc(&func_GenDictExB, sdk.keyxx, "GenDictExB")
	p0 := sdk.getCStrPtr(bkfile)
	p1 := sdk.getCStrPtr(pkbfile)
	p2 := sdk.getCStrPtr(dictbfile)
	p3 := sdk.getCStrPtr(filename)
	func_GenDictExB(p0, p1, p2, p3)
}

func (sdk *Linux_SDK) GenSKF(p int, filename string)  {
	var func_GenSKF func(C.int, *C.char) 
	purego.RegisterLibFunc(&func_GenSKF, sdk.keyxx, "GenSKF")
	p0 := sdk.getCIntPtr(p)
	p1 := sdk.getCStrPtr(filename)
	func_GenSKF(p0, p1)
}

func (sdk *Linux_SDK) GenDictF(skffile string, filename string)  {
	var func_GenDictF func(*C.char, *C.char) 
	purego.RegisterLibFunc(&func_GenDictF, sdk.keyxx, "GenDictF")
	p0 := sdk.getCStrPtr(skffile)
	p1 := sdk.getCStrPtr(filename)
	func_GenDictF(p0, p1)
}

func (sdk *Linux_SDK) EncInt(plain int) string {
	var func_EncInt func(C.int) *C.char
	purego.RegisterLibFunc(&func_EncInt, sdk.keyxx, "EncInt")
	p0 := sdk.getCIntPtr(plain)
	r := func_EncInt(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncPubInt(plain int) string {
	var func_EncPubInt func(C.int) *C.char
	purego.RegisterLibFunc(&func_EncPubInt, sdk.keyxx, "EncPubInt")
	p0 := sdk.getCIntPtr(plain)
	r := func_EncPubInt(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncString(plain string) string {
	var func_EncString func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncString, sdk.keyxx, "EncString")
	p0 := sdk.getCStrPtr(plain)
	r := func_EncString(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncPubString(plain string) string {
	var func_EncPubString func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncPubString, sdk.keyxx, "EncPubString")
	p0 := sdk.getCStrPtr(plain)
	r := func_EncPubString(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncPubStringWithIndex(plain string) string {
	var func_EncPubStringWithIndex func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncPubStringWithIndex, sdk.keyxx, "EncPubStringWithIndex")
	p0 := sdk.getCStrPtr(plain)
	r := func_EncPubStringWithIndex(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncFloat(plain float64) string {
	var func_EncFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncFloat, sdk.keyxx, "EncFloat")
	p0 := sdk.getCFloat64Ptr(plain)
	r := func_EncFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncDouble(plain float64) string {
	var func_EncDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncDouble, sdk.keyxx, "EncDouble")
	p0 := sdk.getCFloat64Ptr(plain)
	r := func_EncDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncPubDouble(plain float64) string {
	var func_EncPubDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncPubDouble, sdk.keyxx, "EncPubDouble")
	p0 := sdk.getCFloat64Ptr(plain)
	r := func_EncPubDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncBinary(plain int, l int) string {
	var func_EncBinary func(C.int, C.int) *C.char
	purego.RegisterLibFunc(&func_EncBinary, sdk.keyxx, "EncBinary")
	p0 := sdk.getCIntPtr(plain)
	p1 := sdk.getCIntPtr(l)
	r := func_EncBinary(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncPubBinary(plain int, l int) string {
	var func_EncPubBinary func(C.int, C.int) *C.char
	purego.RegisterLibFunc(&func_EncPubBinary, sdk.keyxx, "EncPubBinary")
	p0 := sdk.getCIntPtr(plain)
	p1 := sdk.getCIntPtr(l)
	r := func_EncPubBinary(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) DecInt(input string) int {
	var func_DecInt func(*C.char) C.int
	purego.RegisterLibFunc(&func_DecInt, sdk.keyxx, "DecInt")
	p0 := sdk.getCStrPtr(input)
	r := func_DecInt(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) DecFloat(input string) float64 {
	var func_DecFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_DecFloat, sdk.keyxx, "DecFloat")
	p0 := sdk.getCStrPtr(input)
	r := func_DecFloat(p0)
	return sdk.getReturnFloat64(r)
}

func (sdk *Linux_SDK) DecDouble(input string) float64 {
	var func_DecDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_DecDouble, sdk.keyxx, "DecDouble")
	p0 := sdk.getCStrPtr(input)
	r := func_DecDouble(p0)
	return sdk.getReturnFloat64(r)
}

func (sdk *Linux_SDK) Decrypt(input string) string {
	var func_Decrypt func(*C.char) *C.char
	purego.RegisterLibFunc(&func_Decrypt, sdk.keyxx, "Decrypt")
	p0 := sdk.getCStrPtr(input)
	r := func_Decrypt(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) DecBinary(input string) int {
	var func_DecBinary func(*C.char) C.int
	purego.RegisterLibFunc(&func_DecBinary, sdk.keyxx, "DecBinary")
	p0 := sdk.getCStrPtr(input)
	r := func_DecBinary(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) GenSign(input string) string {
	var func_GenSign func(*C.char) *C.char
	purego.RegisterLibFunc(&func_GenSign, sdk.keyxx, "GenSign")
	p0 := sdk.getCStrPtr(input)
	r := func_GenSign(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) VerifySign(input string, sign string) int {
	var func_VerifySign func(*C.char, *C.char) C.int
	purego.RegisterLibFunc(&func_VerifySign, sdk.keyxx, "VerifySign")
	p0 := sdk.getCStrPtr(input)
	p1 := sdk.getCStrPtr(sign)
	r := func_VerifySign(p0, p1)
	return int(C.int(r))
}

func (sdk *Linux_SDK) AddCipherInt(c1 string, c2 string) string {
	var func_AddCipherInt func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_AddCipherInt, sdk.keyxx, "AddCipherInt")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_AddCipherInt(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SubCipherInt(c1 string, c2 string) string {
	var func_SubCipherInt func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_SubCipherInt, sdk.keyxx, "SubCipherInt")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_SubCipherInt(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) MulCipherInt(c1 string, c2 string) string {
	var func_MulCipherInt func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_MulCipherInt, sdk.keyxx, "MulCipherInt")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_MulCipherInt(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) AddCipherFloat(c1 string, c2 string) string {
	var func_AddCipherFloat func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_AddCipherFloat, sdk.keyxx, "AddCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_AddCipherFloat(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SubCipherFloat(c1 string, c2 string) string {
	var func_SubCipherFloat func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_SubCipherFloat, sdk.keyxx, "SubCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_SubCipherFloat(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) MulCipherFloat(c1 string, c2 string) string {
	var func_MulCipherFloat func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_MulCipherFloat, sdk.keyxx, "MulCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_MulCipherFloat(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) DivCipherFloat(c1 string, c2 string) string {
	var func_DivCipherFloat func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_DivCipherFloat, sdk.keyxx, "DivCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_DivCipherFloat(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) AddCipherDouble(c1 string, c2 string) string {
	var func_AddCipherDouble func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_AddCipherDouble, sdk.keyxx, "AddCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_AddCipherDouble(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SubCipherDouble(c1 string, c2 string) string {
	var func_SubCipherDouble func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_SubCipherDouble, sdk.keyxx, "SubCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_SubCipherDouble(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) MulCipherDouble(c1 string, c2 string) string {
	var func_MulCipherDouble func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_MulCipherDouble, sdk.keyxx, "MulCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_MulCipherDouble(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) DivCipherDouble(c1 string, c2 string) string {
	var func_DivCipherDouble func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_DivCipherDouble, sdk.keyxx, "DivCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_DivCipherDouble(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) XORCipherBinary(c1 string, c2 string) string {
	var func_XORCipherBinary func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_XORCipherBinary, sdk.keyxx, "XORCipherBinary")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_XORCipherBinary(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ANDCipherBinary(c1 string, c2 string) string {
	var func_ANDCipherBinary func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_ANDCipherBinary, sdk.keyxx, "ANDCipherBinary")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_ANDCipherBinary(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ORCipherBinary(c1 string, c2 string) string {
	var func_ORCipherBinary func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_ORCipherBinary, sdk.keyxx, "ORCipherBinary")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_ORCipherBinary(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) NOTCipherBinary(c1 string) string {
	var func_NOTCipherBinary func(*C.char) *C.char
	purego.RegisterLibFunc(&func_NOTCipherBinary, sdk.keyxx, "NOTCipherBinary")
	p0 := sdk.getCStrPtr(c1)
	r := func_NOTCipherBinary(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ShiftLeftCipherBinary(c1 string, _bias int) string {
	var func_ShiftLeftCipherBinary func(*C.char, C.int) *C.char
	purego.RegisterLibFunc(&func_ShiftLeftCipherBinary, sdk.keyxx, "ShiftLeftCipherBinary")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCIntPtr(_bias)
	r := func_ShiftLeftCipherBinary(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ShiftRightCipherBinary(c1 string, _bias int) string {
	var func_ShiftRightCipherBinary func(*C.char, C.int) *C.char
	purego.RegisterLibFunc(&func_ShiftRightCipherBinary, sdk.keyxx, "ShiftRightCipherBinary")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCIntPtr(_bias)
	r := func_ShiftRightCipherBinary(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) CompareCipherInt(c1 string, c2 string) int {
	var func_CompareCipherInt func(*C.char, *C.char) C.int
	purego.RegisterLibFunc(&func_CompareCipherInt, sdk.keyxx, "CompareCipherInt")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_CompareCipherInt(p0, p1)
	return int(C.int(r))
}

func (sdk *Linux_SDK) CompareCipherString(c1 string, c2 string) int {
	var func_CompareCipherString func(*C.char, *C.char) C.int
	purego.RegisterLibFunc(&func_CompareCipherString, sdk.keyxx, "CompareCipherString")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_CompareCipherString(p0, p1)
	return int(C.int(r))
}

func (sdk *Linux_SDK) CompareCipherFloat(c1 string, c2 string) int {
	var func_CompareCipherFloat func(*C.char, *C.char) C.int
	purego.RegisterLibFunc(&func_CompareCipherFloat, sdk.keyxx, "CompareCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_CompareCipherFloat(p0, p1)
	return int(C.int(r))
}

func (sdk *Linux_SDK) CompareCipherDouble(c1 string, c2 string) int {
	var func_CompareCipherDouble func(*C.char, *C.char) C.int
	purego.RegisterLibFunc(&func_CompareCipherDouble, sdk.keyxx, "CompareCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	r := func_CompareCipherDouble(p0, p1)
	return int(C.int(r))
}

func (sdk *Linux_SDK) ABSCipherFloat(c1 string) string {
	var func_ABSCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ABSCipherFloat, sdk.keyxx, "ABSCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ABSCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) PowCipherFloat(c1 string, n int) string {
	var func_PowCipherFloat func(*C.char, C.int) *C.char
	purego.RegisterLibFunc(&func_PowCipherFloat, sdk.keyxx, "PowCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCIntPtr(n)
	r := func_PowCipherFloat(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SqrtCipherFloat(c1 string) string {
	var func_SqrtCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_SqrtCipherFloat, sdk.keyxx, "SqrtCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_SqrtCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) PowCipherFloat_FractionOrder(c1 string, n int, m int) string {
	var func_PowCipherFloat_FractionOrder func(*C.char, C.int, C.int) *C.char
	purego.RegisterLibFunc(&func_PowCipherFloat_FractionOrder, sdk.keyxx, "PowCipherFloat_FractionOrder")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCIntPtr(n)
	p2 := sdk.getCIntPtr(m)
	r := func_PowCipherFloat_FractionOrder(p0, p1, p2)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) PowCipherFloat_RealOrder(c1 string, a float64) string {
	var func_PowCipherFloat_RealOrder func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_PowCipherFloat_RealOrder, sdk.keyxx, "PowCipherFloat_RealOrder")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCFloat64Ptr(a)
	r := func_PowCipherFloat_RealOrder(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) LogCipherFloat(c1 string) string {
	var func_LogCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_LogCipherFloat, sdk.keyxx, "LogCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_LogCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ExpCipherFloat(c1 string) string {
	var func_ExpCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ExpCipherFloat, sdk.keyxx, "ExpCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ExpCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SinCipherFloat(c1 string) string {
	var func_SinCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_SinCipherFloat, sdk.keyxx, "SinCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_SinCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) CosCipherFloat(c1 string) string {
	var func_CosCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_CosCipherFloat, sdk.keyxx, "CosCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_CosCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) TanCipherFloat(c1 string) string {
	var func_TanCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_TanCipherFloat, sdk.keyxx, "TanCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_TanCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArcsinCipherFloat(c1 string) string {
	var func_ArcsinCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArcsinCipherFloat, sdk.keyxx, "ArcsinCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArcsinCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArccosCipherFloat(c1 string) string {
	var func_ArccosCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArccosCipherFloat, sdk.keyxx, "ArccosCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArccosCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArctanCipherFloat(c1 string) string {
	var func_ArctanCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArctanCipherFloat, sdk.keyxx, "ArctanCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArctanCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SinhCipherFloat(c1 string) string {
	var func_SinhCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_SinhCipherFloat, sdk.keyxx, "SinhCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_SinhCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) CoshCipherFloat(c1 string) string {
	var func_CoshCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_CoshCipherFloat, sdk.keyxx, "CoshCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_CoshCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) TanhCipherFloat(c1 string) string {
	var func_TanhCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_TanhCipherFloat, sdk.keyxx, "TanhCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_TanhCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArcsinhCipherFloat(c1 string) string {
	var func_ArcsinhCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArcsinhCipherFloat, sdk.keyxx, "ArcsinhCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArcsinhCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArccoshCipherFloat(c1 string) string {
	var func_ArccoshCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArccoshCipherFloat, sdk.keyxx, "ArccoshCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArccoshCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArctanhCipherFloat(c1 string) string {
	var func_ArctanhCipherFloat func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArctanhCipherFloat, sdk.keyxx, "ArctanhCipherFloat")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArctanhCipherFloat(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ABSCipherDouble(c1 string) string {
	var func_ABSCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ABSCipherDouble, sdk.keyxx, "ABSCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ABSCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) PowCipherDouble(c1 string, n int) string {
	var func_PowCipherDouble func(*C.char, C.int) *C.char
	purego.RegisterLibFunc(&func_PowCipherDouble, sdk.keyxx, "PowCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCIntPtr(n)
	r := func_PowCipherDouble(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SqrtCipherDouble(c1 string) string {
	var func_SqrtCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_SqrtCipherDouble, sdk.keyxx, "SqrtCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_SqrtCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) PowCipherDouble_FractionOrder(c1 string, n int, m int) string {
	var func_PowCipherDouble_FractionOrder func(*C.char, C.int, C.int) *C.char
	purego.RegisterLibFunc(&func_PowCipherDouble_FractionOrder, sdk.keyxx, "PowCipherDouble_FractionOrder")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCIntPtr(n)
	p2 := sdk.getCIntPtr(m)
	r := func_PowCipherDouble_FractionOrder(p0, p1, p2)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) PowCipherDouble_RealOrder(c1 string, a float64) string {
	var func_PowCipherDouble_RealOrder func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_PowCipherDouble_RealOrder, sdk.keyxx, "PowCipherDouble_RealOrder")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCFloat64Ptr(a)
	r := func_PowCipherDouble_RealOrder(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) LogCipherDouble(c1 string) string {
	var func_LogCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_LogCipherDouble, sdk.keyxx, "LogCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_LogCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ExpCipherDouble(c1 string) string {
	var func_ExpCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ExpCipherDouble, sdk.keyxx, "ExpCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ExpCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SinCipherDouble(c1 string) string {
	var func_SinCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_SinCipherDouble, sdk.keyxx, "SinCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_SinCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) CosCipherDouble(c1 string) string {
	var func_CosCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_CosCipherDouble, sdk.keyxx, "CosCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_CosCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) TanCipherDouble(c1 string) string {
	var func_TanCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_TanCipherDouble, sdk.keyxx, "TanCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_TanCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArcsinCipherDouble(c1 string) string {
	var func_ArcsinCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArcsinCipherDouble, sdk.keyxx, "ArcsinCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArcsinCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArccosCipherDouble(c1 string) string {
	var func_ArccosCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArccosCipherDouble, sdk.keyxx, "ArccosCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArccosCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArctanCipherDouble(c1 string) string {
	var func_ArctanCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArctanCipherDouble, sdk.keyxx, "ArctanCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArctanCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SinhCipherDouble(c1 string) string {
	var func_SinhCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_SinhCipherDouble, sdk.keyxx, "SinhCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_SinhCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) CoshCipherDouble(c1 string) string {
	var func_CoshCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_CoshCipherDouble, sdk.keyxx, "CoshCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_CoshCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) TanhCipherDouble(c1 string) string {
	var func_TanhCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_TanhCipherDouble, sdk.keyxx, "TanhCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_TanhCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArcsinhCipherDouble(c1 string) string {
	var func_ArcsinhCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArcsinhCipherDouble, sdk.keyxx, "ArcsinhCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArcsinhCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArccoshCipherDouble(c1 string) string {
	var func_ArccoshCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArccoshCipherDouble, sdk.keyxx, "ArccoshCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArccoshCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ArctanhCipherDouble(c1 string) string {
	var func_ArctanhCipherDouble func(*C.char) *C.char
	purego.RegisterLibFunc(&func_ArctanhCipherDouble, sdk.keyxx, "ArctanhCipherDouble")
	p0 := sdk.getCStrPtr(c1)
	r := func_ArctanhCipherDouble(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) ConcatString(data1 string, data2 string) string {
	var func_ConcatString func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_ConcatString, sdk.keyxx, "ConcatString")
	p0 := sdk.getCStrPtr(data1)
	p1 := sdk.getCStrPtr(data2)
	r := func_ConcatString(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) Substring(data string, start int, _end int) string {
	var func_Substring func(*C.char, C.int, C.int) *C.char
	purego.RegisterLibFunc(&func_Substring, sdk.keyxx, "Substring")
	p0 := sdk.getCStrPtr(data)
	p1 := sdk.getCIntPtr(start)
	p2 := sdk.getCIntPtr(_end)
	r := func_Substring(p0, p1, p2)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncPubIntExtend(plain int) string {
	var func_EncPubIntExtend func(C.int) *C.char
	purego.RegisterLibFunc(&func_EncPubIntExtend, sdk.keyxx, "EncPubIntExtend")
	p0 := sdk.getCIntPtr(plain)
	r := func_EncPubIntExtend(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EncPubStringExtend(plain string) string {
	var func_EncPubStringExtend func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncPubStringExtend, sdk.keyxx, "EncPubStringExtend")
	p0 := sdk.getCStrPtr(plain)
	r := func_EncPubStringExtend(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) TranEN(cipher string) string {
	var func_TranEN func(*C.char) *C.char
	purego.RegisterLibFunc(&func_TranEN, sdk.keyxx, "TranEN")
	p0 := sdk.getCStrPtr(cipher)
	r := func_TranEN(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) TranStringEN(_cs string) string {
	var func_TranStringEN func(*C.char) *C.char
	purego.RegisterLibFunc(&func_TranStringEN, sdk.keyxx, "TranStringEN")
	p0 := sdk.getCStrPtr(_cs)
	r := func_TranStringEN(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) DecStringNormal(_cs string) string {
	var func_DecStringNormal func(*C.char) *C.char
	purego.RegisterLibFunc(&func_DecStringNormal, sdk.keyxx, "DecStringNormal")
	p0 := sdk.getCStrPtr(_cs)
	r := func_DecStringNormal(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) DecString(cs string) string {
	var func_DecString func(*C.char) *C.char
	purego.RegisterLibFunc(&func_DecString, sdk.keyxx, "DecString")
	p0 := sdk.getCStrPtr(cs)
	r := func_DecString(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) TranStringENMix(_cs string) string {
	var func_TranStringENMix func(*C.char) *C.char
	purego.RegisterLibFunc(&func_TranStringENMix, sdk.keyxx, "TranStringENMix")
	p0 := sdk.getCStrPtr(_cs)
	r := func_TranStringENMix(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) AddByteExtend(cipher1 string, cipher2 string) string {
	var func_AddByteExtend func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_AddByteExtend, sdk.keyxx, "AddByteExtend")
	p0 := sdk.getCStrPtr(cipher1)
	p1 := sdk.getCStrPtr(cipher2)
	r := func_AddByteExtend(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) SubstractByteExtend(cipher1 string, cipher2 string) string {
	var func_SubstractByteExtend func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_SubstractByteExtend, sdk.keyxx, "SubstractByteExtend")
	p0 := sdk.getCStrPtr(cipher1)
	p1 := sdk.getCStrPtr(cipher2)
	r := func_SubstractByteExtend(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) MultiplyByteExtend(cipher1 string, cipher2 string) string {
	var func_MultiplyByteExtend func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_MultiplyByteExtend, sdk.keyxx, "MultiplyByteExtend")
	p0 := sdk.getCStrPtr(cipher1)
	p1 := sdk.getCStrPtr(cipher2)
	r := func_MultiplyByteExtend(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EqualByteExtend(cipher1 string, cipher2 string) string {
	var func_EqualByteExtend func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_EqualByteExtend, sdk.keyxx, "EqualByteExtend")
	p0 := sdk.getCStrPtr(cipher1)
	p1 := sdk.getCStrPtr(cipher2)
	r := func_EqualByteExtend(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) CatCipherStringMix(_cs1 string, _cs2 string) string {
	var func_CatCipherStringMix func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_CatCipherStringMix, sdk.keyxx, "CatCipherStringMix")
	p0 := sdk.getCStrPtr(_cs1)
	p1 := sdk.getCStrPtr(_cs2)
	r := func_CatCipherStringMix(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) EqualStringExtend(_cs1 string, _cs2 string) string {
	var func_EqualStringExtend func(*C.char, *C.char) *C.char
	purego.RegisterLibFunc(&func_EqualStringExtend, sdk.keyxx, "EqualStringExtend")
	p0 := sdk.getCStrPtr(_cs1)
	p1 := sdk.getCStrPtr(_cs2)
	r := func_EqualStringExtend(p0, p1)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) IsCipher(cipherstr string) int {
	var func_IsCipher func(*C.char) C.int
	purego.RegisterLibFunc(&func_IsCipher, sdk.keyxx, "IsCipher")
	p0 := sdk.getCStrPtr(cipherstr)
	r := func_IsCipher(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) IsCipherInt(cipherstr string) int {
	var func_IsCipherInt func(*C.char) C.int
	purego.RegisterLibFunc(&func_IsCipherInt, sdk.keyxx, "IsCipherInt")
	p0 := sdk.getCStrPtr(cipherstr)
	r := func_IsCipherInt(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) IsCipherFloat(cipherstr string) int {
	var func_IsCipherFloat func(*C.char) C.int
	purego.RegisterLibFunc(&func_IsCipherFloat, sdk.keyxx, "IsCipherFloat")
	p0 := sdk.getCStrPtr(cipherstr)
	r := func_IsCipherFloat(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) IsCipherDouble(cipherstr string) int {
	var func_IsCipherDouble func(*C.char) C.int
	purego.RegisterLibFunc(&func_IsCipherDouble, sdk.keyxx, "IsCipherDouble")
	p0 := sdk.getCStrPtr(cipherstr)
	r := func_IsCipherDouble(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) IsCipherString(cipherstr string) int {
	var func_IsCipherString func(*C.char) C.int
	purego.RegisterLibFunc(&func_IsCipherString, sdk.keyxx, "IsCipherString")
	p0 := sdk.getCStrPtr(cipherstr)
	r := func_IsCipherString(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) IsCipherStringWithIndex(cipherstr string) int {
	var func_IsCipherStringWithIndex func(*C.char) C.int
	purego.RegisterLibFunc(&func_IsCipherStringWithIndex, sdk.keyxx, "IsCipherStringWithIndex")
	p0 := sdk.getCStrPtr(cipherstr)
	r := func_IsCipherStringWithIndex(p0)
	return int(C.int(r))
}

func (sdk *Linux_SDK) CompareStringWithIndex(c1 string, c2 string, max int) int {
	var func_CompareStringWithIndex func(*C.char, *C.char, C.int) C.int
	purego.RegisterLibFunc(&func_CompareStringWithIndex, sdk.keyxx, "CompareStringWithIndex")
	p0 := sdk.getCStrPtr(c1)
	p1 := sdk.getCStrPtr(c2)
	p2 := sdk.getCIntPtr(max)
	r := func_CompareStringWithIndex(p0, p1, p2)
	return int(C.int(r))
}

func (sdk *Linux_SDK) EncStringWithIndex(c1 string) string {
	var func_EncStringWithIndex func(*C.char) *C.char
	purego.RegisterLibFunc(&func_EncStringWithIndex, sdk.keyxx, "EncStringWithIndex")
	p0 := sdk.getCStrPtr(c1)
	r := func_EncStringWithIndex(p0)
	return sdk.getReturnString(r)
}

func (sdk *Linux_SDK) DecStringWithIndex(input string) string {
	var func_DecStringWithIndex func(*C.char) *C.char
	purego.RegisterLibFunc(&func_DecStringWithIndex, sdk.keyxx, "DecStringWithIndex")
	p0 := sdk.getCStrPtr(input)
	r := func_DecStringWithIndex(p0)
	return sdk.getReturnString(r)
}
