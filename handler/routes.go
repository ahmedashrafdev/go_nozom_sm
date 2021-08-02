package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	cashtry := v1.Group("/cashtry")
	cashtry.GET("/stores", h.CashTryStores)
	v1.POST("/get-account", h.GetAccount)
	v1.POST("/get-item", h.GetItem) // done
	v1.POST("/get-doc", h.GetDocNo)
	v1.POST("/get-doc-items", h.GetDocItems)
	v1.POST("/insert-item", h.InsertItem)
	v1.POST("/delete-item", h.DeleteItem)
	v1.POST("/get-docs", h.GetOpenDocs)
	v1.POST("/close-doc", h.CloseDoc)

}
