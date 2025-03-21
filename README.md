# ชื่อโปรเจกต์ Discount Module

## ผู้พัฒนา

ชื่อเต็ม: หัสดินทร์ ชาติเชื้อ Hasadin Chartchue

...


# คู่มือการใช้งานระบบคำนวณส่วนลด (Discount Module)

## คำอธิบาย
ฟังก์ชัน `DisCountModule` ใช้สำหรับคำนวณส่วนลดจากข้อมูลที่ส่งเข้ามา โดยสามารถรองรับส่วนลดประเภทต่าง ๆ ได้แก่:
- **Coupon**: ส่วนลดแบบจำนวนเงินคงที่ (fixed amount) หรือแบบเปอร์เซ็นต์ (percentage)
- **On Top Discount**: ส่วนลดพิเศษเพิ่มเติม เช่น ส่วนลดตามหมวดหมู่ หรือการใช้แต้มสะสม
- **Seasonal Discount**: ส่วนลดตามโปรโมชั่น เช่น ซื้อครบจำนวนที่กำหนดแล้วลดราคา

## วิธีการติดตั้งและใช้งาน

### 1. การติดตั้งโครงการ
หากยังไม่มี Go ติดตั้งในเครื่อง สามารถดาวน์โหลดได้จาก [Go Official Site](https://go.dev/dl/)

จากนั้นใช้คำสั่งต่อไปนี้เพื่อดาวน์โหลด dependency ที่จำเป็น:
```sh
go mod tidy
```

### 2. การรันเซิร์ฟเวอร์
ใช้คำสั่งต่อไปนี้เพื่อรันโปรแกรม:
```sh
go run main.go
```

เมื่อเซิร์ฟเวอร์ทำงานสำเร็จ API จะพร้อมใช้งานที่:
```
http://0.0.0.0:7951/playtorium/discountmodule
```

*(มี Collection สำหรับ Postman ให้ใช้งาน)*

## โครงสร้างของ Request

### Request Body
```json
{
  "cart": [
    {
      "item_id": 1,
      "name": "T-Shirt",
      "category": "Clothing",
      "price": 350,
      "quantity": 1
    },
    {
      "item_id": 2,
      "name": "Hat",
      "category": "Accessories",
      "price": 250,
      "quantity": 1
    },
    {
      "item_id": 3,
      "name": "Hoodie",
      "category": "Clothing",
      "price": 700,
      "quantity": 1
    },
    {
      "item_id": 4,
      "name": "Watch",
      "category": "Electronics",
      "price": 850,
      "quantity": 1
    }
  ],
  "discounts": {
    "coupon": {
      "type": "fixed_amount",
      "amount": 50
    },
    "ontop": {
      "type": "category_discount",
      "category": "Clothing",
      "percentage": 15
    },
    "seasonal": {
      "every_x": 300,
      "discount_y": 40
    }
  }
}
```

### ตัวอย่าง On Top Discount แบบใช้แต้มสะสม (Points)
```json
{
  "ontop": {
    "type": "points",
    "points_used": 100
  }
}
```

### คำอธิบายพารามิเตอร์
- `cart`: รายการสินค้าที่อยู่ในตะกร้า ประกอบด้วย
  - `item_id`: หมายเลขสินค้า
  - `name`: ชื่อสินค้า
  - `category`: หมวดหมู่สินค้า
  - `price`: ราคาสินค้า
  - `quantity`: จำนวนสินค้า
- `discounts`: ส่วนลดที่สามารถใช้ได้
  - `coupon`: คูปองส่วนลด (แบบ fixed amount หรือ percentage)
  - `ontop`: ส่วนลดเพิ่มเติม เช่น ส่วนลดตามหมวดหมู่ หรือใช้แต้มสะสม
  - `seasonal`: ส่วนลดตามโปรโมชั่น เช่น ซื้อครบจำนวนที่กำหนดแล้วลดราคา

## ตัวอย่าง Response

### กรณีสำเร็จ (OK)
```json
{
  "status": "OK",
  "message": "End of process",
  "transactionCode": "eb544488-8e6f-4c95-89b9-effbb403bda0",
  "totalPrice": 1752.5
}
```

### กรณีเกิดข้อผิดพลาด (Error)
```json
{
  "status": "ER",
  "errorCode": "ER995",
  "errorMessage": [
    "couponType is not fixed_amount or percent"
  ],
  "transactionCode": "832c8c62-0c92-4a29-bd7e-e026639b0f92"
}
```

## หมายเหตุ
- หาก `couponType` ไม่ใช่ "fixed_amount" หรือ "percent" ระบบจะคืนค่า error
- หาก `seasonalDiscountY` หรือ `seasonalEveryX` มีค่าต่ำกว่าหรือเท่ากับ 0 ระบบจะไม่คำนวณส่วนลดฤดูกาล
- คูปอง และส่วนลดอื่น ๆ จะถูกนำไปใช้ตามลำดับที่กำหนด

---

## การทดสอบ API ด้วย Postman
สามารถใช้ Postman เพื่อทดสอบ API โดย
1. เปิด Postman
2. ตั้งค่า method เป็น `GET`
3. กรอก URL: `http://0.0.0.0:7951/playtorium/discountmodule`
4. ส่ง request และตรวจสอบ response

---

## ติดต่อผู้พัฒนา
หากพบปัญหาหรือมีข้อเสนอแนะ สามารถติดต่อทีมพัฒนาได้ที่:
- **อีเมล**: hasadin.chartchue1998@gmail.com
- **GitHub**: https://github.com/hasadin10/playtorium-assignment.git