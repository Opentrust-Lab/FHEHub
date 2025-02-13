package api

/*
#cgo LDFLAGS: -L. -lckks -lbgv -lbfv -Wl,-rpath=$ORIGIN/api
#include <stdlib.h>
#include <stdint.h>

// 声明C接口
extern const char* KeyGenCKKS(int multDepth, int scaleModSize, int batchSize);
extern const char* EncryptCKKS(char* ccLoc, size_t length_ccloc,
                const char* pkloc, size_t length_pkloc,
                double* real_parts, size_t length,
                int index);
extern	char* AddCKKS( const char* ccLoc, size_t length_ccloc,
         const char* c1Loc, size_t length_c1loc,  const char* c2Loc, size_t length_c2loc);
extern	char* MulCKKS( const char* ccLoc, size_t length_ccloc,
         const char* multKLoc, size_t length_multkloc,
         const char* c1Loc, size_t length_c1loc,  const char* c2Loc, size_t length_c2loc);
extern	char* RelinearizeCKKS( const char* ccLoc, size_t length_ccloc,
         const char* multKLoc, size_t length_multkloc,
         const char* cLoc, size_t length_cloc);
extern	char* RotCKKS( const char* ccLoc, size_t length_ccloc,
         const char* rotkloc, size_t length_rotkloc,
         const char* cloc, size_t length_cloc,
        int32_t index);
extern  const char* DecryptCKKS( const char* ccLoc, size_t length_ccloc,
         const char* skloc, size_t length_skloc,
         const char* cloc, size_t length_cloc,
        int vectorSize);
extern  void FreeString(const char* str);

extern const char* KeyGenBGV(int multiplicativeDepth, int plaintextModulus);
extern const char* EncryptBGV(char* ccLoc, size_t length_ccloc,
                const char* pkloc, size_t length_pkloc,
                int64_t* data, size_t length,
                int index);
extern	char* AddBGV( const char* ccLoc, size_t length_ccloc,
         const char* c1Loc, size_t length_c1loc,  const char* c2Loc, size_t length_c2loc);
extern	char* MulBGV( const char* ccLoc, size_t length_ccloc,
         const char* multKLoc, size_t length_multkloc,
         const char* c1Loc, size_t length_c1loc,  const char* c2Loc, size_t length_c2loc);
extern	char* RelinearizeBGV( const char* ccLoc, size_t length_ccloc,
         const char* multKLoc, size_t length_multkloc,
         const char* cLoc, size_t length_cloc);
extern	char* RotBGV( const char* ccLoc, size_t length_ccloc,
         const char* rotkloc, size_t length_rotkloc,
         const char* cloc, size_t length_cloc,
        int32_t index);
extern  const char* DecryptBGV( const char* ccLoc, size_t length_ccloc,
         const char* skloc, size_t length_skloc,
         const char* cloc, size_t length_cloc,
        int vectorSize);

extern const char* KeyGenBFV(int multiplicativeDepth, int plaintextModulus);
extern const char* EncryptBFV(char* ccLoc, size_t length_ccloc,
                const char* pkloc, size_t length_pkloc,
                int64_t* data, size_t length,
                int index);
extern	char* AddBFV( const char* ccLoc, size_t length_ccloc,
         const char* c1Loc, size_t length_c1loc,  const char* c2Loc, size_t length_c2loc);
extern	char* MulBFV( const char* ccLoc, size_t length_ccloc,
         const char* multKLoc, size_t length_multkloc,
         const char* c1Loc, size_t length_c1loc,  const char* c2Loc, size_t length_c2loc);
extern	char* RelinearizeBFV( const char* ccLoc, size_t length_ccloc,
         const char* multKLoc, size_t length_multkloc,
         const char* cLoc, size_t length_cloc);
extern	char* RotBFV( const char* ccLoc, size_t length_ccloc,
         const char* rotkloc, size_t length_rotkloc,
         const char* cloc, size_t length_cloc,
        int32_t index);
extern  const char* DecryptBFV( const char* ccLoc, size_t length_ccloc,
         const char* skloc, size_t length_skloc,
         const char* cloc, size_t length_cloc,
        int vectorSize);

*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

func KeyGenCKKS(multDepth, scaleModSize, batchSize int) []string {
	cStr := C.KeyGenCKKS(C.int(multDepth), C.int(scaleModSize), C.int(batchSize))
	defer C.FreeString(cStr) // 确保释放内存
	// 转换为 Go 字符串
	goResult := C.GoString(cStr)	
	// 使用分隔符解析
	stringsList := strings.Split(goResult, ";")
	fmt.Println("Received strings:", stringsList)
	return stringsList
}

func EncryptCKKS(ccLoc, pkLoc string, data []float64, id int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))
	pkLocC := C.CString(pkLoc)
	defer C.free(unsafe.Pointer(pkLocC))

	dataC := (*C.double)(unsafe.Pointer(&data[0])) // 转换为 C 指针

	resultENC := C.EncryptCKKS(
		ccLocC,
		C.size_t(len(ccLoc)),
		pkLocC,
		C.size_t(len(pkLoc)),
		dataC,
		C.size_t(len(data)),
		C.int(id),
	)
	defer C.FreeString(resultENC)
	resultEncGO := C.GoString(resultENC)
	return resultEncGO
}

func AddCKKS(ccLoc, c1Loc, c2Loc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))
	c2LocC := C.CString(c2Loc)
	defer C.free(unsafe.Pointer(c2LocC))
	
	// 调用 Add 函数
	addResult := C.AddCKKS(ccLocC, C.size_t(len(ccLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		c2LocC, C.size_t(len(c2Loc)))
	defer C.FreeString(addResult)
	
	return C.GoString(addResult)
}

func MulCKKS(ccLoc, multKLoc, c1Loc, c2Loc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	multkLocC := C.CString(multKLoc)
	defer C.free(unsafe.Pointer(multkLocC))	
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))
	c2LocC := C.CString(c2Loc)
	defer C.free(unsafe.Pointer(c2LocC))
	
	// 调用 Mul 函数
	mulResult := C.MulCKKS(ccLocC, C.size_t(len(ccLoc)),
		multkLocC, C.size_t(len(multKLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		c2LocC, C.size_t(len(c2Loc)))
	defer C.FreeString(mulResult)
	
	return C.GoString(mulResult)
}

func RelinearizeCKKS(ccLoc, multKLoc, cLoc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	multkLocC := C.CString(multKLoc)
	defer C.free(unsafe.Pointer(multkLocC))	
	cLocC := C.CString(cLoc)
	defer C.free(unsafe.Pointer(cLocC))
	fmt.Println(multKLoc)
	
	// 调用 Mul 函数
	reLineResult := C.RelinearizeCKKS(ccLocC, C.size_t(len(ccLoc)),
		multkLocC, C.size_t(len(multKLoc)),
		cLocC, C.size_t(len(cLoc)))
	defer C.FreeString(reLineResult)
	
	return C.GoString(reLineResult)
}

func RotCKKS(ccLoc, rotkLoc, c1Loc string, index int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	rotkLocC := C.CString(rotkLoc)
	defer C.free(unsafe.Pointer(rotkLocC))
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))

	// 调用 Rot 函数
	rotResult := C.RotCKKS(ccLocC, C.size_t(len(ccLoc)),
		rotkLocC, C.size_t(len(rotkLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		C.int32_t(index)) // index
	defer C.FreeString(rotResult)
	
	return C.GoString(rotResult)
}

func DecryptCKKS(ccLoc, skLoc, cipherloc string, vectorSize int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	skLocC := C.CString(skLoc)
	defer C.free(unsafe.Pointer(skLocC))
	cloc := C.CString(cipherloc)
	defer C.free(unsafe.Pointer(cloc))

	// 调用 Rot 函数
	decryptResult := C.DecryptCKKS(ccLocC, C.size_t(len(ccLoc)),
		skLocC, C.size_t(len(skLoc)),
		cloc, C.size_t(len(cipherloc)),
		C.int(vectorSize)) // vectorSize
	defer C.FreeString(decryptResult)	
	
	return C.GoString(decryptResult)
}

func KeyGenBGV(multiplicativeDepth, plaintextModulus int) []string {
	cStr := C.KeyGenBGV(C.int(multiplicativeDepth), C.int(plaintextModulus))
	defer C.FreeString(cStr) // 确保释放内存
	// 转换为 Go 字符串
	goResult := C.GoString(cStr)	
	// 使用分隔符解析
	stringsList := strings.Split(goResult, ";")
	fmt.Println("Received strings:", stringsList)
	return stringsList
}

func EncryptBGV(ccLoc, pkLoc string, data []int, id int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))
	pkLocC := C.CString(pkLoc)
	defer C.free(unsafe.Pointer(pkLocC))

	dataC := (*C.int64_t)(unsafe.Pointer(&data[0])) // 转换为 C 指针

	resultENC := C.EncryptBGV(
		ccLocC,
		C.size_t(len(ccLoc)),
		pkLocC,
		C.size_t(len(pkLoc)),
		dataC,
		C.size_t(len(data)),
		C.int(id),
	)
	defer C.FreeString(resultENC)
	resultEncGO := C.GoString(resultENC)
	return resultEncGO
}

func AddBGV(ccLoc, c1Loc, c2Loc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))
	c2LocC := C.CString(c2Loc)
	defer C.free(unsafe.Pointer(c2LocC))
	
	// 调用 Add 函数
	addResult := C.AddBGV(ccLocC, C.size_t(len(ccLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		c2LocC, C.size_t(len(c2Loc)))
	defer C.FreeString(addResult)
	
	return C.GoString(addResult)
}

func MulBGV(ccLoc, multKLoc, c1Loc, c2Loc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	multkLocC := C.CString(multKLoc)
	defer C.free(unsafe.Pointer(multkLocC))	
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))
	c2LocC := C.CString(c2Loc)
	defer C.free(unsafe.Pointer(c2LocC))
	
	// 调用 Mul 函数
	mulResult := C.MulBGV(ccLocC, C.size_t(len(ccLoc)),
		multkLocC, C.size_t(len(multKLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		c2LocC, C.size_t(len(c2Loc)))
	defer C.FreeString(mulResult)
	
	return C.GoString(mulResult)
}

func RelinearizeBGV(ccLoc, multKLoc, cLoc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	multkLocC := C.CString(multKLoc)
	defer C.free(unsafe.Pointer(multkLocC))	
	cLocC := C.CString(cLoc)
	defer C.free(unsafe.Pointer(cLocC))

	
	// 调用 Mul 函数
	reLineResult := C.RelinearizeBGV(ccLocC, C.size_t(len(ccLoc)),
		multkLocC, C.size_t(len(multKLoc)),
		cLocC, C.size_t(len(cLoc)))
	defer C.FreeString(reLineResult)
	
	return C.GoString(reLineResult)
}

func RotBGV(ccLoc, rotkLoc, c1Loc string, index int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	rotkLocC := C.CString(rotkLoc)
	defer C.free(unsafe.Pointer(rotkLocC))
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))

	// 调用 Rot 函数
	rotResult := C.RotBGV(ccLocC, C.size_t(len(ccLoc)),
		rotkLocC, C.size_t(len(rotkLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		C.int32_t(index)) // index
	defer C.FreeString(rotResult)
	
	return C.GoString(rotResult)
}

func DecryptBGV(ccLoc, skLoc, cipherloc string, vectorSize int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	skLocC := C.CString(skLoc)
	defer C.free(unsafe.Pointer(skLocC))
	cloc := C.CString(cipherloc)
	defer C.free(unsafe.Pointer(cloc))

	// 调用 Rot 函数
	decryptResult := C.DecryptBGV(ccLocC, C.size_t(len(ccLoc)),
		skLocC, C.size_t(len(skLoc)),
		cloc, C.size_t(len(cipherloc)),
		C.int(vectorSize)) // vectorSize
	defer C.FreeString(decryptResult)	
	
	return C.GoString(decryptResult)
}

func KeyGenBFV(multiplicativeDepth, plaintextModulus int) []string {
	cStr := C.KeyGenBFV(C.int(multiplicativeDepth), C.int(plaintextModulus))
	defer C.FreeString(cStr) // 确保释放内存
	// 转换为 Go 字符串
	goResult := C.GoString(cStr)	
	// 使用分隔符解析
	stringsList := strings.Split(goResult, ";")
	fmt.Println("Received strings:", stringsList)
	return stringsList
}

func EncryptBFV(ccLoc, pkLoc string, data []int, id int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))
	pkLocC := C.CString(pkLoc)
	defer C.free(unsafe.Pointer(pkLocC))

	dataC := (*C.int64_t)(unsafe.Pointer(&data[0])) // 转换为 C 指针

	resultENC := C.EncryptBFV(
		ccLocC,
		C.size_t(len(ccLoc)),
		pkLocC,
		C.size_t(len(pkLoc)),
		dataC,
		C.size_t(len(data)),
		C.int(id),
	)
	defer C.FreeString(resultENC)
	resultEncGO := C.GoString(resultENC)
	return resultEncGO
}

func AddBFV(ccLoc, c1Loc, c2Loc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))
	c2LocC := C.CString(c2Loc)
	defer C.free(unsafe.Pointer(c2LocC))
	
	// 调用 Add 函数
	addResult := C.AddBFV(ccLocC, C.size_t(len(ccLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		c2LocC, C.size_t(len(c2Loc)))
	defer C.FreeString(addResult)
	
	return C.GoString(addResult)
}

func MulBFV(ccLoc, multKLoc, c1Loc, c2Loc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	multkLocC := C.CString(multKLoc)
	defer C.free(unsafe.Pointer(multkLocC))	
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))
	c2LocC := C.CString(c2Loc)
	defer C.free(unsafe.Pointer(c2LocC))
	
	// 调用 Mul 函数
	mulResult := C.MulBFV(ccLocC, C.size_t(len(ccLoc)),
		multkLocC, C.size_t(len(multKLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		c2LocC, C.size_t(len(c2Loc)))
	defer C.FreeString(mulResult)
	
	return C.GoString(mulResult)
}

func RelinearizeBFV(ccLoc, multKLoc, cLoc string) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	multkLocC := C.CString(multKLoc)
	defer C.free(unsafe.Pointer(multkLocC))	
	cLocC := C.CString(cLoc)
	defer C.free(unsafe.Pointer(cLocC))

	
	// 调用 Mul 函数
	reLineResult := C.RelinearizeBFV(ccLocC, C.size_t(len(ccLoc)),
		multkLocC, C.size_t(len(multKLoc)),
		cLocC, C.size_t(len(cLoc)))
	defer C.FreeString(reLineResult)
	
	return C.GoString(reLineResult)
}

func RotBFV(ccLoc, rotkLoc, c1Loc string, index int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	rotkLocC := C.CString(rotkLoc)
	defer C.free(unsafe.Pointer(rotkLocC))
	c1LocC := C.CString(c1Loc)
	defer C.free(unsafe.Pointer(c1LocC))

	// 调用 Rot 函数
	rotResult := C.RotBFV(ccLocC, C.size_t(len(ccLoc)),
		rotkLocC, C.size_t(len(rotkLoc)),
		c1LocC, C.size_t(len(c1Loc)),
		C.int32_t(index)) // index
	defer C.FreeString(rotResult)
	
	return C.GoString(rotResult)
}

func DecryptBFV(ccLoc, skLoc, cipherloc string, vectorSize int) string {
	ccLocC := C.CString(ccLoc)
	defer C.free(unsafe.Pointer(ccLocC))	
	skLocC := C.CString(skLoc)
	defer C.free(unsafe.Pointer(skLocC))
	cloc := C.CString(cipherloc)
	defer C.free(unsafe.Pointer(cloc))

	// 调用 Rot 函数
	decryptResult := C.DecryptBFV(ccLocC, C.size_t(len(ccLoc)),
		skLocC, C.size_t(len(skLoc)),
		cloc, C.size_t(len(cipherloc)),
		C.int(vectorSize)) // vectorSize
	defer C.FreeString(decryptResult)	
	
	return C.GoString(decryptResult)
}