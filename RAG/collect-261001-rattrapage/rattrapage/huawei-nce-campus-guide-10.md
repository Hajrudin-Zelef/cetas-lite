---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-10
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [777, 882]
sha256: d69a4db779cc8fad7577c31fea152d7bb4399f0b25339cd3b44883acdf4f22af
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

Règle : un site = un lieu physique avec une équipe d'astreinte identifiable. Ne pas faire un site « fourre-tout ».

## 50. Les équipements gérés : S310, AP361, AP761, AR720, USG6000

Rappel du parc de Zelef et du rôle de chaque famille sous NCE :

- **CloudEngine S310** : switches d'accès campus — VLAN, 802.1X filaire, PoE (pour les AP/téléphones), QoS. Gérés en NETCONF/YANG + SNMP.
- **AirEngine AP361** : AP Wi-Fi 6 d'intérieur — géré via WAC/CAPWAP, orchestré par NCE (SSID, profils radio).
- **AirEngine AP761** : AP Wi-Fi 6 d'extérieur (IP68) — mêmes principes, contraintes radio/extérieur spécifiques.
- **NetEngine AR720** : routeur de branche — WAN, VPN, SD-WAN de branche, DHCP/DNS local.
- **USG6000** : firewall — politiques de sécurité ; NCE-Campus gère le campus, le firewall peut être supervisé/piloté selon la matrice de compatibilité (vérifier les fonctions exactes supportées).

**Avant tout onboarding** : vérifier pour **chaque modèle** la version logicielle minimale requise par la version NCE cible (matrice de compatibilité). Un équipement en version trop ancienne = onboarding partiel ou refusé.

---

# PARTIE 6 — ONBOARDING (ZTP ET MANUEL)

## 51. ZTP — principe du zero-touch provisioning

Le **Zero-Touch Provisioning** : un équipement neuf (ou réinitialisé) sorti du carton, branché au réseau, **se configure seul** en allant chercher sa configuration sur NCE-Campus. Plus besoin d'envoyer un expert CLI sur site : un technicien local branche, NCE fait le reste.

Séquence logique :

1. L'équipement démarre en configuration d'usine et demande une adresse IP (DHCP).
2. Via DHCP (option dédiée), e-mail d'activation ou scan, il découvre **l'adresse du contrôleur NCE**.
3. Il s'enregistre auprès de NCE (authentification par ESN/numéro de série).
4. NCE l'affecte au **site** prévu, lui pousse le **template** correspondant, puis le **firmware** cible si nécessaire.
5. L'équipement redémarre/applique et passe à l'état « normal » (géré).

Bénéfice : déploiement d'un site distant en **heures** au lieu de jours, sans compétence CLI locale. Condition : le réseau d'amorçage (DHCP, connectivité vers NCE) doit être prêt **avant** l'arrivée des équipements.

## 52. ZTP via DHCP (option 148/option personnalisée) — détail

Méthode la plus courante en entreprise :

1. Le serveur **DHCP** du site (souvent l'AR720 ou un serveur central) attribue à l'équipement neuf une IP **et** une option DHCP contenant l'adresse (IP ou nom) du contrôleur NCE, voire l'identifiant du site.
2. L'équipement utilise cette information pour contacter NCE et s'enregistrer.
3. NCE vérifie l'ESN (l'équipement doit être **pré-déclaré** ou accepté selon la politique : liste blanche d'ESN recommandée en production).

Prérequis : serveur DHCP configuré avec la bonne option **avant** le branchement, connectivité IP équipement → NCE, ESN connus (relevés à la réception). Piège classique : une option DHCP mal formatée = l'équipement ne trouve jamais NCE (voir cas pratique 1).

## 53. ZTP via e-mail d'activation — détail

Variante : l'équipement reçoit ses informations d'amorçage via un **e-mail d'activation** (mécanisme prévu par Huawei pour certains scénarios cloud) :

- L'administrateur déclenche l'envoi depuis NCE vers l'adresse associée au site/technicien.
- Le technicien sur site suit la procédure d'activation (saisie/scan du code) pour lier l'équipement au site NCE.
- Utile quand le DHCP du site ne peut pas porter l'option (prestataire, réseau invité d'amorçage).

Prérequis : connectivité Internet/mail fonctionnelle, procédure écrite pour le technicien local (pas à pas, avec captures), fenêtre de support téléphonique.

## 54. ZTP via scan (QR / code) — détail

Variante terrain : le technicien **scanne** un QR code (ou saisit un code d'activation) affiché dans NCE pour le site concerné, via l'appli mobile/outil prévu :

- Lie l'équipement (ESN scanné ou saisi) au site en quelques secondes.
- Réduit les erreurs de saisie d'ESN (source fréquente d'échec).
- Idéal pour des déploiements en série (dizaines d'AP : on scanne à la chaîne).

Bonnes pratiques : préparer les étiquettes ESN à la réception (contrôle à la livraison), pré-créer les sites dans NCE, tester la procédure sur 2-3 équipements avant la vague.

## 55. Onboarding manuel — quand le ZTP n'est pas possible

Le ZTP échoue ou est inapplicable (site isolé, DHCP non maîtrisé, équipement d'occasion non réinitialisé) → onboarding manuel :

1. Réinitialiser l'équipement en configuration d'usine (si repris).
2. Lui donner une connectivité IP vers NCE (IP statique ou DHCP simple).
3. Dans NCE : ajouter l'équipement (adresse IP, ESN, identifiants), l'affecter au site.
4. Vérifier l'enregistrement, pousser le template, vérifier la conformité.
5. Documenter l'exception (pourquoi pas de ZTP ? action corrective ?).

L'onboarding manuel n'est pas un échec : c'est la **roue de secours**. Mais si > 20 % des équipements passent en manuel, c'est le processus ZTP qu'il faut réparer.

## 56. Ajouter un S310 — procédure détaillée

1. **Pré-déclaration** : dans NCE, site cible → ajouter l'équipement (modèle S310, ESN relevé à la réception, rôle : accès).
2. **Réseau d'amorçage** : port du switch amont en trunk/DHCP vers NCE, ou DHCP local avec option contrôleur.
3. **Branchement** : le S310 démarre, obtient son IP, contacte NCE, s'enregistre.
4. **Affectation** : vérifier qu'il apparaît dans le bon site, état « en enregistrement » puis « normal ».
5. **Template** : appliquer le template du site (VLAN d'accès, VLAN de management, 802.1X, SNMP, NTP, DNS).
6. **Firmware** : si la version n'est pas la cible, planifier la mise à jour (section 85).
7. **Vérification** : ping de management, `display` de conformité via NCE (get-config vs attendu), test d'un port d'accès (802.1X), PoE si AP raccordés.
8. **Documentation** : étiquetage physique (nom NCE = nom d'étiquette), mise à jour du plan de brassage.

## 57. Ajouter un AP361 — procédure détaillée

1. **Prérequis** : le WAC (ou la fonction de contrôle WLAN) doit être lui-même géré par NCE ; le S310 d'accès doit fournir PoE 802.3af (l'AP361 consomme ~8,8 W d'après sa fiche — vérifier l'alimentation : un port non-PoE = AP qui ne démarre pas).
2. **Pré-déclaration** : AP361, ESN, site, profil radio prévu.
3. **Branchement** : l'AP démarre, obtient une IP (DHCP), découvre le WAC (option DHCP ou DNS selon l'architecture), établit le tunnel CAPWAP.
4. **Remontée NCE** : l'AP apparaît sous le WAC dans NCE ; vérifier l'état « normal ».
5. **Profils** : appliquer le profil radio du site (canaux, puissance) et les SSID (section 72-74).
6. **Vérification** : l'AP diffuse les SSID, un client test s'associe, débit et roaming vérifiés, télémétrie visible dans NCE.
7. **Plan** : reporter l'AP sur le plan d'étage (position = qualité du troubleshooting futur).

## 58. Ajouter un AP761 — procédure détaillée

Même base que l'AP361, plus les spécificités **extérieur** :

1. **Physique** : fixation (mât/mur), étanchéité (presse-étoupes), parafoudre et **mise à la terre** (un AP extérieur mal relié à la terre = panne à la première saison des orages), câble extérieur blindé.
2. **Alimentation** : PoE (injecteur ou switch) — vérifier la distance (< 100 m) et la puissance.
3. **Pré-déclaration et ZTP** : identiques (ESN, site).
4. **Radio** : profil extérieur (puissance conforme à la réglementation locale, canaux autorisés en extérieur — **à vérifier** selon le pays).
5. **Vérification** : test d'association à distance réaliste, vérification de l'étanchéité après intervention, photo du montage archivée.
6. **Maintenance** : inspection visuelle annuelle (corrosion, serrage) — mettre au plan de maintenance préventive.

