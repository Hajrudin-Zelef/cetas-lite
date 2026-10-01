---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-18mvjsl-unable-to-connect-to-ssl-vpn-stops-at-40-6a6b84fb
title: "r-fortinet-comments-18mvjsl-unable-to-connect-to-ssl-vpn-stops-at-40-6a6b84fb"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-18mvjsl-unable-to-connect-to-ssl-vpn-stops-at-40-6a6b84fb.md
source_anchor: ""
source_lines: [1, 69]
sha256: 617fbc9a202652e87eb2eb06714761a75660957f3236f3c4486251c9a7fdc0a8
---

# r-fortinet-comments-18mvjsl-unable-to-connect-to-ssl-vpn-stops-at-40-6a6b84fb

Unable to connect to SSL VPN - Stops at 40% "The VPN Server May be Unreachable (-5)" 
        
        
        
    
    
    Strange one. I have configured SSL VPN in the past and believe I have followed the same steps however am stuck.
As per the title not connecting and getting stuck on 40% and generic "Unable to establish the VPN connection" error. I am also unable to connect via the web portal either using the IP and port in place.
- 
      I have configured the SSL-VPN Settings. It's listening on the correct WAN port and I have a custom PORT.
- 
      I have allowed "access from any host"
- 
      I have set "All Other Users / Groups" to the full-access Portal.
- 
      I have configured SSL-VPN Portal for "full-access" and all looks to be correct.
- 
      I have created a Firewall Policy allowing traffic from the SSL-VPN tunnel interface to the Internal interface. I have added the SSL_VPN_TUNNEL_ADDR1 and a group called VPNAccess as the source which has a number of users in it.
      For troubleshooting I have enabled debugging and I see no messages when attempting to connect: diagnose debug application sslvpn -1
diagnose debug enable
    
Forticlient is configured with the correct IP and port details of the external IP of the FW.
Couple of weird things I've noticed. On the log files on Forticlient I can see it has the FGTSERIAL \ DEVID entry as a different one to the actually firewall which is strange. I have no idea where is might be picking up this other Serial \ DEVID from unless this is a serial from Forticlient Instead which I'm not aware of.
We are on firmware version 6.4.14 which I believe is the current highest at this branch.
Unsure how to troubleshoot further on this. I'm raising a ticket with Fortinet but also thought I'd raise on here as well. I'm not aware of needing to do anything else on the FW to allow external access however 🤷♂️
Any ideas on how to further troubleshoot? Am I missing something stupid in term of Firewall Policies and allowing access - the fact we can't even browse the web portal is a weird one however I have double checked and the IP address matches the WAN interface.
Section des commentaires
You know there is an awesome debug log option with export in Forticlient that tell you very detailed what is going on :)
Clear the log, set it on debug and retry.
Most likely it will show up what is the issue.
"Couple of weird things I've noticed. On the log files on Forticlient I can see it has the FGTSERIAL \ DEVID entry as a different one to the actually firewall which is strange."
ok.....
https://community.fortinet.com/t5/FortiGate/Troubleshooting-Tip-Possible-reasons-for-FortiClient-SSL-VPN/ta-p/211965
40%-ish also relates around certificate mismatch/interception.
Did you change/refresh certs perhaps ?
I faffed around with a cert for SAML but we never implemented it. Potentially something to do with that. I’ll check the certs tomorrow.
Too lazy to look at comments but mine pauses at 40% for a bad SSL certificate. On windows the pop up comes to allow connection but doesn't become an active window. You have to find it (probably in your task bar) and click yes, then it will proceed. 45% is the 2fa pause.
This is your answer. The bad cert pop up is infact a pop behind. I normally windows and d followed by opening the client back up and moving it around a bit to find the window behind.
40% and 48% typically means there is not a portal for the user, and not a FW rule in place or the FW rule is not configured properly.
I would start with a
diag sniffer packet any "host (wan/vpn ip) ((or the client's ip) and icmp" 4 0 1
example:
Client IP = 1.1.1.1
WAN/VPN IP= 2.2.2.2
diag sniffer packet any "host2.2.2.2and icmp" 4 0 1diag sniffer packet any "host1.1.1.1and icmp" 4 0 1
ping the FGT and see if you can determine the traffic is making it to the FGT.
How are you handling users? Is this LDAP or local, Azure?
Create a new portal mapping to the user/group to "Full Access" Change the FW rule to include the user/group if it not. Report back results.
Thanks for the post.
The sniffer on the client device IP shows nothing.
The sniffer on the WAN IP shows a couple of items attempting to come in on ICMP but nothing from the client device in question.
I have simplified the user side as much as possible. Using local user. I have created a custom portal and still no luck. Checking the local log on the forticlient shows nothing outside of "SSLVPN tunnel connection failed".
Perhaps a silly question, but you are using the external IP of the client, right, not the internal? You could try and change ICMP in the filter for 'tcp port xxx' where xxx is your custom port.. you can then either test or run a telnet from the client: telnet <WAN IP> xxx
Does your ISP account include any 'security packages'. Comcast does that and it interferes with some encryted traffic.
Not that I’m aware of. Tethering from my phone and using o2.
The fortigate end
Parental Controls from TalkTalk in the UK cause this.
Not that I’m aware of. Tethering from my phone and using o2.
Are you using the forticlient 6.4 and not the 7.x version. Seen some similar causes
Using 7.2.2.0864
Perhaps you can download a older version of it. Also seen some crazy dns bug with these forticlient incompatible versions
Check if packets are even arriving on the WAN port of the FortiGate.
Try to traceroute first to the WAN IP.
Try to visit the SSLVPN portal using a web browser.
If it is disabled turn it on temporarily and test.
Try to visit https://VPNIP/remote/info in browser and see if it is reachable.
Look for any popups on FortiClient to accept Invalid certificate for FortiGate SSLVPN.
Turn off DTLS if it is enabled or turn it on if disabled.
Try logging in using another account.
