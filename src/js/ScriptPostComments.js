let button = document.getElementById("btnSendComment");
if (button) {
    // Функция, которая отправляет комментарий
    function sendComment() {
        const author_id = document.getElementById("NickName");
        const input = document.getElementById("inputComment");
        const url = new URL(window.location.href);
        const pathSegments = url.pathname.split("/");
        const post_id = pathSegments[pathSegments.length - 2];

        if (input) {
            let xhr = new XMLHttpRequest();
            xhr.open("POST", `http://localhost:8080/post/${post_id}/comments`);
            xhr.onload = function (e) {
                console.log(e);
            };

            console.log(
                JSON.stringify({
                    text: input.value,
                    author_id: author_id.textContent,
                })
            );

            //xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
            xhr.send(
                JSON.stringify({
                    text: input.value,
                    author_id: author_id.textContent,
                })
            );
            location.reload();
        }
    }

    // Привязываем функцию к кнопке
    button.onclick = sendComment;

    // Привязываем функцию к полю ввода для события keydown
    let input = document.getElementById("inputComment");
    if (input) {
        input.addEventListener("keydown", function (event) {
            if (event.key === "Enter") {
                sendComment();
            }
        });
    }
}
