package main

import (
	f "fmt"
)

/*
Cheasheet:
tipos_svNew(desc string)
tipos_svEdit(id int, valor_novo string) #vai editar a descricao somente
tipos_svDelete(id int)
*/


//Cria um novo registro de tipo_sv na tabela do DB
func tipos_svNew(desc string) {
	var comando string = f.Sprintf("INSERT INTO tipos_sv (descricao) VALUES ('%s');", desc) 
	result := exec_query(comando);
	if result != true {
		error_sql("REGISTRO TIPO_SV")
	} else {
		f.Printf("Tipo_sv registrado com sucesso!\n")
	}
}

//Edita alguma coisa do tipo_sv baseado em sua ID
func tipos_svEdit(id string, valor_novo string) {
	var comando string = f.Sprintf("UPDATE tipos_sv SET descricao='%s' WHERE id=%s;", valor_novo, id)
	result := exec_query(comando);
	if result != true {
		error_sql("EDIÇÃO TIPO SV")
	} else {
		f.Printf("Tipo_sv editado com sucesso!\n")
	}
}

//Deleta um tipo_sv baseado em sua ID
func tipos_svDelete(id string) {
	var comando string = f.Sprintf("DELETE FROM TABLE tipos_sv WHERE id=%s", id)
	result := exec_query(comando);
	if result != true {
		error_sql("REMOÇÃO DO TIPO SV")
	} else {
		f.Printf("Tipo_sv deletado com sucesso!\n")
	}
}
