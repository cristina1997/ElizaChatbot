package eliza
//Code: https://getbootstrap.com/docs/4.0/components/forms/
//Take from form: https://stackoverflow.com/questions/23282311/parse-input-from-html-form-in-golang
			// https://www.google.ie/search?q=parse+input+in+go+html&rlz=1C1GGRV_enIE769IE769&oq=parse+input+in+go+html&aqs=chrome..69i57.6111j0j4&sourceid=chrome&ie=UTF-8
import (
    "fmt"
    "math/rand"
    "regexp"
    "strings"
    "time"
	
)

type Response struct {
	re      *regexp.Regexp
	answers []string
}

// func checkError(e error) {
//     if e != nil {
//         panic(e)
// }

// Returns a random greeting for Eliza
func ElizaHi() string{
    return random(Greetings)        
}

// Returns a random goodbye for Eliza
func ElizaBye() string{
    return random(Valediction)        
}

func Ask(statement string) string {
    // Trim the statement for more effective matching
    // This method returns a copy of the string, with whitespace omitted.
    statement = trim(statement)

    // Check if this is a end statement
    if IsEndStatement(statement){
        return ElizaBye()
    }

    // We match the statement to a statement recognizable by Eliza 
    // which can then construct an appropriate response.
    for i, responses := range Responses {
        re := regexp.MustCompile("(?i)" + i)
        matches := re.FindStringSubmatch(statement)
        
        if matches != nil {
            var match string
            if len(matches) > 1 {
                match = split(matches[1])
            }
            response := random(responses)

            if strings.Contains(response, "%s") {
                response = fmt.Sprintf(response, match)
            }

            return response
        }
    } // for

    return random(Default)
	
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

func IsEndStatement(statement string) bool {
    statement = trim(statement)
    for _, endStatement := range EndStatements {
        if statement == endStatement {
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
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(len(list)) // returns a random index number related to whatever the array number from the array list
    return list[randomNumber]
}

func split(match string) string{
    words := strings.Split(match, " ")
    for i, word := range words {
        if reflectedWord, ok := Reflected[word]; ok {
            words[i] = reflectedWord
        }
    }
    return strings.Join(words, " ")
}
