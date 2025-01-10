package main

import (
	f "fmt"
	sql "database/sql"

	_ "github.com/go-sql-driver/mysql"
)


func startup() {
	//Conexão basica
	db, err :=  sql.Open("mysql", CONN_STRING_FOR_MANAGER)
	defer db.Close()
	f.Printf("Conexão inicial ao DB em " + HOSTNAME + ": ")
	if err != nil {
		f.Printf("ERRO");
		panic(err)
	} else {
		f.Printf("OK! [PING:  ") 
	}
	//Teste de acesso
	err = db.Ping() //Pong
	if err != nil {
		f.Printf("ERRO]")
		panic(err)
	} else {
		f.Printf("OK]\n")
	}
}

func checkTables() {
	const check_tabela_militares string = "CREATE TABLE IF NOT EXISTS militares (id INTEGER PRIMARY KEY AUTO_INCREMENT, nome VARCHAR(32), numero INTEGER, turma INTEGER, dia_nascimento SMALLINT, mes_nascimento SMALLINT, ano_nascimento INTEGER);"
	const check_tabela_registros string = "CREATE TABLE IF NOT EXISTS registros (id INTEGER PRIMARY KEY AUTO_INCREMENT, id_militar_db INTEGER, ID_sv INTEGER, dia INTEGER, mes INTEGER, ano INTEGER);"
	const check_tabela_tipos_sv string = "CREATE TABLE IF NOT EXISTS tipos_sv (id INTEGER PRIMARY KEY AUTO_INCREMENT, descricao TEXT);"
	var checks = [3]string{check_tabela_militares, check_tabela_registros, check_tabela_tipos_sv}
	for index, comando := range checks {
		f.Printf("Verificando " + STD_TABLES[index] + "\t")  
		result := exec_query(comando)
		if result != false {
			f.Printf("OK!\n");
		} else {
			f.Printf("ERRO!\n");
		}
	}
}
