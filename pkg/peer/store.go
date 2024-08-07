package peer

type Peer interface {
}

type Store interface {
	Reset() (err error)

	Add(peers ...*Peer) (err error)

	Remove(peers ...*Peer) (err error)

	Touch(peers ...*Peer) (err error)

	GetByHost(host string) (peer *Peer, err error)
}
