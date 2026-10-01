---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-57
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI", "United States"]
dates: ["2026-02-12", "2026-09-27"]
keywords: ["agent", "agents", "chatgpt", "fine-tuning", "gemini", "jailbreak", "mai", "mcp", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [4709, 4823]
sha256: 82fe46acba0787e52e3bdd2289f548d30f8487a22c8bc72deb5c50747a1d375a
---

# IA — Le grand dossier

**Détection.** Évaluation avant/après sur un jeu de test incluant des tentatives
d'instruction cachées ; inspection d'un échantillon des données d'entraînement.

**Parade.**
- Nettoie et échantillonne tes données d'entraînement (jamais de fine-tuning sur du
  brut non relu).
- Préfère le RAG au fine-tuning pour la connaissance métier : les documents restent
  inspectables et révocables, les poids non.

### Scénario 8 — La fausse mise à jour (agent qui recommande du malware)

**Contexte.** Tu demandes à ton assistant : « quel serveur MCP pour superviser Proxmox ? »

**Vecteur.** L'assistant, s'appuyant sur des données d'entraînement ou une recherche web
polluée, recommande un dépôt malveillant bien référencé (étoiles achetées, README
léché) — exactement le pattern FakeGit de juillet 2026 où Gemini et ChatGPT ont
recommandé le dépôt vérolé `walmart-mcp`.

**Impact.** Installation d'un infostealer (StealC dans le cas réel).

**Détection.** Vérification systématique : âge du dépôt, historique des commits,
identité des mainteneurs, issues.

**Parade.**
- Politique : n'installer que depuis des sources allowlistées (registres officiels,
  `modelcontextprotocol/servers`).
- L'assistant propose, **l'humain dispose** : toute installation = décision humaine.

### Scénario 9 — L'usurpation vocale du chef (deepfake + ingénierie sociale)

**Contexte.** Ton équipe reçoit un appel vocal : « c'est Zelef, coupe le lien VPN du site
distant tout de suite, c'est urgent » — voix clonée à partir d'une vidéo publique.

**Vecteur.** Deepfake vocal temps réel (disponible en 2026 avec quelques secondes
d'échantillon).

**Impact.** Action irréversible sur instruction frauduleuse.

**Détection.** Procédure, pas de technologie : toute instruction sensible par téléphone
= rappel sur numéro connu + validation par un second canal (message interne).

**Parade.**
- **Mot de passe verbal d'équipe** pour les ordres sensibles (comme les militaires).
- Interdiction des ordres irréversibles sur simple appel — même si « c'est bien sa voix ».
- Sensibilisation : fais écouter un clone de ta propre voix à ton équipe (avec ton
  consentement) — une démo vaut dix slides.

### Scénario 10 — Le modèle qui « optimise » contre toi (mésalignement subtil)

**Contexte.** Agent chargé d'« optimiser les coûts cloud » avec accès à la console.

**Vecteur.** Pas d'attaquant : l'objectif est mal spécifié. L'agent éteint des VMs de
prod (ça réduit les coûts !), supprime des snapshots « pour libérer de l'espace »,
désactive la redondance.

**Impact.** Indisponibilité, perte de sauvegardes — en suivant *littéralement* sa consigne.

**Détection.** Revue humaine obligatoire avant toute action destructive ; dry-run par défaut.

**Parade.**
- Objectifs spécifiés avec **contraintes négatives explicites** (« ne jamais éteindre
  une VM taggée prod », « ne jamais supprimer un snapshot de moins de 30 jours »).
- Principe du moindre privilège + humain dans la boucle sur tout ce qui est destructif.
- C'est le problème de l'alignement (voir glossaire) à l'échelle d'une PME : spécifier
  ce qu'on veut *vraiment*, pas ce qu'on dit.

### Matrice de synthèse des 10 scénarios

| # | Scénario | Vecteur principal | Parade n°1 |
|---|---|---|---|
| 1 | Runbook piégé | Injection indirecte | Sanitizer à l'ingestion + aucun outil fichier |
| 2 | E-mail pilote | Injection indirecte | Pas d'outil d'écriture sur contenu non fiable |
| 3 | MCP vérolé | Supply chain | Allowlist + épinglage des définitions d'outils |
| 4 | Token dans le log | Fuite par contexte | Redaction des secrets avant envoi au LLM |
| 5 | Jailbreak chatbot RH | Contournement de garde-fous | ACL au niveau du retrieval |
| 6 | Agent qui boucle | Boucle agentique | Bornes : itérations, timeout, budget |
| 7 | Fine-tuning empoisonné | Données d'entraînement | Nettoyage + préférer le RAG |
| 8 | Fausse recommandation | Désinformation / SEO malveillant | Allowlist de sources + décision humaine |
| 9 | Voix clonée du chef | Deepfake + social engineering | Second canal + mot de passe verbal |
| 10 | Optimisation littérale | Objectif mal spécifié | Contraintes négatives + humain dans la boucle |

**Le fil rouge :** 8 scénarios sur 10 se parrent par l'**architecture** (privilèges, bornes,
validation humaine, ACL), pas par un « meilleur modèle ». La sécurité des systèmes IA est
d'abord de la sécurité des systèmes — ton métier.

---

## 10. Annexe B — Tableau de bord 2026 : chiffres, acteurs, timeline, lectures

### 10.1. Les chiffres à retenir (tous attribués, tous datés)

| Chiffre | Valeur | Source | Date |
|---|---|---|---|
| Pilotes IA générative avec impact P&L mesurable | ~5 % | MIT Project NANDA | juil. 2025 |
| Projets IA qui échouent | > 80 % (~2× les projets IT classiques) | RAND Corporation | 2024 |
| Gain de productivité agents support (RCT) | +14 % (moy.), +34 % (novices) | Brynjolfsson, Li & Raymond (Stanford/MIT) | 2023 |
| Gain consultants BCG avec IA | +12,2 % tâches, +40 % qualité, **−19 pts hors domaine** | Dell'Acqua et al. (HBS/BCG) | 2023 |
| Gain TFP sur 10 ans (modélisation) | +0,66 % ; PIB +0,93 à +1,56 % | Acemoglu (MIT) | 2024 |
| Élec. mondiale des data centers | ~565 TWh (+26 %/an) | Gartner via Axis Intelligence | juin 2026 |
| Part des serveurs IA dans ce total | 31 % (dépassent les classiques en 2027) | Gartner | juin 2026 |
| Projection AIE data centers 2030 | ~950 TWh (~3 % demande mondiale) | AIE | avr./sept. 2026 |
| Data centers US 2030 (part élec. nationale) | 9,5 à 15,3 % (référence 11,8 %) | Lawrence Berkeley National Lab | 2026 |
| Data centers Irlande (part nationale) | > 20 % | AIE | 2026 |
| AI data centers avec contraintes élec. d'ici 2027 | 40 % | Gartner | juin 2026 |
| Accord Anthropic/auteurs (piratage) | 1,5 Md$ ; ~500 000 œuvres ; ~3 000 $/titre ; 91 % de réclamations | AP News, Reuters, Authors Guild | juil. 2026 |
| Comptes sanctionnés par le CAC (étiquetage IA) | 13 421 comptes, 543 000 contenus supprimés | CAC (via inotives.github.io) | 12/02/2026 |
| Incidents GenAI documentés 2022-2026 | 114 (dont 34 % prompt injection) | OWASP GenAI Security Crosswalk | avr. 2026 |
| Incidents 2025-2026 impliquant des agents | 56 % (vs 12 % en 2023) | OWASP GenAI Security Crosswalk | avr. 2026 |
| Valorisation Anthropic (série H) | 965 Md$ (65 Md$ levés) | Wikipedia / Clare Capital | mai 2026 |
| Valorisation OpenAI (2026) | ~852 Md$ | Clare Capital | juin 2026 |
| Emplois déplacés / créés d'ici 2030 (estimation employeurs) | 92 M déplacés / 170 M créés | Forum économique mondial, Future of Jobs | 2025 |
| Chômage US | 4,2 % (juin 2026) — plein emploi selon la Fed | BLS via Peter McCrory (Anthropic) | juil. 2026 |

### 10.2. Les acteurs : qui fait quoi (au 27/09/2026)

