package twelve

import "fmt"

type data struct {
	day  string
	gift string
}

var words = map[int]data{
	1:  data{day: "first", gift: "a Partridge in a Pear Tree."},
	2:  data{day: "second", gift: "two Turtle Doves, and "},
	3:  data{day: "third", gift: "three French Hens, "},
	4:  data{day: "fourth", gift: "four Calling Birds, "},
	5:  data{day: "fifth", gift: "five Gold Rings, "},
	6:  data{day: "sixth", gift: "six Geese-a-Laying, "},
	7:  data{day: "seventh", gift: "seven Swans-a-Swimming, "},
	8:  data{day: "eighth", gift: "eight Maids-a-Milking, "},
	9:  data{day: "ninth", gift: "nine Ladies Dancing, "},
	10: data{day: "tenth", gift: "ten Lords-a-Leaping, "},
	11: data{day: "eleventh", gift: "eleven Pipers Piping, "},
	12: data{day: "twelfth", gift: "twelve Drummers Drumming, "},
}

func Song() string {
	var song string
	for i := 0; i < 12; i++ {
		song += Verse(i + 1)
		if i < 11 {
			song += "\n"
		}
	}
	return song
}

func Verse(verse int) string {
	v := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", words[verse].day)
	for i := verse; i > 0; i-- {
		v += words[i].gift
	}
	return v
}
