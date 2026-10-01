---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-3
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-03-29", "2026-10-25"]
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [243, 408]
sha256: 53492f4ba8e5dab87a4718fa186a52e7591b057ba1cf1b08117ba73b17886e75
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

```huawei
system-view
[USG] ntp-service unicast-server 192.0.2.10   # IP fictive : ton serveur NTP interne
[USG] clock timezone Paris add 01:00:00
[USG] clock daylight-saving-time Europe repeating 02:00 2026-03-29 03:00 2026-10-25 02:00
# Vérification :
[USG] display clock
[USG] display ntp-service status
save
```

🔧 Si le NTP ne synchronise pas : vérifie que la **politique local→untrust autorise UDP 123** (la zone local, c'est le firewall lui-même — section 19) et que le NAT ne casse pas la source.

## 16. Upgrade initiale du firmware : procédure sans stress

📋 Procédure recommandée :
1. `display version` → noter la version actuelle.
2. Télécharger le firmware **exact** du modèle sur support.huawei.com (compte requis).
3. Vérifier le **checksum** (MD5/SHA) du fichier téléchargé.
4. Sauvegarder la config : `save` puis exporter (`tftp`/`ftp`/USB, section 115).
5. Transférer le firmware : via web (System > Upgrade), ou CLI :
```huawei
# Exemple via FTP (adresses fictives) :
[USG] ftp 192.0.2.50
# login, puis :
get usg6000-firmware.cc
quit
[USG] startup system-software usg6000-firmware.cc
[USG] display startup          # vérifier que le nouveau firmware est bien "next startup"
[USG] reboot
```
6. Après reboot : `display version`, tester le trafic, vérifier les licences (`display license`).
7. ⚠️ Garde l'**ancien firmware en secours** : `startup system-software` ne supprime pas l'ancien — en cas de problème, redémarre sur l'ancienne version (section 117).

## 17. Activation des licences (de base)

```huawei
[USG] display license          # voir les licences actives
[USG] display esn              # numéro de série pour le portail licence
```

- Les licences UTM (AV, IPS, URL, anti-spam) s'activent avec un fichier `.dat` via web (System > License) ou CLI.
- Sans licence UTM valide, les profils UTM **ne protègent pas** (signatures non mises à jour) — le firewall continue de filtrer, mais sans inspection de contenu.
- 📋 Note la **date d'expiration** de chaque licence dans ton planning de maintenance (section 122).

## 18. Sauvegarde de la config « sortie de carton »

Avant TOUTE config métier, sauvegarde l'état initial :
```huawei
[USG] save                     # sauvegarde dans la mémoire flash
[USG] display current-configuration  # pour export vers fichier
```
Exporte aussi vers un serveur TFTP/FTP ou une clé USB (section 115). Nomme le fichier `USG6000_<site>_initiale_<date>.cfg`. Ce fichier, c'est ton **point de retour** si la config métier part en vrille.

---
---

# BLOC C — ZONES DE SÉCURITÉ

## 19. Le concept : pourquoi des zones et pas juste des interfaces

Sur un USG, **on ne filtre pas entre interfaces, on filtre entre ZONES**. Une zone = un ensemble d'interfaces partageant le même niveau de confiance. Avantages :
- La politique suit la **logique métier** (LAN → Internet), pas le câblage.
- Si tu changes de port physique, tu déplaces l'interface de zone — les politiques ne bougent pas.
- Le firewall lui-même est une zone (**local**) : on contrôle qui peut l'administrer, le pinger, lui envoyer du NTP/DNS.

## 20. Les 4 zones prédéfinies (priorités par défaut)

| Zone | Priorité | Rôle | Interfaces typiques |
|------|----------|------|---------------------|
| **local** | 100 | Le firewall lui-même | (implicite, pas d'interface à ajouter) |
| **trust** | 85 | Réseau de confiance | LAN bureautique |
| **dmz** | 50 | Zone démilitarisée | Serveurs exposés |
| **untrust** | 5 | Non fiable | Internet / WAN FAI |

⚠️ Points critiques :
- Une interface ne peut appartenir qu'à **UNE SEULE** zone.
- Une interface **sans zone** ne fait transiter **aucun** trafic (sauf management hors-bande selon config).
- La priorité sert aux **politiques par défaut** et à certaines fonctions (ex. : par défaut, le trafic d'une zone de priorité haute vers basse peut être autorisé selon la politique par défaut configurée — **vérifie toujours ta politique par défaut**, section 29).
- La zone **local** ne contient aucune interface : c'est le trafic **à destination ou en provenance du firewall lui-même** (SSH, HTTPS d'admin, NTP, DNS proxy, VPN terminé sur le boîtier).

## 21. Créer une zone personnalisée (ex. : GUEST, IOT, PROD)

```huawei
system-view
[USG] firewall zone name GUEST
[USG-zone-guest] set priority 20        # entre untrust (5) et dmz (50)
[USG-zone-guest] add interface GigabitEthernet 1/0/5
[USG-zone-guest] quit
[USG] firewall zone name IOT
[USG-zone-iot] set priority 30
[USG-zone-iot] add interface GigabitEthernet 1/0/6
[USG-zone-iot] quit
# Vérification :
[USG] display zone
[USG] display zone name GUEST
save
```

Conseil de nommage : **MAJUSCULES, explicites, sans espaces** (GUEST, IOT, PROD, BACKUP, VIDEO). Évite les accents.

## 22. Construire une DMZ propre

Schéma type :
```
Internet ── [untrust: GE1/0/1] ── USG ── [dmz: GE1/0/2] ── Serveurs (172.16.1.0/24)
                                   └──── [trust: GE1/0/3] ── LAN (192.168.10.0/24)
```

```huawei
system-view
[USG] interface GigabitEthernet 1/0/2
[USG-GigabitEthernet1/0/2] ip address 172.16.1.1 24
[USG-GigabitEthernet1/0/2] quit
[USG] firewall zone dmz
[USG-zone-dmz] set priority 50
[USG-zone-dmz] add interface GigabitEthernet 1/0/2
[USG-zone-dmz] quit
save
```

📋 Checklist DMZ :
- [ ] Les serveurs DMZ ont des IP **fixes** (ou DHCP avec réservation).
- [ ] Politique **untrust→dmz** : uniquement les ports nécessaires (80/443 vers le reverse proxy, jamais d'admin).
- [ ] Politique **dmz→trust** : **DENY par défaut** — un serveur DMZ compromis ne doit pas rebondir vers le LAN.
- [ ] Politique **dmz→untrust** : limitée (mises à jour, DNS, NTP via proxy si possible).
- [ ] Les serveurs DMZ **ne résolvent pas** le DNS interne directement (split-DNS ou forwarder dédié).

## 23. Zone local : protéger le firewall lui-même

La zone local, c'est **toi qui décides qui touche au boîtier** :

```huawei
system-view
[USG] security-policy
[USG-policy-security] rule name local-mgmt-ssh
[USG-policy-security-rule-local-mgmt-ssh] source-zone trust
[USG-policy-security-rule-local-mgmt-ssh] destination-zone local
[USG-policy-security-rule-local-mgmt-ssh] source-address 192.168.10.0 24
[USG-policy-security-rule-local-mgmt-ssh] service ssh https
[USG-policy-security-rule-local-mgmt-ssh] action permit
[USG-policy-security-rule-local-mgmt-ssh] quit
[USG-policy-security] rule name local-deny-untrust
[USG-policy-security-rule-local-deny-untrust] source-zone untrust
[USG-policy-security-rule-local-deny-untrust] destination-zone local
[USG-policy-security-rule-local-deny-untrust] action deny
[USG-policy-security-rule-local-deny-untrust] quit
[USG-policy-security] quit
save
```

⚠️ **Piège classique** : une politique trop restrictive vers local **coupe ton accès SSH/HTTPS** — tu te retrouves à devoir aller en console en baie. Toujours tester une nouvelle règle local depuis une session existante, et garder la console à portée.

## 24. Voir et diagnostiquer les zones

```huawei
[USG] display zone                          # toutes les zones et leurs interfaces
[USG] display zone name trust               # détail d'une zone
[USG] display firewall zone statistic       # stats par zone (selon version)
[USG] display interface brief               # état des interfaces
```

🔧 « Mon interface est UP mais rien ne passe » → dans 50 % des cas, **l'interface n'est dans aucune zone**. `display zone` en 10 secondes tranche.

## 25. Priorités de zones : à quoi ça sert vraiment

