package main

import (
   "fmt"
   "io/ioutil"
   "log" // writes an standard error
   "net/http" // provides an HTTP client and server
   "os"
)

func httpGet(url string) *http.Response{
   response, err := http.Get(url)
   printError(err)
   return response
}

func processBodyResponse(response *http.Response) string{
   body, err := ioutil.ReadAll(response.Body)
   defer response.Body.Close() // the client must close the response body when finished with it
   printError(err)
   return string(body)
}

func processHeaderResponse(response *http.Response) string{
   header := response.Header
   var result string
   for x, y := range header{
      result += fmt.Sprintf("%s\t-->\t%s\n\n",x, y)
   }
   return result
}

func writeFile(fileName string, text string){
   // create file
   file, err := os.Create(fileName)
   printError(err)

   // write
   number, err := file.WriteString(text)
   if printError(err){
      file.Close()
   }

   fmt.Printf("Written successfully (%d bytes)", number)
   err = file.Close()
   printError(err)
}

func printError(err error) bool{
   if err != nil {
      log.Fatalln(err) // is equivalent to Println() followed by a call to os.Exit
      return true
   }
   return false
}

func main() {

   // HTTP GET
   const url = "https://www.epn.edu.ec/"
   response := httpGet(url)

   fmt.Printf("::::::::::::::::::::::::: URL: %s :::::::::::::::::::::::::\n", url)

   // Titles
   stars := "***************************"
   headerTitle := fmt.Sprintf("%s HEADER %s\n", stars, stars)
   bodyTitle := fmt.Sprintf("%s BODY %s\n", stars, stars)

   // Response content
   content := headerTitle + processHeaderResponse(response) +
              bodyTitle + processBodyResponse(response) // EPN's web

   // Write  report
   writeFile("report.txt", content)
}