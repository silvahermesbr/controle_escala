package main

import (
	f "fmt"
	"github.com/inancgumus/screen" // limpar tela
	"time"
)

func menuPessoal(turma string) {
	var choice string
	if (turma == "0") {
		screen.Clear()
		f.Printf("Por favor digite o ano da turma (Aspirantado): ")
		f.Scan(&turma)
	}
	screen.Clear()
	f.Println("Visão geral")
	f.Println(DASH)
	f.Printf("Carregando...\r")
	resultado := exec_sql_return_all_militares_turma(turma)
	if (resultado != "") {
		f.Println(resultado)
	} else {
		f.Println("Não há registros para a turma de: " + turma)
	}
	f.Println(DASH)
	f.Printf("1> Adicionar militar\n2> Alterar info\n3> Deletar Militar\n4> Voltar para o Menu Principal\n#>> ")
	f.Scan(&choice)
	if (choice == "1") {
		screen.Clear()
		var nome, numero, dia, mes, ano string
		f.Println("ADICIONANDO NOVO REGISTRO")
		f.Println(DASH)
		f.Printf("NOME DE GUERRA: ")
		f.Scan(&nome)
		f.Printf("\nNUMERO: ")
		f.Scan(&numero)
		f.Printf("\nDATA DE NASCIMENTO\n")
		f.Printf("\nDIA: ")
		f.Scan(&dia)
		f.Printf("\nMES: ")
		f.Scan(&mes)
		f.Printf("\nANO: ")
		f.Scan(&ano)
		screen.Clear()
		var final_choice string = ""
		f.Printf("\n" + DASH + "\nO registro seguinte está prestes a ser adicionado:\n NOME: %s\n NUMERO: %s\n TURMA: %s\n NASCIDO EM: %s/%s/%s\nIsso está correto?\n", nome, numero, turma, dia, mes, ano)
		for {
			f.Println("#(s/n)>> ")
			f.Scan(&final_choice)
			if (final_choice == "s") {
				screen.Clear()
				militaresNew(nome, numero, turma, dia, mes, ano)
				time.Sleep(2 * time.Second)
				break
			} else if (final_choice == "n") {
				screen.Clear()
				f.Println("Registro Descartado")
				time.Sleep(2 * time.Second)
				break
			} else {
				screen.Clear()
				f.Println("Favor digitar 's' ou 'n'")
			}
		}
		menuPessoal(turma)
	} else if (choice == "2") {
		f.Printf("Por Favor digite o ID do registro que você deseja alterar: ")
		var id string
		f.Scan(&id)
		//NÃO FINALIZAR (PAREI EM MILITARES.GO TAMBEM!)
	} else if (choice == "3") {
		f.Printf("Por Favor digite o ID do registro que você deseja remover da tabela: ")
		var id string
		f.Scan(&id)
		screen.Clear()
		militaresDelete(id)
		time.Sleep(2 * time.Second)
		menuPessoal(turma)
	} else if (choice == "4") {
		mainMenu()
	} else {
	}
}

func menuSV() {
	screen.Clear()
}

func mainMenu() {
	screen.Clear()
	//Primeiro vamos selecionar qual das 3 tabelas vamos usar
	var choice string = ""
	var keeper = true
	for (keeper == true) {
		f.Print(DASH)
		f.Printf("\nMENU PRINCIPAL DE GERENCIAMENTO\n1> Pessoal\n2> SV\n3> Sair\n" + DASH + "\n#>> ")
		f.Scan(&choice)
		if (choice == "1") {
			menuPessoal("0")
			break
		} else if (choice == "2") {
			menuSV()
			break
		} else if (choice == "3") {
			screen.Clear()
			break
		} else {
			screen.Clear()
			f.Println("Por favor selecione umas das opções válidas")
		}
	}
}
