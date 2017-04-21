package storage

import (
	"github.com/aerospike/aerospike-client-go"
	"reflect"
	"errors"
)

type AeroSpike struct {
	Host string
	Port int
	Ns string
	Set string
}

var (
	client *aerospike.Client
	isConnect int
)


// получение записи о посетителе
func (as *AeroSpike) Get(hash string) (map[string]interface{}, error) {

	policy := new(aerospike.BasePolicy)
	policy.Priority = aerospike.HIGH

	if isConnect != 1 {

		conn, err := aerospike.NewClient(as.Host, as.Port)

		if err != nil {
			return nil, err
		}
		client = conn
		isConnect = 1
	}

	key, err := aerospike.NewKey(as.Ns, as.Set, hash)
	if err != nil {
		return nil, err
	}

	record, err := client.Get(policy, key)
	if record == nil {
		return nil, nil
	}

	return record.Bins, nil
}

// сохранение данных о посетителе
func (as *AeroSpike) Put(record map[string]interface{}) (error) {

	// проверяем что данные корректны
	if record == nil {
		return errors.New("Данные о посетителе пустые")
	}

	// формирукем новый полиси для записи
	policy := new(aerospike.WritePolicy)
	policy.Priority = aerospike.HIGH
	policy.Expiration = 8600

	if isConnect != 1 {
		conn, err := aerospike.NewClient(as.Host, as.Port)

		if err != nil {
			return err
		}
		client = conn
		isConnect = 1
	}

	// создаем ключ
	key, err := aerospike.NewKey(as.Ns, as.Set, reflect.ValueOf(record["id"]).String())
	if err != nil {
		return err
	}

	// записываем
	err = client.Put(policy, key, record)
	if err != nil {
		return err
	}

	return nil
}