---
id: collect-261001-general-networking/general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026-1
title: "macOS et Linux"
domain: general-networking
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["acquisition", "agents", "benchmarks", "open source"]
source: docs/RAG/collect-261001-general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 54]
sha256: 7c7784f02dd0e2ee507004b4984e8968fadde799dfea14a681595b6d302044f4
---

# macOS et Linux

*Dernière mise à jour : 19 août 2026.* Si vous écrivez du Python en 2026, le gestionnaire de paquets **uv Python** est probablement l’outil le plus important à adopter cette année. Développé par la société **Astral** – les créateurs du linter Ruff – et écrit intégralement en Rust, uv se présente comme « un gestionnaire de paquets et de projets Python extrêmement rapide ». Concrètement, il remplace à lui seul `pip`, `pip-tools`, `pipx`, `poetry`, `pyenv`, `twine` et `virtualenv`, tout en étant de **10 à 100 fois plus rapide** que pip selon la documentation officielle d’Astral – un écart que le tutoriel UV publié par Han en décembre 2025 chiffre même, sur certains flux de travail, à **10 à 1000 fois plus rapide** que les outils Python traditionnels. La cadence de publication d’Astral s’est encore accélérée depuis : après la 0.12.2 du 5 août 2026, l’éditeur a livré la 0.12.6 le 25 août 2026, puis **uv 0.12.7**, la dernière version en date, le 27 août 2026.

Le sujet est d’autant plus brûlant que, le 19 mars 2026, **OpenAI a annoncé le rachat d’Astral** – et donc de uv, Ruff et de l’outil de typage ty. L’équipe rejoint le projet Codex d’OpenAI, mais s’est engagée à poursuivre le développement de ces outils en open source sous licence permissive, comme le rapporte le développeur Simon Willison. Autrement dit : uv n’est pas près de disparaître, bien au contraire.

Dans ce tutoriel, vous allez installer uv, gérer plusieurs versions de Python, puis **construire un projet Python complet et fonctionnel** de A à Z : une application en ligne de commande baptisée `pypeek` qui interroge l’API publique de PyPI pour afficher la dernière version d’un paquet. Vous verrez les dépendances, le verrouillage reproductible, les tests, le linting, la construction d’un paquet, la publication, la conteneurisation Docker et l’intégration continue. Comptez environ **40 minutes** pour les 12 étapes. Aucune installation préalable de Python n’est nécessaire : uv s’en charge.

## uv, le gestionnaire de paquets Python qui bouscule pip et Poetry

Pendant des années, l’écosystème Python a souffert d’une fragmentation de son outillage. On installait les paquets avec pip, on isolait les environnements avec virtualenv ou venv, on gérait les versions de l’interpréteur avec pyenv, on verrouillait les dépendances avec pip-tools ou Poetry, on lançait les outils CLI avec pipx, et on publiait sur PyPI avec twine. Sept outils, sept syntaxes, sept sources de bugs.

uv unifie tout cela derrière une seule commande. Le dépôt astral-sh/uv sur GitHub dépasse aujourd’hui les **87 000 étoiles**, ce qui en fait l’un des projets d’outillage Python les plus populaires jamais publiés. Il s’installe comme un binaire Rust autonome, sans nécessiter Python ni Rust au préalable, et fonctionne sur macOS, Linux et Windows.

Astral n’est pas un projet amateur : la société a levé **4 millions de dollars** en amorçage en 2023, tour mené par le fonds Accel avec la participation de figures comme Guillermo Rauch (Vercel) et Solomon Hykes (Docker). uv a été testé à grande échelle contre les **10 000 paquets PyPI les plus téléchargés** pour garantir sa compatibilité. C’est cette rigueur qui explique son adoption fulgurante dans les entreprises et les projets open source.

## OpenAI rachète Astral : ce que cela change pour uv (mars 2026)

Le 19 mars 2026, OpenAI a officialisé l’acquisition d’Astral. Le montant de la transaction n’a pas été communiqué. L’équipe d’Astral rejoint la division Codex d’OpenAI, l’assistant de programmation de l’entreprise, ce qui laisse deviner l’ambition : faire de uv et Ruff des briques centrales de l’expérience de développement Python assistée par IA.

La question qui a immédiatement agité la communauté était : uv va-t-il rester libre ? La réponse, réaffirmée par Astral et OpenAI, est oui. Les deux parties se sont engagées à maintenir uv, Ruff et ty en open source, sous une licence permissive qui autorise le fork si nécessaire. Pour les équipes qui ont déjà misé sur **uv Python** en production, il n’y a donc aucun changement de licence à craindre à court terme. Cette continuité est précisément ce qui rend le moment idéal pour adopter l’outil : sa gouvernance est désormais adossée à l’un des acteurs les mieux financés du secteur.

Ce rachat s’inscrit dans une tendance de fond : l’outillage de développement devient un champ de bataille stratégique pour les entreprises d’IA. Après tout, la vitesse d’installation et la reproductibilité des environnements conditionnent directement la productivité des agents de code automatisés. uv, capable de recréer un environnement complet en quelques centaines de millisecondes, est un candidat naturel pour ce rôle.

## uv Python contre pip, Poetry et pyenv : le tableau des remplacements

Avant de plonger dans l’installation, il est utile de cartographier ce que uv remplace. Le tableau ci-dessous met en regard chaque outil traditionnel et son équivalent uv. C’est aussi une bonne feuille de route mentale : si vous connaissez déjà pip ou Poetry, vous retrouverez vos réflexes.

| Outil traditionnel | Rôle | Commande uv équivalente | 
|---|---|---|
| pip | Installer des paquets | `uv add` /`uv pip install` | 
| virtualenv / venv | Créer des environnements isolés | `uv venv` (automatique via`uv run` ) | 
| pyenv | Gérer les versions de Python | `uv python install` /`uv python pin` | 
| pip-tools (pip-compile) | Verrouiller les dépendances | `uv lock` /`uv pip compile` | 
| pipx | Exécuter des outils CLI isolés | `uvx` /`uv tool install` | 
| Poetry / PDM | Gérer projet + dépendances | `uv init` /`uv add` /`uv sync` | 
| twine | Publier sur PyPI | `uv publish` | 
| build | Construire sdist et wheel | `uv build` | 

L’avantage n’est pas seulement la vitesse : c’est la cohérence. Toutes ces commandes partagent le même cache global, le même résolveur de dépendances et le même fichier de verrouillage. Vous n’avez plus à vous demander si Poetry et pip voient les mêmes versions – il n’y a plus qu’une seule source de vérité.

## Benchmarks : à quel point uv est-il plus rapide que pip ?

La promesse d’Astral, « 10 à 100 fois plus rapide que pip », se vérifie dans les mesures. Le facteur exact dépend surtout de l’état du cache : à froid, le gain est déjà net ; à chaud, il devient spectaculaire, car uv déduplique les paquets déjà téléchargés et se contente de créer des liens durs (hardlinks) plutôt que de recopier les fichiers.

| Opération | Outil classique | Accélération avec uv | 
|---|---|---|
| Résolution + installation (cache froid) | pip | 8 à 10× | 
| Réinstallation (cache chaud) | pip | 80 à 115× | 
| Création d’un environnement virtuel | `python -m venv` | ~80× | 
| Création d’un environnement virtuel | virtualenv | ~7× | 

Ces chiffres proviennent des benchmarks publiés par Astral et confirmés par des tests indépendants. En pratique, cela transforme le flux de travail : recréer un environnement propre en intégration continue passe de plusieurs dizaines de secondes à moins d’une seconde. Sur un monorepo avec des centaines de dépendances, l’écart cumulé se compte en heures d’ingénierie économisées chaque mois – une étude de Techplained publiée en avril 2026 chiffre ce gain à environ **58 heures par an** de temps de CI économisé face à pip, sur un échantillon de **10 400 exécutions** avec un cache mixte.

## Prérequis et versions

