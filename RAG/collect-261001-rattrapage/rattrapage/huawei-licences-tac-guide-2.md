---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-2
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "incident", "license", "sandbox"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [142, 292]
sha256: 6e92cf5e89d5b283d95e2dd6973b45ac8feb7bbdb44d793d36df94e23a55fb3f
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

| Produit | Licence de fonctionnalité | Effet |
|---|---|---|
| USG6000 | Licence **IPS** (Intrusion Prevention) | Active la détection/prévention d'intrusion |
| USG6000 | Licence **AV** (Antivirus) | Active l'analyse antivirus des flux |
| USG6000 | Licence **URL Filtering** | Active le filtrage d'URL par catégories |
| USG6000 | Licence **Anti-Spam** | Active l'antispam mail |
| USG6000 | Licence **Sandbox / C&C** (selon modèles) | Détection avancée |
| eSight | Licence de **module** (ex. : gestion WLAN, inventaire) | Déverrouille les modules |
| AR720 | Licences de fonctions (ex. : voix, sécurité) | Selon le modèle et la version VRP |

Sans ces licences, l'équipement route et commute normalement, mais les lames
« UTM » (Unified Threat Management) restent inactives. **C'est le cas le plus
fréquent de « panne » un lundi matin** : la licence 1 an achetée avec
l'équipement a expiré pendant le week-end (voir cas vécu n°101).

## 7. Types de licences (2/3) : les licences de capacité

Elles fixent des **quantités** : nombre d'objets gérés, de sessions, d'utilisateurs.

| Produit | Licence de capacité | Exemple |
|---|---|---|
| Contrôleur WLAN / NCE | Nombre d'**AP gérés** | 32, 64, 128 AP... |
| USG6000 | Nombre d'**utilisateurs SSL VPN** concurrents | 50, 100, 200... |
| USG6000 | Nombre de **tunnels IPsec** (selon modèles) | Quotas |
| eSight | Nombre de **nœuds gérés** | Paliers (50/200/1000...) |
| Switch (selon séries) | Ports 40G→100G (licence **RTU** — Right To Use) | Activation de débits |

Dépasser la capacité autorisée se traduit en général par un refus poli :
l'AP n°65 ne montera pas sur le contrôleur, le client VPN n°101 sera rejeté.
Les logs indiquent alors une cause « license » (voir section 76).

## 8. Types de licences (3/3) : temporaires vs permanentes

Huawei distingue (terminologie constatée sur les guides officiels) :

| Type | Durée typique | Usage | Obtenue via |
|---|---|---|---|
| **Commissioning** | 30 jours | Mise en service, tests | Plateforme ESDP |
| **Temporary** | 60 jours | Avant-vente, démo, reprise après incident | Plateforme ESDP |
| **Commercial** | Permanente (ou durée contractuelle) | Production | Achat classique |

Points d'attention :
- Les licences **UTM** (IPS/AV/URL) des USG sont très souvent vendues **à durée
  limitée** (1 an ou 3 ans) : ce sont des abonnements, pas des achats uniques.
- Une licence temporaire expirée **ne se prolonge pas** : il faut en générer
  une nouvelle.
- En cas de RMA, une licence commerciale peut être **transférée** vers le nouvel
  ESN (section 32) ; une licence temporaire, en général, non.

## 9. Licences : cas de l'USG6000

L'USG6000 (HiSecEngine) est l'équipement du parc le plus « gourmand » en
licences. Cartographie :

```
USG6000
├── Licences de base (souvent incluses) : routage, NAT, firewall stateful, IPsec de base
├── Licences UTM (abonnement 1/3 ans, à renouveler !)
│   ├── IPS          → mise à jour des signatures d'attaques
│   ├── AV           → base virale
│   ├── URL Filter   → base de catégories d'URL
│   └── Anti-Spam    → base antispam
├── Licences de capacité
│   ├── Utilisateurs SSL VPN concurrents
│   └── (selon modèle) tunnels / sessions
└── Mises à jour de signatures : nécessitent une licence UTM valide
    → sans licence valide, les signatures ne se mettent plus à jour,
      la protection se dégrade (détection sur des bases obsolètes).
```

**Commandes de contrôle sur USG6000 :**
```
<USG6000> display license
<USG6000> display utm license   (selon version)
```
Vérifiez systématiquement la colonne « Expire Date » (voir section 27).

## 10. Licences : cas d'eSight

eSight (plateforme de supervision) utilise un modèle de licences par **modules**
et par **paliers de nœuds** :

- Chaque grand domaine (réseau, WLAN, inventaire, etc.) peut nécessiter sa
  licence de module.
- Le nombre d'équipements supervisés est plafonné par palier.
- La licence s'installe depuis l'interface d'administration eSight
  (pas en CLI VRP).

> eSight étant en fin de vie commerciale au profit de NCE dans la stratégie
> Huawei récente — **à vérifier sur le portail officiel** pour votre version —,
> vérifiez que vos renouvellements portent bien sur le bon produit.

## 11. Licences : cas des switches S310 et du routeur AR720

- **S310** : switch d'accès eKit/entreprise. La plupart des fonctions L2/L3 de
  base ne demandent pas de licence. Les fonctions avancées (selon modèle :
  stacking iStack, routage avancé) peuvent nécessiter des licences RTU.
- **AR720** : routeur de la gamme NetEngine AR. Licence de base incluse ;
  fonctions voix, sécurité avancée ou capacité (tunnels) sous licence selon
  la référence commandée.

**Réflexe :** à l'achat, exiger du partenaire la liste écrite des licences
**incluses** vs **optionnelles** avec leur durée (voir checklist section 113).

## 12. Licences : cas des AP361 / AP761 (Wi-Fi)

Bonne nouvelle : sur les AP Huawei (dont AP361 et AP761 eKit), il n'y a en
général **pas de licence par AP** à acheter. Le modèle est :

- Les AP sont gérés par un contrôleur (WAC), par NCE, ou en mode **cloud/local**
  selon la gamme.
- C'est le **contrôleur** qui porte la licence de capacité « nombre d'AP gérés ».
- En mode eKit (gestion via l'application / le cloud eKit), la gestion est
  simplifiée et ne demande pas de licence AP unitaire — **à vérifier sur le
  portail officiel** pour votre mode de gestion exact.

Conséquence pratique : si un AP ne « monte » pas sur le contrôleur alors que
les autres fonctionnent, pensez « capacité de licence du contrôleur atteinte »
avant « AP défectueux ».

## 13. Licences : NCE et autres plateformes

NCE (Network Cloud Engine), successeur logique d'eSight/ Agile Controller dans
la stratégie Huawei, utilise des licences par contrôleur et par ressources
gérées. Les principes restent identiques (ESN, `.dat`, portail), seule
l'interface d'import change. **À vérifier sur le portail officiel** pour la
version exacte que vous déployez.

## 14. Où sont stockées les licences sur l'équipement

- Le fichier `.dat` est copié dans la **mémoire flash** de l'équipement
  (répertoire visible via `dir`).
- L'activation (`license active`) enregistre la licence dans la zone dédiée ;
  elle **survit au reboot** (pas besoin de réactiver après redémarrage).
- Elle **ne survit pas** à un remplacement de carte mère / MPU (nouvel ESN) :
  prévoir le transfert (section 32).
- En cluster actif/standby (ex. : deux USG en HA), **chaque membre** doit avoir
  sa propre licence activée avec son propre ESN.

**Checklist de fin de partie A :**
- [ ] Je sais distinguer licence / support / garantie.
- [ ] Je sais récupérer un ESN par `display esn`.
- [ ] Je connais les licences à durée limitée de mon parc (UTM USG6000 en tête).
- [ ] J'ai identifié qui porte les licences de capacité (contrôleur WLAN...).

---

# B. ACTIVER UNE LICENCE PAS À PAS

## 15. Vue d'ensemble du workflow

