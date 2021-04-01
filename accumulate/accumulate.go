package accumulate

//Accumulate takes a collection of strings and an operation. Then  executes the operation over strings in the collection
func Accumulate(collection []string, operation func(string) string) []string {
	for i, str := range collection {
		collection[i] = operation(str)
	}
	return collection
}
