package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//使用证书池，将自己签发的证书加进去
	pool := x509.NewCertPool()
	myCaPath := "../cert/ca.crt"
	caCrt, _ := ioutil.ReadFile(myCaPath)
	pool.AppendCertsFromPEM(caCrt)

	//创建httpClient，并指定证书池
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            pool,
				InsecureSkipVerify: false, //这次不跳过验证
			},
		},
	}
	//发起GET请求
	r, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	log.Println(string(body))

	//发起POST请求
	r1, err := client.Post("https://localhost:8080/reg", "text/plain", strings.NewReader("hello, i'm client"))
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	log.Println(r1.StatusCode)
}