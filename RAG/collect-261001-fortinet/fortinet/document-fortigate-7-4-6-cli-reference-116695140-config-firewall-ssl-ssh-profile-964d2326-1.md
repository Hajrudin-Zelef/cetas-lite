---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326-1
title: "document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326.md
source_anchor: ""
source_lines: [1, 170]
sha256: f4d531dc5895dadc8fbae83cfa13f3ad04ea6b74df33d404a668693fcd40f0cc
---

# document-fortigate-7-4-6-cli-reference-116695140-config-firewall-ssl-ssh-profile-964d2326

config firewall ssl-ssh-profile
config firewall ssl-ssh-profile
Configure SSL/SSH protocol options.
config firewall ssl-ssh-profile
    Description: Configure SSL/SSH protocol options.
    edit <name>
        set allowlist [enable|disable]
        set block-blocklisted-certificates [disable|enable]
        set caname {string}
        set comment {var-string}
        config dot
            Description: Configure DNS over TLS options.
            set cert-validation-failure [allow|block|...]
            set cert-validation-timeout [allow|block|...]
            set client-certificate [bypass|inspect|...]
            set expired-server-cert [allow|block|...]
            set proxy-after-tcp-handshake [enable|disable]
            set quic [inspect|bypass|...]
            set revoked-server-cert [allow|block|...]
            set sni-server-cert-check [enable|strict|...]
            set status [disable|deep-inspection]
            set unsupported-ssl-cipher [allow|block]
            set unsupported-ssl-negotiation [allow|block]
            set unsupported-ssl-version [allow|block]
            set untrusted-server-cert [allow|block|...]
        end
        config ech-outer-sni
            Description: ClientHelloOuter SNIs to be blocked.
            edit <name>
                set sni {string}
            next
        end
        config ftps
            Description: Configure FTPS options.
            set cert-validation-failure [allow|block|...]
            set cert-validation-timeout [allow|block|...]
            set client-certificate [bypass|inspect|...]
            set expired-server-cert [allow|block|...]
            set min-allowed-ssl-version [ssl-3.0|tls-1.0|...]
            set ports {integer}
            set revoked-server-cert [allow|block|...]
            set sni-server-cert-check [enable|strict|...]
            set status [disable|deep-inspection]
            set unsupported-ssl-cipher [allow|block]
            set unsupported-ssl-negotiation [allow|block]
            set unsupported-ssl-version [allow|block]
            set untrusted-server-cert [allow|block|...]
        end
        config https
            Description: Configure HTTPS options.
            set cert-probe-failure [allow|block]
            set cert-validation-failure [allow|block|...]
            set cert-validation-timeout [allow|block|...]
            set client-certificate [bypass|inspect|...]
            set encrypted-client-hello [allow|block]
            set expired-server-cert [allow|block|...]
            set min-allowed-ssl-version [ssl-3.0|tls-1.0|...]
            set ports {integer}
            set proxy-after-tcp-handshake [enable|disable]
            set quic [inspect|bypass|...]
            set revoked-server-cert [allow|block|...]
            set sni-server-cert-check [enable|strict|...]
            set status [disable|certificate-inspection|...]
            set unsupported-ssl-cipher [allow|block]
            set unsupported-ssl-negotiation [allow|block]
            set unsupported-ssl-version [allow|block]
            set untrusted-server-cert [allow|block|...]
        end
        config imaps
            Description: Configure IMAPS options.
            set cert-validation-failure [allow|block|...]
            set cert-validation-timeout [allow|block|...]
            set client-certificate [bypass|inspect|...]
            set expired-server-cert [allow|block|...]
            set ports {integer}
            set proxy-after-tcp-handshake [enable|disable]
            set revoked-server-cert [allow|block|...]
            set sni-server-cert-check [enable|strict|...]
            set status [disable|deep-inspection]
            set unsupported-ssl-cipher [allow|block]
            set unsupported-ssl-negotiation [allow|block]
            set unsupported-ssl-version [allow|block]
            set untrusted-server-cert [allow|block|...]
        end
        set mapi-over-https [enable|disable]
        config pop3s
            Description: Configure POP3S options.
            set cert-validation-failure [allow|block|...]
            set cert-validation-timeout [allow|block|...]
            set client-certificate [bypass|inspect|...]
            set expired-server-cert [allow|block|...]
            set ports {integer}
            set proxy-after-tcp-handshake [enable|disable]
            set revoked-server-cert [allow|block|...]
            set sni-server-cert-check [enable|strict|...]
            set status [disable|deep-inspection]
            set unsupported-ssl-cipher [allow|block]
            set unsupported-ssl-negotiation [allow|block]
            set unsupported-ssl-version [allow|block]
            set untrusted-server-cert [allow|block|...]
        end
        set rpc-over-https [enable|disable]
        set server-cert <name1>, <name2>, ...
        set server-cert-mode [re-sign|replace]
        config smtps
            Description: Configure SMTPS options.
            set cert-validation-failure [allow|block|...]
            set cert-validation-timeout [allow|block|...]
            set client-certificate [bypass|inspect|...]
            set expired-server-cert [allow|block|...]
            set ports {integer}
            set proxy-after-tcp-handshake [enable|disable]
            set revoked-server-cert [allow|block|...]
            set sni-server-cert-check [enable|strict|...]
            set status [disable|deep-inspection]
            set unsupported-ssl-cipher [allow|block]
            set unsupported-ssl-negotiation [allow|block]
            set unsupported-ssl-version [allow|block]
            set untrusted-server-cert [allow|block|...]
        end
        config ssh
            Description: Configure SSH options.
            set inspect-all [disable|deep-inspection]
            set ports {integer}
            set proxy-after-tcp-handshake [enable|disable]
            set ssh-algorithm [compatible|high-encryption]
            set ssh-tun-policy-check [disable|enable]
            set status [disable|deep-inspection]
            set unsupported-version [bypass|block]
        end
        config ssl
            Description: Configure SSL options.
            set cert-probe-failure [allow|block]
            set cert-validation-failure [allow|block|...]
            set cert-validation-timeout [allow|block|...]
            set client-certificate [bypass|inspect|...]
            set encrypted-client-hello [allow|block]
            set expired-server-cert [allow|block|...]
            set inspect-all [disable|certificate-inspection|...]
            set min-allowed-ssl-version [ssl-3.0|tls-1.0|...]
            set revoked-server-cert [allow|block|...]
            set sni-server-cert-check [enable|strict|...]
            set unsupported-ssl-cipher [allow|block]
            set unsupported-ssl-negotiation [allow|block]
            set unsupported-ssl-version [allow|block]
            set untrusted-server-cert [allow|block|...]
        end
        set ssl-anomaly-log [disable|enable]
        config ssl-exempt
            Description: Servers to exempt from SSL inspection.
            edit <id>
                set address {string}
                set address6 {string}
                set fortiguard-category {integer}
                set regex {string}
                set type [fortiguard-category|address|...]
                set wildcard-fqdn {string}
            next
        end
        set ssl-exemption-ip-rating [enable|disable]
        set ssl-exemption-log [disable|enable]
        set ssl-handshake-log [disable|enable]
        set ssl-negotiation-log [disable|enable]
        config ssl-server
            Description: SSL server settings used for client certificate request.
            edit <id>
                set ftps-client-certificate [bypass|inspect|...]
                set https-client-certificate [bypass|inspect|...]
                set imaps-client-certificate [bypass|inspect|...]
                set ip {ipv4-address-any}
