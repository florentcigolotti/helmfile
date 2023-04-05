package cmd

import (
	"github.com/spf13/cobra"

	"github.com/helmfile/helmfile/pkg/app"
	"github.com/helmfile/helmfile/pkg/config"
)

// NewDagCmd returns dag subcmd
func NewDagCmd(globalCfg *config.GlobalImpl) *cobra.Command {
	dagOptions := config.NewDagOptions()

	cmd := &cobra.Command{
		Use:   "dag",
		Short: "Generate a DAG of releases",
		RunE: func(cmd *cobra.Command, args []string) error {
			dagImpl := config.NewDagImpl(globalCfg, dagOptions)
			err := config.NewCLIConfigImpl(dagImpl.GlobalImpl)
			if err != nil {
				return err
			}

			if err := dagImpl.ValidateConfig(); err != nil {
				return err
			}

			a := app.New(dagImpl)
			return toCLIError(dagImpl.GlobalImpl, a.Dag(dagImpl))
		},
	}

	f := cmd.Flags()
	f.StringVar(&globalCfg.GlobalOptions.Args, "args", "", "pass args to helm diff")
	f.StringArrayVar(&dagOptions.Set, "set", nil, "additional values to be merged into the command")
	f.StringArrayVar(&dagOptions.Values, "values", nil, "additional value files to be merged into the command")
	f.IntVar(&dagOptions.Concurrency, "concurrency", 0, "maximum number of concurrent helm processes to run, 0 is unlimited")
	f.BoolVar(&dagOptions.Validate, "validate", false, "validate your manifests against the Kubernetes cluster you are currently pointing at. Note that this requires access to a Kubernetes cluster to obtain information necessary for validating, like the diff of available API versions")
	f.BoolVar(&dagOptions.SkipNeeds, "skip-needs", true, `do not automatically include releases from the target release's "needs" when --selector/-l flag is provided. Does nothing when --selector/-l flag is not provided. Defaults to true when --include-needs or --include-transitive-needs is not provided`)
	f.BoolVar(&dagOptions.IncludeTests, "include-tests", false, "enable the diffing of the helm test hooks")
	f.BoolVar(&dagOptions.IncludeNeeds, "include-needs", false, `automatically include releases from the target release's "needs" when --selector/-l flag is provided. Does nothing when --selector/-l flag is not provided`)
	f.BoolVar(&dagOptions.IncludeTransitiveNeeds, "include-transitive-needs", false, `like --include-needs, but also includes transitive needs (needs of needs). Does nothing when --selector/-l flag is not provided. Overrides exclusions of other selectors and conditions.`)
	f.BoolVar(&dagOptions.SkipDeps, "skip-deps", false, `skip running "helm repo update" and "helm dependency build"`)
	f.BoolVar(&dagOptions.ShowSecrets, "show-secrets", false, "do not redact secret values in the output. should be used for debug purpose only")
	f.BoolVar(&dagOptions.NoHooks, "no-hooks", false, "do not diff changes made by hooks.")
	f.BoolVar(&dagOptions.DetailedExitcode, "detailed-exitcode", false, "return a detailed exit code")
	f.IntVar(&dagOptions.Context, "context", 0, "output NUM lines of context around changes")
	f.StringVar(&dagOptions.Output, "output", "", "output format for diff plugin")
	f.BoolVar(&dagOptions.SuppressSecrets, "suppress-secrets", false, "suppress secrets in the output. highly recommended to specify on CI/CD use-cases")
	f.StringArrayVar(&dagOptions.Suppress, "suppress", nil, "suppress specified Kubernetes objects in the output. Can be provided multiple times. For example: --suppress KeycloakClient --suppress VaultSecret")
	f.BoolVar(&dagOptions.ReuseValues, "reuse-values", false, `Override helmDefaults.reuseValues "helm diff upgrade --install --reuse-values"`)
	f.BoolVar(&dagOptions.ResetValues, "reset-values", false, `Override helmDefaults.reuseValues "helm diff upgrade --install --reset-values"`)
	f.StringVar(&dagOptions.PostRenderer, "post-renderer", "", `pass --post-renderer to "helm template" or "helm upgrade --install"`)

	return cmd
}
