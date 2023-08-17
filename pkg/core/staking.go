package core

import (
	"github.com/tonkeeper/tongo"
)

type Nominator struct {
	Pool                  tongo.AccountID
	Member                tongo.AccountID
	MemberBalance         int64
	MemberPendingDeposit  int64
	MemberPendingWithdraw int64
	MemberWithdraw        int64
}

type TFPool struct {
	Address         tongo.AccountID
	ValidatorStake  int64
	NominatorsStake int64
	// TotalAmount = ValidatorStake + NominatorsStake
	TotalAmount       int64
	MinNominatorStake int64
	ValidatorShare    uint32
	StakeAt           uint32
	Nominators        int
	MaxNominators     int
	VerifiedSources   bool
}

type LiquidPool struct {
	Name            string
	Address         tongo.AccountID
	TotalAmount     int64
	VerifiedSources bool
	JettonMaster    tongo.AccountID
}
