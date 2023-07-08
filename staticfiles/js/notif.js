const notifWrapperId = "notif-wrapper";
const notifTypes = ["info", "error"];

function hideMessage(id) {
    try {
        const ele = document.getElementById("notif-" + id);
        ele.style.animation = "slide-out ease-in 0.5s forwards";
        setTimeout(() => {
            ele.remove();
        }, 300);
    } catch {}
}

function msgTemplate(id, type, msg) {
    return `
    <div
        id="notif-${id}"
        class="notif notif-${type} drop-shadow-xl"
        style="
            transform: translateX(38rem);
            opacity: 0;
            animation: slide-in cubic-bezier(0.43, 0.69, 0.29, 1) 0.5s forwards;
        "
    >
        <img 
            src="/static/svg/close.svg"
            class="notif-close"
            onclick="hideMessage('${id}')"
        />
        <p class="notif-msg">${msg}</p>
    </div>`;
}

function createNotif(msg, type = "info", timeout) {
    let msgWrapper = document.getElementById(notifWrapperId);
    if (!msgWrapper) {
        console.error(`Object ${notifWrapperId} not found`);
    }
    if (!notifTypes.includes(type)) {
        type = "info";
        console.warn(`${type} is not a valid notif type`);
    }
    if (!msg) {
        msg = `Lorem ipsum dolor sit amet, consectetur adipisicing elit. 
            Adipisc ipsam, earum nam expedita alias aspernatur id harum fugit nemo iure
            ipsa? Aliquam sit nobis voluptatibus iusto animi illum odit! Soluta?`;
    }
    if (msg.length > 10000) {
        console.warn("! Message Content too long\n" + msg);
    }

    let id = Math.floor(Math.random() * 1000);
    msgWrapper.innerHTML = msgTemplate(id, type, msg) + msgWrapper.innerHTML;
    setTimeout(() => {
        const ele = document.getElementById(`notif-${id}`);
        ele.style.transform = "translateX(0)";
        ele.style.opacity = "1";
        ele.style.animation = "";
        if (timeout) {
            setTimeout(() => hideMessage("msg-" + id), timeout * 1000);
        }
    }, 600);
}

function createErrorNotif(msg, timeout) {
    createNotif(msg, "error", timeout);
}
