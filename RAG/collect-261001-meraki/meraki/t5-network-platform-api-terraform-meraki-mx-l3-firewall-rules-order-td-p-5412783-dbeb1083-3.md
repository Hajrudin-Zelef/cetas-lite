---
id: collect-261001-meraki/meraki/t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083-3
title: "t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083.md
source_anchor: ""
source_lines: [307, 488]
sha256: 3dc6a0550c1efcc3ac27879560d11f8cd067da7be6419dddbeea4d691ee0f185
---

# t5-network-platform-api-terraform-meraki-mx-l3-firewall-rules-order-td-p-5412783-dbeb1083

      syslog_enabled = false
    }
  ])
}
resource "meraki_networks_appliance_firewall_l3_firewall_rules" "sb3_fw_l3" {
  network_id = meraki_networks.sb3.id
  rules      = local.firewall_rules
}
Dashboard:
Copy of the state file block for firewall rules.
{
      "mode": "managed",
      "type": "meraki_networks_appliance_firewall_l3_firewall_rules",
      "name": "sb3_fw_l3",
      "provider": "provider[\"registry.terraform.io/cisco-open/meraki\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "network_id": "xxxxxxxxxx",
            "rules": [
              {
                "comment": "Allow-corp-outbound-to-internet.",
                "dest_cidr": "any",
                "dest_port": "any",
                "policy": "allow",
                "protocol": "any",
                "src_cidr": "10.100.2.0/24",
                "src_port": "any",
                "syslog_enabled": false
              },
              {
                "comment": "Allow-guest-outbound-to-internet.",
                "dest_cidr": "any",
                "dest_port": "any",
                "policy": "allow",
                "protocol": "any",
                "src_cidr": "10.120.2.0/24",
                "src_port": "any",
                "syslog_enabled": false
              },
              {
                "comment": "Allow-iot-outbound-to-internet.",
                "dest_cidr": "any",
                "dest_port": "any",
                "policy": "allow",
                "protocol": "any",
                "src_cidr": "10.110.2.0/24",
                "src_port": "any",
                "syslog_enabled": false
              },
              {
                "comment": "Deny-RFC1918.",
                "dest_cidr": "10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16",
                "dest_port": "any",
                "policy": "deny",
                "protocol": "any",
                "src_cidr": "any",
                "src_port": "any",
                "syslog_enabled": false
              }
            ],
            "rules_response": [
              {
                "comment": "Allow-corp-outbound-to-internet.",
                "dest_cidr": "Any",
                "dest_port": "Any",
                "policy": "allow",
                "protocol": "any",
                "src_cidr": "10.100.2.0/24",
                "src_port": "Any",
                "syslog_enabled": false
              },
              {
                "comment": "Allow-guest-outbound-to-internet.",
                "dest_cidr": "Any",
                "dest_port": "Any",
                "policy": "allow",
                "protocol": "any",
                "src_cidr": "10.120.2.0/24",
                "src_port": "Any",
                "syslog_enabled": false
              },
              {
                "comment": "Allow-iot-outbound-to-internet.",
                "dest_cidr": "Any",
                "dest_port": "Any",
                "policy": "allow",
                "protocol": "any",
                "src_cidr": "10.110.2.0/24",
                "src_port": "Any",
                "syslog_enabled": false
              },
              {
                "comment": "Default rule",
                "dest_cidr": "Any",
                "dest_port": "Any",
                "policy": "allow",
                "protocol": "Any",
                "src_cidr": "Any",
                "src_port": "Any",
                "syslog_enabled": false
              },
              {
                "comment": "Deny-RFC1918.",
                "dest_cidr": "10.0.0.0/8,172.16.0.0/12,192.168.0.0/16",
                "dest_port": "Any",
                "policy": "deny",
                "protocol": "any",
                "src_cidr": "Any",
                "src_port": "Any",
                "syslog_enabled": false
              }
            ],
            "syslog_default_rule": null
          },
          "sensitive_attributes": [],
          "identity_schema_version": 0,
          "dependencies": [
            "meraki_networks.sb3"
          ]
        }
      ]
    }
The order in the statefile matches what i see in the dashboard. I just dont get why it reorders the array.
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
07-17-2025 01:16 PM
With the debug information, please.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-17-2025 01:42 PM
Done. Here is a link to the issue.
Thanks for the replies 
https://github.com/cisco-open/terraform-provider-meraki/issues/274
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-21-2025 12:12 PM
Fyi there is a commit do the dev branch of the repo with a possible solution now.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-21-2025 01:54 PM
And now its merged with main. New provider version 1.1.7-beta is out
Thanks everyone!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-21-2025 03:47 PM
Happy coding!
