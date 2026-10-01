---
id: collect-261001-meraki/meraki/zebbkira-merakiautomationskill-blob-head-skills-meraki-automation-skill-md-48709c7d-2
title: "Keyword search (returns up to 20 results by default)"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/zebbkira-merakiautomationskill-blob-head-skills-meraki-automation-skill-md-48709c7d.md
source_anchor: ""
source_lines: [139, 226]
sha256: d20dd133704aad13c1f0b6117be49e6381c9e74392d710b03021a0fab7b90b5b
---

# Keyword search (returns up to 20 results by default)

            return data  # Non-list response, no pagination needed
    return all_results
Encapsulate each API call as an independent function. Use a main() entry point:
def get_organizations():
    """Return list of organizations accessible with this API key."""
    url = f"{BASE_URL}/organizations"
    return make_request("GET", url)
def get_organization_networks(org_id):
    """Return the list of networks in an organization."""
    url = f"{BASE_URL}/organizations/{org_id}/networks"
    return get_all_pages(url)
def update_network_wireless_ssid(network_id, ssid_number, payload):
    """Update the attributes of an MR SSID.
    
    payload example:
        {"name": "My SSID", "enabled": True, "authMode": "psk", "psk": "password"}
    """
    url = f"{BASE_URL}/networks/{network_id}/wireless/ssids/{ssid_number}"
    return make_request("PUT", url, payload=payload)
def main():
    org_id = "YOUR_ORG_ID"
    print("Fetching organizations...")
    orgs = get_organizations()
    for org in orgs:
        print(f"  {org['id']}: {org['name']}")
    print(f"\nFetching networks for org {org_id}...")
    networks = get_organization_networks(org_id)
    print(f"  Found {len(networks)} networks")
if __name__ == "__main__":
    main()
Use the resp field from api_index.json or the full response_examples from parse_collection.py --name to know the exact field names:
# Example: getOrganizationDevices returns list of device objects
# resp keys: ["serial", "name", "model", "mac", "networkId", "lanIp", "firmware", "status", ...]
devices = get_all_pages(f"{BASE_URL}/organizations/{org_id}/devices")
for device in devices:
    print(f"{device['serial']} | {device['name']} | {device['model']} | {device.get('lanIp', 'N/A')}")orgs = make_request("GET", f"{BASE_URL}/organizations")
for org in orgs:
    networks = get_all_pages(f"{BASE_URL}/organizations/{org['id']}/networks")
    for net in networks:
        # process each network
        passerrors = []
for item in items:
    try:
        result = make_request("PUT", url, payload=payload)
    except Exception as e:
        errors.append({"item": item, "error": str(e)})
if errors:
    print(f"Completed with {len(errors)} errors:")
    for err in errors:
        print(f"  {err}")import csv, json
# JSON export
with open("output.json", "w", encoding="utf-8") as f:
    json.dump(results, f, ensure_ascii=False, indent=2)
# CSV export
with open("output.csv", "w", newline="", encoding="utf-8") as f:
    if results:
        writer = csv.DictWriter(f, fieldnames=results[0].keys())
        writer.writeheader()
        writer.writerows(results)
Use these category paths to navigate api_index.json efficiently:
| Task | Category in api_index.json | 
|---|---|
| Organizations | getOrganizations ,platform/configure | 
| Networks | getOrganizationNetworks ,platform/configure/devices | 
| Devices / Inventory | getOrganizationDevices ,platform/configure/inventory | 
| Network Clients | getNetworkClients ,platform/configure/clients | 
| Admins | platform/configure/admins | 
| Config Templates | platform/configure/configTemplates | 
| Firmware | platform/configure/firmwareUpgrades | 
| Alerts & Webhooks | platform/configure/alerts ,platform/configure/webhooks | 
| Appliance (MX) | products/appliance/configure/* ,products/appliance/monitor/* | 
| Switch (MS) | products/switch/configure/* ,products/switch/monitor/* | 
| Wireless (MR) | products/wireless/configure/* ,products/wireless/monitor/* | 
| Camera (MV) | products/camera/configure/* ,products/camera/monitor/* | 
| Sensor (MT) | products/sensor/configure/* ,products/sensor/monitor/* | 
| Cellular Gateway | products/cellularGateway/configure/* | 
| Insight | products/insight/configure/* ,products/insight/monitor/* | 
| Systems Manager | products/sm/configure/* | 
| SASE sites / connectors | platform/configure/sase ,platform/monitor/sase | 
| Global policies | platform/configure/policies/global ,products/appliance/configure/policies/global | 
If the Postman Collection is updated, regenerate the index by running:
python scripts/generate_index.py "<path_to_collection.json>" --output references/api_index.json
Keep the collection unchanged, update its filename and version references, and use
the generated meta.total and meta.categories for documentation counts. Check
the new collection's categories before changing category hints. Validate offline:
python -m unittest discover -s scripts -p "test_*.py"
python scripts/search_index.py --name "getOrganizationSaseSites"
python scripts/parse_collection.py "references/Meraki Dashboard API - v1.69.0.postman_collection.json" --name "getOrganizationSaseSites"
