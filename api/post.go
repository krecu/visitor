package api

import (
	"net/http"
	"encoding/json"
	"visitor/core"
	"fmt"
)

type Method struct {}

// тело запроса
type Body struct{
	Ip string
	Ua string
	Id string
	Extra map[string]interface{}
	Debug int
}

// ошибка
type Error struct {
	code int
	error string
}

// создание записи о посетителе GET & POST
func (api *Method) Post(rw http.ResponseWriter, req *http.Request) {

	// берем контент из тела запроса и создаем структуру
	decoder := json.NewDecoder(req.Body)
	var body Body
	err := decoder.Decode(&body)

	// если неудалось распарсить содержимое body запроса
	if err != nil {
		api.Error(rw, 100, "Неверный формат запроса: " + err.Error())
		return
	}

	// получаем данные о посетителе
	coreVisitor := core.Visitor{Ua: body.Ua, Ip: body.Ip, Id: body.Id, Extra:body.Extra}
	visitor, err := coreVisitor.Identify()

	// если при определении информации о посетителе возникла ошибка
	if err != nil {
		api.Error(rw, 101, "Неудалось получить данные о посетителе: " + err.Error())
		return
	}

	// упаковываем структуру в json
	jsonCode, err := json.Marshal(visitor)
	if err != nil {
		api.Error(rw, 102, "Неудалось преобразовать в JSON: " + err.Error())
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(jsonCode)

}

// @todo - как бы даделать общий оброботчик ошибок
func (api *Method) Error(rw http.ResponseWriter, code int, error string) {

	fmt.Println(error)

	body := new(Error)
	body.error = error
	body.code = code

	jsonCode, _ := json.Marshal(body)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(500)
	rw.Write(jsonCode)
}