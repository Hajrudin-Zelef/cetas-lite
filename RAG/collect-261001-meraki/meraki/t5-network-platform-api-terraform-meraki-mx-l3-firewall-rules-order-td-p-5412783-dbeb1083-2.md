---
id: collect-261001-meraki/meraki/t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083-2
title: "t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083.md
source_anchor: ""
source_lines: [23, 306]
sha256: cc4f1b0dc338fc70d62d96f10c1dfcc4ab8b53f52c15564d10e651af4b31a83a
---

# t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083

			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2025 01:15 PM
Can you kindly open an issue on our GitHub repo? Our developers will take a look.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2025 06:01 AM
You can try using a script (e.g., Python or Bash) that calls the Meraki API directly in the desired order. Terraform can trigger this script using the null_resource provisioner and local-exec.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2025 06:13 AM
Question. When using updateNetworkApplianceFirewallL3FirewallRules, the input is an array of rules.
Are you seeing different behavior between Terraform, Postman, Python, etc’?
If so, can you share the Terraform plan you’re using?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2025 12:07 PM
It is indeed an array.
This is the terraform resource for L3 firewall rules.
resource "meraki_networks_appliance_firewall_l3_firewall_rules" "sb3_fw_l3" {
  network_id = meraki_networks.sb3.id
  rules = [{
    comment        = "Deny-RFC1918."
    dest_cidr      = "10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16"
    dest_port      = "any"
    policy         = "deny"
    protocol       = "any"
    src_cidr       = "any"
    src_port       = "any"
    syslog_enabled = false
  },
  {
    comment        = "Allow-corp-outbound-to-internet."
    dest_cidr      = "any"
    dest_port      = "any"
    policy         = "allow"
    protocol       = "any"
    src_cidr       = var.subnet_prefix_corp
    src_port       = "any"
    syslog_enabled = false
  },
  {
    comment        = "Allow-iot-outbound-to-internet."
    dest_cidr      = "any"
    dest_port      = "any"
    policy         = "allow"
    protocol       = "any"
    src_cidr       = var.subnet_prefix_iot
    src_port       = "any"
    syslog_enabled = false
  },
  {
    comment        = "Allow-guest-outbound-to-internet."
    dest_cidr      = "any"
    dest_port      = "any"
    policy         = "allow"
    protocol       = "any"
    src_cidr       = var.subnet_prefix_guest
    src_port       = "any"
    syslog_enabled = false
  }
  ]
}
This is how it ends up in the Dashboard.
I deleted all rules above and pasted the same array into postman.
Postman body:
{
    "rules": [
        {
            "policy": "deny",
            "protocol": "any",
            "srcCidr": "any",
            "destCidr": "10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16",
            "comment": "Deny-RFC1918.",
            "srcPort": "any",
            "destPort": "any",
            "syslogEnabled": false
        },
        {
            "policy": "allow",
            "protocol": "any",
            "srcCidr": "10.100.2.0/24",
            "destCidr": "any",
            "comment": "Allow-corp-outbound-to-internet.",
            "srcPort": "any",
            "destPort": "any",
            "syslogEnabled": false
        },
        {
            "policy": "allow",
            "protocol": "any",
            "srcCidr": "10.110.2.0/24",
            "destCidr": "any",
            "comment": "Allow-iot-outbound-to-internet.",
            "srcPort": "any",
            "destPort": "any",
            "syslogEnabled": false
        },
        {
            "policy": "allow",
            "protocol": "any",
            "srcCidr": "10.120.2.0/24",
            "destCidr": "any",
            "comment": "Allow-guest-outbound-to-internet.",
            "srcPort": "any",
            "destPort": "any",
            "syslogEnabled": false
        }
    ]
}
Postman response:
200 OK
"rules": [
        {
            "comment": "Deny-RFC1918.",
            "policy": "deny",
            "protocol": "any",
            "srcPort": "Any",
            "srcCidr": "Any",
            "destPort": "Any",
            "destCidr": "10.0.0.0/8,172.16.0.0/12,192.168.0.0/16",
            "syslogEnabled": false
        },
        {
            "comment": "Allow-corp-outbound-to-internet.",
            "policy": "allow",
            "protocol": "any",
            "srcPort": "Any",
            "srcCidr": "10.100.2.0/24",
            "destPort": "Any",
            "destCidr": "Any",
            "syslogEnabled": false
        },
        {
            "comment": "Allow-iot-outbound-to-internet.",
            "policy": "allow",
            "protocol": "any",
            "srcPort": "Any",
            "srcCidr": "10.110.2.0/24",
            "destPort": "Any",
            "destCidr": "Any",
            "syslogEnabled": false
        },
        {
            "comment": "Allow-guest-outbound-to-internet.",
            "policy": "allow",
            "protocol": "any",
            "srcPort": "Any",
            "srcCidr": "10.120.2.0/24",
            "destPort": "Any",
            "destCidr": "Any",
            "syslogEnabled": false
        },
        {
            "comment": "Default rule",
            "policy": "allow",
            "protocol": "Any",
            "srcPort": "Any",
            "srcCidr": "Any",
            "destPort": "Any",
            "destCidr": "Any",
            "syslogEnabled": false
        }
    ]
}
Dashboard:
Everything is identical in the array. The only difference is how i populate the fields in terraform. I'm just pointing to some variables defined in variables.tf
Here they are:
variable "subnet_prefix_corp" {
    type = string
    default = "10.100.2.0/24"
}
variable "appliance_ip_corp" {
    type = string
    default = "10.100.2.1"
}
variable "subnet_prefix_iot" {
    type = string
    default = "10.110.2.0/24"
}
variable "appliance_ip_iot" {
    type = string
    default = "10.110.2.1"
}
variable "subnet_prefix_guest" {
    type = string
    default = "10.120.2.0/24"
}
variable "appliance_ip_guest" {
    type = string
    default = "10.120.2.1"
}
variable "subnet_prefix_mgmt" {
    type = string
    default = "10.130.2.0/24"
}
variable "appliance_ip_mgmt" {
    type = string
    default = "10.130.2.1"
}
It looks like Terraform is somehow re-ordering the array. But i'm not sure how or why
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2025 12:37 PM
Wrap your rules in a tolist() to force Terraform to treat it as an ordered list.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2025 01:10 PM
Thanks for the tip, but unfortunately it did not solve the issue.
rules wrapped in a tolist using locals:
locals {
  firewall_rules = tolist([
    {
      comment        = "Deny-RFC1918."
      dest_cidr      = "10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16"
      dest_port      = "any"
      policy         = "deny"
      protocol       = "any"
      src_cidr       = "any"
      src_port       = "any"
      syslog_enabled = false
    },
    {
      comment        = "Allow-corp-outbound-to-internet."
      dest_cidr      = "any"
      dest_port      = "any"
      policy         = "allow"
      protocol       = "any"
      src_cidr       = var.subnet_prefix_corp
      src_port       = "any"
      syslog_enabled = false
    },
    {
      comment        = "Allow-iot-outbound-to-internet."
      dest_cidr      = "any"
      dest_port      = "any"
      policy         = "allow"
      protocol       = "any"
      src_cidr       = var.subnet_prefix_iot
      src_port       = "any"
      syslog_enabled = false
    },
    {
      comment        = "Allow-guest-outbound-to-internet."
      dest_cidr      = "any"
      dest_port      = "any"
      policy         = "allow"
      protocol       = "any"
      src_cidr       = var.subnet_prefix_guest
      src_port       = "any"
