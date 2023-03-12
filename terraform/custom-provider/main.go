package main

import (
	"github.com/google/uuid"
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
			"executor": executor(),
		},
	}
}

func executor() *schema.Resource {

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

			_, err := exec.Command(entrypoint, cmd).Output()
			data.SetId(uuid.New().String())

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
