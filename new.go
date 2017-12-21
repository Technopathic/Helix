
package main

import (
  "log"
  "os"
  "time"
  "bufio"
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
)

var (
  newSummary *os.File
  newContent *os.File
  err error
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  countfile, err := ioutil.ReadFile("./posts/count.md")
 
  lines := strings.Split(string(countfile), "\n")
  count := string(lines[0])
  newCount, err := strconv.Atoi(count)
  newCount = newCount + 1
  err = ioutil.WriteFile("./posts/count.md", []byte(strconv.Itoa(newCount)),0644)  
 
  if err != nil {
    fmt.Println(err)
    return
  }

  var postName string
  fmt.Println("Name of your Post: ")
  postName, _ = reader.ReadString('\n')
  postName = strings.Replace(postName, "\n", "", -1)
  postName = strings.Replace(postName, " ", "-", -1)
  current_time := time.Now().Local()  
  myTime := current_time.Format("January 2, 2006")

  newSummary, err = os.Create("posts/summary/" + string(count) + "_" + postName + ".md")
  newSummary.WriteString("New Post\n" + myTime + "\n http://via.placeholder.com/300\n A summary goes here\n")
  newContent, err = os.Create("posts/content/" + string(count) + "_" + postName + ".md")
  if err != nil {
    log.Fatal(err)
  }
  
  //countfile.Close()
  newSummary.Close()
  newContent.Close()
}

