---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-22
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1987, 2034]
sha256: 1b8f0d1842eb87f432dcb656df3b44b41bb4dc7e146f301502d7ffa3699f2f6f
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

**Q1.** Quelles sont les trois fonctions intégrées par iMaster NCE-Campus (qui le distinguent d'un NMS traditionnel comme eSight) ?

**Q2.** Citez trois protocoles southbound utilisés par NCE-Campus pour dialoguer avec les équipements, et le rôle de chacun.

**Q3.** Qu'est-ce qu'un « device-day » dans le modèle de licence NCE, et que se passe-t-il quand le compteur passe en négatif ?

**Q4.** Décrivez la séquence complète d'un onboarding ZTP via DHCP (5 étapes).

**Q5.** Un AP361 branché n'apparaît pas dans NCE. Citez les 4 niveaux de diagnostic dans l'ordre, du physique au logique.

**Q6.** Pourquoi faut-il désenrôler un AP d'eKit avant de l'onboarder dans NCE ? Que faire s'il « retourne » vers eKit ?

**Q7.** En 802.1X, quels sont les trois rôles (supplicant, authentificateur, serveur) et quel équipement Huawei joue typiquement l'authentificateur en filaire ?

**Q8.** Un template appliqué en production coupe l'accès d'un site. Quelle est la première action correcte, et quelles sont les trois mesures préventives qui auraient dû être en place ?

**Q9.** Le contrôleur NCE devient injoignable. Le réseau s'arrête-t-il ? Qu'est-ce qui s'arrête vraiment, et quelle brique permet aux branches de continuer à authentifier localement ?

**Q10.** Citez trois périmètres pour lesquels eSight reste plus pertinent que NCE-Campus, et justifiez en une phrase chacun.

## 157. Quiz — réponses détaillées

**R1.** **Management** (gestion : inventaire, supervision, alarmes — le rôle NMS), **Control** (contrôle SDN : orchestration des configurations/politiques via NETCONF/YANG, ZTP, templates), **Analysis** (analyse : télémétrie, big data/ML, détection d'anomalies, cause racine — avec CampusInsight en option avancée). C'est l'intégration des trois qui distingue NCE d'eSight (manager seul).

**R2.** **NETCONF/YANG** : configuration structurée (déploiement, vérification de conformité, notifications de changement). **SNMP** : supervision et télémétrie classique, gestion des équipements traditionnels, base des sondes SLA/NQA. **CAPWAP** : tunnels de contrôle/gestion entre le WAC et les AP AirEngine. (Acceptés aussi : HTTP/2 pour la télémétrie, HTTPS pour les portails/gestion web, TCP pour la synchro des composants d'authentification.)

**R3.** Un device-day = **un équipement géré pendant un jour** ; les licences sont poolées et consommées au fil du temps. Quand le compteur passe en négatif, une **période de grâce de 30 jours** s'enclenche (documentée) : le système continue de fonctionner, c'est le délai pour régulariser (racheter des licences) — pas pour ignorer le problème.

**R4.** (1) L'équipement neuf démarre et demande une IP au **DHCP**. (2) Le DHCP lui attribue une IP **et l'option** contenant l'adresse du contrôleur NCE. (3) L'équipement **contacte NCE et s'enregistre** (ESN vérifié — pré-déclaration/liste blanche). (4) NCE l'**affecte au site** et pousse le **template** (+ firmware si besoin). (5) L'équipement applique, redémarre si nécessaire, passe à l'état **normal** (vérification de conformité).

**R5.** Dans l'ordre : (1) **Énergie** — l'AP est-il alimenté ? (PoE 802.3af du S310 : LED, état PoE du port). (2) **IP** — a-t-il une adresse ? (DHCP fonctionnel ?). (3) **WAC** — voit-il le contrôleur WLAN ? (option DHCP/DNS, CAPWAP UDP/5246-5247). (4) **NCE** — le WAC est-il géré par NCE, l'AP remonte-t-il ? Toujours du physique vers le logique : 80 % des cas se règlent aux niveaux 1-2.

**R6.** Parce qu'un équipement ne peut être piloté que par **une seule** plateforme : tant qu'il est lié à eKit (cloud), il cherche à s'y enregistrer et refuse/ignore NCE. Il faut le **désenrôler proprement d'eKit**, le **réinitialiser** en configuration d'usine, puis l'onboarder dans NCE. S'il « retourne » vers eKit : vérifier que le désenrôlement est effectif côté cloud eKit (pas seulement local).

**R7.** **Supplicant** : le poste qui demande l'accès (client 802.1X). **Authentificateur** : le **switch** (ex : S310) — il relaie l'EAP et applique la décision (ouvre/bloque le port, VLAN dynamique). **Serveur d'authentification** : le RADIUS (composant NCE ou externe) — il décide en vérifiant les identifiants dans l'annuaire et renvoie les attributs d'autorisation.

**R8.** Première action : **rollback** vers la version précédente du template (ne pas improviser d'autres changements ; intervention console locale si le management est coupé). Trois mesures préventives : (1) **déploiement par vagues** (maquette → pilote → généralisation, jamais de big bang) ; (2) **fenêtre de maintenance** planifiée et communiquée ; (3) **template versionné, testé en maquette et relu par un pair**, avec procédure de rollback écrite et testée.

**R9.** **Non, le réseau ne s'arrête pas** : le plan de données est local aux équipements, qui continuent de commuter avec leur dernière configuration (le réseau se « fige », il ne tombe pas). Ce qui s'arrête : supervision, nouveaux déploiements/changements de politiques, authentification **centralisée**. La brique qui permet aux branches de continuer à authentifier localement est le **composant d'authentification** déployé sur les sites distants (jusqu'à 20, synchronisé avec le central via TCP).

**R10.** (1) **Supervision multi-vendeurs** : NCE ne gère pratiquement que du Huawei, eSight supervise les équipements tiers en SNMP. (2) **Énergie/onduleurs** : la supervision des UPS et équipements techniques via SNMP générique est le domaine d'eSight (ou d'un superviseur dédié) — NCE n'est pas un superviseur d'infrastructure technique. (3) **PON/OLT/ONT et data center** : périmètres couverts par eSight, hors scope de NCE-Campus. Bonus accepté : la **simplicité/coût** sur petit parc stable.

---

# PARTIE 17 — GLOSSAIRE

## 158. Glossaire

