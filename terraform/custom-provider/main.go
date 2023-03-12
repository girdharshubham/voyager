package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os/exec"
)

func absolve(_ *schema.ResourceData, i interface{}) error {
	return nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"executor": execute(),
		},
	}
}

func execute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entrypoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"command": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: func(data *schema.ResourceData, i interface{}) error {
			entrypoint := data.Get("entrypoint").(string)
			cmd := data.Get("command").(string)

			out, err := exec.Command(entrypoint, cmd).Output()
			data.SetId(string(out))
			return err
		},
		Read:   absolve,
		Delete: absolve,
	}
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}
