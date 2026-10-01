---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326-2
title: "document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326.md
source_anchor: ""
source_lines: [171, 346]
sha256: 25a21f7f62c2e629e841df56a6ccfe8e5790f7fc13467876255948cba53e3cdf
---

# document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326

                set pop3s-client-certificate [bypass|inspect|...]
                set smtps-client-certificate [bypass|inspect|...]
                set ssl-other-client-certificate [bypass|inspect|...]
            next
        end
        set ssl-server-cert-log [disable|enable]
        set supported-alpn [http1-1|http2|...]
        set untrusted-caname {string}
        set use-ssl-server [disable|enable]
    next
end
                                            config firewall ssl-ssh-profile
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| allowlist | Enable/disable exempting servers by FortiGuard allowlist. | option | - | disable | 
|  |  |  |  |  | 
| block-blocklisted-certificates | Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. | option | - | enable | 
|  |  |  |  |  | 
| caname | CA certificate used by SSL Inspection. | string | Maximum length: 35 | Fortinet_CA_SSL | 
| comment | Optional comments. | var-string | Maximum length: 255 |  | 
| mapi-over-https | Enable/disable inspection of MAPI over HTTPS. | option | - | disable | 
|  |  |  |  |  | 
| name | Name. | string | Maximum length: 35 |  | 
| rpc-over-https | Enable/disable inspection of RPC over HTTPS. | option | - | disable | 
|  |  |  |  |  | 
| server-cert <name> | Certificate used by SSL Inspection to replace server certificate. Certificate list. | string | Maximum length: 79 |  | 
| server-cert-mode | Re-sign or replace the server's certificate. | option | - | re-sign | 
|  |  |  |  |  | 
| ssl-anomaly-log | Enable/disable logging of SSL anomalies. | option | - | enable | 
|  |  |  |  |  | 
| ssl-exemption-ip-rating | Enable/disable IP based URL rating. | option | - | enable | 
|  |  |  |  |  | 
| ssl-exemption-log | Enable/disable logging of SSL exemptions. | option | - | disable | 
|  |  |  |  |  | 
| ssl-handshake-log | Enable/disable logging of TLS handshakes. | option | - | disable | 
|  |  |  |  |  | 
| ssl-negotiation-log | Enable/disable logging of SSL negotiation events. | option | - | enable | 
|  |  |  |  |  | 
| ssl-server-cert-log | Enable/disable logging of server certificate information. | option | - | disable | 
|  |  |  |  |  | 
| supported-alpn | Configure ALPN option. | option | - | all | 
|  |  |  |  |  | 
| untrusted-caname | Untrusted CA certificate used by SSL Inspection. | string | Maximum length: 35 | Fortinet_CA_Untrusted | 
| use-ssl-server | Enable/disable the use of SSL server table for SSL offloading. | option | - | disable | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| disable | Disable FortiGuard certificate blocklist. | 
| enable | Enable FortiGuard certificate blocklist. | 
| Option | Description | 
|---|---|
| enable | Enable inspection of MAPI over HTTPS. | 
| disable | Disable inspection of MAPI over HTTPS. | 
| Option | Description | 
|---|---|
| enable | Enable inspection of RPC over HTTPS. | 
| disable | Disable inspection of RPC over HTTPS. | 
| Option | Description | 
|---|---|
| re-sign | Multiple clients connecting to multiple servers. | 
| replace | Protect an SSL server. | 
| Option | Description | 
|---|---|
| disable | Disable logging of SSL anomalies. | 
| enable | Enable logging of SSL anomalies. | 
| Option | Description | 
|---|---|
| enable | Enable IP based URL rating. | 
| disable | Disable IP based URL rating. | 
| Option | Description | 
|---|---|
| disable | Disable logging of SSL exemptions. | 
| enable | Enable logging of SSL exemptions. | 
| Option | Description | 
|---|---|
| disable | Disable logging of TLS handshakes. | 
| enable | Enable logging of TLS handshakes. | 
| Option | Description | 
|---|---|
| disable | Disable logging of SSL negotiation events. | 
| enable | Enable logging of SSL negotiation events. | 
| Option | Description | 
|---|---|
| disable | Disable logging of server certificate information. | 
| enable | Enable logging of server certificate information. | 
| Option | Description | 
|---|---|
| http1-1 | Enable all ALPN including HTTP1.1 except HTTP2 and SPDY. | 
| http2 | Enable all ALPN including HTTP2 except HTTP1.1 and SPDY. | 
| all | Allow all ALPN extensions except SPDY. | 
| none | Do not use ALPN. | 
| Option | Description | 
|---|---|
| disable | Don't use SSL server configuration. | 
| enable | Use SSL server configuration. | 
config dot
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
| proxy-after-tcp-handshake | Proxy traffic after the TCP 3-way handshake has been established (not before). | option | - | disable | 
|  |  |  |  |  | 
| quic | QUIC inspection status. | option | - | inspect | 
|  |  |  |  |  | 
| revoked-server-cert | Action based on server certificate is revoked. | option | - | block | 
|  |  |  |  |  | 
| sni-server-cert-check | Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. | option | - | enable | 
|  |  |  |  |  | 
| status | Configure protocol inspection status. | option | - | disable | 
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
| inspect | Inspect QUIC traffic. | 
| bypass | Bypass QUIC traffic. | 
| block | Block QUIC traffic. | 
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
