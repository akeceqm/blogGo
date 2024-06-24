const author_id = document.getElementById("NickName");
const input = document.getElementById("inputComment");
if (author_id.title === "") {
    input.disabled = true;
    console.log("Не авторизован");
}

let button = document.getElementById("btnSendComment");
if (button) {
    // Функция, которая отправляет комментарий
    function sendComment() {
        const author_id = document.getElementById("NickName");
        const input = document.getElementById("inputComment");
        const url = new URL(window.location.href);
        const pathSegments = url.pathname.split("/");
        const post_id = pathSegments[pathSegments.length - 2];

        console.log("xxx")
        if (input && input.value) {
            let xhr = new XMLHttpRequest();
            xhr.open("POST", `/post/${post_id}/comments`);

            xhr.onload = function () {
                if (xhr.status >= 200 && xhr.status < 300) {
                    let response = JSON.parse(xhr.responseText);
                    console.log(response);
                    location.reload();
                } else {
                    console.error("Server error:", xhr.status, xhr.statusText);
                    location.reload();
                }
            };

            xhr.onerror = function () {
                // Ошибка сети
                console.error("Network error");
            };

            let commentData = JSON.stringify({
                text: input.value,
                author_id: author_id.title,
            });

            console.log(commentData);

            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.send(commentData);
            input.value = "";
        }
    }

    button.onclick = sendComment;

    const input = document.getElementById("inputComment");
    if (input) {
        input.addEventListener("keydown", function (event) {
            if (event.key === "Enter") {
                sendComment();
            }
        });
    }
}
