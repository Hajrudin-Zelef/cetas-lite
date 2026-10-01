---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972-2
title: "c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972.md
source_anchor: ""
source_lines: [63, 117]
sha256: 7ae964d71a34299f59eb0819c29fce6a474642a1e329c46223ecf500fda30f02
---

# c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972

An IKEv1 transform set combines an encryption method and an authentication method. During the IPsec security association negotiation with ISAKMP, the peers agree to use a particular transform set to protect a particular data flow. The transform set must be the same for both peers.
A transform set protects the data flows for the ACL specified in the associated crypto map entry. You can create transform sets in the ASA configuration, and then specify a maximum of 11 of them in a crypto map or dynamic crypto map entry.
Table 10-1 lists valid encryption and authentication methods.
| esp-des | esp-md5-hmac | 
| esp-3des (default) | esp-sha-hmac (default) | 
| esp-aes (128-bit encryption) |  | 
| esp-aes-192 |  | 
| esp-aes-256 |  | 
| esp-null |  | 
Tunnel Mode is the usual way to implement IPsec between two ASAs that are connected over an untrusted network, such as the public Internet. Tunnel mode is the default and requires no configuration.
To configure a transform set, perform the following site-to-site tasks in either single or multiple context mode:
Step 1 In global configuration mode enter the crypto ipsec ikev1 transform-set command. The following example configures a transform set with the name FirstSet, esp-3des encryption, and esp-md5-hmac authentication. The syntax is as follows:
crypto ipsec ikev1 transform-set transform-set-name encryption-method authentication-method
Creating an IKEv2 Proposal
For IKEv2, you can configure multiple encryption and authentication types, and multiple integrity algorithms for a single policy. The ASA orders the settings from the most secure to the least secure and negotiates with the peer using that order. This allows you to potentially send a single proposal to convey all the allowed transforms instead of the need to send each allowed combination as with IKEv1.
Table 10-1 lists valid IKEv2 encryption and authentication methods.
To configure an IKEv2 proposal, perform the following tasks in either single or multiple context mode:
Step 1 In global configuration mode, use the crypto ipsec ikev2 ipsec-proposal command to enter ipsec proposal configuration mode where you can specify multiple encryption and integrity types for the proposal. In this example, secure is the name of the proposal:
Step 2 Then enter a protocol and encryption types. ESP is the only supported protocol. For example:
Step 3 Enter an integrity type. For example:
Configuring an ACL
The ASA uses access control lists to control network access. By default, the adaptive security appliance denies all traffic. You need to configure an ACL that permits traffic. For more information, see "Information About Access Control Lists" in the general operations configuration guide.
The ACLs that you configure for this LAN-to-LAN VPN control connections are based on the source and translated destination IP addresses. Configure ACLs that mirror each other on both sides of the connection.
An ACL for VPN traffic uses the translated address.
To configure an ACL, perform the following steps:
Step 1 Enter the access-list extended command. The following example configures an ACL named l2l_list that lets traffic from IP addresses in the 192.168.0.0 network travel to the 150.150.0.0 network. The syntax is access-list listname extended permit ip source-ipaddress source-netmask destination-ipaddress destination-netmask.
Step 2 Configure an ACL for the ASA on the other side of the connection that mirrors the ACL. Subnets that are defined in two different crypto ACLs and are attached to the same crypto map should not overlap. In the following example, the prompt for the peer is hostname2.
Note For more information on configuring an ACL with a vpn-filter, see the Specifying a VLAN for Remote Access or Applying a Unified Access Control Rule to the Group Policy.
Defining a Tunnel Group
A tunnel group is a set of records that contain tunnel connection policies. You configure a tunnel group to identify AAA servers, specify connection parameters, and define a default group policy. The ASA stores tunnel groups internally.
There are two default tunnel groups in the ASA: DefaultRAGroup, which is the default IPsec remote-access tunnel group, and DefaultL2Lgroup, which is the default IPsec LAN-to-LAN tunnel group. You can modify them but not delete them.
The main difference between IKE versions 1 and 2 lies in terms of the authentication method they allow. IKEv1 allows only one type of authentication at both VPN ends (that is, either pre-shared key or certificate). However, IKEv2 allows assymetric authentication methods to be configured (that is, pre-shared key authentication for the originator but certificate authentication for the responder) using separate local and remote authentication CLIs. Therefore, with IKEv2 you have assymmetric authentication where one side authenticates with one credential whereas the other side uses another credential (either pre-shared key or certificate).
You can also create one or more new tunnel groups to suit your environment. The ASA uses these groups to configure default tunnel parameters for remote access and LAN-to-LAN tunnel groups when there is no specific tunnel group identified during tunnel negotiation.
To establish a basic LAN-to-LAN connection, you must set two attributes for a tunnel group:
- Set the connection type to IPsec LAN-to-LAN.
- Configure an authentication method for the IP, in the following example, preshared key for IKEv1 and IKEv2.
Note To use VPNs, including tunnel groups, the ASA must be in single-routed mode. The commands to configure tunnel-group parameters do not appear in any other mode.
Step 1 To set the connection type to IPsec LAN-to-LAN, enter the tunnel-group command. The syntax is tunnel-group name type type, where name is the name you assign to the tunnel group, and type is the type of tunnel. The tunnel types as you enter them in the CLI are:
In the following example the name of the tunnel group is the IP address of the LAN-to-LAN peer, 10.10.4.108.
Note LAN-to-LAN tunnel groups that have names that are not an IP address can be used only if the tunnel authentication method is Digital Certificates and/or the peer is configured to use Aggressive Mode.
Step 2 To set the authentication method to preshared key, enter the ipsec-attributes mode and then enter the ikev1 pre-shared-key command to create the preshared key. You need to use the same preshared key on both ASAs for this LAN-to-LAN connection.
The key is an alphanumeric string of 1-128 characters.
In the following example the IKEv1 preshared key is 44kkaol59636jnfx:
In the next example, the IKEv2 preshared key is configured also as 44kkaol59636jnfx:
Note You must configure ikev2 remote-authentication pre-shared-key or a certificate to complete the authentication.
To verify that the tunnel is up and running, use the show vpn-sessiondb summary , show vpn-sessiondb detail l2l , or show cry ipsec sa command.
Creating a Crypto Map and Applying It To an Interface
Crypto map entries pull together the various elements of IPsec security associations, including the following:
- Which traffic IPsec should protect, which you define in an ACL.
- Where to send IPsec-protected traffic, by identifying the peer.
- What IPsec security applies to this traffic, which a transform set specifies.
- The local address for IPsec traffic, which you identify by applying the crypto map to an interface.
For IPsec to succeed, both peers must have crypto map entries with compatible configurations. For two crypto map entries to be compatible, they must, at a minimum, meet the following criteria:
- The crypto map entries must contain compatible crypto ACLs (for example, mirror image ACLs). If the responding peer uses dynamic crypto maps, the entries in the ASA crypto ACL must be “permitted” by the peer’s crypto ACL.
- The crypto map entries each must identify the other peer (unless the responding peer is using a dynamic crypto map).
