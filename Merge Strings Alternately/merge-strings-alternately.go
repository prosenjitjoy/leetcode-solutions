func mergeAlternately(word1 string, word2 string) string {
    var res []byte
    i,j:=0,0;
    for i<len(word1) && j<len(word2) {
        res = append(res, word1[i])
        res = append(res, word2[j])
        i,j=i+1,j+1
    }

    res = append(res, word1[i:]...)
    res = append(res, word2[j:]...)
    return string(res)
}