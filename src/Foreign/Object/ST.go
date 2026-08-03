package Data_Foreign_Object_ST

func NewImpl(_ interface{}) interface{} {
	return make(map[string]interface{})
}

func PeekImpl(just func(interface{}) interface{}, nothing interface{}, k string, m map[string]interface{}, _ interface{}) interface{} {
	if val, ok := m[k]; ok {
		return just(val)
	}
	return nothing
}

func Poke(k string, v interface{}, m map[string]interface{}, _ interface{}) interface{} {
	m[k] = v
	return m
}

func DeleteImpl(k string, m map[string]interface{}, _ interface{}) interface{} {
	delete(m, k)
	return m
}
