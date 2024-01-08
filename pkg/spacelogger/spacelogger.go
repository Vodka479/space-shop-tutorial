package spacelogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Vodka479/space-shop-tutorial/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ISpaceLogger interface {
	Print() ISpaceLogger
	Save()
	SetQuery(c *fiber.Ctx) //http://localhost:3000/v1/products?name=test //name=test คือ query param หรือ name:id query ผ่าน param
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

type spaceLogger struct {
	Time       string `json:"time"`
	Ip         string `json:"tp"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitSpaceLogger(c *fiber.Ctx, res any) ISpaceLogger {
	log := &spaceLogger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: c.Response().StatusCode(),
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)
	return log
}

func (l *spaceLogger) Print() ISpaceLogger {
	utils.Debug(l)
	return l
}

func (l *spaceLogger) Save() {
	data := utils.Output(l) //รับ log แปลง log จาก struct ปกติ เป็น string define แล้วเพิ่มไปใน file txt

	filename := fmt.Sprintf("./assets/logs/spacelogger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //file permission
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n") //สร้าง file
}

func (l *spaceLogger) SetQuery(c *fiber.Ctx) {
	var body any
	if err := c.QueryParser(&body); err != nil {
		log.Printf("query parser error: %v", err)
	}
	l.Query = body
}

func (l *spaceLogger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser error: %v", err)
	}

	switch l.Path {
	case "v1/users/signup":
		l.Body = "never gonna give you up"
	default:
		l.Body = body
	}
}

func (l *spaceLogger) SetResponse(res any) {
	l.Response = res
}
