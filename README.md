# go-base64
### 使用
```
   package  main
   
   import (
   	"base64/base"
   	"fmt"
   )

   func main(){
       enStr := "我爱中国 l love you"
       b := &base.Base64{enStr}
       fmt.Println(enStr+"加密后的字符串为:",b.Encode())
       b.Str = "5oiR54ix5Lit5Zu9IGwgbG92ZSB5b3U="
       de,_:= b.Decode()
       fmt.Println(b.Str+"解密后的字符串为:",de)
   }

```