package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"os"
	"strings"
)

type PluginInterfacePlugin struct{}

func (c *PluginInterfacePlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "plugin-interface-plugin" {
		c.CliCommandWithoutTerminalOutput(cliConnection)
		c.CliCommand(cliConnection)
		c.GetCurrentOrg(cliConnection)
		c.GetCurrentSpace(cliConnection)
		c.Username(cliConnection)
		c.UserGuid(cliConnection)
		c.UserEmail(cliConnection)
		c.IsLoggedIn(cliConnection)
		c.IsSSLDisabled(cliConnection)
		c.HasOrganization(cliConnection)
		c.HasSpace(cliConnection)
		c.ApiEndpoint(cliConnection)
		c.ApiVersion(cliConnection)
		c.HasAPIEndpoint(cliConnection)
		c.LoggregatorEndpoint(cliConnection)
		c.DopplerEndpoint(cliConnection)
		c.AccessToken(cliConnection)
		c.GetApp(cliConnection)
		c.GetApps(cliConnection)
		c.GetOrgs(cliConnection)
		c.GetSpaces(cliConnection)
		c.GetOrgUsers(cliConnection)
		c.GetSpaceUsers(cliConnection)
		c.GetServices(cliConnection)
		c.GetService(cliConnection)
		c.GetOrg(cliConnection)
		c.GetSpace(cliConnection)
	}
}

func (c *PluginInterfacePlugin) CliCommandWithoutTerminalOutput(cliConnection plugin.CliConnection) {
	wrapFunction("CliCommandWithoutTerminalOutput", func() (string, error) {
		CliCommandWithoutTerminalOutputArgs := []string{"feature-flag", "diego_docker"}
		resp, err := cliConnection.CliCommandWithoutTerminalOutput(CliCommandWithoutTerminalOutputArgs...)
		response := strings.Join(resp, "\n\t")
		return response, err
	})
}

func (c *PluginInterfacePlugin) CliCommand(cliConnection plugin.CliConnection) {
	wrapFunction("CliCommand", func() (string, error) {
		CliCommandArgs := []string{"feature-flag", "diego_docker"}
		resp, err := cliConnection.CliCommand(CliCommandArgs...)
		response := strings.Join(resp, "\n\t")
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetCurrentOrg(cliConnection plugin.CliConnection) {
	wrapFunction("GetCurrentOrg", func() (string, error) {
		org, err := cliConnection.GetCurrentOrg()
		response := fmt.Sprintf("%+v", org)
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetCurrentSpace(cliConnection plugin.CliConnection) {
	wrapFunction("GetCurrentSpace", func() (string, error) {
		space, err := cliConnection.GetCurrentSpace()
		response := fmt.Sprintf("%+v", space)
		return response, err
	})
}

func (c *PluginInterfacePlugin) Username(cliConnection plugin.CliConnection) {
	wrapFunction("Username", func() (string, error) {
	username, err := cliConnection.Username()
		return username, err
	})
}

func (c *PluginInterfacePlugin) UserGuid(cliConnection plugin.CliConnection) {
	wrapFunction("UserGuid", func() (string, error) {
		userGuid, err := cliConnection.UserGuid()
		return userGuid, err
	})
}

func (c *PluginInterfacePlugin) UserEmail(cliConnection plugin.CliConnection) {
	wrapFunction("UserEmail", func() (string, error) {
		username, err := cliConnection.UserEmail()
		return username, err
	})
}

func (c *PluginInterfacePlugin) IsLoggedIn(cliConnection plugin.CliConnection) {
	wrapFunction("IsLoggedIn", func() (string, error) {
		loggedIn, err := cliConnection.IsLoggedIn()
		response := fmt.Sprintf("%t", loggedIn)
		return response, err
	})
}

func (c *PluginInterfacePlugin) IsSSLDisabled(cliConnection plugin.CliConnection) {
	wrapFunction("IsSSLDisabled", func() (string, error) {
		sslDisabled, err := cliConnection.IsSSLDisabled()
		response := fmt.Sprintf("%t", sslDisabled)
		return response, err
	})
}

func (c *PluginInterfacePlugin) HasOrganization(cliConnection plugin.CliConnection) {
	wrapFunction("HasOrganization", func() (string, error) {
		hasOrganization, err := cliConnection.HasOrganization()
		response := fmt.Sprintf("%t", hasOrganization)
		return response, err
	})
}

func (c *PluginInterfacePlugin) HasSpace(cliConnection plugin.CliConnection) {
	wrapFunction("HasSpace", func() (string, error) {
		hasSpace, err := cliConnection.HasSpace()
		response := fmt.Sprintf("%t", hasSpace)
		return response, err
	})
}

func (c *PluginInterfacePlugin) ApiEndpoint(cliConnection plugin.CliConnection) {
	wrapFunction("ApiEndpoint", func() (string, error) {
		apiEndpoint, err := cliConnection.ApiEndpoint()
		response := fmt.Sprintf("%s", apiEndpoint)
		return response, err
	})
}

func (c *PluginInterfacePlugin) ApiVersion(cliConnection plugin.CliConnection) {
	wrapFunction("ApiVersion", func() (string, error) {
		apiVersion, err := cliConnection.ApiVersion()
		return apiVersion, err
	})
}

func (c *PluginInterfacePlugin) HasAPIEndpoint(cliConnection plugin.CliConnection) {
	wrapFunction("HasAPIEndpoint", func() (string, error) {
		hasApiEndpoint, err := cliConnection.HasAPIEndpoint()
		response := fmt.Sprintf("%t", hasApiEndpoint)
		return response, err
	})
}

func (c *PluginInterfacePlugin) LoggregatorEndpoint(cliConnection plugin.CliConnection) {
	wrapFunction("LoggregatorEndpoint", func() (string, error) {
		loggregatorEndpoint, err := cliConnection.LoggregatorEndpoint()
		return loggregatorEndpoint, err
	})
}

func (c *PluginInterfacePlugin) DopplerEndpoint(cliConnection plugin.CliConnection) {
	wrapFunction("DopplerEndpoint", func() (string, error) {
		dopplerEndpoint, err := cliConnection.DopplerEndpoint()
		return dopplerEndpoint, err
	})
}

func (c *PluginInterfacePlugin) AccessToken(cliConnection plugin.CliConnection) {
	wrapFunction("AccessToken", func() (string, error) {
		accessToken, err := cliConnection.AccessToken()
		return accessToken, err
	})
}

func (c *PluginInterfacePlugin) GetApp(cliConnection plugin.CliConnection) {
	wrapFunction("GetApp", func() (string, error) {
		appName := "my-app"
		app, err := cliConnection.GetApp(appName)
		response := fmt.Sprintf("%+v", app)
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetApps(cliConnection plugin.CliConnection) {
	wrapFunction("GetApps", func() (string, error) {
		apps, err := cliConnection.GetApps()
		appStrings := stringifyStructs(apps)
		response := strings.Join(appStrings, ",\n")
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetOrgs(cliConnection plugin.CliConnection) {
	wrapFunction("GetOrgs", func() (string, error) {
		orgs, err := cliConnection.GetOrgs()
		orgStrings := stringifyStructs(orgs)
		response := strings.Join(orgStrings, ",\n")
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetSpaces(cliConnection plugin.CliConnection) {
	wrapFunction("GetSpaces", func() (string, error) {
		spaces, err := cliConnection.GetSpaces()
		spaceStrings := stringifyStructs(spaces)
		response := strings.Join(spaceStrings, ",\n")
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetOrgUsers(cliConnection plugin.CliConnection) {
	wrapFunction("GetOrgUsers", func() (string, error) {
		orgName := "org"
		orgUsers, err := cliConnection.GetOrgUsers(orgName)
		orgUserStrings := stringifyStructs(orgUsers)
		response := strings.Join(orgUserStrings, ",\n")
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetSpaceUsers(cliConnection plugin.CliConnection) {
	wrapFunction("GetSpaceUsers", func() (string, error) {
		orgName := "org"
		spaceName := "space"
		spaceUsers, err := cliConnection.GetSpaceUsers(orgName, spaceName)
		spaceUserStrings := stringifyStructs(spaceUsers)
		response := strings.Join(spaceUserStrings, ",\n")
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetServices(cliConnection plugin.CliConnection) {
	wrapFunction("GetServices", func() (string, error) {
		services, err := cliConnection.GetServices()
		serviceStrings := stringifyStructs(services)
		response := strings.Join(serviceStrings, ",\n")
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetService(cliConnection plugin.CliConnection) {
	wrapFunction("GetService", func() (string, error) {
		serviceName := "instance"
		service, err := cliConnection.GetService(serviceName)
		response := fmt.Sprintf("%+v", service)
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetOrg(cliConnection plugin.CliConnection) {
	wrapFunction("GetOrg", func() (string, error) {
		orgName := "org"
		org, err := cliConnection.GetOrg(orgName)
		response := fmt.Sprintf("%+v", org)
		return response, err
	})
}

func (c *PluginInterfacePlugin) GetSpace(cliConnection plugin.CliConnection) {
	wrapFunction("GetSpace", func() (string, error) {
		spaceName := "space"
		space, err := cliConnection.GetSpace(spaceName)
		response := fmt.Sprintf("%+v", space)
		return response, err
	})
}

func wrapFunction(funcName string, function func() (string, error)) {
	startCommand(funcName)
	response, err := function()
	deployTheAlgorithm(response, err)
	endCommand(funcName)
}

func stringifyStructs(structs ...interface{}) []string {
	var someStrings []string
	for _, aStruct := range structs {
		someStrings = append(someStrings, fmt.Sprintf("%+v", aStruct))
	}
	return someStrings
}

func deployTheAlgorithm(response string, err error) {
	if err != nil {
		handleErr(err)
	} else {
		handleResp(response)
	}
}

func startCommand(funcName string) () {
	fmt.Printf("%s\n\n", funcName)
}

func endCommand(funcName string) () {
	fmt.Printf("/%s\n\n", funcName)
}

func handleResp(resp string) {
	fmt.Printf("\tCommand Output:\n\t\t%s\n\n", resp)
}

func handleErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\tError:\n\t\t%+v\n\n", err)
}

func (c *PluginInterfacePlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "PluginInterfacePlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "plugin-interface-plugin",
				HelpText: "Plugin to exercise the plugin interface",
				UsageDetails: plugin.Usage{
					Usage: "plugin-interface-plugin\n   cf plugin-interface-plugin",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(PluginInterfacePlugin))
}
