document.addEventListener("DOMContentLoaded", function () {
    const urlParams = new URLSearchParams(window.location.search);
    const userId = urlParams.get("userId");
    if (userId) {
        console.log("User ID:", userId);
        fetch(`/api/user/${userId}`)
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Failed to load user data");
                }
                return response.json();
            })
            .then((data) => {
                console.log(data);
                const avatarImg = document.querySelector(".avatar");

                // Устанавливаем аватар по умолчанию, если поле пустое или не установлено
                if (!data.avatar || data.avatar === "") {
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

                const nameUser = document.getElementById("name-user");
                const registrationDate = document.getElementById("registration-date");
                const descriptionUser = document.getElementById("description-user");

                if (nameUser) nameUser.textContent = data.nick_name || "Unknown";
                if (registrationDate) registrationDate.textContent = data.date_registration || "Registration date unknown";
                if (descriptionUser) descriptionUser.textContent = data.description || "No description available";

                const btnEdit = document.querySelector(".btn-edit");
                if (btnEdit) {
                    btnEdit.addEventListener("click", function () {
                        window.location.href = "/changeProfile/" + userId;
                    });
                }
            })
            .catch((error) => {
                console.error("Failed to load user data", error);
            });
    }
});
