---
id: collect-261001-fortinet/fortinet/vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd-2
title: "vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license", "valuation"]
source: docs/RAG/collect-261001-fortinet/vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd.md
source_anchor: ""
source_lines: [71, 94]
sha256: 9a84d9bf79caa2660d986fc1866e949dcd40bbfdd4ae3cef21a51e405d7004a8
---

# vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd

Rappel 2 : les licences d'évaluation sont limitées à une seule machine virtuelle à la fois et comportent d'autres restrictions, comme l'absence de support Fortinet, l'utilisation de seulement trois interfaces réseau physiques, un seul CPU, 2 Go de RAM, etc.
Si vous vous reconnectez à la machine virtuelle, vous aurez la possibilité d'effectuer des configurations optionnelles. Celles-ci incluent la personnalisation du tableau de bord, le choix du mode de mise à jour de FortiOS, ou encore la migration via FortiConverter. Une fois cette étape franchie, l'administration de votre pare-feu peut débuter.
III. Conclusion
À mon avis, cette méthode est plus facile à adopter et à mettre en œuvre que celle basée sur GNS3, surtout si vous débutez avec les pare-feu FortiGate et travaillez sur des configurations simples.
Les choses se corsent quand on veut mettre en place des laboratoires qui demandent plusieurs pare-feu, comme pour configurer des tunnels VPN site-à-site, par exemple, ou autres.
Comme solution, vous pourriez soit créer plusieurs comptes FortiCloud différents les uns des autres, pour profiter d'autres licences d'évaluation FortiGate-VM, soit acheter des licences supplémentaires auprès d'un distributeur agréé si vous en avez les moyens…
Voilà qui marque la fin de ce tutoriel. N'hésitez pas à tester cette alternative chez vous et à nous faire un retour en commentaire.
Merci Mr Emile pour le partage…. Moi je suis confronté à un souci.
Je n’ai pas trouvé l’option « Evaluation licence »; Seul « Full licence » s’affiche .
Que faire???
Moi j’ai opté pour une version gratuite à vie avec la version Fortigate 7.0.13 et la commande execute factoryreset2, n’oubliez pas de rechanger ensuite votre mot de passe. A refaire à chaque demande de licence. La config est sauvegardée.
Merci pour cet article, très instructif.
@Euloge: Il faut sélectionner l’ovf Fortigate et pas Fortifirewall sur le site internet pour avoir l’option evaluation license.
Une fois encore merci.
C’est bien l’ovf FortiGate du VMware ESXi version 7.6.2 j’ai choisi sans obtenir l’option d’évaluation.
Je vous prie de m’envoyer si possible la version que vous avez par mail. D’avance merci
C’est bon à présent.
Merci à vous
Bonjour Florian,
Merci pour tous ces articles,
Je rencontre le même problème que dans le commentaire d’Euloge AHOUANDJINOU, après installation (7.6.6) il n’est proposé que la solution licence full…
Actuellement (début février 2026), les versions accessibles via votre lien sont :
7.6.6 / 7.4.11 / 7.2.3
Je n’ai pas encore essayé les autres versions proposées au téléchargement, mais ce sera probablement le même cas de figure…Dommage
