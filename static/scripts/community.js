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

document.getElementById("updateCommunityButton").onclick= function() {
    document.getElementById("updateCommunityForm").style.display = 'grid';
}
  
document.addEventListener("click", (evt) => {
    const formEl = document.getElementById("updateCommunityForm");
    let targetEl = evt.target; // clicked element
    if(targetEl == formEl) {
      document.getElementById("updateCommunityForm").style.display = 'none';
    }
});

document.getElementById("deleteCommunityButton").onclick= function() {
    document.getElementById("deleteCommunityForm").style.display = 'grid';
  }
  
document.addEventListener("click", (evt) => {
    const formEl = document.getElementById("deleteCommunityForm");
    let targetEl = evt.target; // clicked element
    if(targetEl == formEl) {
      document.getElementById("deleteCommunityForm").style.display = 'none';
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