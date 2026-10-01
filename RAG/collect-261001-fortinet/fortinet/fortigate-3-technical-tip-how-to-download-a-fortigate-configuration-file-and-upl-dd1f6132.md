---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-how-to-download-a-fortigate-configuration-file-and-upl-dd1f6132
title: "fortigate-3-technical-tip-how-to-download-a-fortigate-configuration-file-and-upl-dd1f6132"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-how-to-download-a-fortigate-configuration-file-and-upl-dd1f6132.md
source_anchor: ""
source_lines: [1, 155]
sha256: 43500d0905f11f05ecd662e1bb941840a2ca1d6cb9064523fdd6c6a8909a4b8f
---

# fortigate-3-technical-tip-how-to-download-a-fortigate-configuration-file-and-upl-dd1f6132

**Description**


Refer to this KB article: Technical Tip: Backing Up the FortiGate configuration file via SCP with limited Read/Write permissions

**Scope**

**System -> Administrator Users Access Control**with all other Access Control Permissions set to 'Read' to perform an SCP backup.

**Solution**


**Step 1: Enable SCP.**


From the CLI:


**config system global**

    set admin-scp enable

end


**Step 2: Enable SSH access on the interface.**


SCP uses SSH protocol to provide secure file transfer. The interface used for administration must allow SSH access.


In the GUI:


- Go to **System -> Network -> Interface** .
- Select the **Edit** icon for the interface use for administrative access.
- In the **Administrative Access** section, select the SSH check box.
- Select 'OK'.


- Add ssh to the allowaccess setting:

**config system interface**

**Step 3: Confirm the FortiGate SSH Port.**

The default is ssh port 22. If a different SSH Port is being used, add '-P <SSH_Port>' to the SCP commands.

**show full-configuration | grep admin-ssh-port**

**Step 4:**It is now possible to download or upload an image and configuration to the FortiGate.


**To Backup FortiGate configuration use the SCP client.**

The FortiGate unit configuration file name is sys_config. Use the following syntax to download the file:

**Linux:**

If using OpenSSH 9.0 or higher, the scp command will require the '-O' option to use the legacy scp protocol.


**Windows:**

**To r**

**estore configuration to the FortiGate, use the SCP client.**

- FortiGate will **reboot immediately**  after the file is uploaded.
- When uploading (restoring) configuration file to FortiGate, the destination file name is 'fgt-restore-config'. Use the following syntax to upload the file:

**Windows:**

**pscp.exe -scp <path_to_config_file> admin@<FortiGate_IP>:fgt-restore-config**

**Manual firmware upgrade using an SCP client.**

- FortiGate will **reboot immediately**  after the file gets uploaded
- The FortiGate firmware file name is 'firmware.out'. Use the following syntax to upload the file:

**Linux:**

**scp <firmware.out> <admin-user>@<IP>:fgt-image**

**Windows:**


**Optional public-private key authentication**

**Configure public-private key authentication.**

Step-by-step:

- Create a public-private key pair using a key generator tool compatible with the SCP client.
- Save the private key to the location on the computer where the SSH private keys are stored.
This step depends on the SCP product. The SSH Secure Shell key generator automatically stores the private key. It is necessary to manually save the private key in the PuTTY Key Generator.
- Copy the public key to the FortiGate unit. Do this in the FortiGate CLI, as follows:

**config system admin**

edit admin

set ssh-public-key1 "<key-type> <key-value>"

end


<key-type> must be ssh-dss for a DSA key or ssh-rsa for an RSA key. For <key-value>, copy the public key data and paste it into the CLI command.

When copying the key data from Windows Notepad, observe the following principles to copy the key data correctly:

- Copy one line at a time and make sure that the paste each line of key data at the end of the previously pasted data.
- Do not copy the end-of-line characters that appear as small rectangles in Notepad.
- Do not copy the ---- BEGIN SSH2 PUBLIC KEY ---- or Comment: "[2048-bit dsa,...]" lines.
- Do not copy the ---- END SSH2 PUBLIC KEY ---- line.
- Type the closing quotation mark and press Enter.
- Enter the end command.

The SCP client can now authenticate to the FortiGate unit based on SSH keys instead of an administrator password.**Examples of using the SCP Client:**

These examples show how to download the configuration file from a FortiGate unit at IP address 172.20.120.171, using Linux and Windows SCP clients.**Linux client example:**


To download the configuration file to a local directory called ~/config, enter the following command:

Enter the admin password when prompted.**Windows client example:**


To download the configuration file to a local directory called c:\config, enter the following command in a Command Prompt window:

Enter the admin password when prompted.

This example shows how to upload (restore) configuration file to a FortiGate unit with IP address 172.20.120.171, from Windows machine.

Enter the admin password when prompted.

These examples show how to upload the firmware file from a FortiGate unit at IP address 172.20.120.171, using Linux SCP clients.**Linux client example:**


To upload the firmware file to a local directory called firmware.out, enter the following command:

Enter the admin password when prompted.


**scp admin@172.20.120.171:sys_config ~/config
pscp admin@172.20.120.171:sys_config c:\config
pscp.exe -scp C:\Users\<configuration file> admin@172.20.120.171:fgt-restore-config**

pscp admin@172.20.120.171:sys_config c:\config

pscp.exe -scp C:\Users\<configuration file> admin@172.20.120.171:fgt-restore-config

**pscp.exe -scp -P 222 C:\Users\<configuration file> admin@172.20.120.171:fgt-restore-config**

**scp firmware.out admin@172.20.120.171:fgt-image**
