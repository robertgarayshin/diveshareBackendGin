const slides = document.querySelectorAll('.slide');
let currentSlide = [0, 1, 2];

const slidesOnPage = 3;
const style1 = window.getComputedStyle(document.querySelector('.slide.first'))
const style2 = window.getComputedStyle(document.querySelector('.slide.second'))
const style3 = window.getComputedStyle(document.querySelector('.slide.third'))


function changeSlideRight() {
    
    slides[currentSlide[0]].classList.remove('first')
    slides[currentSlide[1]].classList.remove('second')
    slides[currentSlide[2]].classList.remove('third')
    currentSlide[0] += 1
    currentSlide[1] += 1
    currentSlide[2] += 1
    if (currentSlide[2] == 5) {
        currentSlide[2] = 0
    }
    if (currentSlide[1] == 5) {
        currentSlide[1] = 0
    }
    if (currentSlide[0] == 5) {
        currentSlide[0] = 0
    }
    
    slides[currentSlide[2]].classList.add('third')
    slides[currentSlide[1]].classList.add('second')
    slides[currentSlide[0]].classList.add('first')
    }

    function changeSlideLeft() {
    
        slides[currentSlide[0]].classList.remove('first')
        slides[currentSlide[1]].classList.remove('second')
        slides[currentSlide[2]].classList.remove('third')
        currentSlide[0] -= 1
        currentSlide[1] -= 1
        currentSlide[2] -= 1
        console.log(currentSlide)
        if (currentSlide[2] == -1) {
            currentSlide[2] = 4
        }
        if (currentSlide[1] == -1) {
            currentSlide[1] = 4
        }
        if (currentSlide[0] == -1) {
            currentSlide[0] = 4
        }
        console.log(currentSlide)
        slides[currentSlide[0]].classList.add('first')
        slides[currentSlide[1]].classList.add('second')
        slides[currentSlide[2]].classList.add('third')
        }
    

document.getElementById('arrowLeft').addEventListener('click', () => {
    changeSlideLeft();
    console.log(currentSlide[0], currentSlide[1], currentSlide[2])
});

document.getElementById('arrowRight').addEventListener('click', () => {
    changeSlideRight();
    console.log(currentSlide[0], currentSlide[1], currentSlide[2])
});

