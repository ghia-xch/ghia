package coin

import (
	"github.com/ghia-xch/ghia/pkg/coin/program"
	bls12381 "github.com/kilic/bls12-381"
)

type Spend struct {
	Coin         *Coin
	PuzzleReveal program.Program
	Solution     program.Program
}

type SpendBundle struct {
	Spends             []Spend
	AggregateSignature *bls12381.PointG2
}

func (sb *SpendBundle) Aggregate(sbs ...*SpendBundle) *SpendBundle {

	var res SpendBundle

	var das = bls12381.NewG2()

	for _, sb := range sbs {
		res.Spends = append(res.Spends, sb.Spends...)
		res.AggregateSignature = bls12381.NewG2().Add(das.New(), res.AggregateSignature, sb.AggregateSignature)
	}

	return &res
}
