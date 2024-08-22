package platform

import "github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/model"

type Product struct {
	model.PurchaseCatalog
	P_ID       string `json:"id"`
	ItemID     string `json:"item_id"`
	ProviderID string `json:"provider_id"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
}
