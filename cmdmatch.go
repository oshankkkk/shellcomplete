package main

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
func dfs(s1,s2 string, ind1,ind2 int,memo map[[2]int]int)int{
	if ind1==len(s1){
		return len(s2)-ind2
	} 
	if ind2==len(s2){
		return len(s1)-ind1
	}
	if s1[ind1]==s2[ind2]{
		 memo[[2]int{ind1,ind2}]=dfs(s1,s2,ind1+1,ind2+1,memo)
		 return  memo[[2]int{ind1,ind2}]

	}

	a:=dfs(s1,s2,ind1+1,ind2,memo)
	b:=dfs(s1,s2,ind1,ind2+1,memo)
	c:=dfs(s1,s2,ind1+1,ind2+1,memo)
	memo[[2]int{ind1,ind2}]=1+min(a,b,c)
	return memo[[2]int{ind1,ind2}]

}
func y(cmd1,cmd2 string)int{
	memo:=make(map[[2]int]int)
return dfs(cmd1,cmd2,0,0,memo)

}
func editDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],    // deletion
					dp[i][j-1],    // insertion
					dp[i-1][j-1],  // substitution
				)
			}
		}
	}

	return dp[m][n]
}

func filterhistory(){
	
}

