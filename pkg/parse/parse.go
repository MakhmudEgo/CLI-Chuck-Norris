package parse

import (
	"fmt"
	"strconv"
)

const (
	Random = iota + 1
	Dump
)

// Data – структура для комманд
type Data struct {
	Type  int
	Count int
}

// Scan – Проверяет на валидность аргументов программы и возвращает структуру для выполнения комманды
func Scan(args []string) (*Data, error) {
	res := new(Data)
	if len(args) == 0 {
		return res, fmt.Errorf("empty arguments")
	}
	switch args[0] {
	case "random":
		{
			if len(args) != 1 {
				return res, fmt.Errorf("bad arguments")
			} else {
				res.Type = Random
				return res, nil
			}
		}
	case "dump":
		{
			res.Type = Dump
			if len(args) < 2 || len(args) > 3 || args[1] != "-n" {
				return res, fmt.Errorf("bad arguments")
			} else if len(args) == 3 {
				var err error
				res.Count, err = strconv.Atoi(args[2])
				if err != nil {
					return res, fmt.Errorf("bad arguments")
				}
			} else {
				res.Count = 5
			}
			return res, nil
		}
	}
	return res, fmt.Errorf("unknown arguments")
}
