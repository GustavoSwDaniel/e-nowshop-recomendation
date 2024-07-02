package products

import (
	"database/sql"
	"fmt"
	"log"
)

type RepositoryProducts struct {
	Conn *sql.DB
}

func (rp *RepositoryProducts) UpdateRecomendation(recomendation_metrics []byte, product_id int64) {
	result, err := rp.Conn.Exec("UPDATE products SET recomendations = $1 WHERE id = $2", recomendation_metrics, product_id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)
}
