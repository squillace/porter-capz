package capz

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the capz mixin in porter.yaml
// mixins:
// - capz:
//	  clientVersion: "v0.0.0"

type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
}

// This is an example. Replace the following with whatever steps are needed to
// install required components into
// const dockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli

/*
RUN curl -L https://github.com/kubernetes-sigs/cluster-api/releases/download/v0.3.13/clusterctl-linux-amd64 -o clusterctl
RUN chmod +x ./clusterctl
RUN sudo mv ./clusterctl /usr/local/bin/clusterctl
RUN clusterctl version


*/ // `

const dockerfileLines = `RUN mv ./clusterctl /usr/bin/ && chmod +x /usr/bin/clusterctl`

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	// Create new Builder.
	var input BuildInput

	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	suppliedClientVersion := input.Config.ClientVersion

	if suppliedClientVersion != "v0.3.13" {
		m.ClientVersion = "v0.3.13" //suppliedClientVersion
	}

	fmt.Fprintln(m.Out, "RUN apt-get update && apt-get install curl -y")
	// Example of pulling and defining a client version for your mixin
	fmt.Fprintf(m.Out, "\nRUN curl -L https://github.com/kubernetes-sigs/cluster-api/releases/download/%s/clusterctl-linux-amd64 -o clusterctl\n", m.ClientVersion)

	fmt.Fprintln(m.Out, dockerfileLines)

	return nil
}
