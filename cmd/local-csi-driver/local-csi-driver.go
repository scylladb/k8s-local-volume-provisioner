// Copyright (C) 2021 ScyllaDB

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	cmd "github.com/scylladb/k8s-local-volume-provisioner/pkg/cmd/local-csi-driver"
	"github.com/scylladb/k8s-local-volume-provisioner/pkg/genericclioptions"
	"k8s.io/klog/v2"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	klog.InitFlags(flag.CommandLine)
	err := flag.Set("logtostderr", "true")
	if err != nil {
		panic(err)
	}
	defer klog.Flush()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	command := cmd.NewLocalDriverCommand(streams)
	err = command.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
