let topics_cards = document.getElementsByClassName("topic");

for(var index = 0; index < topics_cards.length; index++){
    topics_cards.item(index).addEventListener("click", (event) => {
        window.location.href = "/topic?id=1";
    });
}