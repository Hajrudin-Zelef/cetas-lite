---
id: collect-261001-huawei/huawei/fiche-ap361
title: "FICHE RÉFLEXE — Huawei AP361 (eKit, Wi-Fi 6 d'intérieur, PoE 802.3af)"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-huawei/fiche_ap361.md
source_anchor: ""
source_lines: [1, 80]
sha256: 06a8f16502486fbab8c6d8279b423d9e5d40a500e4fafacaea1c9764d1700ca2
---

# FICHE RÉFLEXE — Huawei AP361 (eKit, Wi-Fi 6 d'intérieur, PoE 802.3af)

> Fiche terrain courte : diagnostic et dépannage sur site. Pas de théorie.

## Caractéristiques express
- Wi-Fi 6 (802.11ax), double bande 2,4 / 5 GHz, 2x2 MIMO
- Alimentation PoE 802.3af (consommation max ~8,8 W) ou adaptateur 12 V — à vérifier selon version
- Modes : Fat AP (autonome, géré en web local) ou Fit AP (géré par contrôleur/AC)
- Montage : plafond ou mur, kit fourni

## Accès
- Console : 9600 8N1 (port console RJ45 ou micro-USB selon version)
- Web usine : http://169.254.1.1 — à vérifier selon version
- Identifiants usine : admin / mot de passe défini à la première connexion — à vérifier selon version
- SSH : port 22, à activer dans la config (souvent désactivé en usine)

## Les 10 commandes qui sauvent
1. `display version` — modèle et version du firmware
2. `display device` — état matériel (radios, ports)
3. `display interface brief` — état UP/DOWN de chaque interface
4. `display ip interface brief` — adresses IP et leurs états
5. `display current-configuration` — configuration active
6. `display logbuffer` — derniers événements (erreurs, reboots)
7. `display cpu-usage` — charge CPU (anormale si > 70 % durable)
8. `display memory-usage` — mémoire utilisée
9. `display users` — qui est connecté sur l'équipement en ce moment
10. `ping -c 5 192.168.1.1` — test de connectivité vers la passerelle (adapter l'IP)

## LED : signification
| LED | Couleur | État | Signification |
|-----|---------|------|---------------|
| SYS | Vert | Clignotement lent | Fonctionnement normal |
| SYS | Vert | Clignotement rapide | Démarrage, mise à jour ou association en cours |
| SYS | Rouge | Fixe ou clignotant | Défaut matériel ou logiciel |
| SYS | — | Éteinte | Pas d'alimentation : vérifier PoE et câble |

## Password recovery en 5 étapes
1. Prévenir : la procédure réinitialise l'appareil en configuration usine (config perdue).
2. Appareil sous tension, maintenir le bouton RESET > 5 s jusqu'au redémarrage.
3. Laisser redémarrer complètement (LED SYS verte, clignotement lent).
4. Se connecter au web usine (http://169.254.1.1 — à vérifier selon version) et définir un nouveau mot de passe admin.
5. Restaurer la configuration sauvegardée, contrôler avec `display current-configuration`.

> Procédure exacte (timing, emplacement du bouton) : voir guide complet.

## Upgrade firmware en 5 étapes
1. Sauvegarder la config (`save` ou export web) et noter la version actuelle (`display version`).
2. Télécharger le firmware correspondant EXACTEMENT au modèle (AP361) sur le support Huawei.
3. Transférer le fichier : page web « Maintenance > Mise à jour » ou TFTP/FTP selon version.
4. Vérifier le fichier chargé et le définir comme image de démarrage (`display startup`).
5. Redémarrer, contrôler `display version`, tester l'association d'un client Wi-Fi.

## Panne → réflexe
- Pas de LED → vérifier l'injecteur/switch PoE, changer de câble, tester un autre port PoE.
- LED rouge → `display logbuffer` puis `display device` ; noter l'heure du défaut.
- Clients non associés → `display interface brief` ; vérifier SSID, mot de passe et bande 2,4/5 GHz.
- Débit faible → `display cpu-usage` ; vérifier la négociation du port (1 Gbps) ; éloigner les sources d'interférences.
- Web inaccessible → ping de l'IP usine ; vérifier le câble et mettre le PC dans le même sous-réseau.
- Reboots en boucle → `display logbuffer` ; suspecter un firmware corrompu → ré-upgrade.

## Câblage et PoE : réflexe
- Câble : Cat5e minimum, 100 m max entre le switch/injecteur et l'AP.
- L'AP361 est alimenté en PoE 802.3af (consommation max ~8,8 W) : un port PoE standard suffit.
- Si l'AP ne démarre pas : tester avec un injecteur PoE connu bon avant d'incriminer l'AP.
- Ne pas alimenter en PoE + adaptateur 12 V en même temps (selon version, vérifier le manuel).

## Numéros utiles
- TAC Huawei : [À COMPLÉTER]
- Partenaire : [À COMPLÉTER]
- Responsable : [À COMPLÉTER]

## Avant de quitter le site (checklist)
- [ ] Configuration sauvegardée (`save` + vérification après reboot)
- [ ] Version du firmware notée
- [ ] Test utilisateur final OK (association, débit, roaming si multi-AP)
- [ ] SSID et mot de passe remis au responsable
- [ ] Photos des étiquettes (SN / MAC) prises
- [ ] Câbles et fixation vérifiés
- [ ] Emplacement de l'AP noté sur le plan du site
- [ ] Mot de passe admin remis au responsable (enveloppe scellée)
