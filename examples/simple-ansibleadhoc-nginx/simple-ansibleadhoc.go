package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
)

func main() {

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
		//Connection: "127.0.0.1:8090",
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  "127.0.0.1:8090,",
		ModuleName: "ping",
		//ModuleName: "nginx",
		//ModulePath:
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		//StdoutCallback:    "oneline",
	}

	fmt.Println("======adhoc======")
	fmt.Println(adhoc.String())

	//return

	err := adhoc.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
