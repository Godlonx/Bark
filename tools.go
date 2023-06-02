package bark

import(
	//"fmt"
	//"unicode"
)
type registerError string
const (
	BadUsername registerError = "bad username"
	BadPassword               = "bad password"
	UnequalPassword           = "unequal password"
	None					  = "none"
)

func Check(registerData RegisterData) (bool,registerError){
	// if (!verifyPassword(registerData.Password)) {
	// 	return false,BadPassword
	// }
	if (registerData.Password != registerData.Passwordverif) {
		return false,UnequalPassword
	}
	return true,None
}

// func verifyPassword(s string) (bool) {
//     letters := 0
//     for _, c := range s {
//         switch {
//         case unicode.IsNumber(c):
//             number := true
//         case unicode.IsUpper(c):
//             upper := true
//             letters++
//         case unicode.IsPunct(c) || unicode.IsSymbol(c):
//             special := true
//         case unicode.IsLetter(c) || c == ' ':
//             letters++
//         default:
//             return false
//         }
//     }
//     sevenOrMore = letters >= 7
// 	if (number&&upper&&special&&sevenOrMore) {
// 		return true
// 	}
//     return false
// }