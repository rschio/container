//go:build darwin

package container

import "testing"

func Test_ipFromInspectOutput(t *testing.T) {
	data := []byte(inspectOutput)
	want := "192.168.65.10"
	got, err := ipFromInspectOutput(data)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

var inspectOutput = `[
  {
    "configuration": {
      "capAdd": [],
      "capDrop": [],
      "creationDate": "2026-06-27T19:14:13Z",
      "dns": {
        "nameservers": [],
        "options": [],
        "searchDomains": []
      },
      "id": "493b7e75-293c-436c-a51d-69fc0ab57ec0",
      "image": {
        "descriptor": {
          "digest": "sha256:48c8ad3a7284b82be4482a52076d47d879fd6fb084a1cbfccbd551f9331b0e40",
          "mediaType": "application/vnd.oci.image.index.v1+json",
          "size": 10293
        },
        "reference": "docker.io/library/postgres:18-alpine"
      },
      "initProcess": {
        "arguments": [
          "postgres"
        ],
        "environment": [
          "PG_SHA256=0d5b903b1e5fe361bca7aa9507519933773eb34266b1357c4e7780fdee6d6078",
          "DOCKER_PG_LLVM_DEPS=llvm19-dev \t\tclang19",
          "GOSU_VERSION=1.19",
          "PGDATA=/var/lib/postgresql/18/docker",
          "POSTGRES_PASSWORD=postgres",
          "LANG=en_US.utf8",
          "PG_MAJOR=18",
          "PG_VERSION=18.0",
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
        ],
        "executable": "docker-entrypoint.sh",
        "rlimits": [],
        "supplementalGroups": [],
        "terminal": false,
        "user": {
          "id": {
            "gid": 0,
            "uid": 0
          }
        },
        "workingDirectory": "/"
      },
      "labels": {},
      "mounts": [],
      "networks": [
        {
          "network": "default",
          "options": {
            "hostname": "493b7e75-293c-436c-a51d-69fc0ab57ec0",
            "mtu": 1280
          }
        }
      ],
      "platform": {
        "architecture": "arm64",
        "os": "linux"
      },
      "publishedPorts": [],
      "publishedSockets": [],
      "readOnly": false,
      "resources": {
        "cpuOverhead": 1,
        "cpus": 4,
        "memoryInBytes": 1073741824
      },
      "rosetta": false,
      "runtimeHandler": "container-runtime-linux",
      "ssh": false,
      "stopSignal": "SIGINT",
      "sysctls": {},
      "useInit": false,
      "virtualization": false
    },
    "id": "493b7e75-293c-436c-a51d-69fc0ab57ec0",
    "status": {
      "networks": [
        {
          "hostname": "493b7e75-293c-436c-a51d-69fc0ab57ec0",
          "ipv4Address": "192.168.65.10/24",
          "ipv4Gateway": "192.168.65.1",
          "ipv6Address": "fda1:b93b:e3f4:e8d9:fcf3:11ff:fe45:e3c9/64",
          "macAddress": "fe:f3:11:45:e3:c9",
          "mtu": 1280,
          "network": "default"
        }
      ],
      "startedDate": "2026-06-27T19:14:15Z",
      "state": "running"
    }
  }
]`
