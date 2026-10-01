---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-kea-html-34416966-3
title: "manual-kea-html-34416966"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/manual-kea-html-34416966.md
source_anchor: ""
source_lines: [174, 316]
sha256: 7e939970a8ab0673cdd9652826fd503ae59ef9d30bdf6ee86e99a385f9157d51
---

# manual-kea-html-34416966

KEA DHCPs main strength is the ability to synchronize leases between multiple servers, which makes it ideal for medium to large HA setups (more than 1000 unique clients) where you cannot use Dnsmasq DHCP.
As example we configure a network with two KEA DHCP instances on a master and backup OPNsense.
To configure KEA with a minimal HA setup for LAN using the 192.168.1.0/24 network follow these steps:
- LAN Network:
  - CARP IPv4 address: 192.168.1.1/24
  - Master IPv4 address: 192.168.1.2/24
  - Backup IPv4 address: 192.168.1.3/24
Attention
All configuration must be done on the master, and afterwards synchronized to the backup via
- Go to :
| Option | Value | 
|---|---|
| Enabled | X | 
| Bind address | 127.0.0.1 | 
| Bind port | 8000 | 
- Press Apply then go to and follow through these tabs:
| Option | Value | 
|---|---|
| Service |  | 
| Enabled | X | 
| General settings |  | 
| Interfaces | LAN | 
| Firewall rules** | X | 
| High Availability |  | 
| Enabled | X | 
| This server name | (It is highly recommended to use the offered default value) | 
- Press Apply and go to Subnets
| Option | Value | 
|---|---|
| Subnet | 192.168.1.0/24 | 
| Pools | 192.168.1.100 - 192.168.1.199 | 
| DHCP option data |  | 
| Auto collect option data | (This must be unchecked for HA) | 
| Routers (gateway) | 192.168.1.1 (use the LAN CARP IP address) | 
| DNS servers | 192.168.1.1 (use the LAN CARP IP address) | 
- Press Save and go to HA Peers
- First entry:
| Option | Value | 
|---|---|
| Name | (Use the name that is displayed in the settings Tab for “This server name” on the master) | 
| Role | primary | 
| URL | http://192.168.1.2:8001/ (Use the LAN interface IP of the master, the port must be different than the control agent) | 
- Second entry:
| Option | Value | 
|---|---|
| Name | (Use the name that is displayed in the settings Tab for “This server name” on the backup) | 
| Role | standby | 
| URL | http://192.168.1.3:8001/ (Use the LAN interface IP of the backup, the port must be different than the control agent) | 
- Press Save and Apply
Now the initial configuration is finished, and we synchronize it with the backup server. Both servers will always share the exact same configuration.
Go to and ensure that KEA is selected in Services to synchronize.
Then go to and press Synchronize and reconfigure all.
Immediately afterwards, KEA will be active on both master and backup, and a bidirectional lease synchronization will be configured.
DHCP Options
Each subnet and reservation has an DHCP option data list available. If Auto collect option data is enabled, some DHCP options like router, DNS server and system domain are added automatically. Additional fields can be filled out with other common options.
In cases where more advanced DHCP options need to be sent, you can use the Options tab found in and .
When adding a new option, you can enter matching and setting parameters:
When matching a DHCP option, a client class with a test is created. The set option will only be sent to clients that pass the test.
When setting a DHCP option, the payload will be sent unconditionally if no match exists in the same input mask.
To send a created option, attach it to a reservation or subnet in their respective tabs with the available Options dropdown menu.
Combining both set and match enables you to create multiple options with the same code, but different payloads. A common example is matching based on client architecture and sending a specific boot file as payload:
- Go to and follow through these tabs:
Create an option for BIOS boot:
| Option | Description | 
|---|---|
| Description | option-bios-bootfile | 
| Match DHCP option |  | 
| Match Code | client-system [93] | 
| Match Encoding | uint16 | 
| Match Data | 0 | 
| Set DHCP option |  | 
| Set Code | bootfile-name [67] | 
| Set Encoding | string | 
| Set Data | undionly.kpxe | 
Create an option for EFI boot:
| Option | Description | 
|---|---|
| Description | option-efi-bootfile | 
| Match DHCP option |  | 
| Match Code | client-system [93] | 
| Match Encoding | uint16 | 
| Match Data | 7 | 
| Set DHCP option |  | 
| Set Code | bootfile-name [67] | 
| Set Encoding | string | 
| Set Data | snponly.efi | 
- Press Save and go to Subnets
Select an available subnet, and add the Options you created:
| Option | Value | 
|---|---|
| Options | option-bios-bootfile ,option-efi-bootfile | 
- Press Save and Apply
With this configuration, any client that sends client-system [93] containing the value 0 will be provided with bootfile-name [67] and undionly.kpxe.
The same logic applies to the efi bootfile.
Note
Matching is optional, leave it empty to send the option out to any client in the subnet it is attached to.
Tip
Any option can be sent as user defined hex. This helps for structured and encapsulated options that may have multiple types or are binary blobs.
A common example is vendor specific [43], which is used for vendor specific information.
Just as with the bootfiles example, if you match client specific option codes, you can send out different vendor specific option codes in the same subnet.
Dynamic DNS (RFC2136)
KEA allows registering client FQDNs via dynamic DNS (RFC2136) to an authoritative DNS server.
Such an authoritative DNS server will be ISC BIND or an alternative like PowerDNS. Recursive DNS servers like Dnsmasq or Unbound are not able to fulfill this role.
Tip
The OPNsense Business Edition includes Authoritative DNS with RFC2136 support.
When clients register their IP address, the DHCP server will receive a Client FQDN (DHCP option 81) that either contains a client hostname or an FQDN. In cases where clients only send a hostname, using the DNS qualifying suffix will construct an FQDN and force an update anyway.
Attention
The client is responsible to send the Dynamic DNS update request via DHCP option 81. Only with this payload, the hostname will be registered in a forward zone. Clients that do not send any hostname cannot be registered, the administrator must ensure all of their devices have unique hostnames configured.
As an example setup, we have configured a zone like this in ISC BIND. The example taken from the KEA DDNS documentation:
:
key "key.four.example.com." {
    algorithm hmac-sha224;
    secret "bZEG7Ow8OgAUPfLWV3aAUQ==";
};
:
To configure the forward zone for a DHCPv4 range, go to and select a subnet:
| Option | Value | 
|---|---|
| Subnet | 192.168.1.0/24 | 
| Pools | 192.168.1.100 - 192.168.1.199 | 
| DHCP option data |  | 
| Auto collect option data | (This must be unchecked) | 
| Routers (gateway) | 192.168.1.1 | 
| DNS servers | 192.168.1.1 | 
| Domain name | four.example.com | 
| Dynamic DNS |  | 
| DNS forward zone | four.example.com. | 
| DNS qualifying suffix | four.example.com. (optional, use if your clients do not send an FQDN via DHCP option 81) | 
| DNS server | 203.0.113.1 | 
| TSIG key name | key.four.example.com. | 
| TSIG key secret | bZEG7Ow8OgAUPfLWV3aAUQ== | 
| TSIG key algorithm | hmac-sha224 | 
Next, enable the KEA DDNS Agent. Go to :
| Option | Value | 
|---|---|
| Enabled | X | 
| Bind address | 127.0.0.1 | 
| Bind port | 53001 | 
After applying the configuration, the DHCP servers construct DDNS update requests, known as NameChangeRequests (NCRs), based on DHCP lease change events and then post them to the DDNS Agent. The DDNS Agent attempts to match each request to the appropriate DNS server and carries out the necessary conversation with those servers to update the DNS data.
Note
The TSIG key name must be unique per DNS forward zone. If you configure multiple subnets with an identical DNS forward zone, but different TSIG key names and TSIG key secrets, only the first one will be taken into account. Best practice would be creating one unique DNS forward zone per subnet, each with a unique TSIG key name.
Attention
Only subnets that have a DNS server configured will send DDNS updates.
