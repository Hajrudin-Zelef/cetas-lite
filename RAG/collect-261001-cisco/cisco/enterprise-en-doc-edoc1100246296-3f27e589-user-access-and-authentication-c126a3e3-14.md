---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-14
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [2144, 2312]
sha256: 8d52e4443da649e94b8fb0cad74f2f6f1b182cb99b242edb220ebd8dd4e62111
---

# Configure the authentication profile p1, bind the 802.1X access profile d1 to the authentication profile, and configure the forcible authentication domain example.com for users using the authentication profile.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] dot1x-access-profile d1
[DeviceA-authen-profile-p1] access-domain example.com force
[DeviceA-authen-profile-p1] quit
# Bind the authentication profile p1 to a downlink interface and enable 802.1X authentication on the interface. The following uses 10GE 1/0/2 as an example. Other interfaces have similar configurations.
#
sysname DeviceA
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example.com force
#
vlan batch 10 20
#
aaa
 authentication-scheme abc    
  authentication-mode radius    
 accounting-scheme scheme2
  accounting-mode radius
  accounting realtime 15
 domain example.com            
  authentication-scheme abc     
  accounting-scheme scheme2
  radius-server rd1      
#  
radius-server template rd1
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 192.168.1.30 1812 weight 80
 radius-server accounting 192.168.1.30 1813 weight 80
#
dot1x-access-profile name d1
 dot1x timer client-timeout 30
#
interface Vlanif10
 ip address 192.168.1.10 255.255.255.0
#
interface Vlanif20
 ip address 192.168.2.10 255.255.255.0
#
interface 10GE1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface 10GE1/0/2
 port link-type access
 port default vlan 20
 authentication-profile p1
#
return 
In Figure 3-122, terminals in a company's office area are connected to the company's intranet through DeviceA. The downlink interface (for example 10GE 1/0/2) of DeviceA is directly connected to terminals in the office area.
To meet the company's high security requirements, 802.1X authentication and local authentication need to be configured to authenticate terminals in the office area. Additionally, authentication points need to be deployed on DeviceA's interfaces (for example, 10GE 1/0/2) that are directly connected to the terminals.
# Configure the service scheme s1. In the service scheme s1, set the maximum number of users who are allowed to access the network using the same user name to 15.
[DeviceA-aaa] service-scheme s1
[DeviceA-aaa-service-s1] access-limit user-name max-num 15
[DeviceA-aaa-service-s1] quit
# Configure the domain example.com, and apply the authentication scheme a1, authorization scheme b1, and service scheme s1 to the domain.
[DeviceA-aaa] domain example.com
[DeviceA-aaa-domain-example.com] authentication-scheme a1
[DeviceA-aaa-domain-example.com] authorization-scheme b1
[DeviceA-aaa-domain-example.com] service-scheme s1
[DeviceA-aaa-domain-example.com] quit
[DeviceA-aaa] quit
[DeviceA] dot1x-access-profile name d1
[DeviceA-dot1x-access-profile-d1] dot1x authentication-method pap
[DeviceA-dot1x-access-profile-d1] dot1x timer client-timeout 30
[DeviceA-dot1x-access-profile-d1] quit
#
sysname DeviceA
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example.com force
#
vlan batch 10 20
#
aaa
 authentication-scheme a1    
  authentication-mode local
 authorization-scheme b1
  authorization-mode local
 service-scheme s1
  access-limit user-name max-num 15
 domain example.com            
  authentication-scheme a1
  authorization-scheme b1
  service-scheme s1
 local-access-user 00e0-fcd4-8828
  password cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.<FvBB,.w;M75IN5Z>'!L8G:n-!!!!!2jp5!!!!!!<!!!!k9&fPO<BSRW}jPT(,ewKyfIL"zVtM1~=>e.!!!!!%+%# 
  service-type dot1x 
#  
dot1x-access-profile name d1
 dot1x authentication-method pap
 dot1x timer client-timeout 30
#
interface Vlanif10
 ip address 192.168.1.10 255.255.255.0
#
interface Vlanif20
 ip address 192.168.2.10 255.255.255.0
#
interface 10GE1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface 10GE1/0/2
 port link-type access
 port default vlan 20
 authentication-profile p1
#
return 
On the enterprise network shown in Figure 3-123, DeviceA functions as the access device, and two RADIUS servers are deployed for 802.1X and RADIUS authentication of users on the enterprise network. Users can access the Internet only after being successfully authenticated. The administrator requires that users can enter the bypass state and access the Internet when the two RADIUS servers are faulty and that users can be re-authenticated and re-authorized by the RADIUS servers after the RADIUS servers recover.
The RADIUS authentication and accounting keys configured on the RADIUS server must be the same as the RADIUS server's shared key configured on the device.
[DeviceA] radius-server dead-interval 7
[DeviceA] radius-server dead-count 1
Run the test-aaa command four times. The RADIUS server then goes down.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 10 20
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] portswitch
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
[DeviceA-10GE1/0/1] quit
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 192.168.1.10 24
[DeviceA-Vlanif10] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] portswitch
[DeviceA-10GE1/0/2] port link-type hybrid
[DeviceA-10GE1/0/2] port hybrid pvid vlan 20
[DeviceA-10GE1/0/2] port hybrid untagged vlan 20
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 20
[DeviceA-Vlanif20] ip address 192.168.2.10 24
[DeviceA-Vlanif20] quit
# Configure a route destined for RADIUS servers. In this example, the next-hop IP address is 192.168.1.1.
[DeviceA] ip route-static 10.7.66.0 255.255.255.0 192.168.1.1
# Create the RADIUS server template controller.
[DeviceA] radius-server template controller
# Configure IP addresses and port numbers of the active and standby RADIUS authentication and accounting servers, set the algorithm for selecting RADIUS servers to master-backup, and set the RADIUS authentication key.
[DeviceA-radius-controller] radius-server authentication 10.7.66.66 1812 weight 80
[DeviceA-radius-controller] radius-server accounting 10.7.66.66 1813 weight 80
[DeviceA-radius-controller] radius-server authentication 10.7.66.67 1812 weight 40
[DeviceA-radius-controller] radius-server accounting 10.7.66.67 1813 weight 40
[DeviceA-radius-controller] radius-server algorithm master-backup
[DeviceA-radius-controller] radius-server shared-key cipher Huawei@123456789
# Configure automatic detection.
[DeviceA-radius-controller] radius-server testuser username test1 password cipher abc@123
# Configure the automatic detection interval and detection packet timeout period for RADIUS servers in down state. (This example uses the default values.)
[DeviceA-radius-controller] radius-server detect-server interval 60
[DeviceA-radius-controller] radius-server detect-server timeout 3
# Configure the number of retransmissions and timeout interval for RADIUS authentication requests. (This example uses the default values.)
[DeviceA-radius-controller] radius-server retransmit 3 timeout 5
[DeviceA-radius-controller] quit
# Configure the conditions for setting the RADIUS server to the down state. (This example uses the default values.)
[DeviceA] radius-server dead-interval 5
[DeviceA] radius-server dead-count 2
[DeviceA] radius-server detect-cycle 2
[DeviceA] radius-server max-unresponsive-interval 300
# Configure the authentication scheme auth and set the authentication mode to RADIUS authentication.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme auth
[DeviceA-aaa-authen-auth] authentication-mode radius
[DeviceA-aaa-authen-auth] quit
# Configure the accounting scheme acc and set the accounting mode to RADIUS accounting.
[DeviceA-aaa] accounting-scheme acc
[DeviceA-aaa-accounting-acc] accounting-mode radius
[DeviceA-aaa-accounting-acc] quit
