---
id: collect-261001-huawei/huawei/fiche-ar720
title: "FICHE RÉFLEXE — Huawei NetEngine AR720 (routeur)"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-huawei/fiche_ar720.md
source_anchor: ""
source_lines: [1, 88]
sha256: bd85c81f134913b95384abe6ba1ed6982425399d17886756dfb69958d4944e5a
---

# FICHE RÉFLEXE — Huawei NetEngine AR720 (routeur)

> Fiche terrain courte : diagnostic et dépannage sur site. Pas de théorie.

## Caractéristiques express
- Routeur d'accès entreprise NetEngine AR720 (ports GE/SFP, slots d'extension selon version) — à vérifier selon version
- Fonctions : routage, NAT, VPN IPsec, QoS, firewall de base
- Double alimentation possible selon version — à vérifier selon version
- Console + port USB pour transfert de fichiers

## Accès
- Console : 9600 8N1 (câble console RJ45)
- Web usine : adresse IP de gestion selon version — à vérifier selon version
- Identifiants usine : admin / mot de passe usine ou défini à la première connexion — à vérifier selon version
- SSH : port 22, à activer dans la config

## Les 10 commandes qui sauvent
1. `display version` — modèle et version du firmware (VRP)
2. `display device` — état matériel (cartes, alimentations, ventilateurs)
3. `display interface brief` — état UP/DOWN et protocole de chaque interface
4. `display ip interface brief` — adresses IP et leurs états
5. `display ip routing-table` — table de routage (route manquante ?)
6. `display current-configuration` — configuration active
7. `display nat session` — sessions NAT en cours (table pleine ?)
8. `display cpu-usage` — charge CPU (anormale si > 70 % durable)
9. `display logbuffer` — derniers événements (link, erreurs)
10. `tracert 8.8.8.8` — où s'arrête le trafic (adapter la cible)

## LED : signification
| LED | Couleur | État | Signification |
|-----|---------|------|---------------|
| PWR | Vert | Fixe | Alimentation OK |
| PWR | — | Éteinte | Pas d'alimentation |
| SYS | Vert | Clignotement lent | Fonctionnement normal |
| SYS | Rouge | Fixe ou clignotant | Défaut système |
| Port (GE/SFP) | Vert | Fixe | Lien établi |
| Port (GE/SFP) | Vert | Clignotant | Trafic en cours |
| Port (GE/SFP) | — | Éteinte | Pas de lien : vérifier câble, SFP et équipement distant |

## Password recovery en 5 étapes
1. Prévenir : selon la méthode, la config peut être conservée (skip) ou perdue (reset usine).
2. Se connecter en console (9600 8N1), redémarrer le routeur.
3. Pendant le boot, appuyer sur Ctrl+B pour entrer dans le menu BootLoad.
4. Choisir l'option « Password recovery » / « Skip current system configuration » selon le menu.
5. Redéfinir le mot de passe admin, sauvegarder (`save`), redémarrer normalement et contrôler.

> Procédure exacte (touche, libellé du menu) : voir guide complet.

## Upgrade firmware en 5 étapes
1. Sauvegarder la config (`save` + copie du fichier `vrpcfg.cfg` via TFTP/FTP) et noter la version (`display version`).
2. Télécharger le firmware correspondant EXACTEMENT au modèle (AR720) sur le support Huawei.
3. Transférer le fichier .cc vers le routeur (TFTP/FTP/SFTP ou clé USB).
4. Définir l'image de démarrage : `startup system-software <fichier>.cc` puis vérifier avec `display startup`.
5. Redémarrer (`reboot`), contrôler `display version`, tester le routage et le NAT.

## Panne → réflexe
- Aucune LED → vérifier l'alimentation (les deux si redondante) et le câble secteur.
- SYS rouge → `display logbuffer` puis `display device` ; noter l'heure du défaut.
- Pas d'Internet → `display ip interface brief` (WAN UP ?) puis `tracert` vers l'extérieur ; vérifier la route par défaut (`display ip routing-table`).
- NAT saturé → `display nat session` ; vérifier le pool d'adresses et les timeouts.
- Un site VPN down → vérifier la phase IPsec (logs) et la connectivité de base avant tout.
- Interface UP mais pas de trafic → `display interface <nom>` (erreurs CRC ?) ; changer de câble/SFP.

## NAT sortant express (à adapter)
```
system-view
acl 2000
rule 5 permit source 192.168.1.0 0.0.0.255
quit
interface GigabitEthernet0/0/1
nat outbound 2000
quit
save
```
- Vérifier : `display nat outbound` et `display nat session`.

## Numéros utiles
- TAC Huawei : [À COMPLÉTER]
- Partenaire : [À COMPLÉTER]
- Responsable : [À COMPLÉTER]

## Avant de quitter le site (checklist)
- [ ] Configuration sauvegardée (`save` + vérification après reboot)
- [ ] Version du firmware notée
- [ ] Test utilisateur final OK (Internet, VPN, applications métier)
- [ ] Routes et NAT vérifiés
- [ ] Câbles WAN/LAN étiquetés
- [ ] Photos des étiquettes (SN / MAC) prises
