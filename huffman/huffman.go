package main



func Encode(s string) string {



    return "1000011101001000110010011101100111001001000111110010011111011111100010001111110100111001001011111011101000111111001";
}

func Counter(s string) map[rune]int {
    counter := make(map[rune]int);

    for _, val := range s {
        _, ok := counter[val];
        if !ok {
            counter[val] = 1;
        } else {
            counter[val] += 1;
        }
    }
    return counter;
}
