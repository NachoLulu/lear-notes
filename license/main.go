package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

const (
	// 根证书信息
	rootCa = `./ca.crt`
	// 许可证证书路径
	license = `./license.crt`
)

// 根证书证书池
var roots *x509.CertPool

// 初始化证书池信息
func init() {
	bytes, err := ioutil.ReadFile(rootCa)
	if err != nil {
		panic(err)
	}
	roots = x509.NewCertPool()
	roots.AppendCertsFromPEM(bytes)
}

func main() {
	fmt.Println("CheckLicense:", checkLicense(license))
}

func checkLicense(filename string) bool {
	// 读取许可证证书
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// 解析证书
	block, _ := pem.Decode(bytes)
	if block == nil {
		panic(err)
	}

	// 转换为证书格式
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 验证有效性
	if _, err := cert.Verify(x509.VerifyOptions{
		Roots: roots,
	}); err != nil {
		return false
	}

	return true
}
