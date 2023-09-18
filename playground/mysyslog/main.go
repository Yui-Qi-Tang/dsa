package main

import (
	"errors"
	"io"
	"log"
	"log/syslog"
	"net"
	"time"

	"github.com/nxadm/tail"
)

func main() {

	syslogClient, err := syslog.New(syslog.LOG_INFO, "MySysLog")
	if err != nil {
		panic(err)
	}
	defer syslogClient.Close()

	go func() {
		for {
			if err := syslogClient.Info("omg"); err != nil {
				panic(err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	syslogPath := "/var/log/syslog"
	t, err := tail.TailFile(syslogPath, tail.Config{
		Follow:   true,
		Location: &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd},
	})

	if err != nil {
		log.Fatalf("Error opening syslog file: %v", err)
	}

	defer t.Stop()

	// Process new syslog messages as they arrive
	for line := range t.Lines {
		// Print the received syslog message
		log.Println("Received syslog message:", line.Text)
		// You can add your custom processing logic here
	}

	select {}

}

func unixSyslog() (net.Conn, error) {
	logTypes := []string{"unixgram", "unix"}
	logPaths := []string{"/dev/log", "/var/run/syslog", "/var/run/log"}
	for _, network := range logTypes {
		for _, path := range logPaths {
			conn, err := net.Dial(network, path)
			if err == nil {
				return conn, nil
			}
		}
	}
	return nil, errors.New("unix syslog delivery error")
}
