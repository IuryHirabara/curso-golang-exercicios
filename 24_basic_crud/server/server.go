package server

import (
	"basic_crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler corpo da requisição!"))
		return
	}

	var user user

	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Falha ao tentar converter corpo da requisição em json para struct!"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao tentar conectar no banco!"))
		return
	}

	// prepare statement
	// para evitar sql injection
	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?);")
	if err != nil {
		w.Write([]byte("Não foi possível preparar o statement"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao tentar fazer insert!"))
		return
	}

	idUser, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Não foi possível recuperar o id do insert!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idUser)))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("Houve um erro ao executar o parse na requisição"))
		return
	}

	id, err := strconv.ParseUint(r.Form.Get("id"), 10, 64)
	if err == nil {
		GetUserById(w, r, id)
	} else {
		GetAllUsers(w, r)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request, id uint64) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Não foi possível conectar ao banco de dados"))
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		w.Write([]byte("Não foi possível realizar a consulta"))
		return
	}
	defer row.Close()

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear"))
			return
		}
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para json"))
		return
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Não foi possível conectar ao banco de dados"))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		w.Write([]byte("Não foi possível consultar os usuários"))
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Falha ao escanear linhas"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Não foi possível converter a resposta para json"))
		return
	}
}
