package middleware

import (
	"encoding/json"
	"errors"
	"github.com/cy77cc/hioshop/configs"
	"github.com/cy77cc/hioshop/models"
	"github.com/cy77cc/hioshop/util"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Code string `json:"code" form:"code" binding:"required"`
}

func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(InitParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}

type WxCode2session struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

// RegisterRoute 注册路由
func RegisterRoute(r *gin.Engine, handle *jwt.GinJWTMiddleware) {
	r.POST("/login", handle.LoginHandler)
	r.NoRoute(handle.MiddlewareFunc(), handleNoRoute())

	auth := r.Group("/auth", handle.MiddlewareFunc())
	auth.GET("/refresh_token", handle.RefreshHandler)
	auth.POST("/main", mainHandler)
}

func HandlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func InitParams() *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:       "hioshop",
		Key:         []byte(configs.SecretKey),
		Timeout:     time.Hour * 24 * 7,
		MaxRefresh:  time.Hour * 24 * 7,
		IdentityKey: configs.IdentityKey,
		PayloadFunc: payloadFunc(),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator(),
		Authorizator:    authorizator(),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		LoginResponse:   loginResponseFunc(),
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName: "hio",
		TimeFunc:      time.Now,
	}
}

func loginResponseFunc() func(c *gin.Context, code int, message string, expire time.Time) {
	return func(c *gin.Context, code int, message string, expire time.Time) {
		parseJWT, _ := ParseJWT(message, []byte(configs.SecretKey))

		db := util.GetDB()
		user := models.User{}
		db.Where("weixin_openid=?", parseJWT.Openid).First(&user)
		c.JSON(code, gin.H{
			"code":   code,
			"errno":  0,
			"errmsg": "",
			"data": map[string]interface{}{
				"token":    message,
				"expire":   expire.Format(time.RFC3339),
				"userInfo": user,
			},
			"is_new": 1,
		})
	}
}

// authenticator -> payloadFunc

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*models.User); ok {
			return jwt.MapClaims{
				configs.IdentityKey: v.WeixinOpenid,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &models.User{
			WeixinOpenid: claims[configs.IdentityKey].(string),
		}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		code := loginVals.Code

		client := resty.New()
		wxRespone := WxCode2session{}
		get, err := client.R().
			SetQueryParams(map[string]string{
				"appid":      configs.Appid,
				"secret":     configs.Secret,
				"js_code":    code,
				"grant_type": "authorization_code",
			}).
			Get("https://api.weixin.qq.com/sns/jscode2session")
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(get.Body(), &wxRespone)
		if err != nil {
			return nil, err
		}
		if wxRespone.Errcode != 0 {
			return nil, errors.New(wxRespone.Errmsg)
		}
		db := util.GetDB()
		user := models.User{}
		uuidv1, _ := uuid.NewUUID()
		user.Username = "微信用户" + uuidv1.String()[:8]
		user.Password = wxRespone.Openid
		user.RegisterTime = time.Now()
		user.LastLoginTime = time.Now()
		user.WeixinOpenid = wxRespone.Openid
		user.Avatar = "/static/images/default_avatar.png"

		db.Find(&user, "weixin_openid=?", wxRespone.Openid)
		if user.ID <= 0 {
			result := db.Create(&user)
			if result.RowsAffected == 1 {
				user.Password = ""
				return &user, nil
			} else {
				return nil, jwt.ErrFailedAuthentication
			}
		} else {
			db.Find(&user, "weixin_openid=?", wxRespone.Openid)
			return &user, nil
		}
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		v, ok := data.(*models.User)
		if v.WeixinOpenid == "" {
			return false
		}
		db := util.GetDB()
		user := models.User{}
		db.First(&user, "weixin_openid=?", v.WeixinOpenid)
		if ok && user.ID != 0 {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func handleNoRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
}

func mainHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	openid, _ := claims[configs.IdentityKey].(string)
	db := util.GetDB()
	user := models.User{}
	db.Where("weixin_openid=?", openid).First(&user)

	c.JSON(200, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": map[string]interface{}{
			"userInfo": user,
		},
	})
}
