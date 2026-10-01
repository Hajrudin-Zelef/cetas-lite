---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-9
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [1491, 1689]
sha256: 1f42fbc1ffd1249cb116b4cc247cb5a5cd77677d3831c23236f27793496e8a26
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

```
[AGENCE-DAKAR-AR720]traffic classifier VOIX-IN
[AGENCE-DAKAR-AR720-classifier-VOIX-IN]if-match acl 3003
[AGENCE-DAKAR-AR720-classifier-VOIX-IN]quit
[AGENCE-DAKAR-AR720]traffic behavior MARQUE-EF
[AGENCE-DAKAR-AR720-behavior-MARQUE-EF]remark dscp ef
[AGENCE-DAKAR-AR720-behavior-MARQUE-EF]quit
[AGENCE-DAKAR-AR720]traffic policy MARQUAGE-LAN
[AGENCE-DAKAR-AR720-trafficpolicy-MARQUAGE-LAN]classifier VOIX-IN behavior MARQUE-EF
[AGENCE-DAKAR-AR720-trafficpolicy-MARQUAGE-LAN]quit
[AGENCE-DAKAR-AR720]interface Vlanif 30
[AGENCE-DAKAR-AR720-Vlanif30]traffic-policy MARQUAGE-LAN inbound
[AGENCE-DAKAR-AR720-Vlanif30]quit
```

## 75. Vérifier la QoS

```
display traffic policy applied-record
display traffic-policy statistics interface Dialer 1 outbound
```

Les statistiques montrent les paquets par classe et les paquets jetés par le CAR.
Si la voix saccade malgré la file prioritaire : vérifier que le classifier matche bien
(`display` des compteurs), et que le lien montant n'est pas saturé à 100 % en permanence
(la QoS ne crée pas de débit).

---
---

# Partie I — Supervision

## 76. SNMP : les bases (v2c puis v3)

Activer SNMP et déclarer la communauté (v2c, lecture seule, restreinte au superviseur) :

```
[AGENCE-DAKAR-AR720]snmp-agent
[AGENCE-DAKAR-AR720]snmp-agent sys-info version v2c v3
[AGENCE-DAKAR-AR720]snmp-agent community read cipher CommunauteFictive123 acl 2001
[AGENCE-DAKAR-AR720]snmp-agent sys-info contact admin-reseau@exemple.sn
[AGENCE-DAKAR-AR720]snmp-agent sys-info location Agence-Dakar-Siege-social
```

Avec l'ACL 2001 qui n'autorise que l'IP du superviseur :

```
[AGENCE-DAKAR-AR720]acl number 2001
[AGENCE-DAKAR-AR720-acl-basic-2001]rule 5 permit source 10.0.0.100 0
[AGENCE-DAKAR-AR720-acl-basic-2001]quit
```

## 77. SNMPv3 : la version à utiliser en production

```
[AGENCE-DAKAR-AR720]snmp-agent
[AGENCE-DAKAR-AR720]snmp-agent sys-info version v3
[AGENCE-DAKAR-AR720]snmp-agent mib-view included iso-view iso
[AGENCE-DAKAR-AR720]snmp-agent group v3 supervision-group privacy read-view iso-view
[AGENCE-DAKAR-AR720]snmp-agent usm-user v3 supervision-user group supervision-group acl 2001
[AGENCE-DAKAR-AR720]snmp-agent usm-user v3 supervision-user authentication-mode sha2-256 MotDePasseAuthFictif privacy-mode aes-256 MotDePassePrivFictif
```

SNMPv3 = authentification + chiffrement. Ne jamais exposer SNMP en v1/v2c sur le WAN.
**Avertissement version** : les algorithmes proposés (sha2-256, aes-256) dépendent de la
version VRP — `sha`/`aes128` sont les valeurs sûres historiques ; vérifier avec `?` dans
la CLI du modèle exact.

## 78. Traps SNMP : être prévenu au lieu de constater

```
[AGENCE-DAKAR-AR720]snmp-agent trap enable
[AGENCE-DAKAR-AR720]snmp-agent target-host trap address udp-domain 10.0.0.100 params securityname supervision-user v3 privacy
```

Traps utiles à vérifier : linkup/linkdown, changements OSPF (neighbor down), basculement
de route, température. Tester en débranchant/rebranchant un câble et en regardant le
superviseur recevoir le trap.

## 79. Syslog : centraliser les logs

```
[AGENCE-DAKAR-AR720]info-center enable
[AGENCE-DAKAR-AR720]info-center loghost 10.0.0.101
[AGENCE-DAKAR-AR720]info-center source default channel loghost log level informational
```

- `10.0.0.101` : serveur syslog du siège (via le tunnel VPN).
- Niveau `informational` en production ; `debugging` uniquement en dépannage ciblé et
  temporaire (ça remplit les disques et charge le CPU).
- Vérifier : `display info-center` et regarder les logs arriver sur le serveur.

## 80. NetStream : voir qui consomme la bande passante

NetStream (l'équivalent NetFlow de Huawei) exporte les flux vers un collecteur :

```
[AGENCE-DAKAR-AR720]ip netstream
[AGENCE-DAKAR-AR720]ip netstream export host 10.0.0.102 9996
[AGENCE-DAKAR-AR720]ip netstream export version 9
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]ip netstream inbound
[AGENCE-DAKAR-AR720-Dialer1]ip netstream outbound
[AGENCE-DAKAR-AR720-Dialer1]quit
```

Ça répond à « qui sature le lien ? » en 2 minutes au lieu de 2 heures. Collecteur type :
nfdump/nfsen, ou le module du superviseur.

## 81. NQA/SLA : sonder en continu les liens et les tunnels

Au-delà du basculement (section 22), NQA sert à mesurer la qualité (latence, perte,
gigue) vers le siège ou vers Internet :

```
[AGENCE-DAKAR-AR720]nqa test-instance admin sla-siege
[AGENCE-DAKAR-AR720-nqa-admin-sla-siege]test-type icmp
[AGENCE-DAKAR-AR720-nqa-admin-sla-siege]destination-address ipv4 192.168.0.1
[AGENCE-DAKAR-AR720-nqa-admin-sla-siege]frequency 60
[AGENCE-DAKAR-AR720-nqa-admin-sla-siege]start now
[AGENCE-DAKAR-AR720-nqa-admin-sla-siege]quit
```

Résultats : `display nqa results test-instance admin sla-siege` (min/avg/max RTT, pertes).
Historiser via SNMP ou syslog pour prouver au FAI que « ça rame depuis mardi ».

## 82. Ce qu'il faut superviser en priorité (checklist)

| Sonde | Seuil d'alerte | Pourquoi |
|---|---|---|
| Ping du routeur (LAN + WAN) | perte | Savoir qu'il est vivant |
| État interfaces WAN | down | Panne lien |
| CPU (`display cpu-usage`) | > 80 % prolongé | Sous-dimensionnement / attaque |
| Mémoire (`display memory`) | > 85 % | Fuite / config trop grosse |
| Température | > 40 °C | Ventilation / local |
| État tunnel IPSec (SNMP) | down | VPN coupé |
| NQA vers siège | perte > 5 % | Qualité lien |
| Traps linkup/linkdown | — | Câbles, FAI |

---
---

# Partie J — Exploitation : sauvegarde, firmware, licences

## 83. Sauvegarde : la procédure complète

1. `save` sur le routeur (config en flash).
2. Copie externe immédiate après chaque changement important :
   ```
   <AGENCE-DAKAR-AR720>tftp 10.0.0.103 put vrpcfg.zip backup-AGENCE-DAKAR-20260927.zip
   ```
   (ou `ftp` / clé USB : `copy vrpcfg.zip usb0:/...`).
3. Versionner les fichiers (Git, ou dossier daté) — jamais un seul fichier « config.cfg »
   écrasé en boucle.
4. Tester la restauration **une fois** sur un équipement de spare ou en labo avant d'en
   avoir besoin pour de vrai.

## 84. Restauration : remettre une config

```
<AGENCE-DAKAR-AR720>tftp 10.0.0.103 get backup-AGENCE-DAKAR-20260927.zip vrpcfg.zip
<AGENCE-DAKAR-AR720>startup saved-configuration backup-AGENCE-DAKAR-20260927.zip
<AGENCE-DAKAR-AR720>reboot
```

Vérifier après reboot que la config chargée est la bonne (`display current-configuration
| include sysname`). En cas de doute, la console est le filet de sécurité : toujours
faire une restauration avec quelqu'un en console (ou un accès KVM), jamais en aveugle
à distance.

## 85. Mise à jour firmware : procédure détaillée + rollback

Reprise et détails de la section 14, en procédure d'exploitation :

1. Fenêtre de maintenance déclarée, utilisateurs prévenus.
2. `save` + sauvegarde externe de la config (section 83).
3. `display version` et `display device` : noter la version actuelle.
4. Copier la nouvelle image en flash, vérifier la taille (`dir`).
5. `startup system-software <nouvelle-image>` puis `display startup` pour confirmer.
6. `reboot`, suivre en console.
7. Post-reboot : `display version`, tester WAN → LAN → VPN → téléphone.
8. Laisser tourner 48 h avant d'effacer l'ancienne image.
9. **Rollback** si problème : `startup system-software <ancienne-image>` + `reboot`.

Ne jamais couper l'alimentation pendant un upgrade. Sur un site avec coupures fréquentes :
onduleur obligatoire + ne pas lancer l'upgrade un jour d'orage.

## 86. Fichiers de démarrage : comprendre `display startup`

```
<AGENCE-DAKAR-AR720>display startup
```

Affiche : image système de démarrage, fichier de config de démarrage, patch éventuel.
Après chaque upgrade ou restauration, c'est LA commande de contrôle. Une erreur ici =
le routeur redémarre sur l'ancienne version ou sans config.

## 87. Licences : ce qu'il faut savoir

