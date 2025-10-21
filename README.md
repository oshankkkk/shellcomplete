# Shell autocomplete project 
##### (kinda like fig but in written in go)

### Topics to Understand
- Understanding bubbletea TUI
- Edit distance algorithms (Levenshtein, Damerau, Myers)/dynamic programming
- Tokenization (splitting on whitespace, quotes, flags)
- Bash history structure (~/.bash_history, HISTTIMEFORMAT)
- Basic ranking logic (min distance → best match)

### 1. Read and Clean Bash History
- Parse `~/.bash_history` (ignore timestamps or metadata).
- Remove duplicates and normalize spacing.
- Store each command as a list of tokens:
- Store each command line as a string list → e.g.
["git", "commit", "-m", "init"].
#### - Bash intro
- Its the CLI,the 1st OS(kinda)
- You interact with the kernal through a layer of abstraction called bash
- You can also automate those cmds with bash scripts
#### Bash ENVVARIABLES
- So they are extra context passed to the program thats running buy bash.
- The stuff we mostly goes into the program as ENV variables
- When a program starts bash clones itself, and then runs the new program in that clone, and envs are passed so that
new program whos whats up
- Env variables are just a way to passed dynamic values made to the user envirnment by the user to programs
- Eg:?
#### Bash history
- Bash stores history in memory and then writes it to a file when exit bash
- There are a lot of env variables that effects bash history
- Timetamps are dinoted after # symbol. Its secounds from 1970's when the unix system was made
#### Single qoutes and double quotes in shell
- when you make a bash cmd, you are passing arguments seperated by spaces.
- Quotes will tell bash to take things inside the quotes as 1 arguemnt 
- Double quotes will take stuff are 1 argument but will still show variables n stuff to bash
- Single qoutes just take all things as strings in 1 argument
#### Escape Characters
- \ is a escape character which means \symbols = mark the symbol as just a string do not use it as a keyword or somthing.

#### Regex
- Its a special string searching sytax built into a lot of programming lanugages. 

#### Text parsing
- This is how most languages understand text. 
- They need to know which parts are tokens (echo, hi $USER)
    What’s inside quotes
    What’s escaped
    Where commands end
- There are 2 parts in text parsing:
    1. Tokenization (Logical analysis)- breaking down the text 
    2. Parsing (Syntactic analysis) - Understanding how the broken stuff work together in a way that pc understands

- Its breaking coplex strings in to a array of substrings to convert to bytes so the pc can understand it.
- When tokenizing bash history you have/ to handle the timestamps.
- 

### 2. Tokenize and Index History
- Split each command into tokens (command + args).
- Build:
    token_index = { token: [command_indices...] }
- Keep frequency counts for each token.

### 3. Implement Base Edit Distance
- Use Levenshtein or Damerau–Levenshtein for typo correction.
- This forms the accuracy core for fuzzy matching.


