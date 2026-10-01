---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-20
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1996, 2086]
sha256: b9adefe18031ea54ac260d12c70cc78c5e2b936253a30f0ebda49252796530c5
---

# Guide ultra-complet — Huawei eKit AP761

**Diagnostic :**
1. Vérifier ce que le switch **négocie** : `display poe interface` (côté switch Huawei) — af ou at ?
2. Vérifier la **classe LLDP** négociée par l'AP.
3. Câble trop long ou de mauvaise qualité : la chute de tension peut faire retomber une négociation at en comportement af.

**Remèdes :**
1. **Passer le port en 802.3at** (PoE+) : la solution dans 95 % des cas (chap. 25).
2. Si le switch ne fait que du af : **injecteur 802.3at 30 W** ou changement de switch.
3. En attendant : identifier les fonctions bridées (logs) et vérifier que l'essentiel (Wi-Fi) tourne — mais **ne pas laisser en l'état**, c'est un mode dégradé.

**Prévention :** à l'installation, **tester la négociation PoE** et la noter dans le dossier de site. Un AP installé en af « parce que le port était libre » = une panne programmée.

## 87. Cas 10 : SFP branché mais pas de lien

**Symptômes :** le module SFP est enfiché, la fibre est brassée, mais pas de lien (LED éteinte côté AP ou côté switch).

**Diagnostic ordonné :**
1. **Rappel combo** (chap. 6) : dès que le SFP est branché, **il prend la priorité** sur le RJ45. Si le SFP ne monte pas, l'AP perd aussi son lien cuivre → prévoir une intervention physique, pas à distance.
2. **Le module SFP est-il compatible ?** Certains AP sont stricts sur les SFP (codage constructeur). Tester avec un module **Huawei d'origine** ou un compatible connu.
3. **TX/RX croisés ?** La fibre doit croiser : TX d'un côté → RX de l'autre. Vérifier le brassage (jarretières croisées ou non selon les modules).
4. **Les niveaux optiques sont-ils bons ?** `display transceiver` côté équipements qui le supportent : TX/RX en dBm. Un RX < −20 dBm = fibre sale, trop longue, ou jarretière HS.
5. **Nettoyer les connecteurs** : 80 % des problèmes optiques = une férule sale. Nettoyage avec les outils adaptés (jamais au souffle ni au chiffon).
6. **Propreté du port SFP** : cache anti-poussière en place quand inutilisé.

**Test décisif :** boucle locale (loopback optique) sur le module : si le lien monte en loopback, le module est bon et le problème est la fibre/le distant.

## 88. Cas 11 : WPA3 activé, certains clients ne se connectent plus

**Symptômes :** après bascule en WPA3-SAE (ou WPA3-only), une partie du parc ne s'associe plus : vieilles imprimantes, caméras, badges, vieux smartphones.

**C'est le comportement attendu** (chap. 46), pas un bug. Conduite à tenir :

1. **Ne pas paniquer, ne pas tout rebasculer** : identifier les clients impactés (liste via la console cloud : qui ne revient pas ?).
2. **Basculer en mode transition WPA2/WPA3** (chap. 47) : les modernes passent en SAE, les anciens restent en WPA2.
3. **Isoler les récalcitrants** : SSID dédié en WPA2-PSK pour l'IoT/les vieux équipements (chap. 43) — c'est aussi plus propre en segmentation.
4. **Plan de remplacement** : noter les équipements à remplacer à leur prochain renouvellement (imprimante de 2015 → prévoir).
5. **Vérifier le PMF** (chap. 51) : en WPA3, le PMF est requis — certains clients « WPA3-compatibles » sur le papier échouent à cause du PMF. Tester PMF optionnel vs obligatoire.

**Règle :** la sécurité se monte **au rythme du parc réel**, pas au rythme des annonces marketing. Un SSID en transition bien géré vaut mieux qu'un WPA3-only qui force les utilisateurs à contourner (partage de connexion 4G = pire).

## 89. Cas 12 : portail captif qui ne s'affiche pas

**Symptômes :** le client se connecte à `INVITE`, mais la page d'authentification ne s'ouvre pas — « connecté, pas d'Internet ».

**Causes par ordre de fréquence :**
1. **Le client a un VPN actif** : le VPN chiffre tout avant la détection du portail → la sonde CNA échoue. Demander de **couper le VPN** pour l'authentification, le réactiver après.
2. **DNS personnalisé** (8.8.8.8 en dur, DNS-over-HTTPS du navigateur) : la redirection du portail repose sur l'interception DNS. Forcer le DNS du réseau pendant l'authentification, ou donner l'**URL directe** du portail (à afficher à l'accueil).
3. **Le navigateur bloque la détection** : ouvrir manuellement un site HTTP (pas HTTPS) quelconque, ex. `http://neverssl.com` — ça déclenche la redirection.
4. **HTTPS intercepté** : si le portail intercepte du HTTPS, le navigateur crie au certificat invalide — normal, mais ça effraie. Utiliser la détection CNA du système plutôt que le navigateur.
5. **Le client est déjà « authentifié » mais expiré** : vider le cache / oublier le réseau / réassocier.

**Prévention (à afficher à l'accueil, sur un chevalet) :**
```
Wi-Fi INVITÉ — mode d'emploi
1. Connectez-vous au réseau « INVITE »
2. Coupez votre VPN le temps de la connexion
3. Ouvrez votre navigateur : la page d'accueil s'affiche
   Si rien ne s'affiche : allez sur http://neverssl.com
4. Acceptez les conditions, cliquez « Se connecter »
5. Vous pouvez réactiver votre VPN
```

## 90. Cas 13 : un client consomme tout l'airtime (slow client)

**Symptômes :** tout est lent pour tout le monde, alors que le signal est bon et le canal pas saturé par les voisins. Un seul client avec un mauvais débit **monopolise le temps de parole**.

**La physique :** le Wi-Fi partage le **temps**, pas le débit. Un client à 1 Mbps qui envoie 1 Mbit occupe le canal **100× plus longtemps** qu'un client à 100 Mbps qui envoie le même Mbit. Un seul client à 1 Mbps peut diviser la capacité de la cellule par 2–3.

**Diagnostic :**
1. Identifier le client lent : console cloud (débit par client) ou `display wlan sta` — chercher les clients avec un **rate très bas** et un **RSSI très faible** (−80 dBm et pire).
2. Vérifier : est-ce un client légitime loin, ou un client bloqué en 802.11b ?

**Remèdes :**
1. **Relever les débits de base** (chap. 13) : interdire 1/2/5.5/11 Mbps en 2.4 GHz. Le client zombie ne pourra plus s'associer à 1 Mbps — il devra avoir un meilleur signal ou partir.
2. **Airtime fairness** (si supporté par la version — à vérifier) : l'AP alloue du **temps** équitable plutôt que du débit égal — les clients rapides ne sont plus pénalisés par les lents.
3. **Limiter le débit par client** sur le SSID invité : un invité n'a pas besoin de 100 Mbps.
4. **Déplacer le problème** : si c'est une caméra lointaine en 2.4 GHz, la rapprocher d'un AP ou la passer en filaire.

**Règle :** sur un site pro, **aucun client ne devrait rester associé sous −75 dBm en 2.4 GHz**. En dessous, c'est de la pollution pour tout le monde.

## 91. Cas 14 : l'AP surchauffe / redémarre en plein soleil

**Symptômes :** redémarrages aux heures chaudes (13h–17h), AP brûlant au toucher, logs avec des alertes température, retour à la normale le soir.

**L'AP761 est spécifié −40 °C à +65 °C** [constructeur] — mais c'est la température **ambiante**, pas la température du boîtier au soleil. Un boîtier métal blanc au soleil d'août peut dépasser 80 °C en surface.

**Diagnostic :**
1. Corréler les redémarrages avec la **température/météo** (logs horodatés — d'où le NTP, chap. 70).
2. Vérifier les alertes température dans les logs et la console.
3. Contrôler l'**exposition** : plein sud sans ombre ? Sous un auvent qui fait four ?

