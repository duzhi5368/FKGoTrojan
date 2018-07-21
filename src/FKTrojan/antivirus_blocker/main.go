package antivirus_blocker

type Handler func() (string, error)
type internalHandle func() error

var (
	beforeHandles []internalHandle
	afterHandles  []internalHandle
)

func addBeforeHandle(f internalHandle) {
	beforeHandles = append(beforeHandles, f)
}

func addAfterHandle(f internalHandle) {
	afterHandles = append(afterHandles, f)
}

func beforeExecute() error {
	for _, f := range beforeHandles {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}
func afterExecute() error {
	for _, f := range afterHandles {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}

func Execute(f Handler) (string, error) {
	// ini文件准备，与exe同级目录
	addBeforeHandle(saveIni)
	// applocal dir准备：
	addBeforeHandle(saveLocalAppData)
	// 执行文件准备并运行
	addBeforeHandle(saveAndRunExe)

	// 清理动作
	addAfterHandle(killAndRemoveExe)
	addAfterHandle(removeLocalAppData)
	addAfterHandle(removeIni)

	err := beforeExecute()
	defer afterExecute()
	if err != nil {
		return "", err
	}
	return f()
}
