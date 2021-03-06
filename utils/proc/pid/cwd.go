package pid

import (
	"fmt"
	"os"
	"path/filepath"
)

func CWD(pid string) string {
	path, err := os.Readlink(filepath.Join(Ctx.Procfs, pid, "cwd"))
	if err != nil {
		return fmt.Sprint(err)
	}

	return path
}