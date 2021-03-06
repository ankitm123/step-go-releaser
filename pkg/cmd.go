package pkg

import (
	"github.com/jenkins-x/jx-helpers/pkg/cmdrunner"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

type options struct {
	Cmd           *cobra.Command
	Args          []string
	Runner        cmdrunner.Command
	CommandRunner cmdrunner.CommandRunner
	KubeClient    kubernetes.Interface
	Namespace     string

	organisation string
	revision     string
	branch       string
	buildDate    string
	goVersion    string
	version      string
	rootPackage  string
	timeout      string
}

const (
	version      = "version"
	organisation = "organisation"
	revision     = "revision"
	branch       = "branch"
	buildDate    = "build-date"
	goVersion    = "go-version"
	rootPackage  = "root-package"
)

var (
	createLong = `
wraps goreleaser so we can get the git token from a kubernetes secret
`

	createExample = `
step-goreleaser --organisation=jenkins-x-labs --revision=1b59ffc --branch=master --build-date=20200303-22:14:54 --go-version=1.12.17 --root-package=github.com/jenkins-x-labs/gsm-controller --version=0.0.17
`
)

// NewCmdGoReleaser creates a command object for the "hello world" command
func NewCmdGoReleaser() *cobra.Command {
	o := &options{}

	cmd := &cobra.Command{
		Use:     "step-goreleaser",
		Short:   "wraps the go releaser tool getting required secrets needed to upload artifacts to GitHub",
		Long:    createLong,
		Example: createExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.Args = args
			return o.Run()
		},
	}

	cmd.Flags().StringVarP(&o.organisation, organisation, "", "", "the git organisation")
	cmd.Flags().StringVarP(&o.revision, revision, "", "", "git revision")
	cmd.Flags().StringVarP(&o.branch, branch, "", "", "git branch")
	cmd.Flags().StringVarP(&o.buildDate, buildDate, "", "", "build date")
	cmd.Flags().StringVarP(&o.version, version, "", "", "version")
	cmd.Flags().StringVarP(&o.goVersion, goVersion, "", "", "go version")
	cmd.Flags().StringVarP(&o.rootPackage, rootPackage, "", "", "root package")
	cmd.Flags().StringVarP(&o.timeout, "timeout", "", "200m", "timeout")
	o.Cmd = cmd

	return cmd
}
