// get query value in url from key
const urlParams = new URLSearchParams(window.location.search);

const searchString = urlParams.get('q');
if (searchString != null) {
    document.getElementById("search").value = searchString
}

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