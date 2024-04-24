let modEle = document.getElementsByClassName("moderation");
let visEle = document.getElementsByClassName("visit")


function getAdminMod(isAdmin){

    console.log(isAdmin)

    if(isAdmin){
        for (var i=0;i<modEle.length;i+=1){
            modEle[i].style.display = 'flex';
        }
        for (var i=0;i<visEle.length;i+=1){
            visEle[i].style.display = 'none';
        }
        console.log("is admin")
    } else {
        for (var i=0;i<modEle.length;i+=1){
            modEle[i].style.display = 'none';
        }
        for (var i=0;i<visEle.length;i+=1){
            visEle[i].style.display = 'flex';
        }
        console.log("is not")
    }

}