---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-unifi-ssh-guide-md-e3f88ff1-3
title: "Newer firmware - APs and switches"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-unifi-ssh-guide-md-e3f88ff1.md
source_anchor: ""
source_lines: [190, 275]
sha256: cf738ab4d75cd1c2832e828d5711d8a47b41f55356bac40a6e79da80e26f980a
---

# On UDM Pro / UDM SE (podman-based)
unifi-os shell
# On older UDM (docker-based)
docker exec -it unifi-network-application /bin/bashHost udm
    HostName <GATEWAY-IP>
    User root
    IdentityFile ~/.ssh/id_ed25519_unas
    Port 22
UniFi switches run proprietary firmware with a restricted shell environment. Full Linux functionality is not available. SSH access is primarily useful for diagnostics, checking link state, and running limited show commands.
ssh admin@<SWITCH-IP>
Username is admin. Password is the SSH password set in UniFi Network > Settings > System > Advanced.
Note: Some switch firmware versions use ubnt as the default username on factory reset devices that have not yet been adopted.
ssh ubnt@<SWITCH-IP>
# Default password: ubnt
| Command | Function | 
|---|---|
| info | Show system information | 
| show interfaces | Show interface status | 
| show mac-address-table | Show MAC address table | 
| show spanning-tree | Show spanning tree status | 
| show lldp neighbors | Show LLDP neighbours | 
| reboot | Reboot the switch | 
Full Linux commands such as ip, ss, systemctl, or package management are not available on switch firmware.
Key-based authentication is not reliably supported across all USW firmware versions. Password authentication should be treated as the primary method for switch access.
Access points run a BusyBox or OpenWrt-based environment. SSH access is available primarily for diagnostics.
ssh admin@<AP-IP>
Same SSH password as all other adopted devices. Username ubnt with password ubnt applies to factory reset or unadopted devices.
# Show system info
info
# Show wireless interface status
iwconfig
# Show connected clients
cat /proc/net/arp
# Show current channel and frequency
iwlist ath0 channel
# Restart the AP management agent
syswrapper.sh restart
# View logs
logread
Key authentication is not persistently supported on access points. Keys written to ~/.ssh/authorized_keys will be lost on reboot as the AP filesystem is largely volatile. Password authentication is the practical method for AP access.
SSH is not enabled on the device or the device has not been adopted by the controller.
Verify SSH is enabled: UniFi Network > Settings > System > Advanced > Device SSH Authentication.
The password entered does not match the SSH password in the controller. Confirm the password in UniFi Network settings. On factory reset devices, try ubnt / ubnt.
The public key is not present in authorized_keys on the target device, or the file permissions are incorrect.
Verify on the target device:
ls -la ~/.ssh/
cat ~/.ssh/authorized_keys
Correct permissions:
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
The authorized_keys file is not stored on the rwfs overlay. Follow the persistent key storage steps in the UNAS Pro section above.
Expected behaviour on gateways and APs. Re-copy the public key after any firmware update:
ssh-copy-id -i ~/.ssh/id_ed25519_unas.pub root@<DEVICE-IP>
Occurs when a device has been factory reset or replaced but the same IP is reused. Remove the stale entry from the known hosts file:
ssh-keygen -R <DEVICE-IP>ssh -vvv root@<DEVICE-IP>
The triple verbose flag shows the full authentication negotiation and will identify exactly where a connection is failing.
A consolidated ~/.ssh/config block for the ORION environment:
Host unas
    HostName 10.2.60.x
    User root
    IdentityFile ~/.ssh/id_ed25519_unas
    Port 22
Host udm
    HostName 10.2.1.x
    User root
    IdentityFile ~/.ssh/id_ed25519_unas
    Port 22
Host sw-core
    HostName 10.2.1.x
    User admin
    Port 22
Host sw-access
    HostName 10.2.1.x
    User admin
    Port 22
Replace IP addresses with actual device addresses. Key authentication entries can be omitted for switches where password auth is the only reliable method.
---
## References
---
- UniFi Default SSH Credentials: https://www.unihosted.com/blog/default-password-in-unifi-devices
- USW-Ultra Adoption Challenge: https://deluisio.com/networking/unifi/2025/02/05/the-missing-ssh-unifi-ultra-switches-and-the-adoption-challenge/
- Cloudflare API - Update DNS Record: https://developers.cloudflare.com/api/resources/dns/subresources/records/methods/update/
- UniFi - Required Ports Reference: https://help.ui.com/hc/en-us/articles/204976094-UniFi-Network-Required-Ports-Reference
- UniFi - Connecting with Debug Tools and SSH: https://help.ui.com/hc/en-us/articles/204909374-Connecting-to-UniFi-with-Debug-Tools-SSH
- OpenSSH Key Types - OpenBSD man page: https://man.openbsd.org/ssh-keygen
- UniFi OS SSH Key Persistence - Community Discussion: https://community.ui.com/questions/SSH-Key-Authentication/a25be39d-9873-4f3e-9a2b-5870ef399d65
