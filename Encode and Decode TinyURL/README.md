# [Problem statement](https://leetcode.com/problems/encode-and-decode-tinyurl)

> Note: This is a companion problem to the [System Design](https://leetcode.com/discuss/interview-question/system-design/) problem: [Design TinyURL](https://leetcode.com/discuss/interview-question/124658/Design-a-URL-Shortener-%28-TinyURL-%29-System/).

TinyURL is a URL shortening service where you enter a URL such as `https://leetcode.com/problems/design-tinyurl` and it returns a short URL such as `http://tinyurl.com/4e9iAk`. Design a class to encode a URL and decode a tiny URL.

There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

Implement the `Solution` class:

* `Solution()` Initializes the object of the system.
* `String encode(String longUrl)` Returns a tiny URL for the given `longUrl`.
* `String decode(String shortUrl)` Returns the original long URL for the given `shortUrl`. It is guaranteed that the given `shortUrl` was encoded by the same object.

**Example 1:**


**Input:** url = "https://leetcode.com/problems/design-tinyurl"
**Output:** "https://leetcode.com/problems/design-tinyurl"

**Explanation:**
Solution obj = new Solution();
string tiny = obj.encode(url); // returns the encoded tiny url.
string ans = obj.decode(tiny); // returns the original url after decoding it.

**Constraints:**

* `1 <= url.length <= 104`
* `url` is guranteed to be a valid URL.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1138487510/)

```go
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

```