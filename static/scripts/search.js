// get query value in url from key
const urlParams = new URLSearchParams(window.location.search);

// set attribute selected on the chosen options
const mediaOption = urlParams.get('media');
if (mediaOption != null) {
    document.getElementById(mediaOption).checked = true
} else {
    if ('URLSearchParams' in window) {
        var searchParams = new URLSearchParams(window.location.search);
        searchParams.set("media", "posts");
        window.location.search = searchParams.toString();
    }
}

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


document.querySelector("div.media-radio").addEventListener("click", function(evt){
    if(evt.target.type === "radio"){

        if ('URLSearchParams' in window) {
            var searchParams = new URLSearchParams(window.location.search);
            searchParams.set("media", evt.target.value);
            window.location.search = searchParams.toString();
        }
    }
});

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