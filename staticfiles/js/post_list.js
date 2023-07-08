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

const maxLen = 500;

function truncateStr(string) {
    let len = string.length;
    string = string.substring(0, maxLen);
    if (len > maxLen) {
        string += "...";
    }
    return string;
}

const dateOptions = {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric",
};

let postList = [];
let focusedPost;

document.addEventListener("DOMContentLoaded", function (event) {
    let postsElement = document.getElementById("posts");
    if (postsElement == null) {
        return;
    }

    httpGetAsync("/api/posts", (res) => {
        res = JSON.parse(res);
        if (res.responseCode != 100) {
            createErrorNotif("Error fetching posts");
            console.error(res);
            return;
        }

        postList = res.data;
        postList.sort((a, b) => a.id - b.id);

        // console.log(postList);

        if (postList.length == 0) {
            postsElement.innerHTML = '<div class="card">No Posts Yet.</div>';
            return;
        }

        httpGetAsync("/static/lib/post.html", (res) => {
            postsElement.innerHTML = "";
            postList.forEach((post) => {
                postsElement.innerHTML += parseTemplate(res, {
                    id: post.id,
                    likes: formatLikes(post.likes),
                    slug: post.slug,
                    title: post.title,
                    body: truncateStr(post.body),
                    created_at: new Date(post.created_at).toLocaleDateString(
                        "en-US",
                        dateOptions
                    ),
                    author: post.author,
                });
            });
        });
    });
});
