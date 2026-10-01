---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-5
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [578, 736]
sha256: 75fd8ee6ce627b7d8fd34c0b8ff405386a595404181b205410df11f832100112
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

```huawei
system-view
[USG] security-policy
# --- 1. DNS sortant (indispensable, souvent oublié) ---
[USG-policy-security] rule name LAN-dns-out
[USG-policy-security-rule-LAN-dns-out] source-zone trust
[USG-policy-security-rule-LAN-dns-out] destination-zone untrust
[USG-policy-security-rule-LAN-dns-out] source-address 192.168.10.0 24
[USG-policy-security-rule-LAN-dns-out] service dns
[USG-policy-security-rule-LAN-dns-out] action permit
[USG-policy-security-rule-LAN-dns-out] quit
# --- 2. Web + UTM (antivirus/IPS/URL, voir bloc G) ---
[USG-policy-security] rule name LAN-web-out
[USG-policy-security-rule-LAN-web-out] source-zone trust
[USG-policy-security-rule-LAN-web-out] destination-zone untrust
[USG-policy-security-rule-LAN-web-out] source-address 192.168.10.0 24
[USG-policy-security-rule-LAN-web-out] service http https
[USG-policy-security-rule-LAN-web-out] profile av-profile-lan          # antivirus
[USG-policy-security-rule-LAN-web-out] profile ips-profile-lan         # IPS
[USG-policy-security-rule-LAN-web-out] profile url-profile-lan        # filtrage URL
[USG-policy-security-rule-LAN-web-out] action permit
[USG-policy-security-rule-LAN-web-out] quit
# --- 3. Le reste, au cas par cas (exemple : NTP) ---
[USG-policy-security] rule name LAN-ntp-out
[USG-policy-security-rule-LAN-ntp-out] source-zone trust
[USG-policy-security-rule-LAN-ntp-out] source-zone trust
[USG-policy-security-rule-LAN-ntp-out] destination-zone untrust
[USG-policy-security-rule-LAN-ntp-out] service ntp
[USG-policy-security-rule-LAN-ntp-out] action permit
[USG-policy-security-rule-LAN-ntp-out] quit
[USG-policy-security] quit
save
```

💡 **Ordre recommandé** dans la politique : 1) règles de management (local), 2) deny spécifiques (P2P, geo-block), 3) règles métier par utilisateur/groupe, 4) règles générales (web, DNS), 5) default deny. Revois l'ordre avec `display security-policy rule all` après chaque ajout.

## 36. Compteurs de règles : savoir ce qui sert vraiment

```huawei
[USG] display security-policy rule all          # colonne "hit" / compteurs
[USG] display security-policy statistics rule LAN-web-out   # détail par règle (selon version)
# Remettre les compteurs à zéro pour une mesure propre :
[USG] reset security-policy statistics
```

- Une règle à **0 hit depuis 90 jours** = à désactiver puis supprimer (section 128).
- Une règle `deny` avec **beaucoup de hits** = soit une attaque, soit un flux légitime mal configuré → **enquêter** avant de se réjouir.
- 📋 Relever les compteurs **avant** toute modification : c'est ta preuve que « ça marchait avant ».

## 37. Politique intra-zone : filtrer à l'intérieur d'une même zone

Par défaut, le trafic entre deux interfaces de la **même zone** passe sans politique interzone (comportement à vérifier selon version). Pour l'isoler (ex. : deux départements sur deux ports en trust), la méthode propre :

```huawei
system-view
# Méthode recommandée : deux zones distinctes + politique explicite entre elles.
# (Le filtrage intra-zone direct dépend de la version — syntaxe à vérifier avec ?)
[USG] firewall zone name COMPTA
[USG-zone-compta] set priority 84
[USG-zone-compta] add interface GigabitEthernet 1/0/4
[USG-zone-compta] quit
[USG] security-policy
[USG-policy-security] rule name trust-compta-limite
[USG-policy-security-rule-trust-compta-limite] source-zone trust
[USG-policy-security-rule-trust-compta-limite] destination-zone COMPTA
[USG-policy-security-rule-trust-compta-limite] service ping
[USG-policy-security-rule-trust-compta-limite] action permit
[USG-policy-security-rule-trust-compta-limite] quit
[USG-policy-security] quit
save
```

💡 Si tu as besoin de filtrer, c'est que ce ne sont pas le même niveau de confiance → **deux zones**, pas du bricolage.

## 38. Filtrage géographique dans les politiques

```huawei
[USG] ip address-set pays-bloques type object
# Renseigner avec des plages par pays (via un feed ou manuellement)
[USG] security-policy
[USG-policy-security] rule name geo-block-dmz
[USG-policy-security-rule-geo-block-dmz] source-zone untrust
[USG-policy-security-rule-geo-block-dmz] destination-zone dmz
[USG-policy-security-rule-geo-block-dmz] source-address address-set pays-bloques
[USG-policy-security-rule-geo-block-dmz] action deny
[USG-policy-security-rule-geo-block-dmz] quit
[USG-policy-security] move rule geo-block-dmz before untrust-web-dmz
[USG-policy-security] quit
save
```

⚠️ Les IP des **CDN/SaaS** sont mondiales : un geo-block trop large casse les mises à jour et les services cloud. **Toujours** tester après activation (section 136).

## 39. ICMP : ping et traceroute à travers l'USG

```huawei
[USG] security-policy
[USG-policy-security] rule name ping-lan-wan
[USG-policy-security-rule-ping-lan-wan] source-zone trust
[USG-policy-security-rule-ping-lan-wan] destination-zone untrust
[USG-policy-security-rule-ping-lan-wan] service ping tracert
[USG-policy-security-rule-ping-lan-wan] action permit
[USG-policy-security-rule-ping-lan-wan] quit
# Autoriser le firewall à répondre au ping depuis le LAN (diagnostic) :
[USG-policy-security] rule name ping-vers-local
[USG-policy-security-rule-ping-vers-local] source-zone trust
[USG-policy-security-rule-ping-vers-local] destination-zone local
[USG-policy-security-rule-ping-vers-local] service ping
[USG-policy-security-rule-ping-vers-local] action permit
[USG-policy-security-rule-ping-vers-local] quit
[USG-policy-security] quit
save
```

⚠️ **Jamais** de ping untrust→local en prod (cartographie du réseau par les attaquants). Le ping sortant, oui ; le ping entrant vers le firewall, non.

## 40. Templates de politiques : industrialiser les multi-sites

Quand tu gères 5 sites avec le même socle :
1. Construis une **politique de référence** sur le site pilote (sections 35-39).
2. Exporte-la (`display current-configuration | include security-policy`).
3. Paramètre les variables (IP LAN, IP WAN, noms) dans un **script**.
4. Rejoue sur chaque site en adaptant les variables.

📋 Le template vit dans ton **gestionnaire de versions** (Git), pas dans un coin de disque. Chaque déploiement = un commit. C'est comme ça qu'on évite les « mais sur le site B on avait mis quoi déjà ? ».

## 41. Documenter la politique : l'audit sans douleur

Pour chaque règle, renseigne la **description** :
```huawei
[USG] security-policy
[USG-policy-security] rule name LAN-web-out
[USG-policy-security-rule-LAN-web-out] description "Acces web sortant bureautique - demande DSI-2026-042 - revue 2026-12"
[USG-policy-security-rule-LAN-web-out] quit
[USG-policy-security] quit
save
```

📋 Règle d'or : **pas de règle sans description** (qui l'a demandée, pourquoi, quand la revoir). Un audit avec des descriptions = 1 jour ; sans = 1 semaine.

## 42. Politique et QoS : marquer le trafic (DSCP)

```huawei
# Marquer la VoIP en EF pour que les équipements en aval la priorisent :
[USG] traffic-policy
[USG-policy-traffic] rule name marquer-voip
[USG-policy-traffic-rule-marquer-voip] source-zone trust
[USG-policy-traffic-rule-marquer-voip] destination-zone untrust
[USG-policy-traffic-rule-marquer-voip] service sip rtp
[USG-policy-traffic-rule-marquer-voip] action dscp ef
[USG-policy-traffic-rule-marquer-voip] quit
[USG-policy-traffic] quit
save
```

💡 Le marquage ne sert que si les équipements **en aval** (routeur FAI, switchs) le respectent — coordonne avec ton opérateur (souvent, le FAI réécrit ou ignore le DSCP sur l'accès Internet).

## 43. Règles de management : SNMP, syslog, NTP vers/depuis local

