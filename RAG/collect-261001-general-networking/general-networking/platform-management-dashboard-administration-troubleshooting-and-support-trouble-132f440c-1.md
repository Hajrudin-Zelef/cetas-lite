---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-troubleshooting-and-support-trouble-132f440c-1
title: "platform-management-dashboard-administration-troubleshooting-and-support-trouble-132f440c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["full-duplex"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-troubleshooting-and-support-trouble-132f440c.md
source_anchor: ""
source_lines: [1, 153]
sha256: cb81a09c57af677607c865f9536a211034792bd6dd29f76bdbbf23636e7d8488
---

# platform-management-dashboard-administration-troubleshooting-and-support-trouble-132f440c

## Overview

Packet loss occurs when data sent from one networked device to another fails to arrive. Packet loss can occur for a variety of reasons. When troubleshooting packet loss, first isolate where the loss is occurring.

## Troubleshooting packet loss detection

This section covers how to confirm whether packet loss is occurring, using tools available on a client PC and via the Cisco Meraki dashboard

### Troubleshooting steps

#### On a PC

1. Open a command prompt on a client PC via the **Start****Menu** and search for**cmd** .
2. Run the following ping command:

ping -n 20 8.8.8.8

The command sends 20 Internet Control Message Protocol (ICMP) requests to 8.8.8.8. Replace 8.8.8.8 with the address you want to test.

1. Review the summary to determine whether packet loss occurred.
2. If no loss occurred, increase the **-n** value (for example, to 100) to test over a longer period.

This only tests for packet loss impacting ICMP or all traffic. Protocol-specific loss may not be reflected.

#### Via the dashboard

Use the dashboard to monitor connectivity to a specific IP address. The MX WAN appliance continuously sends ICMP pings to the target IP address.

1. Navigate to **Security & SD-WAN** > **Monitor** > **Appliance Status** > **Uplink** to view the tool.

The following screenshot shows the historical packet loss graph for the WAN uplink. Use the graph to identify periods where packet loss occurred


1. Review the packet loss section under **Historical data** to check for ICMP packet loss while the WAN appliance was pinging 8.8.8.8.
2. To configure the destination IP address, navigate to **Security & SD-WAN** > **Configure** > **SD-WAN & Traffic Shaping** > **Uplink Statistics**

## Troubleshooting packet loss location on routed links

If packet loss occurs, identify where the loss begins. Use tracert to identify the Layer 3 (L3) device where packet loss begins

### Troubleshooting steps

1. Open a command prompt on a client PC via the **Start Menu** and search for**cmd** .
2. Run the following tracert command:

tracert  -d  8.8.8.8

The command traces the route to 8.8.8.8 and displays each hop as an IP address.

1. Wait for the trace to complete. Press CTRL/CMD+C if multiple lines display **Request timed out** .
2. An asterisk (*) represents a lack of response and indicates either packet loss or a device configured to not respond.
3. Run the test multiple times to identify where loss is occurring.
4. Frequent packet loss after a specific hop usually indicates a problem with that device or the link to the previous hop.

1. Test from multiple clients at different network locations to rule out client-specific issues and identify patterns among affected clients.

1. Use WinMTR to perform a continuous trace. WinMTR displays the packet loss percentage for each hop in the path.

The following screenshot shows packet loss percentages for each hop. Use the results to identify where packet loss begins

## Troubleshooting packet loss on WAN uplink

If packet loss occurs on the WAN uplink, determine whether the loss originates on the WAN appliance or the ISP network. Capture traffic on the LAN and Internet interfaces to identify the affected side.

### Troubleshooting steps

1. Run constant pings from a PC to a public IP address.
2. Take simultaneous packet captures on the LAN and WAN appliance.
3. Filter the traffic by source IP address, destination IP address, and ICMP.
4. Verify that the WAN appliance forwards ICMP requests from the LAN to the WAN.
5. If traffic is forwarded correctly, the WAN appliance is not causing the packet loss. Continue troubleshooting the ISP connection.

## Troubleshooting packet loss in a wireless or switched network

The tracert tool only provides information for Layer 3 (L3) devices in the path, such as routers. If packet loss occurs at the first hop and traffic must pass through a wireless access point (AP) and switch to reach that hop, perform additional testing to isolate the problem.

Repeat the tests while progressively moving closer to the L3 device.

### Troubleshooting steps

1. Ping the AP to test wireless quality.
    
  - If using a Cisco Meraki AP, ping my.meraki.com.
  - If packet loss begins at this stage, see Troubleshooting wireless performance.
2. Ping a client connected to the same VLAN, if configured, on the switch to which the wireless client is connected.
    
  - If multiple switches exist along the path, repeat this step as needed.
  - If loss begins at this step, the most likely causes are:
        
    - Duplex or speed settings mismatch on the link between the AP and the switch, or between the switch and the wired client
    - Bad cable between the AP and switch, or between the switch and the wired client
3. Connect a client directly to the router or firewall on the same VLAN as the wireless client, then ping the router or firewall from the wireless client.
    
  - If loss begins at this step, the most likely causes are:
        
    - Duplex or speed settings mismatch on the link between the switch and the router or firewall, or between the router or firewall and the wired client.
    - Bad cable between the AP and switch, or between the router or firewall and the wired client.
4. If loss begins at this step, the most likely causes are:
        

If testing with Cisco Meraki devices, you can also ping the first MS switch or MX WAN appliance in the path.

Find the IP address of the MS switch by navigating to **switch.meraki.com**, and the MX WAN appliance at **wired.meraki.com**.

The following diagram illustrates the recommended testing sequence for isolating packet loss in a wireless or switched network

 

## Troubleshooting common causes of packet loss

The following sections describe common causes of packet loss and how to resolve them.

### Duplex mismatch

A duplex mismatch occurs when the two ends of a link use different speed or duplex settings, such as 100 Mbps half-duplex and 1000 Mbps full-duplex. When a mismatch occurs, the link drops some or all traffic.

#### Troubleshooting steps

1. Identify the speed and duplex settings on both ends of the affected link.
2. Confirm that both sides of the link have identical settings.
3. Set both ends of the connection to "Auto" for both speed and duplex, if possible.

If you must manually set a speed or duplex setting on one end, configure the same values on the other end.

### Link congestion

Link congestion occurs when more traffic attempts to pass over a network link than the link can support, for example, 60 Mbps of traffic passing over a 20 Mbps link. This condition can result in packet loss.

The following screenshot shows traffic exceeding the available bandwidth, resulting in dropped traffic

#### Troubleshooting steps

Use one or more of the following methods:

- Increase the capacity of the overwhelmed link to accommodate all traffic.
- Apply traffic-shaping rules on MX WAN appliances or MR access points to limit traffic volume, focusing on reducing undesirable traffic.
- Apply traffic-shaping rules on MX WAN appliances or MR access points to prioritize more important traffic.

### Firewall blocking certain traffic

An upstream firewall may filter certain types of traffic even when packet loss is not occurring for all traffic. Firewall filtering can prevent specific websites or services from working correctly.

#### Troubleshooting steps

1. Identify whether a firewall exists between the two devices or locations experiencing these symptoms.
2. Confirm that the firewall is not blocking the traffic experiencing the problem.

The following screenshot shows a firewall allowing some traffic while blocking other traffic

### Bad cable or loose connection

A poorly terminated, incorrectly terminated, or physically damaged cable produces an incomplete or inaccurate electrical signal between devices.

