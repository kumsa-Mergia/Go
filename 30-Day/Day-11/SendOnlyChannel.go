func send(ch chan<- int, value int) {
    ch <- value // Can only send to this channel
}
