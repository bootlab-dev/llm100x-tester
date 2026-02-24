package helpers

import (
	"fmt"
	"os/exec"
)

// CompileC 编译 C 文件
// workDir: 工作目录
// source: 源文件名 (如 "hello.c")
// output: 输出文件名 (如 "hello")
// needBootlab: 是否需要 bootlab.h (使用 -I.. 引入父目录)
func CompileC(workDir, source, output string, needBootlab bool) error {
	args := []string{
		"-o", output,
		source,
		"-lm",
		"-Wall",
		"-Werror",
	}

	// 如果需要 bootlab.h，添加 -I.. 使其能找到父目录的 bootlab.h
	if needBootlab {
		args = append(args, "-I..")
	}

	cmd := exec.Command("clang", args...)
	cmd.Dir = workDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("compilation failed: %s\nOutput:\n%s", err, string(out))
	}

	return nil
}
