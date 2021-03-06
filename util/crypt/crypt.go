package crypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// 对字符串进行sha1 计算
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// 对数据进行md5计算
func MD5(byteMessage []byte) string {
	h := md5.New()
	h.Write(byteMessage)
	return hex.EncodeToString(h.Sum(nil))
}

func HamSha1(data string, key []byte) string {
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}
