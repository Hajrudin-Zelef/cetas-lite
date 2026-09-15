(function () {
  var PALETTES = ["bleu", "violet", "vert", "vert_pur", "bleu_ocean", "jaune_or", "rouge"];
  function apply() {
    var t = "clair";
    try {
      t = localStorage.getItem("cetas-lite-theme") || "clair";
    } catch (e) {}
    if (t === "ocean") t = "clair"; // ancien nom du theme clair
    var cls = "";
    if (t === "sombre") cls = "dark";
    else if (t === "hard_dark") cls = "dark hard-dark";
    document.body.className = cls;
    document.documentElement.dataset.theme = t;
    var p = "bleu";
    try {
      p = localStorage.getItem("cetas-lite-palette") || "bleu";
    } catch (e) {}
    if (PALETTES.indexOf(p) === -1) p = "bleu";
    document.documentElement.dataset.palette = p;
  }
  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", apply);
  } else {
    apply();
  }
})();
