---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-resources-device-md-2ccfe76c
title: "paultyng-terraform-provider-unifi-blob-head-docs-resources-device-md-2ccfe76c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-resources-device-md-2ccfe76c.md
source_anchor: ""
source_lines: [1, 56]
sha256: 58ded753109bd152b6f056391a48cc0c440d941b8671c616812c3a76480d7ef6
---

# paultyng-terraform-provider-unifi-blob-head-docs-resources-device-md-2ccfe76c

| page_title | unifi_device Resource - terraform-provider-unifi | 
|---|---|
| subcategory |  | 
| description | unifi_device manages a device of the network. Devices are adopted by the controller, so it is not possible for this resource to be created through Terraform, the create operation instead will simply start managing the device specified by MAC address. It's safer to start this process with an explicit import of the device. | 
unifi_device manages a device of the network.
Devices are adopted by the controller, so it is not possible for this resource to be created through Terraform, the create operation instead will simply start managing the device specified by MAC address. It's safer to start this process with an explicit import of the device.
data "unifi_port_profile" "disabled" {
  # look up the built-in disabled port profile
  name = "Disabled"
}
resource "unifi_port_profile" "poe" {
  name    = "poe"
  forward = "customize"
  native_networkconf_id = var.native_network_id
  tagged_networkconf_ids = [
    var.some_vlan_network_id,
  ]
  poe_mode = "auto"
}
resource "unifi_device" "us_24_poe" {
  # optionally specify MAC address to skip manually importing
  # manual import is the safest way to add a device
  mac = "01:23:45:67:89:AB"
  name = "Switch with POE"
  port_override {
    number          = 1
    name            = "port w/ poe"
    port_profile_id = unifi_port_profile.poe.id
  }
  port_override {
    number          = 2
    name            = "disabled"
    port_profile_id = data.unifi_port_profile.disabled.id
  }
  # port aggregation for ports 11 and 12
  port_override {
    number              = 11
    op_mode             = "aggregate"
    aggregate_num_ports = 2
  }
}
- allow_adoption (Boolean) Specifies whether this resource should tell the controller to adopt the device on create. Defaults totrue .
- forget_on_destroy (Boolean) Specifies whether this resource should tell the controller to forget the device on destroy. Defaults totrue .
- mac (String) The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
- name (String) The name of the device.
- port_override (Block Set) Settings overrides for specific switch ports. (see below for nested schema)
- site (String) The name of the site to associate the device with.
- disabled (Boolean) Specifies whether this device should be disabled.
- id (String) The ID of the device.
Required:
- number (Number) Switch port number.
Optional:
- aggregate_num_ports (Number) Number of ports in the aggregate.
- name (String) Human-readable name of the port.
- op_mode (String) Operating mode of the port, valid values areswitch ,mirror , andaggregate . Defaults toswitch .
- port_profile_id (String) ID of the Port Profile used on this port.
