---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f-3
title: "how-to-add-adguard-to-opnsense-html-bbe7998f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f.md
source_anchor: ""
source_lines: [98, 193]
sha256: 47d35d45c68457605a558b28ff33bdbe8cf3613a8e6e8b110a2308893e5dc605
---

# how-to-add-adguard-to-opnsense-html-bbe7998f

- Enable the DNS filtering feature to block malicious and advertising domains.
- Configure Filters by selecting predefined blocks such as privacy, ad, or social media filtering, or add custom filters if needed.
Setting Up DNS Forwarding
- Return to the OPNsense web interface.
- Navigate to Services > DNS Resolver or DNS Forwarder, depending on your setup.
- Set the DNS server to point to the local AdGuard instance, typically 127.0.0.1 or your router’s IP address.
- Ensure that DNS traffic is directed through AdGuard to enforce filtering for all network clients.
Configuring Firewall Rules
- Go to Firewall > Rules.
- Create a new rule allowing DNS (port 53) traffic to the AdGuard server.
- Place this rule above existing DNS or DHCP rules to ensure it takes precedence.
- Apply changes and save the configuration.
Final Check
Test your setup by visiting a known advertising or malicious website. Confirm that AdGuard is blocking content as configured. Review logs within the AdGuard dashboard to verify activity and effectiveness.
Outdated Drivers Are Slowing You Down
One free scan finds every outdated or missing driver and matches the right update for your exact hardware.Free scan · exact hardware match
Windows Errors? Fix Them Before They Spread
Repair common Windows errors and clear accumulated junk for a smoother, more stable PC - no reinstall needed.Free scan · no reinstall
Configuring OPNsense to Use AdGuard as DNS Resolver
Integrating AdGuard DNS with OPNsense enhances your network’s privacy, ad-blocking capabilities, and security. Follow these steps to set up AdGuard as your DNS resolver on OPNsense.
Step 1: Access OPNsense Web Interface
Log in to your OPNsense dashboard using your administrator credentials. Ensure you have administrative privileges to modify DNS settings.
Step 2: Configure DNS Server Settings
- Navigate to Services > Unbound DNS > General.
- Uncheck Enable Forwarding Mode if it’s enabled, to prevent conflicts with custom DNS setups.
- In the Custom Options field, add the following configuration to forward DNS queries to AdGuard:
server:
    forward-zone:
        name: "."
        forward-addr: 94.140.14.14 # AdGuard DNS primary
        forward-addr: 94.140.15.15 # AdGuard DNS secondary
Step 3: Apply and Save Settings
Click Save at the bottom of the page, then click Apply Changes. This will direct DNS queries to AdGuard servers, leveraging their filtering capabilities.
Step 4: Configure DHCP and DNS Forwarding
- Navigate to Services > DHCPv4 (or DHCPv6 if applicable).
- Ensure your DHCP server hands out the IP address of your OPNsense device as the primary DNS server.
- Alternatively, under System > General, set DNS servers to your OPNsense’s LAN IP to enforce DNS resolution through your configured resolver.
Step 5: Verify the Setup
Test your DNS resolution by browsing to a site known for ads or using tools like nslookup or dig. Confirm that DNS queries are routed through AdGuard and that ads are being blocked.
The Tool Desk
Outbyte Driver Updater FREEScan for outdated or missing drivers - takes under a minuteDriver Scan →Outbyte PC Repair FREERepair Windows errors before they cause bigger problemsFix Now →
By following these steps, OPNsense now uses AdGuard as its DNS resolver, providing enhanced ad-blocking and privacy features for your network.
Setting Up DNS Forwarder/Resolver for AdGuard in OPNsense
Integrating AdGuard Home with OPNsense requires configuring your DNS settings to route queries through the AdGuard resolver. This setup ensures network-wide ad blocking and enhanced privacy. Follow these steps to add AdGuard to your OPNsense system effectively.
1. Install AdGuard Home
First, deploy AdGuard Home on a suitable device or VM within your network. Note the IP address and port it listens on, typically 192.168.x.x:3000. Ensure AdGuard is operational before proceeding.
2. Access OPNsense DNS Settings
Log into your OPNsense dashboard. Navigate to Services > DNS Resolver or DNS Forwarder, depending on your setup. Most modern OPNsense installations use the DNS Resolver (Unbound), so select that option if available.
3. Configure DNS Resolver
- Enable the DNS Resolver if not already active.
- Scroll to the General Settings section.
- In the Network Interfaces field, select the interfaces you want to use with DNS resolution.
- Disable Network Scope if you want DNS resolution to be available network-wide.
4. Add Forwarding to AdGuard
Locate the Custom options or Advanced Configuration section. Enter the following line to forward DNS queries to AdGuard:
<forward-zone>
  <name>.</name>
  <forward-addr>192.168.x.x#53</forward-addr>
</forward-zone>
Replace 192.168.x.x with the IP address of your AdGuard Home instance. Ensure the port (#53) matches AdGuard’s DNS port.
5. Save and Restart DNS Resolver
Click Save and then restart the DNS Resolver service to apply changes. Navigate to Services > Restart DNS Resolver.
6. Verify Configuration
Test DNS resolution on a client device by querying a domain. Confirm that ads are being blocked and that DNS queries are passing through AdGuard. Use tools like nslookup or dig to verify resolution points to your AdGuard server.
By following these steps, you effectively integrate AdGuard with OPNsense, creating a robust ad-blocking and privacy-enhancing network environment.
Adjusting Firewall Rules to Add AdGuard in OPNsense
Integrating AdGuard with OPNsense requires configuring firewall rules to allow DNS traffic through the filter. Properly adjusting these rules ensures that AdGuard can effectively block unwanted content while maintaining network security.
Step 1: Access Firewall Rules
- Log into your OPNsense dashboard.
- Navigate to Firewall > Rules.
- Select the interface where your DNS traffic will pass (typically LAN).
Step 2: Create a New Rule for DNS Traffic
- Click on the Add button to create a new rule.
- Set Action to Pass.
- Set Interface to your LAN or relevant interface.
- Configure Source as LAN net.
- Set Destination to Single host or alias.
- Enter the IP address of your AdGuard DNS server (e.g., 192.168.1.10).
- Set Destination port range to 53 (standard DNS port).
- Save the rule and ensure it’s ordered appropriately — usually above any blocking rules that might interfere.
Step 3: Allow DNS over HTTPS or DNS over TLS (Optional)
If using DNS over HTTPS or DNS over TLS, create rules permitting outbound traffic on ports 443 or 853, respectively. These rules should specify the AdGuard server or proxy IPs.
Step 4: Apply and Test the Configuration
- Click Apply Changes.
- Test DNS resolution from a client device to confirm AdGuard filters the traffic properly.
- Verify that DNS requests are directed to your AdGuard server and that unwanted content is blocked accordingly.
Adjusting firewall rules accurately is crucial for seamless AdGuard integration. Always verify the rules to prevent unintentional network disruptions and ensure effective filtering.
Ensuring Proper DNS Traffic Flow When Adding AdGuard to OPNsense
Integrating AdGuard Home with OPNsense enhances your network’s ad blocking and privacy capabilities. However, to maximize its effectiveness, you must configure your DNS traffic correctly. Proper DNS traffic flow ensures that all client devices resolve DNS queries through AdGuard without leaks or conflicts.
Step 1: Set AdGuard as the Primary DNS Server
- Navigate to Services > DNS Forwarder or DNS Resolver in OPNsense.
- Configure the DNS service to point all DNS queries to the local AdGuard instance. Typically, this is 127.0.0.1 or the IP address where AdGuard is listening.
- Disable any conflicting DNS services that might override this setting, such as the default DNS Forwarder or Resolver.
Step 2: Prevent DNS Leaks
- Ensure that DHCP settings hand out the AdGuard server’s IP as the primary DNS to all clients. This can be found under Services > DHCPv4.
- Configure the DHCP options to assign the AdGuard IP, or set static DNS entries on individual devices.
