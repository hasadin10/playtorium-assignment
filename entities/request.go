package entities

import (
	// "mime/multipart"
)

type RequestHeader struct {
	Authorization string `json:"Authorization"`
}


// Struct หลักที่ใช้ใน Usecase
type RequestDisCountModule struct {
	Cart      []CartItem `json:"cart" validate:"required,dive"` // รายการสินค้าในตะกร้า
	Discounts Discounts  `json:"discounts"`
}


// Struct สำหรับรายการสินค้าในตะกร้า
type CartItem struct {
	ItemID   int     `json:"item_id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
}


// Struct สำหรับคูปองส่วนลด
type Coupon struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// Struct สำหรับส่วนลดแบบ On Top
type OnTopDiscount struct {
    Type       string  `json:"type"` // ประเภทของส่วนลด On Top
    Category   string  `json:"category,omitempty"`     // สำหรับ category_discount เท่านั้น
    Percentage float64 `json:"percentage,omitempty"`   // สำหรับ category_discount เท่านั้น
    PointsUsed int     `json:"points_used,omitempty"`  // สำหรับ points เท่านั้น
}
// Struct สำหรับส่วนลดตามโปรโมชั่น
type SeasonalDiscount struct {
	EveryX    float64 `json:"every_x"`
	DiscountY float64 `json:"discount_y"`
}

// Struct หลักที่รวมส่วนลดทั้งหมด
type Discounts struct {
	Coupon   Coupon           `json:"coupon"`
	OnTop    OnTopDiscount  `json:"ontop"`
	Seasonal SeasonalDiscount `json:"seasonal"`
}



