package main

import (
	f "fmt"
)

/*
Cheasheet:
registrosNew(desc string)
registrosEdit(id int, valor_novo string) #vai editar a descricao somente
registrosDelete(id int)
*/


//Cria um novo registro de registro na tabela do DB
func registrosNew(id_militar string, id_sv string, dia string, mes string, ano string) {
	var comando string = f.Sprintf("INSERT INTO registros (id_militar_db, ID_sv, dia, mes, ano) VALUES (%s, %s, %s, %s);", id_militar, id_sv, dia, mes, ano) 
	result := exec_query(comando);
	if result != true {
		error_sql("REGISTRO SV")
	} else {
		f.Printf("Sv registrado com sucesso!\n")
	}
}

//Edita alguma coisa do registro baseado em sua ID
func registrosEdit(id string, coluna string, valor_novo string) {
	var comando string = f.Sprintf("UPDATE registros SET %s='%s' WHERE id=%s;", coluna, valor_novo, id)
	result := exec_query(comando);
	if result != true {
		error_sql("EDIÇÃO SV")
	} else {
		f.Printf("Sv editado com sucesso!\n")
	}
}

//Deleta um registro baseado em sua ID
func registrosDelete(id string) {
	var comando string = f.Sprintf("DELETE FROM TABLE registros WHERE id=%s", id)
	result := exec_query(comando);
	if result != true {
		error_sql("REMOÇÃO DO SV")
	} else {
		f.Printf("Sv deletado com sucesso!\n")
	}
}


