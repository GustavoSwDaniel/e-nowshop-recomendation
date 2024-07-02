package orderitens

import "github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/products"

type ServiceOrdersItens struct {
	RepositoryOrdersItems *RepositoryOrdersItems
	PorductService        *products.ServiceProducs
}

func (soi ServiceOrdersItens) GetOrdersMetrics(product_id int64) {
	recomendationMetrics := soi.RepositoryOrdersItems.GetOrdersMetrics(product_id)
	soi.PorductService.UpdateProductRecomendation(recomendationMetrics, product_id)
}
