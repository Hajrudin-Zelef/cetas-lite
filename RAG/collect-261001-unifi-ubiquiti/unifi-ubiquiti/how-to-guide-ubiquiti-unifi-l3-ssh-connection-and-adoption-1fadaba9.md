---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-guide-ubiquiti-unifi-l3-ssh-connection-and-adoption-1fadaba9
title: "how-to-guide-ubiquiti-unifi-l3-ssh-connection-and-adoption-1fadaba9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-guide-ubiquiti-unifi-l3-ssh-connection-and-adoption-1fadaba9.md
source_anchor: ""
source_lines: [1, 62]
sha256: db1d16ceb4923e62edfee5d84bd3fe3213d462f7b8c202126ccdff42bcc89daf
---

# how-to-guide-ubiquiti-unifi-l3-ssh-connection-and-adoption-1fadaba9

Regular UniFi adoption is preformed over Layer 2. L2 adoption on UniFi (where the devices and UniFi Network application are on the same network) uses self discovery, any locally-available, unmanaged UniFi Devices will appear as “Pending Adoption” in the UniFi Network application itself (in the Devices section). This is the most typical adoption method and is perfectly serviceable in most use cases.
But what if you are not on the same network? If your UniFi network controller is at a remote location, such as a head office, second premises or even hosted on AWS you can uses Layer 3 Adoption via SSH.
This is useful if you manage multiple offices centrally on a virtual machine or provide management services to end users. Read our Ubiquiti UniFi Controller guide here to find out more about these methods.
If you can SSH into the device, it’s possible to L3 adopt via CLI command. There are various methods for this depending on the device/OS you use. The main methods are outlined below:
Connect via SSH on Windows
Connect via SSH by using PuTTY
1. Either install the program (.msi) or download and run the standalone executable (.exe).
2. Select Run if prompted by a Windows security warning.
3. Enter the IP address of the Ubiquiti device in the Host Name (or IP address) box and leave the port set to 22.
4. The first connection to a device will prompt a host key warning that asks for confirmation to continue. Select Yes to continue connecting.
5. A command line (terminal) window will open that prompts for a username login.
6. Enter the username of the device and press Enter.
7. A welcome banner may appear, and you will now be prompted to enter the password. As you type while entering the password, no characters will appear on screen for security purposes. Carefully enter the correct password, and press Enter. Example output below:
login as: ubnt
ubnt@192.168.1.1's password:
Connect via SSH on macOS
To connect via SSH by using the built-in macOS SSH client:
1. Open the Terminal by navigating to Applications > Utilities > Terminal.
2. A command line (terminal) window will open that accepts various input commands.
3. Use the ssh command followed by the username and IP address of the Ubiquiti device (separated by the @ symbol).
ssh <username>@<ip-address>
4. Replace <username> and <ip-address> with the username and IP address used by the Ubiquiti device and press Enter. For example:
ssh ubnt@192.168.1.1
5. The first connection to a device will display a host authenticity warning and a prompt that asks for confirmation to continue. Type yes and press Enter to continue connecting to the device. Example output below:
ssh ubnt@192.168.1.1
The authenticity of host '192.168.1.1 (192.168.1.1)' can't be established.
ECDSA key fingerprint is SHA256:...
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.1.1' (ECDSA) to the list of known hosts.
ubnt@192.168.1.1's password:
6. A welcome banner may appear, and you will now be prompted to enter the password. As you type while entering the password, no characters will appear on the screen for security purposes. Carefully enter the correct password, and press Enter.
7. You will be presented with the device’s welcome banner if the entered credentials were valid. A command prompt indicating your username and device name will indicate a successful connection and that the remote device is ready to accept commands.
Connect via SSH on Linux
To connect via SSH by using the built-in Linux SSH client:
1. Open the Terminal which is typically located in the Utilities > Terminal section of the main menu. On most distributions, you can also open the Terminal by pressing CTRL+ALT+T.
2. A command line (terminal) window will open that accepts various input commands.
3. Use the ssh command followed by the username and IP address of the Ubiquiti device (separated by the @ symbol).
ssh <username>@<ip-address>
4. Replace <username> and <ip-address> with the username and IP address used by the Ubiquiti device and press Enter. For example:
ssh ubnt@192.168.1.1
5. The first connection to a device will display a host authenticity warning and a prompt that asks for confirmation to continue. Type yes and press Enter to continue connecting to the device. Example output below:
ssh ubnt@192.168.1.1
The authenticity of host '192.168.1.1 (192.168.1.1)' can't be established.
ECDSA key fingerprint is SHA256:...
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.1.1' (ECDSA) to the list of known hosts.
ubnt@192.168.1.1's password:
6. A welcome banner may appear, and you will now be prompted to enter the password. As you type while entering the password, no characters will appear on the screen for security purposes. Carefully enter the correct password, and press Enter.
7. You will be presented with the device’s welcome banner if the entered credentials were valid. A command prompt indicating your username and device name will indicate a successful connection and that the remote device is ready to accept commands.
Connection established
Once you have successfully connected to your Unifi device, you well need to follow these steps in order to direct the unit to your controller.
1. Make sure the device is running updated firmware. If not, follow the sub steps bullet pointed below.
- Copy the link of your device’s firmware download link from here.
- When connected via SSH into the device run the upgrade command followed by the download link. For Example –
upgrade https://dl.ui.com/unifi/firmware/UAL6/5.60.1.12923/BZ.mt7621_5.60.1+12923.210416.1641.bin
2. Make sure the device is in the factory default state. If it’s not, run the following command:
sudo syswrapper.sh restore-default
3. SSH into the device and type the following and hit enter, substituting “ip-of-host” with the IP address of the host of your UniFi Network application:
set-inform http://ip-of-host:8080/inform
4. After issuing the set-inform, the UniFi device will show up for adoption in the Devices section of UniFi Network. Once you click Adopt, the device will appear to go offline or have the status of “Adopting” then proceed to “Provision” and “Connected”.
If the device does not appear after 10-15 minutes, try step 3 and 4 again.
You now have a remotely adopted UniFi device. For more assistance, call our technical team on 01449 724255 or email shop@linitx.com.
