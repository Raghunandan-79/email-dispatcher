const form = document.getElementById("emailForm");
const responseBox = document.getElementById("response");
const API_URL = location.hostname === "127.0.0.1" || location.hostname === "localhost"
  ? "http://127.0.0.1:8080"
  : "https://email-dispatcher-production-0621.up.railway.app/";


const bodyInput = document.querySelector("textarea");

bodyInput.addEventListener("input", function() {
    this.style.height = "auto";
    this.style.height = this.scrollHeight + "px";
});


form.addEventListener("submit", async function(e) {
    e.preventDefault();

    responseBox.innerText = "Sending...";
    responseBox.style.color = "blue";

    try {
        const formData = new FormData(form);

        const res = await fetch("http://{API_URL}/send", {
            method: "POST",
            body: formData
        });

        if (!res.ok) {
            const text = await res.text();
            responseBox.innerText = "Server Error: " + text;
            responseBox.style.color = "red";
            return;
        }

        const data = await res.json();

        if (data.error) {
            responseBox.innerText = "Error: " + data.error;
            responseBox.style.color = "red";
        } else {
            responseBox.innerText = data.message;
            responseBox.style.color = "green";
        }

    } catch (err) {
        responseBox.innerText = "Network Error: " + err.message;
        responseBox.style.color = "red";
    }
});

function applyTheme() {
    const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    document.documentElement.setAttribute("data-theme", prefersDark ? "dark" : "light");
}

applyTheme();

window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', applyTheme);

