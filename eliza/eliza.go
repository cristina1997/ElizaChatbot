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
    statement = trim(statement)

    // Check if this is a end statement
    if IsEndStatement(statement){
        return ElizaBye()
    }

    // We match the statement to a statement recognizable by Eliza 
    // which can then construct an appropriate response.
    for i, responses := range Responses {
        re := regexp.MustCompile("(?i)" + i) // compile regex and ignore case sensitive at the same time
        matches := re.FindStringSubmatch(statement)
        
        // If the statement matched any statements.
        if matches != nil {
            // The statement was matched then get the first match.
            // The matched regex group will match a captured fragment 
            // That captured fragment will then become part of the response
            var capture string
            if len(matches) > 1 {
                capture = split(matches[1])
            }

            // Choose a random response from the Responses array in replies.go
            response := random(responses)

            if strings.Contains(response, "%s") {
                response = fmt.Sprintf(response, capture)
            }

            return response
        }
    } // for

    // If nothing was matched return a default response from replies.go
    return random(Default)
	
} // Ask

// This method returns true if the statement is a start statement
func IsStartStatement(statement string) bool {
    statement = trim(statement)

    // It loops through the StartStatements array from replies.go
    // and gives the statemends in that array the value of "startStatement"
    for _, startStatement := range StartStatements {
        if statement == startStatement {
            return true
        }
    }
    return false
} // IsStartStatement

// This method returns true if the statement is an end statement
func IsEndStatement(statement string) bool {
    statement = trim(statement)

    // It loops through the EndStatements array from replies.go
    // and gives the statemends in that array the value of "endStatement"
    for _, endStatement := range EndStatements {
        if statement == endStatement {
            return true
        }
    }
    return false
} // IsStartStatement

// This method returns a copy of the string, with whitespace omitted.
func trim(statement string) string {
    statement = strings.TrimRight(statement, "\n.!")
    statement = strings.ToLower(statement)
    return statement
} // trim

// Generates a random index related to an element in the string array and returns that element
func random(list []string) string {
    rand.Seed(time.Now().UnixNano()) // generates a random number every nanosecond
    randomNumber := rand.Intn(len(list)) // returns a random index number related to whatever the array number from the array list
    return list[randomNumber]
}

// Reflect maps words in an input fragment (i.e. "I" -> "you").
// Delimiters (one or more characters that separates text strings) are matched with the split function using regular expressions
func split(match string) string{
    
    // Split is used to break up strings into a list of substrings
    words := strings.Split(match, " ")
    for i, word := range words {
        if reflectedWord, ok := Reflected[word]; ok {
            words[i] = reflectedWord
        }
    }

    // Join takes a list of strings and joins them together again
    return strings.Join(words, " ") 
}
