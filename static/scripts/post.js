if(document.getElementById("updatePostButton")){
    document.getElementById("updatePostButton").onclick= function() {
        document.getElementById("updatePostForm").style.display = 'grid';
    }
  }
    
  document.addEventListener("click", (evt) => {
      const formEl = document.getElementById("updatePostForm");
      let targetEl = evt.target; // clicked element
      if(targetEl == formEl) {
        document.getElementById("updatePostForm").style.display = 'none';
      }
  });
  
  if(document.getElementById("deletePostButton")){
    document.getElementById("deletePostButton").onclick= function() {
      document.getElementById("deletePostForm").style.display = 'grid';
    }
  }
    
  document.addEventListener("click", (evt) => {
      const formEl = document.getElementById("deletePostForm");
      let targetEl = evt.target; // clicked element
      if(targetEl == formEl) {
        document.getElementById("deletePostForm").style.display = 'none';
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
    