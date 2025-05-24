package runner

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func RunDocker(image, command string, input interface{}) (string, error) {
	inputStr := fmt.Sprintf("%v", input)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	containerCmd := exec.CommandContext(ctx, "docker", "run",
		"--rm",
		"-i",
		"--network=none",
		"--memory=64m",
		"--cpus=0.5",
		"-v", "/code:/code", // Internal Docker mount only
		image,
		"sh", "-c", command,
	)

	fmt.Printf("Running: docker run --rm -i --network=none --memory=64m --cpus=0.5 -v /code:/code %s sh -c '%s'\n", image, command)

	var stdout, stderr bytes.Buffer
	containerCmd.Stdout = &stdout
	containerCmd.Stderr = &stderr
	containerCmd.Stdin = bytes.NewBufferString(inputStr)

	err := containerCmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("execution timed out")
	}
	if err != nil {
		return stderr.String(), fmt.Errorf("exec error: %v", err)
	}

	return stdout.String(), nil
}

