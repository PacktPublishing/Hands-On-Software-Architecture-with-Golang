
type MyMessageHandler struct {}

func (h *MyMessageHandler) HandleMessage(m *nsq.Message) error {
    m.DisableAutoResponse()
    delegateChannel <- m
    return nil
}

go func() {
    for m := range delegateChannel {
        err := doSomeWork(m) // some long winded tasks
        if err != nil {
            m.Requeue(-1)
            continue
        }
        m.Finish()
    }
}()

cfg := nsq.NewConfig()
cfg.MaxInFlight = 1000 //Maximum number of messages to allow in flight (concurrency knob)
topicName := "my_topic"
channelName := "my_chan"
cons, err := nsq.NewConsumer(topicName, channelName, cfg)
if err != nil {
    log.Fatalf(err.Error())
}

// the method below is an alternative to AddHandler to enable concurrent processing
// the second argument is the number of goroutines to spawn for processing
cons.AddConcurrentHandlers(&MyMessageHandler{}, 20)

err := cons.ConnectToNSQD(nsqdAddr)
if err != nil {
    log.Fatalf(err.Error())
}
