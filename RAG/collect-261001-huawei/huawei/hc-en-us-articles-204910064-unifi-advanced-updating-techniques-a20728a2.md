---
id: collect-261001-huawei/huawei/hc-en-us-articles-204910064-unifi-advanced-updating-techniques-a20728a2
title: "hc-en-us-articles-204910064-unifi-advanced-updating-techniques-a20728a2"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hc-en-us-articles-204910064-unifi-advanced-updating-techniques-a20728a2.md
source_anchor: ""
source_lines: [1, 86]
sha256: 87a879006704f5a7870cb63776ca2fdc8ed8cd436a47ab2bb90aacc69a6dba78
---

# hc-en-us-articles-204910064-unifi-advanced-updating-techniques-a20728a2

UniFi - Advanced Updating Techniques
Use advanced techniques as a last resort to update UniFi devices and UniFi software when it is not possible to do so via the user interface.
We highly recommend that most users enable automatic updates, or trigger updates through the user interface. See How to Update UniFi to learn more.
Update UniFi Devices via Web Application
Updating via the Device Property Panel
Use Case: You want to try Early Access firmware releases for specific devices, or you want to return to an official release after trying an EA release.
- Head to community.ui.com/releases, find your appropriate firmware release, right-clicking the link and select "Copy Link Address."
- Paste the link in the address bar found in the Settings tab of the device’s properties panel.
Updating via Your Network Cache
Use Case: You prefer to download and store updates in your Network application so they can be used by other devices, as opposed to downloading multiple, device-specific files from the internet. This is an ideal solution for reducing bandwidth within high-volume networks that host a large number of similar UniFi devices. It is also suitable for the advanced users who disable internet access on their UniFi device’s management network.
Device updates can be cached in your Network application’s System Settings. Once an update is cached, you can open to your UniFi Devices page and click Update Available.
Note: The Cache link will appear when you hover your cursor over an update.
Updating via SSH
Please note that SSH updating is not an officially supported process and may prevent your UniFi Console from functioning. Only do this at the request of UI Support. It is only prescribed to work around specific scenarios, such as when:
- Prior, traditional, update attempts have failed, likely due to an incorrect network configuration. For more details, see How to Update UniFi.
- Your UniFi Network device is not being discovered or cannot be adopted because it has been preloaded with outdated firmware.
- Your UniFi Console cannot be set up because it has been preloaded with an outdated version of UniFi OS.
UniFi Access Point/Switch (Internet)
- Copy the update link from community.ui.com/releases.
- SSH into your device.
- Run the following command:
upgrade paste_download_link_here Example:upgrade https://dl.ui.com/unifi/firmware/UAL6/5.60.1.12923/BZ.mt7621_5.60.1+12923.210416.1641.bin
UniFi Access Point/Switch (No Internet)
- Download the desired firmware update from community.ui.com/releases.
- Use the following SCP command to copy the file into the /tmp folder of your device from your workstation. This requires a compatible SCP application (e.g., Terminal on macOS and Linux, PuTTY/PowerShell on Windows).
scp /folder_path/firmwarefile.bin <user>@<IP of device>:/tmp/fwupdate.bin Example:
 scp /Users/alexpro/Desktop/BZ.mt7621_5.60.1+12923.210416.1641.bin Alex@192.168.1.219:/tmp/fwupdate.bin
- Enter your SSH password when prompted.
- SSH into your device.
- Run the following command:
 Note: If you use the USW-Multi firmware image, you'll need to runfwutil -p /tmp/fwupdate.bin before the command below.
 syswrapper.sh upgrade2 &
UniFi Consoles/Gateways (Internet)
- Copy the update link from community.ui.com/releases.
- SSH into your device.
- Run the following command:
ubnt-systool fwupdate paste_download_link_here 
Example:
ubnt-systool fwupdate https://fw-download.ubnt.com/data/unifi-dream/dd49-UDR-2.4.10-cd3afa000ebf4a4fb15374481539961c.bin
If your UDM, UDM Pro or UXG-Pro runs version 1.x, then please use the ubnt-upgrade command.
Example:
ubnt-upgrade https://fw-download.ubnt.com/data/udm/1adc-udmpro-1.12.38-ca8a490ac2b04247abb3f7d3e3eae01a.bin
UniFi Consoles/Gateways (No Internet)
- Download the desired firmware update from community.ui.com/releases.
- Use the following SCP command to copy the file into the /tmp (/mnt/data for version 1.12 and below) folder of your device from your workstation. This requires a compatible SCP application (e.g., Terminal on macOS and Linux, PuTTY/PowerShell on Windows).
scp /folder_path/firmwarefile.bin <user>@<IP of device>:/tmp/fwupdate.bin Example:
 scp /Users/alexpro/Desktop/UDR-2.4.10.bin root@192.168.1.1:/tmp/fwupdate.bin If you are running Version 1.12 and below, use the following command:
 scp /folder_path/firmwarefile.bin <user>@<IP of device>:/mnt/data/fwupdate.bin Example:
 scp /Users/alexpro/Desktop/UDR-2.4.10.bin root@192.168.1.1:/mnt/data/fwupdate.bin
- Enter your SSH password when prompted.
- SSH into your device.
- Run the following command:
ubnt-systool fwupdate /tmp/fwupdate.bin Version 1.12 and below:
 ubnt-upgrade /mnt/data/fwupdate.bin
USG (Internet)
- Copy the update link from community.ui.com/releases.
- SSH into your device.
- Run the following command:
upgrade <paste_download_link_here> Example:
 upgrade https://dl.ui.com/unifi/firmware/UGW3/4.4.56.5449062/UGW3.v4.4.56.5449062.tar
USG (No Internet)
- Download the desired firmware update from community.ui.com/releases.
- Rename the file: upgrade.tar
- Use the following SCP command to copy the file into the /home/<user> folder of your USG from your workstation. This requires a compatible SCP application (e.g., Terminal on macOS and Linux, PuTTY/PowerShell on Windows).
scp /folder_path/upgrade.tar <user>@<IP of device>:/home/<user>/upgrade.tar Example:
 scp /Users/alexpro/Desktop/upgrade.tar Alex@192.168.1.1:/home/Alex/upgrade.tar
- Enter your SSH password when prompted.
- SSH into your device.
- Run the following command:
sudo syswrapper.sh upgrade upgrade.tar
Manually Update the Network Application
- Download the desired application update from community.ui.com/releases.
- SSH into your device.
- Run the following command (UDM/UDM Pro if still running UniFi OS 1.x):
unifi-os shell
- Remove previously installed files:
rm /tmp/unifi_sysvinit_all.deb &> /dev/null
- Store the new application version on your device using the download link:
curl -o "/tmp/unifi_sysvinit_all.deb" <network application link.deb> Example:
 curl -o "/tmp/unifi_sysvinit_all.deb" https://dl.ui.com/unifi/6.2.26-a79cb15f05/unifi_sysvinit_all.deb
- Once downloaded, install the new version:
apt-get install -y /tmp/unifi_sysvinit_all.deb
- Following installation, remove the downloaded file:
rm /tmp/unifi_sysvinit_all.deb
Updating Devices in a Broken State
In rare occurrences, a device may stop functioning. UniFi devices may be updated using Recovery Mode. This should only be used as a last resort prior to submitting an RMA, if your device completely stops functioning.
