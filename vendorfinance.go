package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// type Asset struct {
// 	ID             string `json:"ID"`
// 	Color          string `json:"color"`
// 	Size           int    `json:"size"`
// 	Owner          string `json:"owner"`
// 	AppraisedValue int    `json:"appraisedValue"`
// }

type Quote struct {
	QuoteId               string `json:"quoteId"`
	VendorId              string `json:"vendorId"`
	Desc                  string `json:"desc"`
	ItemsList             string `json:"itemsList"`
	BuyerId               string `json:"buyerId"`
	QuoteDate             string `json:"quoteDate"`
	TotalAmount           string `json:"totalAmount"`
	EstimatedDeliveryDate string `json:"extimatedDeliveryDate"`
	Status                string `json:"status"`
}

type PurchaseOrder struct {
	PoId            string `json:"poId"`
	QuoteId         string `json:"quoteId"`
	VendorLimitLeft string `json:"vendorLimitLeft"`
	VendorLimitUsed string `json:"VendorLimitUsed"`
	BankerId        string `json:"bankerId"`
	PoDate          string `json:"poDate"`
	Currency        string `json:"currency"`
	PoStatus        string `json:"poStatus"`
}

type Invoice struct {
	InvoiceId           string `json:"invoiceId"`
	QuoteId             string `json:"quoteId"`
	VendorId            string `json:"vendorId"`
	BuyerId             string `json:"buyerId"`
	PoId                string `json:"poId"`
	InvoiceStatus       string `json:"invoiceStatus"`
	RequestForFinance   string `json:"requestForFinance"`
	InvoiceRaisedDate   string `json:"invoiceRaisedDate"`
	InvoiceAcceptedDate string `json:"invoiceAcceptedDate"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	quotes := []Quote{
		{QuoteId: "quote1", VendorId: "NA", Desc: "NA", ItemsList: "NA", BuyerId: "NA", QuoteDate: "NA", TotalAmount: "NA", EstimatedDeliveryDate: "NA", Status: "NA"},
	}

	for _, quote := range quotes {
		quoteJSON, err := json.Marshal(quote)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(quote.QuoteId, quoteJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// PlaceQuote issues a new quote to the world state with given details.
func (s *SmartContract) PlaceQuote(ctx contractapi.TransactionContextInterface, quoteId string, vendorId string, desc string,
	itemsList string, buyerId string, quoteDate string, totalAmount string,
	extimatedDeliveryDate string, status string) error {
	exists, err := s.IsExists(ctx, quoteId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the quote %s already exists", quoteId)
	}

	quote := Quote{
		QuoteId:               quoteId,
		VendorId:              vendorId,
		Desc:                  desc,
		ItemsList:             itemsList,
		BuyerId:               buyerId,
		QuoteDate:             quoteDate,
		TotalAmount:           totalAmount,
		EstimatedDeliveryDate: extimatedDeliveryDate,
		Status:                status,
	}
	quoteJSON, err := json.Marshal(quote)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(quoteId, quoteJSON)
}

// RaiseInvoice issues a new Invoice to the world state with given details.
func (s *SmartContract) RaiseInvoice(ctx contractapi.TransactionContextInterface, invoiceId string, quoteId string, poId string, vendorId string,
	buyerId string, invoiceStatus string, requestForFinance string,
	invoiceRaisedDate string, invoiceAcceptedDate string) error {
	exists, err := s.IsExists(ctx, invoiceId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the quote %s already exists", invoiceId)
	}

	invoice := Invoice{
		InvoiceId:           invoiceId,
		QuoteId:             quoteId,
		VendorId:            vendorId,
		BuyerId:             buyerId,
		PoId:                poId,
		InvoiceStatus:       invoiceStatus,
		RequestForFinance:   requestForFinance,
		InvoiceRaisedDate:   invoiceRaisedDate,
		InvoiceAcceptedDate: invoiceAcceptedDate,
	}
	invoiceJSON, err := json.Marshal(invoice)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(invoiceId, invoiceJSON)
}

// RaisePO issues a new Invoice to the world state with given details.
func (s *SmartContract) RaisePo(ctx contractapi.TransactionContextInterface, poId string, quoteId string, vendorLimitLeft string,
	VendorLimitUsed string, bankerId string, poDate string, currency string, poStatus string) error {
	exists, err := s.IsExists(ctx, poId)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the quote %s already exists", poId)
	}

	purchaseOrder := PurchaseOrder{
		PoId:            poId,
		QuoteId:         quoteId,
		VendorLimitLeft: vendorLimitLeft,
		VendorLimitUsed: VendorLimitUsed,
		BankerId:        bankerId,
		PoDate:          poDate,
		Currency:        currency,
		PoStatus:        poStatus,
	}
	poJSON, err := json.Marshal(purchaseOrder)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(poId, poJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetQuoteById(ctx contractapi.TransactionContextInterface, quoteId string) (*Quote, error) {
	quoteJSON, err := ctx.GetStub().GetState(quoteId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if quoteJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", quoteId)
	}

	var quote Quote
	err = json.Unmarshal(quoteJSON, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetPoById(ctx contractapi.TransactionContextInterface, poId string) (*PurchaseOrder, error) {
	poJSON, err := ctx.GetStub().GetState(poId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if poJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", poId)
	}

	var purchaseOrder PurchaseOrder
	err = json.Unmarshal(poJSON, &purchaseOrder)
	if err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}

func (s *SmartContract) GetInvoiceById(ctx contractapi.TransactionContextInterface, invoiceId string) (*Invoice, error) {
	invoiceJSON, err := ctx.GetStub().GetState(invoiceId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if invoiceJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", invoiceId)
	}

	var invoice Invoice
	err = json.Unmarshal(invoiceJSON, &invoice)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

// UpdateQuote updates an existing quote in the world state with provided parameters.
func (s *SmartContract) UpdateQuote(ctx contractapi.TransactionContextInterface, quoteId string,
	vendorId string, desc string, itemsList string, buyerId string, quoteDate string,
	totalAmount string, extimatedDeliveryDate string, status string) error {
	exists, err := s.IsExists(ctx, quoteId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", quoteId)
	}

	// overwriting original asset with new asset
	quote := Quote{
		QuoteId:               quoteId,
		VendorId:              vendorId,
		Desc:                  desc,
		ItemsList:             itemsList,
		BuyerId:               buyerId,
		QuoteDate:             quoteDate,
		TotalAmount:           totalAmount,
		EstimatedDeliveryDate: extimatedDeliveryDate,
		Status:                status,
	}
	quoteJSON, err := json.Marshal(quote)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(quoteId, quoteJSON)
}

// UpdatePO updates an existing PO in the world state with provided parameters.
func (s *SmartContract) UpdatePo(ctx contractapi.TransactionContextInterface, poId string, quoteId string,
	vendorLimitLeft string, VendorLimitUsed string, bankerId string,
	poDate string, currency string, poStatus string) error {
	exists, err := s.IsExists(ctx, poId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", poId)
	}

	// overwriting original asset with new asset
	purchaseOrder := PurchaseOrder{
		PoId:            poId,
		QuoteId:         quoteId,
		VendorLimitLeft: vendorLimitLeft,
		VendorLimitUsed: VendorLimitUsed,
		BankerId:        bankerId,
		PoDate:          poDate,
		Currency:        currency,
		PoStatus:        poStatus,
	}
	poJSON, err := json.Marshal(purchaseOrder)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(poId, poJSON)
}

// UpdateInvoice updates an existing Invoice in the world state with provided parameters.
func (s *SmartContract) UpdateInvoice(ctx contractapi.TransactionContextInterface, invoiceId string,
	quoteId string, vendorId string, buyerId string, poId string, invoiceStatus string,
	requestForFinance string, invoiceRaisedDate string, invoiceAcceptedDate string) error {
	exists, err := s.IsExists(ctx, invoiceId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", invoiceId)
	}

	// overwriting original asset with new asset
	invoice := Invoice{
		InvoiceId:           invoiceId,
		QuoteId:             quoteId,
		VendorId:            vendorId,
		BuyerId:             buyerId,
		PoId:                poId,
		InvoiceStatus:       invoiceStatus,
		RequestForFinance:   requestForFinance,
		InvoiceRaisedDate:   invoiceRaisedDate,
		InvoiceAcceptedDate: invoiceAcceptedDate,
	}
	invoiceJSON, err := json.Marshal(invoice)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(invoiceId, invoiceJSON)
}

// DeleteAsset deletes an given asset from the world state.
// func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
// 	exists, err := s.AssetExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}

// 	return ctx.GetStub().DelState(id)
// }

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) IsExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	quoteJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return quoteJSON != nil, nil
}

// // TransferAsset updates the owner field of asset with given id in world state.
// func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
// 	asset, err := s.ReadAsset(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	asset.Owner = newOwner
// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.GetStub().PutState(id, assetJSON)
// }

// GetAllAssets returns all assets found in world state
// func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
// 	// range query with empty string for startKey and endKey does an
// 	// open-ended query of all assets in the chaincode namespace.
// 	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	var assets []*Asset
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return nil, err
// 		}

// 		var asset Asset
// 		err = json.Unmarshal(queryResponse.Value, &asset)
// 		if err != nil {
// 			return nil, err
// 		}
// 		assets = append(assets, &asset)
// 	}

// 	return assets, nil
// }

func main() {
	vendorChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := vendorChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
