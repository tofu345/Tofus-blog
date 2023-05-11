const baseURL = "http://localhost:8005";
const urlParams = new URLSearchParams(window.location.search);

if (!navigator.cookieEnabled) {
    console.error(
        "Error: Cookies not enabled User authentication will not work"
    );
}

function registerDropdown(trigger, content) {
    let dropdown = document.getElementById(content);
    let dropdownToggle = document.getElementById(trigger);
    if (dropdownToggle && dropdownToggle) {
        dropdownToggle.addEventListener("click", () => {
            if (dropdown.style.display == "block") {
                dropdown.style.display = "";
            } else {
                dropdown.style.display = "block";
            }
        });
        window.onclick = function (event) {
            if (event.target != dropdownToggle) {
                dropdown.style.display = "";
            }
        };
    } else {
        console.error(`Error: ${trigger} and ${content} not found`);
    }
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

// setCookie("firstName", "tofs", 7);
// console.log(document.cookie);
// console.log(getCookie("firstName"));
// console.log(getCookie("lastName"));

registerDropdown("user-dropdown-toggle", "user-dropdown-content");
