package pwt

func (algo Algo) String() string {
	switch algo {
	case ALGO_HS256:
		return "HS256"
	case ALGO_HS384:
		return "HS384"
	case ALGO_HS512:
		return "HS512"
	//case ALGO_RS256:
	//	return "RS256"
	//case ALGO_RS384:
	//	return "RS384"
	//case ALGO_RS512:
	//	return "RS512"
	//case ALGO_ES256:
	//	return "ES256"
	//case ALGO_ES384:
	//	return "ES384"
	//case ALGO_ES512:
	//	return "ES512"
	//case ALGO_PS256:
	//	return "PS256"
	//case ALGO_PS384:
	//	return "PS384"
	default:
		panic("unknown algo type")
	}
}
