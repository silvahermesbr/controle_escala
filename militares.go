package main

import (
	f "fmt"
)

/*
Cheasheet:
militaresNew(nome string, numero int, turma int, dia_nascimento int, mes_nascimento int, ano_nascimento int) returns nothing
militaresEdit(id int, coluna string, valor_novo string)
militaresDelete(id int)
*/


//Cria um novo registro de militar na tabela do DB
func militaresNew(nome string, numero string, turma string, dia_nascimento string, mes_nascimento string, ano_nascimento string) {
	f.Printf("Iniciando registro no banco de dados para [%s]...", nome) 
	var comando string = f.Sprintf("INSERT INTO militares (nome, numero, turma, dia_nascimento, mes_nascimento, ano_nascimento) VALUES ('%s', %s, %s, %s, %s, %s);", nome, numero, turma, dia_nascimento, mes_nascimento, ano_nascimento) 
	result := exec_query(comando);
	if result != true {
		error_sql("REGISTRO MILITAR: " + nome)
	} else {
		f.Printf("\nMilitar registrado com sucesso!")
	}
}

//Edita alguma coisa do militar baseado em sua ID
func militaresEdit(id string, coluna string, valor_novo string) {
	var comando string = f.Sprintf("UPDATE militares SET %s='%s' WHERE id=%s;", coluna, valor_novo, id)
	result := exec_query(comando);
	if result != true {
		error_sql("EDIÇÃO DO REGISTRO DE MILITAR")
	} else {
		f.Printf("Militar editado com sucesso!")
	}
}

//Deleta um militar baseado em sua ID
func militaresDelete(id string) bool {
	f.Printf("Iniciando processo do registro de ID [%s]\n", id)
	var comando string = f.Sprintf("DELETE FROM militares WHERE id=%s", id)
	result := exec_query(comando);
	if result != true {
		error_sql("REMOÇÃO DO REGISTRO DE MILITAR")
	} else {
		f.Printf("\nMilitar deletado com sucesso!")
	}
	return result
}

//Retorna um objeto com todas as info necessarias. //NÃO FINALIZADO:
func query_individual(id string) militar {
	var infos militar
	return infos
}
