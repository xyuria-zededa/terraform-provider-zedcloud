package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func SystemInterfaceModel(d *schema.ResourceData) *models.SysInterface {
	costInt, _ := d.Get("cost").(int)
	cost := int64(costInt)
	var intfUsage *models.AdapterUsage // AdapterUsage
	intfUsageInterface, intfUsageIsSet := d.GetOk("intf_usage")
	if intfUsageIsSet {
		intfUsageModel := intfUsageInterface.(string)
		intfUsage = models.NewAdapterUsage(models.AdapterUsage(intfUsageModel))
	}
	intfname, _ := d.Get("intfname").(string)
	ipaddr, _ := d.Get("ipaddr").(string)
	macaddr, _ := d.Get("macaddr").(string)
	netname, _ := d.Get("netname").(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	return &models.SysInterface{
		Cost:      cost,
		IntfUsage: intfUsage,
		Intfname:  intfname,
		Ipaddr:    ipaddr,
		Macaddr:   macaddr,
		Netname:   netname,
		Tags:      tags,
	}
}

func SystemInterfaceModelFromMap(m map[string]interface{}) *models.SysInterface {
	cost := int64(m["cost"].(int))        // int64 false false false
	intfUsage := m["intf_usage"].(string) // AdapterUsage
	intfname := m["intfname"].(string)
	ipaddr := m["ipaddr"].(string)
	macaddr := m["macaddr"].(string)
	netname := m["netname"].(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}
	return &models.SysInterface{
		Cost:      cost,
		IntfUsage: models.NewAdapterUsage(models.AdapterUsage(intfUsage)),
		Intfname:  intfname,
		Ipaddr:    ipaddr,
		Macaddr:   macaddr,
		Netname:   netname,
		Tags:      tags,
	}
}

func SetSysInterfaceResourceData(d *schema.ResourceData, m *models.SysInterface) {
	d.Set("cost", m.Cost)
	d.Set("intf_usage", m.IntfUsage)
	d.Set("intfname", m.Intfname)
	d.Set("ipaddr", m.Ipaddr)
	d.Set("macaddr", m.Macaddr)
	d.Set("netname", m.Netname)
	d.Set("tags", m.Tags)
}

func SetSysInterfaceSubResourceData(m []*models.SysInterface) (d []*map[string]interface{}) {
	for _, SysInterfaceModel := range m {
		if SysInterfaceModel != nil {
			properties := make(map[string]interface{})
			properties["cost"] = SysInterfaceModel.Cost
			properties["intf_usage"] = SysInterfaceModel.IntfUsage
			properties["intfname"] = SysInterfaceModel.Intfname
			properties["ipaddr"] = SysInterfaceModel.Ipaddr
			properties["macaddr"] = SysInterfaceModel.Macaddr
			properties["netname"] = SysInterfaceModel.Netname
			properties["tags"] = SysInterfaceModel.Tags
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysInterface resource defined in the Terraform configuration
func SystemInterface() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cost": {
			Description: `cost of using this interface. Default is 0.`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"intf_usage": {
			Description: `Adapter Udage`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"intfname": {
			Description: `name of interface in the manifest to which this network or adapter maps to`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ipaddr": {
			Description: `IP address: we will be needing this in cae of static network`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"macaddr": {
			Description: `mac address needs to be over-written in some cases`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"netname": {
			Description: `network name: if attaching a network use netname`,
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
	}
}

// Retrieve property field names for updating the SysInterface resource
func GetSysInterfacePropertyFields() (t []string) {
	return []string{
		"cost",
		"intf_usage",
		"intfname",
		"ipaddr",
		"macaddr",
		"netname",
		"tags",
	}
}