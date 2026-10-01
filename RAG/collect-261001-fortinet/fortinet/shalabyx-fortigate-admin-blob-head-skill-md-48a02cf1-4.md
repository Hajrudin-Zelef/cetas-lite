---
id: collect-261001-fortinet/fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1-4
title: "Use VIP in policy"
domain: fortinet
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["benchmark", "sandbox"]
source: docs/RAG/collect-261001-fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1.md
source_anchor: ""
source_lines: [823, 1088]
sha256: e7ec926797d4e0ddc295eb0d3e5cac827da385c936640fc2cc77feadec80b80a
---

# Use VIP in policy

config system automation-stitch
    edit "IPS-Alert-Stitch"
        set trigger "IPS-Event"
        set action "Email-Alert"
        set status enable
    next
endconfig system automation-trigger
    edit "Failed-Login"
        set event-type event-log
        set logid 44544
    next
end
config system automation-action
    edit "Ban-IP"
        set action-type quarantine-ip
        set quarantine-expiry 1h
    next
end
config system automation-stitch
    edit "Block-Failed-Login"
        set trigger "Failed-Login"
        set action "Ban-IP"
        set status enable
    next
end
Security Fabric → Automation → Create New
- Always backup config before upgrading
- Check Fortinet upgrade path tool: https://docs.fortinet.com/upgrade-tool
- Never skip major versions without following the upgrade path
- Schedule during maintenance window
- Test on non-production unit first if possible
execute backup config tftp fortigate-backup.conf 192.168.1.100
# or via SCP:
execute backup config scp fortigate-backup.conf 192.168.1.100 22 username /backups/# Upload firmware via TFTP
execute restore image tftp FGT_60F-v7.4.x-FW-build.out 192.168.1.100
# Confirm when prompted — unit will rebootget system status
diagnose sys flash list
System → Firmware → Upload Firmware
# Backup to TFTP
execute backup config tftp backup.conf 192.168.1.100
# Backup to USB
execute backup config usb backup.conf
# Backup via SCP
execute backup config scp backup.conf 192.168.1.100 22 user /backups/
# Backup with encryption
execute backup config tftp backup.conf 192.168.1.100 encrypt MyEncryptKeyexecute restore config tftp backup.conf 192.168.1.100
# Unit will reboot after restore
System → Configuration → Backup / Restore
config system automation-trigger
    edit "Daily-Backup"
        set event-type scheduled
        set schedule-type daily
        set schedule-hour 2
        set schedule-minute 0
    next
end
config system automation-action
    edit "Backup-Action"
        set action-type backup-config
        set ftp-server "192.168.1.100"
        set ftp-username "backupuser"
        set ftp-password "backuppass"
        set ftp-directory "/fortigate/"
    next
end
config system automation-stitch
    edit "Auto-Backup"
        set trigger "Daily-Backup"
        set action "Backup-Action"
        set status enable
    next
endconfig system snmp community
    edit 1
        set name "public"
        config hosts
            edit 1
                set ip 192.168.1.50 255.255.255.255
            next
        end
        set events cpu-high mem-low intf-ip vpn-tun-up vpn-tun-down
    next
end
config system snmp user
    edit "snmpv3user"
        set security-level auth-priv
        set auth-proto sha
        set auth-pwd "AuthPass123!"
        set priv-proto aes
        set priv-pwd "PrivPass123!"
        set queries enable
        set trap-status enable
        set trap-lport 162
        set trap-rport 162
    next
endconfig log syslogd setting
    set status enable
    set server "192.168.1.60"
    set port 514
    set facility local7
    set format default
    set mode udp
end
# Send specific log types
config log syslogd filter
    set severity information
    set forward-traffic enable
    set local-traffic enable
    set sniffer-traffic disable
    set ztna-traffic enable
    set anomaly enable
    set voip enable
end
Log & Report → Log Settings → Remote Logging
Apply these settings for CIS/Fortinet benchmark compliance:
# Disable unused admin access
config system interface
    edit "wan1"
        set allowaccess ping        # Remove https/ssh from WAN
    next
end
# Enforce strong admin password policy
config system password-policy
    set status enable
    set min-length 12
    set must-contain upper-case-letter lower-case-letter number non-alphanumeric
    set expire-status enable
    set expire-day 90
end
# Set admin session timeout
config system global
    set admin-timeout 10            # 10 minutes
    set admin-lockout-threshold 5
    set admin-lockout-duration 300
end
# Enable management interface restrictions
config system global
    set admin-sport 8443            # Change default HTTPS port
    set admin-ssh-port 2222         # Change default SSH port
end
# Enable audit logging
config log setting
    set fwpolicy-implicit-log enable
    set local-in-allow enable
    set local-in-deny-unicast enable
    set local-in-deny-broadcast enable
    set local-out enable
end
# Disable unused services
config system global
    set gui-certificates enable
    set endpoint-control-fgt-log-quota 0
end
# NTP hardening
config system ntp
    set status enable
    set type custom
    config ntpserver
        edit 1
            set server "pool.ntp.org"
        next
    end
    set syncinterval 60
end
| Item | Command to Verify | 
|---|---|
| Admin timeout set | get system global \| grep admin-timeout | 
| Password policy enabled | get system password-policy | 
| Unused services disabled | get system global | 
| SSH/HTTPS on WAN disabled | get system interface \| grep allowaccess | 
| Logging enabled on all policies | show firewall policy \| grep logtraffic | 
| SNMP community not "public" | get system snmp community | 
| FortiGuard updates scheduled | get system autoupdate schedule | 
| HA configured | get system ha status | 
| Backup scheduled | get system automation-stitch | 
config log fortianalyzer setting
    set status enable
    set server "192.168.1.200"
    set reliable enable
    set ssl-min-proto-version TLSv1-2
    set enc-algorithm high
    set certificate-verification enable
endconfig system central-management
    set type fortimanager
    set fmg "192.168.1.210"
    set fmg-update-port enable
end# On FortiGate
execute central-mgmt register-device 192.168.1.210 admin password
# Verify connection
diagnose test application oftpd 2
get log fortianalyzer statusconfig system interface
    edit "fortilink"
        set fortilink enable
        set ip 169.254.1.1 255.255.255.0
        set allowaccess ping fabric
    next
end# FortiSwitch will appear automatically after connecting
config switch-controller managed-switch
    edit "S248EPTF19000000"
        set authorized enable
    next
endconfig switch-controller managed-switch
    edit "S248EPTF19000000"
        config ports
            edit "port1"
                set vlan "VLAN10"
                set allowed-vlans "VLAN10" "VLAN20"
                set untagged-vlans "VLAN10"
            next
        end
    next
end
WiFi & Switch Controller → Managed FortiSwitches
config wireless-controller setting
    set country US
endconfig wireless-controller wtp
    edit "FP231F-SERIALNO"
        set admin enable
        set wtp-profile "FAP231F-default"
    next
endconfig wireless-controller vap
    edit "Corporate-WiFi"
        set ssid "Corporate"
        set security wpa2-only-personal
        set passphrase "WifiPass123!"
        set vlanid 10
        set vdom "root"
    next
end
WiFi & Switch Controller → SSIDs / Managed APs
config system dhcp server
    edit 1
        set interface "internal"
        set default-gateway 192.168.1.1
        set netmask 255.255.255.0
        set dns-server1 8.8.8.8
        set dns-server2 1.1.1.1
        config ip-range
            edit 1
                set start-ip 192.168.1.100
                set end-ip 192.168.1.200
            next
        end
        set lease-time 86400
        # Static lease
        config reserved-address
            edit 1
                set ip 192.168.1.10
                set mac 00:11:22:33:44:55
                set description "Printer"
            next
        end
    next
enddiagnose ip dhcp lease list
get system dhcp server
Network → DHCP Servers
config system fortisandbox
    set status enable
    set server "192.168.1.230"
    set enc-algorithm default
end
# Enable sandbox inspection in AV profile
config antivirus profile
    edit "AV-Profile"
        set analytics-max-upload 10
        set analytics-wl-filetype 0
        config http
