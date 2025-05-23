package runner

import (
    "bytes"
    "context"
    "fmt"
    "os/exec"
    "time"
)

// RunDocker runs code inside a docker container using the specified image and command.
// codeFile is the filename inside /tmp folder (mounted as /code).
// input is passed to container stdin.
func RunDocker(image, command, code string, input interface{}) (string, error) {
    // Convert input to string
    inputStr := ""
    if s, ok := input.(string); ok {
        inputStr = s
    } else {
        inputStr = fmt.Sprintf("%v", input)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // Compose docker run command with resource limits and no network
    containerCmd := exec.CommandContext(ctx, "docker", "run",
        "--rm",
        "--network=none",    // isolate container network
        "--memory=64m",      // limit memory usage
        "--cpus=0.5",        // limit CPU usage
        "-v", "/tmp:/code",  // mount /tmp to /code inside container
        image,
        "sh", "-c", command,
    )

    var stdout, stderr bytes.Buffer
    containerCmd.Stdout = &stdout
    containerCmd.Stderr = &stderr
    containerCmd.Stdin = bytes.NewBufferString(inputStr) // pass input via stdin

    err := containerCmd.Run()

    // Timeout error
    if ctx.Err() == context.DeadlineExceeded {
        return "", fmt.Errorf("execution timed out")
    }

    // On error return stderr + err
    if err != nil {
        return stderr.String(), fmt.Errorf("exec error: %v", err)
    }

    return stdout.String(), nil
}

