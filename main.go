package main

import (
	"encoding/json"
	"fmt"
	"gomesher/domain"
	"gomesher/infra"
	"net/http"
	"os"
	"time"

	"github.com/go-stomp/stomp"
)

func produceMessages (connector *infra.QueueConnector) error {
	for {
		var uniName string
		fmt.Println("Type the term to search in the university API:")
		fmt.Scanln(& uniName);
		
		url := os.Getenv("EXTERNAL_API")

		response, err := http.Get(url+"="+uniName)
		if err != nil {
			fmt.Println("Request Error")
			return err
		}
	
		defer response.Body.Close()
	
		if response.StatusCode != http.StatusOK {
			fmt.Println("A problem ocurred, error code: ", response.StatusCode)
			return fmt.Errorf("status code error")
		}
	
		var uniArray []domain.University;
	
		err = json.NewDecoder(response.Body).Decode(&uniArray)
		if err != nil {
			fmt.Println("Unmarshalling json error: ", err)
			return err
		}
	
		for _, uni := range uniArray{
			message := fmt.Sprintf("Universities Informations: %s - %s - %s", uni.Name, uni.Country, uni.Country)
			connector.SendMessage(os.Getenv("QUEUE_NAME"),message)
		}
		time.Sleep(1 * time.Second)
	}
}

func consumeMessages (subChannel *stomp.Subscription, q *infra.QueueConnector) {
		for {
			select {
			case message := <-subChannel.C:	
				if message != nil {
					fmt.Println("Received Message:", string(message.Body))
					err := q.Conn.Ack(message)
					if err != nil {
						fmt.Println("Message confirmation error:", err)
					}				
				}
	
				time.Sleep(1 * time.Second)
			}
		}
}

func main(){
	infra.LoadEnv();

	queueConnector,err := infra.CreateQueueConnector(
		os.Getenv("SERVER_USERNAME"),
		os.Getenv("SERVER_PWD"),
		os.Getenv("SERVER_HOST"),
		);
		
	if err != nil {
		err := fmt.Errorf("file manager config - %w", err)
		fmt.Println(err)
	}

	subChannel, err := queueConnector.SubscribeToQueue("QUEUE_NAME")
	if err != nil {
		err := fmt.Errorf("sub error- %w", err)
		fmt.Println(err)
	}

	fmt.Println("Queue manager was created...")

	go consumeMessages(subChannel, queueConnector)

	go produceMessages(queueConnector)

	select {}
}