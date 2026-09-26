---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1srb7xh-choosing-a-mac-mini-for-local-llms-what-would-you-1cebd62a-3
title: "r-localllama-comments-1srb7xh-choosing-a-mac-mini-for-local-llms-what-would-you-1cebd62a"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["gpu", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1srb7xh-choosing-a-mac-mini-for-local-llms-what-would-you-1cebd62a.md
source_anchor: ""
source_lines: [60, 84]
sha256: a5d7a981ea3f16f940c8a2e1626b7f221df494e4845e1242be2b69f914108a95
---

# r-localllama-comments-1srb7xh-choosing-a-mac-mini-for-local-llms-what-would-you-1cebd62a

À mon avis, ça ne vaut pas la peine d'acheter un Mac de 32 Go, surtout si vous voulez coder dessus. Un PC avec une carte dédiée de 32 Go + 16 Go de RAM pourra faire fonctionner votre IDE local confortablement et avoir un bon contexte (je fais tourner qwen3.6 en quant 4-bit avec un contexte de 260K et il reste encore un peu de marge). Mais à 64 Go de RAM, les choses changent.
32 Go, ce n'est pas suffisant en 2026 si tu t'intéresses aux LLM
d'accord et c'est là où j'ai atterri
64 Go minimum, 96 à 128 Go pour un vrai travail -- besoin d'inférence par clause
Et si mon cas d'utilisation n'est pas la programmation ? Mais plutôt comme un assistant/employé dans mon job de PM ? Est-ce que le M1 Mini avec 16 Go de RAM est suffisant pour un modèle comme Qwen 3.5-9B ?
D'après ce que je comprends, oui, les modèles quantifiés peuvent fonctionner pour votre cas d'utilisation, mais plus de RAM est toujours bon et vous permet de jouer un peu jusqu'à ce que vous trouviez la bonne combinaison.
Si tu achètes quelque chose exprès pour llm, alors c'est soit 64 Go, soit 128 Go. Il n'y a pas d'autre moyen
La RAM est définitivement la priorité ici. Si tu peux obtenir le modèle 64 Go M1 Max, c'est le bon choix pour les modèles plus grands, même si l'efficacité du M4 est tentante. Pour les pipelines RAG et les assistants de codage, tu vas atteindre le plafond de mémoire bien avant la vitesse de la puce.
La bande passante de mémoire sur les puces Max fait une énorme différence pour les tokens par seconde. Comme tu utilises déjà OpenClaw pour éviter la taxe API, tu vas apprécier la vitesse. Concernant les rumeurs sur le M5, elles circulent toujours. Les M4 sont déjà de véritables bêtes. Prends la meilleure RAM que tu peux te permettre maintenant et mets-toi au travail.
Voici mes appareils
M3 max 96 Go M4 mini 24 Go M5 air 16 Go
Le seul qui peut vraiment faire tourner des modèles locaux de manière productive est le M3, la RAM est le facteur le plus important.
Un mini avec 64 Go de RAM est un point de départ, mais il est limité dans ce qui peut être exécuté efficacement.
Les vitesses à cœur unique se sont beaucoup améliorées M3 2724 M4 3432 M5 4167
Les vitesses des disques se sont également beaucoup améliorées.
Je considérerais de regrouper des Mac mini à l'avenir, c'est une façon d'augmenter progressivement les choses.
merci, donc en réalité le M3 Max 96 Go est le seul des trois à vraiment avoir du poids pour l'inférence alors
M1 Mac Mini 2020, 16 Go de RAM. Tu es limité aux modèles de milieu de gamme, mais honnêtement, juste en débutant, tu as juste besoin d'un modèle sur lequel tu peux compter.
Quand tu arriveras au point où tu auras atteint les limites de la machine, les M5 seront déjà une génération en arrière. Et en plus, ils sont bon marché.
C'est un exercice d'équilibre difficile entre la capacité de VRAM et la bande passante de la mémoire. Bien sûr, les GPU sont incroyablement rapides, mais sur le marché actuel, un RTX 5090 coûte environ 4 000 € et vous laisse seulement 32 Go de VRAM.
Si vous visez des modèles denses de 27B ou des MoEs de 30B, vous avez besoin de plus d'espace. Si vous ne pouvez pas faire entrer l'ensemble du modèle, les poids et le cache KV dans la VRAM, votre performance va chuter immédiatement. Bien sûr, vous pourriez brancher quatre 5090 et devenir pro... :D mais alors vous vous retrouvez avec une consommation d'énergie et une chaleur incroyables.
C'est pourquoi j'ai trouvé que le M4 Pro Mac Mini avec 64 Go de RAM était le compromis ultime. Bien que sa bande passante de 273 Go/s ne soit pas à la hauteur d'un GPU discret de haut niveau, c'est amplement suffisant pour une inférence fluide. Vous pouvez charger confortablement des modèles plus grands avec une fenêtre de contexte décente tout en consommant une puissance ridicule de 40W. Même à 15-20 t/s, vous pouvez le faire tourner 24h/24 et 7j/7 sans vous soucier de la facture d'électricité.
C'est la conclusion à laquelle je suis arrivé après avoir pesé les options. J'attends actuellement les prix du M5 Pro, mais le M4 Pro est déjà un monstre pour cela.
Concernant les modèles M1/M2 mentionnés dans le fil : gardez à l'esprit que les versions de base/Pro de ces puces ont une bande passante significativement plus faible. Même avec plus de RAM, vous pourriez voir une génération de tokens beaucoup plus lente par rapport à l'architecture du M4 Pro.
J'attendrais jusqu'à la WWDC
