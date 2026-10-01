---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-administration-guide-256210-example-sd-wan-configuratio-a8158f9c-5
title: "Example SD-WAN configurations using ADVPN 2.0"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-administration-guide-256210-example-sd-wan-configuratio-a8158f9c.md
source_anchor: ""
source_lines: [545, 753]
sha256: 1c725f735a474c5de7338aa6ff37d350e9b12d817ff9575b60ad6266fd3bbf51
---

# Example SD-WAN configurations using ADVPN 2.0

```
**set advpn-select enable 
            set advpn-health-check "HUB"**
        next
    end
    config members
        edit 1
            set interface "H1_T11"
            set zone "overlay"
            **set transport-group 1** 
        next
        edit 2
            set interface "H1_T22"
            set zone "overlay"
            **set transport-group 1** 
        next
        edit 3
            set interface "H1_T33"
            set zone "overlay"
            **set transport-group 2** 
        next
    end
    config health-check
        edit "HUB"
            set server "172.31.100.100"
            set members 1 2 3
            config sla
                edit 1
                    set link-cost-factor latency
                    set latency-threshold 100
                next
            end
        next
    end
    config service
        edit 1
            set name "1"
            **set load-balance enable            
            set mode sla**
            set dst "CORP_LAN"
            set src "CORP_LAN"
            config sla
                edit "HUB"
                    set id 1
                next
            end
            **set priority-members 1 2 3**
        next
    end
end
config vpn ipsec phase1-interface
    edit "H1_T11"
        ...
        set idle-timeout enable
        **set shared-idle-timeout enable**
        set idle-timeoutinterval 5
        ...
    next
end
config vpn ipsec phase1-interface
    edit "H1_T22"
        ...
        set idle-timeout enable
        **set shared-idle-timeout enable**
        set idle-timeoutinterval 5                            
        ...
    next
end
config vpn ipsec phase1-interface
    edit "H1_T33"
        ...
        set idle-timeout enable
        **set shared-idle-timeout enable**
        set idle-timeoutinterval 5
        ...
    next
end # diagnose sys sdwan health-check 
Health Check(HUB):
Seq(1 H1_T11): state(alive), packet-loss(0.000%) latency(0.223), jitter(0.018), mos(4.404), bandwidth-up(999999), bandwidth-dw(999998), bandwidth-bi(1999997) **sla_map=0x1**
Seq(2 H1_T22): state(alive), packet-loss(0.000%) latency(0.191), jitter(0.009), mos(4.404), bandwidth-up(999993), bandwidth-dw(999998), bandwidth-bi(1999991) **sla_map=0x1**
Seq(3 H1_T33): state(alive), packet-loss(0.000%) latency(0.139), jitter(0.007), mos(4.404), bandwidth-up(999999), bandwidth-dw(999998), bandwidth-bi(1999997) **sla_map=0x1**

```
config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
        edit "overlay"
            
```
**set advpn-select enable 
            set advpn-health-check "HUB"**
        next
    end
    config members
        edit 1
            set interface "H1_T11"
            set zone "overlay"
            **set transport-group 1** 
        next
        edit 2
            set interface "H1_T22"
            set zone "overlay"
            **set transport-group 1** 
        next
        edit 3
            set interface "H1_T33"
            set zone "overlay"
            **set transport-group 2** 
        next
    end
    config health-check
        edit "HUB"
            set server "172.31.100.100"
            set members 3 1 2
            config sla
                edit 1
                    set link-cost-factor latency
                    set latency-threshold 100
                next
            end
        next
    end
end
config vpn ipsec phase1-interface
    edit "H1_T11"
        ...
        set idle-timeout enable
        **set shared-idle-timeout enable**
        set idle-timeoutinterval 5
        ...
    next
end
config vpn ipsec phase1-interface
    edit "H1_T22"
        ...
        set idle-timeout enable
        **set shared-idle-timeout enable**
        set idle-timeoutinterval 5                            
        ...
    next
end
config vpn ipsec phase1-interface
    edit "H1_T33"
        ...
        set idle-timeout enable
        **set shared-idle-timeout enable**
        set idle-timeoutinterval 5
        ...
    next
end
# diagnose sys sdwan health-check 
Health Check(HUB):
Seq(3 H1_T33): state(alive), packet-loss(0.000%) latency(0.148), jitter(0.021), mos(4.404), bandwidth-up(999999), bandwidth-dw(999998), bandwidth-bi(1999997) **sla_map=0x1**
Seq(1 H1_T11): state(alive), packet-loss(0.000%) latency(0.183), jitter(0.010), mos(4.404), bandwidth-up(999999), bandwidth-dw(999998), bandwidth-bi(1999997) **sla_map=0x1**
Seq(2 H1_T22): state(alive), packet-loss(0.000%) latency(0.163), jitter(0.005), mos(4.404), bandwidth-up(999994), bandwidth-dw(999998), bandwidth-bi(1999992) **sla_map=0x1**

In this scenario, PC1 connected to Spoke 1 initiates an ICMP ping destined for PC1 connected to Spoke 2. Therefore, this user traffic matches SD-WAN rule 1 and triggers shortcut path selection and establishment.

On Spoke 1, in the IKE debug (diagnose debug application ike -1), debug messages indicate that multiple direct shortcut-query packets are being sent to Spoke 2:

ike :VWL_ADVPN_MSG_T_TRIGGER
ike V=root:0 looking up shortcut by addr 172.31.80.2, resp-name:H1_T11, name H1_T22, peer-addr 172.31.3.101:0
ike V=root:0:H1_T22: send shortcut-query
...
ike :VWL_ADVPN_MSG_T_TRIGGER
ike V=root:0 looking up shortcut by addr 172.31.81.2, resp-name:H1_T22, name H1_T22, peer-addr 172.31.3.105:0
ike V=root:0:H1_T22: send shortcut-query
...
ike :VWL_ADVPN_MSG_T_TRIGGER
ike V=root:0 looking up shortcut by addr 172.31.82.2, resp-name:H1_T33, name H1_T33, peer-addr 172.31.4.101:0
ike V=root:0:H1_T33: send shortcut-query
...

From the diagnostic command on Spoke 1, observe that multiple shortcuts are triggered in **bold** based on the ADVPN 2.0 path management calculation where in-SLA overlays within the same transport group were selected.

```
Branch1_FGT# diagnose system sdwan service4
Service(1): Address Mode(IPV4) flags=0x24200 use-shortcut-sla use-shortcut
 Tie break: cfg
 Shortcut priority: 3
  Gen(69), TOS(0x0/0x0), Protocol(0): src(1->65535):dst(1->65535), Mode(sla  hash-mode=round-robin)
  Member sub interface(8):
    1: seq_num(1), interface(H1_T11):
       1: H1_T11_0(103)
       2: H1_T11_1(104)
    2: seq_num(2), interface(H1_T22):
       1: H1_T22_0(105)
       2: H1_T22_1(106)
    3: seq_num(3), interface(H1_T33):
       1: H1_T33_0(100)
  Members(8):
    1: Seq_num(1 H1_T11 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    2: Seq_num(2 H1_T22 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    3: Seq_num(3 H1_T33 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    
```
**4: Seq_num(3 H1_T33_0 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    5: Seq_num(1 H1_T11_0 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    6: Seq_num(1 H1_T11_1 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    7: Seq_num(2 H1_T22_0 overlay), alive, sla(0x1), gid(2), num of pass(1), selected
    8: Seq_num(2 H1_T22_1 overlay), alive, sla(0x1), gid(2), num of pass(1), selected**
  Src address(1):
        10.0.0.0-10.255.255.255
  Dst address(1):
        10.0.0.0-10.255.255.255
                                            From the diagnostic command on Spoke 2, observe the shortcuts in **bold**:

