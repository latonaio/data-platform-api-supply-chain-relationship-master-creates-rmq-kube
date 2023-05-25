package dpfm_api_processing_formatter

type Message struct {
	GeneralUpdates               *GeneralUpdates                 `json:"GeneralUpdates"`
	TransactionUpdates           *[]TransactionUpdates           `json:"TransactionUpdates"`
	BillingRelationUpdates       *[]BillingRelationUpdates       `json:"BillingRelationUpdates"`
	PaymentRelationUpdates       *[]PaymentRelationUpdates       `json:"PaymentRelationUpdates"`
	DeliveryRelationUpdates      *[]DeliveryRelationUpdates      `json:"DeliveryRelationUpdates"`
	DeliveryPlantRelationUpdates *[]DeliveryPlantRelationUpdates `json:"DeliveryPlantRelationUpdates"`
}

type GeneralUpdates struct {
	Buyer               *int  `json:"Buyer"`
	Seller              *int  `json:"Seller"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type TransactionUpdates struct {
	TransactionCurrency *string `json:"TransactionCurrency"`
	PaymentTerms        *string `json:"PaymentTerms"`
	PaymentMethod       *string `json:"PaymentMethod"`
	Incoterms           *string `json:"Incoterms"`
	QuotationIsBlocked  *bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked      *bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked   *bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked    *bool   `json:"BillingIsBlocked"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
type BillingRelationUpdates struct {
	BillToParty                  *int    `json:"BillToParty"`
	BillFromParty                *int    `json:"BillFromParty"`
	DefaultRelation              *bool   `json:"DefaultRelation"`
	TransactionTaxClassification *string `json:"TransactionTaxClassification"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}

type PaymentRelationUpdates struct {
	Payer                 *int    `json:"Payer"`
	Payee                 *int    `json:"Payee"`
	DefaultRelation       *bool   `json:"DefaultRelation"`
	PayerHouseBank        *string `json:"PayerHouseBank"`
	PayerHouseBankAccount *string `json:"PayerHouseBankAccount"`
	PayeeHouseBank        *string `json:"PayeeHouseBank"`
	PayeeHouseBankAccount *string `json:"PayeeHouseBankAccount"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryRelationUpdates struct {
	DeliverToParty      *int  `json:"DeliverToParty"`
	DeliverFromParty    *int  `json:"DeliverFromParty"`
	DefaultRelation     *bool `json:"DefaultRelation"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type DeliveryPlantRelationUpdates struct {
	DeliverToPlant      *string `json:"DeliverToPlant"`
	DeliverFromPlant    *string `json:"DeliverFromPlant"`
	DefaultRelation     *bool   `json:"DefaultRelation"`
	MRPType             *string `json:"MRPType"`
	MRPController       *string `json:"MRPController"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryPlantRelationProductUpdates struct {
	DeliverToPlantStorageLocation             *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           *string  `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              *string  `json:"DeliveryUnit"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsAutoOrderCreationAllowed                *bool    `json:"IsAutoOrderCreationAllowed"`
	IsOrderAcknowledgementRequired            *bool    `json:"IsOrderAcknowledgementRequired"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type DeliveryPlantRelationProductMRPAreaUpdates struct {
	DeliverToPlantStorageLocation             *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           *string  `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              *string  `json:"DeliveryUnit"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsAutoOrderCreationAllowed                *bool    `json:"IsAutoOrderCreationAllowed"`
	IsOrderAcknowledgementRequired            *bool    `json:"IsOrderAcknowledgementRequired"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type StockConfPlantRelationUpdates struct {
	StockConfirmationBusinessPartner *int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           *string `json:"StockConfirmationPlant"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	MRPType                          *string `json:"MRPType"`
	MRPController                    *string `json:"MRPController"`
	IsMarkedForDeletion              *bool   `json:"IsMarkedForDeletion"`
}

type StockConfPlantRelationProductUpdates struct {
	MRPType             *string `json:"MRPType"`
	MRPController       *string `json:"MRPController"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type ProductionPlantRelationUpdates struct {
	ProductionPlantBusinessPartner *int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                *string `json:"ProductionPlant"`
	DefaultRelation                *bool   `json:"DefaultRelation"`
	MRPType                        *string `json:"MRPType"`
	MRPController                  *string `json:"MRPController"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

type ProductionPlantRelationProductUpdates struct {
	MRPType             *string `json:"MRPType"`
	MRPController       *string `json:"MRPController"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
