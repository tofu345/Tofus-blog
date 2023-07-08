const baseURL = "http://localhost:8005";
const urlParams = new URLSearchParams(window.location.search);

if (!navigator.cookieEnabled) {
    console.error("Cookies disabled user authentication might break.");
}

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

// https://stackoverflow.com/a/4033310/17673872
function httpGetAsync(theUrl, callback) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            callback(xmlHttp.responseText);
    };
    xmlHttp.open("GET", theUrl, true); // true for asynchronous
    xmlHttp.send(null);
}

function parseTemplate(temp, data) {
    for (let i = 0; i + 1 < temp.length; i++) {
        // Get variable name in ${} and replace with data
        if (temp[i] == "$" && temp[i + 1] == "{") {
            let y = 0;
            let found = false;
            for (; y < temp.length; y++) {
                if (temp[y] == "}") {
                    found = true;
                    break;
                }
            }

            let variableName = temp.slice(i + 2, y);
            if (found && variableName in data) {
                temp =
                    temp.slice(0, i) + data[variableName] + temp.slice(y + 1);
            } else {
                console.error(`Variable does not exist in data\n`);
            }
        }
    }

    return temp;
}

// setCookie("firstName", "tofs", 7);
// console.log(document.cookie);
// console.log(getCookie("firstName"));
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
