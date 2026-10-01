---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a-2
title: "en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a.md
source_anchor: ""
source_lines: [57, 85]
sha256: a58047233427d392e3776c9409ebe9407e2beb934edbc1e55f13459a01d19aea
---

# en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a

DHCP Option 43 is useful when several devices in the same scope should receive the controller address and the DNS name unifi is not used. If Sophos Firewall issues the leases itself, edit the relevant DHCPv4 server under Network > DHCP and create a custom option in DHCP options:
- Code: 43
- Type: string
- Value: Hexadecimal value of the controller IP
For the controller IP 192.168.3.10, the value used by Ubiquiti is:
0104c0a8030a
Enter the hexadecimal value without spaces. After saving, the test device must renew its lease. With DHCP Relay, the option belongs on the server that actually issues the lease. An FQDN-based Inform URL uses a different format with the prefix 02, a length byte, and a hex-encoded URL; a plain-text host name in the field is not equivalent. If the device ignores the option, first inspect a DHCP Offer or ACK and confirm that option 43 contains the expected value.
Configure DHCP Options on Sophos Firewall explains the general workflow, data types, and CLI fallbacks.
Roll back the changes without a factory reset
Before making changes, record the current Inform URL, any existing unifi DNS entry, DHCP Option 43, and the affected Sophos objects and rules with their values. Do not overwrite existing objects; use the dedicated names shown above for the test. This preserves a way back without resetting the device.
- If adoption has not yet been confirmed, undo the network changes using the following steps first; the device remains in its previous state. After adoption is complete, another set-inform can change the destination address, but it does not reliably restore the old controller’s configuration and credentials. If state-preserving return is required, use only a previously verified controller backup or the migration method supported by the controller vendor. Do not start adoption without that return path.
- Restore the previous unifi DNS entry and the old DHCP Option43 value exactly. If both were created only for this test, remove them instead. Then renew a lease and verify the DNS response.
- Disable the new UniFi devices to controller firewall rule first. Once the old management path and the intended name resolution work again, delete the rule.
- Delete UniFi-Inform ,UniFi-STUN , andUniFi_Controller only if they were created for this guide and SFOS shows no other references. Leave shared or pre-existing objects in place.
A Sophos Firewall configuration backup is useful additional protection, but it does not replace this inventory: restoring a full backup can overwrite unrelated changes made in the meantime.
Troubleshoot failed adoption
Troubleshooting should begin with the first state that differs from expectations:
- No IP address: Check the DHCP scope, VLAN assignment, and switch port.
- unifi does not resolve: Check which DNS server DHCP distributes and whether DNS is allowed for the device zone underAdministration > Device access .
- No connection to TCP 8080: Check the source zone, source network, destination FQDN, rule order, NAT, and any SD-WAN or policy routing.
- The device appears, but adoption stalls: Run the same set-inform command again and make sure TCP8080 remains reachable.
- Server Reject orManaged by Other : The device may still be assigned to another UniFi Host. If possible, restore the previous controller or remove the device from it cleanly.
- Controller behind a remote gateway: Check TCP 8080 port forwarding, the VPN path, and Double NAT.
Log Viewer and Packet Capture show which Rule ID, destination IP, and service are actually used. The process is described in Test Sophos Firewall rules with Log Viewer and Packet Capture.
Reset only as a last resort
If the previous UniFi application is accessible, Manage > Forget is the cleanest method: the device is removed from the application and reset to factory defaults. A physical reset alone does not remove the old entry from that application.
Without access to the previous host, a factory reset may be necessary. The device must remain powered on. Depending on the model, hold the Reset button for approximately 5 to 10 seconds or until the LED confirms the restore. Pressing it too briefly may only restart the device; holding it too long can trigger TFTP Recovery Mode.
Older runbooks mention the SSH command set-default for this purpose. Current general reset instructions do not document it as a cross-device standard, so it is not recommended here as a safe universal command.
After the reset, adoption starts again with DHCP, DNS, and TCP 8080. If these prerequisites are wrong, a factory reset will not solve the underlying network problem.
