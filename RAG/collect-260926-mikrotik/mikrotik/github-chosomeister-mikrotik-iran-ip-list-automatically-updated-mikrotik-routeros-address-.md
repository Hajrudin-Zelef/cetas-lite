---
id: collect-260926-mikrotik/mikrotik/github-chosomeister-mikrotik-iran-ip-list-automatically-updated-mikrotik-routeros-address-
title: "github-chosomeister-mikrotik-iran-ip-list-automatically-updated-mikrotik-routeros-address-list-of-ir"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/github-chosomeister-mikrotik-iran-ip-list-automatically-updated-mikrotik-routeros-address-list-of-ir.md
source_anchor: ""
source_lines: [1, 122]
sha256: 67a9457fdce5960621086642dc78d131da5da4a692ab4d27384c614de8de20e1
---

# github-chosomeister-mikrotik-iran-ip-list-automatically-updated-mikrotik-routeros-address-list-of-ir

Automatically generated IPv4 and IPv6 address lists for MikroTik RouterOS, based on RIPE country-resource data for Iran (`IR`).

`list.rsc` rebuilds the `IRAN-IP-Address` firewall address list and is ready to import directly into RouterOS.

- `list.rsc` : generated RouterOS script containing Iran IPv4 and IPv6 prefixes
- `get.sh` : Bash generator that downloads RIPE data and produces`list.rsc`
- `.github/workflows/` : repository automation for refreshing generated data
- `README.md` : project documentation

- Includes both IPv4 and IPv6 prefixes
- Ready for direct `/import` on MikroTik
- Clears and rebuilds the same address list on each update
- Uses `:do { add ... } on-error={}` for tolerant imports
- Built from RIPE Stat country resource data for `IR`
- Current generator also adds `10.0.0.0/8` to the IPv4 list

Source data comes from the RIPE Stat `country-resource-list` endpoint for Iran:

`https://stat.ripe.net/data/country-resource-list/data.json?resource=IR&v4_format=prefix`

The repository converts those prefixes into RouterOS commands for:

- `/ip firewall address-list`
- `/ipv6 firewall address-list`

1. `get.sh` downloads the RIPE JSON payload with`curl`
2. It verifies the HTTP response and checks that the IPv4 list is not empty
3. It writes commands to remove the current `IRAN-IP-Address` entries
4. It appends all IPv4 prefixes and explicitly adds `10.0.0.0/8`
5. It writes a second section for IPv6 prefixes using the same list name
6. The generated output is saved as `list.rsc`

- Policy routing for Iran destinations
- NAT bypass or no-NAT matching
- Firewall filtering by source or destination range
- Traffic marking and routing decisions
- Queueing or classification rules

```
:local fileName "IRAN-IP-Address.rsc"
:local url "https://raw.githubusercontent.com/ChosoMeister/Mikrotik-IRAN-IP-List/master/list.rsc"
/tool fetch url=$url dst-path=$fileName
/import file-name=$fileName
/file remove $fileName
```
```
:foreach i in={"IRAN-IP-Address"} do={
  /tool fetch url="https://raw.githubusercontent.com/ChosoMeister/Mikrotik-IRAN-IP-List/master/list.rsc" dst-path=$i
  /import file-name=$i
  /file remove $i
}
```
1. Download `list.rsc`
2. Upload it to the router with Winbox, WebFig, or SCP
3. Run:

```
/import file-name=list.rsc
```
For production use, it is safer to import only if the download succeeded and the file is not empty.

```
:local fileName "IRAN-IP-Address.rsc"
:local url "https://raw.githubusercontent.com/ChosoMeister/Mikrotik-IRAN-IP-List/master/list.rsc"
/tool fetch url=$url dst-path=$fileName
:delay 5
:if ([/file find name=$fileName] != "") do={
  :if ([/file get $fileName size] > 0) do={
    /import file-name=$fileName
    /file remove $fileName
  } else={
    :log warning "Iran IP list download was empty; keeping current address list"
    /file remove $fileName
  }
} else={
  :log warning "Iran IP list download failed; keeping current address list"
}
```
You can place that logic in a RouterOS script and run it daily with `/system scheduler`.

```
/ip firewall filter add chain=forward dst-address-list=IRAN-IP-Address action=accept comment="Example: match Iran IPv4 destinations"
/ipv6 firewall filter add chain=forward dst-address-list=IRAN-IP-Address action=accept comment="Example: match Iran IPv6 destinations"
```
Adjust the action and chain to fit your policy.

- `bash`
- `curl`
- `jq`

```
chmod +x get.sh
./get.sh > list.rsc
```
If you maintain this repository, update `get.sh` first and then regenerate `list.rsc`.

- `list.rsc` removes existing entries from`IRAN-IP-Address` before repopulating it
- The generated file contains an IPv6 section
- Routers without IPv6 support may need the IPv6 block removed before import
- Country-assigned prefixes are not the same as guaranteed traffic geolocation
- IP allocations change over time, so regular refresh is important
- The current generator intentionally includes `10.0.0.0/8` ; keep or remove that behavior based on your routing policy

- `tool fetch` fails:
Check DNS, router clock, CA certificates, and access to`raw.githubusercontent.com`
- Import stops at the IPv6 section:
Enable IPv6 support or remove the IPv6 part of `list.rsc`
- Local generation fails:
Install `jq` and verify the RIPE endpoint is reachable
- The address list is replaced after each update: This is expected; the script rebuilds the list from scratch

Contributions are welcome.

Useful improvements include:

- safer update workflows
- RouterOS compatibility fixes
- automation improvements
- documentation cleanup
- validation improvements in `get.sh`

If you change the generation logic, regenerate `list.rsc` from `get.sh` instead of editing the generated file by hand.
