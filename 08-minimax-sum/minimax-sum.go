package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

type intsarr []int32

func (i intsarr) Len() int {
    return len(i)
}

func (i intsarr) Swap(k, j int) {
    i[k], i[j] = i[j], i[k]
}

func (i intsarr) Less(k, j int) bool {
    return i[k] < i[j]
}

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
    var minSum, maxSum int64
    
    sort.Sort(intsarr(arr))

    for i := 0; i < 4; i++ {
        minSum += int64(arr[i])
    }

    for k := 0; k < 4; k++ {
        maxSum += int64(arr[len(arr)-1-k])
    }

    fmt.Printf("%d %d", minSum, maxSum)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

