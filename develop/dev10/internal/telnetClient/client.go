package telnetClient

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type IClient interface {
	Connect(host, port string, timeout time.Duration, ctx context.Context, out context.CancelFunc) error
	writer(conn net.Conn, ctx context.Context)
	reader(conn net.Conn, ctx context.Context)
}

func NewClient() IClient {
	return &Client{}
}

type Client struct{}

func (c *Client) Connect(host, port string, timeout time.Duration, ctx context.Context, out context.CancelFunc) error {
	defer out()
	log.Printf("Connect to %s:%s \n", host, port)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Printf("Connected to %s:%s\n", host, port)

	go c.reader(conn, ctx)
	go c.writer(conn, ctx)

	<-ctx.Done()

	return nil

}

func (c *Client) writer(conn net.Conn, ctx context.Context) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			log.Println("Signal to stop was received, stop writing...")
			return
		default:
			var data string
			if !sc.Scan() {
				log.Printf("Failed to scan input, stop writing: %s\n", sc.Err())
				return
			}
			data = sc.Text()
			n, err := conn.Write([]byte(data))
			if err != nil {
				log.Println("Failed to write input, stop writing...")
				return
			}
			log.Printf("%d bytes were sent to %s\n", n, conn.RemoteAddr())
		}
	}

}

func (c *Client) reader(conn net.Conn, ctx context.Context) {
	reader := bufio.NewReader(conn)

	for {
		select {
		case <-ctx.Done():
			log.Println("Signal to stop was received, stop reading...")
			return
		default:
			data := make([]byte, 1024)
			n, err := reader.Read(data)
			if err != nil && err != io.EOF {
				log.Printf("Failed to read input, stop reading: %s\n", err)
				return
			}
			if n > 0 {
				fmt.Printf("Reading input from %s: -> %s\n", conn.RemoteAddr(), string(data[:n]))
			}
		}
	}
}
