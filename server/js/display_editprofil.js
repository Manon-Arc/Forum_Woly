
function showEditProfil(modal){
    modal.classList.toggle('hidden');
}

function imgChange(self){
    if(self.files.length > 0){
        var src = URL.createObjectURL(self.files[0]);
        var preview = document.getElementById("renderImg");
        preview.src = src;
      }
}

document.addEventListener("click", (evt) => {
    const creationDiv = document.getElementById("editDiv");
      if(evt.target == creationDiv) {
        showEditProfil(creationDiv);
        return;
      }
      evt.target = evt.target.parentNode;
  });