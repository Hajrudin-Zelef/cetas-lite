---
id: collect-261001-fortinet/fortinet/providers-fortinetdev-fortios-latest-docs-be6bb4cd-1
title: "Configure the FortiOS Provider for FortiGate"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/providers-fortinetdev-fortios-latest-docs-be6bb4cd.md
source_anchor: ""
source_lines: [1, 184]
sha256: 84688e730e01504e5f996161e676e9624127a4609f4d0fc474cd5c7b080fc22b
---

# Configure the FortiOS Provider for FortiGate

FortiOS Provider
The FortiOS provider is used to interact with the resources supported by FortiOS and FortiManager. We need to configure the provider with the proper credentials before it can be used.
Two products are supported now: FortiGate and FortiManager, please use the navigation on the left to read more details about the available resources.
Configuration for FortiGate
Example Usage
# Configure the FortiOS Provider for FortiGate
terraform {
  required_providers {
    fortios = {
      source  = "fortinetdev/fortios"
    }
  }
}
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}
# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
  dst     = "110.2.2.122/32"
  gateway = "2.2.2.2"
  # ...
}
If it is used for testing, you can set insecure to "true" and unset cabundlefile to quickly set the provider up, for example:
provider "fortios" {
  hostname = "192.168.52.177"
  token    = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure = "true"
}
Please refer to the Argument Reference below for more help on insecure and cabundlefile.
Authentication
The FortiOS provider offers a means of providing credentials for authentication. The following methods are supported:
- Static credentials
- Environment variables
Static credentials
Static credentials can be provided by adding credential keys in-line in the FortiOS provider block. Default using token if both credential are provided.
There are two kinds of credentials supported.
- token based authentication (Recommanded). User needs to generate an API token from FortiOS.
- username/password authentication. User provide the username and password of the administrator.
Usage:
terraform {
  required_providers {
    fortios = {
      source  = "fortinetdev/fortios"
    }
  }
}
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}
Generate an API token for FortiOS
See the left navigation: Guides -> Generate an API token for FortiOS.
Environment variables
You can provide your credentials via the FORTIOS_ACCESS_HOSTNAME, FORTIOS_ACCESS_TOKEN, FORTIOS_ACCESS_USERNAME, FORTIOS_ACCESS_PASSWORD, FORTIOS_INSECURE and FORTIOS_CA_CABUNDLE environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.
Usage:
$ export "FORTIOS_ACCESS_HOSTNAME"="192.168.52.177"
$ export "FORTIOS_INSECURE"="false"
$ export "FORTIOS_ACCESS_TOKEN"="09m441wrwc10yGyrtQ4nk6mjbqcfz9"
$ export "FORTIOS_CA_CABUNDLE"="/path/yourCA.crt"
Then configure the FortiOS Provider as following:
provider "fortios" {}
# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
  dst       = "110.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  # …
}
VDOM
If the FortiGate unit is running in VDOM mode, the vdom configuration needs to be added.
Usage:
terraform {
  required_providers {
    fortios = {
      source  = "fortinetdev/fortios"
    }
  }
}
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "q3Hs49jxts195gkd9Hjsxnjtmr6k39"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
  vdom         = "vdomtest"
}
resource "fortios_networking_route_static" "test1" {
  dst       = "120.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  priority  = "3"
  device    = "lbforvdomtest"
  comment   = "Terraform test"
}
By default, each resource inherits the provider's global vdom settings, but it can also set its own vdom through the vdomparam of each resource. See the vdomparam argument of each resource for details.
Argument Reference
The following arguments are supported:
- 
hostname - (Optional) The hostname or IP address of FortiOS unit. It must be provided, but it can also be sourced from theFORTIOS_ACCESS_HOSTNAME environment variable.
- 
token - (Optional) The token of FortiOS unit. If omitted, theFORTIOS_ACCESS_TOKEN environment variable will be used. If neither is set, username/password will be used.
- 
username - (Optional) The username of FortiOS unit. If omitted, theFORTIOS_ACCESS_USERNAME environment variable will be used.
- 
password - (Optional) The password of FortiOS unit.If omitted, theFORTIOS_ACCESS_PASSWORD environment variable will be used.
- 
insecure - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, theFORTIOS_INSECURE environment variable is used. If neither is set, default value isfalse .
- 
cabundlefile - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by theFORTIOS_CA_CABUNDLE environment variable.
- 
cabundlecontent - (Optional) The content of a custom CA bundle file. Note:cabundlefile andcabundlecontent cannot exist at the same time! Please only configure one of them.
- 
vdom - (Optional) If the FortiGate unit is running in VDOM mode, you can use this argument to specify the name of the vdom to be set .
- 
update_if_exist - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. This will apply to all the resources, resource level configuration will overwrite provider level configuration.
- 
http_proxy - (Optional) HTTP proxy address. Set this argument toENV if you want to use environment settings. By setting this argument toENV , the provider will get the environment variableHTTPS_PROXY orHTTP_PROXY . Default is empty, which means no HTTP proxy.
Configuration for FortiManager
Example Usage
provider "fortios" {
  fmg_hostname     = "192.168.88.100"
  fmg_username     = "APIUser"
  fmg_passwd       = "admin"
  fmg_insecure     = false
  fmg_cabundlefile = "/path/yourCA.crt"
}
resource "fortios_fmg_system_dns" "test1" {
  primary   = "208.91.112.52"
  secondary = "208.91.112.54"
}
If it is used for testing, you can set insecure to true and unset cabundlefile to quickly set the provider up, for example:
provider "fortios" {
  fmg_hostname = "192.168.88.100"
  fmg_username = "APIUser"
  fmg_passwd   = "admin"
  fmg_insecure = true
}
Please refer to the Argument Reference below for more help on insecure and cabundlefile.
Authentication
As the same to provider for FortiGate, the following two methods are supported:
- Static credentials
- Environment variables
Static credentials
Static credentials can be provided by adding the fmg_hostname, fmg_username and fmg_passwd key in-line in the FortiOS provider block.
Usage:
provider "fortios" {
  fmg_hostname     = "192.168.88.100"
  fmg_username     = "APIUser"
  fmg_passwd       = "admin"
  fmg_insecure     = false
  fmg_cabundlefile = "/path/yourCA.crt"
}
Environment variables
You can provide your credentials via the FORTIOS_FMG_HOSTNAME, FORTIOS_FMG_USERNAME, FORTIOS_FMG_PASSWORD, FORTIOS_FMG_INSECURE and FORTIOS_FMG_CABUNDLE environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.
Usage:
$ export "FORTIOS_FMG_HOSTNAME"="192.168.88.100"
$ export "FORTIOS_FMG_USERNAME"="admin"
$ export "FORTIOS_FMG_PASSWORD"="admin"
$ export "FORTIOS_FMG_INSECURE"="false"
$ export "FORTIOS_FMG_CABUNDLE"="/path/yourCA.crt"
Then configure the FortiOS Provider as following:
provider "fortios" {}
resource "fortios_fmg_system_dns" "test1" {
  primary   = "208.91.112.33"
  secondary = "208.91.112.44"
}
Multi-Adom
Multi-Adom feature is supported in case of using FortiManager, just take the following example as a reference:
provider "fortios" {
  fmg_hostname = "192.168.88.200"
  fmg_username = "APIUser"
  fmg_passwd   = "admin"
  fmg_product  = "fortimanager"
  fmg_insecure = true
}
resource "fortios_fmg_devicemanager_script" "test1" {
