package linespliter

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func generateString(len int) string {
	r := make([]rune, len)
	for i := range r {
		r[i] = ' '
	}
	return string(r)
}

func copyAt(str, value string, index int) string {
	return strings.Join([]string{string([]rune(str)[:index]), value, string([]rune(str)[index+len([]rune(value)):])}, "")
}

func Marshal(e interface{}) (string, error) {
	el := reflect.TypeOf(e).Elem()
	maxLen := 0
	// je recuper la taille maximum du champs
	for i := 0; i < el.NumField(); i++ {
		if v := el.Field(i).Tag.Get("line"); len(v) > 0 {
			vS := strings.Split(v, ":")
			end, err := strconv.Atoi(vS[1])
			if err != nil {
				return "", err
			}
			if end > maxLen {
				maxLen = end
			}
		}
	}
	str := generateString(maxLen)
	elv := reflect.ValueOf(e).Elem()
	for i := 0; i < el.NumField(); i++ {
		if v := el.Field(i).Tag.Get("line"); len(v) > 0 {
			vS := strings.Split(v, ":")
			start, err := strconv.Atoi(vS[0])
			if err != nil {
				return "", err
			}
			tmp := elv.Field(i).Interface()
			str = copyAt(str, reflect.ValueOf(tmp).String(), start)
		}
	}
	return str, nil
}

func checkSimilar(line, e interface{}) error {
	v, err := Marshal(e)
	if err != nil {
		return err
	}
	// fmt.Printf("'%s'\n'%s'\n", v, line)
	if v == line {
		return nil
	}
	return errors.New("Not same result")
}

func Unmarshal(line string, e interface{}) error {
	el := reflect.TypeOf(e).Elem()
	for i := 0; i < el.NumField(); i++ {
		if v := el.Field(i).Tag.Get("line"); len(v) > 0 {
			vS := strings.Split(v, ":")
			start, err := strconv.Atoi(vS[0])
			if err != nil {
				return err
			}
			end, err := strconv.Atoi(vS[1])
			if err != nil {
				return err
			}
			if len(line) < end-1 {
				return errors.New("Line is so small")
			}
			reflect.ValueOf(e).Elem().Field(i).SetString(string([]rune(line)[start:end]))
		}
	}
	return checkSimilar(line, e)
}
