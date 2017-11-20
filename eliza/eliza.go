package eliza
//Code: https://getbootstrap.com/docs/4.0/components/forms/
//Take from form: https://stackoverflow.com/questions/23282311/parse-input-from-html-form-in-golang
			// https://www.google.ie/search?q=parse+input+in+go+html&rlz=1C1GGRV_enIE769IE769&oq=parse+input+in+go+html&aqs=chrome..69i57.6111j0j4&sourceid=chrome&ie=UTF-8
import (
	// "fmt"
    "math/rand"
    //"regexp"
    "strings"
	
	
)

// func checkError(e error) {
//     if e != nil {
//         panic(e)
// }

func ElizaHi(statement string) string{
    // if IsStartStatement(statement) {
        return random(Greetings)
    // }        

    // return "Hello to you too." 
        
}

func Ask(chat string) string {
	//responses := responseList()

	return ""
	
} // Ask

func IsStartStatement(statement string) bool {
    statement = trim(statement)
    for _, startStatement := range StartStatements {
        if statement == startStatement {
            return true
        }
    }
    return false
} // IsStartStatement

func trim(statement string) string {
    //statement = strings.TrimRight(statement, "\n.!")
    statement = strings.ToLower(statement)
    return statement
} // trim

func random(list []string) string {
    randomNumber := rand.Intn(len(list)) // returns a random index number related to whatever the array number from the array list
    return list[randomNumber]
}