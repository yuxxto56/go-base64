package base

import (
	"encoding/base64"
	"fmt"
	"log"
)

//定义结构体
type Base64 struct {
    Str string
}

//base64加密
func (b *Base64) Encode() string{
	//转换成byte类型
	strB := []byte(b.Str)
	return base64.StdEncoding.EncodeToString(strB)
}

//basr64解密
func (b *Base64) Decode()(string,error){
	dec,err := base64.StdEncoding.DecodeString(b.Str)
	var deStr string;
	if err != nil{
		log.Fatal(fmt.Sprintf("DecodeString Error:%s",err.Error()))
	}
	deStr = string(dec)
	return deStr,nil
}

func init(){
	//初始化
}