---
id: collect-261001-cisco/cisco/configuring-cisco-dynamic-multipoint-vpn-dmvpn-hub-spokes-mgre-protection-and-routing-dmvp-2
title: "Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- ------------- --------------- ----- ------- ----- 1 2.2.2.10 172.16.0.2 UP 00:04:58 D 1 3.3.3.10 172.16.0.3 UP 00:04:12 D"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/configuring-cisco-dynamic-multipoint-vpn-dmvpn-hub-spokes-mgre-protection-and-routing-dmvpn-configur.md
source_anchor: ""
source_lines: [219, 337]
sha256: b1e21d952710963fab1b55ed4f457ad5bbe8c7b14b3b42e90b321ce3061862cf
---

# Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- ------------- --------------- ----- ------- ----- 1 2.2.2.10 172.16.0.2 UP 00:04:58 D 1 3.3.3.10 172.16.0.3 UP 00:04:12 D

R1# show dmvpn Legend: Attrb --> S - Static, D - Dynamic, I - Incomplete N - NATed, L - Local, X - No Socket # Ent --> Number of NHRP entries with same NBMA peer NHS Status: E --> Expecting Replies, R --> Responding UpDn Time --> Up or Down Time for a Tunnel ==========================================================================

# Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- ------------- --------------- ----- ------- ----- 1 2.2.2.10 172.16.0.2 UP 00:04:58 D 1 3.3.3.10 172.16.0.3 UP 00:04:12 D

The output of our command provides us with some valuable information. To start with, the router provides an explanation for each column presented (right under the show command) but we are still going to cover them so that we are not left with any unanswered questions.

The first column #Ent shows the number of entries that exist in the NHRP Database for the same spoke. Usually, we wouldn’t expect to see more than one for each spoke.

The second column Peer NBMA Addr presents the spoke’s public IP address, while the third column, Peer Tunnel Add, shows each spoke’s local Tunnel’s IP address.

Next, the State column shows the current state the tunnel is in. In our case, both tunnels are UP. Right next to the State is the UpDN Tm, which is the Up or Down Time of the current State. This is a very important bit of information as you can clearly see out how long your tunnel has been in its current state.

For our example, both spokes have been up for almost 5 minutes.

Lastly, the Attrib column shows the type of tunnels established by the spokes. D stands for Dynamic, S for Static and I for Incomplete. Usually dynamic spokes will create D type tunnels. Tunnels established from the spokes to the Hub router are expected to be S type, since the Hub remains static.

Verifying DMVPN Functionality At The R2 & R3 Spoke Router

Turning to R2 router, our first spoke, we can repeat the same show dmvpn command and obtain a list of dmvpns currently created:

R2# show dmvpn Legend: Attrb --> S - Static, D - Dynamic, I - Incomplete N - NATed, L - Local, X - No Socket # Ent --> Number of NHRP entries with same NBMA peer NHS Status: E --> Expecting Replies, R --> Responding UpDn Time --> Up or Down Time for a Tunnel ==========================================================================

# Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- ------------- -------------- ----- ------- ----- 1 1.1.1.10 172.16.0.1 UP 00:06:35 S

As expected, R2’s output shows one entry only. When traffic needs to be directed to R3, a second GRE tunnel will come up. We’ll try this soon. For now let’s check our third remote site, R3 spoke router

Using the same show dmvpn command we obtain the following similar output:

R3# show dmvpn Legend: Attrb --> S - Static, D - Dynamic, I - Incomplete N - NATed, L - Local, X - No Socket # Ent --> Number of NHRP entries with same NBMA peer NHS Status: E --> Expecting Replies, R --> Responding UpDn Time --> Up or Down Time for a Tunnel ==========================================================================

# Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- --------------- --------------- ----- -------- ----- 1 1.1.1.10 172.16.0.1 UP 00:06:55 S

Protecting - Encrypting DMVPN mGRE Tunnels With IPSec

Since we have our GRE tunnels up and running, we need to encrypt them using IPSec to ensure data confidentiality. Protecting GRE Tunnels is covered in great depth in our Protected GRE over IPSec article, so we are going to simply display the commands here without repeating the topic.

Notice the command crypto isakmp key firewall.cx address 0.0.0.0 0.0.0.0. The peer address for which the isakmp key is valid is 0.0.0.0 0.0.0.0, which means every possible host on the Internet. When our remote routers (spokes) have dynamic IP addresses, 0.0.0.0 0.0.0.0 must be used.

The following configuration applies to R2 & R3 spoke routers:

Again we’ve defined 0.0.0.0 0.0.0.0 as the isakmp peer address. While the hub’s public IP address is known we must keep in mind that R2 and R3 can build dynamic VPN tunnel between them. Taking into consideration that their public IP address is dynamic it is imperative to use 0.0.0.0 0.0.0.0 for the remote peer.

Verifying the DMVPN Crypto Tunnels

Once all routers are configured IPSec VPN tunnels are brought up. We can verify this by using the show crypto session command at our R1 hub router:

R1# show crypto session

Crypto session current status

Interface: Tunnel0

Session status: UP-ACTIVE

Peer: 2.2.2.10 port 500

IKE SA: local 1.1.1.10/500 remote 2.2.2.10/500 Active

IPSEC FLOW: permit 47 host 1.1.1.10 host 2.2.2.10

Active SAs: 2, origin: crypto map

Interface: Tunnel0

Session status: UP-ACTIVE

Peer: 3.3.3.10 port 500

IKE SA: local 1.1.1.10/500 remote 3.3.3.10/500 Active

IPSEC FLOW: permit 47 host 1.1.1.10 host 3.3.3.10

Active SAs: 2, origin: crypto map

Routing Between DMVPN mGRE Tunnels

Last step involves enabling routing in our DMVPN network. This is required so that the hub and spoke routers are aware which packets need to be sent via the VPN network.

There are two ways this can be achieved: 1) Static routes 2) Routing protocol.

For the sake of simplicity we are going to focus on static routes. DMVPN and routing protocol configuration will be covered in another article.

Configuring the necessary static routes is very simple. All that is required is a set of simply static routes on each router (hub and spoke), pointing to the other networks.

On the R1 hub router:

ip route 192.168.2.0 255.255.255.0 172.16.0.2 ip route 192.168.3.0 255.255.255.0 172.16.0.3

On R2 spoke router:

ip route 192.168.1.0 255.255.255.0 172.16.0.1 ip route 192.168.3.0 255.255.255.0 172.16.0.3

And finally on R3 spoke router:

ip route 192.168.1.0 255.255.255.0 172.16.0.1 ip route 192.168.2.0 255.255.255.0 172.16.0.2

Our DMVPN Network Is Ready!

At this point, our DMVPN network is ready and fully functional. All networks are connected between each other and dynamic VPN tunnels between spokes can be established. GRE tunnels are protected properly, providing data confidentiality and ip routing is enabled.

As a final step, we can try sending traffic between the spokes and verify the dynamic tunnel is being established:

From R2 spoke router, we try to ping R3’s LAN IP address:

R2# ping 192.168.3.1

Type escape sequence to abort. Sending 5, 100-byte ICMP Echos to 192.168.3.1, timeout is 2 seconds: .!!!! Success rate is 80 percent (4/5), round-trip min/avg/max = 1/1/4 ms

It is evident that the two spoke routers have established communication.

The DMVPN is up and routing is working perfectly:

R2# show dmvpn Legend: Attrb --> S - Static, D - Dynamic, I - Incomplete N - NATed, L - Local, X - No Socket # Ent --> Number of NHRP entries with same NBMA peer NHS Status: E --> Expecting Replies, R --> Responding UpDn Time --> Up or Down Time for a Tunnel ==========================================================================

# Ent Peer NBMA Addr Peer Tunnel Add State UpDn Tm Attrb ----- ------------------ --------------- ----- ------- ----- 1 1.1.1.10 172.16.0.1 UP 00:39:05 S 1 3.3.3.10 172.16.0.3 UP 00:00:08 D

This concludes our DMVPN configuration article.

This article showed how to configure a DMVPN network between Cisco routers. We covered the configuration of a Cisco DMVPN including Hub, Spokes, Static Routing and Protecting the mGRE Tunnel. We also provided some useful show commands to help troubleshoot and debug the DMVPN network. More articles on VPN & DMVPN can be found in our Cisco Routers Section and Cisco Services & Technlogies Section.
