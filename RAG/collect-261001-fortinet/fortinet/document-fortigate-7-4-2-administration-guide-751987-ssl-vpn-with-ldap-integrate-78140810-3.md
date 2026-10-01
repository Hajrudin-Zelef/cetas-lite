---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-2-administration-guide-751987-ssl-vpn-with-ldap-integrate-78140810-3
title: "document-fortigate-7-4-2-administration-guide-751987-ssl-vpn-with-ldap-integrate-78140810"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-2-administration-guide-751987-ssl-vpn-with-ldap-integrate-78140810.md
source_anchor: ""
source_lines: [160, 186]
sha256: a34ab04ee6fda9e4dcbf1e889e82c84713a57aa0230a83f57ec8d195499f481b
---

# document-fortigate-7-4-2-administration-guide-751987-ssl-vpn-with-ldap-integrate-78140810

- In a web browser, log into the portal http://172.20.120.123:10443. A message requests a certificate for authentication.
- Select the user certificate.You can connect to the SSL VPN web portal.
To check the SSL VPN connection using the GUI:
- Go to Dashboard > Network and expand the SSL-VPN widget to verify the user's connection.
- Go to Log & Report > VPN Events to view the details of the SSL VPN connection event log.
- Go to Log & Report > Forward Traffic to view the details of the SSL VPN traffic.
To check the SSL VPN connection using the CLI:
Below is a sample output of diagnose debug application fnbamd -1 while the user connects. This is a shortened output sample of a few locations to show the important parts. This sample shows lookups to find the group memberships (three groups total) of the user and that the correct group being found results in a match.
[1148] fnbamd_ldap_recv-Response len: 16, svr: 172.18.60.206
[829] fnbamd_ldap_parse_response-Got one MESSAGE. ID:4, type:search-result
[864] fnbamd_ldap_parse_response-ret=0
[1386] __fnbamd_ldap_primary_grp_next-Auth accepted
[910] __ldap_rxtx-Change state to 'Done'
[843] __ldap_rxtx-state 23(Done)
[925] fnbamd_ldap_send-sending 7 bytes to 172.18.60.206
[937] fnbamd_ldap_send-Request is sent. ID 5
[753] __ldap_stop-svr 'ldap-AD'
[53] ldap_dn_list_del_all-Del CN=test3,OU=Testing,DC=Fortinet-FSSO,DC=COM
[399] ldap_copy_grp_list-copied CN=group3,OU=Testing,DC=Fortinet-FSSO,DC=COM
[399] ldap_copy_grp_list-copied CN=Domain Users,CN=Users,DC=Fortinet-FSSO,DC=COM
[2088] fnbamd_auth_cert_check-Matching group 'sslvpn-group'
[2007] __match_ldap_group-Matching server 'ldap-AD' - 'ldap-AD'
[2015] __match_ldap_group-Matching group 'CN=group3,OU=Testing,DC=Fortinet-FSSO,DC=COM' - 'CN=group3,OU=Testing,DC=Fortinet-FSSO,DC=COM'
[2091] fnbamd_auth_cert_check-Group 'sslvpn-group' matched
[2120] fnbamd_auth_cert_result-Result for ldap svr[0] 'ldap-AD' is SUCCESS
[2126] fnbamd_auth_cert_result-matched user 'test3', matched group 'sslvpn-group'
You can also use diagnose firewall auth list to validate that a firewall user entry exists for the SSL VPN user and is part of the right groups.
