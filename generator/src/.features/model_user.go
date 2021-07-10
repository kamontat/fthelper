package features

import (
	"os"

	"github.com/kamontat/fthelper/shared/models"
)

func User(uid, gid int) Feature {
	var empty = make(models.Mapper)
	if uid >= 0 {
		empty.Set("uid", uid)
	}
	if gid >= 0 {
		empty.Set("gid", gid)
	}
	return Raw(KEY_USER, noDeps, withStaticExecutor(empty))
}

func Chmod(mode os.FileMode) Feature {
	return Raw(KEY_CHMOD, noDeps, withStaticExecutor(mode))
}

func ExecuteChmod() Feature {
	return Chmod(os.ModePerm)
}

func RWChmod() Feature {
	return Chmod(0766)
}
