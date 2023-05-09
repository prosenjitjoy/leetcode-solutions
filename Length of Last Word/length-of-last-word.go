func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    wordList := strings.Split(s, " ")
    word := wordList[len(wordList)-1]
    return len(word)
}