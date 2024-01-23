package main

import (
	"fmt"
	"sync"
)

// Outlet 门店模型
type Outlet struct {
	ID           int
	Address      string
	SpecialOffer string
}

// Product 产品模型
type Product struct {
	ID               int
	Name             string
	Price            float64
	AssociatedOutlet *Outlet
	IsNationalOffer  bool
}

// Accessory 配件模型
type Accessory struct {
	ID    int
	Name  string
	Price float64
}

// Customer 顾客模型
type Customer struct {
	ID               int
	Points           int
	IsSubscribed     bool
	SubscriptionType string
	sync.Mutex       // 保护 Points 字段的锁
}

// Order 订单模型
type Order struct {
	ID        int
	Customer  *Customer
	Product   *Product
	Quantity  int
	OrderDate string
}

// Subscription 订阅模型
type Subscription struct {
	ID           int
	Customer     *Customer
	ServiceType  string
	IsSubscribed bool
}

// MarketingCampaign 营销活动模型
type MarketingCampaign struct {
	ID                   int
	ParticipatingOutlets []*Outlet
	AffectedProducts     []*Product
	AffectedPrice        float64
}

// GadgetPointsProgram 积分计划模型
type GadgetPointsProgram struct {
	Customer *Customer
}

// CalculatePoints 计算积分
func (gp *GadgetPointsProgram) CalculatePoints(order *Order) {
	gp.Customer.Lock()
	defer gp.Customer.Unlock()

	gp.Customer.Points += order.Quantity
}

// RedeemAccessory 兑换配件
func (gp *GadgetPointsProgram) RedeemAccessory() *Accessory {
	gp.Customer.Lock()
	defer gp.Customer.Unlock()

	if gp.Customer.Points >= 10 {
		gp.Customer.Points -= 10
		return &Accessory{ID: 1, Name: "Free Accessory", Price: 0.0}
	}
	return nil
}

func main() {
	// 示例用法
	outlet := &Outlet{ID: 1, Address: "123 Main St", SpecialOffer: "10% off on all accessories"}
	product := &Product{ID: 1, Name: "Laptop", Price: 1000.0, AssociatedOutlet: outlet, IsNationalOffer: true}
	customer := &Customer{ID: 1, Points: 0, IsSubscribed: true, SubscriptionType: "Monthly"}
	order := &Order{ID: 1, Customer: customer, Product: product, Quantity: 2, OrderDate: "2024-01-23"}

	// 计算积分
	gadgetPointsProgram := &GadgetPointsProgram{Customer: customer}
	gadgetPointsProgram.CalculatePoints(order)

	// 兑换配件
	freeAccessory := gadgetPointsProgram.RedeemAccessory()

	// 输出信息
	fmt.Printf("Customer Points after order: %d\n", customer.Points)

	if freeAccessory != nil {
		fmt.Printf("Redeemed Accessory: %s\n", freeAccessory.Name)
	} else {
		fmt.Println("Not enough points to redeem an accessory.")
	}
}
