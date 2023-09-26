package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func TypeVerify(c *gin.Context, typeofuser string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != typeofuser {
		err = errors.New("Unauthorised access")
		return err
	}
	return err
}

func UserVerify(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("user_id")

	if userType != "ADMIN" && uid != userId {
		err = errors.New("Unauthorised access")
		return err
	}

	err = TypeVerify(c, userType)
	return err
}
