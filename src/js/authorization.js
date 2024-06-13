document.addEventListener("DOMContentLoaded", function () {
  let form = document.getElementById("authorizationForm");

  if (form) {
    form.addEventListener("submit", function (e) {
      e.preventDefault();

      let loginInput = document.querySelector("input[name='login']");
      let passwordInput = document.querySelector("input[name='password_hash']");

      let loginError = document.getElementById("loginError");
      let passwordError = document.getElementById("passwordError");

      loginError.textContent = "";
      passwordError.textContent = "";

      if (loginInput.value.trim() === "") {
        loginError.textContent = "Введите логин";
        return; // Прерываем выполнение функции, если есть ошибка
      }

      if (passwordInput.value.trim() === "") {
        passwordError.textContent = "Введите пароль";
        return; // Прерываем выполнение функции, если есть ошибка
      }

      let data = {
        login: loginInput.value,
        password_hash: passwordInput.value,
      };

      let xhr = new XMLHttpRequest();

      xhr.open("POST", "/authorization", true);
      xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

      xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
          let response = JSON.parse(xhr.responseText);

          Swal.fire({
            title: "Успех!",
            text: "Успешная авторизация",
            icon: "success",
            background: "#1e1e1e",
            color: "#fff",
          });
          window.location.href = "/profileUser";
        } else {
          Swal.fire({
            icon: "error",
            title: "Ошибка!",
            text: "Неудачная авторизация",
            background: "#1e1e1e",
            color: "#fff",
          });
        }
      };

      xhr.onerror = function () {
        console.error("Request failed");
        Swal.fire({
          title: "Интернет?",
          text: "Проблемы с сервером либо интернетом",
          icon: "question",
          background: "#1e1e1e",
          color: "#fff",
        });
      };

      xhr.send(JSON.stringify(data));
    });
  }
});
