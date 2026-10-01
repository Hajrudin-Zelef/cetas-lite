---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-5
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [320, 390]
sha256: 9b7abffcdcd31cc2ac62fe0b3bec2d0abaf0b364d77947bd9fde05168bef1bb4
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- L'**AP673H (iGuard)** détecte les caméras espion (y compris non connectées au Wi-Fi) par analyse des ondes radio + IA, 24h/24 d'après Huawei.
- **Limite officielle** (fiche constructeur) : la fonction spycam ne marche qu'en **gestion locale** — pas de remontée cloud, pour protéger la confidentialité. Donc : tu installes l'AP673H, tu le gères en local pour cette fonction, et tu le supervises via l'app pour le reste.
- Le détecteur portable **DF10** + appli **iGuard** : pour des contrôles ponctuels (chambre d'hôtel avant l'arrivée d'un VIP, salle de conseil avant une réunion sensible). Autonomie annoncée ~1,5 h, taux de fausses alertes annoncé 5 % avec le modèle IA multidimensionnel.
- Ne promets jamais « zéro caméra garantie » : c'est un **outil de détection**, pas une certification. Positionne-le comme une couche de sécurité supplémentaire, avec une procédure (qui scanne, quand, que faire en cas d'alerte).

## 30. L'app eKit : limites à connaître avant de t'engager

- **Dépendance Internet** : sans Internet sur le site, pas de gestion cloud (le réseau local continue de tourner — section 38).
- **Dépendance au téléphone** : pas de vrai portail web complet d'administration d'après les infos disponibles (**à vérifier sur la documentation officielle** — le SNC apporte une console web, section 31).
- **Fonctions avancées limitées** : pas de scripting, pas d'automatisation type Ansible, pas d'API publique documentée à ma connaissance.
- **Verrouillage écosystème** : l'app ne gère que du Huawei eKit. Ton parc existant non-Huawei reste à gérer à part (interopérabilité : section 61).
- **Données hébergées chez Huawei** : à signaler au client dans les secteurs sensibles (données de gestion du réseau sur un cloud constructeur — clause contractuelle à prévoir).

---

## 31. Le cloud eKit : SME Network Center (SNC), c'est quoi ?

Le **SME Network Center (SNC)** est la plateforme cloud de gestion des réseaux eKit, complémentaire de l'app mobile : l'app, c'est le tournevis (déploiement et gestes quotidiens) ; le SNC, c'est l'atelier (vue d'ensemble, planification, maintenance multi-sites). D'après Huawei (MWC 2025) : le SNC couvre la **planification, le déploiement, l'inspection et l'O&M en ligne** du réseau, avec une fonction de **dépannage simplifié 2.0** et d'**expansion simplifiée** du réseau.
Concrètement, attends-toi à une console qui montre : l'état du réseau, l'état des équipements et l'état des connexions des clients (STA) **pour tous les sites d'un tenant**, de façon centralisée (c'est ce que décrit la fiche AP361 pour la plateforme cloud).

## 32. Comptes et tenants : comment c'est organisé

- Un **tenant** (locataire cloud) = généralement ton entreprise (prestataire) ou le client final — à trancher selon le contrat (section 28).
- Le tenant contient les **sites**, chaque site contient les **équipements**.
- Un équipement n'appartient qu'à **un seul site** à la fois. Pour le déplacer : le retirer du site A puis l'onboarder sur le site B (ne le fais jamais « à chaud » sans prévenir : le déplacement réinitialise la configuration réseau de l'équipement — **à vérifier sur la documentation officielle** pour le comportement exact).
- **Règle d'or multi-clients :** un tenant par client final, ou un tenant prestataire avec des sites bien séparés ? La pratique la plus sûre : **un tenant par client** quand le client est autonome, **un tenant prestataire multi-sites** quand tu assures l'infogérance. Ne mélange jamais deux clients dans les mêmes vues sans cloisonnement.

## 33. Le modèle économique : « gratuit », vraiment ?

Huawei et ses distributeurs annoncent une gestion **app + cloud gratuite, sans frais de licence**. Ce que ça veut dire en pratique, d'après les sources publiques :
- pas de licence annuelle par AP comme chez certains concurrents ;
- l'app et le SNC sont inclus dans l'achat du matériel.
**Nuances honnêtes :**
- « Gratuit » ne veut pas dire « sans coût » : il te faut de la connectivité Internet sur chaque site, un smartphone, et ton temps.
- Les fonctions avancées (pare-feu USG : IPS/AV/SSL VPN au-delà des seuils par défaut, certaines fonctions IA) peuvent relever de **licences ou d'abonnements de signatures** — **à vérifier sur la documentation officielle** et auprès de ton distributeur avant de chiffrer.
- Vérifie aussi la **durée d'engagement** : le gratuit d'aujourd'hui peut devenir payant demain. Ne construis pas ton business model sur une promesse marketing sans filet contractuel.

## 34. Rôles et permissions : qui peut faire quoi

Principe à appliquer même si le détail des rôles varie selon la version (**à vérifier sur la documentation officielle**) :
| Rôle (logique) | Peut voir | Peut modifier | À donner à |
|---|---|---|---|
| Propriétaire / Admin | Tout, tous les sites | Tout | Toi (compte générique du service) |
| Technicien | Sites assignés | Configurer, redémarrer, mettre à jour | Ton équipe terrain |
| Lecture seule | Sites assignés | Rien (diagnostic uniquement) | Client final, stagiaire |
| Installateur temporaire | 1 site, durée limitée | Onboarding uniquement | Sous-traitant ponctuel |

- **Revue trimestrielle des accès** : qui a encore accès à quoi ? Un ex-technicien avec l'admin sur 15 sites clients, c'est une bombe à retardement.
- Journalise les changements importants (qui a modifié quel SSID, quand) — si la plateforme le propose (**à vérifier**), active-le.

## 35. Multi-sites : le vrai terrain de jeu du cloud eKit

C'est là que le cloud paie : une enseigne avec 12 boutiques, un groupe avec 5 agences.
- **Tableau de bord global** : tous les sites, avec un code couleur (vert/orange/rouge) par site. Le matin, un coup d'œil suffit pour savoir où intervenir.
- **Modèles de configuration** : définis une fois la configuration type « boutique standard » (VLAN, SSID, règles), applique-la aux nouveaux sites. C'est le « déploiement simplifié » mis en avant par Huawei.
- **Mises à jour par vagues** : site pilote → 3 sites → tous les sites. Jamais tout d'un coup un vendredi soir.
- **Alertes centralisées** avec le nom du site dans chaque alerte (section 27).
- Limite honnête : le cloud ne remplace pas une **visite sur site** quand le problème est physique (câble arraché, AP décroché, coupure électrique). La supervision dit *où* ça fait mal, pas *pourquoi* au niveau physique.

## 36. Ce qu'on peut faire à distance (et qui marche bien)

- Voir l'état de chaque équipement (en ligne/hors ligne, charge CPU/mémoire — niveau de détail **à vérifier**) ;
- voir les **clients connectés** par AP et par SSID (utile : « le Wi-Fi est lent » → 45 clients sur un seul AP361 = problème de densité, pas de panne) ;
- redémarrer un équipement à distance ;
- modifier un SSID / un PSK (ex. : changer le mot de passe invités d'un hôtel chaque semaine) ;
- pousser une mise à jour firmware ;
- activer/désactiver le PoE d'un port (caméra qui ne répond plus → cycle PoE à distance avant de se déplacer) ;
- diagnostiquer un lien (inspection en ligne, tests de connectivité — « troubleshooting 2.0 »).

## 37. Ce qu'on NE peut PAS faire à distance (limites honnêtes)

