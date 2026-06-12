package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if (len(a) != len(b)) {
        return 0, errors.New("Diferent lengths")
    }
    if a == "" && b == "" {
		return 0, nil 
	}
	
	if a == "" {
		return 0, errors.New("first string is empty while second string has length")
	}
	
	if b == "" {
		return 0, errors.New("second string is empty while first string has length")
	}
    
    if (len(a) < 2 || len(b) < 2) {
        if a[0] == b[0] {
            return 0, nil
        }
       	return 1, nil
    }
    distance := 0
    for i := range a {
        if a[i] != b[i] {
            distance++
        }
    }
    return distance, nil
}
