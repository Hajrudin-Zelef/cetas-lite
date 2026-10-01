---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-8
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [1320, 1490]
sha256: 40aaebb9432e4fa970cdbe5e475ee7dcdc81019e6cfec4245c721057765f9593
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

Principe : par défaut, **tout est refusé entre zones** sauf ce qu'on autorise
explicitement via des `security-policy`. C'est un firewall stateful : le retour d'une
connexion autorisée est automatiquement accepté.

## 65. Assigner les interfaces aux zones

```
[AGENCE-DAKAR-AR720]firewall zone trust
[AGENCE-DAKAR-AR720-zone-trust]add interface Vlanif 10
[AGENCE-DAKAR-AR720-zone-trust]add interface Vlanif 30
[AGENCE-DAKAR-AR720-zone-trust]quit
[AGENCE-DAKAR-AR720]firewall zone untrust
[AGENCE-DAKAR-AR720-zone-untrust]add interface Dialer 1
[AGENCE-DAKAR-AR720-zone-untrust]add interface GigabitEthernet 0/0/1
[AGENCE-DAKAR-AR720-zone-untrust]quit
[AGENCE-DAKAR-AR720]firewall zone dmz
[AGENCE-DAKAR-AR720-zone-dmz]add interface Vlanif 20
[AGENCE-DAKAR-AR720-zone-dmz]quit
```

(Ici les invités sont mis en DMZ pour illustrer : Internet seul, pas d'accès au trust.)

## 66. Politique trust → untrust : laisser sortir le LAN

```
[AGENCE-DAKAR-AR720]security-policy
[AGENCE-DAKAR-AR720-policy-security]rule name LAN-vers-Internet
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]source-zone trust
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]destination-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]source-address 192.168.10.0 24
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]quit
[AGENCE-DAKAR-AR720-policy-security]quit
```

## 67. Politique untrust → trust : n'ouvrir que le nécessaire

Par défaut : rien n'est autorisé d'Internet vers le LAN. On n'ouvre que les ports
explicitement exposés via NAT server :

```
[AGENCE-DAKAR-AR720]security-policy
[AGENCE-DAKAR-AR720-policy-security]rule name WAN-vers-ServeurVideo
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-ServeurVideo]source-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-ServeurVideo]destination-zone trust
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-ServeurVideo]destination-address 192.168.30.10 32
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-ServeurVideo]service protocol tcp destination-port 8080
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-ServeurVideo]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-ServeurVideo]quit
[AGENCE-DAKAR-AR720-policy-security]quit
```

Chaque `nat server` = une règle miroir ici. **Audit trimestriel** : lister les NAT server
et vérifier qu'il y a une règle pour chacun, et inversement.

## 68. Politique vers la zone local : protéger le routeur lui-même

La zone `local` = le routeur. Autoriser : SSH d'administration depuis le LAN, SNMP depuis
le superviseur, NTP. Refuser : tout le reste depuis untrust.

```
[AGENCE-DAKAR-AR720]security-policy
[AGENCE-DAKAR-AR720-policy-security]rule name Admin-SSH-LAN
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]source-zone trust
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]destination-zone local
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]service protocol tcp destination-port 22
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]quit
[AGENCE-DAKAR-AR720-policy-security]rule name IKE-vers-local
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]source-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]destination-zone local
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]service protocol udp destination-port 500
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]service protocol udp destination-port 4500
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]quit
[AGENCE-DAKAR-AR720-policy-security]quit
```

Sans la règle IKE vers local, les tunnels IPSec ne montent jamais (le routeur jette les
paquets IKE). C'est la cause n°2 des « le tunnel ne monte pas » après la PSK.

## 69. Vérifier les politiques et les sessions

```
display security-policy all
display firewall session table
display firewall statistics
```

`display firewall session table` montre les connexions actives (utile pour voir si un
flux est bien autorisé/NATé). Si une règle ne matche pas, vérifier l'ordre des règles
et les zones source/destination (l'erreur la plus fréquente : zones inversées).

---
---

# Partie H — QoS

## 70. NTP : la base (rappel — sans horloge juste, rien n'est fiable)

```
[AGENCE-DAKAR-AR720]ntp-service unicast-server 10.0.0.5
[AGENCE-DAKAR-AR720]ntp-service unicast-server 192.168.10.5 prefer
```

Vérification : `display ntp-service status` (stratum, offset). En agence sans serveur
local : pointer vers le serveur NTP du siège via le tunnel VPN.

## 71. QoS : les concepts en 2 minutes

- **Classification** : repérer le trafic (ACL, DSCP, port).
- **Marquage** : poser un DSCP/802.1p.
- **CAR (Committed Access Rate)** : limiter un débit (policing).
- **Files d'attente** : prioriser (LLQ pour la voix, CBWFQ pour le reste).
- **Shaping** : lisser en sortie.

Sur un lien WAN asymétrique (cas général), la QoS utile se fait **en sortie** (upload),
là où on contrôle la file. En entrée, on ne peut que policer (jeter).

## 72. CAR : limiter le débit des invités (exemple)

Limiter le VLAN invités à 20 Mbit/s en sortie vers le WAN :

```
[AGENCE-DAKAR-AR720]acl number 3002
[AGENCE-DAKAR-AR720-acl-adv-3002]rule 5 permit ip source 192.168.20.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-adv-3002]quit
[AGENCE-DAKAR-AR720]traffic classifier INVITES
[AGENCE-DAKAR-AR720-classifier-INVITES]if-match acl 3002
[AGENCE-DAKAR-AR720-classifier-INVITES]quit
[AGENCE-DAKAR-AR720]traffic behavior LIMITE-20M
[AGENCE-DAKAR-AR720-behavior-LIMITE-20M]car cir 20000 pir 22000 green pass yellow pass red discard
[AGENCE-DAKAR-AR720-behavior-LIMITE-20M]quit
[AGENCE-DAKAR-AR720]traffic policy QOS-WAN
[AGENCE-DAKAR-AR720-trafficpolicy-QOS-WAN]classifier INVITES behavior LIMITE-20M
[AGENCE-DAKAR-AR720-trafficpolicy-QOS-WAN]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]traffic-policy QOS-WAN outbound
[AGENCE-DAKAR-AR720-Dialer1]quit
```

- `cir 20000` : débit garanti/limite en kbps (20 Mbit/s).
- `pir` : pic toléré. `green/yellow pass, red discard` : on jette au-delà du pic.

## 73. Prioriser la voix (ToIP) : file prioritaire

```
[AGENCE-DAKAR-AR720]acl number 3003
[AGENCE-DAKAR-AR720-acl-adv-3003]rule 5 permit udp destination-port eq 5060
[AGENCE-DAKAR-AR720-acl-adv-3003]rule 10 permit udp destination-port range 10000 20000
[AGENCE-DAKAR-AR720-acl-adv-3003]quit
[AGENCE-DAKAR-AR720]traffic classifier VOIX
[AGENCE-DAKAR-AR720-classifier-VOIX]if-match acl 3003
[AGENCE-DAKAR-AR720-classifier-VOIX]quit
[AGENCE-DAKAR-AR720]traffic behavior PRIORITE-VOIX
[AGENCE-DAKAR-AR720-behavior-PRIORITE-VOIX]queue pq
[AGENCE-DAKAR-AR720-behavior-PRIORITE-VOIX]quit
[AGENCE-DAKAR-AR720]traffic policy QOS-WAN
[AGENCE-DAKAR-AR720-trafficpolicy-QOS-WAN]classifier VOIX behavior PRIORITE-VOIX
[AGENCE-DAKAR-AR720-trafficpolicy-QOS-WAN]quit
```

- `queue pq` : file prioritaire (Priority Queuing) — la voix passe avant tout.
- Adapter les ports RTP (10000–20000 ici, fictifs : prendre la plage du PABX réel).
- **Règle d'or** : la file prioritaire ne doit représenter qu'une petite part du lien
  (10–30 %). Si on y met tout, elle ne priorise plus rien.

## 74. Marquage DSCP en entrée du LAN

Marquer la voix dès l'entrée pour que le marquage suive dans le tunnel VPN :

