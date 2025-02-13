package sdk

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/base64"
)

// var SDKPool map[string]*Windows_SDK
var SDKPool map[string]*Linux_SDK
var tmpFileList []string

func genTempFile(libPath, libFile string) string {
	file, err := os.OpenFile(libPath + libFile, os.O_RDONLY, 0)
	defer file.Close()
	res := ""
	if err != nil {
		panic(err)
	} else {
		str, _ := ioutil.ReadAll(file)
		// tmpFile, _ := ioutil.TempFile(libPath, "temp.*.dll")
		tmpFile, _ := ioutil.TempFile(libPath, "temp.*.so")
		res = tmpFile.Name()
		fmt.Println("Created File: " + res)
		_, err = tmpFile.Write(str)
	    if err != nil {
	        panic("Failed to write to temporary file")
	    }
	    tmpFile.Close()
	    tmpFileList = append(tmpFileList, res)
	}
	return res
}

func Close() {
	for _, v := range SDKPool {
		v.Release()
	}
	for _, v := range tmpFileList {
		os.Remove(v)
	}
}

func InitSDKPool() {
	// SDKPool = make(map[string]*Windows_SDK)
	SDKPool = make(map[string]*Linux_SDK)
	tmpFileList = make([]string, 0)
}

func AddSDK(userKey, libPath, libFile string, m, n, q, p int) {
	newFile := genTempFile(libPath, libFile)
	// sdk := new(Windows_SDK)
	sdk := new(Linux_SDK)
	sdk.SetLibPath(newFile)
	sdk.InitLib()
	sdk.Init(m, n, q, p)
	sdk.UserKey = userKey
	SDKPool[userKey] = sdk
}

func DelSDK(userKey string) {
	delete(SDKPool, userKey)
}

func ListSDK() {
	for k, _ := range SDKPool {
		fmt.Println(k)
	}
}

func GetSDK(userKey string) *Linux_SDK {
	return SDKPool[userKey]
}

func Init(userKey string, m_in int, n_in int, q_in int, p_in int)  {
	p0 := m_in
	p1 := n_in
	p2 := q_in
	p3 := p_in
	SDKPool[userKey].Init(p0, p1, p2, p3)
}

func InitCodeType(userKey string, ct int)  {
	p0 := ct
	SDKPool[userKey].InitCodeType(p0)
}

func InitSerializeType(userKey string, st int)  {
	p0 := st
	SDKPool[userKey].InitSerializeType(p0)
}

func InitStrCaseChangeMode(userKey string, scc int)  {
	p0 := scc
	SDKPool[userKey].InitStrCaseChangeMode(p0)
}

func InitIsFastString(userKey string, iss int)  {
	p0 := iss
	SDKPool[userKey].InitIsFastString(p0)
}

func SM3(userKey string, input string) string {
	p0 := input
	r := SDKPool[userKey].SM3(p0)
	return r
}

func LoadPrivKey(userKey string, skbfile string)  {
	p0 := skbfile
	SDKPool[userKey].LoadPrivKey(p0)
}

func LoadPubKey(userKey string, pkbfile string)  {
	p0 := pkbfile
	SDKPool[userKey].LoadPubKey(p0)
}

func LoadDictionary(userKey string, dictbfile string)  {
	p0 := dictbfile
	SDKPool[userKey].LoadDictionary(p0)
}

func LoadPrivKeyString(userKey string, skbstring []byte)  {
	p0 := base64.URLEncoding.EncodeToString(skbstring)
	SDKPool[userKey].LoadPrivKeyString(p0)
}

func LoadPubKeyString(userKey string, pkbstring []byte)  {
	p0 := base64.URLEncoding.EncodeToString(pkbstring)
	SDKPool[userKey].LoadPubKeyString(p0)
}

func LoadDictionaryString(userKey string, dictbstring []byte)  {
	p0 := base64.URLEncoding.EncodeToString(dictbstring)
	SDKPool[userKey].LoadDictionaryString(p0)
}

func LoadBiasKey(userKey string, bkfile string)  {
	p0 := bkfile
	SDKPool[userKey].LoadBiasKey(p0)
}

func LoadExKey(userKey string, dictExbfile string)  {
	p0 := dictExbfile
	SDKPool[userKey].LoadExKey(p0)
}

func LoadPrivkeyFloat(userKey string, skffile string)  {
	p0 := skffile
	SDKPool[userKey].LoadPrivkeyFloat(p0)
}

func LoadDictionaryFloat(userKey string, dictffile string)  {
	p0 := dictffile
	SDKPool[userKey].LoadDictionaryFloat(p0)
}

func LoadExpressionBlockKeys(userKey string)  {
	SDKPool[userKey].LoadExpressionBlockKeys()
}

func LoadExpressionFloatKeys(userKey string)  {
	SDKPool[userKey].LoadExpressionFloatKeys()
}

func InitExpression(userKey string)  {
	SDKPool[userKey].InitExpression()
}

func ClearExpressionVariable(userKey string)  {
	SDKPool[userKey].ClearExpressionVariable()
}

func AddExpressionVariable(userKey string, _name string, _type string, _value string)  {
	p0 := _name
	p1 := _type
	p2 := _value
	SDKPool[userKey].AddExpressionVariable(p0, p1, p2)
}

func ExpressionCalculation(userKey string, formula string) string {
	p0 := formula
	r := SDKPool[userKey].ExpressionCalculation(p0)
	return r
}

func GenSKB(userKey string, m int, n int, q int, p int, filename string)  {
	p0 := m
	p1 := n
	p2 := q
	p3 := p
	p4 := filename
	SDKPool[userKey].GenSKB(p0, p1, p2, p3, p4)
}

func GenPKB(userKey string, skbfile string, filename string)  {
	p0 := skbfile
	p1 := filename
	SDKPool[userKey].GenPKB(p0, p1)
}

func GenDictB(userKey string, skbfile string, filename string, delta float64)  {
	p0 := skbfile
	p1 := filename
	p2 := delta
	SDKPool[userKey].GenDictB(p0, p1, p2)
}

func GenSKBString(userKey string, m int, n int, q int, p int) []byte {
	p0 := m
	p1 := n
	p2 := q
	p3 := p
	r := SDKPool[userKey].GenSKBString(p0, p1, p2, p3)
	res, _ := base64.URLEncoding.DecodeString(r)
	return res
}

func GenPKBString(userKey string, skbfile []byte) []byte {
	p0 := base64.URLEncoding.EncodeToString(skbfile)
	r := SDKPool[userKey].GenPKBString(p0)
	res, _ := base64.URLEncoding.DecodeString(r)
	return res
}

func GenDictBString(userKey string, skbfile []byte, delta float64) []byte {
	p0 := base64.URLEncoding.EncodeToString(skbfile)
	p1 := delta
	r := SDKPool[userKey].GenDictBString(p0, p1)
	res, _ := base64.URLEncoding.DecodeString(r)
	return res
}

func GenBK(userKey string, n int, q int, p int, filename string)  {
	p0 := n
	p1 := q
	p2 := p
	p3 := filename
	SDKPool[userKey].GenBK(p0, p1, p2, p3)
}

func GenDictExB(userKey string, bkfile string, pkbfile string, dictbfile string, filename string)  {
	p0 := bkfile
	p1 := pkbfile
	p2 := dictbfile
	p3 := filename
	SDKPool[userKey].GenDictExB(p0, p1, p2, p3)
}

func GenSKF(userKey string, p int, filename string)  {
	p0 := p
	p1 := filename
	SDKPool[userKey].GenSKF(p0, p1)
}

func GenDictF(userKey string, skffile string, filename string)  {
	p0 := skffile
	p1 := filename
	SDKPool[userKey].GenDictF(p0, p1)
}

func EncInt(userKey string, plain int) string {
	p0 := plain
	r := SDKPool[userKey].EncInt(p0)
	return r
}

func EncPubInt(userKey string, plain int) string {
	p0 := plain
	r := SDKPool[userKey].EncPubInt(p0)
	return r
}

func EncString(userKey string, plain string) string {
	p0 := plain
	r := SDKPool[userKey].EncString(p0)
	return r
}

func EncPubString(userKey string, plain string) string {
	p0 := plain
	r := SDKPool[userKey].EncPubString(p0)
	return r
}

func EncPubStringWithIndex(userKey string, plain string) string {
	p0 := plain
	r := SDKPool[userKey].EncPubStringWithIndex(p0)
	return r
}

func EncFloat(userKey string, plain float64) string {
	p0 := plain
	r := SDKPool[userKey].EncFloat(p0)
	return r
}

func EncDouble(userKey string, plain float64) string {
	p0 := plain
	r := SDKPool[userKey].EncDouble(p0)
	return r
}

func EncPubDouble(userKey string, plain float64) string {
	p0 := plain
	r := SDKPool[userKey].EncPubDouble(p0)
	return r
}

func EncBinary(userKey string, plain int, l int) string {
	p0 := plain
	p1 := l
	r := SDKPool[userKey].EncBinary(p0, p1)
	return r
}

func EncPubBinary(userKey string, plain int, l int) string {
	p0 := plain
	p1 := l
	r := SDKPool[userKey].EncPubBinary(p0, p1)
	return r
}

func DecInt(userKey string, input string) int {
	p0 := input
	r := SDKPool[userKey].DecInt(p0)
	return r
}

func DecFloat(userKey string, input string) float64 {
	p0 := input
	r := SDKPool[userKey].DecFloat(p0)
	return r
}

func DecDouble(userKey string, input string) float64 {
	p0 := input
	r := SDKPool[userKey].DecDouble(p0)
	return r
}

func Decrypt(userKey string, input string) string {
	p0 := input
	r := SDKPool[userKey].Decrypt(p0)
	return r
}

func DecBinary(userKey string, input string) int {
	p0 := input
	r := SDKPool[userKey].DecBinary(p0)
	return r
}

func GenSign(userKey string, input string) string {
	p0 := input
	r := SDKPool[userKey].GenSign(p0)
	return r
}

func VerifySign(userKey string, input string, sign string) int {
	p0 := input
	p1 := sign
	r := SDKPool[userKey].VerifySign(p0, p1)
	return r
}

func AddCipherInt(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].AddCipherInt(p0, p1)
	return r
}

func SubCipherInt(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].SubCipherInt(p0, p1)
	return r
}

func MulCipherInt(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].MulCipherInt(p0, p1)
	return r
}

func AddCipherFloat(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].AddCipherFloat(p0, p1)
	return r
}

func SubCipherFloat(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].SubCipherFloat(p0, p1)
	return r
}

func MulCipherFloat(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].MulCipherFloat(p0, p1)
	return r
}

func DivCipherFloat(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].DivCipherFloat(p0, p1)
	return r
}

func AddCipherDouble(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].AddCipherDouble(p0, p1)
	return r
}

func SubCipherDouble(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].SubCipherDouble(p0, p1)
	return r
}

func MulCipherDouble(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].MulCipherDouble(p0, p1)
	return r
}

func DivCipherDouble(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].DivCipherDouble(p0, p1)
	return r
}

func XORCipherBinary(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].XORCipherBinary(p0, p1)
	return r
}

func ANDCipherBinary(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].ANDCipherBinary(p0, p1)
	return r
}

func ORCipherBinary(userKey string, c1 string, c2 string) string {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].ORCipherBinary(p0, p1)
	return r
}

func NOTCipherBinary(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].NOTCipherBinary(p0)
	return r
}

func ShiftLeftCipherBinary(userKey string, c1 string, _bias int) string {
	p0 := c1
	p1 := _bias
	r := SDKPool[userKey].ShiftLeftCipherBinary(p0, p1)
	return r
}

func ShiftRightCipherBinary(userKey string, c1 string, _bias int) string {
	p0 := c1
	p1 := _bias
	r := SDKPool[userKey].ShiftRightCipherBinary(p0, p1)
	return r
}

func CompareCipherInt(userKey string, c1 string, c2 string) int {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].CompareCipherInt(p0, p1)
	return r
}

func CompareCipherString(userKey string, c1 string, c2 string) int {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].CompareCipherString(p0, p1)
	return r
}

func CompareCipherFloat(userKey string, c1 string, c2 string) int {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].CompareCipherFloat(p0, p1)
	return r
}

func CompareCipherDouble(userKey string, c1 string, c2 string) int {
	p0 := c1
	p1 := c2
	r := SDKPool[userKey].CompareCipherDouble(p0, p1)
	return r
}

func ABSCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ABSCipherFloat(p0)
	return r
}

func PowCipherFloat(userKey string, c1 string, n int) string {
	p0 := c1
	p1 := n
	r := SDKPool[userKey].PowCipherFloat(p0, p1)
	return r
}

func SqrtCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].SqrtCipherFloat(p0)
	return r
}

func PowCipherFloat_FractionOrder(userKey string, c1 string, n int, m int) string {
	p0 := c1
	p1 := n
	p2 := m
	r := SDKPool[userKey].PowCipherFloat_FractionOrder(p0, p1, p2)
	return r
}

func PowCipherFloat_RealOrder(userKey string, c1 string, a float64) string {
	p0 := c1
	p1 := a
	r := SDKPool[userKey].PowCipherFloat_RealOrder(p0, p1)
	return r
}

func LogCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].LogCipherFloat(p0)
	return r
}

func ExpCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ExpCipherFloat(p0)
	return r
}

func SinCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].SinCipherFloat(p0)
	return r
}

func CosCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].CosCipherFloat(p0)
	return r
}

func TanCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].TanCipherFloat(p0)
	return r
}

func ArcsinCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArcsinCipherFloat(p0)
	return r
}

func ArccosCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArccosCipherFloat(p0)
	return r
}

func ArctanCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArctanCipherFloat(p0)
	return r
}

func SinhCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].SinhCipherFloat(p0)
	return r
}

func CoshCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].CoshCipherFloat(p0)
	return r
}

func TanhCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].TanhCipherFloat(p0)
	return r
}

func ArcsinhCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArcsinhCipherFloat(p0)
	return r
}

func ArccoshCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArccoshCipherFloat(p0)
	return r
}

func ArctanhCipherFloat(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArctanhCipherFloat(p0)
	return r
}

func ABSCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ABSCipherDouble(p0)
	return r
}

func PowCipherDouble(userKey string, c1 string, n int) string {
	p0 := c1
	p1 := n
	r := SDKPool[userKey].PowCipherDouble(p0, p1)
	return r
}

func SqrtCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].SqrtCipherDouble(p0)
	return r
}

func PowCipherDouble_FractionOrder(userKey string, c1 string, n int, m int) string {
	p0 := c1
	p1 := n
	p2 := m
	r := SDKPool[userKey].PowCipherDouble_FractionOrder(p0, p1, p2)
	return r
}

func PowCipherDouble_RealOrder(userKey string, c1 string, a float64) string {
	p0 := c1
	p1 := a
	r := SDKPool[userKey].PowCipherDouble_RealOrder(p0, p1)
	return r
}

func LogCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].LogCipherDouble(p0)
	return r
}

func ExpCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ExpCipherDouble(p0)
	return r
}

func SinCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].SinCipherDouble(p0)
	return r
}

func CosCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].CosCipherDouble(p0)
	return r
}

func TanCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].TanCipherDouble(p0)
	return r
}

func ArcsinCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArcsinCipherDouble(p0)
	return r
}

func ArccosCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArccosCipherDouble(p0)
	return r
}

func ArctanCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArctanCipherDouble(p0)
	return r
}

func SinhCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].SinhCipherDouble(p0)
	return r
}

func CoshCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].CoshCipherDouble(p0)
	return r
}

func TanhCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].TanhCipherDouble(p0)
	return r
}

func ArcsinhCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArcsinhCipherDouble(p0)
	return r
}

func ArccoshCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArccoshCipherDouble(p0)
	return r
}

func ArctanhCipherDouble(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].ArctanhCipherDouble(p0)
	return r
}

func ConcatString(userKey string, data1 string, data2 string) string {
	p0 := data1
	p1 := data2
	r := SDKPool[userKey].ConcatString(p0, p1)
	return r
}

func Substring(userKey string, data string, start int, _end int) string {
	p0 := data
	p1 := start
	p2 := _end
	r := SDKPool[userKey].Substring(p0, p1, p2)
	return r
}

func EncPubIntExtend(userKey string, plain int) string {
	p0 := plain
	r := SDKPool[userKey].EncPubIntExtend(p0)
	return r
}

func EncPubStringExtend(userKey string, plain string) string {
	p0 := plain
	r := SDKPool[userKey].EncPubStringExtend(p0)
	return r
}

func TranEN(userKey string, cipher string) string {
	p0 := cipher
	r := SDKPool[userKey].TranEN(p0)
	return r
}

func TranStringEN(userKey string, _cs string) string {
	p0 := _cs
	r := SDKPool[userKey].TranStringEN(p0)
	return r
}

func DecStringNormal(userKey string, _cs string) string {
	p0 := _cs
	r := SDKPool[userKey].DecStringNormal(p0)
	return r
}

func DecString(userKey string, cs string) string {
	p0 := cs
	r := SDKPool[userKey].DecString(p0)
	return r
}

func TranStringENMix(userKey string, _cs string) string {
	p0 := _cs
	r := SDKPool[userKey].TranStringENMix(p0)
	return r
}

func AddByteExtend(userKey string, cipher1 string, cipher2 string) string {
	p0 := cipher1
	p1 := cipher2
	r := SDKPool[userKey].AddByteExtend(p0, p1)
	return r
}

func SubstractByteExtend(userKey string, cipher1 string, cipher2 string) string {
	p0 := cipher1
	p1 := cipher2
	r := SDKPool[userKey].SubstractByteExtend(p0, p1)
	return r
}

func MultiplyByteExtend(userKey string, cipher1 string, cipher2 string) string {
	p0 := cipher1
	p1 := cipher2
	r := SDKPool[userKey].MultiplyByteExtend(p0, p1)
	return r
}

func EqualByteExtend(userKey string, cipher1 string, cipher2 string) string {
	p0 := cipher1
	p1 := cipher2
	r := SDKPool[userKey].EqualByteExtend(p0, p1)
	return r
}

func CatCipherStringMix(userKey string, _cs1 string, _cs2 string) string {
	p0 := _cs1
	p1 := _cs2
	r := SDKPool[userKey].CatCipherStringMix(p0, p1)
	return r
}

func EqualStringExtend(userKey string, _cs1 string, _cs2 string) string {
	p0 := _cs1
	p1 := _cs2
	r := SDKPool[userKey].EqualStringExtend(p0, p1)
	return r
}

func IsCipher(userKey string, cipherstr string) int {
	p0 := cipherstr
	r := SDKPool[userKey].IsCipher(p0)
	return r
}

func IsCipherInt(userKey string, cipherstr string) int {
	p0 := cipherstr
	r := SDKPool[userKey].IsCipherInt(p0)
	return r
}

func IsCipherFloat(userKey string, cipherstr string) int {
	p0 := cipherstr
	r := SDKPool[userKey].IsCipherFloat(p0)
	return r
}

func IsCipherDouble(userKey string, cipherstr string) int {
	p0 := cipherstr
	r := SDKPool[userKey].IsCipherDouble(p0)
	return r
}

func IsCipherString(userKey string, cipherstr string) int {
	p0 := cipherstr
	r := SDKPool[userKey].IsCipherString(p0)
	return r
}

func IsCipherStringWithIndex(userKey string, cipherstr string) int {
	p0 := cipherstr
	r := SDKPool[userKey].IsCipherStringWithIndex(p0)
	return r
}

func CompareStringWithIndex(userKey string, c1 string, c2 string, max int) int {
	p0 := c1
	p1 := c2
	p2 := max
	r := SDKPool[userKey].CompareStringWithIndex(p0, p1, p2)
	return r
}

func EncStringWithIndex(userKey string, c1 string) string {
	p0 := c1
	r := SDKPool[userKey].EncStringWithIndex(p0)
	return r
}

func DecStringWithIndex(userKey string, input string) string {
	p0 := input
	r := SDKPool[userKey].DecStringWithIndex(p0)
	return r
}
