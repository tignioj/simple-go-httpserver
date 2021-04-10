window.onload =function () {
    let btn = document.createElement("button");
    let adText = document.createTextNode("我是js添加的按钮");
    btn.appendChild(adText);
    btn.onclick = function () {alert("你点击了这个按钮！")}
    document.body.appendChild(btn);
}