let topic_btn = document.getElementById("topic_btn")
let post_btn = document.getElementById("post_btn")
let topic_div = document.getElementById("topic_div")
let post_div = document.getElementById("post_div")

post_btn.classList.toggle('select')
topic_div.classList.toggle('hidden')

function switchMain(){
    topic_btn.classList.toggle('select')
    post_btn.classList.toggle('select')
    topic_div.classList.toggle('hidden')
    post_div.classList.toggle('hidden')
}
