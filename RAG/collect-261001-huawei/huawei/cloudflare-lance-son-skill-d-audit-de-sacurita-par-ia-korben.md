---
id: collect-261001-huawei/huawei/cloudflare-lance-son-skill-d-audit-de-sacurita-par-ia-korben
title: "Cloudflare lance son skill d'audit de sÃ©curitÃ© par IA"
domain: huawei
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "exploit", "sandbox"]
source: docs/RAG/collect-261001-huawei/cloudflare-lance-son-skill-d-audit-de-sacurita-par-ia-korben.md
source_anchor: ""
source_lines: [1, 56]
sha256: 5f55f8c3447816deaff3bf10edd58e893cd37f74d96ca8c10c1d952c313b2966
---

# Cloudflare lance son skill d'audit de sÃ©curitÃ© par IA

## Ce quâil faut retenir RÃ©sumÃ© gÃ©nÃ©rÃ© par IA

1. Cloudflare publie sur GitHub un skill d'audit de sÃ©curitÃ© par IA, basÃ© sur des fichiers Markdown de consignes par famille d'attaque, plus un schÃ©ma JSON et deux validateurs JavaScript sans dÃ©pendance.
2. L'agent dÃ©roule six phases, de la cartographie de l'architecture jusqu'au rapport final, avec des chasseurs isolÃ©s par piste et un vÃ©rificateur indÃ©pendant qui n'a pas participÃ© Ã la chasse.
3. Chez Cloudflare, l'audit d'un dÃ©pÃ´t d'environ 30 000 lignes prend 3 Ã 4 heures via leur systÃ¨me interne, mais une seule passe ne trouve qu'environ la moitiÃ© des failles.

Cloudflare a mis en ligne sur son GitHub un skill trop super gÃ©nial d'audit de sÃ©curitÃ© qui leur a servi de point de dÃ©part Ã leur propre systÃ¨me de chasse aux failles de sÃ©cu. Ce package pour faire des audits de sÃ©cu boostÃ©s Ã l'agent IA c'est un dossier de consignes en Markdown que vous dÃ©posez dans votre agent de code, et qui le fait bosser comme un vrai auditeur et pas simplement comme un "relecteur" de code.

Dans ce dÃ©pÃ´t, vous trouverez un fichier par famille d'attaque : la corruption mÃ©moire, l'injection de prompt, le cadrage des requÃªtes HTTP, l'isolation entre locataires. Et avec Ã§a, un schÃ©ma JSON qui dÃ©crit Ã quoi doit ressembler un rapport, et deux validateurs Ã©crits en JavaScript sans la moindre dÃ©pendance.

*Le dossier du skill sur GitHub : un fichier de consignes par famille d'attaque, et les deux validateurs en bas de liste (
Source
)*

L'agent dÃ©roule ensuite six phases, de la reconnaissance au rapport final. Il commence par cartographier l'architecture, les frontiÃ¨res de confiance et les points d'entrÃ©e, puis il Ã©crit sa propre grille de couverture. Ensuite il envoie des chasseurs isolÃ©s case par case et chaque piste qui remonte part chez un vÃ©rificateur tout frais tout beau qui n'a pas participÃ© Ã la chasse.

Chaque piste se termine ensuite avec un rapport technique. Un rapport `confirmed` qui exige une trace source complÃ¨te et un rÃ©sultat rÃ©ellement observÃ©. Un rapport `needs_validation` qui nomme le fait prÃ©cis qui manque, et surtout qui n'a droit Ã  aucun niveau de gravitÃ©. Et le rapport des `rejected` qui garde la mÃ©moire de ce qui a Ã©tÃ© rÃ©futÃ©, pour que la passe suivante ne vous ressorte pas la mÃªme chose.

Le validateur, lui, c'est du vrai code qui vÃ©rifie la "forme" de la preuve. Les empreintes doivent Ãªtre uniques, et le suivi du parcours des tests qui doit partir d'un point d'entrÃ©e pour finir sur un point d'impact. Ainsi une trouvaille confirmÃ©e dont le parcours ne tient pas debout est recalÃ©e.

En revanche, sachez-le, l'indÃ©pendance du vÃ©rificateur est demandÃ©e dans les consignes, mais pas contrÃ´lÃ©e (chez Cloudflare, c'est leur orchestrateur maison qui s'en charge, et il n'est pas dispo publiquement).

Cloudflare explique sur son site ce que donne un de ces agents lÃ¢chÃ© sur du code sans garde-fou. Il modifie la source pour que son exploit fonctionne, puis il annonce fiÃ¨rement le bug qu'il vient de crÃ©er lui-mÃªme. Ou alors il pond un test qui dÃ©montre que exec() exÃ©cute des choses, donc que c'est forcÃ©ment une faille critique... Et Ã§a vous l'aurez compris, c'est de la merde et c'est absolument ce qu'on ne veut pas.

VoilÃ donc exactement ce que ces consignes cherchent Ã lui interdire.

Pour l'installer, il vous faut le CLI skills de Vercel Labs :

```
npx skills add https://github.com/cloudflare/security-audit-skill --skill security-audit
```
Ensuite, vous lancez votre agent dans le dÃ©pÃ´t Ã auditer et vous lui demandez un audit de sÃ©curitÃ©. Le skill se dÃ©clenche tout seul et c'est le modÃ¨le de votre agent qui pilotera les sous-agents en parallÃ¨le (utilisez donc un modÃ¨le qui sache faire Ã§a, c'est mieux lool). Par dÃ©faut, le rapport atterrit ensuite dans un dossier ~/security-audit-skill, en dehors du dÃ©pÃ´t auditÃ©.

Ce qui est cool c'est que ce skill est conÃ§u pour limiter les bÃªtises. Par exemple il refuse d'exÃ©cuter le code qu'il audite s'il n'a pas de bac Ã  sable (sandbox) imposÃ© par le systÃ¨me d'exploitation, avec le rÃ©seau coupÃ©, la cible en lecture seule et des limites de CPU et de mÃ©moire. Ce bac Ã  sable, le dÃ©pÃ´t ne le fournit pas, et c'est Ã  vous de le mettre en place avec par exemple
un outil d'isolation comme Fence
. Sans lui, la piste de l'audit finira tout bonnement en `needs_validation` au lieu d'Ãªtre vÃ©ritablement suivie et tranchÃ©e.

AprÃ¨s, c'est le temps qui fera la diffÃ©rence. Chez Cloudflare, ils ont un gros systÃ¨me qui leur permet, par exemple avec un dÃ©pÃ´t d'environ 30 000 lignes de code, de ne passer que 3 Ã 4 heures sur l'audit pour avoir un rÃ©sultat correct. Mais chez vous, en local, ce sera beaucoup plus long. Cloudflare prÃ©vient aussi qu'une seule passe n'est pas suffisante puisqu'elle trouve Ã peu prÃ¨s la moitiÃ© des failles Ã chaque fois. Il faudra relancer Ã§a plusieurs fois.

Allez, si vous cherchez Ã©galement un filet de sÃ©curitÃ© pour tout ce qui concerne vos pull requests, sachez qu'Anthropic maintient de son cÃ´tÃ© un reviewer de sÃ©curitÃ© pour Claude Code qui lit le diff et vient commenter directement dessus.

Le dÃ©pÃ´t de Security Audit Skill de Cloudflare est sous licence MIT, et il lui faudra Node.js pour faire tourner ses deux validateurs.

En tout cas, ce que je vous conseille, c'est d'aller lire le SKILL.md avant de lancer la premiÃ¨re passe, pour vous assurer que tout sera OK au sein de votre harness.

Source : security-audit-skill sur GitHub

EntiÃ¨rement dÃ©diÃ©e Ã la cybersÃ©curitÃ©, l'Ã©cole Guardia est accessible soit directement aprÃ¨s le bac (post-bac), soit aprÃ¨s un bac+2 ou bac+3. En rejoignant l'Ã©cole Guardia, vous deviendrez dÃ©veloppeur informatique option cybersÃ©curitÃ© (Bac+3) ou expert en cybersÃ©curitÃ© (Bac+5).

Guardia CS forme aussi les professionnels Ã la cybersÃ©curitÃ© via plusieurs formations en ligne

## Commentaires

starfix!dans Surfshark ne vous rend pas invMorganedans Discord devine votre Ã¢ge sansts3rv1dans Les Ray-Ban Display arrivent eponpondans Openpilot - La NHTSA passe lesfabiendans Claude Code vous fait choisir
