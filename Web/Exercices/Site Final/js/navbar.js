let alreadyReloaded = false;

function verifierMedia() {
  if (window.matchMedia("(max-width: 900px)").matches) {
    const wrap = document.querySelector('.wrapper nav');
    const lien1 = document.getElementById('link-1');
    const lien2 = document.getElementById('link-2');
    const lien3 = document.getElementById('link-3');
    const lien4 = document.getElementById('link-4');
    const lien5 = document.getElementById('link-5');
    const btn = document.querySelector('.btn-img');
    if (!alreadyReloaded) {
      alreadyReloaded = true;
      btn.addEventListener("click", function(){
        if (wrap.style.height == "70px"){
          wrap.style.height = "80vh";
          lien1.style.display = "block";
          lien2.style.display = "block";
          lien3.style.display = "block";
          lien4.style.display = "block";
          lien5.style.display = "block";
        } else {
          wrap.style.height = "70px";
          lien1.style.display = "none";
          lien2.style.display = "none";
          lien3.style.display = "none";
          lien4.style.display = "none";
          lien5.style.display = "none";
        }
      })
    }
  } else {
    if (window.matchMedia("(min-width: 900px)").matches) {
       refresh()
    }
  }
}

window.addEventListener('resize', verifierMedia);

verifierMedia();
