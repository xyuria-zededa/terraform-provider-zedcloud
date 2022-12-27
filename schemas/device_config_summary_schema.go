package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceConfigSummary resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceConfigSummaryModel(d *schema.ResourceData) *models.DeviceConfigSummary {
	adminState := d.Get("admin_state").(*models.AdminState)  // AdminState
	baseImage := d.Get("base_image").([]*models.BaseOSImage) // []*BaseOSImage
	clusterID := d.Get("cluster_id").(string)
	var debugKnob *models.DebugKnobDetail // DebugKnobDetail
	debugKnobInterface, debugKnobIsSet := d.GetOk("debug_knob")
	if debugKnobIsSet {
		debugKnobMap := debugKnobInterface.([]interface{})[0].(map[string]interface{})
		debugKnob = DebugKnobDetailModelFromMap(debugKnobMap)
	}
	deploymentTag := d.Get("deployment_tag").(string)
	description := d.Get("description").(string)
	id := d.Get("id").(string)
	interfaces := d.Get("interfaces").([]*models.SysInterface) // []*SysInterface
	modelID := d.Get("model_id").(string)
	name := d.Get("name").(string)
	projectID := d.Get("project_id").(string)
	serialno := d.Get("serialno").(string)
	tags := d.Get("tags").(map[string]string) // map[string]string
	title := d.Get("title").(string)
	utype := d.Get("utype").(*models.ModelArchType) // ModelArchType
	return &models.DeviceConfigSummary{
		AdminState:    adminState,
		BaseImage:     baseImage,
		ClusterID:     clusterID,
		DebugKnob:     debugKnob,
		DeploymentTag: deploymentTag,
		Description:   description,
		ID:            id,
		Interfaces:    interfaces,
		ModelID:       &modelID,   // string true false false
		Name:          &name,      // string true false false
		ProjectID:     &projectID, // string true false false
		Serialno:      serialno,
		Tags:          tags,
		Title:         &title, // string true false false
		Utype:         utype,
	}
}

func DeviceConfigSummaryModelFromMap(m map[string]interface{}) *models.DeviceConfigSummary {
	adminState := m["admin_state"].(*models.AdminState)  // AdminState
	baseImage := m["base_image"].([]*models.BaseOSImage) // []*BaseOSImage
	clusterID := m["cluster_id"].(string)
	var debugKnob *models.DebugKnobDetail // DebugKnobDetail
	debugKnobInterface, debugKnobIsSet := m["debug_knob"]
	if debugKnobIsSet {
		debugKnobMap := debugKnobInterface.([]interface{})[0].(map[string]interface{})
		debugKnob = DebugKnobDetailModelFromMap(debugKnobMap)
	}
	//
	deploymentTag := m["deployment_tag"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	interfaces := m["interfaces"].([]*models.SysInterface) // []*SysInterface
	modelID := m["model_id"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	serialno := m["serialno"].(string)
	tags := m["tags"].(map[string]string)
	title := m["title"].(string)
	utype := m["utype"].(*models.ModelArchType) // ModelArchType
	return &models.DeviceConfigSummary{
		AdminState:    adminState,
		BaseImage:     baseImage,
		ClusterID:     clusterID,
		DebugKnob:     debugKnob,
		DeploymentTag: deploymentTag,
		Description:   description,
		ID:            id,
		Interfaces:    interfaces,
		ModelID:       &modelID,
		Name:          &name,
		ProjectID:     &projectID,
		Serialno:      serialno,
		Tags:          tags,
		Title:         &title,
		Utype:         utype,
	}
}

// Update the underlying DeviceConfigSummary resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceConfigSummaryResourceData(d *schema.ResourceData, m *models.DeviceConfigSummary) {
	d.Set("admin_state", m.AdminState)
	d.Set("base_image", SetBaseOSImageSubResourceData(m.BaseImage))
	d.Set("cluster_id", m.ClusterID)
	d.Set("debug_knob", SetDebugKnobDetailSubResourceData([]*models.DebugKnobDetail{m.DebugKnob}))
	d.Set("deployment_tag", m.DeploymentTag)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("interfaces", SetSysInterfaceSubResourceData(m.Interfaces))
	d.Set("model_id", m.ModelID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("serialno", m.Serialno)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("utype", m.Utype)
}

// Iterate throught and update the DeviceConfigSummary resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceConfigSummarySubResourceData(m []*models.DeviceConfigSummary) (d []*map[string]interface{}) {
	for _, DeviceConfigSummaryModel := range m {
		if DeviceConfigSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["admin_state"] = DeviceConfigSummaryModel.AdminState
			properties["base_image"] = SetBaseOSImageSubResourceData(DeviceConfigSummaryModel.BaseImage)
			properties["cluster_id"] = DeviceConfigSummaryModel.ClusterID
			properties["debug_knob"] = SetDebugKnobDetailSubResourceData([]*models.DebugKnobDetail{DeviceConfigSummaryModel.DebugKnob})
			properties["deployment_tag"] = DeviceConfigSummaryModel.DeploymentTag
			properties["description"] = DeviceConfigSummaryModel.Description
			properties["id"] = DeviceConfigSummaryModel.ID
			properties["interfaces"] = SetSysInterfaceSubResourceData(DeviceConfigSummaryModel.Interfaces)
			properties["model_id"] = DeviceConfigSummaryModel.ModelID
			properties["name"] = DeviceConfigSummaryModel.Name
			properties["project_id"] = DeviceConfigSummaryModel.ProjectID
			properties["serialno"] = DeviceConfigSummaryModel.Serialno
			properties["tags"] = DeviceConfigSummaryModel.Tags
			properties["title"] = DeviceConfigSummaryModel.Title
			properties["utype"] = DeviceConfigSummaryModel.Utype
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceConfigSummary resource defined in the Terraform configuration
func DeviceConfigSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_state": {
			Description: `administrative state of device`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"base_image": {
			Description: `base images`,
			Type:        schema.TypeList, //GoType: []*BaseOSImage
			Elem: &schema.Resource{
				Schema: BaseOSImageSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"cluster_id": {
			Description: `System defined universally unique clusterInstance ID, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"debug_knob": {
			Description: `debug knob details for the device`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"deployment_tag": {
			Description: `user defined tag for the device, which is used while deploying policies.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `user specified description`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `system generated unique id for a device`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"interfaces": {
			Description: `System Interface list`,
			Type:        schema.TypeList, //GoType: []*SysInterface
			Elem: &schema.Resource{
				Schema: SysInterfaceSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"model_id": {
			Description: `device model`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"name": {
			Description: `user specified device name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `project name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"serialno": {
			Description: `Device serial number`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `user specified title`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"utype": {
			Description: `device model arch type`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the DeviceConfigSummary resource
func GetDeviceConfigSummaryPropertyFields() (t []string) {
	return []string{
		"admin_state",
		"base_image",
		"cluster_id",
		"debug_knob",
		"deployment_tag",
		"description",
		"id",
		"interfaces",
		"model_id",
		"name",
		"project_id",
		"serialno",
		"tags",
		"title",
		"utype",
	}
}
