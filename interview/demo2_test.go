package interview

import "time"

func main() {
	for {
		err := KeepAlive()
		if err != nil {
			continue
		}
		break
	}

	select {}

}

func KeepAlive() error {

	// 建立连接
	connect, err := NewConnect()

	if err != nil {
		return err
	}

	defer connect.Close()

	for {

		err = connect.Send("心跳消息")
		if err != nil {
			for {
				err = KeepAlive()
				if err != nil {
					continue
				}
				return nil
			}
		}

		<-time.After(time.Second)
	}

	return nil

}

type Connect struct {
}

func (c *Connect) Close() {
	println("Close ")
}
func (c *Connect) Send(msg string) error {
	println("send " + msg)
	return nil
}
func NewConnect() (*Connect, error) {
	return &Connect{}, nil
}
