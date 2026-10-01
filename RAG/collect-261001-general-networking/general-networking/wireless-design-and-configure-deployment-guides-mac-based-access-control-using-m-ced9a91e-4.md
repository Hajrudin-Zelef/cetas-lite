---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e-4
title: "wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e.md
source_anchor: ""
source_lines: [154, 220]
sha256: 1b4878f1a2837e06d0c3468811266bf6d4a7c2bcacec14636041526d68e21b37
---

# wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e

  - Host IP or FQDN (the IP address or FQDN the access points will send RADIUS accounting messages to)
  - Auth port (the port on the RADIUS server that is listening for accounting messages; 1813 by default)
  - Secret (the shared key used to authenticate messages between the access points and RADIUS server)
- Click Done and then Save changes.
At this point, "Start" and "Stop" accounting messages will be sent from the access points to the RADIUS server whenever a client successfully connects or disconnects from the SSID, respectively.
Configuring WPA2-Enterprise with RADIUS using Cisco ISE
Cisco Meraki access points can be configured to provide enterprise WPA2 authentication for wireless networks using Cisco Identity Services Engine (ISE) as a RADIUS server. This article will cover instructions for basic integration with this platform. For more detailed information on how to configure Cisco ISE, please refer to the Cisco Identity Services Engine User Guide.
Prerequisites
- Cisco ISE installed and reachable from the access points
- An SSID configured to use WPA2-Enterprise pointing to the Cisco ISE server
Installing Server Certificates
After installation, Cisco ISE generates, by default, a self-signed local certificate and private key, and stores them on the server. This certificate will be used by default for WPA2-Enterprise. In a self-signed certificate, the hostname of Cisco ISE is used as the common name (CN) because it is required for HTTPS communication.
Note: Using a self-signed certificate is not recommended for RADIUS. In order to use the default self-signed cert, clients will need to have RADIUS server's identity validation disabled in order to connect. For certificate options on the RADIUS server you may refer to the RADIUS configuration section in this document.
Adding Managed Network Devices
- In Cisco ISE, choose Administration > Network Resources > Network Devices.
- From the Network Devices navigation pane on the left, click Network Devices.
- 
    Add, edit, or duplicate a device:
- Add: Select Add. Alternatively, click Add new device from the action icon on the Network Devices navigation pane.
- Edit: Select the check box next to a device and click Edit. Alternatively, click a device name from the list to edit.
- Duplicate: select the check box next to a device and click Duplicate.
4. In the right pane, enter the Name and IP Address.
5. Check the Authentication Settings check box and define a Shared Secret for RADIUS authentication. This must match the Secret entered for the RADIUS server when configuring the SSID in dashboard.
6. Click Submit.
Enabling Policy Sets
Cisco ISE supports policy sets (see Cisco ISE: Introduction to Policy Sets), which allows grouping sets of authentication and authorization policies, as opposed to the basic authentication and authorization policy model, which is a flat list of authentication and authorization rules. Policy sets allow for logically defining an organization's IT business use cases into policy groups or services, such as VPN and 802.1X. This makes configuration, deployment, and troubleshooting much easier.
- In Cisco ISE, choose Administration > System > Deployment > Settings > Policy Sets.
- Click the Default policy. The default policy is displayed in the right.
- Click the plus (+) sign on top and choose Create Above.
- Enter the Name, Description and a Condition for this group policy.
- Define the Authentication policy.
- Click Submit. After configuring a policy set, Cisco ISE will log out any administrators. Log in again to access the Admin portal.
Configuring an Authentication Policy
- In Cisco ISE, select the Actions menu and click Insert New Rule Above.
- Give the sub-rule a Name (Example: Dot1X).
- Click the small window icon to open the Conditions menu.
- Select Create New Condition (Advanced Option).
- Select Network Access > EAP Authentication.
- Leave the operator box set to EQUALS.
- In the last box select EAP-MSCHAPv2.
- In the Use field, select Active Directory as the identity store( see Managing External Identity Sources). Configure the Active Directory integration as appropriate for the desired deployment.
Static Client Exclusion Policies
MR 32.1.1 firmware introduces support for network administrators to block communication from clients during the 802.11 authentication phase attempting to connect to the wireless network based on listed mac addresses . Oftentimes clients either don't have the correct security policies installed or are using an incorrect cert/password for authentication. Since these client devices are trying to authenticate over and over again, this causes additional requests on the authentication server seen in the network event logs.
Static client exclusion policies allow network admins to block up to 100 MAC addresses per SSID, preventing all inbound client communication including pre and post authentication.
Note: In order to enable static client exclusion policies please contact Meraki support.
Configuration:
- 
    Navigate to Wireless > Firewall & Traffic Shaping in Dashboard.
- 
    Within the Client Exclusion section, enable the Client exclusion policies for the desired SSID.
- 
    In the provided text box, enter the MAC addresses you wish to block.
Note: You may add up to 100 MAC addresses to the static exclusion list.
API
The static exclusion status and static exclusion MAC list can be configured separately.
Static exclusion status
For the status, there are only index (GET) and update (PUT) endpoints
GET /organizations/{organizationId}/wireless/ssids/policies/clientExclusion/bySsid
GET optional parameters: networkIds, includeDisabledSsids, ssidNumbers
PUT /networks/{networkId}/wireless/ssids/{number}/policies/clientExclusion
Static exclusion MAC list
For the static exclusion MAC list, index (GET), bulkAdd (POST) and bulkRemove (POST) endpoints exist
GET /organizations/{organizationId}/wireless/ssids/policies/clientExclusion/static/exclusions/bySsid
GET optional parameters: networkIds, includeDisabledSsids, ssidNumbers
POST /networks/{networkId}/wireless/ssids/{number}/policies/clientExclusion/static/exclusions/bulkAdd
POST /networks/{networkId}/wireless/ssids/{number}/policies/clientExclusion/static/exclusions/bulkRemove
The API then writes the enable flag or the static exclusion list to Postgres.
