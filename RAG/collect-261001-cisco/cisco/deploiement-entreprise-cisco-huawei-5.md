---
id: collect-261001-cisco/cisco/deploiement-entreprise-cisco-huawei-5
title: "Déploiement entreprise — Cisco & Huawei (référence complète)"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/deploiement_entreprise_cisco_huawei.md
source_anchor: ""
source_lines: [792, 888]
sha256: c3f27f1083fb95abd672687d58d4094fd35d678e264e6298377184507985e505
---

# Déploiement entreprise — Cisco & Huawei (référence complète)

- **SNMPv3** → Zabbix/Centreon : CPU, mémoire, température, statut
  ports, compteurs d'erreurs, voisins OSPF/BGP, sessions.
- **Syslog** → SIEM (Graylog/ELK) : auth, changements de config
  (`logging`/`info-center`), alertes STP, flaps OSPF.
- **NetFlow v9 / sFlow** (Cisco) / **NetStream** (Huawei) → collecteur :
  qui parle à qui, top talkers, détection d'anomalies.
  ```shell
  ! Cisco — Flexible NetFlow
  CISCO(config)# flow record REC1
  CISCO(config-flow-record)# match ipv4 source address
  CISCO(config-flow-record)# match ipv4 destination address
  CISCO(config-flow-record)# collect counter bytes long
  CISCO(config)# flow exporter EXP1
  CISCO(config-flow-exporter)# destination 192.168.99.70
  CISCO(config)# flow monitor MON1
  CISCO(config-flow-monitor)# record REC1
  CISCO(config-flow-monitor)# exporter EXP1
  CISCO(config)# interface TenGigabitEthernet 1/0/1
  CISCO(config-if)# ip flow monitor MON1 input
  ```
  ```shell
  ! Huawei — NetStream
  HUAWEI> system-view
  HUAWEI] ip netstream
  HUAWEI] ip netstream export source LoopBack 0
  HUAWEI] ip netstream export host 192.168.99.70 9996
  HUAWEI] interface 10GE 1/0/1
  HUAWEI-10GE] ip netstream inbound
  HUAWEI-10GE] ip netstream outbound
  ```
- **IP SLA** (Cisco) / **NQA** (Huawei) : sondes de latence/perte vers
  les sites critiques + tracking pour HSRP/VRRP.

---

## 13. Checklist de mise en production

### Avant le jour J
- [ ] Configs relues par un pair (second œil obligatoire).
- [ ] `copy run start` / `save` sur **chaque** équipement.
- [ ] Backups `config.text`/`vrpcfg.zip` exportés hors-site.
- [ ] Plan de rollback écrit (qui fait quoi si ça casse, dans quel ordre).
- [ ] Fenêtre de maintenance communiquée aux utilisateurs.
- [ ] Console/accès OOB testés sur chaque équipement (en cas de perte
      réseau, on administre en console).
- [ ] Étiquetage physique : chaque câble, chaque port uplink.

### Recette (tests à cocher)
- [ ] Chaque VLAN : DHCP OK, passerelle OK, Internet OK.
- [ ] Inter-VLAN : autorisé/refusé **conformément à la politique**
      (tester les deux sens, avec `ping` ET le vrai protocole).
- [ ] HSRP/VRRP : extinction du master → bascule < 10 s, retour
      préemptif OK.
- [ ] OSPF : `show ip ospf neighbor` / `display ospf peer` = Full
      partout ; coupure d'un lien → reconvergence < 5 s.
- [ ] BGP : sessions Established, `show ip bgp` = vos préfixes
      uniquement en annonce.
- [ ] VPN : tunnels UP, trafic chiffré, test de coupure/reconnexion.
- [ ] DHCP snooping : brancher un serveur DHCP pirate → bloqué,
      alerte log.
- [ ] Port-security : 4e MAC → violation `restrict`, log.
- [ ] QoS : saturation du lien + appel VoIP → voix claire
      (test réel, pas théorique).
- [ ] NTP : `show ntp status` / `display ntp-service status` =
      synchronisé sur tous.
- [ ] SNMP : le superviseur voit tous les équipements.
- [ ] Syslog : un `login` test apparaît dans le SIEM.
- [ ] 802.1X (si déployé) : machine non autorisée → VLAN invité ou
      rejet.

### Le jour J (ordre)
1. Core → 2. Distribution → 3. Accès → 4. Edge/WAN → 5. Firewall/DMZ.
2. Valider chaque couche avant de passer à la suivante.
3. Garder l'ancien réseau en parallèle si possible (rollback = rebrancher).

---

## 14. Maintenance et évolutions

- **Firmwares** : 1 version de retard sur la dernière (sauf faille
  critique), jamais de `.0` en production, procédure de rollback
  testée en lab.
- **Config management** : versionner les configs (Git), tout changement
  = ticket + fenêtre + backup avant + vérification après.
- **Revues trimestrielles** : règles firewall/ACL obsolètes, comptes
  inactifs, certificats < 90 jours, capacités (CPU, ports libres).
- **Exercices** : failover HA et restore backup **chronométrés** 2×/an.
  Un plan non testé est un vœu pieux.
- **Documentation vivante** : schéma à jour, plan d'adressage à jour,
  mots de passe dans le coffre (jamais dans la doc).

---

*Fin de la référence. Documents compagnons : [Guide CLI Cisco](cisco_cli_guide.md),
[Guide CLI Huawei USG](huawei_usg_guide.md),
[Troubleshooting Huawei USG](huawei_usg_troubleshooting.md),
références officielles Huawei (tars EDOC) et Cisco (Command Reference IOS).*
