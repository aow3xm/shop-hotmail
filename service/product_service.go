package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aow3xm/shop-hotmail/config"
	"github.com/aow3xm/shop-hotmail/constants"
	"github.com/aow3xm/shop-hotmail/model"
)

type ProductService struct {
	envConfig config.EnvConfig
}

const baseURL = "https://ballmail.shop/api"

func NewProductService(envConfig config.EnvConfig) *ProductService {
	return &ProductService{
		envConfig: envConfig,
	}
}

func (s *ProductService) GetStockService(request model.GetStockRequest) (model.GetStockResponse, error) {
	var count int
	resp, err := http.Get(baseURL + "/stock?api_key=" + s.envConfig.ApiKey)
	if err != nil {
		return model.GetStockResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.GetStockResponse{}, err
	}
	var apiGetStock model.ApiResponse[model.ApiGetStock]
	err = json.Unmarshal(body, &apiGetStock)
	if err != nil {
		return model.GetStockResponse{}, err
	}
	switch request.KioskID {
	case constants.HOTMAIL_KIOSK_ID:
		count = apiGetStock.Data.Hotmail
	case constants.OUTLOOK_KIOSK_ID:
		count = apiGetStock.Data.Outlook
	}

	return model.GetStockResponse{
		Sum: count,
	}, nil
}

func (s *ProductService) PurchaseService(request model.PurchaseRequest) (model.PurchaseResponse, error) {
	url := s.purchaseUrlBuilder(request)
	resp, err := http.Get(url)
	if err != nil {
		return model.PurchaseResponse{}, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.PurchaseResponse{}, nil
	}
	var apiPurchase model.ApiResponse[[]model.ApiPurchase]

	err = json.Unmarshal(body, &apiPurchase)
	if err != nil {
		return model.PurchaseResponse{}, nil
	}

	var products model.PurchaseResponse

	for _, product := range apiPurchase.Data {
		products = append(products, model.ProductItem{
			Product: fmt.Sprintf("%s|%s|%s|%s", product.Login, product.Password, product.RefreshToken, product.ServiceID),
		})
	}

	return products, nil

}

func (s *ProductService) purchaseUrlBuilder(request model.PurchaseRequest) string {
	url := baseURL + "/account"
	switch request.KioskID {
	case constants.HOTMAIL_KIOSK_ID:
		url += "/hotmail"
	case constants.OUTLOOK_KIOSK_ID:
		url += "/outlook"
	}
	url += "?api_key=" + s.envConfig.ApiKey
	url += "&count=" + request.Quantity
	return url
}
