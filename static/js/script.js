const navButton = document.getElementById("nav-button");
const navBar = document.getElementById("navbar");
const navBackDrop = document.getElementById("nav-backdrop");

navButton.addEventListener("click", () => {
  navBackDrop.style.display = "block";
  navBar.classList.toggle("-translate-x-full");
});

navBackDrop.addEventListener("click", () => {
  navBackDrop.style.display = "none";
  navBar.classList.toggle("-translate-x-full");
});
