---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326-6
title: "document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326.md
source_anchor: ""
source_lines: [850, 953]
sha256: 636898712dce2efacf8b72efe25fbbe65effd8930cd6e39f9ea849abc49ddbbc
---

# document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326

| untrusted-server-cert | Action based on server certificate is not issued by a trusted CA. | option | - | allow | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| allow | Bypass the session when unable to retrieve server's certificate for inspection. | 
| block | Block the session when unable to retrieve server's certificate for inspection. | 
| Option | Description | 
|---|---|
| allow | Allow the server certificate. | 
| block | Block the session. | 
| ignore | Re-sign the server certificate as trusted. | 
| Option | Description | 
|---|---|
| bypass | Bypass the session. | 
| inspect | Inspect the session. | 
| block | Block the session. | 
| Option | Description | 
|---|---|
| allow | Pass the session when encrypted-client-hello exists. | 
| block | Block the session when encrypted-client-hello exists. | 
| Option | Description | 
|---|---|
| allow | Allow the server certificate. | 
| block | Block the session. | 
| ignore | Re-sign the server certificate as trusted. | 
| Option | Description | 
|---|---|
| disable | Disable. | 
| certificate-inspection | Inspect SSL handshake only. | 
| deep-inspection | Full SSL inspection. | 
| Option | Description | 
|---|---|
| ssl-3.0 | SSL 3.0. | 
| tls-1.0 | TLS 1.0. | 
| tls-1.1 | TLS 1.1. | 
| tls-1.2 | TLS 1.2. | 
| tls-1.3 | TLS 1.3. | 
| Option | Description | 
|---|---|
| allow | Allow the server certificate. | 
| block | Block the session. | 
| ignore | Re-sign the server certificate as trusted. | 
| Option | Description | 
|---|---|
| enable | Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. If mismatched, use the CN in the server certificate to do URL filtering. | 
| strict | Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. If mismatched, close the connection. | 
| disable | Do not check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. | 
| Option | Description | 
|---|---|
| allow | Bypass the session when the cipher is not supported. | 
| block | Block the session when the cipher is not supported. | 
| Option | Description | 
|---|---|
| allow | Bypass the session when the negotiation is not supported. | 
| block | Block the session when the negotiation is not supported. | 
| Option | Description | 
|---|---|
| allow | Bypass the session when the version is not supported. | 
| block | Block the session when the version is not supported. | 
| Option | Description | 
|---|---|
| allow | Allow the server certificate. | 
| block | Block the session. | 
| ignore | Re-sign the server certificate as trusted. | 
config ssl-exempt
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| address | IPv4 address object. | string | Maximum length: 79 |  | 
| address6 | IPv6 address object. | string | Maximum length: 79 |  | 
| fortiguard-category | FortiGuard category ID. | integer | Minimum value: 0 Maximum value: 255 | 0 | 
| id | ID number. | integer | Minimum value: 0 Maximum value: 512 | 0 | 
| regex | Exempt servers by regular expression. | string | Maximum length: 255 |  | 
| type | Type of address object (IPv4 or IPv6) or FortiGuard category. | option | - | fortiguard-category | 
|  |  |  |  |  | 
| wildcard-fqdn | Exempt servers by wildcard FQDN. | string | Maximum length: 79 |  | 
| Option | Description | 
|---|---|
| fortiguard-category | FortiGuard category. | 
| address | Firewall IPv4 address. | 
| address6 | Firewall IPv6 address. | 
| wildcard-fqdn | Fully Qualified Domain Name with wildcard characters. | 
| regex | Regular expression FQDN. | 
config ssl-server
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| ftps-client-certificate | Action based on received client certificate during the FTPS handshake. | option | - | bypass | 
|  |  |  |  |  | 
| https-client-certificate | Action based on received client certificate during the HTTPS handshake. | option | - | bypass | 
|  |  |  |  |  | 
| id | SSL server ID. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| imaps-client-certificate | Action based on received client certificate during the IMAPS handshake. | option | - | bypass | 
|  |  |  |  |  | 
| ip | IPv4 address of the SSL server. | ipv4-address-any | Not Specified | 0.0.0.0 | 
| pop3s-client-certificate | Action based on received client certificate during the POP3S handshake. | option | - | bypass | 
|  |  |  |  |  | 
| smtps-client-certificate | Action based on received client certificate during the SMTPS handshake. | option | - | bypass | 
|  |  |  |  |  | 
| ssl-other-client-certificate | Action based on received client certificate during an SSL protocol handshake. | option | - | bypass | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| bypass | Bypass the session. | 
| inspect | Inspect the session. | 
| block | Block the session. |
