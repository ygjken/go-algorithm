package readdata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func LoadData() ([]int, int) {
	var tmp string
	var n int
	var data []int

	// read the data files
	f, err := os.Open("test/data.txt")
	if err != nil {
		fmt.Println("Input Error")
	}
	defer f.Close()

	// start scanning
	scanner := bufio.NewScanner(f)
	// scan the beginning of the file
	scanner.Scan()
	tmp = scanner.Text()
	n, _ = strconv.Atoi(tmp)
	// scan every lines
	data = make([]int, n)
	for i := 0; scanner.Scan(); i++ {
		tmp = scanner.Text()
		data[i], _ = strconv.Atoi(tmp)
	}
	// error hadling
	if err := scanner.Err(); err != nil {
		fmt.Println("Read Lines Error")
	}

	return data, n
}
