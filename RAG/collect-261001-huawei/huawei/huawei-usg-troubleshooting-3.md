---
id: collect-261001-huawei/huawei/huawei-usg-troubleshooting-3
title: "Huawei USG — Guide de dépannage (troubleshooting récurrent)"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-huawei/huawei_usg_troubleshooting.md
source_anchor: ""
source_lines: [438, 509]
sha256: 47eb081678628190704e4a433e33b2366f459e277d2ecbd4b02f75ca12b04012
---

# Huawei USG — Guide de dépannage (troubleshooting récurrent)

1. Ne paniquez pas : le USG garde l'ancienne image (`display startup`
   montre les images primaire/secours).
2. En BootWare (Ctrl+B au démarrage) : choisissez l'image précédente.
3. Causes d'échec : image corrompue (revérifiez le hash/MD5 après TFTP),
   espace flash insuffisant (`dir` : nettoyez les vieux fichiers),
   coupure de courant pendant l'écriture (d'où l'onduleur).
4. Après upgrade : `display version` + `display device` pour valider,
   puis `save`.

---

## 19. Mot de passe perdu

- **Console** : au démarrage, Ctrl+B → menu BootWare →
  « Clear password for console user ».
- **Web/SSH** : via un compte admin valide restant, ou en console :
  `aaa` → `local-user <nom> password irreversible-cipher <nouveau>`.
- Dernier recours : `reset saved-configuration` + reboot = retour
  usine (perte de config ! sauvegardez avant toute manip risquée).

---

## 20. Fiche réflexe : la panne Internet en 5 minutes

```shell
# 1. Interfaces
display interface brief
# 2. Passerelle vue ?
ping -a <ip-lan-usg> <ip-pc>
# 3. Route par défaut ?
display ip routing-table 0.0.0.0
# 4. NAT ?
display nat session all
# 5. Politique ?
display security-policy rule all
display firewall session table source-ip <ip-pc>
# 6. DNS ?
nslookup www.google.com
# 7. WAN ?
ping <passerelle-fai>
display pppoe-client session summary   # si PPPoE
```

Si les 7 points sont verts et que ça ne marche toujours pas :
regardez les **logs** (`display logbuffer`) et désactivez
temporairement l'UTM sur la règle pour isoler.

---

## 21. Commandes à connaître par cœur

| Besoin | Commande |
|---|---|
| Sessions actives | `display firewall session table` |
| Sessions NAT | `display nat session all` |
| Règles de sécurité | `display security-policy rule all` |
| Règles NAT | `display nat-policy rule all` |
| Table de routage | `display ip routing-table` |
| Voisins OSPF/BGP | `display ospf peer` / `display bgp peer` |
| Tunnels VPN | `display ike sa` / `display ipsec sa` |
| État HA | `display hrp state verbose` |
| Charge | `display cpu-usage` / `display memory-usage` |
| Logs | `display logbuffer` / `display trapbuffer` |
| Nettoyer sessions | `reset firewall session table` |
| Nettoyer NAT | `reset nat session all` |
| Debug (temporaire !) | `terminal debugging` + `debugging ...` puis `undo debugging all` |

---

*Fin du guide — ~550 lignes. Règle d'or : 80 % des pannes USG =
security-policy ou NAT. Vérifiez toujours dans cet ordre :
interface → zone → politique → NAT → routage → UTM.*
