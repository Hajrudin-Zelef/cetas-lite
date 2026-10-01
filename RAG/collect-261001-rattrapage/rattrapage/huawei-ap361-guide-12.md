---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-12
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["arr", "incident", "intel"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [1850, 2021]
sha256: 19419a12e2c5ec255170a8b25a342ed261342fa3a192a57722165983cb79a83a
---

# Huawei eKit AP361 — Guide ultra-complet

1. **Qualifier** : qui est impacté (1 client ? 1 AP ? tout le site ?), depuis quand,
   quoi qui a changé (travaux, nouvel équipement, upgrade ?).
2. **Segmenter** : le problème est-il **radio** (le client ne voit pas / signal
   faible), **association** (voit mais ne se connecte pas), **réseau** (connecté
   mais pas d'IP / pas d'Internet) ou **alimentation** (l'AP est mort) ?
3. **Tester du simple vers le compliqué** : câble → PoE → IP → VLAN → SSID →
   radio → client.
4. **Une variable à la fois** : changez un paramètre, testez, notez. Sinon vous
   ne saurez jamais ce qui a réparé (ni ce qui a cassé).
5. **Documentez** : fiche d'incident (date, symptômes, cause, solution, durée).
   Dans 6 mois, vous vous remercierez.

**Kit de dépannage Wi-Fi** (dans le sac, toujours) :

- Téléphone avec analyseur Wi-Fi + app eKit
- PC portable + câble console/RJ45 de secours
- Testeur de câble + testeur PoE
- Tournevis, trombone (reset), étiqueteuse
- AP361 de spare pré-configuré

---

## 121. Cas n°1 — Interférences : débit en dents de scie

**Symptômes** : débit qui s'effondre par moments, ping irrégulier, surtout à
certaines heures. Le signal est bon (-55 dBm) mais ça rame.

**Diagnostic** :

1. Analyseur Wi-Fi : regardez l'**utilisation du canal** (channel utilization).
   > 60-70 % en permanence = canal saturé.
2. Listez les SSID voisins sur le même canal (co-canal) ou chevauchant (adjacent).
3. Cherchez les **non-Wi-Fi** : four micro-ondes (2,4 GHz), caméras analogiques,
   Bluetooth massif, variateurs.
4. Corrélez avec l'heure : micro-ondes = heures des repas ; voisin = heures de bureau.

**Solutions** :

- Changez de canal (le moins occupé mesuré, pas deviné).
- En 2,4 GHz : passez en 20 MHz si ce n'est pas fait, canaux 1/6/11 stricts.
- Éloignez/éteignez la source (le micro-ondes du bureau à 1 m de l'AP, c'est du vécu).
- Baissez la puissance pour réduire la zone de contention.
- En dernier recours : ajoutez un AP pour diviser la charge (plus de cellules
  petites = moins de clients par cellule).

## 122. Cas n°2 — Débit faible alors que le signal est bon

**Symptômes** : -55 dBm, mais 15 Mbit/s au speedtest sur un AP361 capable de bien plus.

**Diagnostic** (dans l'ordre) :

1. **Négociation du port** : l'AP est-il en 100 Mbit/s au lieu de 1000 ?
   (eKit ou switch : `display interface`). Si oui → câble à refaire (§124).
2. **Bande** : le client est-il en 2,4 GHz alors que le 5 GHz est dispo ?
   (band steering, §52).
3. **Largeur de canal** : 20 MHz là où 40 était prévu ?
4. **Vieux client** : un client 802.11n ou avec pilote pourri plombe l'airtime
   (airtime fairness, §53).
5. **Charge de la cellule** : 30 clients actifs sur un AP361 = normal que ça
   rame. Mesurez à 3 h du matin pour comparer.
6. **Lien montant** : testez en filaire au switch. Si le filaire rame aussi,
   le problème n'est pas le Wi-Fi (lien Internet, firewall saturé…).

**Solutions** : selon la cause trouvée. Le cas le plus fréquent en PME : **port
négocié à 100 Mbit/s** à cause d'un câble limite. Refaire le câble = débit x10.

## 123. Cas n°3 — Des clients ne s'associent pas (ou plus)

**Symptômes** : le SSID est visible, mais la connexion échoue (« impossible de se
connecter », boucle d'authentification).

**Diagnostic** :

1. **Un seul client ou tous ?** Un seul → problème client (pilote, clé mal
   tapée, certificat). Tous → problème AP/réseau.
2. **Quoi qui a changé ?** Changement de clé ? Passage WPA3 ? PMF activé ?
   Upgrade firmware ?
3. **Type de client** : les vieux clients + WPA3 pur = échec (§61). Les IoT +
   band steering agressif = échec (§52).
4. **Logs** : échecs d'authentification dans eKit / syslog (mauvaise clé vs
   timeout vs rejet EAP).
5. **Serveur RADIUS** (si 802.1X) : est-il joignable ? Certificat expiré ?
   (Le classique : certificat RADIUS expiré un dimanche.)

**Solutions** :

- Mauvaise clé : re-saisir (le copier-coller évite les erreurs).
- Incompatibilité WPA3 : mode transition.
- Pilote Wi-Fi obsolète : mise à jour (surtout Intel AX200/201 et Realtek).
- « Oublier le réseau » côté client + se reconnecter (résout 50 % des cas
  bizarres post-changement de clé).
- Certificat RADIUS expiré : renouveler (et mettre une alerte d'expiration !).

## 124. Cas n°4 — Le PoE ne monte pas (LED éteinte)

**Symptômes** : AP muet, LED éteinte.

**Diagnostic** (5 minutes chrono) :

1. L'AP sur un **autre port** du switch (connu bon) → s'allume ? Oui = port ou
   câble d'origine en cause.
2. Un **autre équipement PoE** sur le port suspect → s'allume ? Non = port/câble.
3. **Testeur PoE** en bout de câble : tension présente ? (44-57 V attendus).
4. **Budget PoE** du switch : dépassé ? (`display poe` / voyant).
5. L'AP sur un **injecteur connu bon** → s'allume ? Non = AP HS (RMA).

**Solutions** : voir tableau §24. Le trio gagnant des causes : câble trop long/
abîmé, budget PoE dépassé, injecteur sous-dimensionné.

## 125. Cas n°5 — L'AP reboote en boucle

**Symptômes** : LED qui s'allume, l'AP apparaît brièvement puis disparaît, en boucle.

**Causes possibles** :

1. **Alimentation limite** : le PoE « tient » au boot puis s'écroule en charge
   (pics de consommation au démarrage radio). → Tester avec injecteur 802.3at.
2. **Firmware corrompu** : suite à upgrade interrompu → §118.
3. **Boucle réseau** : l'AP branché sur deux ports (ou via un petit switch
   bouclé) → storm → watchdog → reboot. Vérifiez le câblage.
4. **Surchauffe** : AP dans un caisson ou sous toiture à 55 °C → protection
   thermique. Mesurez.

## 126. Cas n°6 — Le DHCP ne passe pas (APIPA 169.254.x.x)

**Symptômes** : client associé au Wi-Fi, mais IP en 169.254.x.x → « connecté,
pas d'accès Internet ».

**Diagnostic** :

1. Le problème touche-t-il **tous les SSID/VLAN** ou un seul ? Un seul → VLAN/
   tagging en cause. Tous → DHCP ou trunk.
2. **Test filaire** : un PC en filaire sur le même VLAN obtient-il une IP ?
   Non → le problème n'est pas le Wi-Fi (serveur DHCP down, VLAN mal routé).
3. **Trunk** : le VLAN du SSID est-il dans `allow-pass` du port AP ? (§75).
4. **Port du switch** : PVID correct ? Le serveur DHCP est-il joignable depuis
   ce VLAN (relai DHCP / ip helper si le serveur est sur un autre VLAN) ?
5. **Plage épuisée** : le scope DHCP est-il plein ? (baux trop longs sur un
   VLAN invité à fort turnover).

**Solutions** : corriger le trunk, ajouter le relai DHCP, purger/raccourcir les
baux, étendre le scope. **Piège** : le relai DHCP oublié quand le serveur DHCP
est sur un autre VLAN que les clients (classique en multi-VLAN).

## 127. Cas n°7 — Le roaming casse les appels (voix hachée en marchant)

**Symptômes** : appel OK à l'arrêt, hachures/coupures en se déplaçant.

**Diagnostic** :

1. Couverture : mesurez en marchant — des zones < -67 dBm ? → ajouter/rapprocher
   des AP (§54).
2. 802.11r : activé ? (§84). Sans r, chaque changement d'AP = ré-authentification
   lente.
3. Sticky client : le téléphone reste-t-il accroché trop longtemps à l'ancien AP ?
   (puissance trop forte, §86).
4. QoS : la voix est-elle marquée et priorisée ? (§99). Un appel qui partage
   l'airtime avec un download massif sans QoS = hachures.
5. Le téléphone : certains smartphones ont un roaming agressif/conservateur
   réglable dans les options avancées (seuil de roaming).

**Solutions** : r + k + v activés, couverture -67 dBm partout, QoS voix,
puissances équilibrées. **Test de validation** : appel de 5 minutes en marchant
sur tout le site, zéro coupure.

## 128. Cas n°8 — « Le Wi-Fi est lent depuis les travaux »

**Symptômes** : dégradation datée précisément d'un événement (travaux, nouveau
voisin, nouvel équipement).

**Causes typiques** :

