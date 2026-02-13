window.onload = main;

function main() {
    const input = document.querySelector("input");
    const video = document.querySelector("video");
    input.onchange = (event) => {
        console.log("test", input.value);
        video.src = input.value;
        video.play();
    };

    document.querySelectorAll(".video-speed-btn").forEach(button => {
        button.onclick = () => {
            video.playbackRate = parseFloat(button.value);
        };
    });
}
