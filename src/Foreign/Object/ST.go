package Data_Foreign_Object_ST

func NewImpl() func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		m := make(map[string]interface{})
		return &m
	}
}

func PeekImpl(just func(interface{}) interface{}, nothing interface{}, k string, m *map[string]interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		if val, ok := (*m)[k]; ok {
			return just(val)
		}
		return nothing
	}
}

func Poke(k string, v interface{}, m *map[string]interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		(*m)[k] = v
		return m
	}
}

func DeleteImpl(k string, m *map[string]interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		delete(*m, k)
		return m
	}
}
