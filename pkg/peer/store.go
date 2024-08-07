package peer

type Store interface {
	Reset() (err error)

	Add(peers ...*Client) (err error)

	Remove(peers ...*Client) (err error)

	Touch(peers ...*Client) (err error)

	GetByHost(host string) (peer *Client, err error)
}
