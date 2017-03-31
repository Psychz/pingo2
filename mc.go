package main

import (
	"net"
	"errors"
	"time"
)

func MCQuery(hostname string, config Config) (reply bool, err error) {
		c := make(chan bool, 1)

		go func() {
			conn, err := net.Dial("tcp", hostname)
			if err != nil {
				c <- false
			}

			_, err = PingMC(conn, hostname)
			if err != nil {
				c <- false
			}
			c <- true
		}()
    select {
    case res := <-c:
        return res, nil
    case <-time.After(time.Duration(config.Timeout) * time.Second):
		return false, errors.New("Timeout")
    }
}
