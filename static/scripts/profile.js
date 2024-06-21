document.getElementById("createPostButton2").onclick= function() {
    document.getElementById("createPostForm").style.display = 'grid';
  }
  
  document.addEventListener("click", (evt) => {
    const formEl = document.getElementById("createPostForm");
    let targetEl = evt.target; // clicked element
    if(targetEl == formEl) {
      document.getElementById("createPostForm").style.display = 'none';
    }
  });

if(document.getElementById("updateUserButton")){
  document.getElementById("updateUserButton").onclick= function() {
      document.getElementById("updateUserForm").style.display = 'grid';
  }
}
  
document.addEventListener("click", (evt) => {
    const formEl = document.getElementById("updateUserForm");
    let targetEl = evt.target; // clicked element
    if(targetEl == formEl) {
      document.getElementById("updateUserForm").style.display = 'none';
    }
});

if(document.getElementById("deleteUserButton")){
  document.getElementById("deleteUserButton").onclick= function() {
    document.getElementById("deleteUserForm").style.display = 'grid';
  }
}
  
document.addEventListener("click", (evt) => {
    const formEl = document.getElementById("deleteUserForm");
    let targetEl = evt.target; // clicked element
    if(targetEl == formEl) {
      document.getElementById("deleteUserForm").style.display = 'none';
    }
});

document.querySelector("article.updateProfileOptions").addEventListener("click", function(evt){
    if(evt.target.type === "radio"){
      if (evt.target.value == "link") {
        document.getElementById("updateProfileFile").style.display = 'none';
        document.getElementById("updateProfileLink").style.display = 'flex';
        document.getElementById("updateProfileLink").placeholder = 'Profile picture link';
      } else if (evt.target.value == "file") {
        document.getElementById("updateProfileFile").style.display = 'flex';
        document.getElementById("updateProfileLink").style.display = 'none';
      } else if (evt.target.value == "keep" || evt.target.value == "remove") {
        document.getElementById("updateProfileFile").style.display = 'none';
        document.getElementById("updateProfileLink").style.display = 'flex';
        document.getElementById("updateProfileLink").placeholder = 'Keep this empty';
      }
    }
  });
  
  document.querySelector("article.updateBannerOptions").addEventListener("click", function(evt){
    if(evt.target.type === "radio"){
      if (evt.target.value == "link") {
        document.getElementById("updateBannerFile").style.display = 'none';
        document.getElementById("updateBannerLink").style.display = 'flex';
        document.getElementById("updateBannerLink").placeholder = 'Banner picture link';
      } else if (evt.target.value == "file") {
        document.getElementById("updateBannerFile").style.display = 'flex';
        document.getElementById("updateBannerLink").style.display = 'none';
      } else if (evt.target.value == "keep" || evt.target.value == "remove") {
        document.getElementById("updateBannerFile").style.display = 'none';
        document.getElementById("updateBannerLink").style.display = 'flex';
        document.getElementById("updateBannerLink").placeholder = 'Keep this empty';
      }
    }
  });

  
// USERINPUT WITH MEDIA QUERY

// get query value in url from key
const urlParams = new URLSearchParams(window.location.search);

// set attribute selected on the chosen options

const sortOption = urlParams.get('sort');
if (sortOption != null) {
    document.getElementById(sortOption).selected = true
} else {
    if ('URLSearchParams' in window) {
        var searchParams = new URLSearchParams(window.location.search);
        searchParams.set("sort", "popular");
        window.location.search = searchParams.toString();
    }
}

const timeOption = urlParams.get('time');
if (timeOption != null) {
    document.getElementById(timeOption).selected = true
} else {
    if ('URLSearchParams' in window) {
        var searchParams = new URLSearchParams(window.location.search);
        searchParams.set("time", "all_time");
        window.location.search = searchParams.toString();
    }
}


document.querySelector("div.sort-select").addEventListener(`change`, (e) => {
    const select = e.target;
    const value = select.value;
    if ('URLSearchParams' in window) {
        var searchParams = new URLSearchParams(window.location.search);
        searchParams.set("sort", value);
        window.location.search = searchParams.toString();
    }
});

document.querySelector("div.time-select").addEventListener(`change`, (e) => {
    const select = e.target;
    const value = select.value;
    if ('URLSearchParams' in window) {
        var searchParams = new URLSearchParams(window.location.search);
        searchParams.set("time", value);
        window.location.search = searchParams.toString();
    }
});