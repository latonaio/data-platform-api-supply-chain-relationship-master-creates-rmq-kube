package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-supply-chain-relationship-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var transaction *[]dpfm_api_output_formatter.Transaction
	var billingRelation *[]dpfm_api_output_formatter.BillingRelation
	var deliveryPlantRelationProductMRPArea *[]dpfm_api_output_formatter.DeliveryPlantRelationProductMRPArea
	var deliveryPlantRelation *[]dpfm_api_output_formatter.DeliveryPlantRelation
	var deliveryRelation *[]dpfm_api_output_formatter.DeliveryRelation
	var paymentRelation *[]dpfm_api_output_formatter.PaymentRelation
	var productionPlantRelationMRP *[]dpfm_api_output_formatter.ProductionPlantRelationMRP
	var productionPlantRelation *[]dpfm_api_output_formatter.ProductionPlantRelation
	var stockConfPlantRelationProduct *[]dpfm_api_output_formatter.StockConfPlantRelationProduct
	var stockConfPlantRelation *[]dpfm_api_output_formatter.StockConfPlantRelation
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalCreateSql(nil, mtx, input, output, errs, log)
		case "Transaction":
			transaction = c.transactionCreateSql(nil, mtx, input, output, errs, log)
		case "BillingRelation":
			billingRelation = c.billingRelationCreateSql(nil, mtx, input, output, errs, log)
		case "DeliveryPlantRelationProductMRPArea":
			deliveryPlantRelationProductMRPArea = c.deliveryPlantRelationProductMRPAreaCreateSql(nil, mtx, input, output, errs, log)
		case "DeliveryPlantRelation":
			deliveryPlantRelation = c.deliveryPlantRelationCreateSql(nil, mtx, input, output, errs, log)
		case "DeliveryRelation":
			deliveryRelation = c.deliveryRelationCreateSql(nil, mtx, input, output, errs, log)
		case "PaymentRelation":
			paymentRelation = c.paymentRelationCreateSql(nil, mtx, input, output, errs, log)
		case "ProductionPlantRelationMRP":
			productionPlantRelationMRP = c.productionPlantRelationMRPCreateSql(nil, mtx, input, output, errs, log)
		case "ProductionPlantRelation":
			productionPlantRelation = c.productionPlantRelationCreateSql(nil, mtx, input, output, errs, log)
		case "StockConfPlantRelationProduct":
			stockConfPlantRelationProduct = c.stockConfPlantRelationProductCreateSql(nil, mtx, input, output, errs, log)
		case "StockConfPlantRelation":
			stockConfPlantRelation = c.stockConfPlantRelationCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:                             general,
		BillingRelation:                     billingRelation,
		DeliveryPlantRelationProductMRPArea: deliveryPlantRelationProductMRPArea,
		DeliveryPlantRelation:               deliveryPlantRelation,
		DeliveryRelation:                    deliveryRelation,
		PaymentRelation:                     paymentRelation,
		ProductionPlantRelationMRP:          productionPlantRelationMRP,
		ProductionPlantRelation:             productionPlantRelation,
		StockConfPlantRelationProduct:       stockConfPlantRelationProduct,
		StockConfPlantRelation:              stockConfPlantRelation,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_processing_formatter.GeneralUpdates
	var transaction *[]dpfm_api_output_formatter.Transaction
	var billingRelation *[]dpfm_api_output_formatter.BillingRelationUpdates
	var deliveryPlantRelationProductMRPArea *[]dpfm_api_output_formatter.DeliveryPlantRelationProductMRPAreaUpdates
	var deliveryPlantRelation *[]dpfm_api_output_formatter.DeliveryPlantRelationUpdates
	var deliveryRelation *[]dpfm_api_output_formatter.DeliveryRelationUpdates
	var paymentRelation *[]dpfm_api_output_formatter.PaymentRelationUpdates
	var productionPlantRelationMRP *[]dpfm_api_output_formatter.ProductionPlantRelationMRPUpdates
	var productionPlantRelation *[]dpfm_api_output_formatter.ProductionPlantRelationUpdates
	var stockConfPlantRelationProduct *[]dpfm_api_output_formatter.StockConfPlantRelationProductUpdates
	var stockConfPlantRelation *[]dpfm_api_output_formatter.StockConfPlantRelationUpdates
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalUpdateSql(mtx, input, output, errs, log)
		case "BusinessPartner":
			businessPartner = c.businessPartnerUpdateSql(mtx, input, output, errs, log)
		case "DeliveryPlantRelationProductMRPArea":
			deliveryPlantRelationProductMRPArea = c.deliveryPlantRelationProductMRPArea(mtx, input, output, errs, log)
		case "DeliveryPlantRelation":
			deliveryPlantRelation = c.deliveryPlantRelation(mtx, input, output, errs, log)
		case "DeliveryRelation":
			deliveryRelation = c.deliveryRelation(mtx, input, output, errs, log)
		case "PaymentRelation":
			paymentRelation = c.paymentRelation(mtx, input, output, errs, log)
		case "ProductionPlantRelationMRP":
			productionPlantRelationMRP = c.productionPlantRelationMRP(mtx, input, output, errs, log)
		case "ProductionPlantRelation":
			productionPlantRelation = c.productionPlantRelation(mtx, input, output, errs, log)
		case "StockConfPlantRelationProduct":
			stockConfPlantRelationProduct = c.stockConfPlantRelationProduct(mtx, input, output, errs, log)
		case "StockConfPlantRelation":
			stockConfPlantRelation = c.stockConfPlantRelation(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:                             general,
		BillingRelation:                     billingRelation,
		DeliveryPlantRelationProductMRPArea: deliveryPlantRelationProductMRPArea,
		DeliveryPlantRelation:               deliveryPlantRelation,
		DeliveryRelation:                    deliveryRelation,
		PaymentRelation:                     paymentRelation,
		ProductionPlantRelationMRP:          productionPlantRelationMRP,
		ProductionPlantRelation:             productionPlantRelation,
		StockConfPlantRelationProduct:       stockConfPlantRelationProduct,
		StockConfPlantRelation:              stockConfPlantRelation,
	}

	return data
}

func (c *DPFMAPICaller) generalCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_supply_chain_relationship_master_general_dataの更新
	generalData := input.General
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "SupplyChainRelationshipMasterGeneral", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Header Data cannot insert")
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	return
}

func (c *DPFMAPICaller) generalUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	req := dpfm_api_processing_formatter.ConvertToGeneralUpdates(input.General)

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "InvoiceDocumentHeader", "runtime_session_id": 123})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Header Data cannot insert")
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToGeneralUpdates(req)

	return data
}
