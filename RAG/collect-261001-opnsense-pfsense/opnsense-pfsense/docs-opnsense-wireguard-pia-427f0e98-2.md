---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-opnsense-wireguard-pia-427f0e98-2
title: "docs-opnsense-wireguard-pia-427f0e98"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-opnsense-wireguard-pia-427f0e98.md
source_anchor: ""
source_lines: [178, 198]
sha256: 780704afaef670244133011b746e0e45316f4df55ef4b6818c025ce9dd35d03e
---

# docs-opnsense-wireguard-pia-427f0e98

Make sure you have curl and jq installed. curl is used to transfer data from a server and jq is a lightweight and flexible command-line JSON processor.
sudo apt install jq curl
4 servers closest to you
After installation, enter the command below to fetch the 4 recommended servers closest to you, with support for WireGuard:
curl -s "https://api.nordvpn.com/v1/servers/recommendations?&filters\[servers_technologies\]\[identifier\]=wireguard_udp&limit=4"|jq -r '.[]|.hostname, .station, (.locations|.[]|.country|.city.name), (.locations|.[]|.country|.name), (.technologies|.[].metadata|.[].value), .load'
fr111.nordvpn.com #your endpoint host
132.44.11.81 #its ip address
Paris #city
France #country
K19dhzkuUGsQkosdahaDKjall18/00idianhkplaUJkmk= #Server public key
15 #Server load at the time.
(...)
Display servers in specific country
Bosnia and Herzegovina:
curl --silent https://api.nordvpn.com/server | jq --raw-output '.[] | select(.country == "Bosnia and Herzegovina") | .domain'
Display the country ID:
curl --silent "https://api.nordvpn.com/v1/servers/countries" | jq --raw-output '.[] | select(.name == "Bosnia and Herzegovina") | [.name, .id] | "ID for \(.[0]) is \(.[1])"'
Display 4 servers with WG in specific country
curl -s "https://api.nordvpn.com/v1/servers/recommendations?&filters\[servers_technologies\]\[identifier\]=wireguard_udp&filters\[country_id\]=209&limit=4"|jq -r '.[]|.hostname, .station, (.locations|.[]|.country|.city.name), (.locations|.[]|.country|.name), (.technologies|.[].metadata|.[].value), .load'
Authors
Mr. Johnson
