package controller

import (
	"github.com/gin-gonic/gin"
	s "api/services"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
	u "api/util"
)

type Data struct {
	Url string `json:"url"`
}

type Access struct {
	Access_token string `json:"access_token"`
	Token_type string `json:"token_type"`
	Error string `json:"error"`
}

type User struct {
	Id int `json:"id"`
	Username string `json:"login"`
}

type Token struct {
	Token string `json:"token"`
}

//获取第三方跳转链接
func Login(c *gin.Context)  {
	conf, err := s.GetGits()
	if err != nil {
		panic(err)
	}
	var data Data
	data.Url = fmt.Sprintf(conf.Url, conf.Client_id)
	c.JSON(200, gin.H{
		"data": data,
		"msg": "",
	})
}

//第三方回调地址
func Callback(c *gin.Context)  {
	conf, err := s.GetGits()
	if err != nil {
		panic(err)
	}
	code := c.Query("code")
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
    access := url.Values{}
    access.Set("client_id", conf.Client_id)
    access.Set("client_secret", conf.Client_secret)
	access.Set("code", code)
	url := fmt.Sprintf("%s?%s", conf.Access_url, access.Encode())
    r, _ := http.NewRequest("POST", url, nil) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")
	res, err1 := httpClient.Do(r)
	if err1 != nil {
		panic(err1)
	}
	defer res.Body.Close()
	accRet, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		panic(err2)
	}
	var accs Access;
    if err3 := json.Unmarshal([]byte(string(accRet)), &accs); err != nil {
		panic(err3)
	}
	fmt.Println(access.Encode(), strings.NewReader(string(accRet)),accs)
	if accs.Access_token == "" {
		panic(string(accRet))
	}
	fmt.Printf("%+v", accs)
	fmt.Println(accs, string(accRet))
    requestGet, _:= http.NewRequest("GET", conf.User_url, nil)
	token := fmt.Sprintf("token %s", accs.Access_token)
    requestGet.Header.Add("Authorization", token)
    requestGet.Header.Add("Accept", "application/json")

    resp, err3 := httpClient.Do(requestGet)
	if err3 != nil {
		panic(err3)
	}
	defer resp.Body.Close()
	userRet, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		panic(err4)
	}
	var users User;
    if err5 := json.Unmarshal(userRet, &users); err != nil {
		panic(err5)
	}
	if users.Id == 0 {
		panic(string(userRet))
	}
	fmt.Printf("%+v", users)
	fmt.Println(users, string(userRet))
    claims := &u.JWTClaims{
        UserID: users.Id,
        Username: users.Username,
    }
    claims.IssuedAt = time.Now().Unix()
    claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(3600)).Unix()
    singedToken, err6 := u.GetToken(claims)
    if err6 != nil {
		panic(err6)
    }
	var tokens Token
	tokens.Token = singedToken
	c.JSON(200, gin.H{
		"data": tokens,
		"msg": "",
	})
}

func GetUser(c *gin.Context){
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODYwNTMyOTQsImlhdCI6MTU4NjA0OTY5NCwidXNlcl9pZCI6MjA4NTMxNjksInVzZXJuYW1lIjoiQnJvdWNlU29uZyJ9.wF7aPloy0mZdLaAWwlwkFTuVxHpLZK_FBrq7Dz-lp_o"
	users, _ := u.VerifyToken(token)
	c.JSON(200, gin.H{
		"data": users,
		"msg": "",
	})
}
