package main

func addToOrderedSet(set *[6]uint8, value uint8) {
	// If is present, just return
	for i, v := range set {
		if v == value {
			return
		} else if v == 0 {
			set[i] = value
			return
		}
	}
}

func removeFromOrderedSet(set *[6]uint8, value uint8) {
	found := false
	for i, v := range set {
		if value == v {
			found = true
		}

		if found {
			if i == 5 {
				set[i] = 0
			} else {
				set[i] = set[i+1]

				// Exit early, stop copying 0s
				if set[i] == 0 {
					return
				}
			}
		}
	}
}
