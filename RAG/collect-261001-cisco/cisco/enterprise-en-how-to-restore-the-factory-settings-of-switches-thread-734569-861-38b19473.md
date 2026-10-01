---
id: collect-261001-cisco/cisco/enterprise-en-how-to-restore-the-factory-settings-of-switches-thread-734569-861-38b19473
title: "enterprise-en-how-to-restore-the-factory-settings-of-switches-thread-734569-861-38b19473"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/enterprise-en-how-to-restore-the-factory-settings-of-switches-thread-734569-861-38b19473.md
source_anchor: ""
source_lines: [1, 31]
sha256: f1453962ac6f71bc4940ba91e16d943f142c5278ba1e5c901f601f1228b9cfe0
---

# enterprise-en-how-to-restore-the-factory-settings-of-switches-thread-734569-861-38b19473

Hello everyone,
Today I will share with you how to restore the factory setting of switches.
This post provides three methods for restoring factory settings:
Press and holds the PNP button to restore factory settings;
restore factory settings with one click;
clear the configuration file to restore factory settings.
The differences between the three methods are as follows:
Whether factory settings are restored by using commands or pressing a specific button
Pressing and holding the PNP button to restore factory settings does not require a login to the switch and is easy to operate.
The following describes the terms in the table:
Factory settings: are the basic configurations that are installed before delivery, that is, the configurations contained in the $_default.cfg file in the protected directory.
Current configurations: are the configurations that are in effect when the switch is running. To check the current configurations, run the display current-configuration command.
Configuration file: contains the configurations that are saved in the default storage path of the system using the save command and are used for initialization during switch power-on. To check information about the configuration file used during switch startup, run the display startup or display saved-configuration command.
For a device that has the PNP button, you can press and hold the PNP button for more than 6 seconds to restore the device to factory settings and restart the device. The following figure shows the PNP button.
Run the reset factory-configuration command in the user view.
<HUAWEI> reset factory-configuration
Warning: The command will delete all the configurations and files (except the startup, patch, module, and license files) from the device. Continue? [Y/N]:y //If you enter y, all configuration and data files on the device are cleared.
Warning: The system will reboot after configurations and files are deleted. Continue? [Y/N]:y //If you enter y, the device automatically restarts.
Run the reset saved-configuration command in the user view.
<HUAWEI> reset saved-configuration
Warning: The action will delete the saved configuration in the device.
The configuration will be erased to reconfigure. Continue? [Y/N]:y //If you enter y, the saved configurations on the device are cleared.
Warning: Now clearing the configuration in the device.
Info: Succeeded in clearing the configuration in the device.
Run the reboot command in the user view to restart the device.
<HUAWEI> reboot
Info: The system is now comparing the configuration, please wait.................
Warning: The configuration has been modified, and it will be saved to the next startup saved-configuration file flash:/vrpcfg.zip. Continue? [Y/N]:n //If you enter n, the configuration file is not saved.
Info: If want to reboot with saving diagnostic information, input 'N' and then execute 'reboot save diagnostic-information'.
The system will reboot! Continue?[Y/N]:y //If you enter y, the device restarts.
That is all I want to share with you! Thank you!
