package middleware

import (
	"fmt"
	"net/http" //for http.Status

	"github.com/gin-gonic/gin"
	helper "github.com/y0gesh02/go-jwt/helpers"
)

//only for getuser and getusers func
func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token") //getting token from user
		if clientToken == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":fmt.Sprintf("No Authorization header provided")})
			c.Abort() //abort operations
			return
		}

		claims, err := helper.ValidateToken(clientToken) //checking token
		if err !="" {
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid",claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}