package controllers

import (
	"discountmodule/entities"
	"discountmodule/usecases"
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ServiceHttp struct {
	usc usecases.ServiceUsecase
}

func NewServiceController(usecase usecases.ServiceUsecase) ServiceHttp {
	return ServiceHttp{usc: usecase}
}

func (shtt *ServiceHttp) DisCountModule(context *fiber.Ctx) error {
	// สร้าง UUID สำหรับ transactionCode
	transactionCode := uuid.New().String()
	// สร้างตัวแปร reqBody เพื่อเก็บข้อมูลที่รับมาจาก client
	reqBody := entities.RequestDisCountModule{}
	// พยายาม parse ข้อมูล JSON ที่รับจาก client
	if err := context.BodyParser(&reqBody); err != nil {
		// ถ้ามีข้อผิดพลาดในการ parse JSON จะตอบกลับ error
		response := entities.Response{
			Status:          "ER",
			ErrorCode:       "ERR001",
			ErrorMessage:    []string{"Invalid JSON format"},
			TransactionCode: transactionCode,
		}
		return context.Status(fiber.StatusBadRequest).JSON(response)
	}
	usecase := shtt.usc.DisCountModule(reqBody, transactionCode)
	return context.JSON(usecase)
}
