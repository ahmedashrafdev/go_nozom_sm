package model

type Doc struct {
	DocNo int
}
type InvReq struct {
	BCode string
}
type UpdatePrepareReq struct {
	QPrep   float64
	ISerial int
	HSerial int
	EmpCode int
}
type GetMsgsRequest struct {
	EmpSerial int
	BonSerial int
}

type UpdatePrepareResp struct {
	Prepared    bool
	QntPrepared float64
	Qnt         float64
}
type InventorySessionReq struct {
	StoreCode int
}
type InventorySessionResp struct {
	SessionNo int
}

type InvoiceItem struct {
	BonSer        string
	Qnt           float64
	Price         float64
	IsPrepared    bool
	QntPrepare    float64
	ItemCode      string
	GroupCode     string
	MinorPerMajor int
	ItemName      string
	ItemSerial    string
}

type PrepareDocResp struct {
	DocNo       string
	AccountName string
	AccountCode int
}

type DocReq struct {
	DevNo     int
	TrSerial  int
	StoreCode int
}
type CloseDocReq struct {
	DevNo int
	Trans int
	DocNo int
}
type ClosePrepareDocReq struct {
	HSerial int
	EmpCode int
}
type ClosePrepareDocResp struct {
	Close bool
}

type OpenDoc struct {
	DocNo        int
	StoreCode    int
	AccontSerial int
	TransSerial  int
	AccountName  string
	AccountCode  int
	DevNo        int
}
type OpenDocReq struct {
	DevNo    int
	TrSerial int
	StCode   int
}
