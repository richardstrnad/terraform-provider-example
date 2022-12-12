package filr

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type File struct {
	Folder string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"folder": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TERRA_FOLDER", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"filr_file_state": resourceFile(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "the content of the file",
			},
		},
		CreateContext: resourceCreateItem,
		ReadContext:   resourceReadItem,
		UpdateContext: resourceUpdateItem,
		DeleteContext: resourceDeleteItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func getFilename(d *schema.ResourceData, m interface{}) string {
	folder := m.(File).Folder
	return path.Join(folder, d.Id())
}

func resourceCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id, err := uuid.GenerateUUID()
	if err != nil {
		tflog.Error(ctx, "Couldn't generate UUID")
		return diag.FromErr(err)
	}
	d.SetId(id)
	content := d.Get("content").(string)
	err = os.WriteFile(getFilename(d, m), []byte(content), 0644)
	if err != nil {
		tflog.Error(ctx, "Couldn't write file")
		return diag.FromErr(err)
	}
	return nil
}

func resourceReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fileName := getFilename(d, m)
	data, err := os.ReadFile(fileName)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Couldn't read file: %s", fileName))
		return diag.FromErr(err)
	}
	d.Set("content", string(data))
	return nil
}

func resourceUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	content := d.Get("content").(string)
	err := os.WriteFile(getFilename(d, m), []byte(content), 0644)
	if err != nil {
		tflog.Error(ctx, "Couldn't write the file")
		return diag.FromErr(err)
	}
	return nil
}

func resourceDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fileName := getFilename(d, m)
	err := os.Remove(fileName)
	if err != nil {
		tflog.Error(ctx, "Couldn't delete the file")
		return diag.FromErr(err)
	}
	return nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	folder := d.Get("folder").(string)
	return File{Folder: folder}, nil
}
