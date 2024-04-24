function editPost(id, content, topic){
    document.getElementById('edit_content').textContent = content
    document.getElementById('editDiv').classList.toggle('hidden')
    document.getElementById('editform').action = `/edit_post?id=${id}&idtopic=${topic}`
}

document.addEventListener("click", (evt) => {
    const creationDiv = document.getElementById("editDiv");
      if(evt.target == creationDiv) {
        editPost(0,"");
        return;
      }
      evt.target = evt.target.parentNode;
  });