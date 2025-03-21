package utils

import (
	"discountmodule/entities"
	"errors"
	"fmt"
	"math"
)

func CalculateTotal(cart []entities.CartItem) (float64, error) {
	if len(cart) == 0 {
		return 0, errors.New("ตะกร้าสินค้าไม่พบข้อมูล")
	}

	total := 0.0
	for _, item := range cart {
		// ตรวจสอบว่า Price และ Quantity ต้องไม่ติดลบหรือเป็น 0
		if item.Price <= 0 {
			return 0, fmt.Errorf("ราคาของสินค้าชื่อ %s ไม่ถูกต้อง (ราคา: %.2f) ราคาต้องมากกว่า 0", item.Name, item.Price)
		}
		if item.Quantity <= 0 {
			return 0, fmt.Errorf("จำนวนสินค้าชื่อ %s ไม่ถูกต้อง (จำนวน: %d) จำนวนต้องมากกว่า 0", item.Name, item.Quantity)
		}

		// คำนวณราคารวม
		total += item.Price * float64(item.Quantity)
		// fmt.Println("Item:", item.Name, "Price:", item.Price, "Quantity:", item.Quantity, "Total:", total)
	}
	return total, nil
}

func ApplyFixedAmountCoupon(total float64, fixedAmount float64) (float64, error) {
	// เช็คว่าราคารวมต้องมากกว่า 0
	if total <= 0 {
		return total, fmt.Errorf("ไม่สามารถใช้ส่วนลดได้: ยอดรวมสินค้า (%.2f) ต้องมากกว่า 0", total)
	}

	// เช็คว่าส่วนลดต้องไม่เป็นค่าลบ
	if fixedAmount < 0 {
		return total, fmt.Errorf("ส่วนลดไม่ถูกต้อง: จำนวนเงินส่วนลด (%.2f) ต้องเป็นค่าบวก", fixedAmount)
	}

	// เช็คว่าส่วนลดต้องไม่เกินยอดรวม
	if fixedAmount > total {
		return total, fmt.Errorf("ส่วนลดมากเกินไป: จำนวนเงินส่วนลด (%.2f) ห้ามเกินยอดรวมสินค้า (%.2f)", fixedAmount, total)
	}

	// คำนวณยอดใหม่หลังหักส่วนลด
	total -= fixedAmount
	return total, nil
}

func ApplyPercentageCoupon(total float64, percentage float64) (float64, error) {
	// ตรวจสอบว่า total ต้องเป็นค่าบวก
	if total < 0 {
		return 0, fmt.Errorf("ยอดรวมต้องเป็นค่ามากกว่า 0 (ได้รับ %.2f)", total)
	}

	// ตรวจสอบว่า percentage ต้องอยู่ในช่วงที่ถูกต้อง (0 - 100%)
	if percentage < 0 || percentage > 100 {
		return 0, fmt.Errorf("เปอร์เซ็นต์ส่วนลดต้องอยู่ระหว่าง 0 - 100 (ได้รับ %.2f%%)", percentage)
	}

	// คำนวณส่วนลด
	total = total - (total * percentage / 100)

	return total, nil
}

func ApplyCategoryDiscount(cart []entities.CartItem, category string, discountPercentage float64) (float64, error) {
	// ตรวจสอบว่าหมวดหมู่ต้องไม่เป็นค่าว่าง
	if category == "" {
		return 0, fmt.Errorf("หมวดหมู่สินค้าต้องไม่เป็นค่าว่าง")
	}

	// ตรวจสอบว่าส่วนลดต้องอยู่ในช่วงที่ถูกต้อง (0 - 100%)
	if discountPercentage < 0 || discountPercentage > 100 {
		return 0, fmt.Errorf("ส่วนลดหมวดหมู่ไม่ถูกต้อง: ต้องอยู่ระหว่าง 0 - 100%% (ได้รับ %.2f%%)", discountPercentage)
	}

	total := 0.0
	found := false // ใช้ตรวจสอบว่ามีสินค้าตรงกับหมวดหมู่หรือไม่

	for _, item := range cart {
		price := item.Price

		// ถ้าสินค้าอยู่ในหมวดหมู่ที่ระบุ ให้ทำการหักส่วนลด
		if item.Category == category {
			price -= (price * discountPercentage / 100)
			found = true
		}

		// คำนวณราคารวมของสินค้าแต่ละชิ้น (ราคา x จำนวน)
		total += price * float64(item.Quantity)
	}

	// ถ้าไม่มีสินค้าในหมวดหมู่ที่กำหนดเลย ให้ return error
	if !found {
		return 0, fmt.Errorf("ไม่พบสินค้าหมวด '%s' ในตะกร้า", category)
	}

	return total, nil
}

func ApplyPointsDiscount(total float64, pointsUsed int) (float64, error) {
	if total <= 0 {
		return total, errors.New("ยอดรวมต้องมากกว่า 0")
	}
	if pointsUsed < 0 {
		return total, errors.New("จำนวนแต้มที่ใช้ต้องไม่เป็นค่าลบ")
	}

	// คำนวณส่วนลดสูงสุดที่ใช้ได้ (ไม่เกิน 20% ของยอดรวม)
	maxDiscount := total * 0.2
	discount := float64(pointsUsed)

	if discount > maxDiscount {
		discount = maxDiscount
	}

	// ตรวจสอบว่าหลังหักแต้มแล้วยอดรวมยังไม่ติดลบ
	newTotal := total - discount
	if newTotal < 0 {
		return 0, errors.New("ส่วนลดมากกว่ายอดรวม ไม่สามารถใช้ได้")
	}

	return newTotal, nil
}

func ApplySeasonalDiscount(total float64, seasonalEveryX float64, seasonalDiscountY float64) (float64, error) {
	// ตรวจสอบข้อมูลที่รับเข้ามา
	if total < 0 {
		return 0, errors.New("ยอดรวมต้องไม่เป็นค่าลบ")
	}

	if seasonalEveryX <= 0 {
		return 0, errors.New("ค่าขั้นต่ำที่ใช้คำนวณส่วนลด (every_x) ต้องมากกว่าศูนย์")
	}

	if seasonalDiscountY < 0 {
		return 0, errors.New("มูลค่าส่วนลด (discount_y) ต้องไม่เป็นค่าลบ")
	}

	// ถ้ายอดรวมเป็น 0 หรือน้อยกว่าค่าขั้นต่ำ ไม่ต้องคำนวณส่วนลด
	if total == 0 || total < seasonalEveryX {
		return total, nil
	}

	// จำนวนครั้งที่จะได้รับส่วนลด (ปัดเศษลง)
	discountTimes := math.Floor(total / seasonalEveryX)

	// คำนวณส่วนลดทั้งหมด
	discountAmount := discountTimes * seasonalDiscountY

	// ตรวจสอบว่าส่วนลดไม่มากกว่ายอดเงินรวม
	if discountAmount > total {
		return 0, errors.New("ส่วนลดมากกว่ายอดรวม ไม่สามารถใช้ได้")
	}

	// หักส่วนลดจากราคารวม
	total -= discountAmount

	return total, nil
}

