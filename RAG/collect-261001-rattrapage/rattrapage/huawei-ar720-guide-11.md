---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-11
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [1896, 2079]
sha256: 11fc6af9dfdd03b7c065ea385096c391ac99b32dcf1c22b2200198c879804a32
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

1. `dhcp enable` désactivé (oublié après un `undo` ou une restauration partielle).
2. Pool épuisé : `used` = taille du pool → élargir le réseau ou raccourcir le bail.
3. `dhcp select global` oublié sur l'interface (ou `dhcp select relay` mal pointé).
4. Conflit : une IP du pool configurée en statique sur un équipement → `excluded-ip-address`
   + réservation propre (section 28).

## 97. Cas n°9 — Boucle réseau : tout rame d'un coup

**Symptômes :** latence énorme, pertes, LED des ports qui clignotent frénétiquement,
parfois le routeur inaccessible.

**Diagnostic :** regarder les LED, `display interface brief` (compteurs d'erreurs qui
explosent), `display cpu-usage`.

**Actions immédiates :**

1. Débrancher les câbles suspects un par un jusqu'à ce que ça revienne.
2. Chercher le petit switch de bureau rebouclé ou le câble branché deux fois.
3. Remède durable : `stp bpdu-protection` + `storm-control` (sections 33–34).

Ne pas chercher midi à quatorze heures : une tempête de broadcast, ça se voit aux LED.

## 98. Cas n°10 — Plus d'accès SSH après un changement

**Symptômes :** timeout SSH après avoir touché à la config.

**Causes → solutions :**

1. Politique trust→local qui n'autorise plus le port 22 (ou zones inversées).
2. `protocol inbound ssh` remplacé par erreur, ou VTY en `authentication-mode password`
   sans mot de passe défini.
3. ACL sur le VTY (`acl` sous `user-interface vty`) qui bloque l'IP admin.
4. On a changé l'IP du LAN sans mettre à jour sa propre config → passer par la console.

**Prévention** : toujours garder une session console ouverte quand on touche aux accès
distants, et faire `save` **après** avoir vérifié que le nouvel accès fonctionne (pas avant).

## 99. Cas n°11 — Le basculement WAN ne se fait pas

**Symptômes :** WAN principal en panne (plus d'Internet) mais la route de secours ne
prend pas.

**Diagnostic :**

```
display nqa results test-instance admin pppoe-check
display track 1
display ip routing-table
```

**Causes → solutions :**

- La sonde NQA ping une cible **à travers le lien principal** : si le lien tombe, la
  sonde tombe aussi… mais le track ne bascule que si la sonde est configurée pour.
  Vérifier que le `track` est bien lié à la route (`display ip routing-table verbose`).
- Route de secours avec une préférence **plus basse** (plus prioritaire) que la principale
  → elle prend toujours le dessus. Vérifier les préférences.
- Le NAT n'est pas configuré sur l'interface de secours → le trafic sort sans être
  traduit. Ajouter `nat outbound` sur GE0/0/1.
- Politique de sécurité : l'interface de secours est-elle bien en zone `untrust` ?

## 100. Cas n°12 — L2TP : le client ne se connecte pas

**Symptômes :** le client VPN reste en « connexion… » puis échoue.

**Diagnostic :** `display l2tp tunnel`, `display logbuffer | include L2TP`.

**Causes → solutions :**

1. Ports UDP 500/4500/1701 non ouverts (NAT server + politique untrust→local).
2. Le client est derrière un NAT avec IPSec NAT-T mal traversé : vérifier que le
   NAT-T est actif des deux côtés.
3. Login PPP erroné ou `service-type ppp` oublié sur l'utilisateur AAA.
4. Pool d'adresses VPN épuisé ou mal lié au Virtual-Template.

## 101. Cas n°13 — DNS qui ne résout plus sur le LAN

**Symptômes :** les IP passent (ping 8.8.8.8 OK) mais aucun nom ne résout.

**Diagnostic :** `nslookup exemple.sn 192.168.10.1` depuis un PC, `display dns server`.

**Causes → solutions :**

- `dns-list` du pool DHCP pointe vers une IP qui ne répond plus → mettre le routeur
  (192.168.10.1) + un public en second.
- `dns resolve` désactivé sur le routeur.
- Politique de sécurité qui bloque le DNS (UDP 53) du LAN vers le routeur (zone local)
  ou vers Internet.
- Le FAI qui intercepte/mente les DNS : forcer des DNS publics connus pour tester.

## 102. Cas n°14 — Le routeur reboot en boucle / ne démarre plus

**Symptômes :** SYS rouge ou reboot cyclique, pas d'invite.

**Actions :**

1. Vérifier l'alimentation (autre prise, autre cordon) — un bloc qui s'effondre en
   charge donne exactement ce symptôme.
2. BootROM : interrompre le boot (`Ctrl+B`) → vérifier `display startup` : image
   corrompue ?
3. Démarrer sur l'ancienne image (si conservée) ou réinstaller depuis USB :
   option de mise à jour via clé USB dans le BootROM.
4. Si la flash semble morte : SAV constructeur/fournisseur (noter le numéro de série :
   `display device` / étiquette sous le châssis).

**Prévention** : ne jamais effacer l'ancienne image avant 48 h de validation (section 85).

## 103. Cas n°15 — CPU à 100 % / routeur lent à répondre

**Symptômes :** CLI très lente, pertes de paquets, parfois tunnels qui tombent.

**Diagnostic :**

```
display cpu-usage
display cpu-usage history
display memory
```

**Causes → solutions :**

- Debug oublié actif (`debugging ...` laissé tourner) → `undo debugging all`.
- Tempête de broadcast / boucle (cas n°9).
- Trop de sessions NAT (attaque ou P2P) : `display nat session table` → identifier la
  source, filtrer.
- Logs en `debugging` vers le syslog : repasser en `informational`.
- Si le CPU est structurellement haut en production : la plateforme est sous-dimensionnée
  pour les services activés → revoir le dimensionnement (ou monter en gamme AR730,
  **à valider avec le fournisseur**).

## 104. Cas n°16 — Certificats VPN expirés : tout tombe en même temps

**Symptômes :** tous les tunnels IPSec à certificats tombent simultanément, `display ike
sa` vide partout.

**Diagnostic :** `display pki certificate` (dates de validité), `display clock`
(l'horloge est-elle juste ?).

**Solution :** renouveler les certificats via la CA, les réimporter, renégocier.
**Prévention** : supervision des dates d'expiration (inventaire + alerte à J-60), NTP
obligatoire (une horloge fausse invalide un certificat valide).

## 105. Cas n°17 — Après une coupure électrique, plus rien ne marche

**Symptômes :** au retour du courant, le routeur démarre mais la config semble partielle
ou les tunnels ne remontent pas.

**Causes → solutions :**

1. `save` oublié avant la coupure : la config en cours n'était pas sauvegardée →
   reconfigurer ce qui manque, puis `save`.
2. La flash a été corrompue par une coupure pendant une écriture → restaurer depuis la
   sauvegarde externe (section 84).
3. L'ONT/box FAI met plus longtemps à redémarrer que le routeur : le PPPoE échoue au
   boot puis ne retente pas → vérifier les paramètres de redial du Dialer, ou
   temporiser le démarrage.
4. **Leçon** : onduleur + `save` systématique après chaque changement.

---
---

# Partie L — Maintenance préventive

## 106. Plan de maintenance : le calendrier type

| Fréquence | Actions |
|---|---|
| Mensuelle | Vérifier CPU/mémoire/température, relire les logs d'alerte, contrôler l'état des tunnels et du basculement WAN |
| Trimestrielle | Sauvegarde + vérification hors-ligne, audit des règles firewall/NAT server, test du basculement WAN, dépoussiérage |
| Semestrielle | Revue des licences/signatures (expiration), test de restauration sur spare, revue du plan d'adressage |
| Annuelle | Bilan capacitaire (le lien suffit-il encore ?), mise à jour firmware planifiée, test onduleur, révision du dossier de site |

## 107. Checklist mensuelle (15 minutes par site)

```
[ ] display cpu-usage / display memory            -> < 70 %
[ ] display device temperature (si supporté)      -> < 40 °C
[ ] display ike sa / display ipsec sa             -> tunnels UP
[ ] display nqa results ...                       -> pertes < 5 %
[ ] display logbuffer                             -> pas d'erreur récurrente
[ ] display interface brief                       -> pas d'interface down anormale
[ ] Vérifier que la sauvegarde externe est à jour
```

