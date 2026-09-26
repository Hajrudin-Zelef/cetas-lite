---
id: collect-260926-mikrotik/mikrotik/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d-1
title: "runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d.md
source_anchor: ""
source_lines: [1, 100]
sha256: e92b72f9e4e3a6f8f64ffb9e8fb67a82ceddf987b7cbf26f10fca975a1cff933
---

# runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d

Imports a MikroTik router and — the reason this integration exists — the devices that router has observed: its ARP table, its DHCP leases, the CDP, LLDP, and MNDP neighbors it has discovered, and the stations associated to its radios.
MikroTik is everywhere in homelab, WISP, and small-office networks, and in most of those networks the router is the only thing with a complete picture of what is attached. The discovery neighbor table in particular is unusually rich: a neighbor announces its own identity, platform, board model, and software version, so those assets arrive already described rather than as a bare address.
- Superuser access to the Custom Integrations configuration in runZero.
- An Explorer with network access to the router's web service.
- RouterOS 7.1 or newer. The REST API was introduced in 7.1beta4 as a JSON wrapper over the console. RouterOS 6 has no REST API at all — its only programmatic interface is the binary API on 8728/8729, which this integration does not implement and which is out of scope. Point this at a v6 device and the first request fails; the log says so and names the version floor.
- The web service must be enabled. REST is served by www-ssl (HTTPS) orwww (plain HTTP).
  - www-ssl will not start without a certificate assigned, which is the single most common reason a first attempt fails.
  - Plain HTTP on www only serves REST from RouterOS 7.9 onward (*) www - allow unsecure HTTP access to REST API; ). On 7.1–7.8, HTTPS is the only option.
With a self-signed certificate:
/certificate add name=rest-tmpl common-name=rest-api days-valid=3650 key-size=2048
/certificate sign rest-tmpl name=rest-api
/ip service set www-ssl certificate=rest-api disabled=no
Or, on 7.9+, plain HTTP if you accept that Basic credentials cross the network unencrypted:
/ip service set www disabled=no
rest-api is a separate policy from api, and both are separate from
web. A read-only user needs read and rest-api; the built-in read group
already contains both, but a dedicated group is better because it can be pinned
to the Explorer's address:
/user group add name=runzero-ro policy=read,rest-api,api
/user add name=runzero group=runzero-ro password=<secret> address=192.0.2.10/32
- read andrest-api are the load-bearing pair.
- api is harmless and only matters if you ever fall back to the binary API.
- web is not required for/rest .
- sensitive is not required and should not be granted — it exposes stored passwords and keys, and no table this integration reads needs it.
- address= restricts the account to the Explorer, which is worth doing.
Confirm the credential from the Explorer host:
curl -sk -u 'runzero:<secret>' https://192.168.88.1/rest/system/resource
A success returns a JSON array with one object — every RouterOS REST read returns an array, even for a single-record menu. A wrong password returns HTTP 401 with no body. A missing policy returns HTTP 500, not 403:
{"detail": "not enough permissions (9)", "error": 500, "message": "Internal Server Error"}
RouterOS serves www-ssl with whatever certificate you assigned, which is
normally self-signed. Set the TLS options accordingly rather than assuming
validation will pass. Do not disable validation without deciding that is
acceptable for your environment.
- Create the Custom Integration.
  - Add a Name and Icon for the integration (e.g., "MikroTik RouterOS").
  - Toggle Enable custom integration script to input the finalized script.
  - Click Validate to ensure it has valid syntax.
  - Click Save to create the Custom Integration.
- Create the Credential for the Custom Integration.
  - Select the type Custom Integration Script Secrets .
  - RouterOS URL (url ): base URL of the device, for examplehttps://192.168.88.1 . The API is resolved as<url>/rest .
  - Username (username ) and Password (password ).
  - Import the router itself (import_router ): optional, default enabled.
  - Collect ARP entries (collect_arp ) / Collect DHCP leases (collect_dhcp_leases ) / Collect discovery neighbors (collect_neighbors ) / Collect wireless stations (collect_wireless ): optional, all default enabled.
  - Bound DHCP leases only (bound_leases_only ): optional, default disabled.
  - Maximum hosts (max_hosts ): optional, default 10000,0 removes the cap.
  - TLS options (tls_* ): set these for the device's self-signed certificate.
- Select the type 
- Create the Custom Integration task.
  - Select the Credential and Custom Integration created in steps 1 and 2.
  - Select the Explorer you would like the Custom Integration to run from.
  - Schedule it. ARP entries and wireless associations age out in minutes, so an hourly run sees substantially more of a busy network than a daily one.
  - Click Save to kick off the first task.
- You will see the task kick off on the tasks page like any other integration.
- The task will update existing assets with data pulled from RouterOS.
- The task will create new assets when there are no existing assets that meet merge criteria (hostname, MAC, etc).
- Search for everything this integration touched with custom_integration:mikrotik-routeros .
- Split the two asset kinds with tag:mikrotik-router andtag:mikrotik-host .
- Find how a device was seen with tag:mikrotik-arp ,tag:mikrotik-dhcp-lease ,tag:mikrotik-neighbor , ortag:mikrotik-wireless .
- Find everything on one SSID with tag:ssid:corp-wifi , or find weak wireless clients withmikrotik_wifi_signal_dbm:<-75 .
- Find the router by serial with tag:serial:HGT08ABCDEF .
runzero script --filename mikrotik-routeros/mikrotik-routeros.star \
  --kwargs url=https://192.168.88.1 \
  --kwargs username=runzero \
  --kwargs password='correct horse battery staple' \
  --kwargs tls_disable_validation=true \
  --kwargs collect_wireless=false \
  --kwargs max_hosts=200 \
  --custom-integration-id 1f2e3d4c-5b6a-7988-9a0b-1c2d3e4f5a6b \
  --output /tmp/mikrotik-run --overwrite
--output writes the serialized assets so you can inspect exactly what would be
imported; --overwrite lets you re-run into the same directory.
To check only that the CONFIG block and the HTTP/TLS wiring are sound, without
touching a real device:
runzero script --filename mikrotik-routeros/mikrotik-routeros.star --validate
To run it the way the platform does — as an integration task that uploads its
results — use the scan command with the custom integration flags.
--custom-integration-id is the UUID shown on the integration's page in the
console:
runzero scan --api-key "$RUNZERO_API_KEY" \
  --custom-integration-id 1f2e3d4c-5b6a-7988-9a0b-1c2d3e4f5a6b \
  --custom-integration-script-kwargs 'url=https://192.168.88.1,username=runzero,password=...'
One CLI caveat: --kwargs passes a value through verbatim, commas included,
until the value contains a second = -- at which point it is parsed as CSV,
so a password containing a comma and an = is silently torn into extra
parameters. Test such a credential through the console form rather than the CLI.
Confirm the flag set with runzero script --help and runzero scan --help on
your own scanner build; it does change between releases.
Two asset kinds, and — unusually — they use different match behavior, because MikroTik publishes a genuine hardware identifier for one of them and nothing at all for the other.
There are three declared asset types, because the router itself splits at
runtime: the script picks router when /system/routerboard yielded a
serial and router-unkeyed when it did not, and observed devices are host.
The type is selected per asset with ImportAsset(assetType=...), and each
type's complete merge policy is declared under that key in
CONFIG["assetTypeBehavior"]:
| type | when | match behavior | 
|---|---|---|
| router | /system/routerboard returned a serial | no-mac-break no-ip-break no-name-break | 
| router-unkeyed | CHR / x86, no serial | no-id-match no-id-break | 
| host | an observed ARP, lease, neighbor, or wireless entry | no-id-match no-id-break | 
