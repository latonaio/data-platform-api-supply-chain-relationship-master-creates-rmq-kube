package dpfm_api_processing_formatter

func ConvertToGeneralUpdates(generalUpdates GeneralUpdates) *GeneralUpdates {
	data := generalUpdates

	return &GeneralUpdates{
		Buyer:               data.Buyer,
		Seller:              data.Seller,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func ConvertToTransactionUpdates(transactionUpdates TransactionUpdates) *TransactionUpdates {
	data := transactionUpdates

	return &TransactionUpdates{
		TransactionCurrency: data.TransactionCurrency,
		PaymentTerms:        data.PaymentTerms,
		PaymentMethod:       data.PaymentMethod,
		Incoterms:           data.Incoterms,
		QuotationIsBlocked:  data.QuotationIsBlocked,
		OrderIsBlocked:      data.OrderIsBlocked,
		DeliveryIsBlocked:   data.DeliveryIsBlocked,
		BillingIsBlocked:    data.BillingIsBlocked,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func ConvertToDeliveryRelationUpdates(deliveryRelationUpdates DeliveryRelationUpdates) *DeliveryRelationUpdates {
	data := deliveryRelationUpdates

	return &DeliveryRelationUpdates{
		DeliverToParty:   data.DeliverToParty,
		DeliverFromParty: data.DeliverFromParty,
		DefaultRelation:  data.DefaultRelation,
	}
}

func ConvertToPaymentRelationUpdates(paymentRelationUpdates PaymentRelationUpdates) *PaymentRelationUpdates {
	data := paymentRelationUpdates

	return &PaymentRelationUpdates{
		Payer:                 data.Payer,
		Payee:                 data.Payee,
		DefaultRelation:       data.DefaultRelation,
		PayerHouseBank:        data.PayerHouseBank,
		PayerHouseBankAccount: data.PayerHouseBankAccount,
		PayeeHouseBank:        data.PayeeHouseBank,
		PayeeHouseBankAccount: data.PayeeHouseBankAccount,
		IsMarkedForDeletion:   data.IsMarkedForDeletion,
	}
}

func ConvertToBillingRelationUpdates(billingRelationUpdates BillingRelationUpdates) *BillingRelationUpdates {
	data := billingRelationUpdates

	return &BillingRelationUpdates{
		BillToParty:                  data.BillToParty,
		BillFromParty:                data.BillFromParty,
		DefaultRelation:              data.DefaultRelation,
		TransactionTaxClassification: data.TransactionTaxClassification,
		IsMarkedForDeletion:          data.IsMarkedForDeletion,
	}
}

func ConvertToDeliveryPlantRelationUpdates(deliveryPlantRelationUpdates DeliveryPlantRelationUpdates) *DeliveryPlantRelationUpdates {
	data := deliveryPlantRelationUpdates

	return &DeliveryPlantRelationUpdates{
		DeliverToPlant:      data.DeliverToPlant,
		DeliverFromPlant:    data.DeliverFromPlant,
		DefaultRelation:     data.DefaultRelation,
		MRPType:             data.MRPType,
		MRPController:       data.MRPController,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func ConvertToDeliveryPlantRelationProductUpdates(deliveryPlantRelationProductUpdates DeliveryPlantRelationProductUpdates) *DeliveryPlantRelationProductUpdates {
	data := deliveryPlantRelationProductUpdates

	return &DeliveryPlantRelationProductUpdates{
		DeliverToPlantStorageLocation:             data.DeliverToPlantStorageLocation,
		DeliverFromPlantStorageLocation:           data.DeliverFromPlantStorageLocation,
		DeliveryUnit:                              data.DeliveryUnit,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
		PlanningTimeFence:                         data.PlanningTimeFence,
		MRPPlanningCalendar:                       data.MRPPlanningCalendar,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
		MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
		StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
		IsAutoOrderCreationAllowed:                data.IsAutoOrderCreationAllowed,
		IsOrderAcknowledgementRequired:            data.IsOrderAcknowledgementRequired,
		IsMarkedForDeletion:                       data.IsMarkedForDeletion,
	}
}

func ConvertToDeliveryPlantRelationProductMRPAreaUpdates(deliveryPlantRelationProductMRPAreaUpdates DeliveryPlantRelationProductMRPAreaUpdates) *DeliveryPlantRelationProductMRPAreaUpdates {
	data := deliveryPlantRelationProductMRPAreaUpdates

	return &DeliveryPlantRelationProductMRPAreaUpdates{
		DeliverToPlantStorageLocation:             data.DeliverToPlantStorageLocation,
		DeliverFromPlantStorageLocation:           data.DeliverFromPlantStorageLocation,
		DeliveryUnit:                              data.DeliveryUnit,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
		PlanningTimeFence:                         data.PlanningTimeFence,
		MRPPlanningCalendar:                       data.MRPPlanningCalendar,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
		MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
		StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
		IsAutoOrderCreationAllowed:                data.IsAutoOrderCreationAllowed,
		IsOrderAcknowledgementRequired:            data.IsOrderAcknowledgementRequired,
		IsMarkedForDeletion:                       data.IsMarkedForDeletion,
	}
}

func ConvertToStockConfPlantRelationUpdates(stockConfPlantRelationUpdates StockConfPlantRelationUpdates) *StockConfPlantRelationUpdates {
	data := stockConfPlantRelationUpdates

	return &StockConfPlantRelationUpdates{
		StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:           data.StockConfirmationPlant,
		DefaultRelation:                  data.DefaultRelation,
		MRPType:                          data.MRPType,
		MRPController:                    data.MRPController,
		IsMarkedForDeletion:              data.IsMarkedForDeletion,
	}
}

func ConvertToStockConfPlantRelationProductUpdates(stockConfPlantRelationProductUpdates StockConfPlantRelationProductUpdates) *StockConfPlantRelationProductUpdates {
	data := stockConfPlantRelationProductUpdates

	return &StockConfPlantRelationProductUpdates{
		MRPType:             data.MRPType,
		MRPController:       data.MRPController,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func ConvertToProductionPlantRelationUpdates(productionPlantRelationUpdates ProductionPlantRelationUpdates) *ProductionPlantRelationUpdates {
	data := productionPlantRelationUpdates

	return &ProductionPlantRelationUpdates{
		ProductionPlantBusinessPartner: data.ProductionPlantBusinessPartner,
		ProductionPlant:                data.ProductionPlant,
		DefaultRelation:                data.DefaultRelation,
		MRPType:                        data.MRPType,
		MRPController:                  data.MRPController,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}

func ConvertToProductionPlantRelationProductUpdates(productionPlantRelationProductUpdates ProductionPlantRelationProductUpdates) *ProductionPlantRelationProductUpdates {
	data := productionPlantRelationProductUpdates

	return &ProductionPlantRelationProductUpdates{
		MRPType:             data.MRPType,
		MRPController:       data.MRPController,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}
