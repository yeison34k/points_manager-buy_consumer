package app

import (
	"fmt"
)

type BuyHandler struct {
	BuyApp *BuyApp
}

func NewBuyHandler(pointApp *BuyApp) *BuyHandler {
	return &BuyHandler{
		BuyApp: pointApp,
	}
}

func (h *BuyHandler) HandleBuyCreation() error {
	err := h.BuyApp.CreateBuy()
	if err != nil {
		return fmt.Errorf("failed to create buy: %w", err)
	}

	return nil
}
