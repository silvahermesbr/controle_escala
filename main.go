package main

import (
	f "fmt"
	"github.com/inancgumus/screen" // limpar tela
)


//CONSTANTES DE CONFIGURAÇÃO
const DASH string = "==================================="
const HOSTNAME string = "box.devhermes.com.br"
const DB_PORT string =  "3306"
const SERVER_PORT string = ":8080"
const KEY_MANAGER string = "sgteante:3scalA"
const KEY_USER string = "cadete:cadete"
const DB_NAME string = "escala_aman"
var STD_TABLES = [3]string{"militares", "registros", "tipo_sv"}
var DATA_TRABALHO = [3]int{1, 2, 2025}


//COMANDOS PRE-ESCRITOS
const CONN_STRING_FOR_MANAGER string = KEY_MANAGER + "@tcp(" + HOSTNAME + ":" + DB_PORT + ")/" + DB_NAME 
const CONN_STRING_FOR_USER string = KEY_USER + "@tcp(" + HOSTNAME + ":" + DB_PORT + ")/" + DB_NAME

//STRUCTS BASICAS 1 para cada tabela

type militar struct {
	id	int
	nome	string
	numero	int
	turma	int
	dia	int
	mes	int	
	ano	int
}

type tipo_sv struct {
	id	int
	desc	string
}

type registro struct {
	id	int
	id_mil	int
	id_sv	int
	dia	int
	mes	int
	ano	int
}


func main () {
	screen.Clear()
	f.Println(DASH)
	f.Printf("Iniciando\n")
	//testes preliminares de conexão startup.go
	startup()
	//Verificação do estado das tabelas necessarias
	checkTables()
	screen.Clear()
	f.Println(DASH)
	f.Println("Carregamento Concluido")
	var choice string = "2"
	var keeper = true
	for (keeper == true) {
		f.Println(DASH)
		f.Printf("1> Modo de servidor\n2> Modo de gerenciamento\n3> Sair\n"+ DASH +"\n#>> ")
		f.Scan(&choice)
		if (choice == "1") {
			screen.Clear()
			server_init() //server.go
			break
		} else if (choice == "2") {
			screen.Clear()
			mainMenu()
			break
		} else if (choice == "3") {
			screen.Clear()
			break
		} else {
			screen.Clear()
			f.Println("Por favor escolha um dos modos oferecidos")
		}
	}
}
