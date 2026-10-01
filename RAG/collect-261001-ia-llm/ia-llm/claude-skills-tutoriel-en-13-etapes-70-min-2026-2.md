---
id: collect-261001-ia-llm/ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026-2
title: "claude-skills-tutoriel-en-13-etapes-70-min-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "agent", "license"]
source: docs/RAG/collect-261001-ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026.md
source_anchor: ""
source_lines: [43, 120]
sha256: 561cfca3bbf57cd3c0341da6c7e44ac77c5860843b0565389503e54d551e1266
---

# claude-skills-tutoriel-en-13-etapes-70-min-2026

| Emplacement | Chemin | Portée | 
|---|---|---|
| Entreprise | Paramètres managés (settings.json d’organisation) | Tous les utilisateurs de l’organisation | 
| Personnel | `~/.claude/skills/<nom-skill>/SKILL.md` | Tous vos projets, sur votre machine | 
| Projet | `.claude/skills/<nom-skill>/SKILL.md` | Ce projet uniquement, partageable via Git | 
| Plugin | `<plugin>/skills/<nom-skill>/SKILL.md` | Partout où le plugin est activé | 

Pour ce tutoriel, nous allons commencer par une Skill personnelle : elle est immédiate à tester, disponible dans tous vos projets, et ne nécessite aucune coordination d’équipe. Une fois validée, rien n’empêche de la déplacer vers `.claude/skills/` à la racine d’un dépôt pour la partager avec le reste de l’équipe via un commit Git classique — c’est d’ailleurs la méthode recommandée pour diffuser une convention d’équipe sans passer par un plugin.

Un détail qui surprend souvent les débutants : les Skills se chargent aussi depuis des dossiers `.claude/skills/` imbriqués, en dessous du dossier où vous lancez Claude Code, jusqu’à la racine du dépôt. Dans un monorepo, un sous-dossier comme `apps/web/.claude/skills/` peut donc exposer des Skills propres à cette application, qui ne deviennent actives que lorsque Claude touche un fichier dans ce sous-dossier.

## Étape 2 : créer l’arborescence de dossiers

Le projet fil rouge de ce tutoriel est une Skill nommée `pr-description`, qui génère une description de pull request structurée à partir du diff Git en cours et d’un numéro de ticket. C’est un besoin réel, présent dans quasiment toutes les équipes, et il illustre bien l’injection de contexte dynamique et les arguments que nous verrons plus loin. Créez d’abord le dossier personnel :

```
mkdir -p ~/.claude/skills/pr-description
cd ~/.claude/skills/pr-description
```
Le nom du dossier n’est pas cosmétique : c’est lui qui détermine la commande que vous taperez pour invoquer la Skill manuellement. Un dossier `pr-description` donne la commande `/pr-description`. Le champ `name` du frontmatter, que nous verrons à l’étape suivante, ne sert qu’à l’affichage dans les listes de Skills pour une installation personnelle ou projet ; seule une Skill de plugin utilise ce champ pour construire le nom de commande final.

## Étape 3 : rédiger le frontmatter YAML de SKILL.md

Chaque `SKILL.md` commence par un bloc YAML entre deux lignes `---`. Deux champs sont couverts par la spécification officielle et recommandés dans tous les cas : `name` et `description`. Anthropic impose des contraintes strictes sur le champ `name` : 64 caractères maximum, uniquement des lettres minuscules, des chiffres et des tirets, aucune balise XML, et surtout aucun des mots réservés “anthropic” ou “claude”. Le champ `description` doit être non vide, ne pas dépasser 1024 caractères, et ne contenir aucune balise XML.

La documentation officielle de la plateforme Claude insiste sur ce point : la qualité de ce champ `description` détermine à elle seule si votre Skill sert à quelque chose. C’est le texte que Claude compare à chaque nouvelle demande pour décider s’il doit charger la Skill : il doit donc dire à la fois ce que fait la Skill et quand l’utiliser, pas seulement l’un des deux.

```
---
name: pr-description
description: Génère une description de pull request structurée (résumé, changements, tests, risques) à partir du diff Git en cours. Utiliser quand l'utilisateur demande une description de PR, un message de pull request, ou veut résumer ses changements avant de les pousser.
---
```
Claude Code accepte, en plus de ces deux champs, une liste étendue d’attributs optionnels qui ne font pas partie du standard portable mais permettent un contrôle beaucoup plus fin du comportement de la Skill. Le tableau ci-dessous résume les plus utiles ; nous en utiliserons plusieurs dans les étapes suivantes.

| Champ | Rôle | Valeur par défaut | 
|---|---|---|
| `when_to_use` | Contexte de déclenchement additionnel, ajouté à la description | Absent | 
| `argument-hint` | Indice affiché dans l’autocomplétion, ex. `[numero-ticket]` | Absent | 
| `arguments` | Noms d’arguments positionnels pour la substitution `$nom` | Absent | 
| `disable-model-invocation` | Empêche Claude de déclencher la Skill tout seul | `false` | 
| `user-invocable` | Masque la Skill du menu `/` si`false` | `true` | 
| `allowed-tools` | Outils pré-autorisés sans confirmation pendant l’invocation | Absent | 
| `context` | `fork` exécute la Skill dans un sous-agent isolé | Absent (contexte principal) | 
| `model` | Modèle à utiliser le temps de la Skill | Hérite de la session | 

Important à savoir si vous comptez un jour distribuer la même Skill sur claude.ai ou via l’API : ces deux surfaces, ainsi que le script officiel `package_skill.py` du dépôt anthropics/skills, n’acceptent que six champs de frontmatter — `name`, `description`, `license`, `compatibility`, `metadata` et `allowed-tools`. Tout champ hors de cette liste (comme `argument-hint` dans l’exemple ci-dessus) fait échouer l’empaquetage avec une erreur explicite plutôt que d’être simplement ignoré.

## Étape 4 : rédiger le corps d’instructions

Sous le frontmatter, le corps du `SKILL.md` est le contenu que Claude lit réellement une fois la Skill déclenchée (le “niveau 2” de la divulgation progressive, plafonné à environ 5 000 tokens selon la documentation officielle). C’est du Markdown standard : titres, listes, blocs de code. La règle d’or reste la concision, puisque ce contenu reste dans le contexte pour le reste de la conversation une fois chargé — chaque ligne a un coût récurrent en tokens.

```
## Instructions
À partir du diff fourni ci-dessous, rédige une description de pull
request avec ces sections :
1. **Résumé** — une phrase expliquant le changement
2. **Changements** — liste à puces des modifications significatives
3. **Tests** — comment le changement a été vérifié
4. **Risques** — effets de bord potentiels, à zéro si aucun
Si un numéro de ticket est fourni en argument, ajoute une ligne
"Closes #NUMERO" à la fin.
## Diff en cours
!`git diff HEAD`
```
La dernière ligne de cet exemple, `` !`git diff HEAD` ``, est ce qu’Anthropic appelle l’injection de contexte dynamique : nous la détaillons à l’étape 7. Gardez pour l’instant en tête la structure générale — instructions claires, puis données injectées — car c’est le patron que vous retrouverez dans la quasi-totalité des Skills utiles en développement logiciel.

## Étape 5 : tester votre premier Skill avec Claude Code

Enregistrez le fichier, placez-vous dans un dépôt Git contenant des modifications non commitées, puis lancez Claude Code avec la commande `claude`. Deux façons de vérifier que la Skill fonctionne : la laisser se déclencher toute seule, ou l’appeler explicitement.

- **Déclenchement automatique** : tapez une demande qui correspond à la description, par exemple “Rédige-moi une description de PR pour ces changements”.
- **Invocation directe** : tapez`/pr-description` dans le prompt.

Dans les deux cas, Claude doit répondre avec un résumé structuré de vos modifications réelles, pas un exemple générique. Si rien ne se déclenche automatiquement, le problème vient presque toujours de la `description` : reformulez-la pour qu’elle colle exactement aux mots qu’un utilisateur emploierait naturellement. Claude Code surveille en continu les dossiers de Skills : toute modification apportée à un `SKILL.md` existant est prise en compte immédiatement dans la session en cours, sans redémarrage. Ce n’est que la création d’un tout nouveau dossier de Skills de premier niveau qui exige de relancer Claude Code pour qu’il commence à le surveiller.

## Étape 6 : ajouter des fichiers de support (scripts et références)

