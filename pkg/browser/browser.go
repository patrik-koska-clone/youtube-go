package browser

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func LaunchBrowser(videoID string) error {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return errors.New("unsupported platform")
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open browser:\n%v", err)
	}

	return nil
}
