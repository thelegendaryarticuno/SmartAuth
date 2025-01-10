const inputs = document.querySelectorAll(".input-field");
const toggle_btn = document.querySelectorAll(".toggle");
const main = document.querySelector("main");
const bullets = document.querySelectorAll(".bullets span");
const images = document.querySelectorAll(".image");
const registerForm = document.getElementById('register-form');
const loginForm = document.getElementById('login-form');
const modal = document.getElementById('response-modal');
const closeBtn = document.querySelector('.close-btn');
const responseMessage = document.getElementById('response-message');

// Handle focus and blur on input fields
inputs.forEach((inp) => {
  inp.addEventListener("focus", () => {
    inp.classList.add("active");
  });
  inp.addEventListener("blur", () => {
    if (inp.value != "") return;
    inp.classList.remove("active");
  });
});

// Toggle between Sign In and Sign Up forms
toggle_btn.forEach((btn) => {
  btn.addEventListener("click", () => {
    main.classList.toggle("sign-up-mode");
  });
});

// Slider function for the text slider and images
function moveSlider() {
  let index = this.dataset.value;

  let currentImage = document.querySelector(`.img-${index}`);
  images.forEach((img) => img.classList.remove("show"));
  currentImage.classList.add("show");

  const textSlider = document.querySelector(".text-group");
  textSlider.style.transform = `translateY(${-(index - 1) * 2.2}rem)`;

  bullets.forEach((bull) => bull.classList.remove("active"));
  this.classList.add("active");
}

bullets.forEach((bullet) => {
  bullet.addEventListener("click", moveSlider);
});

// Close modal when the close button is clicked
closeBtn.onclick = function () {
  modal.style.display = 'none';
}

// Close modal if the user clicks outside of the modal
window.onclick = function (event) {
  if (event.target === modal) {
    modal.style.display = 'none';
  }
}

// Handle Register form submission
registerForm.addEventListener('submit', function (event) {
  event.preventDefault(); // Prevent page reload on submit

  const formData = new FormData(registerForm);
  fetch('/register', {
    method: 'POST',
    body: formData
  })
    .then(response => response.json())
    .then(data => {
      responseMessage.textContent = data.message || data.error;
      modal.style.display = 'block';
    })
    .catch(error => {
      responseMessage.textContent = 'Something went wrong!';
      modal.style.display = 'block';
    });
});

// Handle Login form submission
loginForm.addEventListener('submit', function (event) {
  event.preventDefault(); // Prevent page reload on submit

  const formData = new FormData(loginForm);
  fetch('/login', {
    method: 'POST',
    body: formData
  })
    .then(response => response.json())
    .then(data => {
      responseMessage.textContent = data.message || data.error;
      modal.style.display = 'block';
    })
    .catch(error => {
      responseMessage.textContent = 'Something went wrong!';
      modal.style.display = 'block';
    });
});
