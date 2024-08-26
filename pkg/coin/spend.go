package coin

import bls12381 "github.com/kilic/bls12-381"

type Spend struct {
	Coin         *Coin
	PuzzleReveal Program
	Solution     Program
}

type SpendBundle struct {
	Spends             []Spend
	AggregateSignature *bls12381.PointG2
}

func (sb *SpendBundle) Aggregate(sbs ...*SpendBundle) *SpendBundle {

	var res SpendBundle

	var das = bls12381.NewG2()
	var aggSig = das.New()

	for _, sb := range sbs {
		res.Spends = append(res.Spends, sb.Spends...)
		aggSig = bls12381.NewG2().Add(das.New(), aggSig, sb.AggregateSignature)
	}

	res.AggregateSignature = aggSig

	return &res
}
