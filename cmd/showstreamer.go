package cmd

import (
	"html/template"
	"text/tabwriter"

	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

const (
	dsShowFormat = "" +
		"{{range .}}" +
		"Name:\t{{.Name}}\n" +
		"State:\t{{.State}}\n" +
		"Address:\t{{.PrimaryIPV4}}\n" +
		"Management Network:\t{{.PrintMgmtNetwork}}\n" +
		"Cluster:\t{{.PrimaryClusterName}}\n\n" +
		"OS Deployment Plan:\n" +
		"{{range .PrintDeploymentPlan}}" +
		"Name:\t{{.}}\n" +
		"{{end}}" +
		"{{end}}"

	saShowFormat = "" +
		"Name:\t{{range .Name}}{{.}}\t{{end}}\n" +
		"-----\t-----------------\t-----------------\n" +
		"MgmtIP:\t{{range .MgmtIP}}{{.}}\t{{end}}\n" +
		"VSAMgmtIP:\t{{range .VSAMgmtIP}}{{.}}\t{{end}}\n" +
		"AMVMMgmtIP:\t{{range .AMVMMgmtIP}}{{.}}\t{{end}}\n" +
		"*MgmtClusterIP:\t{{range .ClusterIP}}{{.}}\t{{end}}\n" +
		"DataIP:\t{{range .DataIP}}{{.}}\t{{end}}\n" +
		"VSADataIP:\t{{range .VSADataIP}}{{.}}\t{{end}}\n" +
		"*VSAClusterIP:\t{{range .VSAClusterIP}}{{.}}\t{{end}}\n" +
		"AMVMDataIP:\t{{range .AMVMDataIP}}{{.}}\t{{end}}\n" +
		"MgmtActive:\t{{range .MgmtActive}}{{.}}\t{{end}}\n"

	artShowFormat = "" +
		"Name\tNum of DeploymentPlans\tNum of BuildPlans\tNum of PlanScripts\tNum of GoldenImages\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{len .DeploymentPlans}}\t{{len .BuildPlans}}\t{{len .PlanScripts}}\t{{len .Goldenimage}}\n" +
		"{{end}}"

	sdShowFormat = "" +
		"Name\tBuild Plan\tGolden Image\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.PrintBuildPlan}}\t{{.PrintGoldenImage}}\n" +
		"{{end}}"
)

func NewShowDeploymentServerCmd(c *oneview.CLIOVClient) *cobra.Command {

	//var name string

	var cmd = &cobra.Command{
		Use:   "os-deploy-server",
		Short: "show OS deployment server",
		Long:  `show OS deployment server`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var list []oneview.DeploymentServer
			var showFormat string

			list = c.GetDeploymentServer()
			showFormat = dsShowFormat

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()

			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, list)

		},
	}

	//cmd.Flags().StringVarP(&name, "name", "n", "", "Server name: all, <name>")

	return cmd
}

func NewShowStreamerCmd(c *oneview.CLIOVClient) *cobra.Command {

	//var name string

	var cmd = &cobra.Command{
		Use:   "streamer",
		Short: "show image streamer",
		Long:  `show image streamer`,

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(StreamerShowApplianceCmd(c))
	cmd.AddCommand(StreamerShowArtifactCmd(c))
	cmd.AddCommand(StreamerShowDeploymentPlanCmd(c))

	return cmd
}

func StreamerShowArtifactCmd(c *oneview.CLIOVClient) *cobra.Command {

	var name string

	var cmd = &cobra.Command{
		Use:   "artifact",
		Short: "show image streamer artifact",
		Long:  `show image streamer artifact`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var list []oneview.Artifact
			var showFormat string

			if name != "" {
				// list = c.GetServerHWVerbose(name)
				// showFormat = hwShowFormatVerbose

			} else {
				list = c.StreamerGetArtifact()
				showFormat = artShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()
			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, list)

		},
	}

	return cmd
}

func StreamerShowApplianceCmd(c *oneview.CLIOVClient) *cobra.Command {

	var name string

	var cmd = &cobra.Command{
		Use:   "appliance",
		Short: "show image streamer appliance",
		Long:  `show image streamer appliance`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var list oneview.ApplianceComparison
			var showFormat string

			if name != "" {
				// list = c.GetServerHWVerbose(name)
				// showFormat = hwShowFormatVerbose

			} else {
				list = c.GetStreamer()
				showFormat = saShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()
			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, list)

		},
	}
	return cmd
}

func StreamerShowDeploymentPlanCmd(c *oneview.CLIOVClient) *cobra.Command {

	var name string

	var cmd = &cobra.Command{
		Use:   "deployplan",
		Short: "show image streamer deployment plan",
		Long:  `show image streamer deployment plan`,
		Run: func(cmd *cobra.Command, args []string) {

			c := verifyClient(c)

			var list []oneview.StreamerDeploymentPlan
			var showFormat string

			if name != "" {
				// list = c.GetServerHWVerbose(name)
				// showFormat = hwShowFormatVerbose

			} else {
				list = c.StreamerGetDeploymentPlan()
				showFormat = sdShowFormat

			}

			tw := tabwriter.NewWriter(c.Out, 5, 1, 3, ' ', 0)
			defer tw.Flush()
			t := template.Must(template.New("").Parse(showFormat))
			t.Execute(tw, list)

		},
	}

	return cmd
}
