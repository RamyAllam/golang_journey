package main

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/lxc/lxd/shared/api"
)

func reportContainers(containers []api.ContainerFull, format string, output string) {
	t := table.NewWriter()

	if output != "os.Stdout" {
		f, _ := os.Create(output)
		t.SetOutputMirror(f)
	} else {
		t.SetOutputMirror(os.Stdout)
	}

	t.AppendHeader(table.Row{"NAME", "STATE", "IPV4", "MACADDR", "IMAGE"})

	for _, container := range containers {

		t.AppendRow([]interface{}{
			container.Name,
			container.Status,
			getContainerIp(container.Name),
			container.Config["volatile.lxdbr0.hwaddr"],
			container.Config["image.description"],
		},
		)

	}

	if format == "table" {
		t.Render()
	}

	if format == "csv" {
		t.RenderCSV()
	}
}
