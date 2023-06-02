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

function truncateStr(string, maxLen) {
    if (!maxLen) {
        maxLen = 500;
    }
    let len = string.length;
    string = string.substring(0, maxLen);
    if (len > maxLen) {
        string += "...";
    }
    return string;
}

let postBody = document.getElementsByClassName("post-body");
for (let item of postBody) {
    item.innerHTML = truncateStr(item.innerHTML, 500);
}

const dateOptions = {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric",
};

document.addEventListener("DOMContentLoaded", function (event) {
    httpGetAsync("/api/posts", (res) => {
        let postList = JSON.parse(res);
        postList = postList.data;
        postList.sort((a, b) => b.id - a.id);

        let postsWrapper = document.getElementById("posts");
        if (postList.length != 0) {
            httpGetAsync("/static/lib/post.html", (res) => {
                postsWrapper.innerHTML = "";

                for (let key in postList) {
                    let post = postList[key];
                    postsWrapper.innerHTML += parseTemplate(res, {
                        id: post.id,
                        likes: formatLikes(post.likes),
                        slug: post.slug,
                        title: post.title,
                        body: truncateStr(post.body),
                        created_at: new Date(
                            post.created_at
                        ).toLocaleDateString("en-US", dateOptions),
                        author: post.author,
                    });
                }
            });
        } else if (postsWrapper) {
            postsWrapper.innerHTML = '<div class="card">No Posts Yet.</div>';
        }
    });
});
