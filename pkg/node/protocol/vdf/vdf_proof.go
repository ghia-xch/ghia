package vdf

type VDFProof struct {
	WitnessType          uint8
	Witness              []byte
	NormalizedToIdentity bool
}
