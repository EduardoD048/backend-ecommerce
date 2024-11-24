package purchase_model

import "compra/internal/app/api/model/product_model"

type Purchase struct {
	ID       uint                    `json:"id"`
	Products []product_model.Product `json:"products" gorm:"many2many:purchase_products;"`
}
