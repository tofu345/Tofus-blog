const baseURL = "http://localhost:8005";
const urlParams = new URLSearchParams(window.location.search);

if (!navigator.cookieEnabled) {
    console.error("Error: Cookies not enabled");
}

function logout() {
    deleteCookie("token");
    window.location.href = baseURL + "/login";
}

// User Dropdown
dropdownToggle = document.getElementById("user-dropdown-toggle");
if (dropdownToggle) {
    dropdownToggle.addEventListener("change", (e) => {
        dropdownContent = document.getElementById("user-dropdown-content");

        if (e.target.checked) {
            dropdownContent.style.display = "block";
        } else {
            dropdownContent.style.display = "";
        }
    });

    window.onclick = function (event) {
        if (!event.target.matches("#user-dropdown")) {
            dropdownContent = document.getElementById("user-dropdown-content");
            dropdownContent.style.display = "";
        }
    };
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
// setCookie("lastName", "ya", 7);

// console.log(document.cookie);

// console.log(getCookie("firstName"));
// console.log(getCookie("lastName"));
