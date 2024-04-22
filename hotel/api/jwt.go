package api

import (
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type myClaims struct {
	// 可根据需要自行添加字段
	UserID primitive.ObjectID `json:"userID"`
	Email  string             `json:"email"`
	jwt.StandardClaims
}

// 定义过期时间
const TokenExpirDuration = time.Hour * 24 * 3

var mySecret = []byte("little fan")

// GenToken 生成JWT
func GenToken(userID primitive.ObjectID, email string) (token string, err error) {
	// 创建一个我们自己的声明数据
	claims := myClaims{
		userID,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpirDuration).Unix(),
			Issuer:    "hotel-hsiaol1", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
	if err != nil {
		slog.Error("token failed", "err", err)
		return
	}
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token, err
}

// ParseToken 解析JWT
// ParseToken 解析JWT
func ParseToken(tokenString string) (*myClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
