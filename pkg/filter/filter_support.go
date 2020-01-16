package filter

func compareFloat64(v1 float64, v2 float64, flag string) bool {
	switch flag {
	case LT:
		return v1 < v2
	case LTE:
		return v1 <= v2
	case GT:
		return v1 > v2
	case GTE:
		return v1 >= v2
	case EQL:
		return v1 == v2
	default:
		panic("Did not receive proper filter.COMPARE")
	}
}

func compareString(v1 string, v2 string, flag string) bool {
	switch flag {
	case LT:
		return v1 < v2
	case LTE:
		return v1 <= v2
	case GT:
		return v1 > v2
	case GTE:
		return v1 >= v2
	case EQL:
		return v1 == v2
	default:
		panic("Did not receive proper filter.COMPARE")
	}
}
