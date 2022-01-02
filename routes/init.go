package routes

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/api")
}

func Index(c *gin.Context) {
	pageContents, err := ioutil.ReadFile("pages/index.html")
	if err != nil {
		pageContents = []byte("Sorry, an error has occurred while trying to load this page!")
		fmt.Print("Error: ", err)
	}
	c.Data(200, "text/html; charset=utf-8", pageContents)
}

func Docs(c *gin.Context) {
	pageContents, err := ioutil.ReadFile("pages/docs.html")
	if err != nil {
		pageContents = []byte("Sorry, an error has occurred while trying to load this page!")
		fmt.Println("Error: ", err)
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", pageContents)
}

type request_body_template struct {
	Uppercase_included    bool `json:"uppercase_included"`
	Lowercase_included    bool `json:"lowercase_included"`
	Numbers_included      bool `json:"numbers_included"`
	Specialchars_included bool `json:"specialchars_included"`
	Password_length       int  `json:"password_length"`
}

type err_response_body_template struct {
	Error int `json:"error"`
}

type response_body_template struct {
	Password string `json:"password"`
}

func API(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	request_body := request_body_template{}
	binding_err := c.BindJSON(&request_body)
	generated_password := ""
	if !request_body.Lowercase_included && !request_body.Uppercase_included && !request_body.Numbers_included && !request_body.Specialchars_included {
		response_body := err_response_body_template{1}
		c.JSON(http.StatusBadRequest, response_body)
	} else if request_body.Password_length < 8 {
		response_body := err_response_body_template{2}
		c.JSON(http.StatusBadRequest, response_body)
	} else if request_body.Password_length > 64 {
		response_body := err_response_body_template{3}
		c.JSON(http.StatusBadRequest, response_body)
	} else if binding_err != nil {
		response_body := err_response_body_template{4}
		c.JSON(http.StatusBadRequest, response_body)
	} else {
		for len(generated_password) != request_body.Password_length {
			randGenUppercaseChar := ""
			randGenLowercaseChar := ""
			randGenNum := ""
			randGenSpecialChar := ""
			if request_body.Uppercase_included {
				randGenUppercaseChar = string(rune(rand.Intn(90-65) + 65))
			}
			if request_body.Lowercase_included {
				randGenLowercaseChar = string(rune(rand.Intn(122-97) + 97))
			}
			if request_body.Numbers_included {
				randGenNum = string(rune(rand.Intn(57-48) + 48))
			}
			if request_body.Specialchars_included {
				randGenSpecialChar = string(rune(rand.Intn(47-33) + 33))
			}
			if len(generated_password) != request_body.Password_length {
				generated_password += randGenUppercaseChar
			}
			if len(generated_password) != request_body.Password_length {
				generated_password += randGenLowercaseChar
			}
			if len(generated_password) != request_body.Password_length {
				generated_password += randGenNum
			}
			if len(generated_password) != request_body.Password_length {
				generated_password += randGenSpecialChar
			}
		}
		response_body := response_body_template{generated_password}
		c.JSON(http.StatusOK, response_body)
	}
}
