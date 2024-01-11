package app

type BuyService interface {
	CreateBuy() error
}

type BuyApp struct {
	BuyService BuyService
}

func NewBuyApplication(buyService BuyService) *BuyApp {
	return &BuyApp{
		BuyService: buyService,
	}
}

func (a *BuyApp) CreateBuy() error {
	return a.BuyService.CreateBuy()
}
