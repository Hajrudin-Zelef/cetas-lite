---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-2
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [184, 432]
sha256: dac2f8aee4446f467a19dc63d588f2ffe65a21203b29acab60ef30f435389fb6
---

# VRP — Le système d'exploitation transversal Huawei

Chaque commande possède un niveau de privilège par défaut (0-15) ; un
utilisateur de niveau N exécute les commandes de niveau ≤ N. On peut
réassigner une commande avec `command-privilege` (voir section 55).

## 12. Où trouver la documentation officielle

- `display ?` et `?` dans chaque vue : la documentation embarquée, toujours
  à jour pour VOTRE version.
- Fiches EDOC sur support.huawei.com (nécessite un compte) : « Command
  Reference » par produit et version — la source de vérité pour la syntaxe
  exacte.
- Les guides CLI PDF (ex. AR720, S310, USG6000 Command Reference) : Zelef
  les a déjà en partie via ses collectes (voir mémoire du projet RAG).

---

## 13. La user view — le point d'entrée

```text
<AR720>
```

C'est la vue d'accueil après login. On y fait de la **supervision** et du
**diagnostic**, pas de la configuration :

```vrp
<AR720>display version
<AR720>display current-configuration
<AR720>ping 8.8.8.8
<AR720>tracert 8.8.8.8
<AR720>display interface brief
<AR720>save
<AR720>reboot
```

Toutes les commandes `display`, `ping`, `tracert`, `debugging`, `save`,
`reboot`, `dir`, `tftp`, `ftp` vivent ici (ou y sont accessibles).

## 14. Entrer dans la system view

```vrp
<AR720>system-view
Enter system view, return user view with Ctrl+Z.
[AR720]
```

Le message est explicite : **Ctrl+Z** ramène directement à la user view,
d'où que vous soyez. C'est le réflexe panique à connaître.

## 15. La system view — le centre de configuration

```text
[AR720]
```

Toutes les configurations globales partent d'ici : `sysname`, `vlan`,
`interface`, `aaa`, `ospf`, `bgp`, `acl`, `ip route-static`, etc.

```vrp
[AR720]sysname AR720-SIEGE
[AR720-SIEGE]vlan 10
[AR720-SIEGE-vlan10]description USERS
[AR720-SIEGE-vlan10]quit
[AR720-SIEGE]quit
<AR720-SIEGE>
```

Notez comment le prompt suit le sysname : après `sysname AR720-SIEGE`, le
prompt devient `[AR720-SIEGE]`.

## 16. Les vues d'interface

```vrp
[AR720]interface GigabitEthernet 0/0/0
[AR720-GigabitEthernet0/0/0]ip address 192.168.1.1 24
[AR720-GigabitEthernet0/0/0]description LIEN_VERS_S310
[AR720-GigabitEthernet0/0/0]undo shutdown
[AR720-GigabitEthernet0/0/0]quit
[AR720]
```

Types d'interfaces rencontrés :

| Famille | Nommage typique | Exemple |
|---|---|---|
| AR720 | `GigabitEthernet0/0/X`, `Eth-Trunk`, `Dialer`, `Tunnel`, `Vlanif` (si switching) | `GE0/0/0` |
| S310 | `GigabitEthernet0/0/X`, `XGigabitEthernet0/0/X`, `Vlanif`, `Eth-Trunk`, `NULL0` | `GE0/0/1` |
| USG6000 | `GigabitEthernet0/0/X`, `XGigabitEthernet`, `Vlanif`, `Tunnel`, `Eth-Trunk` | `GE0/0/0` |

Astuce : `interface ?` en system view liste tous les types supportés par
VOTRE boîtier. Les abréviations `gi`, `xgi`, `Eth-Trunk`/`Eth-Trunk` sont
acceptées si uniques.

## 17. Les vues de protocoles (protocol views)

Chaque protocole de routage a sa vue dédiée :

```vrp
[AR720]ospf 1 router-id 1.1.1.1
[AR720-ospf-1]area 0
[AR720-ospf-1-area-0.0.0.0]network 192.168.1.0 0.0.0.255
[AR720-ospf-1-area-0.0.0.0]quit
[AR720-ospf-1]quit
[AR720]bgp 65001
[AR720-bgp]router-id 1.1.1.1
[AR720-bgp]peer 10.0.0.2 as-number 65002
[AR720-bgp]quit
[AR720]isis 1
[AR720-isis-1]network-entity 49.0001.0010.0100.1001.00
[AR720-isis-1]quit
[AR720]rip 1
[AR720-rip-1]version 2
[AR720-rip-1]network 192.168.1.0
[AR720-rip-1]quit
```

Autres vues de protocoles/services : `[AR720-aaa]`, `[AR720-acl-basic-2000]`,
`[AR720-ip-pool-POOL1]` (DHCP), `[AR720-nqa-admin-test]`.

## 18. Les vues spécifiques USG6000 : zones et politiques de sécurité

Sur l'USG6000, le pare-feu organise tout autour des **zones de sécurité** :

```vrp
[USG6000]firewall zone trust
[USG6000-zone-trust]add interface GigabitEthernet 0/0/1
[USG6000-zone-trust]quit
[USG6000]firewall zone untrust
[USG6000-zone-untrust]add interface GigabitEthernet 0/0/0
[USG6000-zone-untrust]quit
[USG6000]security-policy
[USG6000-policy-security]rule name ALLOW_LAN_TO_WAN
[USG6000-policy-security-rule-ALLOW_LAN_TO_WAN]source-zone trust
[USG6000-policy-security-rule-ALLOW_LAN_TO_WAN]destination-zone untrust
[USG6000-policy-security-rule-ALLOW_LAN_TO_WAN]action permit
[USG6000-policy-security-rule-ALLOW_LAN_TO_WAN]quit
```

Rappel fondamental USG : **sans règle de security-policy, tout est refusé
entre zones** (deny par défaut). Une interface sans zone = inutilisable.

## 19. La vue WLAN (AR avec fonction AC / contrôleurs)

```vrp
[AR720]wlan
[AR720-wlan-view]ap-group name AP-GROUP-SIEGE
[AR720-wlan-ap-group-AP-GROUP-SIEGE]quit
[AR720-wlan-view]regulatory-domain-profile name DOMAIN-FR
[AR720-wlan-view]quit
```

> Sur l'AR720, la fonction WLAN n'existe que si le modèle/licence le permet
> (à vérifier avec `display version` / `display license`).

## 20. Navigation complète : quit, return, Ctrl+Z

| Commande / touche | Effet |
|---|---|
| `quit` | Remonte d'**un** niveau (ou quitte la session en user view) |
| `return` | Retour direct à la **user view**, d'où que vous soyez |
| `Ctrl+Z` | Idem `return` (ne fonctionne que depuis le clavier, pas dans un script) |
| `Ctrl+C` | Interrompt la commande en cours (équivalent Ctrl+C Cisco) |

```vrp
[AR720-ospf-1-area-0.0.0.0]quit      # -> [AR720-ospf-1]
[AR720-ospf-1]quit                   # -> [AR720]
[AR720]interface GE 0/0/0
[AR720-GigabitEthernet0/0/0]return   # -> <AR720> directement
```

## 21. Exécuter une commande display depuis la system view : `run`

Pas besoin de faire `quit` pour un simple `display` :

```vrp
[AR720-ospf-1]run display ospf peer
[AR720-GigabitEthernet0/0/0]run display this
```

Équivalent Cisco : `do show ...` en mode configuration.

## 22. `display this` — afficher la config de la vue courante

```vrp
[AR720-GigabitEthernet0/0/0]display this
#
interface GigabitEthernet0/0/0
 description LIEN_VERS_S310
 ip address 192.168.1.1 255.255.255.0
#
return
```

Indispensable pour vérifier ce qu'on vient de taper sans relire toute la
config. Fonctionne dans **toutes** les vues.

## 23. L'aide en ligne : `?` (aide complète et partielle)

```vrp
<AR720>?
  User view commands:
    arp-ping         ARP-ping
    autosave         Autosave command group
    backup           Backup information
    ...
    display          Display information
    ...
<AR720>display ?
  ...  current-configuration  Current configuration
       cpu-usage              Cpu usage information
       ...
<AR720>display c?
  clock        System clock
  cpu-defend   CPU defend
  cpu-usage    Cpu usage information
  current-configuration  Current configuration
```

- `?` seul : toutes les commandes de la vue.
- `mot ?` : aide sur ce mot-clé.
- `début?` (collé) : tous les mots-clés commençant par ces lettres.

## 24. Raccourcis clavier — tableau de référence

| Touche | Effet |
|---|---|
| `Tab` | Complète la commande si la frappe est unique |
| `?` | Aide contextuelle |
| `↑` / `↓` | Historique des commandes (précédente / suivante) |
| `Ctrl+A` | Curseur au début de la ligne |
| `Ctrl+E` | Curseur en fin de ligne |
| `Ctrl+U` | Efface tout ce qui est à gauche du curseur — à vérifier sur votre version |
| `Ctrl+K` | Interrompt une opération en cours (ex. copie TFTP) |
| `Ctrl+W` | Efface le mot à gauche du curseur |
| `Alt+D` | Efface le mot à droite du curseur |
| `Ctrl+C` | Interrompt la commande en cours |
| `Ctrl+Z` | Retour à la user view |
| `Espace` | Page suivante (quand `--More--` s'affiche) |
| `q` ou `Ctrl+C` | Quitte l'affichage paginé |
| `Entrée` | Ligne suivante dans l'affichage paginé |

> Les combinaisons exactes peuvent varier légèrement selon la version VRP ;
> `Ctrl+A/E/Z/C` sont les plus stables. `display history-command` montre les
> 10 dernières commandes de la session.

## 25. Filtrer les sorties : `| include | begin | exclude`

Comme le `| include` de Cisco :

