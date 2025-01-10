package main

import (
	f "fmt"
	"net/http"
	"io"
	//si "strconv"
)

func teste (w http.ResponseWriter, r *http.Request) {
	//padrão de teste, vou mandar um html com js para login nas outras
	//versões
	f.Println("Recebi um pedido")
	io.WriteString(w, exec_sql_return_all_militares())
}

func relatorio_militares (w http.ResponseWriter, r *http.Request) {
	var id string = r.PathValue("numero")
	//Aqui vem o codigo para que cada um possa acessar suas informações
	if id != "" {
		io.WriteString(w, id)
	} else {
		io.WriteString(w, exec_sql_return_all_militares())
	}
}

func server_init() {
	http.HandleFunc("/", teste) //root
	//visualizar todos os militares:
	http.HandleFunc("/militares/", relatorio_militares)
	err := http.ListenAndServe(SERVER_PORT, nil)
	if err != nil {
		panic(err)
	} else {
		f.Printf("Servidor Pronto e escutando na porta: %d\n", SERVER_PORT)
	}
}
