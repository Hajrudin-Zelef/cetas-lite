---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-13
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [2542, 2606]
sha256: b8668b5b7f686735b39ad83cf1215f4f0f7fa6dc30235a8ac845be22f7068bac
---

# VRP — Le système d'exploitation transversal Huawei

1. Taper `show` au lieu de `display`.
2. Oublier `save`.
3. Configurer l'OSPF `network` dans la mauvaise vue.
4. Mettre le masque en décimal pointé (`255.255.255.0`) au lieu du CIDR.
5. Laisser Telnet activé en production.
6. Oublier `undo shutdown` (ou laisser `shutdown` par copier-coller).
7. Ne pas vérifier `display startup` avant un reboot d'upgrade.
8. Supprimer l'ancien `.cc` trop tôt après un upgrade.
9. Configurer une interface USG sans zone de sécurité.
10. Faire un `debugging` en production sans fenêtre de maintenance.

## 132. Modèle de fiche d'intervention (à copier)

```text
FICHE D'INTERVENTION VRP
Date/heure : ............    Intervenant : ............
Équipement : ............    Version VRP : ............
Motif : ....................................................
État avant (display version / display startup / display interface brief) :
........................................................................
Sauvegarde config : [ ] locale (save)  [ ] externe (fichier : ............)
Commandes prévues :
  1. ................................................................
  2. ................................................................
Rollback prévu : .....................................................
Vérifications après :
  [ ] display current-configuration (échantillon)
  [ ] save + display saved-configuration
  [ ] Tests : ping ............  display ospf peer ............
  [ ] Supervision OK (Zabbix/syslog)
Journal des changements mis à jour : [ ] oui
```

## 133. Commandes d'urgence — à connaître par cœur

```vrp
<AR720>reboot                                        # redémarrage
<AR720>display diagnostic-information               # tout-en-un pour le support
<AR720>undo debugging all                           # couper les debugs
[AR720-GE0/0/0]shutdown / undo shutdown             # couper/rétablir un lien
<AR720>reset saved-configuration                    # reset usine (puis reboot)
Ctrl+B au boot                                      # BootROM (password recovery)
```

## 134. Note sur les valeurs « à vérifier »

Ce guide privilégie la syntaxe stable entre versions VRP. Quand une
commande varie selon le modèle ou la version (comportement observé sur
certaines déclinaisons uniquement), elle est marquée **« à vérifier »** :
testez-la sur VOTRE version exacte (`display version`) ou dans la Command
Reference EDOC du produit avant de l'utiliser en production. En cas de
doute, `?` dans la vue concernée tranche toujours.

## 135. Conclusion — l'état d'esprit VRP

VRP récompense la méthode : `display` avant d'agir, `display this` après
chaque vue, `compare configuration` avant `save`, `display startup` avant
`reboot`. Trois boîtiers différents (AR720, S310, USG6000), un seul réflexe.
Avec ce socle, chaque guide spécifique (routage, switching, sécurité)
devient une simple couche de vocabulaire métier posée sur les mêmes fondations.

---

*Fin du guide — VRP, le système d'exploitation transversal Huawei.*
*Bon courage, et n'oubliez jamais : `save` !*
