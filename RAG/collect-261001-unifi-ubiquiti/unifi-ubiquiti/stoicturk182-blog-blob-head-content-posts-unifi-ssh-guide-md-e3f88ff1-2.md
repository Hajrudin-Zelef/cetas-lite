---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-unifi-ssh-guide-md-e3f88ff1-2
title: "Newer firmware - APs and switches"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-unifi-ssh-guide-md-e3f88ff1.md
source_anchor: ""
source_lines: [84, 189]
sha256: 0d2d9f9fcfd6fb3e4c80f7005baffdd7f66ce3b9bf7626d3d56265c8aad45c18
---

# Newer firmware - APs and switches

Port 8080 TCP must be open inbound on the controller site's gateway. If the controller is behind a pfSense or UniFi gateway, add a NAT port forward and corresponding firewall rule:
For pfSense:
Firewall > NAT > Port Forward
Interface: WAN
Protocol: TCP
Destination port: 8080
Redirect target IP: <CONTROLLER-LAN-IP>
Redirect target port: 8080
For UniFi gateway, add a port forward rule under Settings > Routing > Port Forwarding targeting the controller's LAN IP on port 8080.
After running set-inform on the remote device, confirm it is attempting to reach the controller:
# On the remote device
info
The info command on UniFi devices returns the current inform URL and status. Look for the inform URL reflecting the address you set and a status of connected or adopting.
If the device does not appear in the controller after 60-90 seconds, check:
# Confirm DNS resolves from the device
nslookup <FQDN>
# Confirm port 8080 is reachable from the device
nc -zv <FQDN> 8080
If nc fails, port 8080 is not reachable from the remote device — the issue is either a firewall rule on the controller site or the DDNS hostname is not resolving correctly.
Exposing port 8080 to the internet is required for remote adoption but represents an attack surface. Options to reduce exposure:
- Restrict inbound port 8080 to known remote site WAN IP ranges if they are static
- Use a VPN between sites instead of exposing port 8080 publicly, then use the VPN tunnel IP for set-inform
- Close port 8080 after adoption is complete if ongoing remote management uses an alternative path
Reference: UniFi - Inform Communication
SSH credentials must be configured before any of the methods below will work. Ubiquiti has moved this setting multiple times across recent versions. Use the table below to find the correct location for your controller version.
| Version | SSH Settings Location | 
|---|---|
| Pre-9.2.87 | Settings > System > Advanced > Device Authentication | 
| 9.2.87 - 9.5.20 | Devices > Device Updates and Settings (top right corner) | 
| 9.5.21+ | Devices > Device Updates and Settings (bottom left corner) | 
In all versions, the setting is labelled "Device SSH Authentication" and allows configuring a site-wide username and password that applies to all adopted devices. SSH keys can also be added from the same panel.
Note: UniFi OS SSH (the gateway console itself) and UniFi Network device SSH (switches, APs) are managed separately. Enabling one does not enable the other.
Reference: UniFi - Connecting with Debug Tools and SSH
The UNAS Pro runs a full Debian-based Linux environment. SSH behaves as it would on any standard Debian host. Key-based authentication is fully supported and can be made persistent across reboots using the rwfs overlay filesystem.
The UNAS Pro is a Tier 1 device and manages its own SSH daemon independently. Changing Device SSH Authentication credentials in the UniFi Network controller has no effect on the UNAS Pro. SSH access is controlled entirely by the local sshd_config and the local root password on the device itself.
ssh root@<UNAS-IP>
Password is the local root password set on the UNAS Pro itself. This is a Linux account password local to the device. If it was never explicitly set or is unknown, set it from an active SSH session:
passwd root
To enable password authentication if it has been disabled:
sed -i 's/^PasswordAuthentication no/PasswordAuthentication yes/' /etc/ssh/sshd_config
systemctl reload ssh
Key-based authentication is the most secure and practical method for regular access. The UNAS Pro supports all standard OpenSSH key types. Ed25519 is preferred over RSA for new keys due to smaller key size and equivalent security at lower computational cost.
Reference: OpenSSH key types comparison - OpenBSD man page
ssh-keygen -t ed25519 -C "orion-unas" -f ~/.ssh/id_ed25519_unas
The -C flag adds a comment for identification. The -f flag specifies the output filename to keep UNAS keys separate from other key pairs.
ssh-copy-id -i ~/.ssh/id_ed25519_unas.pub root@<UNAS-IP>
Or manually append to the authorized_keys file:
cat ~/.ssh/id_ed25519_unas.pub | ssh root@<UNAS-IP> "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"ssh -i ~/.ssh/id_ed25519_unas root@<UNAS-IP>
By default on the UNAS Pro, the root filesystem is partially read-only. Changes written to standard paths may not survive a reboot. The rwfs overlay provides a writable layer that persists across reboots.
Confirm the rwfs mount point:
mount | grep rwfs
The persistent path for SSH authorized keys on the UNAS Pro is typically:
/rwfs/root/.ssh/authorized_keys
Symlink the standard path to the persistent location:
mkdir -p /rwfs/root/.ssh
chmod 700 /rwfs/root/.ssh
cp ~/.ssh/authorized_keys /rwfs/root/.ssh/authorized_keys
chmod 600 /rwfs/root/.ssh/authorized_keys
ln -sf /rwfs/root/.ssh/authorized_keys /root/.ssh/authorized_keys
Verify after reboot that the key is still present:
cat /root/.ssh/authorized_keys
Note: Firmware updates on the UNAS Pro may reset the overlay or overwrite symlinks. Verify key persistence after any firmware update.
Add an entry to ~/.ssh/config on the client machine to avoid specifying the key and user on every connection:
Host unas
    HostName <UNAS-IP>
    User root
    IdentityFile ~/.ssh/id_ed25519_unas
    Port 22
Connect with:
ssh unas
Before relying on any SSH method, verify the daemon configuration is in the expected state:
grep -E "PasswordAuthentication|PubkeyAuthentication|PermitRootLogin" /etc/ssh/sshd_config
Expected output with both methods enabled:
PermitRootLogin yes
PasswordAuthentication yes
PubkeyAuthentication absent from output means it is defaulting to yes — this is normal on the UNAS Pro.
To view the full config:
cat /etc/ssh/sshd_config
The UNAS Pro sshd_config contains an include directive at the top:
Include /etc/ssh/sshd_config.d/*.conf
Any .conf files in this directory are loaded and can override the main config. A directive such as PasswordAuthentication no in an included file will override PasswordAuthentication yes in the main file. Always check this directory if authentication behaves unexpectedly:
ls /etc/ssh/sshd_config.d/
cat /etc/ssh/sshd_config.d/*.conf
The UNAS Pro may contain an AllowGroups directive in /etc/ssh/sshd_config.d/ — for example:
AllowGroups root
This restricts SSH access to users belonging only to the root group. Connecting as root is unaffected, but any other user account cannot SSH in unless added to that group. This is expected and correct for a single-admin homelab NAS.
Always test both methods from a second terminal before closing your active session:
# Test key auth
ssh -i ~/.ssh/id_ed25519_unas root@<UNAS-IP>
# Test password auth explicitly
ssh -o PreferredAuthentications=password -o PubkeyAuthentication=no root@<UNAS-IP>
Use reload rather than restart when applying config changes:
systemctl reload ssh
systemctl status ssh
Check which port SSH is listening on:
ss -tlnp | grep sshd
View active SSH sessions:
who
UniFi gateways run UniFi OS, a Debian-based environment with an additional application layer managing network services. SSH access reaches the UniFi OS shell, from which the UniFi Network application container can also be accessed.
ssh root@<GATEWAY-IP>
Password is the SSH password set in UniFi Network > Settings > System > Advanced. This is the same credential used across all adopted devices.
Key authentication is supported on UniFi gateways but persistence is more limited than on the UNAS Pro.
ssh-copy-id -i ~/.ssh/id_ed25519_unas.pub root@<GATEWAY-IP>
On UniFi gateways, ~/.ssh/authorized_keys is stored in the root filesystem which may be reset on firmware update. There is no officially documented persistent overlay equivalent to the UNAS Pro rwfs path. Keys survive normal reboots but should be re-applied after firmware updates.
Reference: UniFi OS SSH key persistence - community discussion
From the gateway SSH shell, the UniFi Network application runs in a container that can be accessed directly:
