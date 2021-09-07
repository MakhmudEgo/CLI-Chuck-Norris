package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

// Joke – Структура для получения поле value из Resp
type Joke struct {
	Value string `json:"value"`
}

// Random – функция для получения рандомной шутки
func Random() error {
	kek := new(Joke)
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf(err.Error())
	}
	if err := json.Unmarshal(b, kek); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(kek.Value)
	return nil
}

// RandomWithCat – функция для получения рандомной шутки с учетом категории
func RandomWithCat(cat string) (string, error) {
	kek := new(Joke)
	resp, err := http.Get("https://api.chucknorris.io/jokes/random?category=" + cat)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	if err := resp.Body.Close(); err != nil {
		return "", fmt.Errorf(err.Error())
	}
	if err := json.Unmarshal(b, kek); err != nil {
		log.Fatal(err.Error())
	}
	return kek.Value, nil
}

// Dump – функция для получения n шуток из всех существующих категорий и записи в файлы(категория.txt)
func Dump(n int) error {
	cat, err := getCategories()
	if err != nil {
		return err
	}
	wg := &sync.WaitGroup{}
	for _, el := range cat {
		wg.Add(1)
		go handle(wg, el, n)
	}
	wg.Wait()
	return nil
}

// handle – вспомогательная функция Dump
/*
 если ответы [https://api.chucknorris.io/jokes/random?category={category}]
 3 раза(AmountReq) одинаковые, то функция прекращает свою работу, несмотря на флаг -n
 – для предотвращения залипания, в некоторых категориях шуток может быть меньше чем {-n}
*/
func handle(wg *sync.WaitGroup, el string, n int) {
	defer wg.Done()
	file, err := os.OpenFile(el+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}
	jokes := make([]string, n)
	var AmountReq int
	var prevJoke string
	for i := 0; i < n; i++ {
		if AmountReq == 3 {
			return
		}
		res, err := RandomWithCat(el)
		if err != nil {
			log.Fatal(err)
		}
		if prevJoke == res {
			AmountReq++
		}
		prevJoke = res
		if isNotUnique(res, jokes) {
			i--
			continue
		}
		_, err = fmt.Fprintln(file, res)
		if err != nil {
			log.Fatal(err)
		}
		jokes = append(jokes, res)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// isNotUnique – функция для проверки на уникальность
func isNotUnique(joke string, jokes []string) bool {
	for _, s := range jokes {
		if s == joke {
			return true
		}
	}
	return false
}

// getCategories – функция для получения категории
func getCategories() ([]string, error) {
	var cat []string
	resp, err := http.Get("https://api.chucknorris.io/jokes/categories")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	if err := resp.Body.Close(); err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	if err := json.Unmarshal(b, &cat); err != nil {
		log.Fatal(err.Error())
	}
	return cat, nil
}
