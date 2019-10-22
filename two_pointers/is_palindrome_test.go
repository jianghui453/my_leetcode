package two_pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
    var s string
    var h, r bool

    s = "A man, a plan, a canal: Panama"
    h = true
    r = isPalindrome(s)
    t.Logf("%t h=%t r=%t\n", h==r, h, r)

    s = "race a car"
    h = false
    r = isPalindrome(s)
    t.Logf("%t h=%t r=%t\n", h==r, h, r)
}
