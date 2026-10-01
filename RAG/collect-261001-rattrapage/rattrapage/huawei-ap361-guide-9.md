---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-9
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["incident", "intel"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [1297, 1475]
sha256: cefddb65fabb719d929d7de6750fd062f8c3293415a7079c60cd7f22c53ba270
---

# Huawei eKit AP361 — Guide ultra-complet

| Erreur | Symptôme | Diagnostic |
|---|---|---|
| PVID ≠ VLAN management | L'AP ne joint plus le cloud après passage en trunk | Remettre le port en access le temps de corriger, vérifier le PVID |
| VLAN oublié dans allow-pass | Clients associés mais pas d'IP / pas de ping passerelle | `display port trunk`, ajouter le VLAN |
| SSID → mauvais VLAN | Un invité obtient une IP du LAN | Vérifier le mapping SSID/VLAN dans l'app |
| DHCP manquant sur le VLAN | Associé mais « pas d'accès Internet » / APIPA 169.254.x.x | Tester en filaire sur le même VLAN |
| VLAN 1 utilisé partout | Aucun symptôme… jusqu'au jour où | Migrer vers le plan ci-dessus |

## 80. Checklist VLAN

- [ ] Plan de VLAN documenté (ID, nom, sous-réseau, usage)
- [ ] Ports AP en trunk, PVID = VLAN management
- [ ] `allow-pass` = uniquement les VLANs nécessaires (pas de `all` par paresse)
- [ ] DHCP actif sur chaque VLAN utilisateur
- [ ] Réservations DHCP pour les AP (IP fixes)
- [ ] VLAN invité bloqué vers le LAN (règle firewall testée)
- [ ] Test par SSID : bonne IP, bon accès, bonnes restrictions

---
---

# I. ITINÉRANCE (ROAMING)

## 81. Le problème : le client décide, pas l'AP

En Wi-Fi, **c'est le client qui décide** de changer d'AP (contrairement au
cellulaire). L'AP peut l'encourager (802.11k/v) et accélérer la ré-authentification
(802.11r), mais un client têtu restera accroché à un AP lointain. D'où l'importance
de comprendre les trois protocoles.

## 82. 802.11k : le client voit la carte

Le 802.11k permet à l'AP de fournir au client la **liste des AP voisins** avec
leurs canaux. Le client n'a plus à scanner tous les canaux (scan = coupure) :
il sait où aller.

- ✅ Activez-le : roaming plus rapide, moins de coupures.
- Invisible pour l'utilisateur, bénéfique surtout aux smartphones modernes.

## 83. 802.11v : l'AP suggère (BSS Transition)

Le 802.11v permet à l'AP de **suggérer** au client de migrer vers un meilleur AP
(trame BSS Transition Management).

- ✅ Activez-le en complément du k.
- L'AP peut ainsi **délester** un AP surchargé vers un voisin moins chargé
  (load balancing).
- Les clients ne sont pas obligés d'obéir, mais la plupart des clients récents
  suivent la suggestion.

## 84. 802.11r (Fast BSS Transition) : la ré-authentification express

Le 802.11r accélère la poignée de main lors du changement d'AP (utile surtout en
WPA2/WPA3-Enterprise où la ré-authentification 802.1X est lente).

- ✅ **Indispensable pour la voix sur Wi-Fi** (téléphones, softphones) : sans r,
  la coupure peut dépasser 200-500 ms → microcoupure audible.
- En PSK simple, le gain est moindre mais réel.
- ⚠️ Quelques vieux clients supportent mal le 802.11r → s'ils décrochent en
  boucle, testez avec r désactivé **sur un SSID de test** avant de toucher à la prod.

**Trio gagnant** : k + v + r activés sur le SSID bureautique/voix. C'est le
réglage « ne pas réfléchir » qui marche dans 90 % des cas.

## 85. PMK caching / OKC : les cousins du r

- **PMK caching** : le client réutilise sa clé maîtresse en revenant sur un AP
  déjà visité (pas de ré-authentification complète).
- **OKC (Opportunistic Key Caching)** : variante inter-AP.
- En pratique : laissez les valeurs par défaut sauf problème avéré. Le 802.11r
  couvre le besoin principal.

## 86. Sticky client : le fléau des open spaces

**Symptôme** : un utilisateur au fond de l'open space reste accroché à l'AP
d'entrée avec 1 barre, alors qu'un AP est à 5 mètres. Débit minable, plaintes.

**Causes** :

1. Puissance trop forte (le client « entend » toujours l'ancien AP — voir §49).
2. Seuils de roaming côté AP trop permissifs.
3. Client mal élevé (pilote Wi-Fi ancien, notamment sur vieux PC Windows).

**Remèdes** :

- Baisser la puissance (§49) — remède n°1.
- Activer k/v pour encourager le départ.
- **Déconnexion des clients faibles** : certains firmwares permettent de
  déconnecter un client sous un seuil RSSI (ex. -75 dBm) pour le forcer à
  se réassocier au meilleur AP. À utiliser avec doigté (trop agressif = flap).
- Côté client : mettre à jour les pilotes Wi-Fi (surtout Intel/Realtek sur Windows).

## 87. Cas voix : téléphonie sur Wi-Fi

Exigences pour de la voix correcte :

- [ ] Couverture **-67 dBm minimum** partout où l'on téléphone (pas -70, pas -75)
- [ ] 802.11r **activé**
- [ ] 5 GHz privilégié (band steering)
- [ ] QoS voix prioritaire (voir §99)
- [ ] Recouvrement de cellules vérifié en marchant (voir §54)
- [ ] Test d'appel en marchant d'un bout à l'autre du site : **zéro coupure**

> 📞 Si la direction veut « le Wi-Fi pour les téléphones » : faites signer ce
> tableau d'exigences **avant**. La voix ne pardonne pas une couverture « à peu près ».

## 88. Checklist roaming

- [ ] 802.11k activé (SSID bureau/voix)
- [ ] 802.11v activé
- [ ] 802.11r activé (indispensable si voix)
- [ ] Puissances équilibrées entre AP adjacents (pas un AP à fond au milieu de petits)
- [ ] Test de marche : association → déplacement → vérifier le changement d'AP
      (dans l'app eKit : client visible sur le nouvel AP)
- [ ] Pilotes Wi-Fi des PC du parc à jour (campagne annuelle)

---
---

# J. SÉCURITÉ WI-FI (WIDS, ROGUE AP, ISOLATION)

## 89. WIDS/wIPS sur l'AP361 : périmètre honnête

WIDS (détection d'intrusion sans fil) / wIPS (prévention) : l'AP surveille les
trames pour détecter les comportements suspects.

Sur un AP d'entrée de gamme comme l'AP361 en gestion eKit, attendez-vous à
l'essentiel (🔎 détail exact selon version) :

| Fonction | Attente réaliste |
|---|---|
| Détection rogue AP | ✅ Probable (voir §90) |
| Détection d'attaques deauth | ✅/🔎 Selon version |
| Contre-mesures actives (déauth du rogue) | 🔎 Souvent réservé aux gammes supérieures / contrôleur |
| Classification auto autorisé/voisin/rogue | 🔎 Selon version |

> 💡 En PME, le wIPS « de base » suffit à **voir** les problèmes. La
> contre-mesure active (qui est juridiquement délicate — vous émettez vers un
> équipement tiers) reste un sujet d'expert : contentez-vous de détecter et
> d'alerter, sauf politique écrite.

## 90. Détection des rogue AP : la méthode

Un **rogue AP** = un point d'accès non autorisé sur votre réseau (ex. : un
collaborateur qui branche sa box perso « pour avoir du Wi-Fi dans son bureau »,
ou un attaquant).

**Procédure** :

1. Activez la détection rogue dans l'app eKit (si disponible).
2. L'AP361 écoute et remonte les SSID/BSSID vus qui ne font pas partie du site.
3. **Qualifiez** chaque détection :
   - SSID de l'entreprise émis par un BSSID inconnu → **suspect** (evil twin possible).
   - Box du voisin d'en face → **voisin**, à ignorer (mais notez le canal : interférences).
   - Répéteur inconnu branché sur votre LAN → **rogue avéré** : débranchez physiquement.
4. Pour confirmer qu'un rogue est **sur votre LAN** (wired rogue) : cherchez sa
   MAC sur votre switch (`display mac-address`) → le port donne l'emplacement physique.

**Réflexe** : à chaque détection « SSID entreprise / BSSID inconnu », traitez
comme un incident (voir §145) jusqu'à preuve du contraire.

## 91. Evil twin : le jumeau maléfique

Un attaquant diffuse **votre SSID** depuis son propre AP pour capturer des
handshakes ou servir un faux portail.

**Défense** :

- WPA3-SAE / 802.1X : l'attaquant ne peut pas prouver son identité → les clients
  bien configurés refusent (d'où l'importance des certificats valides en EAP).
- En PSK : sensibilisez les utilisateurs (« si votre téléphone redemande la clé
  du Wi-Fi bureau sans raison, appelez l'IT »).
- Détection rogue (§90) + ronde occasionnelle à l'analyseur Wi-Fi.

## 92. Attaques deauth : comprendre et parer

Une trame de **désauthentification** forgée peut déconnecter un client (outil
d'attaque classique, 2 minutes avec un ESP32).

