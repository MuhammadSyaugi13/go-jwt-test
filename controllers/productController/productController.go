package productcontroller

import (
	"go-jwt-web/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]any{
		{
			"id":           1,
			"nama_product": "Kemeja",
			"stok":         120,
		},
		{
			"id":           2,
			"nama_product": "Celana",
			"stok":         110,
		},
		{
			"id":           3,
			"nama_product": "Topi",
			"stok":         90,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
