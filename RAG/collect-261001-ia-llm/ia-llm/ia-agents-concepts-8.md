---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-8
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [1063, 1225]
sha256: 57010c195617fb7c4ea0248c1932ffd35c17cefd2c41f90db7a744a8c969d1e9
---

# Concepts : agents IA, agentic, autonomie

## 57. Open Interpreter existe-t-il encore ? (réponse : oui, et il a changé)

Oui. Vérifié fin septembre 2026 : le dépôt `openinterpreter/openinterpreter`
est actif (dernier commit relevé : 13 septembre 2026, ~68 000 étoiles GitHub).
Le projet a connu une évolution majeure : après la version Python historique
(`pip install open-interpreter`), le développement s'est orienté vers une
**réécriture en Rust** (version 0.0.43 évoquée dans les notes de release),
avec une interface « style Codex », une exécution de commandes en environnements
isolés (sandboxed), et une compatibilité avec l'**ACP (Agent Client Protocol)**
pour l'intégration aux éditeurs.

Ce que ça veut dire pour toi :
- le projet n'est pas mort, il est en pleine refonte ;
- la version Python classique reste utilisable et documentée, mais le futur
  du projet est la version Rust ;
- vérifie la version exacte au moment où tu l'installes (`interpreter --version`,
  changelog du dépôt) : en période de réécriture, la doc peut être en retard
  d'une version.

## 58. Open Interpreter : c'est quoi, concrètement

Open Interpreter, c'est : **tu parles en langage naturel dans ton terminal,
il écrit du code (Python, shell, etc.), il l'exécute sur ta machine, il te montre
le résultat, et il continue.**

Exemple de session (version Python historique) :

```
$ interpreter
> Convertis le fichier inventaire.csv en tableau Markdown trié par IP.
# l'agent écrit un script Python avec pandas, l'exécute, affiche le résultat
> Maintenant envoie-moi un résumé des 5 machines avec le moins de RAM libre.
# il réutilise le contexte, écrit un nouveau script, l'exécute
```

Différence avec un chatbot : il **agit** sur ta machine (fichiers, shell, navigateur
dans certaines configs). Différence avec un agent « cloud » : tout tourne en local,
tu vois le code avant/après exécution, et tu peux brancher le modèle de ton choix
(OpenAI, Anthropic, modèles locaux via Ollama/LM Studio selon config — à vérifier
pour la version que tu installes).

Pour un sysadmin, c'est l'outil « langage naturel → machine locale » par excellence :
exploration de logs, conversions de formats, petits scripts jetables, analyse de CSV
d'inventaire, génération de rapports.

## 59. Installation (version Python historique — à vérifier)

La voie historique, documentée de longue date :

```bash
# environnement dédié recommandé (jamais le Python système)
python3 -m venv ~/.venvs/oi && source ~/.venvs/oi/bin/activate
pip install open-interpreter
# ou : pipx install open-interpreter   (isolation applicative, recommandé)

interpreter
```

Premier lancement : l'outil demande la clé API / le provider et le modèle.
Configuration persistée dans un fichier de config (voir `--config` / docs du dépôt).

Version avec le toolkit de sécurité (voir section 61) :

```bash
pip install "open-interpreter[safe]"
```

⚠️ **À vérifier (sept 2026)** : avec la réécriture Rust en cours, la méthode
d'installation canonique a pu changer (binaire précompilé ? `cargo install` ?
gestionnaire de paquets ?). Consulte le README du dépôt au jour de l'installation
et note la version installée. Ne suis pas un tuto de 2024 sans vérifier.

## 60. Usage : les commandes et réglages qui comptent

En session interactive (version Python) :

| Commande / réglage | Effet |
|---|---|
| `interpreter` | lance la session |
| `interpreter --safe` | active le safe mode (section 61) |
| `%verbose true` (en session) | affiche le code généré en détail |
| `%auto_run true/false` | exécute le code sans demander (dangereux : voir §62) |
| `%model gpt-4o` (exemple) | change de modèle (noms à vérifier selon provider) |
| Fichier de config YAML | `model`, `temperature`, `safe_mode`, `system_message`… |

Réglages recommandés pour un usage sysadmin prudent :

```yaml
# config.yaml (exemple — clés à vérifier selon version)
model: <ton-modele>        # à vérifier : nom exact chez ton provider
temperature: 0             # déterminisme : indispensable en admin
verbose: true              # voir le code avant exécution
safe_mode: ask             # scan du code avant exécution (expérimental)
auto_run: false            # JAMAIS true sur une machine qui compte
```

## 61. Safe mode : ce qui existe (et ses limites)

Le projet documente un « safe mode » expérimental (docs du dépôt, relevé sept 2026) :

- installation : `pip install open-interpreter[safe]` (ajoute notamment `semgrep`) ;
- activation : flag `--safe` ou `safe_mode: ask|auto` dans la config ;
- effets : **désactive l'exécution automatique** du code + **scan du code généré**
  avec Semgrep avant exécution ;
- modes : `off` (défaut), `ask` (propose le scan), `auto` (scan systématique).

Limites, écrites noir sur blanc dans la doc du projet : **« Safe mode is experimental
and does not provide any guarantees of safety or security. »**

Traduction : c'est un filet, pas un mur. Un scan Semgrep ne détecte pas une logique
métier dangereuse (`rm -rf /tmp/*` un jour où /tmp est un lien symbolique…),
et « pas d'auto-run » ne protège que si tu lis vraiment le code avant de valider.
En pratique : safe mode = utile, mais ne remplace ni le sandboxing (section 52),
ni la relecture humaine, ni les sauvegardes.

La feuille de route du projet mentionnait aussi l'exécution en conteneurs
(PR #459 citée dans la doc) : à vérifier si c'est arrivé dans la version Rust.

## 62. Cas d'usage sysadmin (là où ça brille)

Scénarios où Open Interpreter est légitime et efficace :

1. **Exploration de logs** : « trouve les 10 IPs qui font le plus de 404 sur
   l'access.log d'hier, et dis-moi si c'est un scan » → il écrit le parsing,
   l'exécute, interprète.
2. **Inventaire et rapports** : CSV d'inventaire → tableaux, tris, exports,
   détection d'anomalies (machines sans agent de supervision, etc.).
3. **Conversions et migrations de formats** : vieux fichier de conf → nouveau
   format, avec vérification.
4. **Prototypage de scripts** : « écris-moi un script qui vérifie les certificats
   expirant dans < 30 jours sur ces 20 hôtes » → il génère, tu relis, tu valides,
   tu gardes le script.
5. **Analyse de données ponctuelle** : corréler deux exports (supervision +
   tickets) sans monter un notebook.

Point commun : **données locales, actions réversibles, humain qui relit.**
C'est le niveau 2 (supervisé) par excellence.

## 63. Limites d'Open Interpreter (à connaître avant de l'adopter)

- **Exécution de code = risque majeur** (section 64) : c'est le point central,
  pas un détail.
- **Contexte local uniquement** : pas de mémoire long terme native, pas de
  multi-agents, pas d'orchestration — c'est un agent solo interactif.
- **Dépendance au modèle** : la qualité du code généré dépend entièrement du
  modèle branché ; un petit modèle local fera plus d'erreurs (à compenser par
  plus de relecture).
- **Période de transition** (2026) : réécriture Rust en cours → doc et tutos
  potentiellement contradictoires selon la version. Épingle ta version et
  teste avant de t'appuyer dessus en prod.
- **Pas d'audit trail natif poussé** : en usage pro, ajoute ta propre journalisation
  (qui a demandé quoi, quel code a été exécuté — voir section 55).
- **Coût** : chaque session = N appels modèle ; sur de gros traitements,
  un script écrit une fois coûte 100× moins cher que l'agent qui le réinvente
  à chaque fois. Règle : l'agent **produit** le script, le script **tourne** en cron.

## 64. Sécurité Open Interpreter : la doctrine

Règle n°1 : **Open Interpreter exécute du code sur TA machine avec TES droits.**
Tout le reste en découle.

Doctrine minimale :

