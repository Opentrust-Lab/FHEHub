package api

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"fhehub/sdk"
)

func KeyxxInit(m, n, q, p int) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	sdk.InitSDKPool()
	sdk.AddSDK("default", "./sdk/", "libkeyxx.core.so", m, n, q, p)
	fmt.Println("KeyxxAPI V0.1.0.202407062321 initing...")
}

func KeyxxSM3(userKey, input string) string {
	res := sdk.SM3(userKey, input)
	return res
}

func KeyxxGenSKB(userKey string, m, n, q, p int, filename string) {
	sdk.GenSKB(userKey, m, n, q, p, filename)
}

func KeyxxGenPKB(userKey, skbfile, filename string) {
	sdk.GenPKB(userKey, skbfile, filename)
}

func KeyxxGenDictB(userKey, skbfile, filename string, delta float64) {
	sdk.GenDictB(userKey, skbfile, filename, delta)
}

func KeyxxLoadPrivKey(userKey, skbfile string) {
	sdk.LoadPrivKey(userKey, skbfile)
}

func KeyxxLoadPubKey(userKey, pkbfile string) {
	sdk.LoadPrivKey(userKey, pkbfile)
}

func KeyxxLoadDictionary(userKey, dictbfile string) {
	sdk.LoadPrivKey(userKey, dictbfile)
}

func KeyxxEncrypt(userKey, input string) string {
	res := ""
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		val, err := strconv.ParseFloat(input, 64)
		if err != nil {
			res = sdk.EncString(userKey, input)
		} else {
			res = sdk.EncDouble(userKey, val)
		}
	} else {
		res = sdk.EncInt(userKey, int(val))
	}
	return res
}

func KeyxxEncryptPublic(userKey, input string) string {
	res := ""
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		val, err := strconv.ParseFloat(input, 64)
		if err != nil {
			res = sdk.EncPubString(userKey, input)
		} else {
			res = sdk.EncPubDouble(userKey, val)
		}
	} else {
		res = sdk.EncPubInt(userKey, int(val))	
	}
	return res
}

func KeyxxDecrypt(userKey, input string) string {
	res := sdk.Decrypt(userKey, input)
	return res
}

func KeyxxEncryptBinary(userKey, input, length string) string {
	data, _ := strconv.Atoi(input)
	l, _ := strconv.Atoi(length)
	res := sdk.EncBinary(userKey, data, l)
	return res
}

func KeyxxEncryptPublicBinary(userKey, input, length string) string {
	data, _ := strconv.Atoi(input)
	l, _ := strconv.Atoi(length)
	res := sdk.EncPubBinary(userKey, data, l)
	return res
}

func KeyxxDecryptBinary(userKey, input string) string {
	res := sdk.DecBinary(userKey, input)
	return strconv.Itoa(res)
}

func KeyxxEncryptString(userKey, input string) string {
	res := sdk.EncString(userKey, input)
	return res
}

func KeyxxEncryptPublicString(userKey, input string) string {
	res := sdk.EncPubString(userKey, input)
	return res
}

func KeyxxDecryptString(userKey, input string) string {
	res := sdk.DecString(userKey, input)
	return res
}

func KeyxxGenSign(userKey, input string) string {
	res := sdk.GenSign(userKey, input)
	return res
}

func KeyxxVerifySign(userKey, input, sign string) bool {
	ret := sdk.VerifySign(userKey, input, sign)
	res := false
	if ret == 1 {
		res = true
	}
	return res
}

func KeyxxAddCipher(userKey, c1, c2 string) string {
	res := ""
	if (sdk.IsCipherInt(userKey, c1) + sdk.IsCipherInt(userKey, c2) == 0) {
		res = sdk.AddCipherInt(userKey, c1, c2)
	} else if (sdk.IsCipherFloat(userKey, c1) + sdk.IsCipherFloat(userKey, c2) == 0) {
		res = sdk.AddCipherFloat(userKey, c1, c2)
	} else if (sdk.IsCipherDouble(userKey, c1) + sdk.IsCipherDouble(userKey, c2) == 0) {
		res = sdk.AddCipherDouble(userKey, c1, c2)
	} else {
		res = "Cipher Type Error!"
	}
	return res
}

func KeyxxSubstractCipher(userKey, c1, c2 string) string {
	res := ""
	if (sdk.IsCipherInt(userKey, c1) + sdk.IsCipherInt(userKey, c2) == 0) {
		res = sdk.SubCipherInt(userKey, c1, c2)
	} else if (sdk.IsCipherFloat(userKey, c1) + sdk.IsCipherFloat(userKey, c2) == 0) {
		res = sdk.SubCipherFloat(userKey, c1, c2)
	} else if (sdk.IsCipherDouble(userKey, c1) + sdk.IsCipherDouble(userKey, c2) == 0) {
		res = sdk.SubCipherDouble(userKey, c1, c2)
	} else {
		res = "Cipher Type Error!"
	}
	return res
}

func KeyxxMultiplyCipher(userKey, c1, c2 string) string {
	res := ""
	if (sdk.IsCipherInt(userKey, c1) + sdk.IsCipherInt(userKey, c2) == 0) {
		res = sdk.MulCipherInt(userKey, c1, c2)
	} else if (sdk.IsCipherFloat(userKey, c1) + sdk.IsCipherFloat(userKey, c2) == 0) {
		res = sdk.MulCipherFloat(userKey, c1, c2)
	} else if (sdk.IsCipherDouble(userKey, c1) + sdk.IsCipherDouble(userKey, c2) == 0) {
		res = sdk.MulCipherDouble(userKey, c1, c2)
	} else {
		res = "Cipher Type Error!"
	}
	return res
}

func KeyxxDivideCipher(userKey, c1, c2 string) string {
	res := ""
	if (sdk.IsCipherInt(userKey, c1) + sdk.IsCipherInt(userKey, c2) == 0) {
		res = "Cipher1 is CipherInt, but Divide is not supported!"
	} else if (sdk.IsCipherFloat(userKey, c1) + sdk.IsCipherFloat(userKey, c2) == 0) {
		res = sdk.DivCipherFloat(userKey, c1, c2)
	} else if (sdk.IsCipherDouble(userKey, c1) + sdk.IsCipherDouble(userKey, c2) == 0) {
		res = sdk.DivCipherDouble(userKey, c1, c2)
	} else {
		res = "Cipher Type Error!"
	}
	return res
}

func KeyxxXORCipher(userKey, c1, c2 string) string {
	res := sdk.XORCipherBinary(userKey, c1, c2)
	return res
}

func KeyxxANDCipher(userKey, c1, c2 string) string {
	res := sdk.ANDCipherBinary(userKey, c1, c2)
	return res
}

func KeyxxORCipher(userKey, c1, c2 string) string {
	res := sdk.ORCipherBinary(userKey, c1, c2)
	return res
}

func KeyxxNOTCipher(userKey, c1 string) string {
	res := sdk.NOTCipherBinary(userKey, c1)
	return res
}

func KeyxxShiftLeft(userKey, c1, bias string) string {
	bias_in, _ := strconv.Atoi(bias)
	res := sdk.ShiftLeftCipherBinary(userKey, c1, bias_in)
	return res
}

func KeyxxShiftRight(userKey, c1, bias string) string {
	bias_in, _ := strconv.Atoi(bias)
	res := sdk.ShiftRightCipherBinary(userKey, c1, bias_in)
	return res
}

func KeyxxCompare(userKey, c1, c2 string) int {
	res := 0
	if (sdk.IsCipherInt(userKey, c1) + sdk.IsCipherInt(userKey, c2) == 0) {
		res = sdk.CompareCipherInt(userKey, c1, c2)
	} else if (sdk.IsCipherFloat(userKey, c1) + sdk.IsCipherFloat(userKey, c2) == 0) {
		res = sdk.CompareCipherFloat(userKey, c1, c2)
	} else if (sdk.IsCipherDouble(userKey, c1) + sdk.IsCipherDouble(userKey, c2) == 0) {
		res = sdk.CompareCipherDouble(userKey, c1, c2)
	} else if (sdk.IsCipherString(userKey, c1) + sdk.IsCipherString(userKey, c2) == 0) {
		res = sdk.CompareCipherString(userKey, c1, c2)
	}
	return res
}

func KeyxxABSCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ABSCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ABSCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxPowerCipher(userKey, c1, nstr, mstr string) string {
	res := ""
	n, err1 := strconv.Atoi(nstr)
	m, err2 := strconv.Atoi(mstr)
	p := 0.0
	if err1 != nil {
		p, err1 = strconv.ParseFloat(nstr, 64)
		if err1 != nil {
			res = "Power Number Wrong!"
		}
	}
	if err2 != nil {
		m = 1
	}
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		if m == 1 && p == 0.0 {
			res = sdk.PowCipherFloat(userKey, c1, n)
		} else if p != 0.0 {
			m, _ := strconv.ParseFloat(mstr, 64)
			p = p / m
			res = sdk.PowCipherFloat_RealOrder(userKey, c1, p)
		} else {
			res = sdk.PowCipherFloat_FractionOrder(userKey, c1, n, m)
		}
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		if m == 1 && p == 0.0 {
			res = sdk.PowCipherDouble(userKey, c1, n)
		} else if p != 0.0 {
			m, _ := strconv.ParseFloat(mstr, 64)
			p = p / m
			res = sdk.PowCipherDouble_RealOrder(userKey, c1, p)
		} else {
			res = sdk.PowCipherDouble_FractionOrder(userKey, c1, n, m)
		}
	}
	return res
}

func KeyxxSqrtCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.SqrtCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.SqrtCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxLogCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.LogCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.LogCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxExpCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ExpCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ExpCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxSinCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.SinCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.SinCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxCosCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.CosCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.CosCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxTanCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.TanCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.TanCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxAsinCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ArcsinCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ArcsinCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxAcosCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ArccosCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ArccosCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxAtanCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ArctanCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ArctanCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxSinhCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.SinhCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.SinhCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxCoshCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.CoshCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.CoshCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxTanhCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.TanhCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.TanhCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxAsinhCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ArcsinhCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ArcsinhCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxAcoshCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ArccoshCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ArccoshCipherDouble(userKey, c1)
	}
	return res
}

func KeyxxAtanhCipher(userKey, c1 string) string {
	res := ""
	if (sdk.IsCipherFloat(userKey, c1) == 0) {
		res = sdk.ArctanhCipherFloat(userKey, c1)
	} else if (sdk.IsCipherDouble(userKey, c1) == 0) {
		res = sdk.ArctanhCipherDouble(userKey, c1)
	}
	return res
}

func ConcatString(userKey, data1, data2 string) string {
	res := sdk.ConcatString(userKey, data1, data2)
	return res
}

func Substring(userKey, data, start, end string) string {
	start_in, _ := strconv.Atoi(start)
	end_in, _ := strconv.Atoi(end)
    res := sdk.Substring(userKey, data, start_in, end_in)
	return res
}
