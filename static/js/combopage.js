window.onload = main;

function main() {
    const videos = document.querySelectorAll(".combo-video");
    var currentVideo = document.querySelector("#video-container div video");

    const selectorButtons = document.querySelectorAll("#video-selector img");
    var currentSelectorButton = document.querySelector("#video-selector img");

    document.querySelectorAll(".video-speed-btn").forEach(button => {
        button.onclick = () => {
            videos.forEach(video => {
                video.playbackRate = parseFloat(button.value);
            });
        };
    });

    selectorButtons.forEach(element => {
        element.onclick = () => {
            currentSelectorButton.style.borderWidth = "0px";
            element.style.borderWidth = "3px";
            currentSelectorButton = element;

            currentVideo.pause();
            currentVideo.currentTime = 0;
            currentVideo.parentElement.style.display = "none";

            document.querySelector("#video-container ." + element.className).style.display = "block";
            currentVideo = document.querySelector("#video-container ." + element.className + " video");
        };
    });

    document.querySelectorAll(".trick").forEach(element => {
        const video = element.querySelector("video");
        video.pause();
        element.onmouseenter = () => {
            video.currentTime = 0;
            video.play();
        };

        element.onmouseleave = () => {
            video.pause();
            video.currentTime = 0;
        };
    });
}
