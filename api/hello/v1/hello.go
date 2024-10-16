package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
)

type HelloReq struct {
	g.Meta `path:"/ping" tags:"Admin API" method:"get" sum:"Server Ping Check"`
}
type HelloRes struct {
	g.Meta `status:"201" resEg:"resource/test.json"`
	Status string `json:"status" des:"server status" eg:"OK!"`
}

// Example case
type Issue3747Res401 struct {
	g.Meta `resEg:"resource/test.json"`
}

// Override case 1
type Issue3747Res402 struct {
	g.Meta `mime:"application/json"`
}

// Override case 2
type Issue3747Res403 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Common response case
type Issue3747Res404 struct{}

func (r HelloRes) ResponseStatusMap() map[goai.StatusCode]any {
	return map[goai.StatusCode]any{
		401: Issue3747Res401{},
		402: Issue3747Res402{},
		403: Issue3747Res403{},
		404: Issue3747Res404{},
		406: struct{}{},
	}
}
