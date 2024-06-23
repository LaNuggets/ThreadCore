// USERINPUT WITH MEDIA QUERY

// get query value in url from key
const urlParams = new URLSearchParams(window.location.search);

// set attribute selected on the chosen options
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

document.querySelector("div.time-select").addEventListener(`change`, (e) => {
    const select = e.target;
    const value = select.value;
    if ('URLSearchParams' in window) {
        var searchParams = new URLSearchParams(window.location.search);
        searchParams.set("time", value);
        window.location.search = searchParams.toString();
    }
});