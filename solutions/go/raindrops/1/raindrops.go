package raindrops

import(
    "fmt"
    "strconv"
)

func Convert(number int) string {
    var resp string
    resp = ""
	if number%3 == 0 {
        resp = resp + "Pling"
        if number%5 == 0 {
            resp = resp + "Plang"
        }
        if number%7 == 0 {
            resp = resp + "Plong"
        }
    } else if number%5 == 0 {
        resp = resp + "Plang"
        if number%7 == 0 {
            resp = resp + "Plong"
        }
    } else if number%7 == 0 {
        resp = resp + "Plong"
    } else {
        resp = strconv.Itoa(number)
    }
    return resp
}

func main() {
	fmt.Println(Convert(15))
}