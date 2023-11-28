// ****************************Input range  Animation************************************


window.onload = function () {
    closeNav();
    slideLeft();
    slideRight();
    slideLeftAlbum();
    slideRightAlbum();
}

let sliderLeft = document.getElementById("slider-left");
let sliderRight = document.getElementById("slider-right");
let ValueLeft = document.getElementById("range-left");
let ValueRight = document.getElementById("range-right");
let minGap, minGap2 = 0;
let sliderTrack = document.querySelector(".slider-track");
let sliderLeftAlbum = document.getElementById("slider-left-album");
let sliderRightAlbum = document.getElementById("slider-right-album");
let ValueLeftAlbum = document.getElementById("range-left-album");
let ValueRightAlbum = document.getElementById("range-right-album");
let sliderMaxValueAlbum = document.getElementById("slider-left-album").max;
let sliderTrackAlbum = document.querySelector(".slider-track-album");

function slideLeft() {
    if (parseInt(sliderRight.value) - parseInt(sliderLeft.value) <= minGap) {
        sliderLeft.value = parseInt(sliderRight.value) - minGap;
    }
    ValueLeft.textContent = sliderLeft.value;
    fillColor();
}
function slideRight() {
    if (parseInt(sliderRight.value) - parseInt(sliderLeft.value) <= minGap) {
        sliderRight.value = parseInt(sliderLeft.value) + minGap;
    }
    ValueRight.textContent = sliderRight.value;
    fillColor();
}
function fillColor() {
    sliderTrack.style.background = "#f5fced";
}

function slideLeftAlbum() {
    if (parseInt(sliderRightAlbum.value) - parseInt(sliderLeftAlbum.value) <= minGap2) {
        sliderLeftAlbum.value = parseInt(sliderRightAlbum.value) - minGap2;
    }
    ValueLeftAlbum.textContent = sliderLeftAlbum.value;
    fillColorAlbum();
}
function slideRightAlbum() {
    if (parseInt(sliderRightAlbum.value) - parseInt(sliderLeftAlbum.value) <= minGap2) {
        sliderRightAlbum.value = parseInt(sliderLeftAlbum.value) + minGap2;
    }
    ValueRightAlbum.textContent = sliderRightAlbum.value;
    fillColorAlbum();
}
function fillColorAlbum() {
    sliderTrackAlbum.style.background = "#f5fced";
}


function openNav() {
    document.getElementById("filters").style.width = "400px";
}

function closeNav() {
    document.getElementById("filters").style.width = "0";
}