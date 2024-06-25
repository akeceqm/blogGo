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

    const btnsEditPost = document.querySelectorAll('.BtnEditPost');
    const btnCloseEdit = document.getElementById('BtnCloseEdit')
    const popUpEdit = document.getElementById('Pop_upEdit')

    btnCloseEdit.onclick = () => {
        popUpEdit.classList.remove('active')
    }
    btnsEditPost.forEach((btn) => {
        btn.addEventListener('click', () => {
            popUpEdit.classList.add('active')
            const TitleInput = document.getElementById('TitleInputEdit')
            const TextInput = document.getElementById('TextInputEdit')
            const currentUrl = window.location.href;
            const url = new URL(currentUrl);
            const userId = url.searchParams.get('userId');
            let TitleAndText = btn.title.split('|')
            TitleInput.value = TitleAndText[0]
            TextInput.value = TitleAndText[1]
        });
    });
});