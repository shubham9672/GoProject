package token

import (
	"time"
	"net/http"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

// var charlist string= "_${'`'}{|}~123abcde.fmnopqlABCDE@FJKLMNOPQRSTUVWXYZ456789stuvwxyz0!#$%&ijkrgh'*+-/=?^";
var mySigningKey = []byte("captainjacksparrowsayshi")

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}

func GenerateJWT(email string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = email
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        fmt.Errorf("Something Went Wrong: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}







// func GenerateToken(key string) string{
//      token := "";
//     for  i:= 0; i < key.length; i++ {
//         index := strings.index(key[i]) || charlist.length / 2;
//         randomIndex := Math.floor(Math.random() * index);
//         console.log(randomIndex < charlist.length);
      
//         token += charlist[randomIndex] + charlist[index - randomIndex];
//     }
//     return token;
// }
// func compareToken(token,key string) bool{
//     string := "";
//     for i := 0; i < len(token); i = i + 2 {
// 		temp:=string(token[i])
//      	index1 := strings.Index(temp,charlist)
//         index2 := strings.Index(string(token[i + 1]),charlist);
//         string += charlist[index1 + index2];
//     }
//     console.log(string);
//     if string == key{
//         return true;
//     }
//     return false;
// }