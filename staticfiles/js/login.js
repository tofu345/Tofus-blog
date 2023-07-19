const form = document.getElementById("login-form");
const errorText = document.getElementById("login-error");
const signinBtn = document.getElementById("signin-btn");

form.addEventListener("submit", validateLoginForm);

function validateLoginForm(event) {
    event.preventDefault();
    errorText.innerHTML = "";
    signinBtn.disabled = true;
    const formData = new FormData(event.target);
    var object = {};
    formData.forEach((value, key) => (object[key] = value));

    if (!object["email"]) {
        errorText.innerHTML = "Email Address is required";
        signinBtn.disabled = false;
        return;
    }
    if (!object["password"]) {
        errorText.innerHTML = "Password is required";
        signinBtn.disabled = false;
        return;
    }

    fetch("/api/token", {
        method: "POST",
        body: JSON.stringify(object),
    })
        .then((res) => {
            signinBtn.disabled = false;
            return res.json();
        })
        .then((res) => {
            if (res.responseCode != 100) {
                errorText.innerHTML = res.data;
                return;
            }

            setCookie("access", res.access, 7);
            setCookie("refresh", res.refresh, 7);

            if (urlParams.get("next")) {
                window.location.href = baseURL + urlParams.get("next");
            } else {
                window.location.href = baseURL;
            }
        })
        .catch((error) => {
            console.error("Error: " + error);
            signinBtn.disabled = false;
        });
}
