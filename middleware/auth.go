package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		tokenValue, err := ctx.Cookie("session_token")
		contentType := ctx.GetHeader("Content-Type")
		if err != nil && contentType == "application/json" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			return
		} else if err != nil && contentType != "application/json" {
			ctx.AbortWithStatusJSON(http.StatusSeeOther, model.NewErrorResponse("token is missing"))
			return
		}

		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return model.JwtKey, nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse("token is invalid"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			return
		}

		b, _ := json.Marshal(claims)
		var customClaims model.Claims
		json.Unmarshal(b, &customClaims)

		ctx.Set("email", customClaims.Email)
		// TODO: answer here
	})
}
