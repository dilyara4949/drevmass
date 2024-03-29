package middleware

import (
	// "fmt"
	// "fmt"
	"net/http"
	"strings"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/dilyara4949/drevmass/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

// func JwtAuthMiddleware(secret string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.Request.Header.Get("Authorization")
// 		t := strings.Split(authHeader, " ")
// 		if len(t) == 2 {
// 			authToken := t[1]
// 			authorized, err := tokenutil.IsAuthorized(authToken, secret)

// 			if authorized {
// 				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
// 				if err != nil {
// 					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
// 					c.Abort()
// 					return
// 				}
// 				c.Set("x-user-id", userID)
// 				c.Next()
// 				return
// 			}
// 			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not0 authorized"})
// 		c.Abort()
// 	}
// }

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
    
    return func(c *gin.Context) {
        authHeader := c.Request.Header.Get("Authorization")
        t := strings.Split(authHeader, " ")
        if len(t) == 2 {
            authToken := t[1]
			// fmt.Println(authToken)
            // fmt.Println(secret)
            authorized, err := tokenutil.IsAuthorized(authToken, secret)
            if err != nil {
				
				c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
				// c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "2"})
				c.Abort()
				return
			}
            if authorized {

                userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				// fmt.Println(userID)
                if err != nil {
                    c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					// c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "1"})
                    c.Abort()
                    return
                }
                c.Set("x-user-id", userID)
                c.Next()
                return
            }
            c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not0 authorized"})
            c.Abort()
            return
        }
        c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Authorization is required Header"})
        c.Abort()
    }
}
