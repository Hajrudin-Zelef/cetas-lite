---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-1
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [1, 183]
sha256: 40e428be1690219507b411e9279480be2ea98404444cd8bb67910effc4d766d5
---

# VRP — Le système d'exploitation transversal Huawei

**Guide technique ultra-complet** — Versatile Routing Platform (VRP)
Public : Zelef, chef de service systèmes & énergies
Équipements couverts : routeurs AR720, switchs S310, pare-feu USG6000
Date de rédaction : 2026-09-27 — Langue : français
Objectif : faire de VRP un réflexe transversal, quel que soit le boîtier.

> Convention de notation dans tout ce guide :
> - `<AR720>` = invite de la **user view** (prompt commençant par `<` et se terminant par `>`).
> - `[AR720]` = invite de la **system view** (prompt entre crochets).
> - `[AR720-GigabitEthernet0/0/0]` = invite d'une **interface view**.
> - Les commandes sont à taper sans les chevrons ni crochets.
> - Le nom d'hôte varie (`<S310>`, `<USG6000>`) mais la logique reste identique.

---

## 1. VRP : qu'est-ce que c'est ?

VRP (Versatile Routing Platform) est le système d'exploitation commun à
quasiment tout le portefeuille réseau Huawei : routeurs (séries AR, NE),
switchs (séries S, CloudEngine), pare-feu (séries USG), contrôleurs WLAN (AC)
et passerelles eKit. Un seul CLI à apprendre, des dizaines de boîtiers
pilotés. C'est l'équivalent de Cisco IOS/IOS-XE côté Huawei.

Pourquoi c'est important pour Zelef : l'AR720 (routage WAN), le S310
(commutation d'accès) et l'USG6000 (sécurité périmétrique) partagent 90 % des
commandes. Apprendre une fois = exploiter trois familles.

## 2. Architecture logicielle de VRP

VRP est un OS temps réel modulaire :

- **Noyau temps réel** : ordonnancement préemptif, gestion mémoire protégée.
- **Processus par fonction** : chaque protocole (OSPF, BGP, SNMP...) tourne
  dans son processus ; un plantage d'un processus ne fait pas tomber le
  système entier.
- **Base de configuration** : la configuration active vit en RAM ; elle doit
  être explicitement sauvegardée en flash (`save`), sinon elle est perdue au
  reboot. C'est LA différence culturelle majeure avec certains OS qui
  sauvegardent en continu.
- **Fichiers système** : le logiciel système est un fichier `.cc` (ex.
  `AR720-V300R022C00SPC500.cc`) ; la configuration est un fichier `.zip`
  (par défaut `vrpcfg.zip`) ; les patchs sont des `.pat`.

## 3. Les grandes versions : VRP5 vs VRP8

| Point | VRP5 (classique) | VRP8 (nouvelle génération) |
|---|---|---|
| Modèle de config | Direct : chaque commande s'applique immédiatement | **Candidate database** : les commandes vont dans une base candidate, il faut `commit` |
| Prompt si non commité | n/a | `[~Nom]` : le tilde `~` signale des changements non commités |
| Commande de validation | `save` | `commit` puis `save` |
| Annulation | `undo <commande>` | `rollback configuration` / `commit` annulé avant validation |
| Équipements | AR, S (anciennes versions), USG6000 | CloudEngine récents, NE, versions VRP8 des AR/S |
| Compatibilité syntaxe | Référence historique | ~95 % identique à VRP5 pour le quotidien |

Règle d'or : regardez toujours `display version` AVANT de travailler.
Si le prompt affiche `[~...]`, vous êtes en VRP8 avec des modifications en
attente — pensez `commit` ou `rollback`.

## 4. Quels équipements utilisent VRP ?

| Famille | Exemples | Rôle typique | VRP |
|---|---|---|---|
| Routeurs d'accès | AR720, AR2200, AR1000V | Routage WAN, VPN, NAT | VRP5/VRP8 |
| Switchs campus | S310, S5735, S6730 | Commutation L2/L3 | VRP5/VRP8 |
| Pare-feu | USG6000, USG6600 | Filtrage, NAT, VPN IPsec | VRP5 |
| Contrôleurs WLAN | AC6508, AC6805 | Gestion des AP | VRP5 |
| Routeurs cœur | NE40E, NE9000 | Cœur de réseau opérateur | VRP8 |
| eKit | Switchs/AP eKit | PME | VRP simplifié |

Dans ce guide, les exemples utilisent l'AR720 comme référence, avec des
notes « S310 » et « USG6000 » quand le comportement diffère.

## 5. Premier contact : la console

Paramètres série (PuTTY, Tera Term, minicom) :

- Débit : **9600 bauds**
- Bits de données : 8
- Bit de stop : 1
- Parité : aucune
- Contrôle de flux : aucun

```text
Câble console (RJ45 ou USB) -> port CONSOLE de l'équipement
PuTTY : Connection type = Serial, Speed = 9600
```

À la première mise sous tension, VRP propose souvent l'**Auto-Config**
(découverte automatique via DHCP). En production, répondez `n` (non) puis
configurez manuellement, sinon les réglages DHCP/VTY seront écrasés.

> ⚠️ Si le message « Warning: Auto-Config is working... Do you want to stop
> Auto-Config? [y/n]: » apparaît, répondez **y** pour reprendre la main.

## 6. Philosophie du CLI : display, undo, quit

Trois verbes gouvernent 80 % du quotidien :

| Verbe | Rôle | Exemple |
|---|---|---|
| `display` | Afficher (jamais `show` !) | `display version` |
| `undo` | Annuler / désactiver (équivalent du `no` Cisco) | `undo shutdown` |
| `quit` | Remonter d'un niveau de vue | `quit` |

```vrp
<S310>display version          # afficher
[S310]undo telnet server enable # désactiver
[S310-GigabitEthernet0/0/1]quit # remonter
```

## 7. La règle du `save` — lisez ceci deux fois

VRP ne sauvegarde **rien** automatiquement (sauf si `autosave` est
configuré, voir section 63). Toute configuration vit en RAM jusqu'au `save`.

```vrp
<AR720>save
  The current configuration will be written to the device.
  Are you sure to continue? (y/n)[n]:y
  Now saving the current configuration to the slot 1.
  Save the configuration successfully.
```

**Incident classique n°1** (voir cas de dépannage n°2) : on configure un
vendredi soir, on oublie `save`, coupure de courant le week-end, tout est
perdu lundi matin. Prenez le réflexe : **chaque session de configuration se
termine par `save` + `display saved-configuration` pour vérifier.**

## 8. Abréviations et complétion

VRP accepte les formes abrégées tant qu'elles sont **uniques** :

```vrp
<AR720>dis ver              # = display version
<AR720>sys                  # = system-view
[AR720]dis cu               # = display current-configuration
[AR720-GigabitEthernet0/0/1]dis th   # = display this
```

Tabulation = complétion automatique. `?` = aide contextuelle (voir
section 20).

## 9. Sensibilité à la casse et espaces

- Les mots-clés sont **insensibles à la casse** (`Display` = `display`),
  mais prenez l'habitude du minuscule.
- Les noms d'objets (sysname, descriptions, noms de VLAN) conservent la
  casse telle que saisie.
- Un espace sépare mot-clé et paramètre : `display?` (collé) affiche l'aide
  sur `display`, `display ?` (espace) liste les sous-options.

## 10. Les invites (prompts) par vue — mémo visuel

```text
<AR720>                              USER VIEW
[AR720]                              SYSTEM VIEW
[AR720-GigabitEthernet0/0/1]         INTERFACE VIEW
[AR720-aaa]                          AAA VIEW
[AR720-ospf-1]                       OSPF VIEW
[AR720-bgp]                          BGP VIEW
[AR720-ui-vty0-4]                    VTY USER-INTERFACE VIEW
[AR720-ui-console0]                  CONSOLE USER-INTERFACE VIEW
[AR720-vlan10]                       VLAN VIEW
[AR720-wlan-view]                    WLAN VIEW (contrôleur/AR avec WLAN)
[AR720-zone-trust]                   SECURITY ZONE VIEW (USG6000)
[AR720-policy-security]              SECURITY POLICY VIEW (USG6000)
[~AR720]                             SYSTEM VIEW avec modifs non commitées (VRP8)
```

## 11. Niveaux d'utilisateurs 0 à 3 et privilèges 0 à 15

VRP distingue **niveau d'utilisateur** (0-3, grandes familles de droits) et
**niveau de privilège de commande** (0-15, granularité fine).

| Niveau utilisateur | Nom | Commandes accessibles |
|---|---|---|
| 0 | Visit | Diagnostic réseau de base : `ping`, `tracert`, `telnet` client, quelques `display` |
| 1 | Monitoring | Niveau 0 + maintenance système : tous les `display` |
| 2 | Configuration | Niveaux 0-1 + configuration des services (routage, IP, VLAN...) |
| 3 | Management | Niveaux 0-2 + tout : fichiers, FTP/TFTP, utilisateurs, `debugging`, `reboot` |

