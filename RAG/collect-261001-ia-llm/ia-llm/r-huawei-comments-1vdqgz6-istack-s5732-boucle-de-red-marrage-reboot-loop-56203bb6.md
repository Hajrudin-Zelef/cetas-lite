---
id: collect-261001-ia-llm/ia-llm/r-huawei-comments-1vdqgz6-istack-s5732-boucle-de-red-marrage-reboot-loop-56203bb6
title: "r-huawei-comments-1vdqgz6-istack-s5732-boucle-de-red-marrage-reboot-loop-56203bb6"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-07-11"]
keywords: ["incident", "license"]
source: docs/RAG/collect-261001-ia-llm/r-huawei-comments-1vdqgz6-istack-s5732-boucle-de-red-marrage-reboot-loop-56203bb6.md
source_anchor: ""
source_lines: [1, 118]
sha256: d229005c88898f58fd1ff735a74f935f9e38832b3f5993790f5b7f4f5a629862
---

# r-huawei-comments-1vdqgz6-istack-s5732-boucle-de-red-marrage-reboot-loop-56203bb6

[iStack S5732] Boucle de redémarrage (reboot loop) 
        
        
        
    
    
    Hello TEAMS, I've been having a problem for over a week. Please, can you help me?
Environment
      - Model: 2x Huawei S5732 (48 ports)
- VRP Version: V200R022C00SPC500 (identical on both switches — verified)
- Topology: iStack, 2 members, ring-based stack links over 2x 25GE (stack ports 1 and 2 on each switch)
    
      Roles:
- SW1 (slot 0) = Master, priority 150
- SW2 (slot 1) = Standby, priority 120
    
Background / Sequence of Events
Both switches were previously used in standalone mode (each configured individually). We recently set up stacking (iStack) between them for the first time, following the standard procedure:
- 
      Individual configuration of each switch in standalone mode:
      -------------------------- Configuration Sw1 on slot 0 start --------------------------------
system-view
sysname SW1-DC
stack enable
stack slot 0 renumber 0
stack slot 0 priority 150
stack reserved-vlan 4093
stack timer mac-address switch-delay 10
interface stack-port 0/1
port interface 25GE0/1/7 enable
quit
interface stack-port 0/2
port interface 25GE0/1/8 enable
quit
save
--------------------------Configuration Sw1 on slot 0 End----------------------------------
    
      -------------------------- Configuration Sw2 on slot 1 Begin--------------------------------
system-view
sysname SW2-DC
stack enable
stack slot 0 renumber 1
stack slot 1 priority 120
stack reserved-vlan 4093
stack timer mac-address switch-delay 10
interface stack-port 1/1
port interface 25GE0/1/7 enable
quit
interface stack-port 1/2
port interface 25GE0/1/8 enable
quit
save
---------------------------Configuration Sw2 on slot 1 End------------------------------------
    
- 
      Power off both switches, connect the ring-configuration stacking cables to the dedicated 25GE ports, then power them back on simultaneously.
The stack formed correctly on the first attempt: Master (SW1) and Standby (SW2) were elected as expected, and all VLANs and interfaces were operational.
Once the stack was operational, we began the usual resilience tests (restarting a member to verify behavior in the event of a failure):
Test 1 — Restarting the Standby (SW2): The switch restarts, joins the existing stack normally, and no incidents occur. The Master does not change roles during this operation.
Test 2 — restarting the Master (SW1): this is where the problem arose. While SW1 restarts, SW2 (Standby) must temporarily assume the Master role. It was precisely at this moment of role transition that the software crash and reboot loop described below occurred.
This test was repeated a second time (second incident below) to verify reproducibility: the behavior recurred (reboot loop), but this time without the visible crash, and the stack eventually stabilized after several cycles.
Problem
The stack functions normally during regular use. The problem occurs only during a reboot:
1- Restarting the Standby (SW2) → the stack rebuilds normally, without incident.
2- Restarting the Master (SW1) → the stack enters an automatic restart loop and does not stabilize immediately.
Log of the first incident
During the role switch (the Standby must temporarily become the Master while SW1 restarts), a software crash occurs:
      Slot 1 was elected as slave
The stack-port 1 turned to up
...
The stack-port 2 turned to up
    
      Assert at File: srm_func.c, Line: 34662. CallStack:
<-- 0x8da8440(vosAssertRecord+0x80) <-- 0x8da85d0(vos_assert+0xe8)
<-- 0x69e1c0c(+0x6097c0c) <-- 0x86719f4(SRM_GL_IsLicenseAttribDemo+0x124)
<-- 0x8671af0(SRM_GL_IsNeedAddLicItem+0x74) <-- 0x8672290(SRM_GL_AddLicItem+0xc4)
<-- 0x8672440(SRM_GL_RpcAddLicItemCallBack+0x50) <-- 0x865bad8(SRM_DEVRPCCallBack+0x100)
<-- 0x6a9e424(IPC_RPC_Notify+0x3a4) <-- 0x6a9eb94(IPC_RPC_TimerTask+0x118)
<-- 0x8e79c10(VosTskAllTaskEntry+0x164) <-- 0x8da1ecc(tskAllAdaptTaskEntry+0x148)
<-- 0x252a978(+0x1be0978) <-- 0xf4943e40(+0x7ae40)
<-- 0xf49c499c(+0xfb99c)
<TaskID: 46> [2026.07.11 13:27:47.100]
    
      Slot 1 switched to the standby at 2026-07-11 13:27:55.970
...
Slot 1 switched to the master at 2026-07-11 13:31:00.130
Update Old master 0 unregistered.
Then a series of automatic restarts:
    
      Device will restart for stack merge
System reboot ... Startup information : 2
Device will restart for stack merge
System reboot ... Startup information : 3
Device will restart for stack merge
System reboot ... Startup information : 4
Log of the second incident (reproduction)
    
When retesting a Master reboot, the same “Device will restart for stack merge” loop occurred over 3–4 consecutive cycles, but without the srm_func.c assert crash this time. The stack eventually stabilized on the 5th boot, and the configuration was restored normally:
      Startup information : 5
...
The stack-port 1 turned to up
The stack-port 2 turned to up
Device will restart for stack merge <-- dernière occurrence
...
Recover configuration begin ...
Recover configuration end
    
What has already been verified / ruled out
VRP version: identical on both members (V200R022C00SPC500) → rules out the possibility of software incompatibility between members.
Cabling/stack port negotiation: Stack ports 1 and 2 come up successfully on every attempt, before each reboot.
The behavior is reproducible: observed on two separate occasions, only when the Master is rebooted.
Questions for the community
1) In your opinion, what is the cause of this issue?
2) Why does restarting the Master switch disrupt the stack in this way, while restarting the Standby switch causes no issues?
3) Is this crash in the SRM module (license) a known issue with this version?
Section des commentaires
Soyez la première personne à commenter
Personne n’a encore répondu à cette publication. Partage ton avis pour lancer la conversation.
