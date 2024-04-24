
const template = `<div class="topicCreation" id="creationDiv">
    <form action="/creat_topic" method="post" class="card">
        <div class="pancarte">
            <h1>Create a Topic</h1>
        </div>
        <img src="/server/img/woly.png" alt="" class="woly">
        <label for="topic_name">Topic's Name</label>
        <input type="text" name="topic_name" placeholder="Topic's name" placeholder="name">
        <input type="text" name="topic_category" placeholder="category">
        <label for="topic_content">Topic's Description</label>
        <textarea type="text" name="topic_description" placeholder="description" placeholder="Topic's description" class="descText"></textarea>
        <input type="submit" class="button" value="Create It">
    </form>
</div>`;

function showCreation(){
    document.body.innerHTML += template;
}

document.addEventListener("click", (evt) => {
    const creationDiv = document.getElementById("creationDiv");
      if(evt.target == creationDiv) {
        creationDiv.remove();
        return;
      }
      evt.target = evt.target.parentNode;
  });