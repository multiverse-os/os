package slice

func Contains(elements []interface{}, check interface{}) (exists bool) {
	switch v := i.(type) {
	case int:
		return ContainsString(elements.(int), check.(int))
	case string:
		return ContainsString(elements.(string), check.(string))
	default:
		for _, e := range elements {
			if e == check {
				exists = true
			}
		}
		return exists
	}
}

func ContainsString(strings []string, check string) (exists bool) {
	for _, s := range strings {
		if s == check {
			exists = true
		}
	}
	return exists
}

func ContainsInt(ints []int, check int) (exists bool) {
	for _, i := range ints {
		if i == check {
			exists = true
		}
	}
	return exists
}
