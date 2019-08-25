package min_distance

import "testing"

func TestMinDistance(t *testing.T) {
    var word1, word2 string
    var hope, ret int

    word1 = "horse"
    word2 = "ros"
    hope = 3
    ret = minDistance(word1, word2)
    t.Logf("word1=%s word2=%s hope=%d ret=%d", word1, word2, hope, ret)

    word1 = "intention"
    word2 = "execution"
    hope = 5
    ret = minDistance(word1, word2)
    t.Logf("word1=%s word2=%s hope=%d ret=%d", word1, word2, hope, ret)

    word1 = "sea"
    word2 = "eat"
    hope = 2
    ret = minDistance(word1, word2)
    t.Logf("word1=%s word2=%s hope=%d ret=%d", word1, word2, hope, ret)
}