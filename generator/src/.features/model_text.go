package features

func Name(str string) Feature {
	return Raw(KEY_NAME, noDeps, withStaticExecutor(str))
}

func Text(text string) Feature {
	return Raw(KEY_TEXT, noDeps, withStaticExecutor(text))
}
