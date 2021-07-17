package bookings

func MergeBookingIdentifiers(s1 []Identifier, s2 []Identifier) []Identifier {

	if s1 == nil && s2 != nil {
		return s2
	}

	if s2 == nil && s1 != nil {
		return s1
	}

	if s1 == nil && s2 == nil {
		return nil
	}

	for _, s := range s1 {

		alreadyIn := false
		for _, ss2 := range s2 {
			if ss2 == s {
				alreadyIn = true
				break
			}
		}

		if !alreadyIn {
			s2 = append(s2, s)
		}

	}
	return s2
}
