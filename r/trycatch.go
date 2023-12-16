package r

// TryCatch is a struct that provides a way to handle errors within a specific code block.
// It has a single field `try` which is a function that contains the code block to be executed.
type TryCatch struct {
	try func()
}

// Try is a function that takes in a try function and returns a TryCatch instance. The try function
// should be a function without any input parameters and without any return values. The TryCatch
// instance can be used to catch any panics that occur within the try function.
//
// Example usage:
//
//	Try(func() {
//	    // code to try
//	}).Catch(func(err any) {
//	    // code to handle the panic
//	})
func Try(try func()) *TryCatch {
	return &TryCatch{
		try: try,
	}
}

// Catch is a method of the TryCatch struct that takes in a catch function as parameter. The catch function should be
// a function with one parameter of type `any` and no return value
func (tc *TryCatch) Catch(catch func(any)) {
	defer func() {
		if r := recover(); r != nil {
			catch(r)
		}
	}()
	tc.try()
}

// Raise is a function that panics with the given error. It is used to intentionally cause a panic in the code.
// Example usage:
//
//	Raise("Something went wrong")
func Raise(err any) {
	panic(err)
}
