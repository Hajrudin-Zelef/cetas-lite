---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34-2
title: "hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34.md
source_anchor: ""
source_lines: [13, 89]
sha256: 91b71ee5db907d9303cbf6a97804e09d8b1005fcbf32f6f6bbe8a78b3eef4bec
---

# hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34

To access the Internet through DHCP:
- Log in to the AR using STelnet or through the console port.
- Run the system-view command to enter the system view.
- Run the interface interface-type interface-number command to enter the uplink interface view.
- Run the ip address dhcp-alloc command to enable the DHCP client function.By default, the DHCP client function is disabled on an interface.
To access the Internet using a static IP address:
- Log in to the AR using STelnet or through the console port.
- Run the system-view command to enter the system view.
- Run the interface interface-type interface-number command to enter the uplink interface view.
- Run the ip address ip-address { mask | mask-length } command to configure the IP address and subnet mask of the cloud AR. The IP address is a public IP address allocated by the carrier.
- Run the ip route-static 0.0.0.0 0.0.0.0 ip-address command to configure a default route. ip-address in this command is the default gateway address provided by the carrier. Ensure that the AR can ping iMaster NCE-Campus successfully after the configuration.
To access the Internet through PPPoE:
- Log in to the AR using STelnet or through the console port.
- Configure a dialer interface.
  - Run the system-view command to enter the system view.
  - (Optional) Configure a dialer ACL. This step is mandatory only when on-demand dial-up is used.
     - Run the dialer-rule command to enter the dialer-rule view.
    - Run the dialer-rule dialer-rule-number { acl { acl-number | name acl-name } | ip { deny | permit } | ipv6 { deny | permit } } command to configure a dialer ACL for a dialer access group.By default, no dialer ACL is configured.
    - Run the quit command to return to the system view.
  - Run the interface dialer number command to create a dialer interface and enter the dialer interface view.
  - Run the dialer user user-name command to enable the RS-DCC function.By default, the RS-DCC function is disabled and the remote user name is not configured.
  - Run the dialer bundle number command to specify a dialer bundle for the dialer interface.
  - (Optional) Run the dialer-group group-number command to configure a dialer access group to which the dialer interface belongs. This step is mandatory only when on-demand dial-up is used. Make sure that the value of group-number in the dialer-group command is the same as that of dialer-rule-number in the dialer-rule command.
- In the dialer interface view, configure the IP address of the dialer interface.
- Configure a PPPoE client as the supplicant.
  - Configure the user name and password for PPP authentication.
    - Configure PAP authentication.
      - Run the ppp pap local-user username password { cipher | simple } password command to configure the user name and password sent from the PPPoE client to the PPPoE server for PAP authentication.The user name and password must be the same as those configured on the PPPoE server. By default, the local device sends a request without the user name and password to the remote device for PAP authentication.
    - Configure CHAP authentication when a user name is configured on the PPPoE server.
    - Configure CHAP authentication when no user name is configured on the PPPoE server.
      - Run the ppp chap user username command to configure the user name for CHAP authentication.
      - Run the ppp chap password { cipher | simple } password command to configure the password for CHAP authentication.
- Enable PPPoE client on an interface.
  - Run the quit command to return to the system view.
  - Run the interface interface-type interface-number command to enter the interface view.
  - Run the pppoe-client dial-bundle-number number [ on-demand ] [ no-hostuniq ] [ ppp-max-payload value ] [ service-name name ] command to specify a dialer bundle for the PPPoE session. 
    - The number of the specified dialer bundle must be the same as that of the dialer bundle configured using the dialer bundle command in 2.e.
    - If on-demand is not specified, the PPPoE dial-up mode is permanently online. If on-demand is specified, the PPPoE dial-up mode is on-demand dial-up. Currently, the on-demand dial-up mode supported by the device is the packet-triggered mode.
    - If no-hostuniq is not specified, a Host-Uniq field is carried in the call initiated by the PPPoE client to associate with a particular host request, making the check strict.
  - Run the quit command to return to the system view.
  - Run the ip route-static 0.0.0.0 0 { nexthop-address | interface-type interface-number } [ preference preference ] command to configure a static route destined for the PPPoE server.
To access the Internet in 3G mode (complying with the WCDMA standard):
- Log in to the AR using STelnet or through the console port.
- Search for and select a PLMN.
  - Run the system-view command to enter the system view.
  - Run the interface cellular interface-number command to enter the 3G cellular interface view.
  - Run the plmn search command to search for a PLMN. After a period of time, the PLMN information will be displayed on the device.
  - Select a PLMN.
    - Run the plmn auto command to configure automatic selection of a PLMN.
    - Run the plmn select manual mcc mnc [ fail-over-auto ] command to configure manual selection of a PLMN.
 By default, the device automatically selects a PLMN.
- Run the mode wcdma { gsm-only | gsm-precedence | wcdma-only | wcdma-precedence } command to configure the WCDMA network connection mode for a 3G modem.
- Configure the Internet dial-up access (single APN scenario).
  - Run the quit command to return to the system view.
  - Create an APN profile.
    - Run the apn profile profile-name command to create an APN profile and enter the APN profile view.By default, no APN profile is created, and the APN is configured dynamically.
    - Run the apn apn-name command to configure an APN.By default, no APN is configured in the APN profile.  
      - To obtain an APN, contact the local network carrier.
      - After an APN is configured, it is permanently recorded in the 3G modem. If the APN changes, reconfigure it.
    - Run the quit command to return to the system view.
  - Enable C-DCC.
    - Run the interface cellular interface-number command to enter the 3G cellular interface view.
    - Run the ip address negotiate command to configure the 3G cellular interface to obtain an IP address dynamically.
    - Run the dialer enable-circular command to enable the C-DCC function.
  - Run the dialer number dial-number [ autodial ] command to configure a dialer number.You can obtain the dialer number from the carrier.
  - Bind the APN profile to the 3G cellular interface.
    - Run the apn-profile profile-name [ track nqa { admin-name test-name } &<1-2> ] command to configure the APN profile bound to the 3G cellular interface.
    - Run the quit command to return to the system view.
  - Run the ip route-static 0.0.0.0 0 cellular interface-number [ preference preference ] command to configure a default route.interface-number indicates the number of a 3G cellular interface.
- Authenticate a PIN.
  - Run the interface cellular interface-number command to enter the 3G cellular interface view.
  - Run the pin verification enable [ auto ] command to enable PIN authentication on a 3G modem.By default, PIN authentication is disabled. To perform this step, you must enter a PIN to enable PIN authentication on the 3G data card.
  - Run the pin verify [ auto ] command to authenticate the PIN.In this step, the PIN is input automatically or manually. When the message PIN has been verified successfully is displayed on the interface after a period of time, the PIN has been authenticated successfully.
To access the Internet in 3G mode (complying with the CDMA2000 standard):
- Log in to the AR using STelnet or through the console port.
- Configure a network connection mode.
  - Run the system-view command to enter the system view.
