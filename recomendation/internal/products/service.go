package products

type ServiceProducs struct {
	RepositoryProducts *RepositoryProducts
}

func (sp ServiceProducs) GetOrdersMetrics(product_id int64) {
	sp.RepositoryProducts.GetOrdersMetrics(product_id)
}
