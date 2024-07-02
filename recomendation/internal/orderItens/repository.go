package orderitens

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/dto/products"
)

type RepositoryOrdersItems struct {
	Conn *sql.DB
}

func (r *RepositoryOrdersItems) GetOrdersMetrics(product_id int64) []byte {
	rows, err := r.Conn.Query(`WITH orders_with_product AS (
		SELECT DISTINCT order_id
		FROM order_items
		WHERE product_id = $1
	),
	products_bought_with_product AS (
		SELECT oi.product_id, COUNT(DISTINCT oi.order_id) AS product_count
		FROM order_items oi
		JOIN orders_with_product om ON oi.order_id = om.order_id
		WHERE oi.product_id != $1
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
	ORDER BY percentage DESC;`, product_id)
	if err != nil {
		log.Fatalf("Erro na consulta no banco: %v", err)
	}
	defer rows.Close()
	var metrics = products.ProductsMetrics{}
	data := make(map[int64]float64)
	for rows.Next() {
		rows.Scan(&metrics.ProductId, &metrics.Percentage)
		fmt.Printf("%v", metrics)
		data[metrics.ProductId] = metrics.Percentage
	}
	jsonString, _ := json.Marshal(data)
	return jsonString
}
