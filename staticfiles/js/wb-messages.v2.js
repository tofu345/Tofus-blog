const messageStyles = `
@import url("https://fonts.googleapis.com/css2?family=Mulish:ital,wght@0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap");
#wb-msg-wrapper {
    position: fixed;
    top: 0;
    right: 0;
    z-index: 50;
}
.wb-error-msg {
    background: linear-gradient(
        180deg,
        rgba(250, 240, 234, 0.78) 0%,
        rgba(252, 239, 234, 0.99) 19.26%
    );
    border: 1.3183px solid #de6d69;
    box-shadow: inset 0px 0px 2.63659px rgba(255, 255, 255, 0.63);
    backdrop-filter: blur(11.2055px);
    border-radius: 26.3659px;
}
.wb-success-msg {
    background: linear-gradient(180deg, #eefdf7 0%, #ebfdf5 19.26%);
    border: 1.3183px solid #76ad94;
    box-shadow: inset 0px 0px 2.63659px rgba(255, 255, 255, 0.63);
    backdrop-filter: blur(11.2055px);
    border-radius: 26.3659px;
}
.wb-info-msg {
    background: linear-gradient(180deg, #e7eefa 0%, #e7eefa 19.26%);
    border: 1.3183px solid #9eb1cb;
    box-shadow: inset 0px 0px 2.63659px rgba(255, 255, 255, 0.63);
    backdrop-filter: blur(11.2055px);
    border-radius: 26.3659px;
}
.wb-warning-msg {
    background: linear-gradient(180deg, #fef7ea 0%, #fef7ea 19.26%);
    border: 1.3183px solid #e3bc74;
    box-shadow: inset 0px 0px 2.63659px rgba(255, 255, 255, 0.63);
    backdrop-filter: blur(11.2055px);
    border-radius: 26.3659px;
}
.wb-msg {
    transform: translateX(0);
    font-family: "Mulish" !important;
    z-index: 50;
    margin: 0.5rem;
    padding: 15px;
    border-radius: 10px;
    width: 34rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    overflow: auto;
    position: relative;
}
.wb-msg-svg {
    margin: 0.5rem 1rem 0.5rem 0.5rem;
    filter: drop-shadow(0px 13.183px 27.0251px rgba(209, 63, 0, 0.14));
    flex: none;
    order: 0;
    flex-grow: 0;
    justify-self: center;
}
.wb-msg-close {
    position: absolute;
    top: 7px;
    right: 5px;
    margin: 10px;
    border-radius: 50%;
    display: none;
    cursor: pointer;
}

.wb-msg:hover .wb-msg-close {
    display: block;
}

.wb-msg-content {
    align-content: space-around;
    height: 100%;
    width: 100%;
}
.wb-msg-title {
    margin: 0 0 0.25rem 0;
    font-style: normal;
    font-weight: 600;
    font-size: 16px;
    line-height: 19px;
}
.wb-msg-body {
    margin: 0;
    padding: 0;
    font-style: normal;
    font-weight: 400;
    font-size: 14px;
    line-height: 19px;
}
@keyframes slide-out {
    50% {
        transform: translateX(38rem);
        opacity: 0;
    }
    100% {
        transform: translateX(38rem);
        margin: 0;
        height: 0px;
        padding: 0;
    }
}
@keyframes slide-in {
    0% {
        transform: translateX(38rem);
    }
    100% {
        transform: translateX(0);
        opacity: 1;
    }
}
`;

const errorSvg = `
<svg class="wb-msg-svg" width="48" height="49" viewBox="0 0 48 49" fill="none">
    <circle cx="23.8446" cy="24.5001" r="23.7293" fill="#DD5A56" />
    <path
        fill-rule="evenodd"
        clip-rule="evenodd"
        d="M23.8445 39.0018C31.8533 39.0018 38.3458 32.5093 38.3458 24.5005C38.3458 16.4917 31.8533 9.99927 23.8445 9.99927C15.8357 9.99927 9.34326 16.4917 9.34326 24.5005C9.34326 32.5093 15.8357 39.0018 23.8445 39.0018ZM18.6087 29.7372C18.0705 29.1991 18.0705 28.3266 18.6087 27.7884L21.6943 24.7029L18.6087 21.6173C18.0705 21.0791 18.0705 20.2066 18.6087 19.6685C19.1468 19.1304 20.0193 19.1304 20.5575 19.6685L23.643 22.7541L26.7286 19.6685C27.2668 19.1304 28.1393 19.1304 28.6774 19.6685C29.2156 20.2066 29.2156 21.0791 28.6774 21.6173L25.5918 24.7029L28.6774 27.7884C29.2156 28.3266 29.2156 29.1991 28.6774 29.7372C28.1393 30.2754 27.2668 30.2754 26.7286 29.7372L23.643 26.6516L20.5575 29.7372C20.0193 30.2754 19.1468 30.2754 18.6087 29.7372Z"
        fill="white"
    />
</svg>`;

const successSvg = `
<svg class="wb-msg-svg" width="48" height="49" viewBox="0 0 48 49" fill="none">
    <circle cx="23.8446" cy="24.5001" r="23.7293" fill="#38CB89" />
    <ellipse cx="24" cy="23.417" rx="5.83333" ry="5.25" fill="#38CB89" />
    <path
        d="M31.6301 14.8068L25.2134 12.4035C24.5484 12.1585 23.4634 12.1585 22.7984 12.4035L16.3817 14.8068C15.1451 15.2735 14.1417 16.7201 14.1417 18.0385V27.4885C14.1417 28.4335 14.7601 29.6818 15.5184 30.2418L21.9351 35.0368C23.0667 35.8885 24.9217 35.8885 26.0534 35.0368L32.4701 30.2418C33.2284 29.6701 33.8467 28.4335 33.8467 27.4885V18.0385C33.8584 16.7201 32.8551 15.2735 31.6301 14.8068ZM28.0601 21.3401L23.0434 26.3568C22.8684 26.5318 22.6467 26.6135 22.4251 26.6135C22.2034 26.6135 21.9817 26.5318 21.8067 26.3568L19.9401 24.4668C19.6017 24.1285 19.6017 23.5685 19.9401 23.2301C20.2784 22.8918 20.8384 22.8918 21.1767 23.2301L22.4367 24.4901L26.8351 20.0918C27.1734 19.7535 27.7334 19.7535 28.0717 20.0918C28.4101 20.4301 28.4101 21.0018 28.0601 21.3401Z"
        fill="white"
    />
</svg>`;

const warningSvg = `
<svg class="wb-msg-svg" width="48" height="49" viewBox="0 0 48 49" fill="none">
    <circle cx="23.8446" cy="24.5001" r="23.7293" fill="#FDBE20" />
    <path
        d="M38.3458 24.5005C38.3458 32.5093 31.8533 39.0018 23.8445 39.0018C15.8357 39.0018 9.34326 32.5093 9.34326 24.5005C9.34326 16.4917 15.8357 9.99927 23.8445 9.99927C31.8533 9.99927 38.3458 16.4917 38.3458 24.5005Z"
        fill="white"
    />
    <path
        d="M16.5373 31.6791C18.1727 33.3154 20.3248 34.3343 22.6272 34.5619C24.9295 34.7894 27.2394 34.2116 29.1636 32.9269C31.0876 31.6423 32.5068 29.7303 33.1792 27.5165C33.8517 25.303 33.7358 22.9245 32.8513 20.787C31.9667 18.6493 30.3684 16.8841 28.3286 15.793C26.2887 14.7017 23.9335 14.3513 21.6642 14.8017C19.3949 15.2521 17.3518 16.4754 15.8836 18.2631C14.4152 20.0508 13.612 22.2925 13.611 24.606C13.6079 25.9199 13.8649 27.2214 14.3671 28.4354C14.8695 29.6494 15.6068 30.7516 16.5373 31.6791ZM17.484 18.479C18.5504 17.3628 19.8955 16.5517 21.3804 16.1296C22.8653 15.7075 24.4358 15.6895 25.9298 16.0778C27.4238 16.466 28.7871 17.2461 29.8787 18.3376C30.9702 19.4293 31.7503 20.7925 32.1385 22.2866C32.5268 23.7807 32.5088 25.3512 32.0868 26.836C31.6646 28.3208 30.8536 29.6659 29.7373 30.7324C28.671 31.8486 27.3259 32.6596 25.841 33.0818C24.3561 33.5038 22.7856 33.5218 21.2916 33.1336C19.7976 32.7453 18.4343 31.9653 17.3427 30.8738C16.2512 29.7821 15.4711 28.4188 15.0829 26.9248C14.6946 25.4307 14.7126 23.8602 15.1346 22.3754C15.5568 20.8906 16.3678 19.5455 17.484 18.479ZM22.9439 25.9389V19.892C22.9439 19.6539 23.071 19.4338 23.2772 19.3147C23.4835 19.1955 23.7376 19.1955 23.9439 19.3147C24.1501 19.4338 24.2772 19.6539 24.2772 19.892V25.9389C24.2772 26.1771 24.1501 26.3972 23.9439 26.5162C23.7376 26.6353 23.4834 26.6353 23.2772 26.5162C23.071 26.3972 22.9439 26.1771 22.9439 25.9389ZM22.7239 29.2722C22.7239 29.0371 22.8172 28.8115 22.9836 28.6453C23.1498 28.4789 23.3754 28.3856 23.6105 28.3856C23.8457 28.3856 24.0713 28.4789 24.2375 28.6453C24.4038 28.8115 24.4971 29.0371 24.4971 29.2722C24.4971 29.5073 24.4038 29.7329 24.2375 29.8992C24.0712 30.0655 23.8457 30.1588 23.6105 30.1588C23.3748 30.1606 23.1483 30.0679 22.9815 29.9012C22.8148 29.7344 22.7221 29.508 22.7239 29.2722Z"
        fill="#FDBE20"
    />
</svg>`;

const infoSvg = `
  <svg width="48" class="wb-msg-svg" height="49" viewBox="0 0 48 49" fill="none">
    <circle cx="23.8446" cy="24.5001" r="23.7293" fill="#3087E9" />
    <path
        d="M38.3458 24.5005C38.3458 32.5093 31.8533 39.0018 23.8445 39.0018C15.8357 39.0018 9.34326 32.5093 9.34326 24.5005C9.34326 16.4917 15.8357 9.99927 23.8445 9.99927C31.8533 9.99927 38.3458 16.4917 38.3458 24.5005Z"
        fill="white"
    />
    <path
        d="M30.6829 17.5411C29.0475 15.9048 26.8954 14.8859 24.593 14.6584C22.2907 14.4308 19.9808 15.0087 18.0566 16.2933C16.1326 17.5779 14.7134 19.4899 14.041 21.7037C13.3685 23.9172 13.4844 26.2957 14.369 28.4332C15.2535 30.5709 16.8519 32.3361 18.8916 33.4272C20.9315 34.5185 23.2867 34.8689 25.556 34.4186C27.8254 33.9681 29.8684 32.7449 31.3366 30.9571C32.8051 29.1694 33.6082 26.9277 33.6092 24.6142C33.6123 23.3003 33.3553 21.9988 32.8531 20.7848C32.3507 19.5708 31.6134 18.4686 30.6829 17.5411ZM29.7362 30.7412C28.6698 31.8575 27.3247 32.6685 25.8398 33.0907C24.3549 33.5127 22.7844 33.5307 21.2904 33.1424C19.7964 32.7542 18.4331 31.9741 17.3415 30.8826C16.25 29.7909 15.4699 28.4277 15.0817 26.9336C14.6934 25.4396 14.7114 23.869 15.1335 22.3843C15.5556 20.8994 16.3666 19.5543 17.4829 18.4878C18.5492 17.3716 19.8943 16.5606 21.3793 16.1384C22.8641 15.7164 24.4346 15.6984 25.9286 16.0866C27.4227 16.4749 28.7859 17.2549 29.8776 18.3465C30.9691 19.4381 31.7491 20.8014 32.1374 22.2954C32.5256 23.7895 32.5076 25.36 32.0856 26.8448C31.6634 28.3297 30.8524 29.6747 29.7362 30.7412ZM24.2763 23.2814V29.3283C24.2763 29.5664 24.1493 29.7864 23.943 29.9055C23.7368 30.0247 23.4826 30.0247 23.2763 29.9055C23.0701 29.7864 22.943 29.5663 22.943 29.3283V23.2814C22.943 23.0431 23.0701 22.823 23.2763 22.704C23.4826 22.5849 23.7368 22.5849 23.943 22.704C24.1493 22.823 24.2763 23.0431 24.2763 23.2814ZM24.4963 19.948C24.4963 20.1831 24.403 20.4087 24.2366 20.5749C24.0704 20.7413 23.8448 20.8346 23.6097 20.8346C23.3746 20.8346 23.149 20.7413 22.9827 20.5749C22.8164 20.4087 22.7231 20.1831 22.7231 19.948C22.7231 19.7129 22.8164 19.4873 22.9827 19.3211C23.149 19.1547 23.3746 19.0614 23.6097 19.0614C23.8454 19.0596 24.0719 19.1523 24.2387 19.319C24.4054 19.4858 24.4981 19.7123 24.4963 19.948Z"
        fill="#3087E9"
    />
</svg>`;

var styleSheet = document.createElement("style");
styleSheet.innerText = messageStyles;
document.head.appendChild(styleSheet);

let msgWrapper = document.getElementById("wb-msg-wrapper");
if (!msgWrapper) {
    msgWrapper = document.createElement("div");
    msgWrapper.id = "wb-msg-wrapper";
    document.body.appendChild(msgWrapper);
}

function hideMessage(id) {
    try {
        const ele = document.getElementById(id);
        ele.style.transform = "translate-x(38rem)";
        ele.style.animation = "slide-out ease-in-out 0.5s forwards";
        setTimeout(() => {
            ele.remove();
        }, 600);
    } catch {}
}

function msgTemplate(id, title, content, className, svg) {
    return `
    <div
        id="msg-${id}"
        class="wb-msg ${className}"
        style="
            transform: translateX(38rem);
            opacity: 0;
            animation: slide-in cubic-bezier(0.43, 0.69, 0.29, 1) 0.5s forwards;
        "
    >
        ${svg}
        <div class="wb-msg-content">
            <p class="wb-msg-title">${title}</p>
            <p class="wb-msg-body">${content}</p>
        </div>
        <svg
            onclick="hideMessage('msg-${id}')"
            class="wb-msg-close"
            width="16"
            height="16"
            fill="black"
            viewBox="0 0 16 16"
        >
            <path
                d="M2.146 2.854a.5.5 0 1 1 .708-.708L8 7.293l5.146-5.147a.5.5 0 0 1 .708.708L8.707 8l5.147 5.146a.5.5 0 0 1-.708.708L8 8.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 8 2.146 2.854Z"
            />
        </svg>
    </div>`;
}

function createSimpleMessage(title, content, timeout, className, svg) {
    if (!className && !svg) {
        console.warn(
            "This function is for internal use only\nUse create(Success/Error/Warning/Info)Message functions instead"
        );
        return;
    }

    let msgWrapper = document.getElementById("wb-msg-wrapper");
    if (!msgWrapper) {
        console.error("Object 'wb-msg-wrapper' not found");
    }

    const id = Math.floor(Math.random() * 1000000000);
    let defaultTitle = "The title of the message";
    let defaultContent = `Lorem ipsum dolor sit amet, consectetur adipisicing elit. Adipisc
    ipsam, earum nam expedita alias aspernatur id harum fugit nemo iure
    ipsa? Aliquam sit nobis voluptatibus iusto animi illum odit! Soluta?`;

    if (!title) {
        title = defaultTitle;
    }
    if (!content) {
        content = defaultContent;
    }
    if (content.length > 10000) {
        console.warn("! Message Content too long\n" + content);
    }
    msgWrapper.innerHTML += msgTemplate(id, title, content, className, svg);
    setTimeout(() => {
        const ele = document.getElementById(`msg-${id}`);
        ele.style.transform = "translate(0, 0)";
        ele.style.opacity = "1";
        ele.style.animation = "";
        if (timeout) {
            setTimeout(() => hideMessage("msg-" + id), timeout * 1000);
        }
    }, 600);
}

function createInfoMessage(
    title,
    content,
    timeout,
    className = "wb-info-msg",
    svg = infoSvg
) {
    createSimpleMessage(title, content, timeout, className, svg);
}

function createWarningMessage(
    title,
    content,
    timeout,
    className = "wb-warning-msg",
    svg = warningSvg
) {
    createSimpleMessage(title, content, timeout, className, svg);
}

function createErrorMessage(
    title,
    content,
    timeout,
    className = "wb-error-msg",
    svg = errorSvg
) {
    createSimpleMessage(title, content, timeout, className, svg);
}

function createSuccessMessage(
    title,
    content,
    timeout,
    className = "wb-success-msg",
    svg = successSvg
) {
    createSimpleMessage(title, content, timeout, className, svg);
}
