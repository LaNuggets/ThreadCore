// GET URL QUERY VALUES
const urlParams = new URLSearchParams(window.location.search);

// SEARCH BAR VALUE FROM URL QUERY
const searchString = urlParams.get('q');
if (searchString != null) {
    document.getElementById("search").value = searchString
}

// URL QUERY FROM SEARCH BAR INPUT
let search = document.getElementById("search");

search.addEventListener("keypress", function(event) {
  if (event.key === "Enter") {
    event.preventDefault();
    if ('URLSearchParams' in window) {
      var searchParams = new URLSearchParams(window.location.search);
      searchParams.set("q", search.value);
      window.location.search = searchParams.toString();
    }
  }
});

// GET ERROR MESSAGE FROM URL QUERY
const messageString = urlParams.get('message');
const messageType = urlParams.get('type');
urlParams.delete("message")
urlParams.delete("type")
if (messageString != null) {
  document.getElementById("message").style.display = 'grid';
  let message_title = document.getElementsByClassName("message_title")
  for (let i = 0; i < message_title.length; i++) {
    message_title[i].textContent = messageString
  }
  if (messageType == "error") {
    document.getElementById("error").style.display = 'flex';
  } else if (messageType == "success"){
    document.getElementById("success").style.display = 'flex';
  } else if (messageType == "warning"){
    document.getElementById("warning").style.display = 'flex';
  } else if (messageType == "info"){
    document.getElementById("info").style.display = 'flex';
  }
}

const message_close = document.getElementsByClassName("message_close")
for (let i = 0; i < message_close.length; i++) {
  message_close[i].onclick= function() {
    document.getElementById("message").style.animation = "moveOpen 1s reverse";
    setTimeout(function(){
      document.getElementById("message").style.display = 'none';
    }, 1000);
  }
}



// If any images are not found, display default image
// const images = document.querySelectorAll('img');

// images.forEach(img => {
//   img.addEventListener('error', function handleError() {
//     const defaultImage =
//       'https://bobbyhadz.com/images/blog/javascript-show-div-on-select-option/banner.webp';

//     img.src = defaultImage;
//     img.alt = 'default';
//   });
// });

if(document.getElementById("logout")){
  document.getElementById("logout").onclick= function() {
    document.getElementById("disconnect").submit();
  }
}

if(document.getElementById("connectionButton")){
  document.getElementById("connectionButton").onclick= function() {
    document.getElementById("connection").style.display = 'grid';
  }
}

document.addEventListener("click", (evt) => {
  const formEl = document.getElementById("connection");
  let targetEl = evt.target; // clicked element
  if(targetEl == formEl) {
    document.getElementById("connection").style.display = 'none';
  }
});

document.getElementById("createCommunityButton").onclick= function() {
  document.getElementById("createCommunityForm").style.display = 'grid';
}

document.addEventListener("click", (evt) => {
  const formEl = document.getElementById("createCommunityForm");
  let targetEl = evt.target; // clicked element
  if(targetEl == formEl) {
    document.getElementById("createCommunityForm").style.display = 'none';
  }
});

document.getElementById("createPostButton").onclick= function() {
  document.getElementById("createPostForm").style.display = 'grid';
}

document.addEventListener("click", (evt) => {
  const formEl = document.getElementById("createPostForm");
  let targetEl = evt.target; // clicked element
  if(targetEl == formEl) {
    document.getElementById("createPostForm").style.display = 'none';
  }
});

document.querySelector("article.mediaOptions").addEventListener("click", function(evt){
  if(evt.target.type === "radio"){
    if (evt.target.value == "link") {
      document.getElementById("mediaFile").style.display = 'none';
      document.getElementById("mediaLink").style.display = 'flex';
    } else if (evt.target.value == "file") {
      document.getElementById("mediaFile").style.display = 'flex';
      document.getElementById("mediaLink").style.display = 'none';
    }
  }
});

document.querySelector("article.profileOptions").addEventListener("click", function(evt){
  if(evt.target.type === "radio"){
    if (evt.target.value == "link") {
      document.getElementById("profileFile").style.display = 'none';
      document.getElementById("profileLink").style.display = 'flex';
    } else if (evt.target.value == "file") {
      document.getElementById("profileFile").style.display = 'flex';
      document.getElementById("profileLink").style.display = 'none';
    }
  }
});

document.querySelector("article.bannerOptions").addEventListener("click", function(evt){
  if(evt.target.type === "radio"){
    if (evt.target.value == "link") {
      document.getElementById("bannerFile").style.display = 'none';
      document.getElementById("bannerLink").style.display = 'flex';
    } else if (evt.target.value == "file") {
      document.getElementById("bannerFile").style.display = 'flex';
      document.getElementById("bannerLink").style.display = 'none';
    }
  }
});
