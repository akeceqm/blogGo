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

        console.log(data); // Отладочная информация
        document.getElementById("name-user").textContent = data.nick_name;
        document.getElementById("registration-date").textContent =
          data.date_registration || "Дата регистрации неизвестна";
        document.getElementById("description-user").textContent =
          data.description || "Нет описания";
      })
      .catch((error) => {
        console.error("Ошибка загрузки данных о пользователе", error);
      });
    const btnEdit = document.querySelector(".btn-edit");
    btnEdit.addEventListener("click", function () {
      window.location.href = `/changeProfile?userId=`+data.id;
    });
  }
});
