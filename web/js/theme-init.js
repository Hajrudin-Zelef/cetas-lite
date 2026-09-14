(function () {
  function apply() {
    var t = "ocean";
    try {
      t = localStorage.getItem("cetas-lite-theme") || "ocean";
    } catch (e) {}
    var cls = "ocean-theme";
    if (t === "sombre") cls = "ocean-theme dark";
    else if (t === "clair") cls = "";
    document.body.className = cls;
    document.documentElement.dataset.theme = t;
  }
  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", apply);
  } else {
    apply();
  }
})();
