package main
import(
	"fmt"
	"os"
	"os/user"
	"bufio"
	"strings"
	"regexp"
	"unicode"
)


func loadhistory()(*os.File,error){

	history:=os.Getenv("HISTFILE")

	if history==""{
		currentuser,err:=user.Current()
		if err!=nil{
			fmt.Println("error",err)
			return nil,err
		}
		history=currentuser.HomeDir+"/.bash_history"
	}
	bashhistory, err:=os.Open(history)
	if err!=nil{
		fmt.Println("error",err)
		return nil,err
	}

	return bashhistory,nil

}

func readbashhistory2(historyFile *os.File) map[string][]string {
	defer historyFile.Close()

	cmdCache := make(map[string][]string)
	scanner := bufio.NewScanner(historyFile)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip empty lines and timestamps (#1676427720)
		}

		// Tokenize line into fields (split on any whitespace)
		fields := strings.Fields(line)

		if len(fields) > 0 {
			cmdCache[line] = fields
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return cmdCache
}

func readbashhistory1(historyfile *os.File) map[string][]string{
	//historylist:=[]string{}
	defer historyfile.Close() 

	cmdcashe:=make(map[string][]string)
	scanner := bufio.NewScanner(historyfile)
	
	for scanner.Scan() {
		line := scanner.Text()      
		_,ok:=cmdcashe[line]
		if ok==false{
		cmdcashe[line]=strings.Fields(line)

	}
//		historylist=append(historylist,line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return cmdcashe
//	return historylist

}

// ReadBashHistory reads a bash history file and returns a map
// from the cleaned command line -> tokens (quote/escape-aware).
func readbashhistory(historyFile *os.File) map[string][]string {
	defer historyFile.Close()

	cmdCache := make(map[string][]string)
	scanner := bufio.NewScanner(historyFile)

	// regex to detect lines like: "   42  some command"
	leadingNumberRe := regexp.MustCompile(`^\s*\d+\s+`)
	// regex to detect timestamp lines like: "#1697169300"
	timestampRe := regexp.MustCompile(`^#\d+$`)

	for scanner.Scan() {
		raw := scanner.Text()
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		// skip pure timestamp lines (Debian-style)
		if timestampRe.MatchString(line) {
			continue
		}
		// strip leading history numbers if present (from `history` output)
		line = leadingNumberRe.ReplaceAllString(line, "")
		if line == "" {
			continue
		}

		tokens := shellTokenize(line)
		if len(tokens) > 0 {
			cmdCache[line] = tokens
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
	}
	return cmdCache
}

// shellTokenize tokenizes a shell-like command line handling:
// - double quotes "...", single quotes '...'
// - backslash escaping \x
// - treats any unicode.IsSpace(r) as separator when not inside quotes
func shellTokenize(line string) []string {
	var tokens []string
	var b strings.Builder
	inSingle := false
	inDouble := false
	escaped := false

	for _, r := range line {
		if escaped {
			// append the next rune literally
			b.WriteRune(r)
			escaped = false
			continue
		}

		switch r {
		case '\\':
			// start escape (unless inside single quotes, where backslash is literal)
			if inSingle {
				b.WriteRune(r)
			} else {
				escaped = true
			}
		case '\'':
			if inDouble {
				// inside double quotes, single quote is literal
				b.WriteRune(r)
			} else {
				// toggle single-quote mode
				if inSingle {
					inSingle = false
				} else {
					inSingle = true
				}
				// DON'T add the quote char to token
			}
		case '"':
			if inSingle {
				// inside single quotes double-quote is literal
				b.WriteRune(r)
			} else {
				// toggle double-quote mode
				if inDouble {
					inDouble = false
				} else {
					inDouble = true
				}
				// DON'T add the quote char to token
			}
		default:
			if !inSingle && !inDouble && unicode.IsSpace(r) {
				// separator when not in quotes
				if b.Len() > 0 {
					tokens = append(tokens, b.String())
					b.Reset()
				}
			} else {
				b.WriteRune(r)
			}
		}
	}

	// add last token
	if b.Len() > 0 {
		tokens = append(tokens, b.String())
	}

	return tokens
}

