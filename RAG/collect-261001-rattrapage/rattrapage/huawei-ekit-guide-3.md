---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-3
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Huawei"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [161, 233]
sha256: 00b2156601aa897dbdc85322eabb797165d7ba955689ffb093ef3f53316f857e
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- Portail officiel : **ekit.huawei.com** (cité dans toutes les fiches techniques).
- Documentation technique : **support.huawei.com** (fiches, guides d'installation, notes de version) et **e.huawei.com** (présentations solutions).
- **À vérifier sur la documentation officielle** systématiquement avant un déploiement : la fiche exacte du modèle (les gammes évoluent vite — AR281, AP772, F700D sont des nouveautés 2025-2026), les notes de version du firmware que tu vas installer, et la disponibilité des fonctions cloud dans ton pays.
- Ton **distributeur Gold eKit** : c'est ton premier niveau de support (stock, RMA, firmwares). Garde son contact dans ton téléphone — section 101.

## 11. Vocabulaire eKit : Fit, Fat, Cloud — ne plus confondre

| Mode | Signification | Quand l'utiliser |
|---|---|---|
| **Fit** | L'AP est piloté par un contrôleur (WAC) | Avec un AR eKit qui fait office de WAC (l'AR280/AR281 intègre la fonction WAC), ou un WAC externe |
| **Fat** | L'AP est autonome, configuré unitairement | Petit site, dépannage, ou quand tu veux garder la main en local |
| **Cloud** | L'AP est piloté par le cloud eKit (app/SNC) | Le mode standard des déploiements eKit multi-sites |

Un AP eKit peut **changer de mode** selon la fiche (AP361 : Fit, Fat, cloud). En pratique eKit : tu onboardes en **Cloud** via l'app, et tu bascules en local si besoin. L'erreur classique : laisser des AP en Fat « parce que ça marche » puis perdre la supervision centralisée. Choisis ton mode **avant** le déploiement et tiens-le.

## 12. Ce que l'app eKit n'est PAS

Pour cadrer les attentes (les tiennes et celles du client) :
- Ce n'est pas un **NMS** complet (pas de NetFlow/sFlow, pas de corrélation d'événements avancée, pas d'API publique documentée à ma connaissance — **à vérifier sur la documentation officielle**).
- Ce n'est pas un outil de **planification radio** (pas de heatmap prédictive comme Ekahau) : le placement des AP reste ton métier, au mètre et au plan.
- Ce n'est pas un **contrôleur local redondé** : si Internet tombe, la gestion cloud tombe (le réseau local, lui, continue de fonctionner — le data plane ne dépend pas du cloud, voir section 38).

## 13. Les 9 secteurs verticaux eKit (pour parler au client)

Huawei décline eKit en solutions par secteur : **bureaux, hôtellerie, éducation, santé, commerce/retail, restauration**, et d'autres. Intérêt pratique : quand tu réponds à un besoin, cherche la « solution eKit » du secteur — tu y trouveras une topologie de référence et une liste de matériel type déjà pensées par le constructeur. Ça t'évite de réinventer la roue et ça rassure le client (« c'est la solution hôtel officielle, pas un bricolage »).

## 14. eKitStor et IdeaHub : ce que tu dois en savoir (en 2 minutes)

- **eKitStor Xtreme 200E** (SSD) et **Shield 210** (stockage portable) : pour les postes et la sauvegarde locale. L'AR281 intègre une fonction **NVR** avec compression « SuperCoding 3.0 » annoncée à -85 % de stockage vidéo — intéressant si le client couple réseau + vidéosurveillance.
- **IdeaHub S3 / B3** : écrans collaboratifs 4K pour salles de réunion (visio, tableau blanc, projection sans fil). Ils se connectent à ton réseau eKit comme des clients Wi-Fi/filaires gourmands — prévois-le dans le dimensionnement (une visio 4K, c'est du débit réservé).
- Tu n'as pas besoin d'être expert de ces gammes, mais tu dois savoir qu'elles existent quand le client dit « et pour nos salles de réunion ? ».

## 15. Check-list « avant de proposer eKit » (à valider avec le client)

- [ ] Nombre d'utilisateurs et de terminaux (avec marge +30 % à 3 ans) ?
- [ ] Surface, nombre de pièces/étages, nature des murs (béton, placo, verre) ?
- [ ] Débit Internet disponible et type de raccordement (fibre, 4G/5G, ADSL) ?
- [ ] Besoin de PoE : combien d'AP, caméras, téléphones IP ?
- [ ] Invités / portail captif ? (oui → prévoir la fonction portail de l'AP ou de l'USG)
- [ ] Vidéosurveillance sur le même réseau ? (oui → VLAN dédié + NVR, attention au dimensionnement)
- [ ] Multi-sites ? (oui → eKit cloud prend tout son sens)
- [ ] Contraintes de sécurité : invités isolés, 802.1X, VPN site-à-site ?
- [ ] Le client a-t-il déjà du matériel à réutiliser ? (voir section 61)
- [ ] Qui administre au quotidien ? (si « personne », l'app eKit + un contrat de maintenance, c'est toi)

---

## 16. L'app eKit : c'est quoi, où la trouver

L'**app HUAWEI eKit** (Android/iOS) est l'outil central du déploiement : c'est par elle que tu **onboardes** les équipements (tu les déclares au cloud), que tu configures les sites, et que tu fais la maintenance courante. D'après les fiches techniques (AP361, AP266, switchs S310/S620) :
- déploiement **par Wi-Fi** : le téléphone se connecte au Wi-Fi de gestion de l'équipement, l'app configure le projet réseau, les équipements sont automatiquement onboardés et gérables à distance ;
- déploiement **par scan de code-barres** : tu scannes le numéro de série (SN) sur le châssis, l'info est synchronisée vers le système eKit, l'équipement est onboardé — idéal pour ajouter des équipements à un projet existant ou quand le modèle ne supporte pas le déploiement par Wi-Fi.
- Après le déploiement, « davantage d'opérations de maintenance du projet peuvent être effectuées sur l'app » (fiche AP361).

**Installation :** cherche « HUAWEI eKit » sur le Play Store / App Store. Vérifie que c'est bien l'app officielle Huawei (éditeur Huawei). **À vérifier sur la documentation officielle** : la disponibilité de l'app et du cloud eKit dans ton pays — certaines fonctions cloud peuvent varier selon la région.

## 17. Création du compte : la base à ne pas rater

1. Crée **un compte par structure d'intervention**, pas un compte par technicien avec son adresse perso. Si ton technicien part, tu perds l'accès aux sites clients.
2. Recommandation terrain : une adresse e-mail **générique de ton service** (type `reseau@tonentreprise.xx` — exemple fictif) comme compte « propriétaire », puis tu invites les techniciens avec leurs comptes nominatifs en rôle restreint (rôles : section 34).
3. Active la **double authentification** si l'app la propose (**à vérifier sur la documentation officielle** — si elle n'existe pas encore, mot de passe long unique + gestionnaire de mots de passe).
4. Note le compte propriétaire dans ton **registre des accès** (section 106) : client, site, compte, rôle, date. C'est ton assurance-vie le jour où le client t'appelle dans 2 ans.

## 18. Organisation par sites : le concept central

L'app organise tout par **sites** (un site = un bâtiment, une boutique, un hôtel). Chaque site contient ses équipements, ses SSID, ses VLAN, ses alertes. Règles d'organisation :
- **1 site physique = 1 site eKit.** Ne mets jamais deux bâtiments dans le même site « pour simplifier » : tu perdrais la supervision par bâtiment et les configs divergeraient.
- Nomme les sites de façon **explicite et durable** : `CLIENT – Ville – Bâtiment` (ex. fictif : `HOTEL PALM – Douala – Bâtiment A`). Pas de « Site 1 », « Test », « Nouveau site (2) ».
- Un compte peut gérer **plusieurs sites** (multi-sites, section 35) : c'est tout l'intérêt pour un prestataire qui suit 10 boutiques d'une enseigne.
- Documente pour chaque site : adresse, contact sur place, plan de câblage, plages d'adresses (section 66).

## 19. Onboarding par Wi-Fi : la méthode rapide (AP)

