package subaccount

import "github.com/aiviaio/okex"

type (
	ViewList struct {
		SubAcct string `json:"subAcct,omitempty"`
		Enable  bool   `json:"enable,omitempty"`
		After   int64  `json:"after,omitempty,string"`
		Before  int64  `json:"before,omitempty,string"`
		Limit   int64  `json:"limit,omitempty,string"`
	}
	CreateAPIKey struct {
		Pwd        string            `json:"pwd"`
		SubAcct    string            `json:"subAcct"`
		Label      string            `json:"label"`
		Passphrase string            `json:"Passphrase"`
		IP         []string          `json:"ip,omitempty"`
		Perm       okex.APIKeyAccess `json:"perm,omitempty"`
	}
	QueryAPIKey struct {
		APIKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	DeleteAPIKey struct {
		Pwd     string `json:"pwd"`
		APIKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	GetBalance struct {
		SubAcct string `json:"subAcct"`
	}
	HistoryTransfer struct {
		Ccy     string            `json:"ccy,omitempty"`
		SubAcct string            `json:"subAcct,omitempty"`
		After   int64             `json:"after,omitempty,string"`
		Before  int64             `json:"before,omitempty,string"`
		Limit   int64             `json:"limit,omitempty,string"`
		Type    okex.TransferType `json:"type,omitempty,string"`
	}
	ManageTransfers struct {
		Ccy            string           `json:"ccy"`
		FromSubAccount string           `json:"fromSubAccount"`
		ToSubAccount   string           `json:"tiSubAccount"`
		Amt            float64          `json:"amt,string"`
		From           okex.AccountType `json:"from,string"`
		To             okex.AccountType `json:"to,string"`
	}
	CreateSubAccount struct {
		SubAcct string `json:"subAcct"`
		Label   string `json:"label,omitempty"`
	}
	DeleteSubAccount struct {
		SubAcct string `json:"subAcct"`
	}
	CreatAPIKeySubAccount struct {
		SubAcct    string            `json:"subAcct"`
		Label      string            `json:"label"`
		Passphrase string            `json:"passphrase"`
		IP         []string          `json:"ip,omitempty"`
		Perm       okex.APIKeyAccess `json:"perm,omitempty"`
	}
	UpdateAPIKEySubAccount struct {
		SubAcct string `json:"subAcct"`
		Label   string `json:"label,omitempty"`
		APIKey  string `json:"apiKey"`
		Perm    string `json:"perm,omitempty"`
		IP      string `json:"ip,omitempty"`
	}
	DeleteAPIKeySubAccount struct {
		SubAcct string `json:"subAcct"`
		APIKey  string `json:"apiKey"`
	}
	SetLevelSubAccount struct {
		SubAcct string `json:"subAcct"`
		AcctLv  string `json:"acctLv"` //Account level 1: Simple 2: Single-currency margin 3: Multi-currency margin 4：Portfolio margin
	}
	SetFeeRateSubAccount struct {
		SubAcct  string `json:"subAcct,omitempty"`
		InstType string `json:"instType,omitempty"`
		MgnType  string `json:"mgnType,omitempty"`
		ChgType  string `json:"chgType"`
		ChgTaker string `json:"chgTaker"`
		ChgMaker string `json:"chgMaker"`
		EffDate  string `json:"effDate,omitempty"`
	}
	CreateDepositAddress struct {
		SubAcct  string `json:"subAcct"`
		Ccy      string `json:"ccy"`
		Chain    string `json:"chain,omitempty"`
		AddrType string `json:"addrType,omitempty"` // 1: Regular address, 2:SegWit address (Only applicable to BTC/LTC), Default is 1
		TO       string `json:"to,omitempty"`       // 6:Funding, 18:Trading account, Default is 6
	}
	UpdateDepositAddress struct {
		SubAcct  string `json:"subAcct"`
		Ccy      string `json:"ccy"`
		Chain    string `json:"chain,omitempty"`
		AddrType string `json:"addrType"` // 1: Regular address, 2:SegWit address (Only applicable to BTC/LTC), Default is 1
		TO       string `json:"to"`       // 6:Funding, 18:Trading account, Default is 6
	}
	GetDepositAddress struct {
		SubAcct string `json:"subAcct"`
		Ccy     string `json:"ccy"`
	}
)
