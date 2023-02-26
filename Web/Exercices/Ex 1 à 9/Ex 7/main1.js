const   container=document.getElementById("container");

let slideWidth = 800;
function CheckSize(){
    if (window.innerWidth <= 850) {
        slideWidth = 400;
    } else if (window.innerWidth <= 1200) {
        slideWidth = 600;
    } else {
        slideWidth = 800;
    }
}

document.body.onload=function(){
    nbr=17;
    p=0;
    CheckSize();
    window.addEventListener("resize", CheckSize);

    for(i=1;i<=nbr;i++){
        div=document.createElement("div");
        div.className="image photo";
        div.style.backgroundImage="url('../img/map/Map"+i+".webp')";
        div.style.width = slideWidth + "px";
        container.appendChild(div);
        container.style.width=(slideWidth*nbr)+"px";
    }
}

b.onclick=function(){
    if(p>-nbr+1)
        p--;
    container.style.transform="translate("+p*slideWidth+"px)";
    container.style.transition="all 0.5s ease"
}
a.onclick=function(){
    if(p<0)
        p++;
    container.style.transform="translate("+p*slideWidth+"px)";
    container.style.transition="all 0.5s ease"
}