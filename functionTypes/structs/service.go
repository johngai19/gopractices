package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func (s Service) getName() string {
	return s.description
}
func (s Service) getCost(annual bool) float64 {
	if annual {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
