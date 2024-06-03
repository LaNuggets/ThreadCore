<<<<<<< HEAD
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
=======
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
const images = document.querySelectorAll('img');

images.forEach(img => {
  img.addEventListener('error', function handleError() {
    const defaultImage =
      'https://bobbyhadz.com/images/blog/javascript-show-div-on-select-option/banner.webp';

    img.src = defaultImage;
    img.alt = 'default';
  });
});
>>>>>>> main
