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
