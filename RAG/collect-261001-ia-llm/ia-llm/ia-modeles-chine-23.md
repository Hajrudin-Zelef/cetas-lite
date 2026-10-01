---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-23
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "LongCat", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Xiaomi", "Z.ai"]
dates: []
keywords: ["agent", "agents", "apache", "arr", "asic", "attention", "benchmark", "benchmarks", "claude", "cyber", "deepseek", "embedding"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [2290, 2396]
sha256: cb4f617ef852edf070a22d7c9469b05ef74873ba086846e8291ec242f16bc438
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

Le dosage « couches linéaires vs attention standard » varie selon les modèles Qwen, et c'est un paramètre
de design à part entière : **Qwen3.6-35B-A3B** — 40 couches en 10 groupes de (**3× Gated DeltaNet + 1×
Gated Attention**), soit 75 % linéaire ; **Qwen3-Coder-Next** — **75 % linéaire / 25 % full** (même ratio,
explicite) ; **Qwen3.8-27B** — **48 couches GDN + 16 couches full-attention** (même 75/25). La constance
du ratio 3:1 à travers trois modèles suggère un optimum empirique Alibaba : trois couches linéaires
bon marché pour le « gros œuvre » du contexte, une couche d'attention pleine pour la précision des
dépendances fines. C'est de l'ingénierie de compromis quantifiée — et c'est ce ratio, plus que le nombre
de paramètres, qui fait l'économie du 256K–1M chez Qwen.

## 201. Kimi : le vocabulaire 163 840 et la tokenisation

Kimi K2.6 utilise un vocabulaire de **163 840 tokens** — entre MiniMax-M2 (200 064) et les ~250K de Qwen.
La taille du vocabulaire est un levier sous-estimé : un gros vocabulaire fragmente moins les langues
non-anglaises (le chinois notamment, où chaque caractère peut être un token) et le code, ce qui réduit
le nombre de tokens par document — donc le coût et la latence à qualité égale. Le revers : une matrice
d'embedding plus grosse (163 840 × hidden size) qui pèse dans les poids et la mémoire. Moonshot a aussi
fait le choix du **natif INT4** (pas de release haute précision) : le modèle est distribué directement
dans sa forme servable, pas dans sa forme d'entraînement. Philosophie : le poids publié est un artefact
de déploiement, pas un artefact de recherche.

## 202. GLM : anatomie du positionnement « cyber »

Le dossier « cyber » du GLM-5.3 mérite le détail car c'est une première assumée : post-training avec
**environnements de découverte de vulnérabilités**, capacités revendiquées — analyse de vulns, exploit
scripting, élévation de privilèges, chaînes d'exploitation multi-étapes — et chiffres : **2 436
vulnérabilités trouvées sur 269 projets open-source, 53 CVE**. Benchmarks : CyberGym 77,2 % → **84,5 %**,
ExploitBench 24,4 % → **54,4 %** (×2,2), Terminal-Bench 3.0 4,6 % → 28,3 % (×6). Z.ai a retenu les poids
deux semaines pour durcissement sécurité avant publication (28 août). Trois lectures : (1) le post-training
spécialisé peut multiplier une capacité par 2–6 sans changer l'architecture (même base 5.2) ; (2) le
dual-use offensif est désormais un argument marketing ; (3) la « retenue sécurité » devient un précédent
de gouvernance des releases ouvertes. Tous les chiffres sont vendor-reported : l'ampleur réelle reste à
mesurer indépendamment.

## 203. MiniMax : l'hallucination divisée par 2,6

Le progrès le plus spectaculaire de la série M2 n'est pas un benchmark de code, c'est la **chute du taux
d'hallucination** : de **88 % (M2.5) à 34 % (M2.7)** sur AA-Omniscience (Artificial Analysis) — meilleur
que Claude Sonnet 4.6 (46 %). En une génération (février → mars 2026), MiniMax a divisé par 2,6 le taux
d'hallucination de son modèle ouvert. C'est le genre de métrique qui compte plus que SWE-bench pour un
RAG : un modèle qui n'invente pas est un modèle qu'on peut brancher sur des documents d'entreprise. Le
moteur probable : le « self-evolution harness » (self-feedback, self-optimization) et un post-training
orienté factualité. Corollaire : GDPval-AA **1495 ELO** (#1 open-source) et la parité avec GPT-5.3-Codex
sur SWE-Pro (56,22 %) font de M2.7 le « tout-terrain » ouvert le plus équilibré du printemps 2026 —
dommage que sa licence exacte reste non vérifiée.

## 204. Tencent : preserved reasoning et context caching (le détail qui tue)

Deux features Hy3 méritent qu'on s'y arrête car elles ciblent le vrai coût des agents : le **preserved
reasoning** (le raisonnement survit entre les appels d'outils — pas besoin de le régénérer à chaque tour)
et le **context caching** (le cache de contexte survit entre les tours). Dans une boucle agentique
classique, chaque appel d'outil force à renvoyer tout l'historique + à refaire le raisonnement : c'est
du token gaspillé et de la latence. Hy3 attaque les deux : raisonnement persistant + cache persistant +
structured output garanti + trois niveaux de thinking (`no_think` pour les appels triviaux, `think_high`
pour la planification). C'est une conception « système » de l'agent : le modèle est un composant
logiciel fiable avec des garanties d'interface, pas un oracle brillant mais imprévisible. Pour un
développeur d'agents, ces deux features valent plus qu'un point de benchmark.

## 205. Les chiffres qui résument 2026 (table)

| Record / fait | Valeur | Modèle |
|---|---|---|
| Plus gros total open-weight | 2,8T paramètres | Kimi K3 |
| Plus gros actifs/token | ~104B (divergence : ~50B) | Kimi K3 |
| Plus petit actifs/token « sérieux » | 1,3B | Ling-3.0-tiny (7,9B total) |
| Input le moins cher | $0,075 / 1M | Ling-3.0-flash |
| Plus grand ratio total/actifs | ~27× | Kimi K3 (2,8T/104B) |
| Plus gros téléchargement | 1,4–1,56 To (96 shards) | Kimi K3 |
| Contexte max standard | 1M (12 modèles du volume) | — |
| RL le moins cher documenté | ~$3,47M | MiMo-V2.6 |
| Plus longue vie d'un Flash | 4,5 mois (retiré) | DeepSeek V4-Flash |
| Plus longue tâche agentique revendiquée | 35 heures | Qwen3.7-Max |
| Plus gros essaim d'agents revendiqué | 300 agents / 4 000 étapes | Kimi K2.6 |
| Plus forte sparsité d'experts | 16/896 | Kimi K3 |
| Plus forte chute d'hallucination | 88 % → 34 % (une génération) | MiniMax M2.5 → M2.7 |
| Plus gros cluster domestique revendiqué | ~50 000 ASIC | LongCat-2.0 |
| Plus gros serving domestique revendiqué | 100 000 puces chinoises | GLM-5.3-Flash |

## 206. Note finale : ce volume est une photographie

Ce volume fige au 27 septembre 2026 un paysage qui bouge chaque semaine : DeepSeek V4.1 Pro est attendu
d'un jour à l'autre, Qwen4 est teasé, la licence du Kimi K3 reste à confirmer, la controverse
Anthropic–Xiaomi n'est pas tranchée. Les fiches « non vérifié » ne sont pas des échecs de recherche,
ce sont des **positions ouvertes** : chaque mention est une question à reposer dans trois mois. La
méthode (aucune spec inventée, sources tracées, divergences signalées) vaut pour les volumes suivants :
mieux vaut une encyclopédie honnête et incomplète qu'une encyclopédie complète et fausse. Rendez-vous au
volume 2 — l'Occident (NVIDIA Nemotron, NousResearch Hermes) — puis au volume 3, le choc des deux
hémisphères.

---

*Fin du volume 1 — Encyclopédie des modèles IA : la Chine (février 2026 → 27 septembre 2026).*
*Sections numérotées de 1 à 206. Vérification finale : `wc -l` ≥ 2500, 100+ sections, tableaux par famille, glossaire 32 termes, quiz 15 questions + réponses.*

## 207. Qwen : l'open-weight comme arme géopolitique douce

La valse d'Alibaba (tout-ouvert → frontier-closed → réouverture du 3.8-Max) n'est pas qu'une hésitation
commerciale : c'est de la **géopolitique par les poids**. Chaque release Apache 2.0 (Qwen3.5 intégral,
3.6-35B, 3.8-27B, Coder-Next) équipe gratuitement des milliers d'équipes dans le monde — y compris
occidentales — et crée une dépendance d'écosystème (harnais, fine-tunes, benchmarks calibrés sur Qwen).
La parenthèse « frontier-closed » (avril–juillet 2026) a testé la monétisation directe ; sa fin (août)
suggère que la valeur de l'influence open-weight dépasse la rente API du flagship. Même logique chez
Xiaomi (RL public, 7 000 environnements publiés) et Zhipu (MIT sans restriction territoriale, contrairement
au H3 de MiniMax). En 2026, **publier des poids ouverts est un acte de politique étrangère technologique**
— et la Chine est le seul acteur à le faire à l'échelle frontier, systématiquement.

## 208. Récapitulatif : les sorties mois par mois

