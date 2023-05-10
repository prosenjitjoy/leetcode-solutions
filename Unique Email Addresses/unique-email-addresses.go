func numUniqueEmails(emails []string) int {
    m := map[string]int{}
    for _, email := range emails {
        parts := strings.Split(email, "@")
        localName := strings.ReplaceAll(strings.Split(parts[0], "+")[0], ".", "")
        domainName := parts[1]
        mail := localName + "@" + domainName
        m[mail]++
    }
    return len(m)
}