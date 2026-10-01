---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13-1
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13.md
source_anchor: ""
source_lines: [1, 53]
sha256: 98d322a2ee7f5c47760627638e67baa73300889c2c34b77f1cc285dcf96e2ab9
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-808afc13

Recover Enable and Telnet Passwords
If you forget the enable or Telnet passwords, you can recover them for ASA models. The procedure differs by device type. You must perform the task using the CLI.
| Note | For Firepower platforms, you cannot recover lost passwords. You can only restore the factory default configuration, and reset the passwords to the default. For Firepower 4100/9300, see the FXOS configuration guide. For Firepower and 2100, see the FXOS troubleshooting guide. | 
Recover Passwords on the ASA 5500-X
This procedure works for the ASA 5512-X, 5515-X, 5525-X, 5545-X, 5555-X, and 5585-X.
To recover passwords for the ASA, perform the following steps.
Procedure
| Step 1 | Connect to the ASA console port. | 
| Step 2 | Power off the ASA, then power it on. | 
| Step 3 | After startup, press the Escape key when you are prompted to enter ROMMON mode. | 
| Step 4 | To update the configuration register value, enter the following command:  rommon #1> confreg 0x41 Update Config Register (0x41) in NVRAM...   | 
| Step 5 | To set the ASA to ignore the startup configuration, enter the following command:  rommon #1> confreg  The ASA displays the current configuration register value, and asks whether you want to change it:  Current Configuration Register: 0x00000041 Configuration Summary:   boot default image from Flash   ignore system configuration  Do you wish to change this configuration? y/n [n]: y   | 
| Step 6 | Record the current configuration register value, so you can restore it later. | 
| Step 7 | At the prompt, enter Y to change the value. The ASA prompts you for new values. | 
| Step 8 | Accept the default values for all settings, except for the "disable system configuration?" value. | 
| Step 9 | At the prompt, enter Y. | 
| Step 10 | Reload the ASA by entering the following command:  rommon #2> boot Launching BootLoader... Boot configuration file contains 1 entry.  Loading disk0:/asa800-226-k8.bin... Booting...Loading...  The ASA loads the default configuration instead of the startup configuration. | 
| Step 11 | Access the privileged EXEC mode by entering the following command:  ciscoasa# enable   | 
| Step 12 | When prompted for the password, press Enter. The password is blank. | 
| Step 13 | Load the startup configuration by entering the following command:  ciscoasa# copy startup-config running-config   | 
| Step 14 | Access the global configuration mode by entering the following command:  ciscoasa# configure terminal   | 
| Step 15 | Change the passwords, as required, in the default configuration by entering the following commands:  ciscoasa(config)# password password ciscoasa(config)# enable password password ciscoasa(config)# username name password password   | 
| Step 16 | Load the default configuration by entering the following command:  ciscoasa(config)# no config-register   The default configuration register value is 0x1. See the command reference for more information about the configuration register. | 
| Step 17 | Save the new passwords to the startup configuration by entering the following command:  ciscoasa(config)# copy running-config startup-config   | 
Recover Passwords on the ASA 5506-X, ASA 5508-X, and ASA 5516-X
To recover passwords for the ASA 5506-X, ASA 5508-X, and ASA 5516-X perform the following steps:
Procedure
| Step 1 | Connect to the ASA console port. | 
| Step 2 | Power off the ASA, then power it on. | 
| Step 3 | After startup, press the Escape key when you are prompted to enter ROMMON mode. | 
| Step 4 | To update the configuration register value, enter the following command:  rommon #1> confreg 0x41  You must reset or power cycle for new config to take effect  The ASA displays the current configuration register value and a list of configuration options. Record the current configuration register value, so you can restore it later.  Configuration Register: 0x00000041  Configuration Summary  [ 0 ] password recovery  [ 1 ] display break prompt  [ 2 ] ignore system configuration  [ 3 ] auto-boot image in disks  [ 4 ] console baud: 9600  boot: ...... auto-boot index 1 image in disks   | 
| Step 5 | Reload the ASA by entering the following command:  rommon #2> boot Launching BootLoader... Boot configuration file contains 1 entry.  Loading disk0:/asa932-226-k8.bin... Booting...Loading...  The ASA loads the default configuration instead of the startup configuration. | 
| Step 6 | Access the privileged EXEC mode by entering the following command:  ciscoasa# enable   | 
| Step 7 | When prompted for the password, press Enter. The password is blank. | 
| Step 8 | Load the startup configuration by entering the following command:  ciscoasa# copy startup-config running-config   | 
| Step 9 | Access the global configuration mode by entering the following command:  ciscoasa# configure terminal   | 
| Step 10 | Change the passwords, as required, in the default configuration by entering the following commands:  ciscoasa(config)# password password ciscoasa(config)# enable password password ciscoasa(config)# username name password password   | 
| Step 11 | Load the default configuration by entering the following command:  ciscoasa(config)# no config-register   The default configuration register value is 0x1. See the command reference for more information about the configuration register. | 
| Step 12 | Save the new passwords to the startup configuration by entering the following command:  ciscoasa(config)# copy running-config startup-config   | 
Recover Passwords or Images on the ASAv
To recover passwords or images on the ASAv, perform the following steps:
Procedure
| Step 1 | Copy the running configuration to a backup file on the ASAv: copy running-config filename Example:  ciscoasa# copy running-config backup.cfg   | 
| Step 2 | Restart the ASAv: reload | 
| Step 3 | From the GNU GRUB menu, press the down arrow, choose the <filename> with no configuration load option, then press Enter. The filename is the default boot image filename on the ASAv. The default boot image is never automatically booted through the fallback command. Then load the selected boot image.  GNU GRUB version 2.0(12)4 bootflash:/asa100123-20-smp-k8.bin bootflash: /asa100123-20-smp-k8.bin with no configuration load  Example:  GNU GRUB version 2.0(12)4 bootflash: /asa100123-20-smp-k8.bin with no configuration load   | 
| Step 4 | Copy the backup configuration file to the running configuration. copy filename running-config Example:  ciscoasa (config)# copy backup.cfg running-config   | 
| Step 5 | Reset the password. enable password password Example:  ciscoasa(config)# enable password cisco123   | 
| Step 6 | Save the new configuration. write memory Example:  ciscoasa(config)# write memory   | 
Disable Password Recovery for ASA Hardware
| Note | You cannot disable password recovery on the ASAv or Firepower models. | 
To disable password recovery to ensure that unauthorized users cannot use the password recovery mechanism to compromise the ASA, perform the following steps.
Before you begin
On the ASA, the no service password-recovery command prevents you from entering ROMMON mode with the configuration intact. When you enter ROMMON mode, the ASA prompts you to erase all Flash file systems. You cannot enter ROMMON mode without first performing this erasure. If you choose not to erase the Flash file system, the ASA reloads. Because password recovery depends on using ROMMON mode and maintaining the existing configuration, this erasure prevents you from recovering a password. However, disabling password recovery prevents unauthorized users from viewing the configuration or inserting different passwords. In this case, to restore the system to an operating state, load a new image and a backup configuration file, if available.
