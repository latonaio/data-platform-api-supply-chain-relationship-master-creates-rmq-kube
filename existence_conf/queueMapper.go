package existence_conf

import "data-platform-api-supply-chain-relationship-creates-rmq-kube/config"

type exconfQueueMapper map[string]string

func NewExconfQueueMapper(c *config.Conf) exconfQueueMapper {
	m := exconfQueueMapper{}
	ecQNames := c.RMQ.QueueToExConf()
	m["SupplyChainRelationshipGeneral"] = ecQNames[0%len(ecQNames)]
	m["BusinessPartnerGeneral"] = ecQNames[0%len(ecQNames)]
	m["TransactionCurrency"] = ecQNames[1%len(ecQNames)]
	m["Incoterms"] = ecQNames[1%len(ecQNames)]
	m["PaymentMethod"] = ecQNames[1%len(ecQNames)]
	m["PaymentTerms"] = ecQNames[1%len(ecQNames)]
	return m
}

func (m exconfQueueMapper) getQueueNameByConfContent(content string) string {
	return m[content]
}
