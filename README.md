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

### 2. Tokenize and Index History
- Split each command into tokens (command + args).
- Build:
    token_index = { token: [command_indices...] }
- Keep frequency counts for each token.

### 3. Implement Base Edit Distance
- Use Levenshtein or Damerau–Levenshtein for typo correction.
- This forms the accuracy core for fuzzy matching.


