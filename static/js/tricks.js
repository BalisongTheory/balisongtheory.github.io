window.onload = main;

function main() {
    const searchInput = document.getElementById("trick-search");
    const trickList = document.querySelectorAll(".trick");
    searchInput.onkeyup = () => {
        trickList.forEach(trick => {
            const name = trick.querySelector(".trick-name");
            const tags = trick.querySelector(".trick-tags");
            const filter = searchInput.value.trim();
            let resultElement;
            if (name.innerText.toLowerCase().includes(filter.toLowerCase())) {
                resultElement = name;
                tags.innerHTML = tags.innerText;
                tags.style.display = "none";
            } else if (tags.innerText.toLowerCase().includes(filter.toLowerCase())) {
                resultElement = tags;
                name.innerHTML = name.innerText;
                tags.style.display = "block";
            }

            if (typeof resultElement !== "undefined") {
                trick.style.display = "";

                resultElement.innerHTML = highlightText(resultElement.innerText, filter);
            } else {
                trick.style.display = "none";
                name.innerHTML = name.innerText;
                tags.innerHTML = tags.innerText;
            }
        });
    }

    trickList.forEach(element => {
        const video = element.querySelector("video");
        video.pause();
        element.onmouseenter = () => {
            video.fastSeek(0);
            video.play();
        };

        element.onmouseleave = () => {
            video.pause();
            video.fastSeek(0);
        };
    });
}

function highlightText(text, highlightText) {
    let index = text.toLowerCase().indexOf(highlightText);
    if (index == -1) {
        return text
    }
    let length = highlightText.length;
    let start = text.substring(0, index);
    let middle = text.substring(index, index + length);
    let end = text.substring(index + length);
    return `${start}<mark>${middle}</mark>${end}`;
}
