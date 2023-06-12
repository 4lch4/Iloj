package iloj

// If the given error is not nil, logs the error message to the console using my glogger package.
// It will log the error message with the context "Iloj#HandleError".
func HandleError(e error, ctxIn *string) {
	var ctx string

	// Check if the context is nil and set it to a default value if it is.
	if ctxIn == nil {
		ctx = "Iloj#HandleError"
	} else {
		ctx = *ctxIn
	}

	// If the error is not nil, log the error message.
	if e != nil {
		logger.Error(e.Error(), ctx)
	}
}
