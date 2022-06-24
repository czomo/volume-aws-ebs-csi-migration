package main

import (
	"github.com/czomo/volume-aws-ebs-csi-migration/cmd/plugin/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // required for GKE
)

func main() {
	cli.InitAndExecute()
}
