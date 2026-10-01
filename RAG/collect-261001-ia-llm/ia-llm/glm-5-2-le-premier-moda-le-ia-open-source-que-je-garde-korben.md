---
id: collect-261001-ia-llm/ia-llm/glm-5-2-le-premier-moda-le-ia-open-source-que-je-garde-korben
title: "GLM 5.2 - Le premier modÃ¨le IA open source que je garde"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "Moonshot", "Z.ai"]
dates: []
keywords: ["glm", "open source", "benchmark", "claude", "deepseek", "fable 5", "kimi", "leaderboard", "llama", "moe", "open weights", "qwen"]
source: docs/RAG/collect-261001-ia-llm/glm-5-2-le-premier-moda-le-ia-open-source-que-je-garde-korben.md
source_anchor: ""
source_lines: [1, 54]
sha256: bc8a3fc72f175af042732fda12caa888d6b65b64bcd625f98d844ca762bd9608
---

# GLM 5.2 - Le premier modÃ¨le IA open source que je garde

## Ce quâil faut retenir RÃ©sumÃ© gÃ©nÃ©rÃ© par IA

1. GLM 5.2, le modÃ¨le open weights de Z.ai (744 milliards de paramÃ¨tres en MoE, 1 million de tokens de contexte), fonctionne sans bugs pour les usages de codage de l'auteur, contrairement aux autres modÃ¨les open source testÃ©s.
2. GLM 5.2 s'intÃ¨gre directement dans Claude Code via l'API compatible Anthropic de Z.ai, permettant d'utiliser les mÃªmes skills et scripts que avec Claude.
3. Sur le leaderboard Arena.ai dÃ©diÃ© au code front-end, GLM 5.2 arrive en deuxiÃ¨me place, premier modÃ¨le open weights Ã ce niveau (tous les autres dans le top sont propriÃ©taires).

Les amis, il faut que je vous parle de GLM 5.2 . Je l'utilise en ce moment mÃªme Ã travers Z.ai, et c'est la premiÃ¨re fois qu'un modÃ¨le open weights me donne satisfaction sur ce que je lui demande de faire. Et dieu sait que j'en ai testÃ© de ces putains de modÃ¨les !

GLM 5.2, c'est le dernier-nÃ© de Z.ai, le lab chinois connu avant sous le nom de Zhipu AI. Il est sorti en ce mois-ci (en juin), et c'est un gros bÃ©bÃ© avec ses 744 milliards de paramÃ¨tres en Mixture-of-Experts (MoE), dont Ã peu prÃ¨s 40 milliards qui s'activent pour chaque token, ainsi qu'une fenÃªtre de contexte qui monte Ã 1 million de tokens via la dÃ©clinaison glm-5.2[1m]. Le tout publiÃ©, comme toujours, sous licence MIT, avec les poids tÃ©lÃ©chargeables sur HuggingFace.

Bref, j'y croyais pas trop, mais j'ai quand mÃªme pris le petit abonnement Z.ai et j'ai lancÃ© mes outils habituels et codÃ© quelques nouvelles features sur mes logiciels. Et Ã surprise, il s'en sort trÃ¨s trÃ¨s bien pour mes usages (je dis bien pour mes usages !). J'ai eu aucun bug, pas de discussion Ã l'infini qui tourne autour du pot, ni de fin de conversation qui part en caractÃ¨res chinois comme me faisait souvent Qwen.

AprÃ¨s, le truc chouette, c'est que je l'ai branchÃ© directement dans Claude Code. Si Ã§a vous intÃ©resse, je me suis fait un petit launcher spÃ©cifique. C'est cadeau :

```
#!/usr/bin/env bash
export ANTHROPIC_BASE_URL="https://api.z.ai/api/anthropic"
export ANTHROPIC_AUTH_TOKEN=VOTRE_CLE_API
export ANTHROPIC_DEFAULT_SONNET_MODEL="glm-5.2[1m]"
export ANTHROPIC_DEFAULT_OPUS_MODEL="glm-5.2[1m]"
export CLAUDE_CODE_AUTO_COMPACT_WINDOW="1000000"
claude "$@"
```
Vous le sauvegardez sous le nom de votre choix, par exemple "glm". Puis vous faites un :

```
chmod +x glm
```
Et ensuite vous le lancez comme ceci :

```
./glm
```
L'idÃ©e, c'est que comme l'API de Z.ai est compatible Anthropic, il suffit de pointer Claude Code vers leur endpoint, de glisser votre clÃ©, et il cause Ã GLM 5.2 comme il causerait Ã Claude. Mes skills, mes scripts, tout marche pareil, c'est le feu !

Je regrette juste une chose, c'est de ne pas pouvoir le faire tourner en local chez moi. Parce que le bestiau, il est TROP gros. MÃªme rabotÃ© et quantifiÃ© en 2-bit pour la maison , il vous bouffe dans les 240 Go de RAM. Chez moi, j'ai pas le matos, et vous probablement pas non plus. Donc pour le moment, l'API, c'est la seule porte d'entrÃ©e rÃ©aliste et abordable.

Que ce soit Qwen, Llama, Kimi, DeepSeek, peu importe ce que j'ai testÃ© en local, pour mes usages un peu chiadÃ©s, Ã chaque fois je suis super dÃ©Ã§u. Alors celui-lÃ , pour ce que je lui demande, il tient trÃ¨s bien la route.

Maintenant, je vais pas vous vendre Ã§a non plus comme un Claude Killer mais j'ai quand mÃªme trouvÃ© un benchmark qui confirme mon ressenti. Sur le leaderboard Arena.ai dÃ©diÃ© au code front-end, GLM 5.2 pointe Ã la deuxiÃ¨me place, juste derriÃ¨re Fable 5. Et comme tout ce qui le prÃ©cÃ¨de est propriÃ©taire, Ã§a en fait le premier modÃ¨le open weights Ã ce niveau du classement.

Donc c'est pas la meilleure IA du monde, hein, mais c'est la premiÃ¨re open source qui me donne un rÃ©sultat qui me convient. Et vous savez tous Ã quel point je suis chiant et exigeant avec ce genre d'outil. En tout cas, c'est la premiÃ¨re fois que je me dis que l'IA open source pourrait vraiment entrer dans mon flux du quotidien, et pas juste rester un joujou pour classer des trucs ou faire du slop sur des blogs de SEO. Maintenant, entre nous, j'attends surtout que Fable 5, ou son Ã©quivalent, revienne mettre le feu !!

Si Ã§a vous tente d'essayer, il y a donc le GLM Coding Plan de Z.ai, qui dÃ©marre Ã 18 dollars par mois et qui est surtout taillÃ© pour le code. Il se branche sur Claude Code, Cline et une vingtaine d'outils du mÃªme acabit. Petit conseil au passage, ce lien vers le Plan GLM est un lien affiliÃ© certes, mais il vous offre 10 % de rÃ©duc si vous l'utilisez, et Ã§a me file un petit truc aussi, donc tout le monde y gagne.

VoilÃ , si vous codez avec autre chose jusqu'ici, Ã§a vaut le coup d'y jeter un Åil par curiositÃ©.

Source : Z.ai

## Commentaires

starfix!dans Surfshark ne vous rend pas invMorganedans Discord devine votre Ã¢ge sansts3rv1dans Les Ray-Ban Display arrivent eponpondans Openpilot - La NHTSA passe lesfabiendans Claude Code vous fait choisir
