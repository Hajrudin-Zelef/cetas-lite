---
id: collect-261001-ia-llm/ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026-5
title: "claude-skills-tutoriel-en-13-etapes-70-min-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "mai", "open source"]
source: docs/RAG/collect-261001-ia-llm/claude-skills-tutoriel-en-13-etapes-70-min-2026.md
source_anchor: ""
source_lines: [295, 330]
sha256: bc887151f330ac72c6a85879347a61fcf97e1056ab7ff65906dc043a64883876
---

# claude-skills-tutoriel-en-13-etapes-70-min-2026

- **Une description trop vague.** “Aide avec les PR” ne dit ni ce que fait la Skill ni quand l’invoquer. Claude ne la déclenchera jamais de façon fiable. Formulez toujours la description comme une réponse aux deux questions : que fait cette Skill, et à quel moment dois-je l’utiliser.
- **Un champ `name` qui viole les contraintes.** Majuscules, espaces, ou les mots réservés “claude”/”anthropic” font échouer la validation lors de l’empaquetage vers claude.ai ou l’API. Restez sur des minuscules, des chiffres et des tirets.
- **Confondre les emplacements de stockage.** Une Skill dans`~/.claude/skills/` n’est visible que sur votre machine. Beaucoup de développeurs découvrent trop tard qu’un collègue qui clone le même dépôt ne voit rien, faute d’avoir committé la Skill dans`.claude/skills/` du projet.
- **Un SKILL.md surchargé.** Tout mettre dans le corps principal plutôt que dans des fichiers de référence annexes annule l’intérêt de la divulgation progressive : chaque chargement de la Skill consomme alors inutilement du contexte. La recommandation officielle est de rester sous 500 lignes dans le fichier principal.
- **Faire confiance à une Skill dont l’origine est inconnue.** Une Skill malveillante peut orienter Claude vers des appels d’outils ou des commandes qui ne correspondent pas à son objectif affiché. N’installez que des Skills que vous avez écrites vous-même ou qui proviennent directement d’Anthropic.
- **Oublier `disable-model-invocation` sur une action à effet de bord.** Sans cette protection, rien n’empêche Claude de déclencher tout seul une Skill de déploiement parce que le contexte de la conversation “ressemblait” à une demande de mise en production.

## Sécurité : ce qu’Anthropic recommande

La documentation officielle est directe sur ce point : une Skill donne à Claude de nouvelles capacités via des instructions et du code, ce qui signifie qu’une Skill malveillante peut le pousser à exécuter des actions qui ne correspondent pas à l’objectif affiché. Cette vigilance est d’autant plus nécessaire que les sources tierces se multiplient : la collection GitHub d’Obviousworks, mise à jour le 30 mai 2026, agrège déjà des dizaines de Skills provenant à la fois d’Anthropic et de contributeurs communautaires, sans garantie d’audit homogène. Anthropic recommande de traiter l’installation d’une Skill exactement comme l’installation d’un logiciel, avec le même niveau de méfiance envers les sources inconnues.

- **N’utilisez que des sources fiables** : vos propres Skills, ou celles publiées directement par Anthropic.
- **Auditez systématiquement** le contenu complet d’une Skill externe — le SKILL.md, les scripts, les images et toute autre ressource jointe — à la recherche d’appels réseau, d’accès fichiers ou d’opérations qui ne correspondent pas à l’objectif annoncé.
- **Méfiez-vous particulièrement des Skills qui récupèrent du contenu depuis des URL externes** : ce contenu peut inclure des instructions malveillantes, et une dépendance externe pourtant fiable au départ peut changer de comportement dans le temps.
- **Gardez à l’esprit que les Skills ne sont pas couvertes par les accords de rétention zéro des données (ZDR)** : les définitions de Skills et les données d’exécution suivent la politique de rétention standard d’Anthropic, pas une politique zéro-rétention même si votre organisation en bénéficie par ailleurs.
- **Traitez toute intégration en production avec prudence renforcée** , en particulier lorsque la Skill a accès à des données sensibles ou à des opérations critiques.

## Dépannage : 8 problèmes fréquents et solutions

- **La Skill n’apparaît pas dans le menu `/`.** Vérifiez que le fichier s’appelle exactement`SKILL.md` (sensible à la casse) et qu’il se trouve directement à la racine de son dossier, pas dans un sous-dossier. Vérifiez aussi que`user-invocable` n’est pas positionné à`false` .
- **Claude ne déclenche jamais la Skill automatiquement.** La cause numéro un est une`description` mal formulée. Réécrivez-la en incluant les mots exacts qu’un utilisateur emploierait pour formuler sa demande.
- **Erreur “Unexpected key(s) in SKILL.md frontmatter” lors de l’empaquetage.** Vous utilisez un champ Claude Code (comme`argument-hint` ou`context` ) qui ne fait pas partie des six champs portables. Retirez-le ou gardez deux versions du frontmatter selon la destination.
- **La Skill fonctionne dans Claude Code mais pas sur claude.ai.** Normal : les Skills ne se synchronisent jamais automatiquement entre les surfaces. Il faut la téléverser séparément sur chaque canal où vous voulez l’utiliser.
- **Le champ `allowed-tools` avec `${CLAUDE_SKILL_DIR}` continue à demander une confirmation.** Cette substitution dans`allowed-tools` nécessite une version relativement récente de Claude Code. Mettez à jour avec`claude --version` pour vérifier, puis mettez à jour l’outil si besoin.
- **Deux Skills du même nom entrent en conflit dans un monorepo.** C’est un comportement attendu : la variante imbriquée apparaît sous un nom qualifié par son chemin (par exemple`apps/web:deploy` ), tandis que le nom court reste réservé à la Skill de la racine du projet.
- **Un script s’exécute mais son résultat n’apparaît nulle part.** Vérifiez que le SKILL.md demande explicitement à Claude d’exécuter le script et d’utiliser sa sortie, plutôt que de simplement mentionner son existence.
- **Une nouvelle Skill ajoutée en cours de session reste invisible.** La détection à chaud couvre les modifications d’un SKILL.md existant, mais la création d’un tout nouveau dossier racine de Skills nécessite un redémarrage de Claude Code pour être détectée.

## Astuces avancées pour aller plus loin

Une fois les bases maîtrisées, plusieurs mécanismes permettent d’aller nettement plus loin dans la spécialisation de vos Skills. Le champ `paths` accepte des motifs glob qui limitent l’activation automatique d’une Skill aux fichiers correspondants : une Skill de conventions React ne se déclenchera ainsi que lorsque Claude travaille sur des fichiers `.tsx`, sans jamais polluer le contexte système prompt pour le reste du projet.

Le champ `hooks` permet d’attacher des scripts au cycle de vie propre d’une Skill (avant ou après son exécution), utile pour journaliser chaque déclenchement d’une Skill sensible ou valider un état avant de la laisser s’exécuter. Le champ `metadata`, lui, accepte n’importe quelle donnée clé-valeur libre : pratique pour votre propre tooling interne (numéro de version, équipe propriétaire, catalogue de Skills), sans que Claude Code n’agisse dessus directement.

Pour les équipes qui gèrent un catalogue de Skills conséquent, le dépôt open source anthropics/skills mérite un détour : il embarque notamment une Skill “Claude API” livrant une documentation de référence à jour pour huit langages de programmation, directement intégrée par défaut dans Claude Code. Le succès de ce dépôt officiel donne une idée de l’engouement suscité par le format : selon Taskade, il totalisait plus de 157 000 étoiles GitHub en juillet 2026, dépassé par la Skill communautaire Superpowers pour Claude Code qui affichait de son côté plus de 243 000 étoiles à la même date. C’est un bon modèle à étudier pour structurer vos propres Skills de référence technique, avec la même logique de fichiers annexes chargés à la demande plutôt qu’un unique fichier monolithique.

