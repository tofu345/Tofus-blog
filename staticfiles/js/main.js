const baseURL = "http://localhost:8005";
const urlParams = new URLSearchParams(window.location.search);

if (!navigator.cookieEnabled) {
    console.error("Cookies disabled user authentication might break.");
}

let dropdowns = {};

function dropdownSetup(toggle) {
    let dropdownToggle = document.getElementById(toggle);
    if (dropdownToggle) {
        dropdownToggle.addEventListener("click", function () {
            let dropdown = document.getElementById(this.dataset.content);
            if (dropdown.style.display == "block") {
                dropdown.style.display = "";
            } else {
                dropdown.style.display = "block";
            }
        });

        dropdowns[toggle] = dropdownToggle.dataset.content;
    } else {
        console.error(`Object ${toggle} not found.`);
    }
}

window.onclick = function (event) {
    for (const key in dropdowns) {
        if (event.target != document.getElementById(key)) {
            document.getElementById(dropdowns[key]).style.display = "";
        }
    }
};

function setCookie(name, value, daysToLive) {
    const date = new Date();
    // Convert date to milliseconds
    date.setTime(date.getTime() + daysToLive * 24 * 60 * 60 * 1000);
    let expires = "expires=" + date.toUTCString();
    document.cookie = `${name}=${value}; expires=${expires}; path=/; SameSite=Lax`;
}

function deleteCookie(name) {
    setCookie(name, null, null);
}

function logout() {
    deleteCookie("token");
    window.location.href = baseURL + "/login";
}

function getCookie(name) {
    const cookieString = decodeURIComponent(document.cookie);
    const cookies = cookieString.split("; ");
    let result = null;

    cookies.forEach((element) => {
        if (element.indexOf(name) == 0) {
            result = element.substring(name.length + 1);
        }
    });

    return result;
}

// setCookie("firstName", "tofs", 7);
// console.log(document.cookie);
// console.log(getCookie("firstName"));
// console.log(getCookie("lastName"));

/*
    Ajax
    fetch(`/posts/${id}/likes`, {
        method: "POST",
        body: JSON.stringify({ vote: "like" }),
    })
        .then((res) => res.json())
        .then((res) => {
            if (res.responseCode != 100) {
                console.error(`Error: ${res.data}`);
                return;
            }

            updateLikes("+", elementId);
        })
        .catch((error) => {
            console.error("Error: " + error);
        });
 */

dropdownSetup("user-dropdown-toggle");
