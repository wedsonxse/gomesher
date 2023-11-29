package main

import (
	"encoding/json"
	"fmt"
	"gomesher/domain"
	"gomesher/infra"
	"net/http"
	"os"
)

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

	fmt.Println("Queue manager was created...")

	for {
		var uniName string
		fmt.Println("Type the term to search in the university API:")
		fmt.Scanln(& uniName);
		
		url := os.Getenv("EXTERNAL_API")

		response, err := http.Get(url+"="+uniName)
		if err != nil {
			fmt.Println("Request Error")
			return
		}
	
		defer response.Body.Close()
	
		if response.StatusCode != http.StatusOK {
			fmt.Println("A problem ocurred, error code: ", response.StatusCode)
			return
		}
	
		var uniArray []domain.University;
	
		err = json.NewDecoder(response.Body).Decode(&uniArray)
		if err != nil {
			fmt.Println("Unmarshalling json error: ", err)
			return
		}
	
		for _, uni := range uniArray{
			message := fmt.Sprintf("Universities Informations: %s - %s - %s", uni.Name, uni.Country, uni.Country)
			queueConnector.SendMessage(os.Getenv("QUEUE_NAME"),message)
		}
	}
}