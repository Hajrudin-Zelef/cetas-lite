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
    // Coloration syntaxique : theme highlight.js assorti au mode clair/sombre.
    var darkMode = t === "sombre" || t === "hard_dark";
    var hlDark = document.getElementById("hljs-theme-dark");
    var hlLight = document.getElementById("hljs-theme-light");
    if (hlDark) hlDark.disabled = !darkMode;
    if (hlLight) hlLight.disabled = darkMode;
    window.cetasHljsTheme = function (tt) {
      var d = tt === "sombre" || tt === "hard_dark";
      var a = document.getElementById("hljs-theme-dark");
      var b = document.getElementById("hljs-theme-light");
      if (a) a.disabled = !d;
      if (b) b.disabled = d;
    };
  }
  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", apply);
  } else {
    apply();
  }
})();
