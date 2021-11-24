package test

import (
	"encoding/json"
	"gebi/app/Http/Serializer"
	"gebi/core"
	_ "gebi/utils/database"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var r = core.SetupRouter()

func TestAddress_List(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    int
		wantErr bool
	}{
		{"first case", "http://127.0.0.1/address/list", 0, false},
		{"second case", "http://127.0.0.1/address/list?id=7", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := tt.args
			body := Get(url, r, t)

			var resp Serializer.Response
			err := json.Unmarshal(body, &resp)

			assert.Nil(t, err)
			assert.Equal(t, tt.want, resp.Code)
		})
	}
}

func TestAddress_Add(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string][]string
		want    int
		wantErr bool
	}{
		{"first case", url.Values{"uid": {"58"}, "code": {"1234"}, "phone": {"53422131"}, "address": {"test address 1"}}, 0, false},
		{"second case", url.Values{"uid": {"26"}, "code": {"95"}, "address": {"test address 2"}}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "http://127.0.0.1/address/add"
			body := PostForm(url, tt.args, r, t)

			var resp Serializer.Response
			err := json.Unmarshal(body, &resp)

			assert.Nil(t, err)
			assert.Equal(t, tt.want, resp.Code)
		})
	}
}
