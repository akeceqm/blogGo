document.addEventListener("DOMContentLoaded", function () {
  const urlParams = new URLSearchParams(window.location.search);
  const userId = urlParams.get("userId");

  if (userId) {
    fetch(`/api/user/${userId}`)
        .then((response) => {
          if (!response.ok) {
            throw new Error("Не удалось загрузить данные пользователя");
          }
          return response.json();
        })
        .then((data) => {
          const avatarImg = document.querySelector(".avatar");

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

          console.log(data); // Отладочная информация

          const nameUser = document.getElementById("name-user");
          const registrationDate = document.getElementById("registration-date");
          const descriptionUser = document.getElementById("description-user");

          // Преобразуем строку в объект, если необходимо
          let nickName = data.nick_name;

          if (typeof nickName === 'string') {
            try {
              nickName = JSON.parse(nickName); // Попытка преобразования строки в объект
            } catch (e) {
              // Обработка ошибки парсинга, если это не объект JSON
              console.error('Ошибка парсинга никнейма:', e);
            }
          }

          // Если nickName все еще строка, просто установим ее как текст
          if (typeof nickName === 'string') {
            nameUser.textContent = nickName;
          } else {
            // Иначе, если это объект, обрабатываем соответствующим образом
            nameUser.textContent = nickName.String || "Неизвестный пользователь";
          }

          if (registrationDate) registrationDate.textContent = data.date_registration || "Дата регистрации неизвестна";
          if (descriptionUser) descriptionUser.textContent = data.description || "Нет описания";

          const btnEdit = document.querySelector(".btn-edit");
          if (btnEdit) {
            btnEdit.addEventListener("click", function () {
              window.location.href = "/changeProfile?userId=" + userId;
            });
          }
        })
        .catch((error) => {
          console.error("Ошибка загрузки данных о пользователе", error);
        });
  }
});
