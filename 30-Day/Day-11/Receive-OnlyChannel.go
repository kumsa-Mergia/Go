func receive(ch <-chan int) {
    fmt.Println(<-ch) // Can only receive from this channel
}
