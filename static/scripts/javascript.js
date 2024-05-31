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
