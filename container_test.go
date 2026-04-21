package container

import (
	"testing"
	"time"
)

func TestContainer(t *testing.T) {
	c, err := Start("pgvector/pgvector:pg17", "5432", "-e", "POSTGRES_PASSWORD=password")
	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}

	t.Log(c.Host)
	time.Sleep(5 * time.Second)
	t.Logf("%q", Logs(c.ID))

	if err := Stop(c.ID); err != nil {
		t.Fatalf("failed to stop container: %v", err)
	}
}

func TestContainer2(t *testing.T) {
	image := "postgis/postgis:17-3.5-alpine"
	port := "5432"
	args := []string{"-a", "amd64", "-e", "POSTGRES_PASSWORD=postgres"}

	c, err := Start(image, port, args...)
	if err != nil {
		t.Fatal(err)
	}

	if err := Stop(c.ID); err != nil {
		t.Fatalf("failed to stop container: %v", err)
	}
}
