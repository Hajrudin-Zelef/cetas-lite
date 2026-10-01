---
id: collect-261001-huawei/huawei/fiche-ap761
title: "FICHE RÉFLEXE — Huawei AP761 (AP Wi-Fi 6 d'EXTÉRIEUR, IP68)"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "memory"]
source: docs/RAG/collect-261001-huawei/fiche_ap761.md
source_anchor: ""
source_lines: [1, 81]
sha256: 48aed54656df01bee731106da2985a6aab32edf7c47bfd4a0342a19f58634015
---

# FICHE RÉFLEXE — Huawei AP761 (AP Wi-Fi 6 d'EXTÉRIEUR, IP68)

> Fiche terrain courte : diagnostic et dépannage sur site. Pas de théorie.
> ⚠️ L'AP761 est un AP Wi-Fi 6 (802.11ax) d'extérieur — PAS du Wi-Fi 7.

## Caractéristiques express
- Wi-Fi 6 (802.11ax), double bande 2,4 / 5 GHz, usage extérieur
- Indice de protection IP68 : étanche à la poussière et à l'immersion — ne jamais ouvrir le boîtier sur site
- Alimentation PoE (802.3at conseillé, vérifier la classe PoE requise) — à vérifier selon version
- Température de fonctionnement étendue (plage exacte : voir datasheet) — à vérifier selon version
- Montage : mât ou mur ; terre/parafoudre obligatoire en extérieur
- Modes : Fat AP (autonome) ou Fit AP (contrôleur/AC)

## Accès
- Console : 9600 8N1 (port console protégé par bouchon étanche — refermer après usage)
- Web usine : http://169.254.1.1 — à vérifier selon version
- Identifiants usine : admin / mot de passe défini à la première connexion — à vérifier selon version
- SSH : port 22, à activer dans la config

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
| SYS | — | Éteinte | Pas d'alimentation : vérifier PoE, parafoudre et câble |

## Password recovery en 5 étapes
1. Prévenir : la procédure réinitialise l'appareil en configuration usine (config perdue).
2. Appareil sous tension, maintenir le bouton RESET > 5 s jusqu'au redémarrage (accès via trappe étanche — refermer après).
3. Laisser redémarrer complètement (LED SYS verte, clignotement lent).
4. Se connecter au web usine (http://169.254.1.1 — à vérifier selon version) et définir un nouveau mot de passe admin.
5. Restaurer la configuration sauvegardée, contrôler avec `display current-configuration`.

> Procédure exacte (timing, emplacement du bouton) : voir guide complet.

## Upgrade firmware en 5 étapes
1. Sauvegarder la config (`save` ou export web) et noter la version actuelle (`display version`).
2. Télécharger le firmware correspondant EXACTEMENT au modèle (AP761) sur le support Huawei.
3. Transférer le fichier : page web « Maintenance > Mise à jour » ou TFTP/FTP selon version.
4. Vérifier le fichier chargé et le définir comme image de démarrage (`display startup`).
5. Redémarrer, contrôler `display version`, tester l'association d'un client Wi-Fi.

## Panne → réflexe
- Pas de LED → vérifier l'injecteur/switch PoE, le parafoudre et le câble (oxydation des connecteurs en extérieur).
- LED rouge → `display logbuffer` puis `display device` ; noter l'heure du défaut.
- Clients non associés → `display interface brief` ; vérifier SSID, mot de passe et bande 2,4/5 GHz.
- Débit faible → `display cpu-usage` ; vérifier la négociation du port ; contrôler l'orientation et les obstacles.
- Web inaccessible → ping de l'IP usine ; vérifier le câble et mettre le PC dans le même sous-réseau.
- Reboots en boucle → `display logbuffer` ; suspecter firmware corrompu ou alimentation PoE instable → ré-upgrade / changer d'injecteur.

## Points d'attention extérieur
- Ne jamais laisser la trappe console ouverte : perte de l'étanchéité IP68.
- Vérifier la continuité de la terre et l'état du parafoudre après chaque orage.
- Contrôler le serrage du collier de mât et l'état des presse-étoupes.

## Numéros utiles
- TAC Huawei : [À COMPLÉTER]
- Partenaire : [À COMPLÉTER]
- Responsable : [À COMPLÉTER]

## Avant de quitter le site (checklist)
- [ ] Configuration sauvegardée (`save` + vérification après reboot)
- [ ] Version du firmware notée
- [ ] Test utilisateur final OK (association, débit)
- [ ] Trappe console refermée (étanchéité OK)
- [ ] Terre et parafoudre contrôlés
- [ ] Fixation mât/mur vérifiée (serrage)
- [ ] Photos des étiquettes (SN / MAC) prises
