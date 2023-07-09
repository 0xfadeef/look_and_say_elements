package look_and_say_elements

import "fmt"
import "testing"


func decay(sequence string) string {
	if len(sequence) == 0 {
		return sequence
	}
	var result []byte
	ch := sequence[0]
	count := 1

	for n := 1; n < len(sequence); n++ {
		if ch == sequence[n] {
			count++
		} else {
			result = fmt.Appendf(result, "%d%c", count, ch)
			ch = sequence[n]
			count = 1
		}
	}
	result = fmt.Appendf(result, "%d%c", count, ch)
	return string(result)
}

func TestNumbers(t *testing.T) {
	for n, elem := range Elements {
		if n != elem.Number {
			t.Fatalf(
				"Element number %d/%s does not match its index position (%d)",
				elem.Number, elem.Name, n)
		}
	}
}

func TestDecay(t *testing.T) {
	var N = len(Elements)

	for _, elem := range Elements {
		product := decay(elem.Sequence)
		s := 0

		for _, num := range elem.Products {
			if 0 >= num || num >= N {
				t.Fatalf(
					"Invalid number (%d) in Products of element %d/%s " +
					"(expected to be in range 1..%d)",
					num, elem.Number, elem.Name, N-1)
			}
			sequence := Elements[num].Sequence
			e := s + len(sequence)

			if e > len(product) || product[s:e] != sequence {
				t.Fatalf(
					"Element %d/%s does not decay into sum of its Products!",
					elem.Number, elem.Name)
			} else {
				s = e
			}
		}
	}
}

