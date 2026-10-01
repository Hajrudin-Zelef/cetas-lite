---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-2
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [159, 345]
sha256: bd9b6bb1533e4a885b4f3890770040ad2755bc6148abcea2c7d949f4130205f6
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

- Alimentation AC intégrée unique, **non redondante**. Sur un site critique, on protège en
  amont : onduleur (voir le guide onduleurs/UPS) + parafoudre. Le routeur ne fait que 33 W
  max : un petit UPS 1 kVA tient largement.
- Plage de tension large (90–264 V) : tolère bien les réseaux instables, mais pas les
  micro-coupures — d'où l'onduleur.
- Température 0–45 °C : dans une armoire technique sans clim en zone chaude, ça se surveille.
  Prévoir une sonde ou au minimum la supervision SNMP de température si disponible.
- Humidité 5–95 % sans condensation : éviter les locaux avec climatisation qui condense
  directement sur l'équipement.
- Flux d'air gauche → droite : dans un rack, ne pas mettre un équipement qui souffle chaud
  directement sur l'entrée d'air gauche de l'AR720.

## 7. La console : câble, paramètres, premier contact

Le port console est un RJ45 série. Il faut un câble console RJ45 ↔ DB9 (ou USB-série selon
le PC). Paramètres du terminal :

| Paramètre | Valeur |
|---|---|
| Débit | 9600 bauds |
| Bits de données | 8 |
| Parité | Aucune |
| Bits de stop | 1 |
| Contrôle de flux | Aucun |

Outils : PuTTY (Windows), `screen /dev/ttyUSB0 9600` ou `minicom` (Linux).

Au premier boot, VRP demande la configuration initiale (assistant). **Ne jamais laisser le
routeur avec le mot de passe par défaut ou sans mot de passe**, même 5 minutes : en DHCP
client sur le WAN, il peut récupérer une IP publique et être scanné dans la minute.

## 8. Architecture logicielle : VRP en deux mots

VRP (Versatile Routing Platform) est l'OS de Huawei, équivalent d'IOS chez Cisco ou JunOS
chez Juniper. Concepts de base :

- **Vues (views)** : la CLI est hiérarchique. `<AR720>` = vue utilisateur (user view),
  `[AR720]` = vue système (system view), puis des sous-vues par fonction
  (`[AR720-GigabitEthernet0/0/0]`, `[AR720-ike-proposal-1]`, etc.).
- `system-view` : passer en vue système (depuis `<AR720>`, taper `system-view` ou `sys`).
- `quit` : remonter d'un niveau. `return` : revenir directement en vue utilisateur
  (raccourci `Ctrl+Z`).
- `display` : commandes d'affichage (équivalent `show` chez Cisco). Ex. `display
  current-configuration`.
- `?` : aide contextuelle partout. `Tab` : complétion.
- La configuration se fait en **mode candidat implicite** : les commandes prennent effet
  immédiatement mais ne sont **pas sauvegardées** tant qu'on n'a pas fait `save`.
  Oublier `save` = perdre la config au reboot. C'est l'erreur n°1 des débutants.

## 9. Premier boot : procédure pas à pas

1. Brancher le câble console, ouvrir le terminal (9600 8N1).
2. Brancher l'alimentation. Observer les LED : PWR vert, SYS clignote puis se stabilise.
3. À l'invite, répondre à l'assistant de configuration initiale **ou** le quitter pour
   configurer à la main (recommandé : plus de contrôle).
4. Définir le nom d'hôte :
   ```
   <AR720>system-view
   [AR720]sysname AGENCE-DAKAR-AR720
   ```
5. Configurer le mot de passe console + activer l'authentification.
6. Configurer l'utilisateur admin pour SSH/telnet (voir section 14).
7. Configurer la date/heure (`clock timezone`, `clock datetime` ou NTP — voir Partie I).
8. Configurer le WAN (Partie B), le LAN (Partie C), le NAT (Partie D).
9. **`save`** puis vérifier avec `display saved-configuration` vs `display
   current-configuration`.

## 10. Sécuriser l'accès console et le mot de passe enable

```
[AGENCE-DAKAR-AR720]user-interface console 0
[AGENCE-DAKAR-AR720-ui-console0]authentication-mode password
[AGENCE-DAKAR-AR720-ui-console0]set authentication password cipher MotDePasseConsoleFictif123
[AGENCE-DAKAR-AR720-ui-console0]idle-timeout 10 0
[AGENCE-DAKAR-AR720-ui-console0]quit
```

- `idle-timeout 10 0` : déconnexion après 10 minutes d'inactivité (minutes, secondes).
- Toujours `cipher` (mot de passe chiffré dans la config), jamais en clair.
- Mot de passe = fictif dans ce guide, évidemment. En production : long, unique par site,
  stocké dans le coffre de l'équipe.

## 11. Créer l'utilisateur admin (SSH)

```
[AGENCE-DAKAR-AR720]aaa
[AGENCE-DAKAR-AR720-aaa]local-user admin password cipher MotDePasseAdminFictif456
[AGENCE-DAKAR-AR720-aaa]local-user admin privilege level 15
[AGENCE-DAKAR-AR720-aaa]local-user admin service-type ssh telnet terminal
[AGENCE-DAKAR-AR720-aaa]quit
[AGENCE-DAKAR-AR720]stelnet server enable
[AGENCE-DAKAR-AR720]ssh user admin authentication-type password
[AGENCE-DAKAR-AR720]user-interface vty 0 4
[AGENCE-DAKAR-AR720-ui-vty0-4]authentication-mode aaa
[AGENCE-DAKAR-AR720-ui-vty0-4]protocol inbound ssh
[AGENCE-DAKAR-AR720-ui-vty0-4]idle-timeout 10 0
[AGENCE-DAKAR-AR720-ui-vty0-4]quit
```

Note : `telnet` dans `service-type` est montré pour compatibilité, mais **désactiver telnet
en production** (voir Partie M). SSH uniquement.

## 12. Nom d'hôte, fuseau horaire, horloge

```
[AGENCE-DAKAR-AR720]sysname AGENCE-DAKAR-AR720
[AGENCE-DAKAR-AR720]clock timezone Dakar add 00:00
[AGENCE-DAKAR-AR720]clock datetime 10:30:00 2026-09-27
```

Le fuseau horaire s'écrit avec un nom libre + décalage. Pour une synchro fiable, configurer
NTP (voir section 70). Une horloge fausse = logs inexploitables, certificats VPN qui
échouent, dépannage impossible. C'est un détail qui coûte cher.

## 13. Les fichiers du système : vrpcfg.zip, firmware, USB

En vue utilisateur :

```
<AGENCE-DAKAR-AR720>dir
<AGENCE-DAKAR-AR720>display version
<AGENCE-DAKAR-AR720>display startup
```

- `display version` : version VRP, uptime, modèle, mémoire. **Première commande à donner
  quand on ouvre un ticket** ou qu'on demande de l'aide.
- `display startup` : quel fichier de config et quel firmware sont chargés au boot.
- `dir` : liste les fichiers (config `vrpcfg.zip`, image système `.cc`).
- Les 2 ports USB servent à : mettre à jour le firmware depuis une clé, sauvegarder la
  config, ou booter en secours si la flash est corrompue.

Convention de nommage des firmwares Huawei : `AR720-V300R024C00SPC100.cc` (format
indicatif — **à vérifier sur la fiche du modèle exact** et sur le portail de téléchargement
avec le numéro de série).

## 14. Upgrade initiale : mettre le firmware à niveau avant la mise en service

Règle : **on ne met jamais en production un équipement avec le firmware d'usine** sans
vérifier qu'il n'y a pas une version corrective connue. Procédure type :

1. Télécharger la version cible depuis le portail Huawei (compte + contrat de support
   requis en général).
2. Copier sur clé USB (FAT32), insérer dans le routeur.
3. En vue utilisateur :
   ```
   <AGENCE-DAKAR-AR720>dir usb0:
   <AGENCE-DAKAR-AR720>copy usb0:/AR720-V300R0XXCXX.cc flash:/
   ```
4. Déclarer la nouvelle image comme image de démarrage :
   ```
   <AGENCE-DAKAR-AR720>startup system-software AR720-V300R0XXCXX.cc
   ```
5. Vérifier : `display startup`.
6. **`save`** la configuration actuelle avant tout reboot.
7. `reboot` (confirmer). Observer le boot en console.
8. Après reboot : `display version` pour confirmer, puis re-tester WAN/LAN/VPN.

**Rollback** : garder l'ancienne image en flash (il y a 512 Mo utilisables, ça passe en
général pour 2 images). Si la nouvelle version pose problème :
`startup system-software <ancienne-image>` + `reboot`. Ne jamais effacer l'ancienne image
avant d'avoir validé la nouvelle en production pendant au moins 48 h.

## 15. Sauvegarder la configuration (la base de tout)

```
<AGENCE-DAKAR-AR720>save
```

Répondre `y` aux deux questions (écraser le fichier de config). Vérification :

```
<AGENCE-DAKAR-AR720>display saved-configuration
<AGENCE-DAKAR-AR720>display current-configuration
```

Les deux doivent être identiques après un `save`. Pour une sauvegarde externe :

```
<AGENCE-DAKAR-AR720>copy vrpcfg.zip usb0:/backup-AGENCE-DAKAR-2026-09-27.zip
```

Ou via TFTP/FTP vers un serveur (voir Partie J). Fréquence minimale : après chaque
changement + sauvegarde hebdomadaire automatisée.

---
---

