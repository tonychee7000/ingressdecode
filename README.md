Ingress Decoder
============

Ingress 多功能解码器

支持 atbash hexatbash（十六进制 atbash） rot13 morse 和 fence（栅栏密码，列出各种可能）

其中 morse 支持:
* tomorse（转为 Morse）
* frommorse（从 Morse 转换）
* swapmorse（点划互换）

用法：
```bash
ingressdecoder -i passcode -m method
  passcode 输入可疑的passcode
  method 解码方法
```
