package private

import (
	"github.com/LIJI-MAX/okx/events"
	"github.com/LIJI-MAX/okx/models/account"
	"github.com/LIJI-MAX/okx/models/trade"
)

type (
	Account struct {
		Arg      *events.Argument   `json:"arg"`
		Balances []*account.Balance `json:"data"`
	}
	Position struct {
		Arg       *events.Argument    `json:"arg"`
		Positions []*account.Position `json:"data"`
	}
	BalanceAndPosition struct {
		Arg                 *events.Argument              `json:"arg"`
		BalanceAndPositions []*account.BalanceAndPosition `json:"data"`
	}
	Order struct {
		Arg    *events.Argument `json:"arg"`
		Orders []*trade.Order   `json:"data"`
	}
	AlgoOrder struct {
		Arg    *events.Argument   `json:"arg"`
		Orders []*trade.AlgoOrder `json:"data"`
	}
)
