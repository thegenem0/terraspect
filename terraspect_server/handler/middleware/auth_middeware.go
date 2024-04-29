package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/service"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func AuthMiddleware(s service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		err := c.ShouldBindHeader(&h)
		if err != nil {
			return
		}

		fmt.Println("ID_TOKEN: ", h.IDToken)

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			c.JSON(400, gin.H{
				"error": "Authorization header must be provided",
			})

			c.Abort()
			return
		}

		// validate ID token here
		_, err = s.VerifyToken(idTokenHeader[1])

		if err != nil {
			c.JSON(401, gin.H{
				"error": "Token verification failed",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
