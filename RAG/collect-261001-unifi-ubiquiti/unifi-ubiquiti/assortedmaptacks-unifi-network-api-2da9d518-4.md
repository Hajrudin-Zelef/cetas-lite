---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/assortedmaptacks-unifi-network-api-2da9d518-4
title: "Clone the repository"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "mit license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/assortedmaptacks-unifi-network-api-2da9d518.md
source_anchor: ""
source_lines: [356, 463]
sha256: 0bc613659d06dac1e30da7e2dbeef886c7f2d352ff38d75e9db8f9119b2a04ed
---

# Clone the repository

|  | /api/s/{site}/upd/firewallgroup/{id} | PUT | Update Firewall Group | 
|  | /api/s/{site}/upd/portconf/{id} | PUT | Update Port Configuration | 
|  | /api/s/{site}/upd/radiusprofile/{id} | PUT | Update RADIUS Profile | 
|  | /api/s/{site}/upd/account/{id} | PUT | Update Account | 
|  | /api/s/{site}/upd/wlangroup/{id} | PUT | Update WLAN Group | 
|  | /api/s/{site}/upd/hotspotop/{id} | PUT | Update Hotspot Operator | 
|  | /api/s/{site}/upd/hotspotpackage/{id} | PUT | Update Hotspot Package | 
|  | /api/s/{site}/upd/portforward/{id} | PUT | Update Port Forwarding Rule | 
|  | /api/s/{site}/upd/dynamicdns/{id} | PUT | Update Dynamic DNS Configuration | 
|  | /api/s/{site}/upd/portprofile/{id} | PUT | Update Port Profile | 
|  | /api/s/{site}/upd/dpiapp/{id} | PUT | Update DPI Application | 
|  | /api/s/{site}/upd/dpigroup/{id} | PUT | Update DPI Group | 
|  | /api/s/{site}/del/user/{id} | DELETE | Delete User | 
|  | /api/s/{site}/del/usergroup/{id} | DELETE | Delete User Group | 
|  | /api/s/{site}/del/wlanconf/{id} | DELETE | Delete WLAN Configuration | 
|  | /api/s/{site}/del/networkconf/{id} | DELETE | Delete Network Configuration | 
|  | /api/s/{site}/del/firewallrule/{id} | DELETE | Delete Firewall Rule | 
|  | /api/s/{site}/del/firewallgroup/{id} | DELETE | Delete Firewall Group | 
|  | /api/s/{site}/del/portconf/{id} | DELETE | Delete Port Configuration | 
|  | /api/s/{site}/del/radiusprofile/{id} | DELETE | Delete RADIUS Profile | 
|  | /api/s/{site}/del/account/{id} | DELETE | Delete Account | 
|  | /api/s/{site}/del/wlangroup/{id} | DELETE | Delete WLAN Group | 
|  | /api/s/{site}/del/hotspotop/{id} | DELETE | Delete Hotspot Operator | 
|  | /api/s/{site}/del/hotspotpackage/{id} | DELETE | Delete Hotspot Package | 
|  | /api/s/{site}/del/voucher/{id} | DELETE | Delete Voucher | 
|  | /api/s/{site}/del/portforward/{id} | DELETE | Delete Port Forwarding Rule | 
|  | /api/s/{site}/del/dynamicdns/{id} | DELETE | Delete Dynamic DNS Configuration | 
|  | /api/s/{site}/del/portprofile/{id} | DELETE | Delete Port Profile | 
|  | /api/s/{site}/del/dpiapp/{id} | DELETE | Delete DPI Application | 
|  | /api/s/{site}/del/dpigroup/{id} | DELETE | Delete DPI Group | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Content Filtering | /v2/api/site/{site}/content-filtering | GET | Get Content Filtering Configuration | 
|  | /v2/api/site/{site}/content-filtering | PUT | Update Content Filtering Configuration | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| DHCP | /v2/api/site/{site}/excluded-ips | GET | List Excluded IPs | 
|  | /v2/api/site/{site}/excluded-ips | POST | Create Excluded IP | 
|  | /v2/api/site/{site}/excluded-ips/{id} | GET | Get Excluded IP Details | 
|  | /v2/api/site/{site}/excluded-ips/{id} | PUT | Update Excluded IP | 
|  | /v2/api/site/{site}/excluded-ips/{id} | DELETE | Delete Excluded IP | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Hotspot | /v2/api/site/{site}/hotspot2-conf | GET | List Hotspot 2.0 Configurations | 
|  | /v2/api/site/{site}/hotspot2-conf | POST | Create Hotspot 2.0 Configuration | 
|  | /v2/api/site/{site}/hotspot2-conf/{id} | GET | Get Hotspot 2.0 Configuration Details | 
|  | /v2/api/site/{site}/hotspot2-conf/{id} | PUT | Update Hotspot 2.0 Configuration | 
|  | /v2/api/site/{site}/hotspot2-conf/{id} | DELETE | Delete Hotspot 2.0 Configuration | 
|  | /v2/api/site/{site}/hotspot-operators | GET | List Hotspot Operators | 
|  | /v2/api/site/{site}/hotspot-operators | POST | Create Hotspot Operator | 
|  | /v2/api/site/{site}/hotspot-operators/{id} | GET | Get Hotspot Operator Details | 
|  | /v2/api/site/{site}/hotspot-operators/{id} | PUT | Update Hotspot Operator | 
|  | /v2/api/site/{site}/hotspot-operators/{id} | DELETE | Delete Hotspot Operator | 
|  | /v2/api/site/{site}/hotspot-packages | GET | List Hotspot Packages | 
|  | /v2/api/site/{site}/hotspot-packages | POST | Create Hotspot Package | 
|  | /v2/api/site/{site}/hotspot-packages/{id} | GET | Get Hotspot Package Details | 
|  | /v2/api/site/{site}/hotspot-packages/{id} | PUT | Update Hotspot Package | 
|  | /v2/api/site/{site}/hotspot-packages/{id} | DELETE | Delete Hotspot Package | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Static DNS | /v2/api/site/{site}/static-dns | GET | List Static DNS Records | 
|  | /v2/api/site/{site}/static-dns | POST | Create Static DNS Record | 
|  | /v2/api/site/{site}/static-dns/{id} | GET | Get Static DNS Record Details | 
|  | /v2/api/site/{site}/static-dns/{id} | PUT | Update Static DNS Record | 
|  | /v2/api/site/{site}/static-dns/{id} | DELETE | Delete Static DNS Record | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Teleport | /v2/api/site/{site}/teleport | GET | Get Teleport Configuration | 
|  | /v2/api/site/{site}/teleport | PUT | Update Teleport Configuration | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Topology | /v2/api/site/{site}/topology | GET | Get Site Topology | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| Users | /v2/api/site/{site}/user-records | GET | List User Records | 
|  | /v2/api/site/{site}/user-records | POST | Create User Record | 
|  | /v2/api/site/{site}/user-records/{id} | GET | Get User Record Details | 
|  | /v2/api/site/{site}/user-records/{id} | PUT | Update User Record | 
|  | /v2/api/site/{site}/user-records/{id} | DELETE | Delete User Record | 
|  | /v2/api/site/{site}/user-groups | GET | List User Groups | 
|  | /v2/api/site/{site}/user-groups | POST | Create User Group | 
|  | /v2/api/site/{site}/user-groups/{id} | GET | Get User Group Details | 
|  | /v2/api/site/{site}/user-groups/{id} | PUT | Update User Group | 
|  | /v2/api/site/{site}/user-groups/{id} | DELETE | Delete User Group | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| VPN | /v2/api/site/{site}/vpn | GET | Get VPN Configuration | 
|  | /v2/api/site/{site}/vpn | PUT | Update VPN Configuration | 
| Category | Endpoint | Method | Description | 
|---|---|---|---|
| WireGuard | /v2/api/site/{site}/wireguard-users | GET | List WireGuard Users | 
|  | /v2/api/site/{site}/wireguard-users | POST | Create WireGuard User | 
|  | /v2/api/site/{site}/wireguard-users/{id} | GET | Get WireGuard User Details | 
|  | /v2/api/site/{site}/wireguard-users/{id} | PUT | Update WireGuard User | 
|  | /v2/api/site/{site}/wireguard-users/{id} | DELETE | Delete WireGuard User | 
Use the OpenAPI specification to generate client SDKs:
# Generate Python client
openapi-generator-cli generate -i openapi/openapi.yaml -g python -o clients/python
# Generate JavaScript client
openapi-generator-cli generate -i openapi/openapi.yaml -g javascript -o clients/javascript
# Generate Go client
openapi-generator-cli generate -i openapi/openapi.yaml -g go -o clients/go
Use tools like Postman or Insomnia with the OpenAPI spec for interactive testing:
- Import openapi/openapi.yaml into your API testing tool
- Configure authentication (API key or session cookies)
- Test endpoints against your UniFi Controller
This project is licensed under the MIT License - see the LICENSE file for details.
Disclaimer: This is an unofficial, community-maintained documentation project. Ubiquiti does not officially support or endorse this API documentation. Use at your own risk and always test thoroughly in development environments before production deployment.
