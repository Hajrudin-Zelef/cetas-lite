---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-21
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["exploit", "sandbox"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1908, 1986]
sha256: 99ee00fda1d26f03d221db235b579ad138d2d7cddd782dc310d79febda71a8f4
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

**Mise en œuvre** : (1) NCE expose ses alarmes via l'**API** (section 90) ; un script (ou connecteur) les injecte dans Zabbix (ou Zabbix interroge l'API périodiquement). (2) Dans Zabbix : corrélation « switch down + onduleur sur batterie sur le même site » = cause probable électrique. (3) Zabbix supervise **NCE lui-même** (services, disque, certificats — section 91).

**Leçon** : NCE ne remplace pas la supervision transverse — il l'**alimente**. L'architecture « NCE = réseau Huawei, Zabbix = transverse, GLPI = tickets » donne à chacun son rôle sans doublon.

## 147. Cas pratique 24 — PPSK pour les invités d'un hôtel

**Contexte** : un site de type hôtelier veut un Wi-Fi invité simple mais traçable, sans portail compliqué.

**Mise en œuvre** : (1) SSID `INVITES`, VLAN dédié, isolation clients, débit plafonné. (2) **PPSK** : une clé par chambre (ou par lot de chambres), générée dans NCE, imprimée sur la fiche d'accueil, **à durée limitée** (durée du séjour). (3) Révocation au départ (ou expiration automatique). (4) Journalisation des connexions (obligations locales — **à vérifier**).

**Leçon** : le PPSK est le bon compromis entre le PSK partagé (qui fuit) et le portail (lourd) — à condition d'en gérer le **cycle de vie** (création, expiration, révocation), ce que NCE fait nativement.

## 148. Cas pratique 25 — Coupure électrique et reprise (lien onduleurs)

**Contexte** : coupure secteur sur un site ; l'onduleur tient 30 minutes ; le groupe ne démarre pas.

**Ce qui se passe côté NCE** : les équipements s'éteignent par zone (alarmes « hors ligne » en cascade) ; à la reprise, ils **redémarrent** et se reconnectent au contrôleur (qui n'a rien perdu — il était sur onduleur+groupe au central, évidemment).

**Conduite** : (1) Côté énergie : procédure onduleur/groupe (le métier de Zelef — voir son guide onduleurs). (2) Côté NCE : **ne rien toucher** pendant la coupure (les alarmes sont le symptôme, pas le problème) ; à la reprise, vérifier la **resynchronisation** (tous les équipements « normal », conformité OK — un redémarrage brutal peut révéler un firmware corrompu). (3) REX croisé énergie/réseau.

**Leçon** : la **priorité d'extinction/allumage** se pense à l'avance (d'abord les charges non critiques, les équipements réseau en dernier à l'extinction et en premier à la reprise si possible) — et le NCE du site central doit être sur la chaîne ondulée la plus protégée (voir cas 11).

---

# PARTIE 15 — LIMITES HONNÊTES

## 149. Limite 1 — Courbe d'apprentissage

NCE-Campus n'est pas « eSight en mieux », c'est un **nouveau métier** : concepts SDN (YANG, intent), templates versionnés, politiques 802.1X, ZTP, API, licences device-day. Compter **plusieurs semaines** de formation par administrateur + **2-3 mois** de pratique accompagnée avant l'autonomie — et un **référent** qui monte en expertise sur 6-12 mois.

Conséquence managériale : ne pas lancer NCE en même temps qu'un autre gros projet ; sanctuariser du temps de formation (pas « entre deux tickets ») ; accepter une phase de productivité réduite. Une équipe qui n'a pas le temps d'apprendre fera du NCE un eSight cher et compliqué.

## 150. Limite 2 — Dépendance au contrôleur (single point of management)

Centraliser le pilotage, c'est centraliser le risque **de gestion** (pas de trafic — le plan de données reste local, rappel section 22). Si NCE est indisponible : plus de supervision, plus de déploiement, authentification centralisée fragilisée.

Mitigations réelles (pas des slogans) : **HA** déployée et testée, **composants d'authentification locaux** sur les sites critiques, **backups testés**, supervision du contrôleur, procédures dégradées écrites. Avec ça, la dépendance est **gérée** — sans ça, c'est une bombe à retardement. Le dire en comité de pilotage n'est pas du pessimisme, c'est du professionnalisme.

## 151. Limite 3 — Coût des licences

Le modèle device-day en souscription, c'est un **abonnement à vie** : on paie chaque année pour continuer à gérer son propre réseau. Sur 5 ans, le TCO peut dépasser largement un eSight perpétuel — surtout si le parc est stable (l'automatisation ne « rembourse » pas grand-chose sur un réseau qui ne change pas).

Points de négociation : durée (3 ans souvent mieux que 1 an), cotermination, marge de croissance incluse, **clause de sortie** (que se passe-t-il si on ne renouvelle pas ? — à écrire noir sur blanc), prix des extensions. Et surtout : **modéliser le TCO avant de signer** (section 121), pas après.

## 152. Limite 4 — Maturité sur petit site (eKit suffit)

Sur un petit site (quelques switches, une poignée d'AP, besoins simples), NCE-Campus est **surdimensionné** : coût, complexité et formation pour un bénéfice marginal — eKit fait le travail pour une fraction du prix et de l'effort.

Le vrai risque : déployer NCE « pour voir » sur un petit site, le sous-exploiter, et en conclure que « NCE ne sert à rien » — alors que c'est le **périmètre** qui était mauvais, pas l'outil. NCE se juge sur un périmètre à sa mesure (multi-sites, 802.1X, automatisation) — partout ailleurs, eKit/eSight sont des choix **respectables**, pas des aveux d'échec.

## 153. Limite 5 — Écosystème et documentation

Constats honnêtes :

- La documentation est **abondante mais inégale** : brochures marketing nombreuses, guides techniques parfois ardus, versions qui se succèdent vite (R020/R022/R024...) — toujours travailler sur la **doc de la version exacte** déployée.
- La **communauté** francophone/utilisateurs est plus petite que celles de Cisco/Aruba — moins de retours d'expérience, moins de scripts partagés.
- L'écosystème API existe (500+ API, sandbox) mais demande un investissement pour être exploité.
- Le support passe largement par le **partenaire/intégrateur** : sa qualité fait la moitié du projet — le choisir avec autant de soin que le produit.

## 154. Limite 6 — Dépendance au support Huawei/partenaire

NCE-Campus n'est pas un produit qu'on achète et qu'on oublie : versions fréquentes, matrices de compatibilité à suivre, licences à renouveler, évolutions de l'écosystème. Cela crée une **dépendance durable** au constructeur et à l'intégrateur.

La gérer : exiger la **documentation** de tout ce qui est fait par l'intégrateur (templates, procédures — pas de boîte noire), former un **référent interne** (jamais 100 % dépendant), maintenir un **contrat de support** avec des SLA écrits (délais de réponse/résolution), et garder une **porte de sortie** documentée (export des configurations, procédure de retour à une gestion locale si nécessaire — même si on ne s'en sert jamais, l'avoir change le rapport de force).

## 155. Synthèse honnête — faut-il y aller ?

**Oui, si** : parc Huawei campus significatif et/ou multi-sites, besoin de 802.1X/automatisation, équipe (petite mais) formable, budget projet avec TCO modélisé, sponsor direction. Dans ce cas, NCE-Campus est un vrai **saut capacitaire** : ZTP, templates, politiques centralisées, assurance — on change de division.

**Non (ou pas maintenant), si** : petit site stable (eKit), parc multi-vendeurs/énergie dominant (eSight/Zabbix), pas de budget projet, équipe saturée sans perspective de formation. Dans ce cas, NCE serait un **boulet** : cher, compliqué, sous-exploité — et c'est le produit qui serait blâmé pour une mauvaise décision de périmètre.

**La voie sage** : pilote sur un site représentatif (3-6 mois), REX honnête, décision go/no-go sur données réelles — pas sur des promesses commerciales. Et dans tous les cas : eSight reste **légitime** sur ses périmètres (supervision transverse, énergie) — la coexistence n'est pas un échec, c'est une architecture.

---

# PARTIE 16 — QUIZ

## 156. Quiz — 10 questions

