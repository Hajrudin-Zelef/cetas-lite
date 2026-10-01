---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-5
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [606, 763]
sha256: 00978c4b2b349a7bb635090f6a95bddff882d964637599debfc51cc68004d45e
---

# E. GESTION : APP EKIT, CLOUD, FAT / FIT / CLOUD

## 33. Les trois modes de fonctionnement : Fit, Fat, Cloud

✅ La datasheet annonce trois modes. Voici ce que ça veut dire concrètement :

| Mode | Pilotage | Idéal pour | Limites |
|---|---|---|---|
| **Cloud** | Portail/app **eKit** (cloud Huawei) | PME, multi-sites, 1 à ~50 AP | Dépend d'Internet pour la gestion (le Wi-Fi local continue si le cloud tombe — à vérifier selon version) |
| **Fat** (autonome) | Web/CLI local sur chaque AP | 1-3 AP, site isolé sans Internet, maquette | Gestion AP par AP, fastidieux au-delà de 5 |
| **Fit** | Contrôleur WLAN (AC) Huawei | Parcs entreprise classiques | Nécessite un AC (coût, complexité) — hors positionnement eKit |

> 💡 **En pratique eKit** : 95 % des déploiements AP361 se font en mode **Cloud**.
> Le mode Fat sert de secours / maquette. Le mode Fit ne se justifie que si vous
> avez déjà un AC Huawei dans l'entreprise.

## 34. Mode Cloud eKit : quand l'utiliser (et quand ne pas)

**Utilisez le mode Cloud quand** :

- Le site a un accès Internet fiable (même modeste : la gestion consomme peu).
- Vous gérez plusieurs sites (le cloud donne une vue unifiée).
- L'équipe sur place n'est pas spécialiste réseau (l'app guide pas à pas).
- Vous voulez superviser à distance (alertes, état des AP, statistiques).

**Évitez-le (ou prévoyez un plan B) quand** :

- Site sans Internet ou avec Internet instable (carrière, zone rurale) → mode Fat.
- Contrainte réglementaire interdisant la gestion cloud hors UE/zone (données de
  gestion hébergées chez Huawei — vérifiez votre politique de conformité).
- Besoin de fonctionnalités avancées type WIDS poussé, NAC 802.1X complexe :
  vérifiez que le cloud eKit les expose pour l'AP361 (🔎 selon version).

## 35. Mode Fat/autonome : quand l'utiliser

- **Maquette** avant déploiement : 1 AP sur un bureau, on teste SSID/VLAN/roaming.
- **Site isolé** : pas d'Internet, pas de cloud → configuration locale web/CLI.
- **Secours** : si le cloud est injoignable durablement, basculer temporairement
  en Fat pour garder la main localement (procédure à tester **avant** d'en avoir
  besoin).

**Inconvénient majeur** : chaque AP se configure séparément. À 10 AP, c'est déjà
une corvée et une source d'incohérences. Documentez chaque AP (fichier de config
sauvegardé, voir §112).

## 36. Mode Fit : quand l'utiliser

- Vous avez **déjà** un contrôleur WLAN Huawei (AC) qui gère un parc AirEngine.
- L'AP361 peut alors rejoindre le parc existant comme AP léger (léger = la
  config vient de l'AC).
- 🔎 Vérifiez la **compatibilité de version** AC ↔ AP361 avant tout achat :
  un AP trop récent pour un AC trop vieux ne s'enregistrera pas.

## 37. App eKit : onboarding pas à pas (procédure détaillée)

Prérequis : compte Huawei, app eKit à jour, Bluetooth/GPS activés si l'app les
demande pour la découverte locale.

1. **Créer l'organisation / le site** : `eKit > Sites > +` → nom explicite
   (`ClientX-Siege-Etage2`), adresse, fuseau horaire.
2. **Ajouter les équipements** : scannez le QR code de chaque AP361
   (ou saisissez le SN manuellement). Vérifiez que le SN scanné correspond à
   l'étiquette physique — une erreur de saisie = un AP « fantôme ».
3. **Brancher** : PoE + lien vers le switch, DHCP actif (voir §27).
4. **Attendre l'état « En ligne »** (2 à 5 min après boot).
5. **Assistant SSID** : créez le SSID principal (WPA2/WPA3, clé robuste fictive
   type `Bureau-2026-Exemple!Changez-Moi`), choisissez les bandes.
6. **Appliquer au site** : la config est poussée vers tous les AP du site.
7. **Vérifier** : topologie dans l'app, chaque AP vert, test d'association (§31).

**Conseil d'équipe** : faites l'onboarding **à deux** — un qui scanne/étiquette
en haut de l'échelle, un qui valide dans l'app en bas. Ça divise les erreurs par deux.

## 38. Portail cloud eKit : structurer vos sites

Le portail web eKit (ekit.huawei.com) donne la vue d'ensemble :

- **Organisation** → **Sites** → **Équipements** : nommez de façon hiérarchique
  (`SIEGE/ETAGE2/AP-03-ZoneOpenSpace`).
- **Modèles de configuration** : créez un template « SSID Bureau standard » et
  appliquez-le aux nouveaux sites plutôt que de tout refaire à la main.
- **Comptes** : un compte admin par technicien (pas de compte partagé
  « admin » — voir §144), avec journal des actions si disponible.

## 39. Rôles et comptes : qui peut faire quoi

| Rôle (typique) | Droits | Pour qui |
|---|---|---|
| Super admin / Owner | Tout, y compris suppression du site | Vous (chef de service) |
| Admin réseau | Config SSID/radio, upgrades | Techniciens réseau |
| Lecture seule | Supervision, pas de modif | Astreinte, direction |
| Installateur | Onboarding (ajout d'AP) sans config | Prestataire de pose |

> ⚠️ **Retirez les accès des prestataires** à la fin du chantier. Un compte
> installateur oublié, c'est une porte dérobée.

## 40. Topologie dans eKit : lire la carte

L'app/le cloud affichent la topologie (AP ↔ switch ↔ Internet). Apprenez à la lire :

- AP **vert** : en ligne, OK.
- AP **gris** : hors ligne → vérifier PoE/réseau (voir §124).
- AP **orange/rouge** : alarme → cliquer pour le détail (forte utilisation CPU,
  interférences, clients en échec d'association).
- Lien **dégradé** : négocié à 100 Mbit/s → câble à refaire.

Faites une **capture d'écran de la topologie saine** le jour de la mise en
service : c'est votre référence « avant » pour comparer en cas de panne.

## 41. Mise en service multi-sites : méthode

1. **Site pilote** (1 site, 2-3 AP) : validez le template de config complet
   (SSID, VLAN, sécurité, radio).
2. **Documentez le template** : exportez/sauvegardez la config du site pilote.
3. **Dupliquez** vers les autres sites via le cloud (même SSID, mêmes VLAN si
   l'adressage est homogène).
4. **Adaptez localement** : canaux radio (l'environnement diffère), puissance,
   nom du site.
5. **Vague d'upgrade firmware** synchronisée sur tous les sites (même version
   partout — voir §30).

## 42. Sauvegarde de la configuration cloud

- Le cloud eKit conserve la configuration du site : c'est déjà une forme de
  sauvegarde. Mais **ne comptez pas que sur lui**.
- Exportez périodiquement la configuration (si la fonction existe dans votre
  version — 🔎) ou documentez-la dans un fichier versionné (voir §112).
- Notez les **versions firmware** par site dans un tableau de suivi.

## 43. Quand le cloud est injoignable : plan B

Scénario : Internet coupé, vous devez modifier un SSID **maintenant**.

1. Accès local : l'AP en mode Cloud garde-t-il une interface web locale ?
   🔎 **À tester sur votre version** : sur certains firmwares, l'accès local
   reste possible ; sur d'autres, tout passe par le cloud.
2. Si pas d'accès local : bascule temporaire en mode Fat (reset + config locale),
   puis ré-onboarding cloud quand Internet revient.
3. **Testez cette procédure une fois par an** en exercice (comme un exercice
   incendie). Le jour où vous en aurez besoin, il sera trop tard pour découvrir
   qu'elle ne marche pas.

## 44. Checklist de gestion

- [ ] Organisation/sites nommés de façon hiérarchique et homogène
- [ ] Un compte par technicien, pas de compte partagé
- [ ] Accès prestataires révoqués en fin de chantier
- [ ] Template de configuration documenté et versionné
- [ ] Capture « topologie saine » archivée (référence)
- [ ] Procédure « cloud injoignable » testée une fois

---
---

# F. CONFIGURATION RADIO

## 45. Les deux bandes : à quoi sert chacune

