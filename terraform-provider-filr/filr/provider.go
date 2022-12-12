package filr

import (
	"log"
	"os"
	"path"

	"github.com/hashicorp/go-uuid"
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
		Create: resourceCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func getFilename(d *schema.ResourceData, m interface{}) string {
	folder := m.(File).Folder
	return path.Join(folder, d.Id())
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	id, err := uuid.GenerateUUID()
	if err != nil {
		log.Print(err)
	}
	d.SetId(id)
	content := d.Get("content").(string)
	os.WriteFile(getFilename(d, m), []byte(content), 0644)
	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	fileName := getFilename(d, m)
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Print(err)
	}
	d.Set("content", string(data))
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	content := d.Get("content").(string)
	os.WriteFile(getFilename(d, m), []byte(content), 0644)
	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	fileName := getFilename(d, m)
	os.Remove(fileName)
	return nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	folder := d.Get("folder").(string)
	return File{Folder: folder}, nil
}
