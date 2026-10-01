---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34-3
title: "hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34.md
source_anchor: ""
source_lines: [90, 164]
sha256: 4059f67b23c47211de7c08e5618fe8b42b7ab6076ac7cf1b577a6f637b1c311c
---

# hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34

  - Run the interface cellular interface-number command to enter the 3G cellular interface view.
  - Run the mode cdma { 1xrtt-only | evdo-only | hybrid } command to configure the CDMA2000 network connection mode for a 3G modem.
- Configure the Internet dial-up access.
  - Run the quit command to return to the system view.
  - Enable C-DCC.
    - Run the interface cellular interface-number command to enter the 3G cellular interface view.
    - Run the dialer enable-circular command to enable the C-DCC function.
  - Run the dialer number dial-number [ autodial ] command to configure a dialer number.You can obtain the dialer number from the carrier.
  - Run the quit command to return to the system view.
  - Run the ip route-static 0.0.0.0 0 cellular interface-number [ preference preference ] command to configure a default route.interface-number indicates the number of a 3G cellular interface.
- Authenticate a PIN.
  - Run the interface cellular interface-number command to enter the 3G cellular interface view.
  - Run the pin verification enable [ auto ] command to enable PIN authentication on a 3G modem.By default, PIN authentication is disabled. To perform this step, you must enter a PIN to enable PIN authentication on the 3G data card.
  - Run the pin verify [ auto ] command to authenticate the PIN.In this step, the PIN is input automatically or manually. When the message PIN has been verified successfully is displayed on the interface after a period of time, the PIN has been authenticated successfully.
To access the Internet through LTE:
- Log in to the AR using STelnet or through the console port.
- Search for and select a PLMN.
  - Run the system-view command to enter the system view.
  - Run the interface cellular interface-number command to enter the LTE cellular interface view.
  - Run the plmn search command to search for a PLMN.
  - Select a PLMN.
    - Run the plmn auto command to configure automatic selection of a PLMN.
    - Run the plmn select manual mcc mnc [ fail-over-auto ] command to configure manual selection of a PLMN.
 By default, the device automatically selects a PLMN.
- Run the mode lte { auto | gsm-only | 1xrtt-only | evdo-only | hybrid | lte-only | tdscdma-only | umts-gsm | umts-only | wcdma-gsm | wcdma-only } command to configure the 3G/LTE network connection mode for an LTE modem.By default, the 3G/LTE network connection mode is auto for an LTE modem.
- Configure an APN profile (single SIM card and single APN scenario).
  - Run the quit command to return to the system view.
  - Create an APN profile.
    - Run the apn profile profile-name command to create an APN profile and enter the APN profile view.By default, no APN profile is created, and the APN is configured dynamically.
    - Run the apn apn-name command to configure an APN.By default, no APN is configured in the APN profile.  
      - To obtain an APN, contact the local network carrier.
      - After an APN is configured, it is permanently recorded in the LTE data card. If the APN changes, reconfigure it.
    - Run the quit command to return to the system view.
  - Run the apn-profile profile-name [ track nqa { admin-name test-name } &<1-2> ] command to configure the APN profile bound to an LTE cellular interface.
- Configure the MTU.
  - Run the interface cellular interface-number command to enter the LTE cellular interface view or the LTE channel interface view.
  - Run the mtu mtu command to configure the MTU for the LTE cellular interface or the LTE channel interface.The default MTU of an LTE cellular interface or an LTE channel interface is 1500 bytes.
  - Run the quit command to return to the system view.
- Configure C-DCC for dial-up connection.
  - Configure a dialer ACL.
    - Run the dialer-rule command to enter the dialer-rule view.
    - Run the dialer-rule dialer-rule-number { acl { acl-number | name acl-name } | ip { deny | permit } | ipv6 { deny | permit } } command to configure a dialer ACL for a dialer access group to specify the conditions for initiating DCC calls.
    - Run the quit command to return to the system view.
  - Enable C-DCC.
    - Run the interface cellular interface-number command to enter the LTE cellular interface view or the LTE channel interface view.When the multi-APN function is configured, the LTE channel interface view is displayed. If the multi-APN function is not configured, the LTE cellular interface view is displayed.
    - Run the dialer enable-circular command to enable the C-DCC function.By default, the C-DCC function is disabled on an interface.
    - Run the dialer-group group-number command to configure a dialer access group for the dialer interface.By default, no dialer access group is configured.
  - Obtain an IP address.
    - When a user uses two E392 data cards to connect to the Internet, the LTE link uses the PPP dial-up mode.Run the ip address ppp-negotiate command to configure the interface to obtain an IP address from the remote device through PPP negotiation.
    - In other cases, the WWAN dial-up mode is used.Run the ip address negotiate command to configure an LTE cellular interface or an LTE channel interface to obtain an IP address dynamically.
  - Run the dialer number dial-number [ autodial ] command to configure a dialer number.You can obtain the dialer number from the carrier.
  - Run the quit command to return to the system view.
  - Run the ip route-static 0.0.0.0 0 { nexthop-address | interface-type interface-number } [ preference preference ] command to configure a default route.
- Authenticate a PIN.
  - Run the interface cellular interface-number command to enter the LTE cellular interface view.
  - Run the pin verification enable [ auto ] command to enable PIN authentication on an LTE modem.By default, PIN authentication is disabled. In this step, you must enter a PIN to enable PIN authentication for the LTE modem.
  - Run the pin verify [ auto ] command to authenticate the PIN.In this step, you must enter the PIN. When the message PIN has been verified successfully is displayed on the interface after a period of time, the PIN has been authenticated successfully.
To access the Internet through ADSL:
- Log in to the AR using STelnet or through the console port.
- Deactivate the ADSL interface.
  - Run the system-view command to enter the system view.
  - Run the interface atm interface-number command to enter the ADSL interface view.
  - Run the shutdown command to deactivate the ADSL interface.
- Set uplink parameters for the ADSL interface.
  - Run the adsl standard { adsl2 [ annexm | annexj | annexl ] | adsl2+ [ annexm | annexj ] | auto | gdmt | t1413 } command to configure the transmission standard for the ADSL interface.By default, the transmission standard for an ADSL interface is auto.  The device supports both ADSL-A/M and ADSL-B/J. annexm, annexl, and t1413 are supported only by ADSL-A/M. annexj is supported only by ADSL-B/J.
  - Run the adsl bitswap { off | on } command to enable or disable bit exchange on the ADSL interface.By default, bit exchange is enabled on an ADSL interface.
  - Run the adsl sra { off | on } command to enable or disable seamless rate adaptation on the ADSL interface.By default, seamless rate adaptation is disabled on an ADSL interface.
  - Run the adsl trellis { off | on } command to enable or disable trellis coding on the ADSL interface.By default, trellis coding is enabled on an ADSL interface.
- Run the undo shutdown command to activate the ADSL interface.
To access the Internet through VDSL (in ATM mode):
- Log in to the AR using STelnet or through the console port.
- Configure the VDSL interface to work in ATM mode.
  - Run the system-view command to enter the system view.
  - Run the set workmode slot slot-id vdsl atm command to configure the VDSL interface to work in ATM mode.By default, a VDSL interface works in PTM mode.
- Deactivate the VDSL interface.
