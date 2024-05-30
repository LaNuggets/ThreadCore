var search = document.getElementById("search");

// Execute a function when the user presses a key on the keyboard
search.addEventListener("keypress", function(event) {
  // If the user presses the "Enter" key on the keyboard
  if (event.key === "Enter") {
    event.preventDefault();
    let searchString = document.forms["search"]["search"].value
    document.location.href="/search/" + searchString
  }
});