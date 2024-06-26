# terraform-provider-zedcloud

The __terraform-provider-zedcloud__ provides access to Zededa's public API for the zedcloud cloud services.

Supported resources are:
- [x] Edge-Node (also called Device-Config)
- [x] Edge-App
- [x] Edge-App-Instance
- [x] Network
- [x] Network-Instance
- [x] Datastore
- [x] Volume-Instance
- [ ] Image

## Documentation

### Schemas

Documentation of all API endpoints and data schemas can be found under https://zedcontrol.zededa.net/api/v1/docs/.

Schema documentation generated from the schema files contained in this repo can be found under:
1. [Provider](https://github.com/xyuria-zededa/terraform-provider-zedcloud/blob/main/docs/index.md)
1. [Resources](https://github.com/xyuria-zededa/terraform-provider-zedcloud/tree/main/docs/resources)
1. [Data Sources](https://github.com/xyuria-zededa/terraform-provider-zedcloud/tree/main/docs/data-sources)

> Note, the resource schemas match the API schemas but support for some fields might be incomplete in the provider. The testdata directory contains examples with the full set of supported fields for the supported resources of the latest version deployed in the Terraform provider registry.

---

### Product

Product documentation including explanations of workflows and data schemas can be found under https://help.zededa.com/hc/en-us

## Installation and configuration

Information on how to install and configure a Terraform provider can be found under https://developer.hashicorp.com/terraform/language/providers. If you need help setting up the __terraform-provider-zedcloud__ please reach out to Zededa support.
The latest version of the provider can be found in the official Terraform provider registry under https://registry.terraform.io/providers/xyuria-zededa/zedcloud/latest.

## Breaking changes in v2

- Authentication uses the API-token only. Basic-Auth via username and password has been removed. You can find the API-token for your user under https://zedcontrol.zededa.net/profile/user
- The resource and data schemas have changed. The schemas now map to the ones in API documentation. For reference on supported fields see the testdata directory which contains Terraform configuration files for all supported resources. If you have trouble porting your configuration, please reach out to Zededa support.
- `adminstate_config` has been removed. `admin_state` can now be configured directly. Diffs will be suppressed if the `EdgeNode` is in `ADMINSTATE_REGISTERED` state and your configuration is set to `ADMINSTATE_ACTIVE`.
- CSRF checks have been removed.

## Environment configuration

The following environment variables need to be set to use the provider.
```
export TF_VAR_zedcloud_url="zedcontrol.zededa.net"
export TF_VAR_zedcloud_token=<YOUR-API-TOKEN>
export TF_LOG=ERROR
export TF_LOG_PATH=./terraform.log
```

