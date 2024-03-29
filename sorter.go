package main

import "flag"
import "fmt"
import "bufio"
import "io"
import "os"
import "strconv"
import "time"

import "GoStudy/algorithm/bubblesort"
import "GoStudy/algorithm/qsort"

var infile *string = flag.String("i", "infile", "File contains value for sorted")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm  qsort|bubble")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Faield to open the input file", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}

		str := string(line) //转换字符输出为字符串

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println(" infile = ", *infile, " outfile = ", *outfile, " algotith = ", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "bubble":
			bubblesort.Bubblesort(values)
		case "qsort":
			qsort.QuickSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()

		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		fmt.Println("Read Values", values)

		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}

}
