const   container=document.getElementById("container");
const   container2=document.getElementById("container2");

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
    nbr=15;
    nbr1=16;
    p=0;
    CheckSize();
    window.addEventListener("resize", CheckSize);

    for(i=1;i<=nbr;i++){
        div=document.createElement("div");
        div.className="image photo";
        div.style.backgroundImage="url('../src/img/skin/att"+i+".jpg')";
        div.style.width = slideWidth + "px";
        container.appendChild(div);
        container.style.width=(slideWidth*nbr)+"px";
    }
    for(i=1;i<=nbr1;i++){
        div=document.createElement("div");
        div.className="image photo";
        div.style.backgroundImage="url('../src/img/skin/def"+i+".jpg')";
        div.style.width = slideWidth + "px";
        container2.appendChild(div);
        container2.style.width=(slideWidth*nbr1)+"px";
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
d.onclick=function(){
    if(p>-nbr1+1)
        p--;
    container2.style.transform="translate("+p*slideWidth+"px)";
    container2.style.transition="all 0.5s ease"
}
c.onclick=function(){
    if(p<0)
        p++;
    container2.style.transform="translate("+p*slideWidth+"px)";
    container2.style.transition="all 0.5s ease"
}