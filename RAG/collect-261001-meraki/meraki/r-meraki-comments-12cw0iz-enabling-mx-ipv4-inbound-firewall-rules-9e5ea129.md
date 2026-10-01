---
id: collect-261001-meraki/meraki/r-meraki-comments-12cw0iz-enabling-mx-ipv4-inbound-firewall-rules-9e5ea129
title: "r-meraki-comments-12cw0iz-enabling-mx-ipv4-inbound-firewall-rules-9e5ea129"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-12cw0iz-enabling-mx-ipv4-inbound-firewall-rules-9e5ea129.md
source_anchor: ""
source_lines: [1, 42]
sha256: c85b2739d1a0a6e61f4e79ad97cbcef7ec0e235776c0a91a204dc493122d1e52
---

# r-meraki-comments-12cw0iz-enabling-mx-ipv4-inbound-firewall-rules-9e5ea129

Enabling MX IPv4 Inbound Firewall Rules 
        
        
        
    
    
    We are currently undergoing an external vulnerability assessment, and the third party has asked that we whitelist their scanner IPs so they can effectively scan our external IP ranges. By default, this traffic is blocked by the Meraki's inbound deny all rule.
Googling indicated that asking support to enable the inbound firewall rule module would be pretty straightforward, however, I'm being told by support that I would also HAVE to enable No-NAT in order to get the inbound rule functionality. No-NAT is still in beta, and is not something we would want enabled. Has anyone previously experienced a similar situation?
Section des commentaires
Whitelist for what? If they’re doing a scan from the outside in then that scan only applies to any port forwards or services open to the internet. If there aren’t any open services then that test isn’t applicable.
Internal testing is different. They can use a third party or company system within the network for internal pen testing.
To me it sounds like this third party doesn’t fully understand what they’re doing either..
That's the thing, there are open services though - none of our 1:1 NAT or 1:Many NAT'd services are showing up when the scans are being performed. Our syslog shows everything from their IP's being blocked.
Is the IP they're scanning from added to the 'allowed inbound connections' list in each 1:1/1:Many rule?
Agreed with the other commenter. The third party has no idea what they're talking about.
If they're testing your outside interface, then we're only testing vulnerabilities on the edge itself (e.g. port 500 for IPsec) or port forwards/NAT forwarding rules toward internal servers.
And if the port forward(s) and 1:1 NAT(s) are locked down to specific remote IPs, then you're protected. They could be checking the internal device for vulnerabilities from the outside, but unless the remote IP's network gets compromised, it's a useless test.
I'm also experiencing this and it makes no sense to me why would we would open up access. Does anyone have a clue of the basis behind this?
This is what they say for reference:
Sysnet access
In order to run the scan, we need you to grant access to the IP addresses listed below.
If you use security software such as a firewall in your organization, you may need to white-list the below addresses in order for the scan to run successfully. Otherwise, you may block access to the scan, meaning it will fail. This will result in you being unable to successfully report your compliance.
If you are unsure how to do this, consult the help section of your firewall or contact your internet service provider for assistance.
What is an IP address?
An IP address is a series of numbers and dots that is your address on the internet. We need the correct address for your internet connection, to allow us to scan the correct connection – otherwise, we may scan someone else’s network.
Dynamic IP addresses
Some internet service providers will assign you a “Dynamic IP address.” This is an IP address that changes every time you connect and disconnect your internet router.
If you have a dynamic IP address, you need to update us with this new number every time you run your scan. This allows us to scan the correct connection.
If you are unsure as to whether you have a dynamic IP address, please contact your internet service provider who will be able to advise you. If you do have a dynamic IP, it’s advisable to refrain from scheduling scans in advance, as your IP address may have changed by the time the scheduled scan runs.
64.39.96.0/20
154.59.121.0/24
139.87.104.123/32
139.87.117.66/32
139.87.112.0/23
141.144.196.156/32
158.101.209.126/32
I'm seeing this exact same message including the IPs so assuming we're doing the same PCI scan. Did you end up doing anything or just have them run the scan?
what I found out is the reason why they’re asking for the scan in the first place is because the preliminary questions were answered incorrectly so if you go back through the questionnaire you most likely don’t even have to run the scan in the first place
We just went through our scans for external - I don't open or whitelist anything. If nothing is allowed in except for stateful packets you'll be all set.
Just read you do have 1:1 NAT enabled - are you only allowing certain ip addresses access to the NAT'd services? If so you will need to add the ip address their scan is coming from to that list.
Hey, can you open up these closed ports on the firewall so we can call them vulnerabilities? Lol. Umm, no… if you are doing external pen testing, the point is to not pre-enable anything; they should pen test against your current setup. Go home third party, you’re drunk and unqualified.
Had this with a few third parties, can you open up your routers so we can pen test them. FWIW most of them are happy if you change ‘drop’ to reject so they get a ‘no’ response. Of course this isn’t an option on Meraki.
