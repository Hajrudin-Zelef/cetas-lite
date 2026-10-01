---
id: collect-261001-fortinet/fortinet/document-fortigate-6-2-0-cookbook-632796-ospf-with-ipsec-vpn-for-network-redunda-290a294b-1
title: "OSPF with IPsec VPN for network redundancy"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-6-2-0-cookbook-632796-ospf-with-ipsec-vpn-for-network-redunda-290a294b.md
source_anchor: ""
source_lines: [1, 56]
sha256: fd196efb508d176bca7f958fe11c0f52a9689656ab4ddbe009894253ceafab94
---

# OSPF with IPsec VPN for network redundancy

# OSPF with IPsec VPN for network redundancy

This is a sample configuration of using OSPF with IPsec VPN to set up network redundancy. Route selection is based on OSPF cost calculation. You can configure ECMP or primary/secondary routes by adjusting OSPF path cost.


Because the GUI can only complete part of the configuration, we recommend using the CLI.

###### To configure OSPF with IPsec VPN to achieve network redundancy using the CLI:

1. Configure the WAN interface and static route.Each FortiGate has two WAN interfaces connected to different ISPs. The ISP1 link is for the primary FortiGate and the IPS2 link is for the secondary FortiGate. 
  1. Configure HQ1.
						config system interface edit "port1" set alias to_ISP1 set ip 172.16.200.1 255.255.255.0 next edit "port2" set alias to_ISP2 set ip 172.17.200.1 255.255.255.0 next end config router static edit 1 set gateway 172.16.200.3 set device "port1" next edit 2 set gateway 172.17.200.3 set device "port2" set priority 100 next end
  2. Configure HQ2.
						config system interface edit "port25" set alias to_ISP1 set ip 172.16.202.1 255.255.255.0 next edit "port26" set alias to_ISP2 set ip 172.17.202.1 255.255.255.0 next end config router static edit 1 set gateway 172.16.202.2 set device "port25" next edit 2 set gateway 172.17.202.2 set device "port26" set priority 100 next end
2. Configure HQ1.
						
3. Configure the internal (protected subnet) interface.
  1. Configure HQ1.
						config system interface edit "dmz" set ip 10.1.100.1 255.255.255.0 next end
  2. Configure HQ2.
						config system interface edit "port9" set ip 172.16.101.1 255.255.255.0 next end
4. Configure HQ1.
						
5. Configure IPsec phase1-interface and phase-2 interface. On each FortiGate, configure two IPsec tunnels: a primary and a secondary.
  1. Configure HQ1.
						config vpn ipsec phase1-interface edit "pri_HQ2" set interface "port1" set peertype any set net-device enable set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1 set remote-gw 172.16.202.1 set psksecret sample1 next edit "sec_HQ2" set interface "port2" set peertype any set net-device enable set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1 set remote-gw 172.17.202.1 set psksecret sample2 next end config vpn ipsec phase2-interface edit "pri_HQ2" set phase1name "pri_HQ2" set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305 set auto-negotiate enable next edit "sec_HQ2" set phase1name "sec_HQ2" set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305 set auto-negotiate enable next end
  2. Configure HQ2.
						config vpn ipsec phase1-interface edit "pri_HQ1" set interface "port25" set peertype any set net-device enable set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1 set remote-gw 172.16.200.1 set psksecret sample1 next edit "sec_HQ1" set interface "port26" set peertype any set net-device enable set proposal aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1 set remote-gw 172.17.200.1 set psksecret sample2 next end config vpn ipsec phase2-interface edit "pri_HQ1" set phase1name "pri_HQ1" set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305 set auto-negotiate enable next edit "sec_HQ1" set phase1name "sec_HQ1" set proposal aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305 set auto-negotiate enable next end
6. Configure HQ1.
						
7. Configure an inbound and outbound firewall policy for each IPsec tunnel.
  1. Configure HQ1.
						config firewall policy edit 1 set name "pri_inbound" set srcintf "pri_HQ2" set dstintf "dmz" set srcaddr "172.16.101.0" set dstaddr "10.1.100.0" set action accept set schedule "always" set service "ALL" next edit 2 set name "pri_outbound" set srcintf "dmz" set dstintf "pri_HQ2" set srcaddr "10.1.100.0" set dstaddr "172.16.101.0" set action accept set schedule "always" set service "ALL" next edit 3 set name "sec_inbound" set srcintf "sec_HQ2" set dstintf "dmz" set srcaddr "172.16.101.0" set dstaddr "10.1.100.0" set action accept set schedule "always" set service "ALL" next edit 4 set name "sec_outbound" set srcintf "dmz" set dstintf "sec_HQ2" set srcaddr "10.1.100.0" set dstaddr "172.16.101.0" set action accept set schedule "always" set service "ALL" next end
  2. Configure HQ2.
						config firewall policy edit 1 set name "pri_inbound" set srcintf "pri_HQ1" set dstintf "port9" set srcaddr "10.1.100.0" set dstaddr "172.16.101.0" set action accept set schedule "always" set service "ALL" next edit 2 set name "pri_outbound" set srcintf "port9" set dstintf "pri_HQ1" set srcaddr "10.1.100.0" set dstaddr "172.16.101.0" set action accept set schedule "always" set service "ALL" next edit 3 set name "sec_inbound" set srcintf "sec_HQ1" set dstintf "port9" set srcaddr "10.1.100.0" set dstaddr "172.16.101.0" set action accept set schedule "always" set service "ALL" next edit 4 set name "sec_outbound" set srcintf "port9" set dstintf "sec_HQ1" set srcaddr "172.16.101.0" set dstaddr "10.1.100.0" set action accept set schedule "always" set service "ALL" next end
8. Configure HQ1.
						
9. Assign an IP address to the IPsec tunnel interface.
			
  1. Configure HQ1.
						config system interface edit "pri_HQ2" set ip 10.10.10.1 255.255.255.255 set remote-ip 10.10.10.2 255.255.255.255 next edit "sec_HQ2" set ip 10.10.11.1 255.255.255.255 set remote-ip 10.10.11.2 255.255.255.255 next end
  2. Configure HQ2.
				config system interface edit "pri_HQ1" set ip 10.10.10.2 255.255.255.255 set remote-ip 10.10.10.1 255.255.255.255 next edit "sec_HQ1" set ip 10.10.11.2 255.255.255.255 set remote-ip 10.10.11.1 255.255.255.255 next end
10. Configure HQ1.
						
11. Configure OSPF.
  1. Configure HQ1.
						config router ospf set router-id 1.1.1.1 config area edit 0.0.0.0 next end config ospf-interface edit "pri_HQ2" set interface "pri_HQ2" set cost 10 set network-type point-to-point next edit "sec_HQ2" set interface "sec_HQ2" set cost 20 set network-type point-to-point next end config network edit 1 set prefix 10.10.10.0 255.255.255.0 next edit 2 set prefix 10.10.11.0 255.255.255.0 next edit 3 set prefix 10.1.100.0 255.255.255.0 next end end
  2. Configure HQ2.config router ospf set router-id 2.2.2.2 config area edit 0.0.0.0 next end config ospf-interface edit "pri_HQ1" set interface "pri_HQ1" set cost 10 set network-type point-to-point next edit "sec_HQ1" set interface "sec_HQ1" set cost 20 set network-type point-to-point next end config network edit 1 set prefix 10.10.10.0 255.255.255.0 next edit 2 set prefix 10.10.11.0 255.255.255.0 next edit 3 set prefix 172.16.101.0 255.255.255.0 next end end
12. Configure HQ1.
						

###### To check VPN and OSPF states using diagnose and get commands:

