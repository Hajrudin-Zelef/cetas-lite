---
id: collect-261001-general-networking/general-networking/fiche-s310
title: "FICHE RÉFLEXE — Huawei S310 (switch eKit)"
domain: general-networking
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fiche_s310.md
source_anchor: ""
source_lines: [1, 92]
sha256: 8e456f7078659ec3358e047fd380302f0e3a915a66f21f915d01740b70cbc306
---

# FICHE RÉFLEXE — Huawei S310 (switch eKit)

> Fiche terrain courte : diagnostic et dépannage sur site. Pas de théorie.

## Caractéristiques express
- Switch manageable eKit (ex. S310-24T4S / S310-24P4S selon version) — à vérifier selon version
- Ports RJ45 Gigabit + uplinks SFP/SFP+ selon modèle — à vérifier selon version
- Modèles PoE disponibles (budget PoE à noter sur site) — à vérifier selon version
- Empilable / administrable en web local (mode eKit)

## Accès
- Console : 9600 8N1 (câble console RJ45)
- Web usine : http://192.168.1.253 — à vérifier selon version
- Identifiants usine : admin / mot de passe défini à la première connexion — à vérifier selon version
- SSH/Telnet : à activer dans la config ; Telnet déconseillé (non chiffré)

## Les 10 commandes qui sauvent
1. `display version` — modèle et version du firmware (VRP)
2. `display device` — état des cartes/ports et alimentations
3. `display interface brief` — état UP/DOWN de tous les ports
4. `display current-configuration` — configuration active
5. `display ip interface brief` — adresses IP des interfaces VLAN
6. `display mac-address` — table MAC (pour localiser un équipement)
7. `display stp brief` — état Spanning Tree (boucle si port en discarding anormal)
8. `display cpu-usage` — charge CPU (anormale si > 70 % durable)
9. `display logbuffer` — derniers événements (link up/down, erreurs)
10. `ping -c 5 192.168.1.1` — test de connectivité (adapter l'IP)

## LED : signification
| LED | Couleur | État | Signification |
|-----|---------|------|---------------|
| PWR | Vert | Fixe | Alimentation OK |
| PWR | — | Éteinte | Pas d'alimentation |
| SYS | Vert | Clignotement lent | Fonctionnement normal |
| SYS | Rouge | Fixe ou clignotant | Défaut système |
| Port (link/act) | Vert | Fixe | Lien établi |
| Port (link/act) | Vert | Clignotant | Trafic en cours |
| Port (link/act) | — | Éteinte | Pas de lien : vérifier câble et équipement distant |

## Password recovery en 5 étapes
1. Prévenir : selon la méthode, la config peut être conservée (skip) ou perdue (reset usine).
2. Se connecter en console (9600 8N1), redémarrer le switch.
3. Pendant le boot, appuyer sur Ctrl+B pour entrer dans le menu BootROM/BootLoad.
4. Choisir l'option « Skip current system configuration » (ou équivalent) pour démarrer sans la config.
5. Redéfinir le mot de passe admin, sauvegarder (`save`), redémarrer normalement et contrôler.

> Procédure exacte (touche, libellé du menu) : voir guide complet.

## Upgrade firmware en 5 étapes
1. Sauvegarder la config (`save` + copie du fichier `vrpcfg.cfg` via TFTP/FTP) et noter la version (`display version`).
2. Télécharger le firmware correspondant EXACTEMENT au modèle (S310-…) sur le support Huawei.
3. Transférer le fichier .cc vers le switch (TFTP/FTP/SFTP).
4. Définir l'image de démarrage : `startup system-software <fichier>.cc` puis vérifier avec `display startup`.
5. Redémarrer (`reboot`), contrôler `display version`, tester les ports et les VLAN.

## Panne → réflexe
- Aucune LED → vérifier l'alimentation et le câble secteur ; tester une autre prise.
- SYS rouge → `display logbuffer` puis `display device` ; noter l'heure du défaut.
- Un port DOWN → `display interface <port>` ; changer de câble ; tester le port sur un autre équipement ; vérifier la négociation.
- Boucle réseau (tempête broadcast) → `display stp brief` ; chercher le port en forwarding anormal ; débrancher un à un.
- Un VLAN ne passe pas → `display current-configuration` (vérifier `port trunk allow-pass vlan`) ; contrôler le PVID du port d'accès.
- PoE absent sur un port → `display power` (ou équivalent) ; vérifier le budget PoE total et la classe de l'équipement.

## VLAN express (à adapter)
```
system-view
vlan 10
quit
interface GigabitEthernet0/0/5
port link-type access
port default vlan 10
quit
interface GigabitEthernet0/0/24
port link-type trunk
port trunk allow-pass vlan 10 20
quit
save
```
- Vérifier : `display vlan` et `display port vlan`.

## Numéros utiles
- TAC Huawei : [À COMPLÉTER]
- Partenaire : [À COMPLÉTER]
- Responsable : [À COMPLÉTER]

## Avant de quitter le site (checklist)
- [ ] Configuration sauvegardée (`save` + vérification après reboot)
- [ ] Version du firmware notée
- [ ] Test utilisateur final OK (chaque VLAN testé, un port par usage)
- [ ] Plan de ports / étiquetage à jour
- [ ] Budget PoE vérifié si équipements alimentés
- [ ] Photos des étiquettes (SN / MAC) prises
