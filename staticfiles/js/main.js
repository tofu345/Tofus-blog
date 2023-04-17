const baseURL = "http://localhost:8005";

if (!navigator.cookieEnabled) {
    console.error("Error: Cookies not enabled");
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

setCookie("firstName", "tofs", 7);
setCookie("lastName", "ya", 7);

// console.log(document.cookie);

console.log(getCookie("firstName"));
console.log(getCookie("lastName"));
