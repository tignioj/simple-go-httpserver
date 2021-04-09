window.onload =function () {
    let ad = document.createElement("a");
    let adText = document.createTextNode("我是广告");
    ad.setAttribute("href", "http://localhost:80/adlink");
    ad.appendChild(adText);
    document.body.appendChild(ad);
}