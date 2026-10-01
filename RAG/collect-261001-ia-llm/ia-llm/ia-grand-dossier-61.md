---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-61
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Moonshot", "OpenAI", "United States", "vLLM"]
dates: ["2025-09-01", "2026-08-02", "2026-09-27", "2026-12-02"]
keywords: ["agents", "chatgpt", "claude", "deepseek", "kimi", "open source", "qwen", "rlhf", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5192, 5321]
sha256: a00dd7e048f4dd46582cf3d04b4bd784e2a4b4ef405c6e3ce0338f443b1a57bd
---

# IA — Le grand dossier

**Étapes.**
1. Choisir 3 scénarios fermés (disque plein, service web down, certificat expirant).
2. Pour chacun : runbook écrit et validé par toi (le LLM ne *crée* pas la procédure, il la
   *retrouve et l'adapte*).
3. Intégration : alerte → runbook → plan → validation → exécution → ticket.
4. Garde-fous : dry-run systématique, rollback documenté, périmètre Ansible limité aux
   hôtes taggés `auto-remediation-ok`.
5. Métriques : MTTR avant/après, taux d'escalade humaine, incidents causés par l'automation
   (zéro toléré au début).

**Budget.** 0 € logiciel ; 2-3 semaines.

**Risques.** C'est le projet où le « lethal trifecta » est maximal (données sensibles +
contenu externe + actions) → c'est aussi celui où les garde-fous doivent être les plus
stricts. **Ne jamais commencer par celui-ci** : fais les projets 1-4 d'abord, ils construisent
les briques (RAG, whitelist, validation humaine, métriques).

### Tableau de synthèse des 5 projets

| Projet | Difficulté | Durée | Prérequis | Valeur |
|---|---|---|---|---|
| 1. RAG local | ★☆☆☆☆ | 1 week-end | VM + Ollama | Base de tout le reste |
| 2. Astreinte lecture seule | ★★☆☆☆ | 1-2 jours | Projet 1 | Diagnostic 24/7 sans risque |
| 3. Prédictif onduleurs | ★★★☆☆ | 2-4 j + 6 mois données | NUT + TSDB | Ton métier, valeur directe |
| 4. Tri tickets GLPI | ★★★☆☆ | 3-5 jours | Export GLPI | Gain de temps niveau 1 |
| 5. Supervision augmentée | ★★★★☆ | 2-3 semaines | Projets 1-4 | MTTR réduit, à faire en dernier |

**Ordre conseillé : 1 → 2 → 4 → 3 → 5.** Chaque projet finance la confiance pour le suivant —
c'est comme ça qu'on passe du « pilote qui échoue » (MIT NANDA, 95 %) au « déploiement qui
tient » : petit périmètre, métriques, humain dans la boucle, itération.

---

## 15. Annexe G — FAQ : les questions qu'on te posera (et leurs réponses)

> Les questions que ta direction, ton équipe ou tes collègues te poseront sur l'IA — avec
> des réponses courtes, sourcées, sans langue de bois. Utile en réunion.

**« L'IA va-t-elle supprimer nos emplois ? »**
Personne ne sait avec certitude. Ce qu'on sait : Dario Amodei (Anthropic) prédit 50 % des
emplois de bureau débutants supprimés en 1-5 ans ; son propre économiste en chef (Peter
McCrory, juillet 2026) montre qu'à ce jour le chômage US est à 4,2 % sans impact mesurable
de l'IA. Le Forum économique mondial (2025) prévoit 92 M d'emplois déplacés mais 170 M créés
d'ici 2030. Réponse honnête : les *tâches* changent plus vite que les *métiers* ; la
formation est la variable d'ajustement. (Voir 1.3.)

**« Les gains de productivité sont-ils réels ? »**
Oui, mais hétérogènes : +14 % en support client (Brynjolfsson et al., 2023), jusqu'à +60 %
avec des agents (Aral & Ju, MIT, 2025), concentrés sur les débutants et les tâches de
rédaction/code. Mais 95 % des pilotes d'entreprise n'atteignent pas d'impact financier
mesurable (MIT NANDA, 2025) — l'écart vient de l'organisation, pas du modèle. Exige une
baseline mesurée avant tout investissement. (Voir 1.2.)

**« Peut-on mettre nos documents dans une IA sans risque ? »**
Pas dans un outil grand public sans contrat d'entreprise (risque de réutilisation pour
l'entraînement + fuite). En instance dédiée (ex. : Claude/ChatGPT Team/Enterprise avec
exclusion d'entraînement) ou en local (Ollama), oui — avec ACL par document, journalisation
et pas de secrets dans les prompts. Le risque n°1 n'est pas le modèle, c'est l'absence de
contrôle d'accès au niveau du retrieval. (Voir 2.1.3.)

**« L'open source, c'est moins sûr que les modèles propriétaires ? »**
Ça dépend de quoi on parle. Côté *capacités*, l'écart s'est réduit (DeepSeek, Qwen, Kimi).
Côté *sécurité d'usage*, l'open source auto-hébergé garde tes données chez toi (bon pour la
confidentialité) mais c'est à toi de gérer les garde-fous (pas de filtres intégrés par
défaut). Côté *régulation*, la Chine ne fait pas de différence (mêmes obligations) et l'UE
a un régime allégé pour l'open source sauf risque systémique. Il n'y a pas de réponse
universelle : c'est un arbitrage données/coûts/compétences. (Voir vol. IA + 2.4.)

**« Faut-il interdire l'IA dans l'entreprise en attendant d'y voir clair ? »**
Non — l'interdiction crée du shadow AI (usages non déclarés, sans garde-fous), pire que
l'usage encadré. Mieux : inventaire des usages, charte simple, instance d'entreprise,
formation. (Voir 6, conseil n°13, et 12.4.)

**« L'IA consomme-t-elle vraiment autant d'électricité qu'on le dit ? »**
Les data centers mondiaux : ~565 TWh en 2026 (+26 %/an, Gartner), ~1,5 % de l'électricité
mondiale aujourd'hui, ~3 % projetés en 2030 (AIE). Le problème n'est pas la part mondiale,
c'est la *concentration* : files de raccordement de 5 ans, 40 % des data centers IA sous
contrainte électrique d'ici 2027. Et 31 % de cette consommation vient déjà des serveurs IA.
(Voir 1.5.)

**« Un chatbot doit-il dire qu'il est une IA ? »**
Oui, dans l'UE depuis le 2 août 2026 (article 50 de l'AI Act) — c'est du droit applicable,
avec amendes possibles. Même hors UE, c'est une bonne pratique (confiance, traçabilité).
(Voir 2.4.1.)

**« Les images générées par IA doivent-elles être marquées ? »**
UE : oui depuis le 02/08/2026 (art. 50), filigrane machine-readable au 02/12/2026 pour les
systèmes existants. Chine : oui depuis le 01/09/2025 (double marquage visible + métadonnées,
sanctions réelles). US : pas d'obligation fédérale générale, mais le TAKE IT DOWN Act
criminalise les deepfakes intimes non consentis. (Voir 2.2.1, 2.4.)

**« Que vaut la "Constitutional AI" d'Anthropic ? »**
C'est une méthode d'alignement réelle et publiée (2022) : le modèle se corrige lui-même selon
des principes écrits et auditable — mieux que le RLHF pur sur la scalabilité et la
transparence des principes. Mais ça ne résout ni les jailbreaks ni la question politique
(qui écrit la Constitution ?). C'est un progrès d'ingénierie, pas une garantie. (Voir 3.2.)

**« Par quoi commencer concrètement ? »**
Le projet 1 de l'annexe F : un RAG local sur tes documents avec Ollama, en un week-end,
0 €. Puis la check-list des 20 conseils (section 6) : 2 newsletters, un créneau hebdo de
45 minutes, les releases GitHub des outils que tu utilises. La régularité bat l'intensité.

**« Comment répondre à un commercial qui vend "l'IA magique" ? »**
Trois questions : (1) Quelles métriques avant/après sur *mon* cas d'usage ? (2) Qui
l'exploite en production, depuis quand ? (3) Où vont mes données (entraînement ? sous-traitants ?).
Sans réponses écrites : ne pas signer. (Voir 1.7, règle d'or.)

**« Et si je n'ai qu'une heure par mois pour la veille IA ? »**
Lis Import AI (hebdo, Jack Clark) et survole les releases GitHub de vLLM/Ollama/LangChain.
Une heure bien investie par mois avec ce filtre vaut mieux que dix heures de flux X.
(Voir 5.2.)

**« L'IA va-t-elle rendre les sysadmins obsolètes ? »**
Les tâches d'exécution répétitives (déploiements standard, tri de tickets, rédaction de
comptes-rendus) seront de plus en plus assistées/automatisées. Mais l'architecture, la
sécurité, la responsabilité en cas de panne, la compréhension du métier — ça ne se délègue
pas à un modèle. Le sysadmin qui pilote des agents vaut plus que celui qui les subit :
c'est exactement ce que ce dossier t'apprend à devenir. (Voir 1.3, 4.4, annexe F.)

---

*Fin de la partie 3 — 27/09/2026. 15 sections : débat (1), dangers (2), Claude (3), stacks (4),
communautés (5), conseils (6), glossaire (7), quiz (8), scénarios sécu (9), tableau de bord (10),
énergie (11), conformité AI Act (12), mémos (13), projets (14), FAQ (15).*
*Maintenance : revérifier 1.5, 2.4, 10.1 tous les 6 mois.*

---

## 16. Journal de maintenance du dossier

