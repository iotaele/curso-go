package raindrops

import(
    "fmt"
    "strconv"
)

func Convert(number int) string {
    var resp string
    resp = ""
	if number%3 == 0 {
        resp += "Pling"
    }
    if number%5 == 0 {
        resp += "Plang"
    }
    if number%7 == 0 {
        resp += "Plong"
    }
    if resp == "" {
        resp = strconv.Itoa(number)
    }
    return resp
}

func main() {
	fmt.Println(Convert(15))
}