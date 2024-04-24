let pancarte = document.getElementById("pancarte");
let emailInput = document.getElementById("emailInput");
let emailLabel = document.getElementById("emailLabel");
let confpInput = document.getElementById("confpinput");
let confpLabel = document.getElementById("confplabel");
let changeConnectionText = document.getElementById("changeConnectionText");
let connectionText = document.getElementById("connectionText");

var action = ""


const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);

const type = urlParams.get("type");

action = type == "l" ?"r" : "l";
pancarte.textContent = type == "l" ? "Login":"Register";
emailInput.style.display = type == "l" ? "none":"block";
emailLabel.style.display = type == "l" ? "none":"block";
confpInput.style.display = type == "l" ? "none":"block";
confpLabel.style.display = type == "l" ? "none":"block";
connectionText.textContent = type == "l" ? "Don't have account ?":"Already have an account ?";
changeConnectionText.textContent = type == "l" ? "Create it":"Use it";

function changeAction(){
    action = action == "l" ?"r" : "l";
    pancarte.textContent = action == "l" ? "Login":"Register";
    emailInput.style.display = action == "l" ? "none":"block";
    emailLabel.style.display = action == "l" ? "none":"block";
    confpInput.style.display = action == "l" ? "none":"block";
    confpLabel.style.display = action == "l" ? "none":"block";
    connectionText.textContent = action == "l" ? "Don't have account ?":"Already have an account ?";
    changeConnectionText.textContent = action == "l" ? "Create it":"Use it";

}