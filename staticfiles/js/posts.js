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
