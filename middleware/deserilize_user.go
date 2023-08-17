package middleware

import (
	"net/http"
	"strings"

	"cinema.com/data/response"
	"cinema.com/repository"
	"cinema.com/utils"
	"github.com/gin-gonic/gin"
)

func DeserilizeUser(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) >= 2 && fields[0] == "Bearer" {
			token = fields[1]
		}

		unauthorized := response.ErorrResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		forbiddenResponse := response.ErorrResponse{
			Code:   http.StatusForbidden,
			Status: "Forbidden",
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized)
			ctx.Abort()
			return
		}

		sub, err := utils.ValidateToken(token, "secret")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized)
			ctx.Abort()
			return
		}

		subMap, ok := sub.(map[string]interface{})
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized)
			ctx.Abort()
			return
		}

		id, ok := subMap["id"]
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorized)
			ctx.Abort()
			return
		}

		result, err := userRepository.FindById(int(id.(float64)))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, forbiddenResponse)
			ctx.Abort()
			return
		}

		ctx.Set("currentUser", result)
		ctx.Next()
	}
}
