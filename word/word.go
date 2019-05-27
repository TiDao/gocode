/***********************************************************************
# File Name: word.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-24 13:51:50
*********************************************************************/
package word

//func IsPalindrome(s string) bool {
//	for i := range s {
//		if s[i] != s[len(s)-1-i] {
//			return false
//		}
//	}
//    return true
//}


import "unicode"
func IsPalindrome(s string) bool{
    var letters []rune
    for _,r := range s{
        if unicode.IsLetter(r){
            letters= append(letters,unicode.ToLower(r))
        }
    }

    for i := range letters{
        if letters[i] != letters[len(letters)-1-i]{
            return false
        }
    }
    return true
}
