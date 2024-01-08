package entities // เก็บ struct ที่เป็น global หรือที่เรียกใช้จากทุก module

import (
	"github.com/Vodka479/space-shop-tutorial/pkg/spacelogger"
	"github.com/gofiber/fiber/v2"
)

type IResponse interface {
	Success(code int, data any) IResponse
	Error(code int, tractId, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data       any
	ErrorRes   *ErrorResponse
	Context    *fiber.Ctx
	IsError    bool
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg     string `json:"message"`
}

func NewResponse(c *fiber.Ctx) IResponse {
	return &Response{
		Context: c,
	}
}

func (r *Response) Success(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	spacelogger.InitSpaceLogger(r.Context, &r.Data).Print().Save()
	return r
}
func (r *Response) Error(code int, tractId, msg string) IResponse {
	r.StatusCode = code
	r.ErrorRes = &ErrorResponse{
		TraceId: tractId,
		Msg:     msg,
	}
	r.IsError = true
	spacelogger.InitSpaceLogger(r.Context, &r.ErrorRes).Print().Save()
	return r
}
func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() any { //status คล้ายกลับของ fiber c.Status(fiber.StatusOK).JSON(res)
		if r.IsError {
			return &r.ErrorRes //r. = reference
		}
		return &r.Data
	}())
}
