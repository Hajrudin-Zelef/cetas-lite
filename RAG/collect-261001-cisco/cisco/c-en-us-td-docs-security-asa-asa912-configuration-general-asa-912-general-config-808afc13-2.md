---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13-2
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13.md
source_anchor: ""
source_lines: [54, 56]
sha256: 66057212d516328b8c77973101fcb47272ac6fc9f839866823dedc9c0ded56fb
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13

The service password-recovery command appears in the configuration file for information only. When you enter the command at the CLI prompt, the setting is saved in NVRAM. The only way to change the setting is to enter the command at the CLI prompt. Loading a new configuration with a different version of the command does not change the setting. If you disable password recovery when the ASA is configured to ignore the startup configuration at startup (in preparation for password recovery), then the ASA changes the setting to load the startup configuration as usual. If you use failover, and the standby unit is configured to ignore the startup configuration, then the same change is made to the configuration register when the no service password- recovery command replicates to the standby unit.
Procedure
| Disable password recovery. no service password-recovery Example:  ciscoasa (config)# no service password-recovery   |
