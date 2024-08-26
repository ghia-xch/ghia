package full_node

import "github.com/ghia-xch/ghia/pkg/node/protocol"

type NewTransaction struct {
	TransactionId protocol.Hash
	Cost          uint64
	Fees          uint64
}

func (n *NewTransaction) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {

	enc.Reset(protocol.NewPeak, nil)

	return enc.Encode(
		n.TransactionId,
		n.Cost,
		n.Fees,
	)
}

func (n *NewTransaction) Decode(dec *protocol.MessageDecoder) (err error) {

	if n.TransactionId, err = dec.ParseHash(); err != nil {
		return
	}

	if n.Cost, err = dec.ParseUint64(); err != nil {
		return
	}

	if n.Fees, err = dec.ParseUint64(); err != nil {
		return
	}

	return
}

type RequestTransaction struct {
	TransactionId protocol.Hash
}

func (n RequestTransaction) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {

	enc.Reset(protocol.RequestTransaction, nil)

	return enc.Encode(
		n.TransactionId,
	)
}

func (n *RequestTransaction) Decode(dec *protocol.MessageDecoder) (err error) {

	if n.TransactionId, err = dec.ParseHash(); err != nil {
		return
	}

	return nil
}

func CreateRequestTransaction(transactionId protocol.Hash) (em protocol.EncodedMessage) {
	em, _ = RequestTransaction{TransactionId: transactionId}.Encode(protocol.NewMessageEncoder(38))
	return
}

//type CoinSpend struct {
//}
//
//type SpendBundle struct{
//	CoinSpends []CoinSpend
//	AggSignature []byte
//}
//
//type RespondTransaction struct {
//	Transaction SpendBundle
//}
//
//func (n *RespondTransaction) Encode(enc *protocol.MessageEncoder) (em protocol.EncodedMessage, err error) {
//	return enc.Encode(
//		n.Transaction,
//	)
//}
//
//func (n *RespondTransaction) Decode(dec *protocol.MessageDecoder) (err error) {
//
//	if n.Transaction.
//}