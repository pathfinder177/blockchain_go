package tcp

type tcpGateway struct {
	thisAddr  string
	otherAddr string
}

func New(thisAddr, otherAddr string) *tcpGateway {
	return &tcpGateway{
		thisAddr:  thisAddr,
		otherAddr: otherAddr,
	}
}
