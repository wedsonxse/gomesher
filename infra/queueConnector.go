package infra

import (
	"fmt"
	"net"

	"github.com/go-stomp/stomp"
)

type QueueConnector struct {
	conn *stomp.Conn
}

func CreateQueueConnector(user,password,host string) (*QueueConnector,error){
	conn, err := net.Dial("tcp",host)
	if err != nil{
		err := fmt.Errorf("erro no dial - %w",err)
		return nil, err
	}

	stompConn, err := stomp.Connect(
		conn,
		stomp.ConnOpt.Login(user,password),
		stomp.ConnOpt.Host(host),
	)

	if err != nil {
		err := fmt.Errorf("erro na conexão - %w",err)
		return nil,err
	}

	return &QueueConnector{ conn: stompConn }, nil
}

func (q *QueueConnector) SendMessage(queueName , message string) error {
	err := q.conn.Send(queueName, "text/plain", []byte(message), stomp.SendOpt.Receipt)
	if err != nil{
		return err
	}

	fmt.Println("Mensagem enviada com sucesso para a fila!")
	return nil
}


func (q *QueueConnector) CloseConnection () {
	q.conn.Disconnect()
}