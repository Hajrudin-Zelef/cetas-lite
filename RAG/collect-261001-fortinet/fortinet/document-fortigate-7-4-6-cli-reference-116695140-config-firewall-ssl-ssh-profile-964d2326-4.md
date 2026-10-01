---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326-4
title: "document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326.md
source_anchor: ""
source_lines: [521, 680]
sha256: 04f0372a9e14b3b3b88797ad037f80a4ffd3b383e7fc7908e412da7ade946837
---

# document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326

|---|---|
| enable | Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. If mismatched, use the CN in the server certificate to do URL filtering. | 
| strict | Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. If mismatched, close the connection. | 
| disable | Do not check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. | 
| Option | Description | 
|---|---|
| disable | Disable. | 
| certificate-inspection | Inspect SSL handshake only. | 
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
config imaps
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
config pop3s
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
