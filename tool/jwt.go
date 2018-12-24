package tool

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var rsaPublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0Y+HSS6nEk34VfpQyxBk
/8WrcY7tHXYQKZ1mlyjhBn6QmcF6An6fUU/2BxyzyivnPB/Aku9NuJSgUKN+ccKD
0EoHw+3c+9qj2hckTCphnVsOMmYltDHXYEbSRy6k+lXJZNWTpIoTwSW63HANbwAZ
/ZV6wVXIpHtPjRfKD8OYraFx7v56AfMnA+47oQCnJ+MjGJfXqRS1qEPliIpODwJn
vQCRUdRRyj1dTyB5TnMzdzle+Bp70D40ghBfJQlRcyqrUEF9RcS/K21Uo9EKZN0r
d7G8+bv9hbJu3qZCoGddcBuIowQSSVctdDQxcBFmm34tUiJX24b5vGZT8WJCXdiZ
8QIDAQAB
-----END PUBLIC KEY-----
`)
var rsaPrivateKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDRj4dJLqcSTfhV
+lDLEGT/xatxju0ddhApnWaXKOEGfpCZwXoCfp9RT/YHHLPKK+c8H8CS7024lKBQ
o35xwoPQSgfD7dz72qPaFyRMKmGdWw4yZiW0MddgRtJHLqT6Vclk1ZOkihPBJbrc
cA1vABn9lXrBVcike0+NF8oPw5itoXHu/noB8ycD7juhAKcn4yMYl9epFLWoQ+WI
ik4PAme9AJFR1FHKPV1PIHlOczN3OV74GnvQPjSCEF8lCVFzKqtQQX1FxL8rbVSj
0Qpk3St3sbz5u/2Fsm7epkKgZ11wG4ijBBJJVy10NDFwEWabfi1SIlfbhvm8ZlPx
YkJd2JnxAgMBAAECggEBAIpwlkAbaNXn2nn6LqunvgFWdcdEU5LV7yF+0UplABuJ
Oz/IGDFTsy5Fc9QI+yFBVbZbh53KJ38b+Lis+ZhvlyLf9TdTiVx8eShXh2wlx2zr
vXpADwJ2tsDsqHfL8cNLLwvFKI9XhGth4ItGHynIWGM0/+5/HFjE+a3NHHjDj7aL
siH82wirKhjHHCiDg8uZsmPuacbB82fnrCwasgrnyIGk4pnNEB1URNP2peiul5KH
VUrtDzcAxUpeiWWOBXIOF9HK5q/UsLjY7Ki+ryzOfb21hLk8WIkKcbfsih69HOtd
urJFTqy0EUQn1DAMqMcCZ9KuY5rfZF1vCknk/Q/UggECgYEA7BxE9EeepAo3Ap15
5aJ+d9KUWFuM7H17maV8kDdADYb7/xjGBGKLdWy5RWUw89QTyMSaBvJvBRi/gqpu
ZNZmt9djBmWLPLdpQ6zpd+HtdmMdoWqRd4NnTpC1t+dWt+AWPte0THmnhtEIOyVb
sk9/0F2tSKcSv61B0x3yEEpqyBECgYEA4za2WNV2nmZh1ym81W3f5/7YzUIAd7Sk
SUlcH1ZircmjOCi58vhyavXvxjl4dO4ETc7HaqqRaUG2TgMt7OnZWcW1hr9sZhp8
B9M2TywVWggyAnvmwb6z9TP2bk9GigRxOO43L+UcNNqLsjkVkk+zLmhqEk09skgf
fyKd1/hNk+ECgYAoOHiWlEy17PeJ/oFxWMjqaHjUxGOxNX8EoWgiuQa0RzZOMk4p
GRBgpHC7HY3FO29TWApU/J9k4t5cx+OeUdeVoZ9Ay7N4WFYaLS5oOZhV/nRhtuGo
1j2W6JfMCCCBmi6v7dbM4Dxlft5EFQ0Vxu0i9ZI7Ohq7hCweZo+BruuZwQKBgGRY
ElnTFArKbmVGooWje6IKc5bakjhLh85x3oRcc0IK7dFscwqx3F4OdTEWynLhdvch
3gHmiEdy5N85Gjwp3np+DTgQLAA/rA0Mo0x9zIQxYFymeXeREAq3QxreWem0Ioig
87BsP3O8HLu96B/woNJa/JVOlXO51GciU5FAsIohAoGAO2xlTrs2PJcClFZ887l/
UuNDDwgvQNTkdLZ530165guvwfDJfSF4QdYS1O3mltzmwBAUNV/h+K8alKhh8nVK
6KFFch9SyOz9pR+Jq2gQ5hEdrN2xpbVGFoWR/D22LnSismM+8k8LUxgC0mAmieq1
clhkKRoA92JhB8fr7fCjAd4=
-----END PRIVATE KEY-----
`)
type JwtToken struct {
	Token string

}
func EncodeToken(id int64) (string,int32,error) {
	exp:=time.Now().Unix()+7*24*3600
	token:=jwt.NewWithClaims(jwt.SigningMethodRS256,jwt.MapClaims{
		"iss":"loopyun.com",
		"exp":exp,
		"user_id":id,

	})
	PrivateKey,_:=jwt.ParseRSAPrivateKeyFromPEM(rsaPrivateKey)

	ss,err:=token.SignedString(PrivateKey)
	if err!=nil{
		return "",0,err
	}
	return ss,int32(exp),nil

}
