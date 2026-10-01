---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-13
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [1892, 2061]
sha256: 58f1f9f2d14bb8bae40e5cb9eb17d56d3ccb169a4deb46fdf56c2892f9b71ab0
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

```huawei
system-view
[USG] info-center loghost 192.168.10.100 facility local7   # IP fictive : ton syslog
[USG] info-center loghost 192.168.10.100 port 514
[USG] info-center source default channel loghost log level informational
save
```

🔧 Tester : générer un log (`ping` bloqué par une règle avec logging) et vérifier sa réception côté serveur. **Sans test, tu découvriras la panne de logs le jour où tu en auras besoin.**

## 110. Logging par règle : quoi logger (et quoi ne pas logger)

```huawei
[USG] security-policy
[USG-policy-security] rule name LAN-web-out
# ... critères ...
[USG-policy-security-rule-LAN-web-out] action permit
# Activer le log :
[USG-policy-security-rule-LAN-web-out] counting enable
[USG-policy-security-rule-LAN-web-out] log enable      # syntaxe à vérifier selon version
[USG-policy-security-rule-LAN-web-out] quit
[USG-policy-security] quit
save
```

Stratégie :
- **Logger** : les `deny` (qui essaie de passer ?), les règles sensibles (DMZ, VPN, admin).
- **Ne PAS logger en détail** : les `permit` à fort volume (LAN→Internet web) — sinon tu noies le syslog. Utilise les **compteurs** (`counting`) + **échantillonnage** ou logs agrégés.
- ⚠️ Un syslog qui reçoit 10 000 logs/seconde = un syslog qui **perd** des logs. Dose.

## 111. Rapports intégrés : ce que l'USG sait faire seul

Avec disque dur : **rapports de trafic, top applications, top utilisateurs, top menaces** — générables depuis le web (Monitor > Reports).
- Planifier un **rapport hebdo PDF** envoyé par mail (selon version).
- 📋 Le rapport hebdo, c'est ton **radar** : une appli inconnue dans le top 10 = enquête immédiate.

## 112. SNMP : supervision dans Zabbix/Centreon/Prometheus

```huawei
system-view
[USG] snmp-agent
[USG] snmp-agent sys-info version v3
[USG] snmp-agent local-user snmp-admin
[USG] snmp-agent local-user snmp-admin password   # mot de passe fictif à saisir
[USG] snmp-agent local-user snmp-admin privilege level 3
[USG] snmp-agent target-host trap address udp-domain 192.168.10.100 params securityname snmp-admin v3
# Autoriser SNMP vers la zone local :
[USG] security-policy
[USG-policy-security] rule name snmp-local
[USG-policy-security-rule-snmp-local] source-zone trust
[USG-policy-security-rule-snmp-local] destination-zone local
[USG-policy-security-rule-snmp-local] source-address 192.168.10.100 32
[USG-policy-security-rule-snmp-local] service snmp
[USG-policy-security-rule-snmp-local] action permit
[USG-policy-security-rule-snmp-local] quit
[USG-policy-security] quit
save
```

Superviser en priorité : **CPU, mémoire, sessions actives, état HA, état des interfaces, état des tunnels VPN, expiration des licences, état des disques**. MIB Huawei : `1.3.6.1.4.1.2011` (l'OID du template USG6620 vu en recherche : `.1.3.6.1.4.1.2011.6.133.55`).

## 113. Traps SNMP : être prévenu au lieu de constater

Configurer les **traps** pour : bascule HA, interface down, CPU > 80 %, mémoire > 85 %, tunnel VPN down, échec de MAJ de signatures, disque en erreur.

```huawei
[USG] snmp-agent trap enable
# Activer les traps par module (syntaxe à vérifier selon version) :
[USG] snmp-agent trap enable feature-name hrp
[USG] snmp-agent trap enable feature-name ipsec
save
```

📋 **Tester chaque trap** après config (provoquer l'événement en maquette ou en heures creuses). Un trap jamais testé = une alerte qui n'arrivera jamais.

## 114. Sondes applicatives : vérifier que « ça marche vraiment »

Au-delà du ping :
- **NQA/SLA** : l'USG peut sonder périodiquement (ping, TCP connect, HTTP GET) vers des cibles critiques (serveur DMZ, DNS, passerelle FAI) et **déclencher des actions** (bascule de route, trap).
- Exemple : si le ping vers 8.8.8.8 via le lien principal échoue 3 fois → bascule sur le lien de secours (voir section 161, double WAN).

```huawei
system-view
[USG] nqa test-instance admin ping-fai
[USG-nqa-admin-ping-fai] test-type icmp
[USG-nqa-admin-ping-fai] destination-address ipv4 8.8.8.8
[USG-nqa-admin-ping-fai] frequency 30
[USG-nqa-admin-ping-fai] probe-count 3
[USG-nqa-admin-ping-fai] start now
[USG-nqa-admin-ping-fai] quit
[USG] display nqa results test-instance admin ping-fai
save
```

---
---

# BLOC K — SAUVEGARDE, FIRMWARE, LICENCES

## 115. Sauvegarde de configuration : les 3 méthodes

**Méthode 1 — CLI + TFTP** (la plus rapide) :
```huawei
[USG] save
[USG] tftp 192.168.10.50 put vrpcfg.cfg USG6000_siteA_20260927.cfg
```

**Méthode 2 — Web** : System > Configuration File Management > Export.

**Méthode 3 — USB** : brancher la clé, web ou CLI :
```huawei
[USG] copy vrpcfg.cfg usb0:/USG6000_siteA_20260927.cfg
```

📋 Nommer : `USG6000_<site>_<date>_<raison>.cfg` (ex. `USG6000_Siege_20260927_avant-MAJ.cfg`). Garder **au moins 3 générations** + la version « sortie de carton » (section 18).

## 116. Restauration : revenir en arrière proprement

```huawei
# Depuis un fichier sur TFTP :
[USG] tftp 192.168.10.50 get USG6000_Siege_20260927_avant-MAJ.cfg vrpcfg.cfg
[USG] startup saved-configuration USG6000_Siege_20260927_avant-MAJ.cfg
[USG] reboot
# Ou appliquer à chaud (DANGEREUX : peut couper l'accès) :
[USG] configuration replace file USG6000_Siege_20260927_avant-MAJ.cfg   # selon version
```

⚠️ La restauration à chaud peut **couper ta session** et le trafic. Préfère le reboot planifié, avec **console à portée** et **quelqu'un sur site** (ou en astreinte joignable).

## 117. Firmware : procédure complète + rollback

**Procédure d'upgrade** (détail de la section 16) :
1. Vérifier compatibilité modèle/version (matrice Huawei).
2. Sauvegarder config (section 115) + noter `display version` et `display license`.
3. Transférer le firmware (FTP/TFTP/web/USB), vérifier le checksum.
4. `startup system-software <fichier>` → `display startup` → `reboot`.
5. Vérifier : version, config intacte, trafic, VPN, HA, licences.

**Rollback** :
```huawei
[USG] display startup              # voir les firmwares disponibles
[USG] startup system-software <ancien-fichier>   # repointer vers l'ancienne version
[USG] reboot
```

💡 L'USG garde généralement **l'ancien firmware** disponible — ne le supprime pas « pour faire de la place » avant d'avoir validé le nouveau en prod pendant **au moins une semaine**.

## 118. Restauration d'urgence : mot de passe perdu

Si plus aucun accès admin :
1. Accès **console physique** obligatoire.
2. Au boot, interrompre pour entrer dans le **BootROM/BootLoad** (touche indiquée à l'écran, souvent `Ctrl+B`).
3. Option **« Clear password » / skip config** (intitulé exact à vérifier selon version) → reboot **sans charger la config**.
4. Reconfigurer le mot de passe, puis **recharger la config sauvegardée** (d'où l'importance de la section 115).
5. ⚠️ Cette procédure **efface la config courante** si mal exécutée — ne la faire qu'avec une sauvegarde valide sous la main.

## 119. Licences UTM : gestion et renouvellement

```huawei
[USG] display license               # état des licences
[USG] display license resource      # ressources couvertes (selon version)
```

- Les licences UTM (AV, IPS, URL, anti-spam) sont **annuelles** : sans renouvellement, les **signatures ne se mettent plus à jour** (la protection se dégrade, le firewall continue de filtrer).
- 📋 **Calendrier** : noter chaque expiration à **J-60** (commande), **J-30** (relance), **J-7** (alerte). Un renouvellement oublié = 3 mois de signatures obsolètes avant de s'en rendre compte.
- Activation : fichier `.dat` via web (System > License > Import) — lier au **numéro de série** (ESN) du boîtier.
- En HA : **une licence par nœud** (section 103).

## 120. Fichiers de signatures : sauvegarde et restauration

