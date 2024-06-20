document.addEventListener("DOMContentLoaded", function () {
    const currentUrl = window.location.pathname;
    const urlSegments = currentUrl.split('/');
    const userId = urlSegments[urlSegments.length - 1];

    if (userId){
        fetch(`/api/user/${userId}`)
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Не удалось загрузить данные пользователя");
                }
                return response.json();
            }) .then((data) => {
            const avatarImg = document.querySelector(".avatar");

            if (!data.avatar) {
                avatarImg.src = "assets/img/avatar.svg";
                avatarImg.alt = "Default Avatar";
            } else {
                avatarImg.src = data.avatar;
                avatarImg.alt = "Avatar";
            }

            avatarImg.onerror = function () {
                avatarImg.src = "assets/img/avatar.svg";
                avatarImg.alt = "Default Avatar";
            };

        })
            .catch((error) => {
                console.error("Ошибка загрузки данных о пользователе", error);
            });
    }
});
