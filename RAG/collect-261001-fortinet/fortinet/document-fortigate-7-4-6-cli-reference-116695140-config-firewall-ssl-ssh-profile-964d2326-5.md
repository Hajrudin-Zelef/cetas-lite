---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326-5
title: "document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326.md
source_anchor: ""
source_lines: [681, 849]
sha256: a9d644d60b1fb60b841535c1c8620d17ea5308baa7d8f10b32f80ba2df1078f4
---

# document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326

| disable | Do not check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. | 
| Option | Description | 
|---|---|
| disable | Disable. | 
| deep-inspection | Full SSL inspection. | 
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
config smtps
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| cert-validation-failure | Action based on certificate validation failure. | option | - | block | 
|  |  |  |  |  | 
| cert-validation-timeout | Action based on certificate validation timeout. | option | - | allow | 
|  |  |  |  |  | 
| client-certificate | Action based on received client certificate. | option | - | inspect | 
|  |  |  |  |  | 
| expired-server-cert | Action based on server certificate is expired. | option | - | block | 
|  |  |  |  |  | 
| ports | Ports to use for scanning. | integer | Minimum value: 1 Maximum value: 65535 |  | 
| proxy-after-tcp-handshake | Proxy traffic after the TCP 3-way handshake has been established (not before). | option | - | disable | 
|  |  |  |  |  | 
| revoked-server-cert | Action based on server certificate is revoked. | option | - | block | 
|  |  |  |  |  | 
| sni-server-cert-check | Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. | option | - | enable | 
|  |  |  |  |  | 
| status | Configure protocol inspection status. | option | - | deep-inspection | 
|  |  |  |  |  | 
| unsupported-ssl-cipher | Action based on the SSL cipher used being unsupported. | option | - | allow | 
|  |  |  |  |  | 
| unsupported-ssl-negotiation | Action based on the SSL negotiation used being unsupported. | option | - | allow | 
|  |  |  |  |  | 
| unsupported-ssl-version | Action based on the SSL version used being unsupported. | option | - | block | 
|  |  |  |  |  | 
| untrusted-server-cert | Action based on server certificate is not issued by a trusted CA. | option | - | allow | 
|  |  |  |  |  | 
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
| allow | Allow the server certificate. | 
| block | Block the session. | 
| ignore | Re-sign the server certificate as trusted. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
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
| disable | Disable. | 
| deep-inspection | Full SSL inspection. | 
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
config ssh
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| inspect-all | Level of SSL inspection. | option | - | disable | 
|  |  |  |  |  | 
| ports | Ports to use for scanning. | integer | Minimum value: 1 Maximum value: 65535 |  | 
| proxy-after-tcp-handshake | Proxy traffic after the TCP 3-way handshake has been established (not before). | option | - | disable | 
|  |  |  |  |  | 
| ssh-algorithm | Relative strength of encryption algorithms accepted during negotiation. | option | - | compatible | 
|  |  |  |  |  | 
| ssh-tun-policy-check | Enable/disable SSH tunnel policy check. | option | - | disable | 
|  |  |  |  |  | 
| status | Configure protocol inspection status. | option | - | disable | 
|  |  |  |  |  | 
| unsupported-version | Action based on SSH version being unsupported. | option | - | bypass | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| disable | Disable. | 
| deep-inspection | Full SSL inspection. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| compatible | Allow a broader set of encryption algorithms for best compatibility. | 
| high-encryption | Allow only AES-CTR, AES-GCM ciphers and high encryption algorithms. | 
| Option | Description | 
|---|---|
| disable | Disable SSH tunnel policy check. | 
| enable | Enable SSH tunnel policy check. | 
| Option | Description | 
|---|---|
| disable | Disable. | 
| deep-inspection | Full SSL inspection. | 
| Option | Description | 
|---|---|
| bypass | Bypass the session. | 
| block | Block the session. | 
config ssl
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| cert-probe-failure | Action based on certificate probe failure. | option | - | block | 
|  |  |  |  |  | 
| cert-validation-failure | Action based on certificate validation failure. | option | - | block | 
|  |  |  |  |  | 
| cert-validation-timeout | Action based on certificate validation timeout. | option | - | allow | 
|  |  |  |  |  | 
| client-certificate | Action based on received client certificate. | option | - | bypass | 
|  |  |  |  |  | 
| encrypted-client-hello | Block/allow session based on existence of encrypted-client-hello. | option | - | block | 
|  |  |  |  |  | 
| expired-server-cert | Action based on server certificate is expired. | option | - | block | 
|  |  |  |  |  | 
| inspect-all | Level of SSL inspection. | option | - | disable | 
|  |  |  |  |  | 
| min-allowed-ssl-version | Minimum SSL version to be allowed. Flow-based inspection does not support SSL version control. | option | - | tls-1.1 | 
|  |  |  |  |  | 
| revoked-server-cert | Action based on server certificate is revoked. | option | - | block | 
|  |  |  |  |  | 
| sni-server-cert-check | Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. | option | - | enable | 
|  |  |  |  |  | 
| unsupported-ssl-cipher | Action based on the SSL cipher used being unsupported. | option | - | allow | 
|  |  |  |  |  | 
| unsupported-ssl-negotiation | Action based on the SSL negotiation used being unsupported. | option | - | allow | 
|  |  |  |  |  | 
| unsupported-ssl-version | Action based on the SSL version used being unsupported. | option | - | block | 
|  |  |  |  |  | 
