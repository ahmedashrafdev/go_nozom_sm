package handler

import (

	"fmt"
	"net/http"
	"github.com/ahmedashrafdev/reports/model"
	"github.com/labstack/echo/v4"
)


func (h *Handler) GetDocNo(c echo.Context) error {

	req := new(model.DocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	// return c.JSON(http.StatusOK, "test")

	var DocNo []model.Doc
	rows, err := h.db.Raw("EXEC GetSdDocNo @DevNo = ?, @TrSerial = ?,@StoreCode = ?;", req.DevNo, req.TrSerial, req.StoreCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var doc model.Doc
		err = rows.Scan(
			&doc.DocNo,
		)
		print(rows)
		if err != nil {
			return c.JSON(http.StatusOK, 1)
		}
		DocNo = append(DocNo, doc)
	}

	return c.JSON(http.StatusOK, DocNo[0].DocNo+1)
}


func (h *Handler) GetOpenDocs(c echo.Context) error {

	req := new(model.OpenDocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var OpenDocs []model.OpenDoc
	rows, err := h.db.Raw("EXEC GetOpenSdDocNo @DevNo = ?, @TrSerial = ?;", req.DevNo, req.TrSerial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var openDoc model.OpenDoc
		err = rows.Scan(
			&openDoc.DocNo,
			&openDoc.StoreCode,
			&openDoc.AccontSerial,
			&openDoc.TransSerial,
			&openDoc.AccountName,
			&openDoc.AccountCode,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		OpenDocs = append(OpenDocs, openDoc)
	}

	return c.JSON(http.StatusOK, OpenDocs)
}
func (h *Handler) GetDocItems(c echo.Context) error {

	req := new(model.DocItemsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var DocItems []model.DocItem
	rows, err := h.db.Raw("EXEC GetSdItems @DevNo = ?, @TrSerial = ?,@StoreCode = ? , @DocNo = ?;", req.DevNo, req.TrSerial, req.StoreCode, req.DocNo).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var docItem model.DocItem
		err = rows.Scan(
			&docItem.Serial,
			&docItem.Qnt,
			&docItem.Item_BarCode,
			&docItem.ItemName,
			&docItem.MinorPerMajor,
			&docItem.ByWeight,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		DocItems = append(DocItems, docItem)
	}

	return c.JSON(http.StatusOK, DocItems)
}

func (h *Handler) DeleteItem(c echo.Context) error {

	req := new(model.DeleteItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	print(req)
	rows, err := h.db.Raw("EXEC DeleteSdItem  @Serial = ?; ", req.Serial).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) InsertItem(c echo.Context) error {

	req := new(model.InsertItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	print(req)
	rows, err := h.db.Raw(
		"EXEC InsertSdDocNo  @DNo = ? ,@TrS = ? ,@AccS = ? ,@ItmS =?  ,@Qnt = ? ,@StCode = ? ,@InvNo = ? ,@ItmBarCode = ? ,@DevNo = ?,@StCode2 = ?; ", req.DNo, req.TrS, req.AccS, req.ItmS, req.Qnt, req.StCode, req.InvNo, req.ItmBarCode, req.DevNo, req.StCode2).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rows)
}


func (h *Handler) CashTryStores(c echo.Context) error {
	var stores []model.CashtryStores
	// return c.JSON(http.StatusOK, "test")
	rows, err := h.db.Raw("EXEC GetStoreName").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer rows.Close()
	for rows.Next() {
		var store model.CashtryStores
		rows.Scan(&store.StoreCode, &store.StoreName)
		stores = append(stores, store)
	}

	return c.JSON(http.StatusOK, stores)
}

func (h *Handler) GetAccount(c echo.Context) error {

	req := new(model.GetAccountRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var accounts []model.Account
	rows, err := h.db.Raw("EXEC GetAccount @Code = ?, @Name = ? , @Type = ?", req.Code, req.Name, req.Type).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var account model.Account
		rows.Scan(&account.Serial, &account.AccountCode, &account.AccountName)
		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}

func (h *Handler) GetItem(c echo.Context) error {

	req := new(model.GetItemRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var items []model.Item
	rows, err := h.db.Raw("EXEC GetItemData @BCode = ?, @StoreCode = ?", req.BCode, req.StoreCode).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.Serial, &item.ItemName, &item.MinorPerMajor, &item.POSPP, &item.POSTP, &item.ByWeight)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

