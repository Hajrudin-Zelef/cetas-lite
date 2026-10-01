---
id: collect-261001-fortinet/fortinet/itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff-1
title: "itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff.md
source_anchor: ""
source_lines: [1, 107]
sha256: 5fcccf54740ee567b46bdc33097bf2402c13a6bd9a7be38a1b594d51856d83c4
---

# itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff

FortiGate is Fortinet's network security appliance line, providing firewall, VPN, and routing functionality configured either via CLI (SSH) or the FortiOS REST API.
This project provides two ways to automate against FortiGate: netmiko over SSH/CLI (the same pattern used for Cisco IOS and Arista EOS), or a custom IG5 device driver over the FortiOS REST API for HTTP-based access — see Inventory Manager Configuration below.
Requirements: Itential Platform >= 6.4 · Itential Gateway >= 5.0 (only if using the fortigate-rest device driver)
| Asset | Description | 
|---|---|
| device-drivers/fortigate-rest | IG5 Python FortiOS REST driver — is-alive, get-config, plus a generic REST passthrough | 
| Configuration Parsers/fortios.json | Config Manager parser defining the fortios device type | 
| Golden Configurations/FortiGate - Simple | Golden config tree covering admin access, password policy, NTP, syslog, and interface hardening — requires the fortios parser | 
| Studio Projects/FortiGate | 4 workflows — golden config compliance, interface configuration (with pre/post checks and diff), interface reset, and backup/config snapshot | 
FortiGate devices can be automated two ways. Pick per device based on whether you want CLI (SSH) or REST (FortiOS API) access — both can coexist across different nodes in the same inventory.
Itential Platform ships with a netmiko driver for FortiGate out of the box — no additional driver install required. Broker actions (is-alive, run-command, get-config, set-config) are wired automatically when the inventory is created with createBrokerActions: true, the same as Cisco IOS and Arista EOS.
{
  "name": "my-fortigate-device",
  "attributes": {
    "itential_host": "192.0.2.200",
    "itential_port": 22,
    "itential_driver": "netmiko",
    "itential_platform": "fortinet",
    "itential_user": "admin",
    "itential_password": "changeme"
  }
}
| Attribute | Type | Description | 
|---|---|---|
| itential_host | string | Management IP or hostname of the device | 
| itential_port | integer | SSH port (default: 22 ) | 
| itential_driver | string | Driver to use — must be netmiko | 
| itential_platform | string | Netmiko device type — fortinet for FortiOS | 
| itential_user | string | SSH username | 
| itential_password | string | SSH password | 
Use this instead of netmiko when you want HTTP-based access to FortiGate's REST API rather than SSH/CLI — no SSH access required, and all operations go over HTTPS to the management interface. Requires the fortigate-rest device driver registered in Itential Gateway 5.x.
See device-drivers/fortigate-rest/README.md for full setup details, the API-token authentication model, and a generic REST passthrough (fortigate-rest-call) for calling any FortiOS REST endpoint directly from a workflow.
Node attributes:
{
  "name": "fortigate-01",
  "attributes": {
    "itential_host": "192.0.2.200",
    "itential_port": 443,
    "itential_driver": "fortigate-rest",
    "itential_password": "<FortiOS API token>",
    "itential_driver_options": {
      "fortigate-rest": {
        "vdom": "root",
        "verify_ssl": false,
        "timeout": 30,
        "backup_scope": "global"
      }
    }
  }
}
Action mapping — wire the supported broker contracts to this driver's IG5 services, replacing cluster-itential with your own cluster_id:
[
  {
    "name": "is-alive",
    "action_type": "iag5-service",
    "action_config": {
      "service_name": "fortigate-rest-is-alive",
      "cluster_id": "cluster-itential"
    },
    "action_parameters": {}
  },
  {
    "name": "get-config",
    "action_type": "iag5-service",
    "action_config": {
      "service_name": "fortigate-rest-get-config",
      "cluster_id": "cluster-itential"
    },
    "action_parameters": {}
  }
]
FortiOS's REST API is a structured, object-model interface (/api/v2/cmdb/*, /api/v2/monitor/*) — Fortinet's own support guidance is explicit that it isn't a CLI replacement, so it has no endpoint for executing arbitrary commands. That's a deliberate product design choice, not a driver limitation: fortigate-rest implements is-alive and get-config faithfully, but can't offer real run-command/set-config broker actions. Only SSH/CLI gives you those.
Use fortigate-rest (REST/API) when:
- SSH access to the management interface isn't available or is restricted by policy.
- You want structured JSON responses instead of parsing CLI text output.
- You're doing targeted object-level changes — firewall policies, addresses, and similar cmdb objects — viafortigate-rest-call .
- You want a lighter-weight reachability check or config snapshot without opening a CLI session.
Use netmiko (SSH/CLI) when:
- You need arbitrary operational commands (run-command ) or Config Manager CLI-diff remediation (set-config ).
- You need CLI-only functionality that has no cmdb REST equivalent.
Hybrid — no broker override needed (recommended for most deployments): fortigate-rest's connection resolution (main.py) reads itential_host, itential_password, and itential_driver_options.fortigate-rest.* directly off the node — it never looks at the node's top-level itential_driver or its actions array. That means a single node can be set up as Option 1 (itential_driver: netmiko, broker actions auto-wired for Config Manager/generic device operations) and also carry an itential_driver_options.fortigate-rest block for the REST token/settings:
{
  "name": "fortigate-01",
  "attributes": {
    "itential_host": "192.0.2.200",
    "itential_port": 22,
    "itential_driver": "netmiko",
    "itential_platform": "fortinet",
    "itential_user": "admin",
    "itential_password": "changeme",
    "itential_driver_options": {
      "fortigate-rest": {
        "api_token": "<FortiOS API token>",
        "vdom": "root",
        "verify_ssl": false
      }
    }
  }
}
Note the api_token override here — itential_password is already doing double duty as the SSH password for netmiko, so the REST token (a separate credential in FortiOS) goes in itential_driver_options.fortigate-rest.api_token instead. With this in place, a Studio Project workflow can call fortigate-rest-is-alive, fortigate-rest-get-config, or fortigate-rest-call as an Automation Gateway task against this same node at any time, resolving host/token from these attributes — independent of whatever the node's broker actions are wired to. No changes to the actions array, and no second node needed.
Hybrid — broker-level override: only reach for this if you specifically want the platform's own broker-driven features (Config Manager compliance/backup, generic device get-config/is-alive calls) to use REST instead of CLI, rather than just calling the service directly from a workflow:
- Create the node with itential_driver: netmiko /itential_platform: fortinet (Option 1) so Itential Platform auto-wires all four broker actions viacreateBrokerActions: true .
- In Inventory Manager, edit just the is-alive andget-config action entries to override theiraction_config tofortigate-rest-is-alive /fortigate-rest-get-config (Option 2's services) instead.
- Leave run-command andset-config on their auto-generated netmiko mapping, since those need CLI access anyway.
This still requires the same itential_driver_options.fortigate-rest block as above.
A native Python FortiOS REST driver for IG5. Use this for HTTP-based FortiGate automation instead of SSH/CLI — useful when SSH access to the management interface isn't available, or when you'd rather work directly against FortiOS REST endpoints from a workflow.
See device-drivers/fortigate-rest/README.md for full documentation including all operations, the API-token authentication model, and local testing. That README also has the Inventory Manager action mapping JSON needed to wire the broker contracts (is-alive, get-config) to this driver's IG5 services.
