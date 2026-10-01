---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1bi10go-fortigate-ipsec-tunnel-phase-2-updown-since-4e4dd97e
title: "r-fortinet-comments-1bi10go-fortigate-ipsec-tunnel-phase-2-updown-since-4e4dd97e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1bi10go-fortigate-ipsec-tunnel-phase-2-updown-since-4e4dd97e.md
source_anchor: ""
source_lines: [1, 43]
sha256: 55eb642c3d68ce6ad3dcb4756e5945ef7e6d5a2eaf25a838396328e46b89b132
---

# r-fortinet-comments-1bi10go-fortigate-ipsec-tunnel-phase-2-updown-since-4e4dd97e

Fortigate IPSEC Tunnel Phase 2 up/down since upgrade 
        
        
        
    
    
    I believe when we upgraded 7.0.14 something broke with one of our tunnels. (Only one person uses it and only as necessary for a vendor...)
While one tunnel to another vendor still works fine, the other one phase 2 keeps install_sa - Negotiate every few minutes and no traffic is passing.
No firewall or other rules have changed, I've troubleshot with the vendors side and they claim no changes on their side, I'm not familiar enough with the debugging to be 100% sure I'm good to push them harder to check traffic on their side. When I try to ping destinations there I see the traffic hit/increase counters on the policy rule for the tunnel as well as the tunnel itself.
Any help would be appreciated.
Section des commentaires
Can you update from which firmware the FGT was upgraded to 7.0.14 before having this issue.
Apart from install_sa do you see any other logs such as tunnel-down in the "action" category.
Can you try flushing the tunnel on both ends to reestablish the tunnel, so the SAs can be re-established.
Also, If there are more than one subnets (both local and remote) configured over the IPsec VPN, there should be more than one phase2 selector configured instead of including multiple firewall addresses in a single firewall deal with group and defining it as a single phase2 selector
Ref:
https://community.fortinet.com/t5/FortiGate/Technical-Tip-IPsec-VPN-between-FortiGate-and-other-Vendor-with/ta-p/205118
To flush the tunnel:
diagnose vpn tunnel flush <my-phase1-name>
If the above doesn't work, kindly collect the below logs along with the latest config file and share it to sferoz@fortinet.com for further analysis.
Logs:
dia vpn tunnel list name xyz (xyz is the name of the tunnel)
diag vpn ike gateway list name xyz (xyz is the name of the tunnel)
When IPSEC is down, kindly run the IPSEC debug on the FGT side:
diag deb reset
diag vpn ike log-filter dst-addr4 x.x.x.x (x.x.x.x is the remote IP address)
diag debug application ike -1
diag debug console timestamp enable
diag debug enable
To disable the debug :
di de dis
di de reset
Thanks.
Thanks replied with info and a question!
We had this issue as well. The default DPD on demand setting needed to be changed. We set DPD on-idle on both sides and the tunnels are stable now.
https://docs.fortinet.com/document/fortigate/7.4.3/hardware-acceleration/636026/disabling-np-offloading-for-individual-ipsec-vpn-phase-1s give this a try!
moving all processing off the NPUs can create a lot of CPU overhead, just be aware.
Is the idea that the phase 2 wont be cycling if the fortigate isnt passing traffic back and forth between the CPU and NPU?
I also am having this problem starting 7.0.14
Check this: https://community.fortinet.com/t5/FortiGate/Technical-Tip-Dialup-IPSEC-issues-after-upgrading-7-2-6-7-0-13/ta-p/283504
Same issue as OP and also found this solution but sceptical to try this
We did the add-route disable only on the side where the dialup connect and it worked perfectly, resolved the issue
Been having this same issue since updating to 7.0.14, but using static IPSEC, not dialup. Haven't been able to find a solution. Had the tunnels already doing DPD on idle, TAC advised to switch to the default of on-demand but still issue persist. Random tunnels drop while none of the other tunnels do to the same interface.
