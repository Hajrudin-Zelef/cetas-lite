---
id: collect-261001-meraki/meraki/r-meraki-comments-b84y3o-mx-sitetosite-and-client-vpn-issues-ca9e0652
title: "r-meraki-comments-b84y3o-mx-sitetosite-and-client-vpn-issues-ca9e0652"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-b84y3o-mx-sitetosite-and-client-vpn-issues-ca9e0652.md
source_anchor: ""
source_lines: [1, 40]
sha256: d8b87e5e5986dbaeb5a197341c6880af307508735c507208fdee49c189b208f5
---

# r-meraki-comments-b84y3o-mx-sitetosite-and-client-vpn-issues-ca9e0652

MX site-to-site and client vpn issues 
        
    Hello, recently got several MX64's for our offices. I have set up a template for all of them to be based on, assigned them IPs, provided static IPs on our servers, and set up RADIUS for them.
When I am connected physically to the MX at our main office, I can view all of our shared files and folders. Once I attempt to access from behind an MX at a remote location, I cannot see the files or folders.
The Client VPN connects successfully, but then says "No network access" when I hover over the connection. When using the client vpn connection I cannot ping the MX, or view our shared files and folders. I'm not seeing myself show up in the Clients on the MX, or seeing the logon Audit Success or failure on the DC.
Using the default allow any any rules for the moment until I can get a solid connection going.
We are in the process of migrating from ASA 5505's the the MX's.
Any troubleshooting steps you can provide?
I'm not seeing any error codes in Event Viewer on the server
EDIT 1
Ok, I placed the MX in my home lab in the DMZ. That is letting me RDP from the main site into my home computer. I can ping the home lab mx from the main site. Can't ping the main MX from the home lab, because apparently I'm an A-Hole and allowed the computer to shut down mid-process. Whoo Boy!
However, before I did that, I still wasn't able to ping the main MX from the home lab.
EDIT 2
I CAN ping the main sites mx from the home lab, and the home labs MX from the main site, but still can't ping the main sites server.
Final Edit*
Got it working thanks to all of your suggestions.
Section des commentaires
Define "see". Are you trying to browse the network through Windows Explorer? Is the auto-vpn enabled for the template in the dashboard? Are the VPNs established? Can you ping across by name to the main site? Can you ping across by IP?
They are set up in a hub vpn on the template, the office subnet is set to use VPN as is the client, with automatic NAT traversal.
While the VPN's show as established in the VPN status page for the devices, I cannot ping across by name or IP address.
By "see" I mean I cannot navigate to \server\share in Windows Explorer. I get the "Windows cannot access \server\share" network error
Are there any other routers involved in the topology? Or layer-3 switches doing inter-vlan routing? Are the clients at the main site using the MX as their gateway? If you're not getting IP connectivity across, don't bother troubleshooting file shares or anything else.
If the MX doesn't see your client, that sounds like a problem. Is this at the remote site you're referring to? Can you ping the MX from the client on the same network?
Did you allow access to the local subnet on the server side MX?
I believe so. Pretend like you don't believe me :) Where would I find those settings to double check.
So a site to site vpn using Meraki’s in the same Org they are not setup as a non-Meraki peer correct?
Also, are you attempting from behind the home lab Meraki to use the client vpn? If the home lab Meraki is configured as a site to site peer with the office and it shows connected you do not need to use the client vpn.
What are is the home lab network scope?
Office network scope?
Are you using any vlans?
Try to ping the office Meraki from your lab Meraki without using the client vpn.
Does your WAN address change when you connect to the remote site?
Would this be my public IP address?
That does not change. It isn't changing while using the Cisco VPN Client either.
All of the remote sites have a static ip address assigned from the ISP, except for my home lab, where most of the testing is being done.
Right now the main MX is not using the static IP we have, because the ASA is still using it while we get everything set up.
We are in the same process, I have 1 mx100 , 2 mx 84, and about soon to be over 30 mx 64 for remote sites. We currently running the mx100 as a hub and MX 64 as spokes. Mx84 will be we hubs for certain region locations but right now the are spokes until we migrate those sights. The 3 main mx will be hubs due to having domain controllers.
When setting up plans make sure to advertise the DNS in the mx that house the controllers.
Also you might have to advertise in the asa where to route the specific route that you want to go to the main mx at your location.
I use ipchicken site to find what my want ip on the inside, also should be in your asa
