package main

import (
	"fmt"
	"time"
	"os"
	"io"
	"math/rand"
	// "io/ioutil"
	"net"
    "net/http"
	"strconv"
	"fhehub/server"
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/gin-gonic/gin"
)

var port string = ":8878"
var db *leveldb.DB

func main() {
	rand.Seed(time.Now().UnixNano())
    fmt.Println("服务启动..")

	server.InitServer(db)

	// gin.DisableConsoleColor()

    // Logging to a file.
    f, _ := os.Create("log/server_gin." + strconv.Itoa(int(time.Now().UnixNano())) + ".log")
    gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	// r.LoadHTMLGlob("site/pages/template/*")
	// r.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/site/login.html")
	// })
	// r.StaticFS("/site", http.Dir("./site/"))
	server.ProcessRequest(r, db)
	fmt.Println("ProcessRequest")

	server := &http.Server{Addr: port, Handler: r}
	ln, err := net.Listen("tcp4", port)
	if err != nil {
		panic(err)
	}
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}
	server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
	
    // fmt.Println("服务启动成功。", port)
}