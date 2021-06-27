package playground

import (
	"context"

        "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"playground_id": dataSourcePlaygroundId(),
		},
	}
}

func dataSourcePlaygroundId() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePlaygroundIdRead,
		Schema: map[string]*schema.Schema{
			"playground_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePlaygroundIdRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if err := d.Set("playground_id", "10ab8aca-30e8-4f98-9af7-d44355b33c9c"); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
