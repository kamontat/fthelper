package features

func Debug() Feature {
	return Raw(MergeKeys(KEY_MODE, KEY_DEBUG), noDeps, withStaticExecutor(true))
}

func Disabled() Feature {
	return Raw(MergeKeys(KEY_MODE, KEY_DISABLE), noDeps, withStaticExecutor(true))
}

// Toggle is for switch enabled/disabled of current generator
func Toggle(b bool) Feature {
	return Raw(MergeKeys(KEY_MODE, KEY_DISABLE), noDeps, withStaticExecutor(b))
}

func ToggleConfig(configpath string) Feature {
	return Raw(MergeKeys(KEY_MODE, KEY_DISABLE), Dependencies{
		KEY_CONFIG: REQUIRE,
	}, withConfigExecutor(configpath))
}
