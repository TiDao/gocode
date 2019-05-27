/***********************************************************************
# File Name: word_test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-24 13:53:28
*********************************************************************/
package word 

import(
    "testing"
    "math/rand"
)

func randomPalindrome(rng *rand.Rand) string{
    n := rng.Intn(25)
    runes := make([]rune,n)
    for i := 0; i<(n+1)/2; i++{
        r := rune(rng.Intn(0x1000))
        runes[i] = r
        runes[n-1-i] = r
    }
    return string(runes)
}
func TestPalindrome(t *testing.T){
    if !IsPalindrome("detartrated"){
        t.Error(`IsPalindrome("datartrated")= false`)
    }

    if !IsPalindrome("kayak"){
        t.Error(`IsPalindrome("kayak") = false`)
    }
}

func TestNonPalindrome(t *testing.T) {
    if IsPalindrome("palindrome"){
        t.Error(`IsPalindrome("palindrome") = true`)
    }
}

func TestFrenchPalindrome(t *testing.T){
    if !IsPalindrome("été"){
        t.Error(`IsPalindrome("été") = false`)
    }
}

func TestCandlPalinrome(t *testing.T){
    input := "A man,a plan, a canal:Panama"
    if !IsPalindrome(input){
        t.Errorf(`IsPalindrome(%q) = false`,input)
    }
}

func BenchmarkIsPalindrome(b *testing.B){
    for i := 0;i < b.N; i++{
        IsPalindrome("A man ,a plan,a canal:Panma")
    }
}
