package v0

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/business/gen"
	"github.com/metamatex/metamate/metactl/pkg/v0/business/sdk"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate sdks using the asg",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		errs := func() (errs []error) {
			c, err := requireProjectConfig(gArgs)
			if err != nil {
				errs = append(errs, err)

				return
			}

			errs = gen.Gen(d.MessageReport, d.FileSystem, d.Version, d.RootNode, c.V0.Gen.Tasks)
			if len(errs) != 0 {
				return
			}

			var sdkNames []string
			for _, sdk0 := range c.V0.Gen.Sdks {
				sdkNames = append(sdkNames, sdk0.Names...)
			}

			errs = sdk.Reset(sdkNames)
			if len(errs) != 0 {
				return
			}

			for _, sdk0 := range c.V0.Gen.Sdks {
				for _, name := range sdk0.Names {
					errs = sdk.Gen(d.MessageReport, d.FileSystem, d.Version, d.RootNode, name, sdk0.Data, sdk0.Endpoints, nil)
					if len(errs) != 0 {
						return
					}
				}
			}

			return
		}()
		if len(errs) != 0 {
			d.MessageReport.AddError(errs)
		}

		handleReport(*d.MessageReport, types.Output{}, gArgs.VerbosityLevel)

		return
	},
}

func addGen(parentCmd *cobra.Command) {
	parentCmd.AddCommand(genCmd)
}
