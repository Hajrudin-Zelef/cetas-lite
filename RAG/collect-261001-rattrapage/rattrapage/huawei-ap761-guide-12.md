---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-12
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1139, 1242]
sha256: 99d97313b5d5e49612c48657bf188cea877879b33a5f55efc529e818ac30e033
---

# Guide ultra-complet — Huawei eKit AP761

**Limite constructeur : 16 SSID par radio** [constructeur]. C'est un maximum technique, pas un objectif : **chaque SSID actif consomme de l'airtime** (beacons à ~10/s par SSID, sondes, etc.). Au-delà de 3–4 SSID, la surcharge devient mesurable.

**Stratégie recommandée (3 SSID types) :**
| SSID | Usage | Sécurité | VLAN |
|---|---|---|---|
| `ENTREPRISE` | Salariés | WPA2-Enterprise (802.1X) ou WPA3-SAE | 20 (réseau interne) |
| `INVITE` | Visiteurs | Portail captif | 30 (Internet seul, isolé) |
| `IOT` | Caméras, portiers, capteurs | WPA2-PSK clé forte, 2.4 GHz uniquement | 40 (restreint) |

**Règles de nommage :**
- Noms **courts, sans espaces ni accents** : `ENTREPRISE`, pas `WiFi de l'Entreprise Dupont & Fils`.
- **Ne pas mettre d'infos sensibles** dans le SSID (`COMPTA-SERVEUR` = indication gratuite pour un attaquant).
- SSID invité **distinct** : jamais le même SSID pour salariés et visiteurs avec un simple VLAN différent « invisible » — source d'erreurs.
- Éviter les SSID « pièges » trop communs (`FreeWifi`, `Linksys`) : ils attirent les confusions et les evil twins (chap. 64).

**SSID masqué (non diffusé) :** le datasheet mentionne le masquage de SSID par AP [constructeur]. **Ça n'apporte aucune sécurité** (le SSID se capture en 30 secondes avec un sniffer) et ça casse l'expérience (saisie manuelle, certains clients IoT ne gèrent pas). À réserver à des cas très spécifiques (SSID IoT qu'on ne veut pas voir dans la liste des voisins).

## 44. Sécurité : panorama WPA/WPA2/WPA3 sur l'AP761

**Supporté par l'AP761** [constructeur/revendeurs] : WEP, WPA/WPA2-PSK, WPA/WPA2-Enterprise, **WPA3-SAE**, WPA3-Enterprise, WAPI, 802.11w.

| Mode | Usage | Verdict terrain |
|---|---|---|
| WEP / WPA-TKIP | Historique | ❌ **Interdit** : cassé depuis 2001/2008. Ne jamais activer, même « pour un vieux truc ». |
| WPA2-PSK (AES) | PME, IoT, invités simples | ✅ Le standard de fait. Clé forte ≥ 20 caractères. |
| WPA3-SAE | PME moderne | ✅ Mieux que WPA2-PSK (résiste au brute-force hors ligne). **Vérifier la compatibilité du parc** (chap. 46). |
| WPA2/WPA3 transition | Parc mixte | ✅ Le bon compromis pendant la migration (chap. 47). |
| WPA2-Enterprise (802.1X) | Salariés, EAP-TLS/PEAP | ✅ Le top pour un parc géré : un identifiant par utilisateur, révocable. |
| WPA3-Enterprise | Exigeant | ✅ Si le RADIUS et les clients suivent. |
| Portail captif | Invités | ⚠️ Authentification, pas chiffrement : à combiner avec de l'isolation (chap. 49). |

**Question posée par la mission : « WPA3 obligatoire en Wi-Fi 7 ? »**
Réponse honnête : **le WPA3 n'est pas « obligatoire » au sens réglementaire**, mais :
- La certification **Wi-Fi CERTIFIED 7** exige WPA3.
- La bande **6 GHz** n'accepte QUE WPA3 (pas de WPA2, pas de transition) — pas de sujet en outdoor UE (chap. 19).
- **Sur l'AP761 (Wi-Fi 6), WPA3 est optionnel** : tu peux rester en WPA2 sans problème. Le WPA3-SAE est un plus, pas une obligation.

## 45. WPA2-PSK : configuration pas à pas

Le mode le plus courant en PME. Recette en CLI (Fat AP) — clé **fictive** :

```
[AP761-wlan-view] security-profile name SEC-PSK
[AP761-wlan-sec-prof-SEC-PSK] security wpa2 psk pass-phrase VoiciUneCleForte_2026! aes
[AP761-wlan-sec-prof-SEC-PSK] quit
```

**Règles pour une clé PSK sérieuse :**
- **≥ 20 caractères**, mélange imposé (majuscules, minuscules, chiffres, symboles). `VoiciUneCleForte_2026!` = 22 caractères, correct.
- **Une clé par SSID**, jamais la même clé pour salariés et invités.
- **Rotation** : changer la clé invité régulièrement (chaque trimestre, ou à chaque départ sensible pour le SSID salariés si pas de 802.1X).
- **Ne jamais** : `12345678`, le nom de l'entreprise, `password`, une date de naissance.

**En mode cloud :** la clé se saisit dans la fiche du SSID, avec option « afficher/masquer ». La plateforme ne stocke que le hash — mais **la clé transite** : la changer depuis un poste sûr, pas depuis le Wi-Fi invité d'un café.

**Limite du PSK à connaître :** avec une clé partagée, **quiconque connaît la clé peut déchiffrer le trafic des autres** (si la clé fuite, tout le monde est exposé). Pour les salariés, le 802.1X (chap. 48) supprime ce problème : chaque utilisateur a ses propres clés de session.

## 46. WPA3-SAE : configuration et limites

**SAE (Simultaneous Authentication of Equals)** remplace la poignée de main WPA2-PSK par un échange résistant aux attaques **hors ligne** : même avec une clé faible, un attaquant qui capture la poignée de main ne peut pas la brute-forcer à la maison — il doit attaquer en ligne, essai par essai, et l'AP le détecte.

```
[AP761-wlan-view] security-profile name SEC-SAE
[AP761-wlan-sec-prof-SEC-SAE] security wpa3 sae pass-phrase VoiciUneCleForte_2026! aes
[AP761-wlan-sec-prof-SEC-SAE] quit
```

**Limites à connaître avant d'activer :**
1. **Compatibilité clients :** les clients antérieurs à ~2018–2020 (et beaucoup d'objets connectés) ne connaissent pas SAE. En mode **WPA3-only**, ils ne se connectent plus du tout (cas 11).
2. **Pas de mélange avec WEP/TKIP** : le WPA3 impose AES (GCMP pour l'Enterprise 192 bits).
3. **Le roaming 802.11r** fonctionne avec SAE, mais certains vieux clients 802.11r ont des bugs d'interop — tester.

**Recommandation :** sur un SSID **salariés avec parc moderne** (smartphones récents, PC récents), le WPA3-SAE est un vrai plus. Sur un SSID **IoT/invités** avec du matériel hétéroclite, rester en WPA2-PSK ou utiliser le mode **transition** (chap. 47).

## 47. Mode transition WPA2/WPA3 : pour qui, pourquoi

Le mode **transition** diffuse un SSID acceptant **à la fois** WPA2-PSK et WPA3-SAE : les clients modernes négocient SAE, les anciens restent en WPA2. C'est le mode de migration idéal.

```
[AP761-wlan-view] security-profile name SEC-TRANSITION
[AP761-wlan-sec-prof-SEC-TRANSITION] security wpa2 psk pass-phrase VoiciUneCleForte_2026! aes
[AP761-wlan-sec-prof-SEC-TRANSITION] security wpa3 sae pass-phrase VoiciUneCleForte_2026! aes
# (syntaxe exacte du mode transition : à valider avec ? selon la version)
```

**Avantages :** zéro exclusion de clients, montée en sécurité progressive.
**Inconvénient honnête :** la sécurité globale reste celle du maillon faible — un attaquant forcera la négociation vers WPA2 s'il le peut (downgrade). Le mode transition est une **étape**, pas un état final : objectif = WPA3-only quand le parc le permet.

**Plan de migration type (6 mois) :**
1. Mois 1–2 : inventaire du parc (quels clients supportent WPA3 ? — la console cloud ou `display wlan sta` donne les capacités).
2. Mois 3 : bascule du SSID en **transition**, communication aux utilisateurs.
3. Mois 4–5 : remplacement/mise à jour des clients récalcitrants.
4. Mois 6 : bascule en **WPA3-only**, vérification qu'il ne reste aucun client WPA2.

## 48. 802.1X / WPA2-Enterprise avec RADIUS

Pour le SSID **salariés**, le 802.1X est le niveau professionnel : chaque utilisateur s'authentifie avec **son** identifiant (login/mot de passe via PEAP, ou mieux : certificat via EAP-TLS), et reçoit des **clés de chiffrement uniques**. Un départ = on désactive le compte, pas besoin de changer la clé de tout le monde.

**Architecture :**
```
Client ---(EAP/802.1X)--- AP761 ---(RADIUS)--- Serveur RADIUS (NPS, FreeRADIUS...)
                                              └── Annuaire (AD/LDAP)
```

