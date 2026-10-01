---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6-1
title: "setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-unifi-ubiquiti/setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6.md
source_anchor: ""
source_lines: [1, 141]
sha256: c257acec465ff26d8d24ca17e3a42206ce3dffca02d52b4f1663b2770611a747
---

# setup-a-wireguard-vpn-on-unifi-dream-machine-udm-udm-pro-and-use-macos-as-a-clie-b436e7b6

WireGuard is quite a hyped new VPN server that has some pros why you would like to use it. Its part of the Linux kernel now and therefore you can setup a WireGuard VPN server on almost every piece of hardware. 
But some people might want to run it on their router, and if you are using Unifi gear, they did not support it official yet as of date end of 2021. Its unclear when Ubiquiti will add WireGuard to their systems, but in the meantime there is a relatively easy way to install and run WireGuard directly on your UDM/UDM-Pro. Here is how:
1. Install latest kernel extension
First we need to install the latest kernel extension from github (check this release page for newer versions).
ssh into your UDM and make sure you are on the bare underlying OS, and not in the unifi-os shell mode. To be sure you can first type exit into the console and of you only get a # back, you are in the right mode.
Now we are running a couple of commands. We start by making sure there is no earlier installation of wireguard and remove maybe leftovers:
rm -rf /mnt/data/wireguard
Next cd into the data folder and download the wireguard udm package from github:
Important: Please check before copying the below command for a more up to date release url!
Edit: link updated Oct 2022
cd /mnt/data && curl -LJo wireguard-kmod.tar.Z 
https://github.com/tusc/wireguard-kmod/releases/download/v09-24-22/wireguard-kmod-09-24-22.tar.Z
Now when the download is finished we unpack the archive to the data folder:
tar -C /mnt/data -xvzf wireguard-kmod.tar.Z
Next we go to the newly extracted and created wireguard folder and set the right permissions to run the scripts:
cd /mnt/data/wireguard
chmod +x setup_wireguard.sh
Now we can launch the wireguard vpn service with:
./setup_wireguard.sh
To verify that it is successfully running enter:
dmesg
Scroll down to the very end of the output and the last two lines need to show something like:
[3315037.792695] wireguard: WireGuard 1.0.20210606 loaded. See www.wireguard.com for information.
[3315037.792699] wireguard: Copyright (C) 2015-2019 Jason A. Donenfeld <Jason@zx2c4.com>. All Rights Reserved.
Congratulations, you are running a WireGurd server on your UDM/UDM-Pro now!
2. Make the WireGuard server install persistent
Unfortunately will the above install not survive any reboots, as the symlinks set by the setup_wireguard.sh install script are gone after every UDM reboot.
To make the install persistent, we are using a github script by boostchicken that enables us to autorun any script on boot. This way we are always automatically executing the setup_wireguard.sh script on boot.
To install the script we have to switch into the unifi-os shell mode by entering in the console:
unifi-os shell
In the unifi-os shell mode download the latest udm-boot_1.x.x_all.deb release and install it:
cd /home/
curl -L https://udm-boot.boostchicken.dev -o udm-boot_1.0.5_all.deb
dpkg -i udm-boot_1.0.5_all.deb
exit
Don’t forget the last exit command so you are back on the base UDM to continue.
Now we have a new auto execute folder in /mnt/data/on_boot.d and whatever script we will put in there, will be executed automatically on boot of our UDM. Pretty cool 🙂
Next we move the wireguard setup script into our new on_boot directory:
mv /mnt/data/wireguard/setup_wireguard.sh /mnt/data/on_boot.d/setup_wireguard.sh
Also make sure its executable:
chmod +x /mnt/data/on_boot.d/setup_wireguard.sh
Great, wiregurad will now run persistent on your UDM!
configuration the server
3. Configure the WireGuard server
First we need to create a wireguard config file:
cp /etc/wireguard/wg0.conf.sample /etc/wireguard/wg0.conf
Go into the folder and setup access rights and generate ssh keys:
cd /etc/wireguard/
umask 077
wg genkey | tee privatekey | wg pubkey > publickey
This will generate two files, privatekey and publickey on each of the computers. The publickey file is for telling the world, the privatekey file is secret and should stay on the computer it was generated on. Verify it by listing the results and you should see private key and public key files now:
ls /etc/wireguard/
Next you need to paste the contents of these files into the config file. For this we read out our servers public key with:
cat /etc/wireguard/publickey
Copy & paste the string result into a file and save it for later use on your client.
Now edit the wireguard config file with:
nano /etc/wireguard/wg0.conf
And add the following:
Important: make sure to change the variables to your server needs!
Address = the address/subnet wireguard will create for your connected clients and the IP of the server (in our case here 1)
ListenPort = wireguards vpn port
PublicKey = the publickey from your clients device. It is usually generated by the wireguard application on your client. Copy and paste that key over to this config.
AllowedIPs = the address/subnet wireguard will give to your connected client (in our case here .3) 
[Interface]
Address = 10.0.0.1/24
PrivateKey = YOUR-UDMP-WIREGUARD-PRIVATE-KEY
ListenPort = 51820
[Peer]
PublicKey = YOUR-CLIENT-PUBLIC-KEY
AllowedIPs = 10.0.0.3/32
Save and exit the file with STRG+X and confirm to save with Y.
4. Configure WireGuard client
Next we configure our client, in our case a Mac OS device. We download the Wireguard VPN app, install it and then add a new connection with this config file:
Important: make sure to change the variables to your clients needs!
PrivateKey = the private key of your client. This should be generated by your wireguard app automatically.
Address = your client IP address
PublicKey = the publickey from your server, that we generated earlier and saved for later use. 
Endpoint = the FQDN / public IP address of your wireguard server, followed by its port
AllowedIPs = the range of IPs / subnets that are allowed to be passed
[Interface]
PrivateKey = YOUR-CLIENT-WIREGUARD-PRIVATE-KEY
Address = 10.0.0.3/32
DNS = 10.0.0.1
[Peer]
PublicKey = YOUR-UDMP-WIREGUARD-PUBLIC-KEY
Endpoint = server.ipaddress.public:51820
AllowedIPs = 10.0.0.0/24
Save the connection in your clients app.
5. Setup your UDM to forward the wireguard network.
As a last step to make WireGuard work on your UDM(P), we have to open up the necessary ports and create firewall rules to Unifi:
Rule #1: Internet/WAN Local - forward external traffic to VPN server
- In the Network App, go to "Settings" -> "Security" -> "Internet Threat Management" -> "Firewall"
- Click "Create New Rule"
- Type: "Internet Local"
- Description: "WG: allow incoming VPN traffic"
- Enabled: oh yes
- Rule Applied: "Before predefined rules"
- Action: "Accept"
- IP4 protocol: "UDP"
- Source
  - Type: "Address/Port Group",
  - "Any",
  - "Any"
- Destination:
  - "Address/Port Group",
  - "Wireguard VPN Server IP 192.168.240.1",
  - "Wireguard VPN Port 51820"
Rule #2: LAN Out - allow external traffic to VPN server
- In the Network App, go to "Settings" -> "Security" -> "Internet Threat Management" -> "Firewall"
- Click "Create New Rule"
- Type: "Internet Local"
- Description: "WG: allow access to intranet"
- Enabled: oh yes
- Rule Applied: "Before predefined rules"
- Action: "Accept"
- IP4 protocol: "UDP"
- Source
  - Source Type: "Network",
  - Network: "VLAN240 Wireguard VPN",
  - Network Type: "IPv4 Subnet"
- Destination:
  - Destination Type: "Address/Port Group",
  - IPv4 Address Group: "RFC1918 (All Private Networks)",
  - Port Group: "Any"
Rule #3: LAN Out - allow internal traffic to VPN clients
- In the Network App, go to "Settings" -> "Security" -> "Internet Threat Management" -> "Firewall"
- Click "Create New Rule"
- Type: "Internet Local"
- Description: "WG: allow access to VPN clients"
- Enabled: oh yes
- Rule Applied: "Before predefined rules"
- Action: "Accept"
- IP4 protocol: "UDP"
- Source:
  - Source Type: "Address/Port Group",
  - IPv4 Address Group: "RFC1918 (All Private Networks)",
  - Port Group: "Any"
- Destination:
  - Destination Type: "Network",
  - Network: "VLAN240 Wireguard VPN",
  - Network Type: "IPv4 Subnet"
