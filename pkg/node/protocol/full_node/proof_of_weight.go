package full_node

type RequestProofOfWeight struct {
	TotalNumberOfBlocks uint32
	Tip                 [32]byte
}

type WeightProof struct{}

type RespondProofOfWeight struct {
	WeightProof WeightProof
	Tip         [32]byte
}
