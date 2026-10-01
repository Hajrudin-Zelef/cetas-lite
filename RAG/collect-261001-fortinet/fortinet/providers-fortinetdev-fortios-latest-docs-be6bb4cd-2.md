---
id: collect-261001-fortinet/fortinet/providers-fortinetdev-fortios-latest-docs-be6bb4cd-2
title: "Configure the FortiOS Provider for FortiGate"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/providers-fortinetdev-fortios-latest-docs-be6bb4cd.md
source_anchor: ""
source_lines: [185, 219]
sha256: c645c4b7dc2bd8cef0d786a3397dd2b811f05f87cb8838425337acd3843f5fb9
---

# Configure the FortiOS Provider for FortiGate

  name        = "config-intf3"
  description = "configure interface3"
  content     = "config system interface \n edit port3 \n\t set vdom \"root\"\n\t set ip 10.10.0.200 255.255.0.0 \n\t set allowaccess ping http https\n\t next \n end"
  target      = "device_database"
  adom        = "test-adom"
}
resource "fortios_fmg_devicemanager_script_execute" "test1" {
  script_name    = fortios_fmg_devicemanager_script.test1.name
  target_devname = "FGVM64-test"
  adom           = "test-adom"
  vdom           = "root"
}
resource "fortios_fmg_devicemanager_install_device" "test1" {
  target_devname = fortios_fmg_devicemanager_script_execute.test1.target_devname
  adom           = "test-adom"
  vdom           = "root"
}
This will install the script from the FMG(test-adom) to the FGT(root).
Note that one resource supports Multi-Adom feature if it has 'adom' argument.
Argument Reference
The following arguments are supported:
- 
fmg_hostname - (Optional) The hostname or IP address of FortiManager. It must be provided, but it can also be sourced from theFORTIOS_FMG_HOSTNAME environment variable.
- 
fmg_username - (Optional) The username of FortiManager. It must be provided, but it can also be sourced from theFORTIOS_FMG_USERNAME environment variable.
- 
fmg_passwd - (Optional) The password of FortiManager, it can also be sourced from theFORTIOS_FMG_PASSWORD environment variable.
- 
fmg_insecure - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, theFORTIOS_FMG_INSECURE environment variable is used. If neither is set, default value isfalse .
- 
fmg_cabundlefile - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by theFORTIOS_FMG_CABUNDLE environment variable.
Release
Check out the FortiOS provider release notes and additional information from: the FortiOS provider releases.
Versioning
The provider can cover FortiOS 6.0, 6.2, 6.4, 7.0, 7.2, 7.4, 7.6 versions, the configuration of all parameters should be based on the relevant FortiOS version manual. The provider can cover FortiManager 6.0 and 6.2 versions. When using FortiManager, make sure the versions of FortiManager and the FortiGates controlled by it are the same.
