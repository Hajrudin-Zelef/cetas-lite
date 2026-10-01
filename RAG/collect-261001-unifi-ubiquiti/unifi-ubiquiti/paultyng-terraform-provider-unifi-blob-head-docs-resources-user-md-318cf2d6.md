---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-resources-user-md-318cf2d6
title: "paultyng-terraform-provider-unifi-blob-head-docs-resources-user-md-318cf2d6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/paultyng-terraform-provider-unifi-blob-head-docs-resources-user-md-318cf2d6.md
source_anchor: ""
source_lines: [1, 28]
sha256: d249b4da39cde8de3526cf79d227d5c9d7fd48171c3190ac1f00e2976a67e5ae
---

# paultyng-terraform-provider-unifi-blob-head-docs-resources-user-md-318cf2d6

| page_title | unifi_user Resource - terraform-provider-unifi | 
|---|---|
| subcategory |  | 
| description | unifi_user manages a user (or "client" in the UI) of the network, these are identified by unique MAC addresses. Users are created in the controller when observed on the network, so the resource defaults to allowing itself to just take over management of a MAC address, but this can be turned off. | 
unifi_user manages a user (or "client" in the UI) of the network, these are identified by unique MAC addresses.
Users are created in the controller when observed on the network, so the resource defaults to allowing itself to just take over management of a MAC address, but this can be turned off.
resource "unifi_user" "test" {
  mac  = "01:23:45:67:89:AB"
  name = "some client"
  note = "my note"
  fixed_ip   = "10.0.0.50"
  network_id = unifi_network.my_vlan.id
}
- mac (String) The MAC address of the user.
- name (String) The name of the user.
- allow_existing (Boolean) Specifies whether this resource should just take over control of an existing user. Defaults totrue .
- blocked (Boolean) Specifies whether this user should be blocked from the network.
- dev_id_override (Number) Override the device fingerprint.
- fixed_ip (String) A fixed IPv4 address for this user.
- local_dns_record (String) Specifies the local DNS record for this user.
- network_id (String) The network ID for this user.
- note (String) A note with additional information for the user.
- site (String) The name of the site to associate the user with.
- skip_forget_on_destroy (Boolean) Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults tofalse .
- user_group_id (String) The user group ID for the user.
- hostname (String) The hostname of the user.
- id (String) The ID of the user.
- ip (String) The IP address of the user.
