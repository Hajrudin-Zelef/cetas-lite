---
id: collect-261001-general-networking/general-networking/send-feedback-38-1
title: "SD-WAN traffic shaping and QoS"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/send-feedback-38.md
source_anchor: ""
source_lines: [1, 171]
sha256: b348966e30db8d3202606b3393266d3290dfad880cbd2ffffc55dd74f787e548
---

# SD-WAN traffic shaping and QoS

# SD-WAN traffic shaping and QoS

Use a traffic shaper in a firewall shaping policy to control traffic flow. You can use it to control maximum and guaranteed bandwidth, or put certain traffic to one of the three different traffic priorities: high, medium, or low.

An advanced shaping policy can classify traffic into 30 groups. Use a shaping profile to define the percentage of the interface bandwidth that is allocated to each group. Each group of traffic is shaped to the assigned speed limit based on the outgoing bandwidth limit configured on the interface.

For more information, see Traffic shaping.

### Sample topology


### Sample configuration

This example shows a typical customer usage where the customer's SD-WAN uses the default zone, and has two member: wan1 and wan2, each set to 10Mb/s.

An overview of the procedures to configure SD-WAN traffic shaping and QoS with SD-WAN includes:

1. Give HTTP/HTTPS traffic high priority and give FTP low priority so that if there are conflicts, FortiGate will forward HTTP/HTTPS traffic first.
2. Even though FTP has low priority, configure FortiGate to give it a 1Mb/s guaranteed bandwidth on each SD-WAN member so that if there is no FTP traffic, other traffic can use all the bandwidth. If there is heavy FTP traffic, it can still be guaranteed a 1Mb/s bandwidth.
3. Traffic going to specific destinations such as a VOIP server uses wan1 to forward, and SD-WAN forwards with an Expedited Forwarding (EF) DSCP tag 101110.

###### To configure SD-WAN traffic shaping and QoS with SD-WAN in the GUI:

1. On the FortiGate, add wan1 and wan2 as SD-WAN members, then add a policy and static route.See SD-WAN quick start.
2. Add a firewall policy with *Application Control* enabled. See Security.
3. Go to *Policy & Objects > Traffic Shaping* , select the*Traffic Shaper* tab, and edit*low-priority* .
  1. Enable *Guaranteed Bandwidth* and set it to*1000* kbps.
4. Enable 
5. Go to *Policy & Objects > Traffic Shaping* , select the*Traffic Shaping Policy* tab, and click*Create New* .
  1. Name the traffic shaping policy, for example, *HTTP-HTTPS* .
  2. Set the following:Source *all*Destination *all*Service *HTTP* and*HTTPS*Outgoing interface *virtual-wan-link*Shared Shaper Enable and set to *high-priority*Reverse Shaper Enable and set to *high-priority*
  3. Click *OK* .
6. Name the traffic shaping policy, for example, 
7. Go to *Policy & Objects > Traffic Shaping* , select the*Traffic Shaping Policy* tab, and click*Create New* .
  1. Name the traffic shaping policy, for example, *FTP* .
  2. Set the following:Source *all*Destination *all*Service *FTP* ,*FTP_GET* , and*FTP_PUT*Outgoing interface *virtual-wan-link*Shared Shaper Enable and set to *low-priority*Reverse Shaper Enable and set to *low-priority*
  3. Click *OK*
8. Name the traffic shaping policy, for example, 
9. Go to *Network > SD-WAN* , select the*SD-WAN Rules* tab, and click*Create New* .
  1. Enter a name for the rule, such as *Internet* .
  2. In the *Destination* section, click*Address* and select the VoIP server that you created in the firewall address.
  3. Under *Outgoing Interfaces* select*Manual* .
  4. For *Interface preference* select*wan1* .
  5. Click *OK* .
10. Enter a name for the rule, such as 
11. Use CLI commands to modify DSCP settings. See the DSCP CLI commands below.

###### To configure the firewall policy using the CLI:

```
config firewall policy
    edit 1
        set name "1"
        set srcintf "dmz"
        set dstintf "virtual-wan-link"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set utm-status enable
        set ssl-ssh-profile "certificate-inspection"
        set application-list "default"
        set nat enable
    next
end
```
                                            ###### To configure the firewall traffic shaper priority using the CLI:

```
config firewall shaper traffic-shaper
    edit "high-priority"
        set maximum-bandwidth 1048576
        set per-policy enable
    next
    edit "low-priority"
        set guaranteed-bandwidth 1000
        set maximum-bandwidth 1048576
        set priority low
        set per-policy enable
    next
end
```
                                            ###### To configure the firewall traffic shaping policy using the CLI:

```
config firewall shaping-policy
    edit 1
        set name "http-https"
        set service "HTTP" "HTTPS"
        set dstintf "virtual-wan-link"
        set traffic-shaper "high-priority"
        set traffic-shaper-reverse "high-priority"
        set srcaddr "all"
        set dstaddr "all"
    next
    edit 2
        set name "FTP"
        set service "FTP" "FTP_GET" "FTP_PUT"
        set dstintf "virtual-wan-link"
        set traffic-shaper "low-priority"
        set traffic-shaper-reverse "low-priority"
        set srcaddr "all"
        set dstaddr "all"
    next
end
```
                                            ###### To configure SD-WAN traffic shaping and QoS with SD-WAN in the CLI:

```
config system sdwan
    set status enable
    config members
        edit 1
            set interface "wan1"
            set gateway 172.16.20.2
        next
        edit 2
            set interface "wan2"
            set gateway 10.100.20.2
        next
    end
    config service
        edit 1
            set name "SIP"
            set priority-members 1
            set dst "voip-server"
            set dscp-forward enable
            set dscp-forward-tag 101110
        next
    end
end
```
                                            If no SD-WAN zone is specified, members are added to the default *virtual-wan-link* zone.

###### To use the diagnose command to check if specific traffic is attached to the correct traffic shaper:

```
# diagnose firewall iprope list 100015
policy index=1 uuid_idx=0 action=accept
flag (0):
shapers: orig=high-priority(2/0/134217728) reply=high-priority(2/0/134217728)
cos_fwd=0  cos_rev=0
group=00100015 av=00000000 au=00000000 split=00000000
host=0 chk_client_info=0x0 app_list=0 ips_view=0
misc=0 dd_type=0 dd_mode=0
zone(1): 0 -> zone(2): 36 38
source(1): 0.0.0.0-255.255.255.255, uuid_idx=6,
dest(1): 0.0.0.0-255.255.255.255, uuid_idx=6,
service(2):
        [6:0x0:0/(1,65535)->(80,80)] helper:auto
        [6:0x0:0/(1,65535)->(443,443)] helper:auto
policy index=2 uuid_idx=0 action=accept
flag (0):
shapers: orig=low-priority(4/128000/134217728) reply=low-priority(4/128000/134217728)
cos_fwd=0  cos_rev=0
group=00100015 av=00000000 au=00000000 split=00000000
host=0 chk_client_info=0x0 app_list=0 ips_view=0
misc=0 dd_type=0 dd_mode=0
zone(1): 0 -> zone(2): 36 38
source(1): 0.0.0.0-255.255.255.255, uuid_idx=6,
dest(1): 0.0.0.0-255.255.255.255, uuid_idx=6,
service(3):
        [6:0x0:0/(1,65535)->(21,21)] helper:auto
        [6:0x0:0/(1,65535)->(21,21)] helper:auto
        [6:0x0:0/(1,65535)->(21,21)] helper:auto
```
                                            ###### To use the diagnose command to check if the correct traffic shaper is applied to the session:

