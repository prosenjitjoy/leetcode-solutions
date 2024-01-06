type Codec struct {
    encoder map[string]string
    decoder map[string]string
    baseUrl string
}


func Constructor() Codec {
    return Codec {
        encoder: map[string]string{},
        decoder: map[string]string{},
        baseUrl: "http://tinyurl.com",
    }
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if _, ok := this.encoder[longUrl]; !ok {
        shortUrl := fmt.Sprintf("%s/%v", this.baseUrl, len(this.encoder)+1)
        this.encoder[longUrl] = shortUrl
        this.decoder[shortUrl] = longUrl
        return shortUrl
    }

    return this.encoder[longUrl]
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    return this.decoder[shortUrl]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
