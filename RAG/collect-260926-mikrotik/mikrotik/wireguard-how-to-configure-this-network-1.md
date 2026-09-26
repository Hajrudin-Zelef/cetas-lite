---
id: collect-260926-mikrotik/mikrotik/wireguard-how-to-configure-this-network-1
title: "wireguard-how-to-configure-this-network"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-how-to-configure-this-network.md
source_anchor: ""
source_lines: [1, 188]
sha256: fb276fc53ff4511ba72739b6dc522744c032d550ad01d65dc3d80e162f00db5b
---

# wireguard-how-to-configure-this-network

Hello! I started with Mikrotik not so long ago and still have a poor understanding of routes and other things. Please help me create a network configuration like this:

There are 3 networks and a server with Wireguard installed. I need:

- Users from the network 172.20.21.0/24 have access to the server and network 172.20.20.0/24
 
 
- Users from the network 172.20.20.0/24 have access to the server and network 172.20.21.0/24
 
 
- Users from the 172.20.22.0/24 network can access servers and networks on both 172.20.20.0/24 and 172.20.21.0/24

Additional Internet access (by routing marks):

- Users from network 172.20.20.0/24 via VPS
 
 
- Users from network 172.20.21.0/24 via VPS
 
 
- Users from network 172.20.22.0/24 via router 0 (priority 1) and via VPS (priority 2)

VPS Wireguard config like this:

```
[Interface]
Address = 10.66.66.1/24,fd42:42:42::1/64
ListenPort = 12345
PrivateKey = xxxxx
PostUp = iptables -I FORWARD -i eth0 -o wg0 -j ACCEPT
PostUp = iptables -I FORWARD -i wg0 -j ACCEPT
PostUp = iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown = iptables -D FORWARD -i eth0 -o wg0 -j ACCEPT
PostDown = iptables -D FORWARD -i wg0 -j ACCEPT
PostDown = iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
[Peer]
PublicKey = xxxxx
PresharedKey = xxxxx
AllowedIPs = 10.66.66.2/32,fd42:42:42::2/128
[Peer]
PublicKey = xxxxx
PresharedKey = xxxxx
AllowedIPs = 10.66.66.3/32,fd42:42:42::3/128
[Peer]
PublicKey = xxxxx
PresharedKey = xxxxx
AllowedIPs = 10.66.66.4/32,fd42:42:42::4/128
```

             
            
           
          
            
            
              Since I dont know what VPS is the diagram makes no sense to me.

Why do all routers have a hard connection to this magical VPS

What is a square box not named a router.

It appears you have two different devices wired together with connections to the internet.

Nothing makes sense to me.;

             
            
           
          
            
            
              
Sorry, I fixed: square box is Router #1

VPS: Virtual Private Server, external web-server with white IP address, where I run WG and my branches  (2/3 with grey IPs) connected to. And it is not a hard connections, it thorugh the internet, sorry for my diagram

             
            
           
          
            
            
              I still dont see which routers have internet access directly it appears at the  moment only Router 0 ???

             
            
           
          
            
            
              Is VPS the only wireguard device?  If not would need config of all MT devices (full)

             
            
           
          
            
            
              **VPS Discussion**

(1) Looking at the VPS it is seemingly configured properly as a WG Server for the initial handshake.  The only comment I would make is ensure the single peer on the local Routers, to the VPS uses the nomenclature for wireguard as follows:  **10.66.66.0/24**.

(2) I dont see any allowance for any of the routers using WG to reach **subnets on** other routers other than via wireguard IP.

Not sure how iptable handles the wireguard to wireguard traffic but this single rule covers ALL what I would call **RELAY traffic** through the VPS, both wireguard and subnet!!  If  you do end up using other subnets across routers, then the below rule also covers that.

**add action=accept chain=forward in-interface=WG-interface  out-interface=WG-interface**

(3) If using wg to reach subnets on other routers, then you need more work on VPS but not for firewall routes but for routes.

Every subnet that is remote to the VPS coming over wireguard needs a route into the tunnel this covers both destination traffic and return traffic.

*/ip route
**add dst-address=172.168.20.0/24  gwy=wireguard-interface  table=main
add dst-address=172.168.21.0/24  gwy=wireguard-interface  table=main
add dst-address=172.168.22.0/24  gwy=wireguard-interface  table=main**
(4) If you want to be able to configure the VPS from any router or even a REMOTE wireguard user ( aka admin from laptop at hotel etc.)
**On VPS**
add chain=input action=accept in-interface=wireguard-interface **src-address-list=Authorized**
/ip firewall address
add address=10.66.66.5  list=**Authorized**  comment=“remote RW connection laptop”
add address=10.66.66.6  list=**Authorized**  comment=“remote RW connection ipad/iphone”
add address=172.16.20.X list=**Authorized** comment=“router0 admin Desktop IP”
add address=172.16.21.Y  list=**Authorized** comment=“router1 admin Desktop IP”
add address=172.16.22.Z   list=**Authorized** comment=“router2 admin Desktop IP”
Note for each PEER on VPS then logically you would add the IP allowed address…
[Peer - RO]
AllowedIPs = 10.66.66.2/32,**172.16.20.X,** fd42:42:42::2/128
[Peer-R1]
AllowedIPs = 10.66.66.3/32,**172.16.21.**Y,fd42:42:42::3/128
[Peer-R2]
AllowedIPs = 10.66.66.4/32,**172.16.22.Z**,fd42:42:42::4/128
[Peer-RW1]
AllowedIPs = 10.66.66.5/32
[Peer-RW2]
AllowedIPs = 10.66.66.6/32
Other Routers - simply ensure that the adminIP or the subnet the admin is on,  can access the local wireguard interface via a forward chain rule.
If traversing VPS towards other remote subnets, or remote subnets are coming to your local router, then you will need similar routes as VPS ones, but on the local router.
(5) If you want to be able to configure Other Routers (assuming all MT) via WG then nothing else is required on VPS.
On each router you would need to add the two other subnets as allowed IPs on the single local PEER to the VPS.  This would permit traffic both ways as well.
On each router you would need probably the same firewall **list=Authorized**  for a local input chain rule with in-interface=wireguard-interface
As you can see its a matter of ensuring…  forward chain rules allow traffic, ip routes direct traffic,  wg rules allow matching to a peer outbound(enter tunnel) and filtering a peer inbound (exit tunnel)
++++++++++++++++++++++++++++++++++++++++++++++++++++
Last is the internet access needs.  Need more clarity on the requirements to complete…
Almost there but there are potential complications.
Q1 -  some users from R0 need internet via VPS or entire subnet?  if some, how many?
Q2 - some users from R1  need internet via VPS or entire subnet?  if some, how many?
Q3.  For R2 router,  is there only one group of users needing special routing or two?
Q4.  If it is only one group is it some users or the entire subnet?
Q5.  If one,  what do you mean priority 0 to R0 and priority 1 to VPS,  ie,  primary and secondary in case R0 is not linked to VPS for some reason and not reachable???
Q6.  if two,  which users are going to R0 and which are going to VPS?*

             
            
           
          
            
            
              

I made a simpler diagram. All routers are connected to the internet. All routers are connected to the VPS using WireGuard over the internet.

The VPS is a Debian based server with Wireguard software (not a ROS based device). I pointed the circle as the gate that should provide Internet access (in green or purple direction)

All routers now configured as:

```
/ip address
add address=10.66.66.x/24 interface=wireguard1 network=10.66.66.0
/interface wireguard
add listen-port=12345 mtu=1420 name=wireguard1
/interface wireguard peers
add allowed-address=0.0.0.0/0 endpoint-address=VPS_IP endpoint-port=12345 interface=wireguard1 public-key=xxxxxx
```

Also I added each subnet to the [PEER] section on the VPS config almost as you wrote (I used full subnet /24 here):

```
[Peer - RO]
AllowedIPs = 10.66.66.2/32,172.16.20.0/24, fd42:42:42::2/128
[Peer-R1]
AllowedIPs = 10.66.66.3/32,172.16.21.0/24,fd42:42:42::3/128
[Peer-R2]
AllowedIPs = 10.66.66.4/32,172.16.22.0/24,fd42:42:42::4/128
```

Now I can ping, for example server 172.20.20.8 from VPS 10.66.66.1. But can not reach it from the 172.20.22.0/24 subnet.

