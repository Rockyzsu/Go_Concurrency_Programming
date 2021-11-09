package sse

//NewSSE 实例化Broker并启动监听
func NewSSE() (broker *Broker) {
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}
	//开启监听
	go broker.listen()

	return
}
