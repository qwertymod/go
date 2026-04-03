package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	//file, err := os.Open("files.txt") // чтение по байтам

	data, err := os.ReadFile(name) // - вместо файла получаем всё содержимое в виде массива байтов
	if err != nil {
		fmt.Println(err)
		return nil , err
	}

	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer file.Close() // при добавлении новых defer они будут выполняться last in - first out
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Println("Запись успешна")
}