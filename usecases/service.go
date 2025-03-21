package usecases

import (
	"discountmodule/entities"
	"discountmodule/interfaces"
	"discountmodule/utils"
	"discountmodule/validations"
	// "fmt"
)

type ServiceUsecase struct {
	httpRequest interfaces.HttpRequest
}

func NewServiceUsecase(httpRequest interfaces.HttpRequest) ServiceUsecase {
	return ServiceUsecase{httpRequest: httpRequest}
}

func (uso *ServiceUsecase) DisCountModule(reqBody entities.RequestDisCountModule, transactionCode string) entities.Response {

	// ตรวจสอบว่า reqBody มีข้อมูลที่ถูกต้องหรือไม่
	// ถ้าไม่ถูกต้องจะส่งกลับ error message
	errorMessages := validations.ValidateRequest(reqBody)

	if len(errorMessages) > 0 {
		return entities.Response{
			Status:          "ER",
			ErrorCode:       "ER999",
			ErrorMessage:    errorMessages,
			TransactionCode: transactionCode,
		}
	}

	// คำนวณราคารวมของตะกร้า
	total, err := utils.CalculateTotal(reqBody.Cart)
	if err != nil {
		return entities.Response{
			Status:          "ER",
			ErrorCode:       "ER998",
			ErrorMessage:    []string{err.Error()},
			TransactionCode: transactionCode,
		}
	}

	couponType := reqBody.Discounts.Coupon.Type
	couponAmount := reqBody.Discounts.Coupon.Amount

	ontopType := reqBody.Discounts.OnTop.Type
	ontopCagtegory := reqBody.Discounts.OnTop.Category
	ontopPercentage := reqBody.Discounts.OnTop.Percentage
	ontopPointsUsed := reqBody.Discounts.OnTop.PointsUsed

	seasonalEveryX := reqBody.Discounts.Seasonal.EveryX
	seasonalDiscountY := reqBody.Discounts.Seasonal.DiscountY

	// Apply Coupon (Fixed Amount or Percentage)
	if couponType != "" {
		if couponType == "fixed_amount" && couponAmount > 0 {
			total, err = utils.ApplyFixedAmountCoupon(total, couponAmount)
			if err != nil {
				return entities.Response{
					Status:          "ER",
					ErrorCode:       "ER997",
					ErrorMessage:    []string{err.Error()},
					TransactionCode: transactionCode,
				}
			}
		} else if couponType == "percent" && couponAmount > 0 {
			total, err = utils.ApplyPercentageCoupon(total, couponAmount)
			if err != nil {
				return entities.Response{
					Status:          "ER",
					ErrorCode:       "ER996",
					ErrorMessage:    []string{err.Error()},
					TransactionCode: transactionCode,
				}
			}
		} else {
			return entities.Response{
				Status:          "ER",
				ErrorCode:       "ER995",
				ErrorMessage:    []string{"couponType is not fixed_amount or percent"},
				TransactionCode: transactionCode,
			}
		}
	}

	// Apply On Top Discount
	if ontopType != "" {
		if ontopType == "category_discount" && ontopCagtegory != "" && ontopPercentage > 0 {
			total, err = utils.ApplyCategoryDiscount(reqBody.Cart, ontopCagtegory, ontopPercentage)
			if err != nil {
				return entities.Response{
					Status:          "ER",
					ErrorCode:       "ER994",
					ErrorMessage:    []string{err.Error()},
					TransactionCode: transactionCode,
				}
			}
		} else if ontopType == "points" && ontopPointsUsed > 0 {
			total, err = utils.ApplyPointsDiscount(total, ontopPointsUsed)
			if err != nil {
				return entities.Response{
					Status:          "ER",
					ErrorCode:       "ER993",
					ErrorMessage:    []string{err.Error()},
					TransactionCode: transactionCode,
				}
			}
		}else {
			return entities.Response{
				Status:          "ER",
				ErrorCode:       "ER992",
				ErrorMessage:    []string{"ontopType is not category_discount or points"},
				TransactionCode: transactionCode,
			}
		}
	} 


	if seasonalEveryX > 0 && seasonalDiscountY > 0 {
		total, err = utils.ApplySeasonalDiscount(total, seasonalEveryX, seasonalDiscountY)
		if err != nil {	
			return entities.Response{
				Status:          "ER",
				ErrorCode:       "ER992",
				ErrorMessage:    []string{err.Error()},
				TransactionCode: transactionCode,
			}
		}
	}

	return entities.Response{
		Status:          "OK",
		Message:         "End of process",
		TransactionCode: transactionCode,
		TotalPrice:      total,
	}
}
