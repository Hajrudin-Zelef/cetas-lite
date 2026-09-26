---
id: collect-260926-mikrotik/mikrotik/github-samsesh-github-ip-address-list-automatically-generate-mikrotik-address-list-from-gi
title: "Auto-generated MikroTik address list – GitHub IPs"
domain: mikrotik
role: reference
task: reference
actors: ["Meta"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/github-samsesh-github-ip-address-list-automatically-generate-mikrotik-address-list-from-github-s-pub.md
source_anchor: ""
source_lines: [1, 49]
sha256: 66284c4c8a2a59987a4c25f5cdffd2722ba4d4e15e44131ea356a8005d5a9707
---

# Auto-generated MikroTik address list – GitHub IPs

This repository automatically fetches the latest list of GitHub IP addresses from https://api.github.com/meta, and converts them into a MikroTik-compatible `.rsc` script file.

The script updates daily using GitHub Actions and is ideal for firewall rules, NAT exceptions, or routing GitHub traffic via specific interfaces.

- **`github-ip-list.rsc`** A MikroTik-compatible RouterOS script that adds all GitHub IP ranges into an address list named`GitHub` .

Example content:

```
# Auto-generated MikroTik address list – GitHub IPs
/ip firewall address-list
add address=140.82.112.0/20 list=GitHub
add address=185.199.108.0/22 list=GitHub
add address=192.30.252.0/22 list=GitHub
...
```
To automatically fetch and import the IP list into your MikroTik router, run the following script:

```
:foreach i in={"GitHub"} do={
  /tool fetch url="https://raw.githubusercontent.com/samsesh/github-ip-address-list/Localhost/github-ip-list.rsc" dst-path=$i
  /ip firewall address-list remove [/ip firewall address-list find list=$i]
  /import file-name=$i
  /file remove $i
}
```
✅ This script will:

- Download the latest IP list from your GitHub repo
- Remove old entries in the `GitHub` address list
- Import the updated list
- Clean up the downloaded file

This repo uses a GitHub Actions workflow that:

- Runs daily at midnight UTC
- Fetches the latest GitHub IP ranges from the GitHub Meta API
- Builds and commits a MikroTik `.rsc` file (`github-ip-list.rsc` ) to the`Localhost` branch

You can also run the workflow manually from the Actions tab.

- GitHub API: https://api.github.com/meta
- MikroTik Scripting Guide: https://wiki.mikrotik.com/wiki/Manual:Scripting
- MikroTik `/import` Command:`/import file-name=yourfile.rsc`

If you find this useful, you can support development here: 👉 https://donate.samsesh.net

**Author:** samsesh
**Website:** samsesh.net
