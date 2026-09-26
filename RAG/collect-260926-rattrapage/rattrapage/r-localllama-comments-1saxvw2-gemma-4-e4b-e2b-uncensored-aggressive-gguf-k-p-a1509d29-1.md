---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1saxvw2-gemma-4-e4b-e2b-uncensored-aggressive-gguf-k-p-a1509d29-1
title: "r-localllama-comments-1saxvw2-gemma-4-e4b-e2b-uncensored-aggressive-gguf-k-p-a1509d29"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Google", "Hugging Face", "Nvidia"]
dates: []
keywords: ["gguf", "llama", "attention", "jailbreak", "llama.cpp", "moe", "multimodal", "nvidia", "sol"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1saxvw2-gemma-4-e4b-e2b-uncensored-aggressive-gguf-k-p-a1509d29.md
source_anchor: ""
source_lines: [1, 52]
sha256: 477e301b5b3c83a681ba666b76ba7d5dcf1efb9ef7afe7387ab9062bb35cac0a
---

# r-localllama-comments-1saxvw2-gemma-4-e4b-e2b-uncensored-aggressive-gguf-k-p-a1509d29

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
Gemma 4 E4B + E2B Non censuré (Agressif) — GGUF + K_P Quants (Multimodal : Vision, Vidéo, Audio)
Mes premières Gemma 4 uncensors sont disponibles. Deux modèles sortent aujourd'hui, le E4B (4B) et le E2B (2B). Les deux sont des variantes Aggressive, entièrement multimodales.
Aggressive signifie pas de refus. Je ne fais aucune modification de personnalité ou de changements. La version ORIGINALE de Google, juste uncensurée.
Gemma 4 E4B (4B) : https://huggingface.co/HauhauCS/Gemma-4-E4B-Uncensored-HauhauCS-Aggressive
Gemma 4 E2B (2B) : https://huggingface.co/HauhauCS/Gemma-4-E2B-Uncensored-HauhauCS-Aggressive
0/465 refus* sur les deux. Entièrement débloqué sans perte de capacité.
Ces modèles sont nativement multimodaux, donc texte, image, vidéo et audio tout en un. Le fichier mmproj est inclus pour le support vision/audio.
Ce qui est inclus :
E4B : Q8_K_P, Q6_K_P, Q5_K_P, Q5_K_M, Q4_K_P, Q4_K_M, IQ4_XS, Q3_K_P, Q3_K_M, IQ3_M, Q2_K_P + mmproj
E2B : Q8_K_P, Q6_K_P, Q5_K_P, Q4_K_P, Q3_K_P, IQ3_M, Q2_K_P + mmproj
Tous les quants générés avec imatrix. Les quants K\_P utilisent une analyse spécifique au modèle pour préserver la qualité là où cela compte le plus, ce qui donne effectivement 1 à 2 niveaux de quantification de mieux pour seulement ~5 à 15 % de taille de fichier plus grande. Entièrement compatible avec llama.cpp, LM Studio, ou tout ce qui lit GGUF (Ollama pourrait nécessiter des ajustements par l'utilisateur).
Spécifications rapides (les deux modèles) :
- 42 couches (E4B) / 35 couches (E2B)
- Fenêtre glissante mixte + attention complète
- Contexte natif de 131K
- Nativement multimodal (texte, image, vidéo, audio)
- Couches KV partagées pour l'efficacité de la mémoire
Échantillonnage depuis Google : temp=1.0, top_p=0.95, top_k=64. Utilisez le drapeau --jinja avec llama.cpp.
Remarque : Le widget de compatibilité matérielle de HuggingFace ne reconnaît pas les quants K_P, donc cliquez sur "Voir +X variantes" ou allez dans Fichiers et versions pour voir tous les téléchargements. K_P montrant "?" dans LM Studio est uniquement cosmétique, le modèle se charge bien.
À venir : Gemma 4 E31B (dense) et E26B-A4B (MoE). Je travaille dessus maintenant et je les publierai dès que je serai satisfait de la qualité. Les petits modèles étaient simples, les grands nécessitent plus d'attention.
*Google utilise désormais des techniques similaires à celles de GenRM de NVIDIA, des modèles de récompense générative qui agissent comme des critiques internes, rendant la censure complète et véritable un domaine de plus en plus difficile. Ces modèles n'ont pas eu autant de temps de test manuel à long contexte que mes autres sorties. Je m'attends à ce que 99,999 % des utilisateurs ne rencontrent pas de cas limites, mais l'astérisque est là par honnêteté. Aussi : le E2B est un modèle 2B. Ajustez vos attentes en conséquence, il est impressionnant pour sa taille, mais ne vous attendez pas à rivaliser avec quoi que ce soit au-dessus de 7B.
Tous mes modèles : HuggingFace-HauhauCS
En passant, je travaille actuellement sur un projet très cool, que je reprendrai dès que je publierai les 2 autres modèles Gemma. J'ai hâte de tous les partager une fois que je les aurai terminés.
Section des commentaires
Ce n'est pas du tout vrai, Gemma 4 est bien moins aligné que Gemma 3. Je n'ai rien eu à changer et ça a presque instantanément éliminé les refus. Je n'ai vu aucune indication qu'ils utilisent l'une des techniques proposées dans la littérature, comme l'entraînement multidirectionnel ou l'injection de bruit résiduel.
Les modèles Gemma 3 avaient également le problème infamous des activations massives qui causaient des artefacts lors de l'ablitération et inspiraient des hacks comme la Winsorisation. Ceux-ci semblent inutiles avec Gemma 4.
Dans l'ensemble, les modèles Gemma 4 semblent très faciles et agréables à abliter.
Google eux-mêmes disent que Gemma 4 est beaucoup plus "sûr" que Gemma 3
Eh bien, ils ont tort. J'ai explosé Gemma 4 en 90 minutes après sa sortie, et c'était le modèle le plus facile à utiliser depuis Llama 3. D'autres ont signalé que Gemma 4 devient complètement non censuré avec une simple invite du système (ce qui ne fonctionne pas pour Gemma 3), donc cette assertion est définitivement fausse.
Merci. GenRM et des choses similaires provoquent instantanément une sortie hallucinatoire du LLM qui déformera (ou mieux encore, tordra) ce que votre entrée était. À moins que vous ne fassiez réellement des tests manuels, ces éléments peuvent passer pour des 'non-refus'.
D'après ce que j'ai vu, pratiquement tous les grands laboratoires font cette affirmation chaque fois qu'ils sortent une nouvelle version. Dans la plupart des cas, les modèles sont en réalité les mêmes, voire moins censurés qu'avant. C'est littéralement juste une déclaration qu'ils insèrent pour se couvrir légalement.
À ce stade, la plupart des laboratoires ont également réalisé la simple réalité que pousser trop loin le réglage de la sécurité nuit pas mal aux performances du modèle. Et comme la plupart des laboratoires souhaitent battre leurs concurrents, il est donc naturel que beaucoup aient en fait commencé à se concentrer moins sur l'aspect sécurité, malgré ce qu'ils déclarent publiquement.
Ouais. Je teste sur les versions d'ollama (parce que c'était le plus simple pour moi). Je le place au milieu de trucs NSFW, des contenus avec avertissement, un peu de cybersécurité... Je n'ai pas encore eu de refus catégorique. C'est du jailbreak assez basique. Ce n'est pas très pimenté, mais ça va dans des directions que je m'attendais à ce qu'il refuse. Un peu surpris.
Vos modèles sont assez impressionnants, mais c'est dommage que vous soyez si fermé à ce sujet. Merci de considérer le partage de vos techniques, ou à tout le moins de considérer le partage de versions non quantifiées. Je suis sûr que les gens adoreraient également construire sur celles-ci.
Le responsable du classement UGI n'a pas pu évaluer ces modèles, car ils ne sont disponibles que sous forme de GGUF, ce qui explique qu'il n'y ait toujours pas de confirmation indépendante des affirmations de « perte de capacité nulle ».
Je pense certainement qu'il serait souhaitable que de telles affirmations fortes, jamais faites par un autre chercheur, soient validées avec des chiffres concrets.
Je suis désolé, mais est-ce que quelqu'un peut me dire ce que signifie K_P ? Je comprends k_m, k_s, k_l, k_xl, etc. où se situe k_p ? De plus, est-ce que lcpp a commencé à prendre en charge l'audio ?
bartowski a détecté un problème avec la conversion, es-tu sûr de ne pas être affecté ? Il semble aussi que tu n'as pas de Q4 ks quants, ce qui est beaucoup plus important pour moi que la version non censurée.
Q4_K_P = BPW 5.2 !!! 🧐🔥
J'ai du mal à le faire recevoir de l'audio. Quelqu'un peut m'aider ? J'utilise llama.cpp sur Windows.
Merci. Serait-il possible de quantifier le fichier mmproj en Q8 également ? Ou penses-tu que cela dégradera trop la qualité ?
Quelqu'un a un guide sur comment installer ça sur Google Edge ? 🫣
J'espère pouvoir rejoindre le soutien de mlx, merci pour tes efforts
merci
est-ce que quelqu'un d'autre reçoit "Erreur : 500 Erreur interne du serveur : impossible de charger le modèle : F:\Ollama\models\blobs\sha256-0796d58372742ef8ddc76dd64cd2fde217b7ef32e3dc58e10873253e569cad6b" dans ollama ?
Je suis très nouveau dans les LLms locaux et j'essaie de faire fonctionner ça avec Ollama, mais j'obtiens "Erreur : 500 Erreur interne du serveur : impossible de charger le modèle".
Est-ce que je fais quelque chose de travers ? (P.S. : Je peux faire fonctionner l'officielle Gemma 4 E2B sans problème)
waouh
Je suis impatient de l'essayer
Comment est-ce que je fais tourner ça sur iOS ?
