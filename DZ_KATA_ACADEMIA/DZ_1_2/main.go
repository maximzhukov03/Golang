for _, elem := range fizzBuzzTest {
	if got := fizzBuzz(elem.num); got != elem.res {
		t.Errorf("fizzBuzz(%d): %q != %q", elem.num, got, elem.res)
	}
}