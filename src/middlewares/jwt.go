package middlewares

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWY() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// authHeader := ctx.GetHeader("Authorization")
		//从URL中获取username,从Form中获取token
		token := ctx.PostForm("token")
		username := ctx.Query("username")
		// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		if token == "" {
			ctx.JSON(401, map[string]string{
				"message": "请登陆",
			})
			ctx.Abort()
			return
		} else {
			// token := authHeader[7:]
			// fmt.Println(token)
			claims, err := ParseToken(token)
			if err != nil {
				ctx.JSON(401, map[string]string{
					"message": "无效Token！",
				})
				ctx.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				err = errors.New("Token已过期")
				ctx.JSON(401, map[string]string{
					"message": "Token已过期",
				})
				ctx.Abort()
				return
			}
			log.Println(claims.Username)
			log.Println(username)
			if claims.Username != username {
				ctx.JSON(401, map[string]string{
					"message": "身份认证失败",
				})
				ctx.Abort()
				return
			}
			log.Printf("用户:%s ，登陆成功", username)
		}
	}
}

func GetToken(username string) (string, error) {
	nowTime := time.Now()
	lifeTime := nowTime.Add(1 * time.Hour)

	clamis := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: lifeTime.Unix(),
		},
	}
	tokenClamis := jwt.NewWithClaims(jwt.SigningMethodHS512, clamis)
	token, err := tokenClamis.SignedString(secretKey)
	return token, err
}
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Println("parse token faild!", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil

	}
	return nil, err
}
