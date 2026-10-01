---
id: collect-261001-cisco/cisco/r-cisco-comments-bnxfob-vlan-in-running-config-but-not-in-vlan-database-a2477fb2
title: "r-cisco-comments-bnxfob-vlan-in-running-config-but-not-in-vlan-database-a2477fb2"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["pruning"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-bnxfob-vlan-in-running-config-but-not-in-vlan-database-a2477fb2.md
source_anchor: ""
source_lines: [1, 8]
sha256: 73aa294e55ca9ae56530f0082f9dd7611d23a7acf962b00c731a1586ab2a587f
---

# r-cisco-comments-bnxfob-vlan-in-running-config-but-not-in-vlan-database-a2477fb2

VLAN in running config, but not in VLAN database 
        
    Been seeing an odd issue as of late. I have a lot of remote networks and every now and then one of them takes a power hit. Their routers come back up with the correct running config, VLAN interfaces intact and assigned to their proper ports, they show assigned and up/ up, but they disapear from the VLAN database and I have to have the remote tech and them directly to the database to get things moving again. We don't join these to a VTP domain for reasons either. Any idea why the VLAN database would empty, but everytihng looks fine in the running config? Or how to stop this?
Section des commentaires
Do you have your switches in VTP transparent mode? If so, you need to do "copy run start" after creating the VLANs as they are saved in the running-config and not automatically in VLAN.dat like they are for VTP server/client.
They are in server mode and the copy run start is done once the configuration is loaded in initailly, however once the power goes out and comes back up the VLAN database is set back to default.
if you don't join your switches to vtp-domain - make them vtp transparent. after that vlan database wouldn't be used and all vlans info will be stored in config.
situation looks weird and requires deeper investigations. vlan db shouldn't being reseted, please doublecheck boot logs of device - maybe it will say something like - flash drive is corrupted or something like that. Also VTP-pruning can be the source of such problems.
