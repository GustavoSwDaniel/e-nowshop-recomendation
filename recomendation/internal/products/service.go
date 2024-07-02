package products

type ServiceProducs struct {
	RepositoryProducts *RepositoryProducts
}

func (sp *ServiceProducs) UpdateProductRecomendation(recomendation []byte, product_id int64) {
	sp.RepositoryProducts.UpdateRecomendation(recomendation, product_id)
}
