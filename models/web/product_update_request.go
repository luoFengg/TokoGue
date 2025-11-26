package web

type ProductUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}