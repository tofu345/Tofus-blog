const form = document.getElementById("login-form");
const errorText = document.getElementById("login-error");

form.addEventListener("submit", validateLoginForm);

function validateLoginForm(event) {
    event.preventDefault();
    errorText.innerHTML = "";
    const formData = new FormData(event.target);
    var object = {};
    formData.forEach((value, key) => (object[key] = value));

    fetch("/login", {
        method: "POST",
        body: JSON.stringify(object),
    })
        .then((res) => res.json())
        .then((res) => {
            if (res.responseCode != 100) {
                console.error(`Error: ${res.data}`);

                if (res.data == "record not found") {
                    errorText.innerHTML =
                        "No user was found with entered email address and password";
                } else {
                    errorText.innerHTML = res.data;
                }
                return;
            }

            if (res.data == null) {
                console.error(`Error: No data in response`);
                return;
            }

            if (res.data.token == null) {
                console.error(`Error: No token in response`);
                return;
            }

            setCookie("token", res.data.token, 7);
        })
        .catch((error) => {
            console.error("Error: " + error);
        });
}
