package atlas

import (
	"fmt"
	"log"
	
	"github.com/hashicorp/atlas-go/v1"
	"github.com/hashicorp/terraform/helper/schema"
)

/*

resource "atlas_workspace" "foo" {
  name = "megacorp/coolworkspace"
  
}
*/

func resourceWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkspaceCreate,
		Read:   resourceWorkspaceRead,
		Update: resourceWorkspaceUpdate,
	}
}

func resourceWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*atlas.Client)

	user, name, err := atlas.ParseSlug(d.Get("name").(string))
	if err != nil {
		return err
	}
	
	app, err := client.App(user, name)
	if err != nil {
		// TODO: We should mark the resource as nonexistent
		// here instead if this is a 404
		return fmt.Errorf("Error fetching workspace %s/%s: %s", user, name, err)
	}

	d.SetId(app.Slug()) // TODO: is this good as an ID? Is it
			    // unique? Should we include the Atlas URL
			    // in it?

	d.Set("user", app.User)
	d.Set("workspace", app.Name)
	return nil
}

func resourceWorkspaceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*atlas.Client)

	user, name, err := atlas.ParseSlug(d.Get("name").(string))
	if err != nil {
		return err
	}

	app, err := client.CreateApp(user, name)
	if err != nil {
		return fmt.Errorf("Error creating workspace %s/%s: %s", user, name, err)
	}
	err := resourceWorkspaceUpdate(d, meta)
	if err != nil {
		return err
	}
	return resourceWorkspaceRead(d, meta)
}

func resourceWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*atlas.Client)

	user, name, err := atlas.ParseSlug(d.Get("name").(string))
	if err != nil {
		return err
	}
	
}
