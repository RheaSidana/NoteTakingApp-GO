package sessions

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Cookie interface{
	SetCookie(c *gin.Context, session string)
	DeleteCookie(c *gin.Context)
	GetCookie(c *gin.Context) (string, error)
}

type cookie struct{}

func NewCookie() Cookie{
	return &cookie{}
}

func (ck *cookie) SetCookie(c *gin.Context, session string){
	cookie := http.Cookie{
		Name: "sid",
		Value: session,
		HttpOnly: true,
		Secure: false,
		Path: "/notes",
		Domain: "localhost",
		Expires: time.Now().Add(time.Hour*24),
	}

	c.SetCookie(
		cookie.Name, 
		cookie.Value,
		int(cookie.Expires.Unix()),
		cookie.Path,
		cookie.Domain,
		cookie.Secure,
		cookie.HttpOnly,
	)
}

func (ck *cookie) DeleteCookie(c *gin.Context){
	cookie := &http.Cookie{
        Name:     "sid",
        Value:    "",
        Expires:  time.Now().AddDate(0, 0, -1), 
        MaxAge:   -1,                           
        HttpOnly: true,
        Secure:   false,                         
        SameSite: http.SameSiteStrictMode,      
        Path:     "/notes",                      
        Domain:   "localhost",               
    }

	c.SetCookie(
		cookie.Name, 
		cookie.Value,
		int(cookie.Expires.Unix()),
		cookie.Path,
		cookie.Domain,
		cookie.Secure,
		cookie.HttpOnly,
	)
}

func (ck *cookie) GetCookie(c *gin.Context) (string, error){
	cookieValue, err := c.Cookie("sid")
	fmt.Println("Value",cookieValue)
    if err != nil {
		return "", errors.New("something went wrong")
    }

	return cookieValue, nil
}