package features

import "github.com/kamontat/fthelper/shared/fs"

func FileSystemFormat(featureName string, format *fs.FormatModel) Feature {
	return Raw(MergeKeys(featureName, KEY_FORMAT), Dependencies{
		featureName: REQUIRE,
		KEY_KEY:     OPTIONAL,
	}, withStaticExecutor(format))
}

func FileSystemFileTplFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Filename: "{{ .filename.name }}.{{ .filename.ext }}.tpl",
	})
}
func FileSystemFileTplWithEnvFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Filename: "{{ name .filename.name .env }}.{{ .filename.ext }}.tpl",
	})
}
func FileSystemFileTplWithKeyFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Filename: "{{ name .filename.name .key }}.{{ .filename.ext }}.tpl",
	})
}
func FileSystemFileTplWithEnvKeyFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Filename: "{{ name .filename.name .env .key }}.{{ .filename.ext }}.tpl",
	})
}

func FileSystemFileWithEnvFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Filename: "{{ name .filename.name .env }}.{{ .filename.ext }}",
	})
}
func FileSystemFileWithKeyFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Filename: "{{ name .filename.name .key }}.{{ .filename.ext }}",
	})
}
func FileSystemFileWithEnvKeyFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Filename: "{{ name .filename.name .env .key }}.{{ .filename.ext }}",
	})
}

func FileSystemDirWithEnvFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Dirname: "{{ .env }}/{{ .path.dirname }}",
	})
}

func FileSystemDirWithClusterFormat(featureName string) Feature {
	return FileSystemFormat(featureName, &fs.FormatModel{
		Dirname: "{{ .cluster }}/{{ .path.dirname }}",
	})
}
