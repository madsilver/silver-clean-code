package entity

type operationType struct {
	BuyCash             int
	InstallmentPurchase int
	WithDraw            int
	Payment             int
}

var OperationType = operationType{
	BuyCash:             1,
	InstallmentPurchase: 2,
	WithDraw:            3,
	Payment:             4,
}
