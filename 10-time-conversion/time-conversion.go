package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

func timeConversion(s string) string {
    inputFormat := "03:04:05PM"
    outputFormat := "15:04:05"
    t, err := time.Parse(inputFormat, s)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    return t.Format(outputFormat)

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
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

