---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-5
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-03-29", "2026-09-27", "2026-10-25"]
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [870, 1080]
sha256: aafac3c22eaa9729e798841609f770686971bc90ec44ec494f4f9f1e08644bcd
---

# VRP — Le système d'exploitation transversal Huawei

| Thème | Commandes |
|---|---|
| Système | `display version`, `display device`, `display device elabel`, `display clock`, `display startup` |
| Interfaces | `display interface brief`, `display interface X`, `display ip interface brief` |
| L2 | `display mac-address`, `display vlan`, `display stp brief`, `display eth-trunk` |
| L3 | `display ip routing-table`, `display fib`, `display arp`, `display ip pool` |
| Routage dyn. | `display ospf peer`, `display bgp peer`, `display isis peer` |
| Logs | `display logbuffer`, `display trapbuffer` |
| Perf | `display cpu-usage`, `display memory`, `display health` |
| Sécurité | `display users`, `display ssh server status`, `display acl all` |
| USG | `display firewall session table`, `display security-policy rule all`, `display hrp state` |
| Fichiers | `dir`, `display startup`, `display saved-configuration` |

---


## 51. Configuration minimale d'un équipement neuf

```vrp
<AR720>system-view
[AR720]sysname AR720-SIEGE
[AR720-SIEGE]clock timezone Paris add 02:00:00
[AR720-SIEGE]clock daylight-saving-time Paris repeating 03:00:00 2026-03-29 03:00:00 2026-10-25 one-hour
[AR720-SIEGE]interface GigabitEthernet 0/0/0
[AR720-SIEGE-GigabitEthernet0/0/0]ip address 192.168.1.1 24
[AR720-SIEGE-GigabitEthernet0/0/0]description MGMT
[AR720-SIEGE-GigabitEthernet0/0/0]quit
[AR720-SIEGE]ip route-static 0.0.0.0 0.0.0.0 192.168.1.254
[AR720-SIEGE]quit
<AR720-SIEGE>save
```

## 52. `sysname` et `clock` — les bases

```vrp
[AR720]sysname S310-ETAGE1
[AR720]clock timezone Paris add 01:00:00
[AR720]clock datetime 14:30:00 2026-09-27      # réglage manuel (si pas de NTP)
```

Convention de nommage recommandée (voir section 108) :
`<Type>-<Site>-<Rôle>-<Num>` ex. `AR720-SIEGE-WAN-01`, `S310-ETAGE1-ACC-02`,
`USG6000-SIEGE-FW-01`.

## 53. Utilisateurs locaux et niveaux (AAA)

```vrp
[AR720]aaa
[AR720-aaa]local-user admin password irreversible-cipher MotDePasseFort123!
[AR720-aaa]local-user admin privilege level 15
[AR720-aaa]local-user admin service-type terminal ssh telnet http ftp
[AR720-aaa]local-user operateur password irreversible-cipher AutreMotDePasse456!
[AR720-aaa]local-user operateur privilege level 1
[AR720-aaa]local-user operateur service-type terminal ssh
[AR720-aaa]quit
```

- `irreversible-cipher` : hachage non réversible (préféré à `cipher`).
- `privilege level 15` = droits complets ; `level 1` = supervision seule.
- `service-type` limite les modes d'accès autorisés pour cet utilisateur.

> ⚠️ Ne laissez JAMAIS un compte sans mot de passe ou avec le mot de passe
> par défaut en production. À la première connexion, VRP force souvent la
> définition d'un mot de passe — ne contournez pas cette étape.

## 54. Sécuriser la console et les VTY

```vrp
[AR720]user-interface console 0
[AR720-ui-console0]authentication-mode aaa
[AR720-ui-console0]user privilege level 3
[AR720-ui-console0]idle-timeout 10 0
[AR720-ui-console0]quit
[AR720]user-interface vty 0 4
[AR720-ui-vty0-4]authentication-mode aaa
[AR720-ui-vty0-4]protocol inbound ssh          # SSH uniquement, pas de Telnet
[AR720-ui-vty0-4]user privilege level 3
[AR720-ui-vty0-4]idle-timeout 10 0
[AR720-ui-vty0-4]acl 2000 inbound              # restreint les IP sources (optionnel)
[AR720-ui-vty0-4]quit
```

`idle-timeout 10 0` = déconnexion après 10 minutes d'inactivité (défaut
système : 10 minutes). `authentication-mode aaa` délègue l'authentification
à la base locale (ou RADIUS, voir section 88).

## 55. `command-privilege` — granularité fine

Permet de baisser le niveau requis d'une commande (ex. pour un outil de
sauvegarde qui ne doit pas être admin) :

```vrp
[AR720]command-privilege level 1 view shell display current-configuration
[AR720]command-privilege level 1 view shell screen-length
```

Syntaxe : `command-privilege level <0-15> view <vue> <commande>`.
Après cela, un utilisateur de niveau 1 peut exécuter ces commandes.
Vérifiez avec `display command-privilege` (à vérifier selon version).

## 56. Banner et messages légaux

```vrp
[AR720]header login information "ACCES RESERVE - Service Systemes & Energies"
[AR720]header shell information "Toute connexion est journalisee."
```

Variantes selon version : `banner`, `header login`, `header shell`. Le
banner légal protège juridiquement en cas d'intrusion (mention « accès
réservé »).

## 57. DNS et nom de domaine

```vrp
[AR720]dns resolve
[AR720]dns server 192.168.100.10
[AR720]dns server 192.168.100.11 secondary
[AR720]ip host mon-serveur 192.168.10.50     # entrée statique (à vérifier)
```

## 58. NTP — synchroniser l'horloge

```vrp
[AR720]ntp-service unicast-server 192.168.100.1
[AR720]ntp-service unicast-server 192.168.100.2
[AR720]clock timezone Paris add 01:00:00
[AR720]quit
<AR720>display ntp-service status
```

Sans NTP, impossible de corréler les logs entre AR720, S310 et USG6000.
Pointez tous les équipements vers la même source (serveur interne ou
`pool.ntp.org` si autorisé).

## 59. Le système de fichiers : `dir`, `cd`, `pwd`

```vrp
<AR720>dir
Directory of flash:/
  Idx  Attr     Size(Byte)  Date        Time       FileName
    0  -ro-     84,934,656  Sep 26 2026 22:10:00   AR720-V300R022C00SPC500.cc
    1  -rw-          8,192  Sep 27 2026 01:00:00   vrpcfg.zip
    2  -rw-        123,456  Sep 20 2026 10:00:00   backup_avant_upgrade.zip
    3  drw-              -  Sep 27 2026 00:00:00   logfile
42,123,456 bytes total (28,000,000 bytes free)
<AR720>dir flash:/logfile/
<AR720>pwd
flash:
<AR720>cd logfile/
<AR720>pwd
flash:/logfile
<AR720>cd ..
```

Unités : `flash:` (mémoire interne), `usb0:`/`usbf:` (clé USB si présente),
`sd0:` (carte SD sur certains modèles).

## 60. Lire, copier, renommer, supprimer

```vrp
<AR720>more vrpcfg.zip                         # affiche le contenu (limité pour un .zip)
<AR720>copy flash:/vrpcfg.zip flash:/backup_2026-09-27.zip
<AR720>rename backup_2026-09-27.zip backup_old.zip
<AR720>delete backup_old.zip
Delete flash:/backup_old.zip? [y/n]:y
<AR720>undelete backup_old.zip                 # restaure depuis la corbeille
<AR720>reset recycle-bin                       # vide DÉFINITIVEMENT la corbeille
```

Points clés :
- `delete` envoie dans la **corbeille** (recyclable avec `undelete`).
- `delete /unreserved` ou `reset recycle-bin` = suppression définitive.
- `more` pagine automatiquement le contenu.

## 61. `save` en détail et fichiers de configuration

```vrp
<AR720>save                                     # sauvegarde vers le fichier par défaut
<AR720>save backup_manuelle.zip                # sauvegarde vers un fichier nommé
<AR720>display saved-configuration             # contenu du fichier de démarrage
<AR720>display saved-configuration | include sysname
<AR720>display current-configuration           # config active en RAM
<AR720>display current-configuration interface GigabitEthernet 0/0/0
```

Fichiers en jeu :
- **Fichier de démarrage** : celui chargé au boot (défaut `vrpcfg.zip`,
  modifiable avec `startup saved-configuration`).
- **Config courante** : en RAM, ce que `display current-configuration`
  montre.
- Si les deux diffèrent et qu'on reboot → on perd les modifs non sauvées.

## 62. `compare configuration` — le diff avant/après

```vrp
<AR720>compare configuration
```

Affiche les différences entre la configuration courante (RAM) et la
configuration sauvegardée (fichier de démarrage). À exécuter **avant**
chaque `save` en intervention : on relit ce qu'on s'apprête à figer.

## 63. Sauvegarde automatique : `autosave`

```vrp
[AR720]autosave interval on                    # sauvegarde périodique
[AR720]autosave interval 60                    # toutes les 60 minutes (défaut 1440)
[AR720]autosave time on                        # sauvegarde à heure fixe
[AR720]autosave time 03:00:00                  # tous les jours à 3h00
```

