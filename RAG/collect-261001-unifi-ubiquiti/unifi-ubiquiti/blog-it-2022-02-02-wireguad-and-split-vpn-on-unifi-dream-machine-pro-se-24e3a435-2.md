---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435-2
title: "This script downloads the latest split-vpn and installs it"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-it-2022-02-02-wireguad-and-split-vpn-on-unifi-dream-machine-pro-se-24e3a435.md
source_anchor: ""
source_lines: [71, 152]
sha256: c24a1381b432077d0521125d2bf36aed0d1bbdc585d1ea22da6468fafbf6f0bc
---

# This script downloads the latest split-vpn and installs it

ln -sf "${DATA_DIR}/split-vpn" /etc/split-vpn 
echo split-vpn has been installed to "${DATA_DIR}/split-vpn" and linked to /etc/split-vpn.
- For the UDM, UDM Pro, UDM Pro SE, and UXG Pro, the script will be installed to /mnt/data/split-vpn .
- The installation will also link the script directory to /etc/split-vpn , which will be used for configuration below.root@udm:~# ls -la /etc/split-vpn
lrwxrwxrwx 1 root root 19 Jan 23 08:59 /etc/split-vpn -> /mnt/data/split-vpn/
3. Now you can follow instructions below to set-up the script.
WireGuard kernel module for Unifi OS Permalink
WireGuard support was added with UniFi OS v3.x. Thus, this part can be forgotten if your version of Unifi OS is greater than or equal to 3.0.13 version.
Install Permalink
1. We first need to download the tar file onto the UDM Pro SE. Connect to it via SSH and type the following command to download the tar file. You need to download the following tar file.
Always check this link for the latest release.
curl -LJo wireguard-kmod.tar.Z https://github.com/tusc/wireguard-kmod/releases/download/v01-22-22/wireguard-kmod-01-22-22.tar.Z
2. Type the following to extract the files:
tar -C /mnt/data -xvzf wireguard-kmod.tar.Z
3. Once the extraction is complete, go to /mnt/data/wireguard folder and run the script setup_wireguard.sh as shown below
cd /mnt/data/wireguard
chmod +x setup_wireguard.sh
./setup_wireguard.sh
This will setup the symbolic links for the various binaries to the /usr/bin path as well as create a symlink for the /etc/wireguard folder and finally load the kernel module. You’ll want to run dmesg | grep wireguard to verify the kernel module was loaded. You should see something like the following:
[   22.280358] WireGuard: WireGuard 1.0.20210606 loaded. See www.wireguard.com for information.
[   22.280361] WireGuard: Copyright (C) 2015-2019 Jason A. Donenfeld <Jason@zx2c4.com>. All Rights Reserved.
The script will first try to load the built-in WireGuard module if it exists. If it doesn’t exist, the external module provided by this package will be loaded instead. You can set LOAD_BUILTIN=0 at the top of the setup_wireguard.sh script to always load the external module. Note that only recent UDM releases since 1.11.0 have the built-in module, and it is not always up-to-date.
The tar file includes other useful utils such as htop, iftop and qrencode.
Before continuing, test the installation of the module by running modprobe wireguard which should return nothing and no errors, and running wg-quick which should return the help and no errors.
- After you load the module, run ip link add dev wg0 type wireguard to test if you can add a WireGuard interface successfully. If your UDM locks up and restarts when you do this, then the module is not compatible with your kernel. Check thewireguard-kmod github or your custom kernel for more information.
- If adding the interface succeeded, type ip link del wg0 to delete the WireGuard interface before continuing with the steps below.
The kernel module is dependent on the software version of your UDM Pro because each software update usually brings a new kernel version. If you update the UDM Pro software, you also need to update the kernel module to the new version once it is released (or compile your own module for the new kernel). The module will fail to run on a kernel it was not compiled for. Hence, you have to be careful that the UDM Pro doesn’t perform a sofware update unexpectedly if you use this module.
Surviving at Reboots Permalink
You will need to run setup_wireguard.sh whenever the UDM is rebooted as the symlinks have to be recreated.
For the UDM Pro SE, create a systemd boot service to run the setup script at boot by running the following commands:
curl -Lo /etc/systemd/system/setup-wireguard.service https://raw.githubusercontent.com/tusc/wireguard-kmod/main/src/boot/setup-wireguard.service
systemctl daemon-reload
systemctl enable setup-wireguard
[Unit]
Description=Run WireGuard setup script
Wants=network.target
After=network.target
[Service]
Type=oneshot
ExecStartPre=sh -c 'DIR="$(find /mnt/data/wireguard /data/wireguard -maxdepth 1 -type d -name "wireguard" 2>/dev/null | head -n1)"; ln -sf "$DIR/setup_wireguard.sh" /etc'
ExecStart=/etc/setup_wireguard.sh
[Install]
WantedBy=multi-user.target
Note this only adds the setup script to start at boot. If you also want to bring up your WireGuard interface at boot, you will need to add another boot script with your wg-quick up command.
Wireguad® Configuration: On your VPN service provider Permalink
1. Prererquiste/Configuration need by your VPN service provider.
 Your VPN service provider can require some specific configuration. This configuration may be necessary at the level of your VPN provider, but also at the level of your VPN client router, which in our case is the UDM PRO SE. In this part I will only talk about the setting required to use IVPN’s WireGuard® service.
- SSH into your router (UDM PRO SE) as root
- And generate a WireGuard keys:
cd
wg genkey | tee privatekey | wg pubkey > publickey
chmod 600 privatekey
- Note your Private & your Public keys, you will need them later:
cat privatekey
cat publickey
2. Obtain WireGuard IP address from IVPN
- Log into the Client Area
- Navigate to WireGuard tab and click theAdd a new key button
3. Copy and paste the Public key obtained previously, give it any name, then click the Add key button and note the assigned IP address
Wireguad® configuration: On router Permalink
1. Create a directory for your WireGuard configuration files, copy the sample vpn.conf from /etc/split-vpn/vpn/vpn.conf.sample, and copy your WireGuard configuration file (wg0.conf) or create it. As an example below, we are creating the wg0.conf file that IVPN provider and pasting the contents into it. You can use any name for your config instead of wg0 (e.g.: ivpn-fr.conf) and this will be the interface name of the WireGuard tunnel.
mkdir -p /etc/split-vpn/wireguard/ivpn
cd /etc/split-vpn/wireguard/ivpn
cp /etc/split-vpn/vpn/vpn.conf.sample /etc/split-vpn/wireguard/ivpn/vpn.conf
vi wg0.conf
2. In your WireGuard config (wg0.conf) file, set PostUp and PreDown to point to the updown.sh script, and Table to a custom route table number that you will use in this script’s vpn.conf. Here is an example wg0.conf file:
[Interface]
PrivateKey = <router's private key>               # Is the private key generated above for your router (UDM PRO SE)
Address = 172.20.xxx.yyy/32                       # Is Wireguard IP Address obtained above in the IVPN Client Area
PostUp = sh /etc/split-vpn/vpn/updown.sh %i up
PreDown = sh /etc/split-vpn/vpn/updown.sh %i down
Table = 101
[Peer]
PublicKey = g7BuMzj3r<redacted>                   # Is the public key of your Wireguard IP VPN server `Endpoint`
AllowedIPs = 0.0.0.0/1,128.0.0.0/1
Endpoint = <wireguard service provider IP>:2049   # Is IP of Wireguard service provider
In the above config, make sure to:
- Comment out or remove the DNS line. Use the DNS settings in your vpn.conf file instead if you want to force your clients to use a certain DNS server.
- Set AllowedIPs to0.0.0.0/1,128.0.0.0/1,::/1,8000::/1 to allow allIPv4 andIPv6 traffic through the VPN. Do not use0.0.0.0/0,::/0 because it will interfere with theblackhole routes and won’t allow WireGuard to start. If you prefer to use0.0.0.0/0,::/0 , disableblackhole routes by settingDISABLE_BLACKHOLE=1 in yourvpn.conf file so WireGuard can start successfully.
- Address : Set here the WireGuard IP Address obtained in the IVPN Client Area ending with/32 (e.g.172.20.xxx.yyy/32 ). See above bullet 3.
- Remove any extra PreUp /PostUp /PreDown /PostDown lines that could interfere with the VPN script.
- You can remove or comment out the PreUp line if you do not want VPN-forced clients to lose Internet access if WireGuard does not start correctly.
