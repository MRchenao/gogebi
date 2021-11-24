package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// Get 根据特定请求uri，发起get请求返回响应
func Get(uri string, router *gin.Engine, t *testing.T) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 校验状态码是否符合预期
	assert.Equal(t, http.StatusOK, w.Code)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(uri string, param url.Values, router *gin.Engine, t *testing.T) []byte {
	// 构造post请求
	req := httptest.NewRequest("POST", uri, strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)

	// 校验状态码是否符合预期
	assert.Equal(t, http.StatusOK, w.Code)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
