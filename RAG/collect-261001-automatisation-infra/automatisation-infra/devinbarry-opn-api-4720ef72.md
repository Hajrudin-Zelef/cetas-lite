---
id: collect-261001-automatisation-infra/automatisation-infra/devinbarry-opn-api-4720ef72
title: "Configure API client"
domain: automatisation-infra
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-automatisation-infra/devinbarry-opn-api-4720ef72.md
source_anchor: ""
source_lines: [1, 101]
sha256: af3389c71dd6161724b18b24994f986459d3f4fa558aca716aa0fa028dbdb1b6
---

# Configure API client

A Python client for the OPNsense REST API.

This library is based on opn-cli (v1.7.0) by Andreas Stürz, with additional code from python-opnsense by Dylan Turnbull. It has been stripped down to focus solely on API implementation while maintaining an extensible structure for easy addition of new API functions.

Tested against OPNsense versions 24 and 25. Fully supports OPNsense 25.7+ API format changes.

- Supports OPNsense API calls.
- Built for Python 3.12+ (but likely compatible with older versions).
- Designed for easy extension and customization.
- Includes support for core OPNsense functionalities like firewall, routing, VPN, syslog, and plugins.

You can install the OPN API client using `pip`:

`pip install opn-api`
Alternatively, if you want to install from source:

```
git clone https://github.com/devinbarry/opn-api.git
cd opn-api
pip install .
```
```
from opn_api.api.client import OPNAPIClient, OPNsenseClientConfig
# Configure API client
config = OPNsenseClientConfig(
    api_key="your_api_key",
    api_secret="your_api_secret",
    base_url="https://your-opnsense-instance/api",
    ssl_verify_cert=False  # Set to True if using valid SSL certs
)
# Create API client instance
client = OPNAPIClient(config)
```
```
from opn_api.client import OPNFirewallClient
from opn_api.models.firewall_alias import FirewallAliasCreate
fw = OPNFirewallClient(client)
# Retrieve a list of all aliases
aliases = fw.alias.list()
for alias in aliases:
    print(alias)
new_alias = fw.alias.add(FirewallAliasCreate(
    name="MyAlias",
    type="host",
    content=["192.168.1.100"],
    description="My test alias"
))
print(new_alias)
```
```
from opn_api.client import OPNFirewallClient
from opn_api.models.firewall_models import FirewallFilterRule
firewall = OPNFirewallClient(client)
# Get all rules
rules = firewall.filter.list_rules()
print(rules)
# Retrieve firewall rule by UUID
rule = firewall.filter.get_rule("your-rule-uuid")
print(rule)
# Add a new firewall rule
new_rule = firewall.filter.add_rule(FirewallFilterRule(
    sequence=1,
    action="pass",
    protocol="TCP",
    source_net="192.168.1.0/24",
    destination_net="8.8.8.8/32",
    description="Allow Google DNS"
))
print(new_rule)
```
```
from opn_api.client import OPNFirewallClient
fw = OPNFirewallClient(client)
# Fetch DHCP leases
leases = fw.dhcp.list_leases()
print(leases)
```
```
from opn_api.api.core.configbackup import Backup
backup = Backup(client)
config_data = backup.download()
with open("opnsense_backup.xml", "wb") as file:
    file.write(config_data.encode("utf-8"))
```
```
from opn_api.api.core.firmware import Firmware
firmware = Firmware(client)
info = firmware.info()
print(info)
```
To run the test suite, ensure `pytest` is installed:

`pip install pytest`
Then execute:

`pytest tests/`
Contributions are welcome! Please open an issue or a pull request if you would like to add features or fix bugs.

This project is licensed under the AGPLv3 License. See the LICENSE file for details.

For further details and documentation, visit OPNsense API Reference.
