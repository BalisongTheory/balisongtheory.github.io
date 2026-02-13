window.onload = main;

function main() {
    const videos = document.querySelectorAll(".trick-video");

    document.querySelectorAll(".video-speed-btn").forEach(button => {
        button.onclick = () => {
            videos.forEach(video => {
                video.playbackRate = parseFloat(button.value);
            });
        };
    });

    document.querySelectorAll("img").forEach(element => {
        element.onclick = () => {
            console.log(element.className);
            videos.forEach(video => {
                video.style.display = "none";
                video.pause();
                video.fastSeek(0);
            });
            document.querySelectorAll("#video-container .video-credit").forEach(credit => {
                credit.style.display = "none";
            });

            document.querySelector("#video-container ." + element.className).style.display = "block";
            document.querySelector("#video-container ." + element.className + "-credit").style.display = "block";
        };
    });

    document.querySelectorAll(".combo").forEach(element => {
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
