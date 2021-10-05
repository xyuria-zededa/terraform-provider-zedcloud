// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
)

var EdgeNodeDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceEdgeNode,
	Schema:      zschemas.EdgeNodeSchema,
	Description: "Schema for data source zedcloud_edgenode. Must specify id or name",
}

// The schema for a resource group data source
func getEdgeNodeDataSourceSchema() *schema.Resource {
	// Use the same schema for data source as well.
	return EdgeNodeDataSourceSchema
}

func flattenGeoLocation(entry interface{}) []interface{} {
	loc := entry.(*swagger_models.GeoLocation)
	if loc == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"city":        loc.City,
		"country":     loc.Country,
		"freeloc":     loc.Freeloc,
		"hostname":    loc.Hostname,
		"loc":         loc.Loc,
		"org":         loc.Org,
		"postal":      loc.Postal,
		"region":      loc.Region,
		"underlay_ip": loc.UnderlayIP,
	}}
}

func flattenSysInterfaces(cfgList []*swagger_models.SysInterface) []interface{} {
	entryList := make([]interface{}, 0)
	for _, intf := range cfgList {
		entryList = append(entryList, map[string]interface{}{
			"cost":       int(intf.Cost),
			"intfname":   intf.Intfname,
			"intf_usage": ptrValStr(intf.IntfUsage),
			"ipaddr":     intf.Ipaddr,
			"macaddr":    intf.Macaddr,
			"netname":    intf.Netname,
			"tags":       flattenStringMap(intf.Tags),
		})
	}
	return entryList
}

func flattenEDConfigItems(cfgList []*swagger_models.EDConfigItem) map[string]interface{} {
	dataMap := make(map[string]interface{})
	for _, item := range cfgList {
		dataMap[item.Key] = item.StringValue
	}
	return dataMap
}

func flattenDeviceConfig(cfg *swagger_models.DeviceConfig, computedOnly bool) map[string]interface{} {
	eveImageVersion := ""
	for _, image := range cfg.BaseImage {
		if image.Activate {
			eveImageVersion = *image.ImageName
			break
		}
	}

	data := map[string]interface{}{
		"adminstate":    ptrValStr(cfg.AdminState),
		"cluster_id":    cfg.ClusterID,
		"cpu":           int(cfg.CPU),
		"id":            cfg.ID,
		"memory":        int(cfg.Memory),
		"model_id":      ptrValStr(cfg.ModelID),
		"project_id":    ptrValStr(cfg.ProjectID),
		"reset_counter": int(cfg.ResetCounter),
		"reset_time":    cfg.ResetTime,
		"revision":      flattenObjectRevision(cfg.Revision),
		"serialno":      cfg.Serialno,
		"storage":       int(cfg.Storage),
		"thread":        int(cfg.Storage),
		"utype":         ptrValStr(cfg.Utype),
	}
	if !computedOnly {
		data["asset_id"] = cfg.AssetID
		data["client_ip"] = cfg.ClientIP
		data["config_items"] = flattenEDConfigItems(cfg.ConfigItem)
		data["description"] = cfg.Description
		data["dev_location"] = flattenGeoLocation(cfg.DevLocation)
		data["eve_image_version"] = eveImageVersion
		data["interface"] = flattenSysInterfaces(cfg.Interfaces)
		data["name"] = ptrValStr(cfg.Name)
		data["project_id"] = ptrValStr(cfg.ProjectID)
		data["tags"] = flattenStringMap(cfg.Tags)
		data["title"] = ptrValStr(cfg.Title)
	}
	flattenedDataCheckKeys(zschemas.EdgeNodeSchema, data, computedOnly)
	return data
}

// Read the Resource Group
func readEdgeNode(ctx context.Context, d *schema.ResourceData,
	meta interface{}, resource bool) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
	}
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	cfg, err := getEdgeNodeConfig(client, name, id)
	if err != nil {
		log.Printf("[ERROR] Failed to read config. Error: %s", err.Error())
		return diag.Errorf("[ERROR] edge node %s ( id: %s) not found. Err: %s",
			name, id, err.Error())
	}
	// Take the Config and convert it to terraform object
	marshalData(d, flattenDeviceConfig(cfg, resource))
	return diags
}

func readDataSourceEdgeNode(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readEdgeNode(ctx, d, meta, true)
}
