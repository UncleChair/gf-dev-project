package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/ping" tags:"Admin API" method:"get" sum:"Server Ping Check"`
}
type HelloRes struct {
	g.Meta      `status:"201"`
	Status      string    `json:"status" des:"server status" eg:"OK!"`
	StatusError *Error404 `status:"403" `
}

type Error404 struct {
	Code    string `json:"code" eg:"1"`
	Message string `json:"message" eg:"Not Found"`
}
