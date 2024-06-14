document.addEventListener("DOMContentLoaded", function () {
  // Получаем параметры URL для получения id пользователя
  const urlParams = new URLSearchParams(window.location.search);
  const userId = urlParams.get("userId");

  if (userId) {
    // Отправляем запрос на сервер для получения данных пользователя
    fetch(`/api/user/${userId}`)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Не удалось загрузить данные пользователя");
        }
        return response.json();
      })
      .then((data) => {
        const imagePath = data.image_path;

        // Отобразить изображение на странице
        const avatarImg = document.querySelector(".avatar");
        avatarImg.src = imagePath;
        avatarImg.alt = "Avatar";

        // Если изображение не загружено, можно установить дефолтное изображение
        avatarImg.onerror = function () {
          avatarImg.src = "assets/img/avatar.svg";
          avatarImg.alt = "Default Avatar";
        };
        console.log(data); // Отладочная информация
        document.getElementById("name-user").textContent = data.nick_name;
        document.getElementById("registration-date").textContent =
          data.date_registration; // дата уже в нужном формате
        document.getElementById("description-user").textContent =
          data.description || "Нет описания"; // проверка на null и установка пустой строки, если null
        // Пример установки изображения аватара
        document
          .querySelector(".avatar")
          .setAttribute(
            "src",
            data.avatar_url || "/path/to/default/avatar.png"
          );
      })
      .catch((error) => {
        console.error("Ошибка загрузки данных о пользователе", error);
      });
  }
});
