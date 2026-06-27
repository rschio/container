//go:build darwin

package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// Start starts the specified container for running tests.
func Start(image, port string, args ...string) (*Container, error) {
	arg := []string{"run", "-d"}
	arg = append(arg, args...)
	arg = append(arg, image)

	out := new(bytes.Buffer)
	outErr := new(bytes.Buffer)
	cmd := exec.Command("container", arg...)
	cmd.Stdout = out
	cmd.Stderr = outErr
	if err := cmd.Run(); err != nil {
		if strings.Contains(outErr.String(), "XPC connection error: Connection invalid") {
			return nil, fmt.Errorf("container system not initiated. Run:\n$ container system start")
		}
		return nil, fmt.Errorf("could not start container %s: %w", image, err)
	}

	id := out.String()[:36]
	ip, err := extractIP(id)
	if err != nil {
		return nil, fmt.Errorf("could not extract ip: %w", err)
	}

	c := Container{
		ID:   id,
		Host: net.JoinHostPort(ip, port),
	}

	return &c, nil
}

// Stop stops and removes the specified container.
func Stop(id string) error {
	if err := exec.Command("container", "stop", id).Run(); err != nil {
		return fmt.Errorf("could not stop container %q: %w", id, err)
	}
	if err := exec.Command("container", "rm", id).Run(); err != nil {
		return fmt.Errorf("could not remove container %q: %w", id, err)
	}

	return nil
}

// Logs output logs from the running container.
func Logs(id string) []byte {
	out, err := exec.Command("container", "logs", id).CombinedOutput()
	if err != nil {
		return nil
	}

	return out
}

func extractIP(id string) (string, error) {
	out := new(bytes.Buffer)
	cmd := exec.Command("container", "inspect", id)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("could not inspect container %s: %w", id, err)
	}

	return ipFromInspectOutput(out.Bytes())
}

func ipFromInspectOutput(data []byte) (string, error) {
	type inspectContainer []struct {
		Status struct {
			Networks []struct {
				Ipv4Address string `json:"ipv4Address"`
			} `json:"networks"`
		} `json:"status"`
	}

	var out inspectContainer
	if err := json.Unmarshal(data, &out); err != nil {
		return "", fmt.Errorf("could not decode json: %w", err)
	}
	if len(out) == 0 || len(out[0].Status.Networks) == 0 {
		return "", fmt.Errorf("could not get container IP")
	}

	ipWithMask := out[0].Status.Networks[0].Ipv4Address
	ip, _, _ := strings.Cut(ipWithMask, "/")

	return ip, nil
}
