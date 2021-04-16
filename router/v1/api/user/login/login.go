package login

import (
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/middlerware"
	"github.com/Peterliang233/debate/model"
	Service "github.com/Peterliang233/debate/service/v1/api/user/login"
	"github.com/Peterliang233/debate/service/v1/api/user/validate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


//登录
func Login(c *gin.Context) {
	var login model.Login
	err := c.ShouldBind(&login)
	msg, code := validate.Validate(&login)
	if code != errmsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg": msg,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrRequest,
			"detail": errmsg.CodeMsg[errmsg.ErrRequest],
		})
		return
	}
	StatusCode, code := Service.CheckLogin(&login)
	if code != errmsg.Success {
		c.JSON(StatusCode, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[code],
			},
		})
		return
	}
	var token string
	//获取token
	token, code = middlerware.GenerateToken(login.Email)
	if code == errmsg.Success {

		//将token暂时放到redis里面缓存
		_, err := dao.Conn.Do("SET", login.Email + "token", token)
		if err != nil {
			log.Println("set token error:", err)
		}
		_, err = dao.Conn.Do("expire", login.Email + "token", 7200)  //放到redis里面缓存分120钟
		if err != nil {
			log.Println("set expire error: ", err)
		}
	}
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"token": token,
			"detail": errmsg.CodeMsg[code],
		},
	})
}