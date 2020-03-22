package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecrectStr string = "go-mongo-admin"

func CreateToken(name string, password string) (tokens string, err error) {
	//自定义claim
	claim := jwt.MapClaims{
		"name":     name,
		"password": password,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 签名后的token格式说明
	// 示例：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzUyODUyNzUsImlkIjoxLCJuYmYiOjE1NzUyODUyNzUsInVzZXJuYW1lIjoicGVuZ2oifQ.bDe8UZYLxvmrK7gHcuK8TrlnoiMsIm3Jo_f0-YYle7E
	// 使用符号.，被分割成了三段
	// 第一段base64解码之后：{"alg":"HS256","typ":"JWT"}
	// 第二段base64解码之后：{"iat":1575285275,"id":1,"nbf":1575285275,"username":"pengj"}，是原始的数据。
	// 第三段是使用SigningMethodHS256加密之后的文本
	tokens, err = token.SignedString([]byte(SecrectStr))
	return
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (i interface{}, e error) {
		return []byte(SecrectStr), nil
	}
}

func ParseToken(tokens string) (name string, password string, err error) {
	token, err := jwt.Parse(tokens, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	name = claim["name"].(string)
	password = claim["password"].(string)
	return
}

//func main() {
//	user := UserInfo{ID: 1, Username: "pengj"}
//	tokenStr, err := CreateToken(&user)
//	if err != nil {
//		panic(err)
//	}
//	log.Println("tokenStr:", tokenStr)
//	u, err := ParseToken(tokenStr)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(u.ID, u.Username)
//}
//
