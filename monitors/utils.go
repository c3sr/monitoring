package monitors

func memberQ(set []string, elem string) bool {
	for _, e := range set {
		if e == elem {
			return true
		}
	}
	return false
}
