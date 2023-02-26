
function verifierMedia() {
    if (window.matchMedia("(max-width: 900px)").matches) {
        const wrap = document.querySelector('.wrapper nav');
        const lien1 = document.getElementById('link-1');
        const lien2 = document.getElementById('link-2');
        const lien3 = document.getElementById('link-3');
        const lien4 = document.getElementById('link-4');
        const btn = document.querySelector('.btn-img');
        btn.addEventListener("click", function(){
            if (lien1.style.opacity == "0"){
                wrap.style.height = "100vh";
                lien1.style.opacity = "1";
                lien2.style.opacity = "1";
                lien3.style.opacity = "1";
                lien4.style.opacity = "1";
              } else {
                wrap.style.height = "70px";
                lien1.style.opacity = "0";
                lien2.style.opacity = "0";
                lien3.style.opacity = "0";
                lien4.style.opacity = "0";
              }
        })
    }
}

verifierMedia()