---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-58-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-58.md
source_anchor: ""
source_lines: [128, 192]
sha256: b7c052db33cc6a98b84233b31634f0ca9c5723f5d196341ff4f0d03d5c5fdc70
---

# Overview

| Property | Description | 
|---|---|
| **from-line** | Start executing the script from the specified line number. This option is only available in verbose mode. | 
| **file-name** | Name of the script (.rsc) file to be executed. | 
| **verbose** | Reads each line from the file and executes individually, allowing to debug syntax or other errors more easily. | 
| **dry-run** | Simulates the import without making any configuration changes. This helps in catching syntax errors. This option is only available in verbose mode. | 

If the device has a default or existing configuration that requires replacement, it is necessary to initiate a configuration reset.

This involves applying a clean, empty configuration using the command /system/reset-configuration no-defaults=yes, followed by a device reboot.

## Auto Import

It is also possible to **automatically** execute scripts after uploading to the router with FTP or SFTP. The script file must be named with the extension *.auto.rsc. Once the commands in the file are executed, a new *.auto.log file is created which contains import success or failure information.

".auto.rsc" in the filename is mandatory for a file to be automatically executed.

## Import troubleshooting

### Configuration parts to watch out for in exported .rsc files

Things that should be removed from export files that were created with "/export", before attempting import on a new device.

- Interface renaming conflicts with the default ethernet naming scheme.

- In older versions "export" default entries might show with "add" instead of the "set" command. That should be edited before import to avoid errors.
- Check if the total number of physical interfaces count matches the new and old devices. If there are some missing that will end up in error during .rsc import.

In case of problematic import, attempt the following:

- Use the **dry-run** parameter to simulate the import without making any configuration changes. This helps in catching syntax errors. This option is only available in verbose mode.
- Reset the configuration on that device.
- Run the import command again with the "verbose=yes" argument. It will also stop the import process on a problem that you already encountered, but will also show the place where the export failed. This way shows you the place where things need to be edited in the .rsc import file.

### Startup delay

If your configuration relies on interfaces that might not yet have started up upon command execution, it is suggested to introduce delays or to monitor until all needed interfaces are available. This example script allows you to set how many interfaces you are expecting, and how long to wait until they become available:

The above script will wait until there are 10 interfaces visible, or 30 seconds. If there are no 10 interfaces at this time, it will put a message in the log. Modify the variables according to your needs.

# Configuration Reset

RouterOS allows resetting configuration with `/system reset-configuration` command

This command clears all configuration of the router and sets it to the factory defaults including the login name and password ('admin' with an empty password or, for some models, check user and wireless passwords on the sticker). For more details on the default configuration see the list.

After executing the configuration reset command, the router will reboot and load the default configuration. Starting from version 7.13, following the reset, a license prompt will be displayed with the option to view the end-user license agreement.

The backup file of the existing configuration is stored before reset. That way you can easily restore any previous configuration if the reset is done by mistake.

If the router was installed using Netinstall and had a script specified as the initial configuration, the reset command executes this script after purging the configuration. To stop it from doing so, you will have to reinstall the router.

It is possible to override the default reset behavior with the parameters below:

| Property | Description | 
|---|---|
| **keep-users** | Do not remove existing users from the configuration | 
| **no-defaults** | Do not load the default configuration, just clear the configuration | 
| **skip-backup** | Skip automatic backup file generation before reset | 
| **run-after-reset** | Run specified .rsc file after reset. That way you can load your custom configuration. If a specific .rsc file execution takes more than 2 minutes, a script will fail, and LOG will contain *"runtime limit exceeded"*  or in rare cases*"std failure: timeout"* error. | 
| **caps-mode** | Run *caps-mode* script after configuration reset. | 

For example hard reset configuration without loading default config and skipping backup file:

And the same using Winbox:
