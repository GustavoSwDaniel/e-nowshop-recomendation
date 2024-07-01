package products

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/dto/products"
)

type RepositoryProducts struct {
	Conn *sql.DB
}

func (r *RepositoryProducts) GetOrdersMetrics(product_id int64) {
	rows, err := r.Conn.Query(`WITH orders_with_product AS (
		SELECT DISTINCT order_id
		FROM order_items
		WHERE product_id = $1
	),
	products_bought_with_product AS (
		SELECT oi.product_id, COUNT(DISTINCT oi.order_id) AS product_count
		FROM order_items oi
		JOIN orders_with_product o54 ON oi.order_id = o54.order_id
		WHERE oi.product_id != 54
		GROUP BY oi.product_id
	),
	total_orders_with_product AS (
		SELECT COUNT(DISTINCT order_id) AS total_orders
		FROM orders_with_product
	)
	SELECT p.product_id,
		   ROUND((p.product_count::decimal / t.total_orders) * 100, 2) AS percentage
	FROM products_bought_with_product p
	CROSS JOIN total_orders_with_product t
	ORDER BY percentage DESC;`, 2)
	if err != nil {
		log.Fatalf("Erro na consulta no banco: %v", err)
	}
	defer rows.Close()
	var metrics = products.ProductsMetrics{}
	for rows.Next() {

		rows.Scan(&metrics.Percentage, &metrics.Percentage)
		fmt.Printf("%v", metrics)
	}
}
