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
		err := fmt.Errorf("erro na hora de configurar o Gerenciador de fila - %w", err)
		fmt.Println(err)
	}

	fmt.Println("O gerenciador de fila foi criado...")

	for {
		var uniName string
		fmt.Println("Digite o termo que deseja buscar as universidades:")
		fmt.Scanln(& uniName);
		
		url := os.Getenv("EXTERNAL_API")

		response, err := http.Get(url+"="+uniName)
		if err != nil {
			fmt.Println("Erro na hora da requisição")
			return
		}
	
		defer response.Body.Close()
	
		if response.StatusCode != http.StatusOK {
			fmt.Println("Ocorreu um problema, código de erro: ", response.StatusCode)
			return
		}
	
		var uniArray []domain.University;
	
		err = json.NewDecoder(response.Body).Decode(&uniArray)
		if err != nil {
			fmt.Println("Deu erro na hora de desserializar o json: ", err)
			return
		}
	
		for _, uni := range uniArray{
			message := fmt.Sprintf("Informações das universidades: %s - %s - %s", uni.Name, uni.Country, uni.Country)
			queueConnector.SendMessage(os.Getenv("QUEUE_NAME"),message)
		}
	}
}