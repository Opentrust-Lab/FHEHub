package server

import (
	"fmt"
	"os"

	// "io"
	// "math"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"fhehub/api"
)

var filedir string = "file"
var dbfile string = "mngtdb"
var logdir string = "log"
var default_apikey string = "y93bjoQhsjSYSAVw19dTMFOKaIL139w7ZBFcZH8zKMQ="

var p int = 512
var m int = 2
var n int = 4
var q int = 2147483647
var delta float64 = 0.01

type JsonResult  struct{
    IsSuccess bool `json:"isSuccess"`
    Result interface{} `json:"result"`
    Log string `json:"log"`
}

func InitServer(db *leveldb.DB) {
    var err error
    db, err = leveldb.OpenFile(dbfile, nil)
    defer db.Close()
    if err != nil {
        fmt.Println(err)
    }
}

func ProcessRequest(engine *gin.Engine, db *leveldb.DB) {
    suc := true
    res := ""
    engine.POST("/init", func(c *gin.Context) {
        api.KeyxxInit(m, n, q, p)
        res = "init success."
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/tool/sm3", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxSM3(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/file/getFile", func(c *gin.Context) {
        filename := c.PostForm("filename")
        // fmt.Println(filename)
        file, _ := os.Open(filename);
        defer file.Close();
        fileHeader := make([]byte, 512)
        file.Read(fileHeader)
        fileStat, _ := file.Stat()
        c.Writer.Header().Set("Content-Disposition", "attachment; filename=" + filename)
        c.Writer.Header().Set("Content-Type", http.DetectContentType(fileHeader))
        c.Writer.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))
        file.Seek(0, 0)
        io.Copy(c.Writer, file)
    })
    engine.POST("/km/genSK", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        m_in, err := strconv.Atoi(c.PostForm("m"))
        if err != nil {
            m_in = m
        }
        n_in, err := strconv.Atoi(c.PostForm("n"))
        if err != nil {
            n_in = n
        }
        q_in, err := strconv.Atoi(c.PostForm("q"))
        if err != nil {
            q_in = q
        }
        p_in, err := strconv.Atoi(c.PostForm("p"))
        if err != nil {
            p_in = p
        }
        filename := c.PostForm("filename")
        res = filename
        api.KeyxxGenSKB(userKey, m_in, n_in, q_in, p_in, filename)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)
    })
    engine.POST("/km/genPK", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        skbfile := c.PostForm("skfile")
        filename := c.PostForm("filename")
        res = filename
        api.KeyxxGenPKB(userKey, skbfile, filename)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)
    })
    engine.POST("/km/genDict", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        skbfile := c.PostForm("skfile")
        filename := c.PostForm("filename")
        delta_in, err := strconv.ParseFloat(c.PostForm("delta"), 64)
        if err != nil {
            delta_in = delta
        }
        res = filename
        api.KeyxxGenDictB(userKey, skbfile, filename, delta_in)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)
    })
    engine.POST("/km/loadSK", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        skbfile := c.PostForm("skfile")
        res = skbfile
        api.KeyxxLoadPrivKey(userKey, skbfile)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)
    })
    engine.POST("/km/loadPK", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        pkbfile := c.PostForm("pkfile")
        res = pkbfile
        api.KeyxxLoadPubKey(userKey, pkbfile)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)
    })
    engine.POST("/km/loadDict", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        dictbfile := c.PostForm("dictfile")
        res = dictbfile
        api.KeyxxLoadDictionary(userKey, dictbfile)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)
    })
    engine.POST("/enc/encrypt", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxEncrypt(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)
    })
    engine.POST("/enc/encryptBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        inputList := c.PostForm("inputList")
        inputs := make([]string, 0)
        json.Unmarshal([]byte(inputList), &inputs)
        ret := make([]string, len(inputs))
        for i := 0; i < len(inputs); i++ {
            resi := api.KeyxxEncrypt(userKey, inputs[i])
            ret[i] = resi
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptPublic", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxEncryptPublic(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptPublicBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        inputList := c.PostForm("inputList")
        inputs := make([]string, 0)
        json.Unmarshal([]byte(inputList), &inputs)
        ret := make([]string, len(inputs))
        for i := 0; i < len(inputs); i++ {
            resi := api.KeyxxEncryptPublic(userKey, inputs[i])
            ret[i] = resi
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/decrypt", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxDecrypt(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/decryptBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        inputList := c.PostForm("inputList")
        inputs := make([]string, 0)
        json.Unmarshal([]byte(inputList), &inputs)
        ret := make([]string, len(inputs))
        for i := 0; i < len(inputs); i++ {
            resi := api.KeyxxDecrypt(userKey, inputs[i])
            ret[i] = resi
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptBinary", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        length := c.PostForm("length")
        res = api.KeyxxEncryptBinary(userKey, input, length)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptBinaryPublic", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        length := c.PostForm("length")
        res = api.KeyxxEncryptPublicBinary(userKey, input, length)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/decryptBinary", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxDecryptBinary(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptString", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxEncryptString(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptStringBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        inputList := c.PostForm("inputList")
        inputs := make([]string, 0)
        json.Unmarshal([]byte(inputList), &inputs)
        ret := make([]string, len(inputs))
        for i := 0; i < len(inputs); i++ {
            resi := api.KeyxxEncryptString(userKey, inputs[i])
            ret[i] = resi
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptStringPublic", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxEncryptPublicString(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/encryptStringPublicBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        inputList := c.PostForm("inputList")
        inputs := make([]string, 0)
        json.Unmarshal([]byte(inputList), &inputs)
        ret := make([]string, len(inputs))
        for i := 0; i < len(inputs); i++ {
            resi := api.KeyxxEncryptPublicString(userKey, inputs[i])
            ret[i] = resi
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/decryptString", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxDecryptString(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/decryptStringBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        inputList := c.PostForm("inputList")
        inputs := make([]string, 0)
        json.Unmarshal([]byte(inputList), &inputs)
        ret := make([]string, len(inputs))
        for i := 0; i < len(inputs); i++ {
            resi := api.KeyxxDecryptString(userKey, inputs[i])
            ret[i] = resi
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/genSign", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        res = api.KeyxxGenSign(userKey, input)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/enc/verifySign", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        sign := c.PostForm("sign")
        res = strconv.FormatBool(api.KeyxxVerifySign(userKey, input, sign))
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/add", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.KeyxxAddCipher(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/addBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        c2 := c.PostForm("c2list")
        cs1 := make([]string, 0)
        cs2 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        json.Unmarshal([]byte(c2), &cs2)
        ret := make([]string, 0)
        minlen := len(cs1)
        if len(cs1) > len(cs2) {
            minlen = len(cs2)
        }
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxAddCipher(userKey, cs1[i], cs2[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/substract", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.KeyxxSubstractCipher(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/substractBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        c2 := c.PostForm("c2list")
        cs1 := make([]string, 0)
        cs2 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        json.Unmarshal([]byte(c2), &cs2)
        ret := make([]string, 0)
        minlen := len(cs1)
        if len(cs1) > len(cs2) {
            minlen = len(cs2)
        }
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxSubstractCipher(userKey, cs1[i], cs2[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/multiply", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.KeyxxMultiplyCipher(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/multiplyBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        c2 := c.PostForm("c2list")
        cs1 := make([]string, 0)
        cs2 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        json.Unmarshal([]byte(c2), &cs2)
        ret := make([]string, 0)
        minlen := len(cs1)
        if len(cs1) > len(cs2) {
            minlen = len(cs2)
        }
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxMultiplyCipher(userKey, cs1[i], cs2[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/divide", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.KeyxxDivideCipher(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/divideBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        c2 := c.PostForm("c2list")
        cs1 := make([]string, 0)
        cs2 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        json.Unmarshal([]byte(c2), &cs2)
        ret := make([]string, 0)
        minlen := len(cs1)
        if len(cs1) > len(cs2) {
            minlen = len(cs2)
        }
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxDivideCipher(userKey, cs1[i], cs2[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/xor", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.KeyxxXORCipher(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/and", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.KeyxxANDCipher(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/or", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.KeyxxORCipher(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/not", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxNOTCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/shiftLeft", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        bias := c.PostForm("bias")
        res = api.KeyxxShiftLeft(userKey, c1, bias)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/shiftRight", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        bias := c.PostForm("bias")
        res = api.KeyxxShiftRight(userKey, c1, bias)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)       
    })
    engine.POST("/opt/compare", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = strconv.Itoa(api.KeyxxCompare(userKey, c1, c2))
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/compareBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        c2 := c.PostForm("c2list")
        cs1 := make([]string, 0)
        cs2 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        json.Unmarshal([]byte(c2), &cs2)
        ret := make([]string, 0)
        minlen := len(cs1)
        if len(cs1) > len(cs2) {
            minlen = len(cs2)
        }
        for i := 0; i < minlen; i++ {
            resi := strconv.Itoa(api.KeyxxCompare(userKey, cs1[i], cs2[i]))
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/abs", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxABSCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/absBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxABSCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/power", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        nstr := c.PostForm("n")
        mstr := c.PostForm("m")
        res = api.KeyxxPowerCipher(userKey, c1, nstr, mstr)
        // fmt.Println(res)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/powerBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        nstr := c.PostForm("n")
        mstr := c.PostForm("m")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxPowerCipher(userKey, cs1[i], nstr, mstr)
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        // fmt.Println(res)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/sqrt", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxSqrtCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/sqrtBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxSqrtCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/log", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxLogCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/logBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxLogCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/exp", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxExpCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/expBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxExpCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/sin", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxSinCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/sinBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxSinCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/cos", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxCosCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/cosBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxCosCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/tan", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxTanCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/tanBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxTanCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/asin", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxAsinCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/asinBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxAsinCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/acos", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxAcosCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/acosBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxAcosCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/atan", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxAtanCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/atanBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxAtanCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/sinh", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxSinhCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/sinhBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxSinhCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/cosh", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxCoshCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/coshBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxCoshCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/tanh", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxTanhCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/tanhBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxTanhCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/asinh", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxAsinhCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/asinhBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxAsinhCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/acosh", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxAcoshCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/acoshBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxAcoshCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/atanh", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        res = api.KeyxxAtanhCipher(userKey, c1)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/atanhBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        cs1 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        ret := make([]string, 0)
        minlen := len(cs1)
        for i := 0; i < minlen; i++ {
            resi := api.KeyxxAtanhCipher(userKey, cs1[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/concat", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1")
        c2 := c.PostForm("c2")
        res = api.ConcatString(userKey, c1, c2)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/concatBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        c1 := c.PostForm("c1list")
        c2 := c.PostForm("c2list")
        cs1 := make([]string, 0)
        cs2 := make([]string, 0)
        json.Unmarshal([]byte(c1), &cs1)
        json.Unmarshal([]byte(c2), &cs2)
        ret := make([]string, 0)
        minlen := len(cs1)
        if len(cs1) > len(cs2) {
            minlen = len(cs2)
        }
        for i := 0; i < minlen; i++ {
            resi := api.ConcatString(userKey, cs1[i], cs2[i])
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/substring", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        start := c.PostForm("start")
        end := c.PostForm("end")
        res = api.Substring(userKey, input, start, end)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = res
        c.JSON(200, jsonResult)        
    })
    engine.POST("/opt/substringBatch", func(c *gin.Context) {
        userKey := c.PostForm("userKey")
        input := c.PostForm("input")
        start := c.PostForm("start")
        end := c.PostForm("end")
        inputs := make([]string, 0)
        json.Unmarshal([]byte(input), &inputs)
        ret := make([]string, 0)
        minlen := len(inputs)
        for i := 0; i < minlen; i++ {
            resi := api.Substring(userKey, inputs[i], start, end)
            ret = append(ret, resi)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/gen/randInt", func(c *gin.Context) {
        min := c.PostForm("min")
        max := c.PostForm("max")
        length := c.PostForm("length")
        min_v, _ := strconv.Atoi(min)
        max_v, _ := strconv.Atoi(max)
        length_v, _ := strconv.Atoi(length)
        ret := make([]string, length_v)
        for i := 0; i < length_v; i++ {
            reti := min_v + rand.Intn(max_v)
            ret[i] = strconv.Itoa(reti)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/gen/randFloat", func(c *gin.Context) {
        min := c.PostForm("min")
        max := c.PostForm("max")
        length := c.PostForm("length")
        min_v, _ := strconv.ParseFloat(min, 64)
        max_v, _ := strconv.ParseFloat(max, 64)
        length_v, _ := strconv.Atoi(length)
        ret := make([]string, length_v)
        for i := 0; i < length_v; i++ {
            reti := min_v + rand.Float64() * (max_v - min_v)
            ret[i] = strconv.FormatFloat(reti, 'f', -1, 64)
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    engine.POST("/gen/randString", func(c *gin.Context) {
        strlen := c.PostForm("strlen")
        length := c.PostForm("length")
        strlen_v, _ := strconv.Atoi(strlen)
        length_v, _ := strconv.Atoi(length)
        ret := make([]string, length_v)
        for i := 0; i < length_v; i++ {
            ret[i] = api.KeyxxSM3("", strconv.FormatFloat(rand.Float64(), 'f', -1, 64))[:strlen_v]
        }
        res, _ := json.Marshal(ret)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        jsonResult.Result = string(res)
        c.JSON(200, jsonResult)        
    })
    // CKKS
    engine.POST("/ckks/genKey", func(c *gin.Context) {
        multDepth, _ := strconv.Atoi(c.PostForm("multDepth"))
        scaleModSize, _ := strconv.Atoi(c.PostForm("scaleModSize"))
        batchSize, _ := strconv.Atoi(c.PostForm("batchSize"))
        result := api.KeyGenCKKS(multDepth, scaleModSize, batchSize)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "cryptocontextLoc":  result[0],
            "pkLoc":             result[1],
            "skLoc":             result[2],
            "multKLoc":          result[3],
            "rotKLoc":           result[4],
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/ckks/encrypt", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        pkLoc := c.PostForm("pkLoc")
        id, _ := strconv.Atoi(c.PostForm("id"))
    
        // 接收 data，并解析成 []float64
        dataStr := c.PostForm("data") // e.g., "1.23,4.56,7.89"
        dataParts := strings.Split(dataStr, ",")
        var data []float64
        for _, part := range dataParts {
            value, err := strconv.ParseFloat(part, 64)
            if err == nil {
                data = append(data, value)
            }
        }
    
        // 调用 EncryptCKKS
        result := api.EncryptCKKS(ccLoc, pkLoc, data, id)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "ciphertextLoc":  result,
        }
        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/ckks/add", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        c1Loc := c.PostForm("c1Loc")
        c2Loc := c.PostForm("c2Loc")
        result := api.AddCKKS(ccLoc, c1Loc, c2Loc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "addResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/ckks/mul", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        multKLoc := c.PostForm("multKLoc")
        c1Loc := c.PostForm("c1Loc")
        c2Loc := c.PostForm("c2Loc")
        result := api.MulCKKS(ccLoc, multKLoc, c1Loc, c2Loc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "mulResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/ckks/relinearize", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        multKLoc := c.PostForm("multKLoc")
        cLoc := c.PostForm("cLoc")
        result := api.RelinearizeCKKS(ccLoc, multKLoc, cLoc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "relinearizeResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/ckks/rot", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        rotKLoc := c.PostForm("rotKLoc")
        cLoc := c.PostForm("cLoc")
        index, _ := strconv.Atoi(c.PostForm("index"))
        result := api.RotCKKS(ccLoc, rotKLoc, cLoc, index)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "rotResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/ckks/decrypt", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        skLoc := c.PostForm("skLoc")
        cLoc := c.PostForm("cLoc")
        vectorSize, _ := strconv.Atoi(c.PostForm("vectorSize"))
        result := api.DecryptCKKS(ccLoc, skLoc, cLoc, vectorSize)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "decryptResult":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    // BGV
    engine.POST("/bgv/genKey", func(c *gin.Context) {
        multiplicativeDepth, _ := strconv.Atoi(c.PostForm("multiplicativeDepth"))
        plaintextModulus, _ := strconv.Atoi(c.PostForm("plaintextModulus"))
        result := api.KeyGenBGV(multiplicativeDepth, plaintextModulus)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "cryptocontextLoc":  result[0],
            "pkLoc":             result[1],
            "skLoc":             result[2],
            "multKLoc":          result[3],
            "rotKLoc":           result[4],
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bgv/encrypt", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        pkLoc := c.PostForm("pkLoc")
        id, _ := strconv.Atoi(c.PostForm("id"))
    
        // 接收 data，并解析成 []float64
        dataStr := c.PostForm("data") // e.g., "1.23,4.56,7.89"
        dataParts := strings.Split(dataStr, ",")
        var data []int
        for _, part := range dataParts {
            value, err := strconv.Atoi(part) // 转换为 int
            if err == nil { 
                data = append(data, value)
            }
        }
    
        // 调用 EncryptBGV
        result := api.EncryptBGV(ccLoc, pkLoc, data, id)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "ciphertextLoc":  result,
        }
        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bgv/add", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        c1Loc := c.PostForm("c1Loc")
        c2Loc := c.PostForm("c2Loc")
        result := api.AddBGV(ccLoc, c1Loc, c2Loc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "addResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bgv/mul", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        multKLoc := c.PostForm("multKLoc")
        c1Loc := c.PostForm("c1Loc")
        c2Loc := c.PostForm("c2Loc")
        result := api.MulBGV(ccLoc, multKLoc, c1Loc, c2Loc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "mulResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bgv/relinearize", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        multKLoc := c.PostForm("multKLoc")
        cLoc := c.PostForm("cLoc")
        result := api.RelinearizeBGV(ccLoc, multKLoc, cLoc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "relinearizeResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bgv/rot", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        rotKLoc := c.PostForm("rotKLoc")
        cLoc := c.PostForm("cLoc")
        index, _ := strconv.Atoi(c.PostForm("index"))
        result := api.RotBGV(ccLoc, rotKLoc, cLoc, index)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "rotResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bgv/decrypt", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        skLoc := c.PostForm("skLoc")
        cLoc := c.PostForm("cLoc")
        vectorSize, _ := strconv.Atoi(c.PostForm("vectorSize"))
        result := api.DecryptBGV(ccLoc, skLoc, cLoc, vectorSize)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "decryptResult":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })    
    // BFV
    engine.POST("/bfv/genKey", func(c *gin.Context) {
        multiplicativeDepth, _ := strconv.Atoi(c.PostForm("multiplicativeDepth"))
        plaintextModulus, _ := strconv.Atoi(c.PostForm("plaintextModulus"))
        result := api.KeyGenBFV(multiplicativeDepth, plaintextModulus)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "cryptocontextLoc":  result[0],
            "pkLoc":             result[1],
            "skLoc":             result[2],
            "multKLoc":          result[3],
            "rotKLoc":           result[4],
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bfv/encrypt", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        pkLoc := c.PostForm("pkLoc")
        id, _ := strconv.Atoi(c.PostForm("id"))
    
        // 接收 data，并解析成 []int
        dataStr := c.PostForm("data") // e.g., "1.23,4.56,7.89"
        dataParts := strings.Split(dataStr, ",")
        var data []int
        for _, part := range dataParts {
            value, err := strconv.Atoi(part) // 转换为 int
            if err == nil { 
                data = append(data, value)
            }
        }
    
        // 调用 EncryptBFV
        result := api.EncryptBFV(ccLoc, pkLoc, data, id)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "ciphertextLoc":  result,
        }
        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bfv/add", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        c1Loc := c.PostForm("c1Loc")
        c2Loc := c.PostForm("c2Loc")
        result := api.AddBFV(ccLoc, c1Loc, c2Loc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "addResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bfv/mul", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        multKLoc := c.PostForm("multKLoc")
        c1Loc := c.PostForm("c1Loc")
        c2Loc := c.PostForm("c2Loc")
        result := api.MulBFV(ccLoc, multKLoc, c1Loc, c2Loc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "mulResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bfv/relinearize", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        multKLoc := c.PostForm("multKLoc")
        cLoc := c.PostForm("cLoc")
        result := api.RelinearizeBFV(ccLoc, multKLoc, cLoc)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "relinearizeResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bfv/rot", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        rotKLoc := c.PostForm("rotKLoc")
        cLoc := c.PostForm("cLoc")
        index, _ := strconv.Atoi(c.PostForm("index"))
        result := api.RotBFV(ccLoc, rotKLoc, cLoc, index)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "rotResultLoc":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })
    engine.POST("/bfv/decrypt", func(c *gin.Context) {
        ccLoc := c.PostForm("ccLoc")
        skLoc := c.PostForm("skLoc")
        cLoc := c.PostForm("cLoc")
        vectorSize, _ := strconv.Atoi(c.PostForm("vectorSize"))
        result := api.DecryptBFV(ccLoc, skLoc, cLoc, vectorSize)
        jsonResult := JsonResult {IsSuccess: suc, Result: "", Log: ""}
        // jsonResult := JsonResult {IsSuccess: suc, CryptocontextLoc: result[0], PkLoc: result[1],
        //                             SkLoc: result[2], multKLoc: result[3], rotKLoc: result[4]}
        jsonObj := map[string]interface{}{
            "decryptResult":  result,
        }

        // 将 map 转换为 JSON 字符串
        jsonResult.Result = jsonObj
        c.JSON(200, jsonResult)        
    })     
}

func GetFile(filename string) ([]byte, error) {
    file, err := os.OpenFile(filename, os.O_RDONLY, 0)
    defer file.Close()
    if err != nil {
        fmt.Println(err)
        return nil, err
    } else {
        content, _ := ioutil.ReadAll(file)
        return content, nil
    }
}
