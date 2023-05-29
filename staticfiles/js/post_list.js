const dropdown_checkboxes =
    document.getElementsByClassName("dropdown-checkbox");

for (let i = 0; i < dropdown_checkboxes.length; i++) {
    const ele = dropdown_checkboxes[i];
    ele.addEventListener("change", (e) => {
        parent = e.target.parentNode;

        if (e.target.checked) {
            parent.style.display = "block";
        } else {
            parent.style.display = "";
        }
    });
}

function formatLikes(likes) {
    // < 1k
    if (likes <= 999) {
        return likes;
    }

    // < 1m
    if (likes <= 999999) {
        return `${Math.floor(likes / 1000)}K`;
    }

    return `${Math.floor(likes / 1000000)}M`;
}

let postLikes = document.getElementsByClassName("post-like");
for (let item of postLikes) {
    item.innerHTML = formatLikes(item.innerHTML);
}

let postBodyList = document.getElementsByClassName("post-body");
for (let item of postBodyList) {
    let len = item.innerHTML.length;
    let maxLen = 500;
    item.innerHTML = item.innerHTML.substring(0, maxLen);
    if (len > maxLen) {
        item.innerHTML += "...";
    }
}
