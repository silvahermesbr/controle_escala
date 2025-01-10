package main

import (
	f "fmt"
	sql "database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

//Função basica para retornar "verbose errors"
func error_sql(local string) {
	f.Printf("\n[" + local + "] ERRO NA EXECUÇÃO\n\n")
}


//Essa função usa a struct militares definida em main para retornar uma string formatada em tabela de todos os registros de militares no DB 
func exec_sql_return_all_militares() string {	
	var comando string = f.Sprintf("SELECT * FROM militares ORDER BY numero DESC");
	db, err:= sql.Open("mysql", CONN_STRING_FOR_USER)
	defer db.Close()
	if err != nil {
		error_sql("RETURN ALL CONN")
		panic(err)
		return "erro"
	}
	var result string = ""
	var m militar
	rows, err := db.Query(comando)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&m.id, &m.nome, &m.numero, &m.turma, &m.dia, &m.mes, &m.ano)
		if err != nil {
			panic(err)
			return "erro"
		}
		unidade := f.Sprintf("[%d - %d] Nome: %s\t\tNumero: %d\tNascido: %d/%d/%d ", m.id, m.turma, m.nome, m.numero, m.dia, m.mes, m.ano)
		result += unidade
		result += "\n"
	}
	return result
}

func exec_sql_return_all_militares_turma(turma string) string {	
	var comando string = f.Sprintf("SELECT * FROM militares WHERE turma=%s ORDER BY numero DESC", turma);
	db, err:= sql.Open("mysql", CONN_STRING_FOR_USER)
	defer db.Close()
	if err != nil {
		error_sql("RETURN ALL CONN")
		panic(err)
		return "erro"
	}
	var result string = ""
	var m militar
	rows, err := db.Query(comando)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&m.id, &m.nome, &m.numero, &m.turma, &m.dia, &m.mes, &m.ano)
		if err != nil {
			panic(err)
			return "[[A turma nao existe]]"
		}
		unidade := f.Sprintf("[ID: %d] Nome: %s\t\tNumero: %d\tNascido: %d/%d/%d ", m.id, m.nome, m.numero, m.dia, m.mes, m.ano)
		result += unidade
		result += "\n"
	}
	return result
}


//Função basica para simplificar a execução de comandos SQL simples.
func exec_query(comando string) bool {
	db, err := sql.Open("mysql", CONN_STRING_FOR_MANAGER)
	defer db.Close()
	if err != nil {
		error_sql("RUN QUERY CONN")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		error_sql("RUN QUERY PING")
		panic(err)
	} else {
		f.Printf("[Conexão bem sucedida] ")
	}
	var result bool = false
	_, err = db.Query(comando)
	if err != nil {
		error_sql("RUN QUERY EXEC")
		panic(err)
	} else {
		result = true
	}
	return result
}
