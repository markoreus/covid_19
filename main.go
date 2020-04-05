package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

func main() {
	countries := JSON{Data: Data{}}
	fmt.Println("Fetching Data....")
	response, err := getResponse()
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done")
	err = json.Unmarshal(data, &countries)
	if err != nil {
		log.Fatalln(err)
	}

	sort.Sort(countries.Data)
	err = writeCsv(countries.Data)
	if err != nil {
		log.Fatalln(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Print("Press enter to exit")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		wg.Done()
	}()
	wg.Wait()
}
