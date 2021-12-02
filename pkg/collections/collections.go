package collections

func Map[T, Y any](list []T, f func(T) (Y, error)) ([]Y, error) {
	var err error
	result := make([]Y, len(list))
	for i, v := range list {
		result[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
