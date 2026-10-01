---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326-3
title: "document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326.md
source_anchor: ""
source_lines: [347, 520]
sha256: d2852223fcb658eed2acc48b5c499d5867bec369cf3de9ae80e8bb25e75b165b
---

# document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326

| allow | Bypass the session when the version is not supported. | 
| block | Block the session when the version is not supported. | 
| Option | Description | 
|---|---|
| allow | Allow the server certificate. | 
| block | Block the session. | 
| ignore | Re-sign the server certificate as trusted. | 
config ech-outer-sni
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| name | ClientHelloOuter SNI name. | string | Maximum length: 79 |  | 
| sni | ClientHelloOuter SNI to be blocked. | string | Maximum length: 255 |  | 
config ftps
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| cert-validation-failure | Action based on certificate validation failure. | option | - | block | 
|  |  |  |  |  | 
| cert-validation-timeout | Action based on certificate validation timeout. | option | - | allow | 
|  |  |  |  |  | 
| client-certificate | Action based on received client certificate. | option | - | bypass | 
|  |  |  |  |  | 
| expired-server-cert | Action based on server certificate is expired. | option | - | block | 
|  |  |  |  |  | 
| min-allowed-ssl-version | Minimum SSL version to be allowed. Flow-based inspection does not support SSL version control. | option | - | tls-1.1 | 
|  |  |  |  |  | 
| ports | Ports to use for scanning. | integer | Minimum value: 1 Maximum value: 65535 |  | 
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
config https
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
| min-allowed-ssl-version | Minimum SSL version to be allowed. Flow-based inspection does not support SSL version control. | option | - | tls-1.1 | 
|  |  |  |  |  | 
| ports | Ports to use for scanning. | integer | Minimum value: 1 Maximum value: 65535 |  | 
| proxy-after-tcp-handshake | Proxy traffic after the TCP 3-way handshake has been established (not before). | option | - | disable | 
|  |  |  |  |  | 
| quic | QUIC inspection status. | option | - | inspect | 
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
| ssl-3.0 | SSL 3.0. | 
| tls-1.0 | TLS 1.0. | 
| tls-1.1 | TLS 1.1. | 
| tls-1.2 | TLS 1.2. | 
| tls-1.3 | TLS 1.3. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| inspect | Inspect QUIC traffic. | 
| bypass | Bypass QUIC traffic. | 
| block | Block QUIC traffic. | 
| Option | Description | 
|---|---|
| allow | Allow the server certificate. | 
| block | Block the session. | 
| ignore | Re-sign the server certificate as trusted. | 
| Option | Description | 
