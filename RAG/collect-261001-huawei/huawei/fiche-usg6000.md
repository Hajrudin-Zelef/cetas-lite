---
id: collect-261001-huawei/huawei/fiche-usg6000
title: "FICHE RÉFLEXE — Huawei USG6000 (firewall)"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/fiche_usg6000.md
source_anchor: ""
source_lines: [1, 81]
sha256: 02a69d8c35e5d167ecff3e6afe47d41eebdf20f0f5fa5a49600191b9d42b73ac
---

# FICHE RÉFLEXE — Huawei USG6000 (firewall)

> Fiche terrain courte : diagnostic et dépannage sur site. Pas de théorie.

## Caractéristiques express
- Firewall nouvelle génération USG6000 (série : modèle exact à noter sur site) — à vérifier selon version
- Fonctions : security-policy, NAT, VPN IPsec/SSL, antivirus/IPS/URL filtering (licences)
- Ports GE/SFP ; interfaces en zone de sécurité (trust / untrust / dmz)
- Toujours intervenir en heures creuses : toute erreur coupe le trafic

## Accès
- Console : 9600 8N1 (câble console RJ45)
- Web usine : https://192.168.0.1:8443 — à vérifier selon version
- Identifiants usine : admin / mot de passe usine ou défini à la première connexion — à vérifier selon version
- SSH : port 22, à activer dans la config

## Les 10 commandes qui sauvent
1. `display version` — modèle et version du firmware (VRP)
2. `display device` — état matériel (cartes, alimentations, ventilateurs)
3. `display interface brief` — état UP/DOWN de chaque interface
4. `display current-configuration` — configuration active
5. `display security-policy rule all` — règles de sécurité (la règle qui bloque ?)
6. `display firewall session table` — sessions en cours (table pleine ?)
7. `display ip routing-table` — table de routage
8. `display cpu-usage` — charge CPU (anormale si > 70 % durable)
9. `display logbuffer` — derniers événements (deny, erreurs)
10. `ping -a <ip_source> 8.8.8.8` — test de connectivité depuis une zone précise (adapter)

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
1. Prévenir : toute manipulation en BootROM comporte un risque ; config à sauvegarder avant si possible.
2. Se connecter en console (9600 8N1), redémarrer le firewall.
3. Pendant le boot, appuyer sur Ctrl+B pour entrer dans le menu BootROM.
4. Choisir l'option de récupération / démarrage sans la configuration courante selon le menu.
5. Redéfinir le mot de passe admin, sauvegarder (`save`), redémarrer normalement et contrôler.

> Procédure exacte (touche, libellé du menu) : voir guide complet.

## Upgrade firmware en 5 étapes
1. Sauvegarder la config (`save` + export du fichier de config via web/FTP) et noter la version (`display version`) + les licences.
2. Télécharger le firmware correspondant EXACTEMENT au modèle (USG60xx) sur le support Huawei.
3. Transférer le fichier .cc (web « System > Upgrade » ou FTP/SFTP).
4. Définir l'image de démarrage : `startup system-software <fichier>.cc` puis vérifier avec `display startup`.
5. Redémarrer (`reboot`) en heures creuses, contrôler `display version`, tester chaque flux métier (Internet, VPN, serveurs).

## Panne → réflexe
- Aucune LED → vérifier l'alimentation (les deux si redondante) et le câble secteur.
- SYS rouge → `display logbuffer` puis `display device` ; noter l'heure du défaut.
- Trafic bloqué après un changement → `display security-policy rule all` ; vérifier la règle et la zone source/destination.
- Sessions qui ne passent plus → `display firewall session table` (table pleine ?) ; `display cpu-usage`.
- VPN IPsec down → vérifier la connectivité de base d'abord, puis les logs de phase 1/2.
- Web d'admin inaccessible → ping https://192.168.0.1:8443 ; vérifier le câble et l'IP du PC (même sous-réseau).

## Zones et policy : réflexe
- Chaque interface appartient à une zone : `display zone` pour vérifier.
- Une règle de policy se lit toujours : zone source → zone destination, service, action (permit/deny).
- En dépannage : `display security-policy rule all` puis tester avec un `ping` ciblé.
- Règle d'or : ne jamais laisser une règle « any → any : permit » en production.

## Numéros utiles
- TAC Huawei : [À COMPLÉTER]
- Partenaire : [À COMPLÉTER]
- Responsable : [À COMPLÉTER]

## Avant de quitter le site (checklist)
- [ ] Configuration sauvegardée (`save` + vérification après reboot)
- [ ] Version du firmware et licences notées
- [ ] Test utilisateur final OK (Internet, VPN, chaque flux métier)
- [ ] Règles de sécurité relues (pas de règle « any-any » oubliée)
- [ ] Câbles WAN/LAN/DMZ étiquetés
- [ ] Photos des étiquettes (SN / MAC) prises
