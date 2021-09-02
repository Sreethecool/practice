package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'RemainderSorting' function below (and other code for sorting if needed).
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */


var writer
func RemainderSorting(strArr []string) []string {

  sortedList := []string{}
  var a = map[int][]string{}
    for _,v := range strArr {
        l :=len(v)%3
        if s,ok := a[l];ok{
            for i:=0;i<len(s);i++{
                if s[i] > v {
        fmt.Fprintf(writer, "%v", s)
                    fmt.Println(s)
                    last := len(s)-1
                    s = append(s,s[last])
                    copy(s[i+1:],s[i:last])
                    s[i]=v
                    fmt.Fprintf(writer, "%v", s)
                }
            }
            a[l]=s
        }else{
            a[l] = []string{v}
        }
    }
    for i:=0;i<3;i++{
        
        sortedList = append(sortedList,a[i]...)    
    }
    return sortedList
}

func main() {

   strArr := []string{"a","ab","bc","abc"}


    result := RemainderSorting(strArr)
	fmt.Println(result)
}

